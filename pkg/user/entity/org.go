package entity

import (
	"context"
	"time"
)

type Company struct {
	ID          int64
	Name        string
	CompanyType int
	Address     string
	Contact     string
	Phone       string
	CityID      int64
	UpdateTime  time.Time
	CreateTime  time.Time
	DeletedAt   time.Time
}

type Department struct {
	ID         int64
	Name       string
	Path       string
	Sort       int
	ParentID   int64
	CompanyID  int64
	IsFixed    int
	UpdateTime time.Time
	CreateTime time.Time
	DeletedAt  time.Time
}

type Role struct {
	ID        int64
	Name      string
	Remark    string
	Router    string
	RoleID    int64
	CompanyID int64

	/* DataScope
	{value: '1',label: '全部数据权限'},
	{value: '2',label: '自定数据权限'},
	{value: '3',label: '本部门数据权限'},
	{value: '4',label: '本部门及以下数据权限'},
	{value: '5',label: '仅本人数据权限'}
	*/
	DataScope  DataScope
	UpdateTime time.Time
	CreateTime time.Time
	DeletedAt  time.Time
}

type OrgRepo interface {
	SaveCompany(ctx context.Context, e *Company) error
	GetCompanyByID(ctx context.Context, id int64) (*Company, error)

	SaveDepartment(ctx context.Context, e *Department) error
	GetDepartmentByID(ctx context.Context, id int64) (*Department, error)

	SaveRole(ctx context.Context, e *Role) error
	GetRoleByID(ctx context.Context, id int64) (*Role, error)

	//GetPermissionByUserID(ctx context.Context, uid int64) (*Permission, error)
}

type OrgService interface {
	//GetCompanyByID(ctx context.Context, id int64) (*Company, error)
	//GetDepartmentByID(ctx context.Context, id int64) (*Department, error)
	//GetRoleByID(ctx context.Context, id int64) (*Role, error)
}

type CasbinRepo interface {
}

// PermService 权限
type PermService interface {
	Enforce(ctx context.Context, sub, obj, act string) bool
}
