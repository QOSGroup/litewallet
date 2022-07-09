package service

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/QOSGroup/litewallet/pkg/user/entity"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/v2/model"
)

var _ entity.PermService = (*CasbinService)(nil)
var once sync.Once

type CasbinService struct {
	permRepo entity.OrgRepo
	enforcer *casbin.SyncedEnforcer
	log      *logrus.Entry
}

func NewCasbinService(permRepo entity.OrgRepo, log *logrus.Entry) *CasbinService {
	s := &CasbinService{
		permRepo: permRepo,
		log:      log.WithField("module", "domain.CasbinService"),
	}

	once.Do(func() {
		//a, _ := gormadapter.NewAdapterByDB()
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		enforcer := casbin.NewSyncedEnforcer(m)
		_ = enforcer.LoadPolicy()
		enforcer.StartAutoLoadPolicy(time.Second * 60) // 60秒更新一次
		s.enforcer = enforcer
	})

	return s
}

func (s CasbinService) Enforce(ctx context.Context, sub, obj, act string) bool {
	return s.enforcer.Enforce(sub, obj, act)
}
