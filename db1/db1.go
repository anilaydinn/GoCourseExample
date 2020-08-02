package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/demodb")

	if err != nil {
		panic(err.Error())
	}

	//sql
	defer db.Close()

	var (
		ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		Birthdate string
		IsActive  bool
	)

	/*
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

			lastID, err1 := res.LastInsertId()

			if err1 != nil {
				log.Fatal(err)
			}

			log.Printf("Last Inserted ID: %d", lastID)




		//Eklenen kayıtları getir.
		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		rows.Close()

	*/

	/*
		rows, errQ := db.Query("SELECT * FROM users WHERE ID = ?", 9)
		if err != nil {
			log.Fatal(errQ)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))
		}
		errQ = rows.Err()
		if errQ != nil {
			log.Fatal(errQ)
		}
	*/

	/*
		err = db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
		if err != nil {
			if err == sql.ErrNoRows {

			} else {
				log.Fatal(err)
			}
		}
		log.Println("Bir satır bulundu : ", ID)
	*/

	//Multiple Active Result Set : MARS
	//_, err = db.Exec("DELETE FROM xTable1; DELETE FROM xTable2;")

	//Preparing Queries
	stmt, errQ := db.Prepare("SELECT * FROM users WHERE id = ?")
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer stmt.Close()

	rows, errX := stmt.Query(5)
	if errX != nil {
		log.Fatal(errX)
	}
	defer rows.Close()

	for rows.Next() {
		scanErr := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &Birthdate, &IsActive)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+Birthdate+" "+strconv.FormatBool(IsActive))
	}
}
