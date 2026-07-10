package app

import (
	"database/sql"
	"golang-database-migration/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate -database "mysql://root:password@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:password@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state

	// karena error saat membuat table wrong, maka table correct harus di delete secara manual : DROP TABLE correct;
	// dan kita harus pindahkan versi di schema_migrations ke versi 20260710153215
	// dengan perintah : migrate -database "mysql://root:password@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20260710153215
	// perintah melihat versi saat ini : migrate -database "mysql://root:password@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
}
