package health

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	errorsMap = make(map[int]*callbackError)
}

type registeredCallback struct {
	isHardDependency bool
	description      string
	callback         func() error
}

type callbackError struct {
	IsHardDependency bool   `json:"isHardDependency"`
	Description      string `json:"description"`
	Error            string `json:"error"`
}

var mutex sync.Mutex
var nextIndex int
var errorsMap map[int]*callbackError

type response struct {
	Errors    []callbackError `json:"errors,omitempty"`
	IsHealthy bool            `json:"isHealthy"`
}

func RegisterHealthCheck(hc HealthCheck) {
	go registerCallback(hc.HealthCheck, hc.IsHardDependency(), hc.Description())
}

func registerCallback(callback func() error, isHardDependency bool, description string) {

	mutex.Lock()
	index := nextIndex
	nextIndex++
	mutex.Unlock()

	errChan := make(chan error)
	for {

		go func() {
			errChan <- callback()
		}()

		time.Sleep(time.Second)

		select {
		case err := <-errChan:
			mutex.Lock()
			if err != nil {
				errorsMap[index] = &callbackError{
					IsHardDependency: isHardDependency,
					Description:      description,
					Error:            err.Error(),
				}
			} else {
				errorsMap[index] = nil
			}
			mutex.Unlock()

		case <-time.After(time.Second * 5):

			mutex.Lock()
			errorsMap[index] = &callbackError{
				IsHardDependency: isHardDependency,
				Description:      description,
				Error:            "Check Callback timed out",
			}
			mutex.Unlock()
			fmt.Println(errorsMap[index])
			<-errChan
		}
	}
}

func HealthHandler() http.Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	errs := []*callbackError{}
	mutex.Lock()
	for _, err := range errorsMap {
		if err == nil {
			continue
		}
		errs = append(errs, err)
	}
	mutex.Unlock()

	var resp response
	isHealthy := true
	for _, err := range errs {
		resp.Errors = append(resp.Errors, *err)
		if err.IsHardDependency {
			isHealthy = false
		}
	}
	resp.IsHealthy = isHealthy

	data, _ := json.MarshalIndent(resp, "", "  ")
	w.Write(data)

	return
}

func Init(listenAddress string) error {
	http.Handle("/debug/metrics", promhttp.Handler())
	http.Handle("/debug/health", HealthHandler())
	go func() {
		err := http.ListenAndServe(listenAddress, nil)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
