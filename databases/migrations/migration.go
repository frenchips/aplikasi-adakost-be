package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DB *sql.DB

	//go:embed sql_migrations/*.sql
	dbMigrations embed.FS
)

func DBMigrate(dbParam *sql.DB) {

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DB = dbParam

	fmt.Println("Migration success, applied", n, "migrations!")

}
