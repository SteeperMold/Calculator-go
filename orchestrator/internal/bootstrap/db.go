package bootstrap

import (
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewSqlDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "data/sqlite.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	const q = `
		CREATE TABLE IF NOT EXISTS statuses (
		    id   INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT    NOT NULL UNIQUE
		);
		INSERT OR IGNORE INTO statuses(name)
	   	VALUES (?), (?), (?);

		CREATE TABLE IF NOT EXISTS expressions (
			id        INTEGER PRIMARY KEY AUTOINCREMENT,
			status_id INTEGER NOT NULL,
			result    REAL,
		   	FOREIGN KEY(status_id) REFERENCES statuses(id)
		);

		CREATE TABLE IF NOT EXISTS nodes (
			id            INTEGER PRIMARY KEY AUTOINCREMENT,
			expression_id INTEGER NOT NULL,
			parent_id     INTEGER REFERENCES nodes(node_id),
			status_id     INTEGER NOT NULL,
			value         TEXT    NOT NULL,
			left_id       INTEGER REFERENCES nodes(node_id),
			right_id      INTEGER REFERENCES nodes(node_id),
		    FOREIGN KEY(expression_id) REFERENCES expressions(id),
		    FOREIGN KEY(status_id)     REFERENCES statuses(id)
		);
		CREATE INDEX IF NOT EXISTS idx_nodes_expr ON nodes(expression_id);

		CREATE TABLE IF NOT EXISTS users (
		    id            INTEGER      PRIMARY KEY AUTOINCREMENT,
		    login         VARCHAR(32)  NOT NULL UNIQUE,
		    password_hash VARCHAR(255) NOT NULL,
		    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err = db.Exec(q, domain.StatusInProgress, domain.StatusGivenToAgent, domain.StatusFinished)
	if err != nil {
		log.Fatalf("failed to create tables in database: %v", err)
	}

	return db
}

func CloseDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("failed to close db: %v", err)
	}
	log.Println("database closed successfully")
}
