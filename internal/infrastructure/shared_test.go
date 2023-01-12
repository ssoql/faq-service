package infrastructure

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-service/internal/infrastructure/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type MockedDatabaseSuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *faqReadRepository
}

func (r *MockedDatabaseSuite) SetupSuite() {
	var err error

	r.conn, r.mock, err = sqlmock.New()
	require.NoError(r.T(), err)

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.1.1")

	r.mock.ExpectQuery(`SELECT VERSION()`).WillReturnRows(rows)

	dialector := mysql.New(mysql.Config{Conn: r.conn, DriverName: "mysql", DSN: "sql_mock_db0"})

	r.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(r.T(), err)

	r.repo = &faqReadRepository{db: &db.ClientDB{DB: r.DB}}

}
func (r *MockedDatabaseSuite) TestGet() {
	rows := sqlmock.NewRows([]string{"id", "name", "age"}). // adiciona o nome das colunas
								AddRow( // adiciona o primeiro registro
			int64(1),
			"xxxx",
			"xxxx")
	//r.mock.ExpectBegin()
	//r.mock.ExpectQuery("SELECT * FROM `faqs` WHERE id = ? ORDER BY `faqs`.`id` LIMIT 1").
	r.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`faqs`" + ` WHERE id = ? ORDER BY ` + "`faqs`.`id`" + ` LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows) // valida os registros retornados
	_, err := r.repo.GetByID(context.Background(), int64(1)) // chama o método Find do repository
	assert.NoError(r.T(), err)                               // valida se houve algum erro
	//assert.Contains(r.T(), people, *)
}

func TestNewReadRepository(t *testing.T) {
	suite.Run(t, new(MockedDatabaseSuite))
	//ts := &MockedDatabaseSuite{}
	//ts.SetupSuite()
	////ts.TestGet()
}
