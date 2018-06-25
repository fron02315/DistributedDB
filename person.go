package models

import (
	"database/sql"
	"fmt"
)

const (
	// FileIndexTable is the name of the table for the Fileindex model
	FileIndexTable = "FileIndex"
	// Keyword is the column name of the model's keyword
	Keywordcol = "keyword"
	// PersonLastNameCol is the column name of the model's last name
	Bitmapcol = "file_bitmap"
)

// Person is the database model for a person
type FileIndex struct {
	Keyword string
	Bitmap  string
}

// CreateTable uses db to create a new table for Person models, and returns the result
func CreateIndexTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255), %s varchar(255))",
			FileIndexTable,
			Keywordcol,
			Bitmapcol,
		),
	)
}

// InsertPerson inserts person into db
func InsertIndex(db *sql.DB, fileIndex FileIndex) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES(?, ?)", FileIndexTable),
		fileIndex.Keyword,
		fileIndex.Bitmap,
	)
}

// SelectPerson selects a person with the given first & last names and age. On success, writes the result into result and on failure, returns a non-nil error and makes no modifications to result
func SelectIndex(db *sql.DB, keyword, file_bitmap string, result *FileIndex) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=? AND %s=?",
			FileIndexTable,
			Keywordcol,
			Bitmapcol,
		),
		keyword,
		file_bitmap,
	)
	var retKeyword, retBitmap string
	if err := row.Scan(&retKeyword, &retBitmap); err != nil {
		return err
	}
	result.Keyword = retKeyword
	result.Bitmap = retBitmap
	return nil
}

// UpdateIndex updates the index with the given keyword & bitmap. Returns a non-nil error if the update failed, and nil if the update succeeded
func UpdateIndex(db *sql.DB, keyword, file_bitmap string, newFileIndex FileIndex) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=?,%s=? WHERE %s=? AND %s=?",
			FileIndexTable,
			Keywordcol,
			Bitmapcol,
			Keywordcol,
			Bitmapcol,
		),
		newFileIndex.Keyword,
		newFileIndex.Bitmap,
		keyword,
		file_bitmap,
	)
	return err
}

// DeleteIndex deletes the keyword with the given keyword and bitmap. Returns a non-nil error if the delete failed, and nil if the delete succeeded
func DeleteIndex(db *sql.DB, keyword, file_bitmap string) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=? AND %s=?",
			FileIndexTable,
			Keywordcol,
			Bitmapcol,
		),
		keyword,
		file_bitmap,
	)
	return err
}
