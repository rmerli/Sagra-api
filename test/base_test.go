package test

import (
	"context"
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/migration"
	"gtmx/src/server"
	"os"
	"syscall"
	"testing"

	"github.com/antonlindstrom/pgstore"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	conn    *pgx.Conn
	db      *database.Queries
	store   *pgstore.PGStore
	context context.Context
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestSuite{})
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.NoError(godotenv.Load("../.env-test"))

	ctx := context.Background()
	s.context = ctx

	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	s.NoError(err)

	s.conn = conn

	store, err := pgstore.NewPGStore(os.Getenv("DB_URL"), []byte(os.Getenv("STORE_KEY")))
	s.NoError(err)
	s.store = store

	s.db = database.New(conn)

	serverReady := make(chan bool)
	server := server.New(s.db, store, serverReady)

	go server.Start(fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT")))
	<-serverReady
}

func (s *IntegrationTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *IntegrationTestSuite) SetupTest() {
	migration := migration.Migration{
		Db: s.conn,
	}
	migration.MigrateFresh()
}

func (s *IntegrationTestSuite) TearDownTest() {
}
