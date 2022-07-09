package repo

import "errors"

var (
	// ErrCreateZeroRowsAffected 插入0行数据
	ErrCreateZeroRowsAffected = errors.New("ErrCreateZeroRowsAffected")
	// ErrUpdateZeroRowsAffected 更新0行数据
	ErrUpdateZeroRowsAffected = errors.New("ErrUpdateZeroRowsAffected")
	// ErrDeleteZeroRowsAffected 删除0行数据
	ErrDeleteZeroRowsAffected = errors.New("ErrDeleteZeroRowsAffected")
	// ErrObjectNotExist 对象不存在
	ErrObjectNotExist = errors.New("ErrObjectNotExist")
)
