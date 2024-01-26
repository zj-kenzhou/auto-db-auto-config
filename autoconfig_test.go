package autoconfig

import (
	"github.com/zj-kenzhou/auto-db/datasource"
	"gorm.io/gorm"
	"testing"
)

func TestDatasourceAutoConfig(t *testing.T) {
	db := datasource.Db()
	db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		_, err := tx.Table("test").Rows()
		if err != nil {
			t.Error(err)
		}
		return tx
	})
}
