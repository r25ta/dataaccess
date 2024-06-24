package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"dataaccess.com/constant"
	model "dataaccess.com/model"
)

func main() {
	var conDb *sql.DB

	//Capture connection properties
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", constant.USER, constant.PWD, constant.SERVER, constant.PORT, constant.DATABASE)
	conDb, conErr := sql.Open("postgres", conStr)

	if conErr != nil {
		log.Fatal("Error connecting to the database", conErr)
	}

	defer conDb.Close()

	pingErr := conDb.Ping()

	if pingErr != nil {
		log.Fatal("Error ")
	}

	fmt.Println("Connected in database!")

	albums, error := albumsByArtists("John Coltrane", conDb)

	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Albums found %v\n", albums)

	// Hard-code ID 1 and 10 here to test the query.
	var album1 model.Album

	album1, error = albumById(1, conDb)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("Album found: %v\n", album1)

	/*
		album10, error = albumById(10, conDb)
		if error != nil {
			log.Fatal(error)
		}

		fmt.Printf("Album found: %v\n", album10)
	*/

	alb := model.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}

	idNewAlbum, error := addNewAlbum(alb, conDb)

	if error != nil {
		log.Fatal(error)

	}

	fmt.Printf("ID of added album: %v\n", idNewAlbum)

}
func albumsByArtists(name string, db *sql.DB) ([]model.Album, error) {

	/*Declare an albums slice of the Album type you defined.
	This will hold data from returned rows.
	Struct field names and types correspond to database column names and types.
	*/
	var albums []model.Album

	/*Use conDB.Query to execute a SELECT statement to query for albums with the specified artist name.
	Query’s first parameter is the SQL statement.
	After the parameter, you can pass zero or more parameters of any type.
	These provide a place for you to specify the values for parameters in your SQL statement.
	By separating the SQL statement from parameter values (rather than concatenating them with, say, fmt.Sprintf),
	you enable the database/sql package to send the values separate from the SQL text,
	removing any SQL injection risk.
	*/
	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtists %q: %v", name, err)

	}
	/*Defer closing rows so that any resources it holds will be released when the function exits.
	 */
	defer rows.Close()

	//Loop through the returned rows, using Rows.Scan to assign each row’s column values to Album struct fields.
	for rows.Next() {
		var alb model.Album

		/*Scan takes a list of pointers to Go values, where the column values will be written.
		Here, you pass pointers to fields in the alb variable, created using the & operator.
		Scan writes through the pointers to update the struct fields.
		*/
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

		//Inside the loop, check for an error from scanning column values into the struct fields.
		if err != nil {
			return nil, fmt.Errorf("albumsByArtists %q: %v", name, err)
		}

		//Inside the loop, append the new alb to the albums slice.
		albums = append(albums, alb)
	}

	/*After the loop, check for an error from the overall query, using rows.Err.
	Note that if the query itself fails,
	checking for an error here is the only way to find out that the results are incomplete.
	*/
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtists %q: %v", name, err)
	}

	return albums, nil
}

func albumById(id int64, db *sql.DB) (model.Album, error) {
	var alb model.Album
	/*Use DB.QueryRow to execute a SELECT statement to query for an album with the specified ID.
	It returns an sql.Row. To simplify the calling code (your code!),
	QueryRow doesn’t return an error. Instead,
	it arranges to return any query error (such as sql.ErrNoRows) from Rows.Scan later.
	*/
	row := db.QueryRow("SELECT *FROM album WHERE id = $1", id)

	/*Use Row.Scan to copy column values into struct fields.*/
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	/*Check for an error from Scan.
	The special error sql.ErrNoRows indicates that the query returned no rows.
	Typically that error is worth replacing with more specific text, such as “no such album” here.
	*/
	if err == sql.ErrNoRows {
		return alb, fmt.Errorf("albumsById %d: no such album", id)

	} else if err != nil {
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	} else {
		return alb, nil

	}

}

func addNewAlbum(album model.Album, db *sql.DB) (int64, error) {
	/*Use DB.Exec to execute an INSERT statement.
	Like Query, Exec takes an SQL statement followed by parameter values for the SQL statement.
	*/
	var lastInsertId int64
	row := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", album.Title, album.Artist, album.Price)

	err := row.Scan(&lastInsertId)

	//*Check for an error from the attempt to INSERT.
	//Check for an error from the attempt to retrieve the ID
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("addNewAlbum: %v", err)

	} else if err != nil {
		return 0, fmt.Errorf("addNewAlbum: %v", err)

	} else {
		return lastInsertId, nil

	}

}
