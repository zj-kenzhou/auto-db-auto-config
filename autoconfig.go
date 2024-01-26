package autoconfig

import (
	"fmt"
	"github.com/zj-kenzhou/auto-db/datasource"
	"github.com/zj-kenzhou/fast-config"
	"log"
)

func loadConfig() map[string]datasource.Config {
	configAny := fastconfig.GetValue("datasource", map[string]datasource.Config{})
	configMap, ok := configAny.(map[string]any)
	if !ok {
		return map[string]datasource.Config{}
	}
	resMap := make(map[string]datasource.Config, len(configMap))
	for key := range configMap {
		config := datasource.Config{
			Type:            fastconfig.GetString(fmt.Sprintf("datasource.%s.type", key), ""),
			Host:            fastconfig.GetString(fmt.Sprintf("datasource.%s.host", key), ""),
			Username:        fastconfig.GetString(fmt.Sprintf("datasource.%s.username", key), ""),
			Password:        fastconfig.GetString(fmt.Sprintf("datasource.%s.password", key), ""),
			LogLevel:        fastconfig.GetInt(fmt.Sprintf("datasource.%s.log-level", key), 3),
			Dbname:          fastconfig.GetString(fmt.Sprintf("datasource.%s.dbname", key), ""),
			MaxIdleConns:    fastconfig.GetInt(fmt.Sprintf("datasource.%s.max-idle-conns", key), 5),
			MaxOpenConns:    fastconfig.GetInt(fmt.Sprintf("datasource.%s.max-open-conns", key), 10),
			ConnMaxLifetime: fastconfig.GetInt(fmt.Sprintf("datasource.%s.conn-max-lifetime", key), 180),
			ConnMaxIdleTime: fastconfig.GetInt(fmt.Sprintf("datasource.%s.conn-max-idle-time", key), 30),
			AutoMigrate:     fastconfig.GetBool(fmt.Sprintf("datasource.%s.auto-migrate", key), false),
		}
		resMap[key] = config
	}
	return resMap
}
func init() {
	configMap := loadConfig()
	err := datasource.InitDataSource(configMap)
	if err != nil {
		log.Println(err)
	}
}
