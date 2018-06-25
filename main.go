package main

import (
	"database/sql"
	"log"

	"./models"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

const (
	sqlite3Str = "sqlite3"
	memStr     = ":memory:"
)

func Openconnection()(*sql.DB){
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}

	return db
}

func CreateTable(db *sql.DB){
	log.Printf("Creating new table")
	if _, crErr := models.CreateIndexTable(db);
	crErr != nil {
		log.Fatalf("Error creating table (%s)", crErr)
	}
	log.Printf("Created")
}

func InsertTable(keyword string, bitmap string, db *sql.DB)(models.FileIndex){
	me := models.FileIndex{Keyword: keyword, Bitmap: bitmap}
	log.Printf("Inserting %+v into the DB", me)
	if _, insErr := models.InsertIndex(db, me);
	insErr != nil {
		log.Fatalf("Error inserting new index into the DB (%s)", insErr)
	}
	log.Printf("Inserted")

	return me
}

func SelectTable(db *sql.DB, me models.FileIndex){
	log.Printf("Selecting keyword from the DB")
	selectedMe := models.FileIndex{}
	if err := models.SelectIndex(db, me.Keyword, me.Bitmap, &selectedMe);
	err != nil {
		log.Fatalf("Error selecting person from the DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", selectedMe)
}

func UpdateTable(db *sql.DB){
	log.Printf("Updating person in the DB")
	updatedMe := models.FileIndex{
		Keyword: "Fon",
		Bitmap:  "001011000",
		// make this update after my add file!
	}
	//selectedMe := models.FileIndex{}
	if err := models.UpdateIndex(db, updatedMe.Keyword, updatedMe.Bitmap, updatedMe);
	err != nil {
		log.Fatalf("Error updating person in the DB (%s)", err)
	}
}

func DelTable(db *sql.DB){
	selectedMe := models.FileIndex{}
	log.Printf("Deleting person from the DB")
	if delErr := models.DeleteIndex(db, selectedMe.Keyword, selectedMe.Bitmap); delErr != nil {
		log.Fatalf("Error deleting person from the DB (%s)", delErr)
	}
	log.Printf("Deleted")
}

func main() {
	// change this line to use a different database and connection string to connect to a different database
	db := Openconnection()

	CreateTable(db)

	me := InsertTable("Fon","0010110",db)

	SelectTable(db,me)

	UpdateTable(db)

	SelectTable(db,me)

	DelTable(db)
}
