//https://ericchiang.github.io/post/testing-dbs-with-docker/

package database

import (
	"database/sql"
	config "github.com/curtischong/lizzie_server/config"
	_ "github.com/lib/pq"
	"testing"
)

func createTable(conn *sql.DB) error {
	statement := `
		CREATE TABLE Worker (
			WorkerId SERIAL PRIMARY KEY,
			Host     TEXT NOT NULL,
			UsingTLS BOOLEAN NOT NULL
		);`
	_, err := conn.Exec(statement)
	return err
}

// Test creating a table
func testCreateTable(t *testing.T, conn *sql.DB) {
	if err := createTable(conn); err != nil {
		t.Error(err)
	}
}

// Test creating a table then inserting a row
func testInsertWorker(t *testing.T, conn *sql.DB) {
	if err := createTable(conn); err != nil {
		t.Error(err)
		return
	}
	insertWorker := `INSERT INTO Worker (Host, UsingTLS) VALUES ($1, $2);`
	if _, err := conn.Exec(insertWorker, "10.0.0.3", true); err != nil {
		t.Error(err)
	}
}

// Test multiple versions of Postgres
const (
	version11_3 = "11.3"
)

// The actual Go tests just immediately call RunDBTest
func TestCreateTable(t *testing.T) {
	config := config.LoadConfiguration("../configSecrets/server_config.json")
	RunDBTest(t, version11_3, testCreateTable, config)
}
func TestInsertWorker(t *testing.T) {
	RunDBTest(t, version11_3, testInsertWorker, config)
}
