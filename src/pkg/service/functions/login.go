package functions

import (
	"fmt"
	"AkwardSilents/pkg/service/dbfunctions"
)

func Login(content map[string]string) {
	val, ok := content["Key"]
	// If the key exists
	if !ok {
		// send back login failed
		fmt.Println("Key not here")
	}
	fmt.Println(val);
	if ok {
		err := db.InitDB()
		if err != nil {
			panic(err)
		}

		_, err = db.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)`)
		if err != nil {
			panic(err)
		}


  // Daten aus der Tabelle abrufen
  rows, err := db.DB.Query("SELECT * FROM users")
  if err != nil {
	panic(err)
}
  defer rows.Close()

	// Ergebnisse ausgeben
		for rows.Next() {
			var id int
			var name string
			var email string
			err = rows.Scan(&id, &name, &email)
			if err != nil {
				panic(err)			}
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
		}
	}
}
