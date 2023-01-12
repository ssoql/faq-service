package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockClientDb() (sqlmock.Sqlmock, *db.ClientDB, error) {

	conn, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.1.1")

	mock.ExpectQuery(`SELECT VERSION()`).WillReturnRows(rows)

	dialector := mysql.New(mysql.Config{Conn: conn, DriverName: "mysql", DSN: "sql_mock_db0"})

	dbClient, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return mock, &db.ClientDB{DB: dbClient}, nil
}
