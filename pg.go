package pg_client

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var connect *gorm.DB
var once sync.Once

type Config struct {
	Port     string `env:"DB_SERVER_DB_PORT" env-default:"5432"`
	Host     string `env:"DB_SERVER_HOST" env-default:"localhost"`
	Name     string `env:"DB_SERVER_DB_NAME" env-default:"postgres"`
	User     string `env:"DB_SERVER_USER_NAME" env-default:"user"`
	Password string `env:"DB_SERVER_USER_PASS" env-default:"password"`
	SqlDebug bool   `env:"SQL_DEBUG" env-default:"false"`
}

func GetInstance(cfg Config) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		connect, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.User, cfg.Name, cfg.Password)), &gorm.Config{})
		if cfg.SqlDebug {
			connect = connect.Debug()
		}
	})
	return connect, err
}
