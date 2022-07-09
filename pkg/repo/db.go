package repo

import (
	"context"
	"database/sql"
)

type DB interface {
	DB() (*sql.DB, error)

	AutoMigrate(dst ...interface{}) (err error)
	WithContext(ctx context.Context) DB

	// Exec execute raw sql
	Exec(sql string, values ...interface{}) (rowsAffected int64, err error)

	First(dest interface{}, conds ...interface{}) (err error)
	Create(value interface{}) (rowsAffected int64, err error)
	Update(value interface{}) (rowsAffected int64, err error)
	Delete(value interface{}) (rowsAffected int64, err error)
	Raw(dest interface{}, sql string, values ...interface{}) (rowsAffected int64, err error)
	Transaction(fc func(tx DB) error, opts ...*sql.TxOptions) error

	Where(query interface{}, args ...interface{}) (tx DB)
	Find(dest interface{}, conds ...interface{}) (err error)
	Order(value interface{}) (tx DB)
	Limit(limit int) (tx DB)
	Offset(offset int) (tx DB)

	Close() error
}
