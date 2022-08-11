package metrics

import (
	"fmt"
	"testing"

	"time"
)

// 测试 Timer的Time() 函数使用
func TestTimer(t *testing.T) {

	// new 一个timer实例并且注册到promtheus，
	timer := NewTimer("gmu", "http_request", "help", []string{"http_code", "url"})

	// 处理http 请求
	mockHttpHandleTime(timer, t)
	mockHttpHandleTime(timer, t)
}

// 测试Timer的Observe() 函数使用
func TestTimerObserve(t *testing.T) {
	// new 一个timer实例并且注册到promtheus，
	timer := NewTimer("gmu", "http_request", "help", []string{"http_code", "url"})
	// 处理http 请求
	mockHttpHandleObserve(timer, t)
	mockHttpHandleObserve(timer, t)
}

func TestTimerBucket(t *testing.T) {
	// new 一个timer实例并且注册到promtheus，
	timer := NewTimer("gmu", "http_request", "help", []string{"http_code", "url"}, WithTimerBuckets([]float64{0.1, 1, 5, 10, 50, 100}))
	// 处理http 请求
	timer.Observe(50*time.Second, "200", "http://gmu.com/users")
	data, err := GetPromethuesAsFmtText()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}

// 模拟一个http handle 函数
func mockHttpHandleTime(timer *Timer, t *testing.T) {
	// 开始计时
	tf := timer.Timer()

	// 模拟处理请求的时间
	time.Sleep(100 * time.Millisecond)

	// 结束计时，
	tf("200", "http://gmu.com/users")

	// 在终端打印出promtheus指标
	// 将看见前缀为 : gmu_http_request_h_*  和 gmu_http_request_s* 的指标
	data, err := GetPromethuesAsFmtText()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	t.Logf("%s", data)
}

// 模拟一个http handle 函数
func mockHttpHandleObserve(timer *Timer, t *testing.T) {
	// 开始计时
	startTime := time.Now()

	// 模拟处理请求的时间
	time.Sleep(100 * time.Millisecond)

	// 结束计时， time.Now().Sub(startTime)
	timer.Observe(time.Now().Sub(startTime), "200", "http://gmu.com/users")

	// 在终端打印出promtheus指标
	// 将看见前缀为 : gmu_http_request_h_*  和 gmu_http_request_s* 的指标
	data, err := GetPromethuesAsFmtText()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	t.Logf("%s", data)
}
