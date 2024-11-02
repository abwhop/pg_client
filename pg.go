package pg_client

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once

type Config struct {
	Port     string `env:"DB_SERVER_DB_PORT" env-default:"5432"`
	Host     string `env:"DB_SERVER_HOST" env-default:"localhost"`
	Name     string `env:"DB_SERVER_DB_NAME" env-default:"postgres"`
	User     string `env:"DB_SERVER_USER_NAME" env-default:"user"`
	Password string `env:"DB_SERVER_USER_PASS" env-default:"password"`
	SqlDebug bool   `env:"SQL_DEBUG" env-default:"false"`
}

type PgClient struct {
	db *gorm.DB
}

func GetInstance(cfg Config) (*PgClient, error) {
	var err error
	var connection *gorm.DB
	once.Do(func() {
		connection, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.User, cfg.Name, cfg.Password)), &gorm.Config{})
		if cfg.SqlDebug {
			connection = connection.Debug()
		}
	})
	if err != nil {
		return nil, err
	}
	return &PgClient{db: connection}, err
}
