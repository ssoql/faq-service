package db

import (
	"fmt"
	"github.com/ssoql/faq-service/internal/app/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBFactory interface {
	getDbClient() (*gorm.DB, error)
}

type dbFactory struct {
	connStr string
}

func NewDBFactory(connStr string) *dbFactory {
	return &dbFactory{connStr: connStr}
}

type ClientDB struct {
	*gorm.DB
	options DatabaseOptions
}

type DatabaseOptions struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSchema   string
	DatabasePort     string
	DatabaseHost     string
}

func (r *dbFactory) GetClient() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(r.connStr), &gorm.Config{})
}

func InitializeDB(cfg DatabaseOptions) (*ClientDB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseSchema,
	)
	dbClient, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err := dbClient.AutoMigrate(&entities.Faq{}); err != nil {
		return nil, err
	}

	return &ClientDB{dbClient, cfg}, nil
}
