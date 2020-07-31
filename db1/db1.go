package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/demodb")

	if err != nil {
		panic(err.Error())
	}

	//sql
	defer db.Close()

	createStatement := "`users`(`ID` int(11) NOT NULL AUTO_INCREMENT, `Username` varchar(45) NOT NULL, `Email` varchar(45) NOT NULL, `Password` varchar(45) NOT NULL, `FirstName` varchar(45) NOT NULL, `LastName` varchar(45) NOT NULL, `BirthDate` varchar(45) NOT NULL, IsActive tinyint(1) DEFAULT NULL, PRIMARY KEY (`ID`), UNIQUE KEY `ID_UNIQUE` (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + createStatement)

	if err != nil {
		log.Fatal(err)
	}

	//Veri ekleme
	res, err := db.Exec("INSERT INTO users(Username, Email, Password, FirstName, LastName, BirthDate, IsActive) VALUES('Deneme1','12345','deneme@dene.com','Anıl','Aydın','2017.1.1',1)")

	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted %d rows", rowCount)
}
