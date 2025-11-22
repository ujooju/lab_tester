package storage

import (
	"database/sql"
	"log"

	"github.com/ujooju/lab_tester/webInterface/models"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitSQLite() error {
	log.Println("Initializing database...")
	var err error
	DB, err = sql.Open("sqlite", "./lab_tester.db")
	if err != nil {
		return err
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS test_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		owner text NOT NULL,
		name text NOT NULL,
		branch text NOT NULL,
		status text NOT NULL,
		report text NOT NULL
		);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		return err
	}
	log.Println("Database initialized")
	return nil
}

func GetTestsByOwnerAndName(owner string, taskName string) ([]models.TestRecord, error) {
	result := []models.TestRecord{}
	rows, err := DB.Query("SELECT id, owner, name, branch, status, report FROM test_records WHERE owner = ? AND name = ?", owner, taskName)
	if err != nil {
		return []models.TestRecord{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.TestRecord
		err := rows.Scan(&record.ID, &record.Owner, &record.RepoName, &record.Branch, &record.Status, &record.Report)
		if err != nil {
			return []models.TestRecord{}, err
		}
		result = append(result, record)
	}
	return result, nil
}

func SubmutTest(owner string, name string, branch string) error {
	submitQuery := `INSERT INTO test_records (owner, name, branch, status, report) VALUES (?, ?, ?, ?, ?);`
	_, err := DB.Exec(submitQuery, owner, name, branch, "submited", "")
	if err != nil {
		return err
	}
	return nil
}

func NextTest() (models.TestRecord, error) {
	nextTest := models.TestRecord{}
	query := `SELECT id, owner, name, branch, status, report FROM test_records WHERE status = "submited" ORDER BY id ASC`
	row := DB.QueryRow(query)
	err := row.Scan(&nextTest.ID, &nextTest.Owner, &nextTest.RepoName, &nextTest.Branch, &nextTest.Status, &nextTest.Report)
	if err == sql.ErrNoRows {
		return models.TestRecord{}, nil
	}
	if err != nil {
		log.Println(err.Error())
		return models.TestRecord{}, err
	}
	return nextTest, nil
}

func UpdateRecord(record *models.TestRecord) error {
	q := `UPDATE test_records SET status = ?, report = ? WHERE id = ?`
	_, err := DB.Exec(q, record.Status, record.Report, record.ID)
	if err != nil {
		return err
	}
	return nil
}
