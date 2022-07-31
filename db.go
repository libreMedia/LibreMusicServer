package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func removeDb(dbName string) {
	os.Remove("./" + dbName + ".db")
}

type DbModel struct {
	Name       string
	RoutePath  string
	Path       string
	Title      string
	Artist     string
	Album      string
	Year       string
	GivenGenre string
	VotedGenre string
	Comment    string
	Composer   string
	Lyrics     string
}

func dbCreate(dbName string) {

	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table musicDB (name text not null primary key, routePath text, path text, title text, artist text, album text, year text, givenGenre text, votedGenre text, comment text, composer text, lyrics text);
	delete from musicDB;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func readDb() []DbModel {
	database, _ := sql.Open("sqlite3", "./musicDB.db")
	rows, _ := database.Query("SELECT name, routePath, path, title, artist, album, year, givenGenre, votedGenre, comment, composer, lyrics FROM musicDB")
	var modelArray []DbModel
	var name string
	var routePath string
	var path string
	var title string
	var artist string
	var album string
	var year string
	var givenGenre string
	var votedGenre string
	var comment string
	var composer string
	var lyrics string

	for rows.Next() {
		rows.Scan(&name, &routePath, &path, &title, &artist, &album, &year, &givenGenre, &votedGenre, &comment, &composer, &lyrics)
		modelRow := DbModel{Name: name, RoutePath: routePath, Path: path, Title: title, Artist: artist, Album: album, Year: year, GivenGenre: givenGenre, VotedGenre: votedGenre, Comment: comment, Composer: composer, Lyrics: lyrics}
		modelArray = append(modelArray, modelRow)
	}
	return modelArray
}

func dbInsert(dbName string, name string, routePath string, path string, title string, artist string, album string, year string, givenGenre string, votedGenre string, comment string, composer string, lyrics string) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vals := `VALUES("` + name + `" , "` + routePath + `" , "` + path + `" , "` + title + `" , "` + artist + `" , "` + album + `" , "` + year + `" , "` + givenGenre + `" , "` + votedGenre + `" , "` + comment + `" , "` + composer + `" , "` + lyrics + `" )`
	sqlStmt := `
	INSERT INTO musicDB (name,routePath,path,title,artist,album,year,givenGenre,votedGenre,comment,composer,lyrics) ` + vals
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
