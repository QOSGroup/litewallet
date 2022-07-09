package entity

type DataScope int

func (ds DataScope) Int() int {
	return int(ds)
}

func (ds DataScope) String() string {
	switch ds {
	case DataScopeAll:
		return "全部数据权限"
	case DataScopeCustomize:
		return "自定数据权限"
	case DataScopeDepartment:
		return "本部门数据权限"
	case DataScopeDepartmentAndChild:
		return "本部门及以下数据权限"
	case DataScopePersonal:
		return "仅本人数据权限"
	default:
		return ""
	}
}

const (
	DataScopeAll                DataScope = iota + 1 // 全部数据权限
	DataScopeCustomize                               // 自定数据权限
	DataScopeDepartment                              // 本部门数据权限
	DataScopeDepartmentAndChild                      // 本部门及以下数据权限
	DataScopePersonal                                // 仅本人数据权限
)
