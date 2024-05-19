package database

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

type Migration struct {
	Db *pgx.Conn
}

func (m *Migration) MigrateFresh() {
	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations/postgres",
	}
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			m.Db.Config().User,
			m.Db.Config().Password,
			m.Db.Config().Host,
			m.Db.Config().Database,
		))
	if err != nil {
	}
	defer db.Close()

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		logrus.Error(err)
	}
	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("Migrated fresh")
}
