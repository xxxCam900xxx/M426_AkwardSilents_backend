package functions

import (
	"fmt"
	"AkwardSilents/pkg/service/dbfunctions"
)

func Login(content map[string]string, userName *string) string {
	key, ok := content["Key"]
	// If the key exists
	if !ok {
		// send back login failed
		return "login failed"
	}
	fmt.Println(key);
	if ok {
		// Daten aus der Tabelle abrufen
		// Daten aus der Tabelle abrufen
		rows, err := db.DB.Query("SELECT name FROM User WHERE key = ?", key)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// Ergebnisse ausgeben
		for rows.Next() {
			var name string
			err = rows.Scan(&name)
			*userName = name	
			if err != nil {
				panic(err)			
			}
			return "login success"
		}
		return "login failed"
	}
	return "login failed"
}
