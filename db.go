package quotes

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Quote struct {
	ID int
	Text string
	Author string
}

type QDB struct {
	db *sql.DB
}

func New(dbcreds string) (*QDB, error) {
	db, err := sql.Open("mysql", dbcreds)
	if err != nil {
		return nil, err
	}

	qdb := &QDB{db: db}

	return qdb, nil
}

func (qdb *QDB) Random() (*Quote, error) {
	var q Quote
	row := qdb.db.QueryRow("SELECT id, text, author FROM quotes ORDER BY RAND() LIMIT 1")

	if err := row.Scan(&q.ID, &q.Text, &q.Author); err != nil {
		return nil, err
	}

	return &q, nil
}

// add adds a new quote to the database.
func (qdb *QDB) Add(q *Quote) error {
	// INSERT INTO quotes VALUES..
	_, err := qdb.db.Exec("INSERT INTO quotes (text, author) VALUES (?, ?)", q.Text, q.Author)
	return err
}
