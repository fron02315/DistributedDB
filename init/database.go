package init

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

type Kindex struct {
	Keyword string
	Bitmap string
}

func Openconnection()(*sql.DB){
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
  //sql.Open("sql platform name", "username:password@tcp(serverIP:Port)/database")
	db, err := sql.Open("mysql", "root:00000000@tcp(127.0.0.1:3306)/project")

	// if there is an error opening the connection, handle it
	if err != nil {
			panic(err.Error())
	}
	fmt.Println("Successfully connected")
	return db
}

func Search(db *sql.DB, target string){

	// Execute the query
	results, err := db.Query("SELECT Bitmap FROM test WHERE Keyword = ?", target)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var kindex Kindex
		// for each row, scan the result into our tag composite object
		err = results.Scan(&kindex.Bitmap)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	   // and then print out the tag's Name attribute
		fmt.Println(kindex.Bitmap)
	}
}

func InsertKey(db *sql.DB, key string, bit string){
	//test for insert multiple variable
	// insert
	stmt, err := db.Prepare("INSERT test SET Keyword=?, Bitmap=?")
	checkErr(err)

	res, err := stmt.Exec(key, bit)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func UpdateTable(db *sql.DB, key string, bit string){
  stmt, err := db.Prepare("update test set Bitmap=? where Keyword=?")
  checkErr(err)

  res, err := stmt.Exec(bit, key)
  checkErr(err)

  affect, err := res.RowsAffected()
  checkErr(err)

  fmt.Println(affect)
}

func main() {
    /*fmt.Println("Go MySQL Tutorial")

    db := Openconnection()

		//Insert the query in form of (db,Keyword,Bitmap)
		InsertKey(db, "Hell", "11111100")

		//search the target keyword in form of (db,Keyword)
		//and retuern the Bitmap
		Search(db, "Hell")

		// defer the close till after the main function has finished
		// executing at the end of the query
		defer db.Close() */

}
