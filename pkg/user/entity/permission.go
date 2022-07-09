package entity

type Permission struct {
	UserID     int64
	Department Department
	Role       Role

	DataScope       DataScope // 数据权限
	PermissionWhole int       // 数据权限auth 是否全部权限 0-否 1-是
	PermissionAuth  []int64
}
