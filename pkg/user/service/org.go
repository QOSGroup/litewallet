package service

import (
	"github.com/QOSGroup/litewallet/pkg/user/entity"
	"github.com/sirupsen/logrus"
)

var _ entity.OrgService = (*PermService)(nil)

type PermService struct {
	permRepo entity.OrgRepo
	log      *logrus.Entry
}

func NewOrgService(permRepo entity.OrgRepo, log *logrus.Entry) *PermService {
	return &PermService{
		permRepo: permRepo,
		log:      log.WithField("module", "domain.OrgService"),
	}
}

//
//func (s *OrgService) GetCompanyByID(ctx context.Context, id int64) (*entity.Company, error) {
//	s.log.Infof(ctx, "GetCompanyByID input=%v", id)
//	return s.permRepo.GetCompanyByID(ctx, id)
//}
//
//func (s *OrgService) GetDepartmentByID(ctx context.Context, id int64) (*entity.Department, error) {
//	s.log.Infof(ctx, "GetDepartmentByID input=%v", id)
//	return s.permRepo.GetDepartmentByID(ctx, id)
//}
//
//func (s *OrgService) GetRoleByID(ctx context.Context, id int64) (*entity.Role, error) {
//	s.log.Infof(ctx, "GetRoleByID input=%v", id)
//	return s.permRepo.GetRoleByID(ctx, id)
//}
