package autoconfig

import (
	"github.com/zj-kenzhou/auto-db/datasource"
	"gorm.io/gorm"
	"testing"
)

func TestDatasourceAutoConfig(t *testing.T) {
	db := datasource.Db()
	db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx.Table("test").Rows()
		return tx
	})
}
