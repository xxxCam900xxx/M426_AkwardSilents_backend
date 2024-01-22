package functions

import (
	"fmt"
	"AkwardSilents/pkg/service/dbfunctions"
	"AkwardSilents/pkg/tools"
)

func Login(content map[string]string, userName *string, ws *websocket.Conn) string {
	key, ok := content["Key"]
	// If the key exists
	if !ok {
		// send back login failed
		return "login failed"
	}
	fmt.Println(key);
	if ok {
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
			tools.addName(name, ws)

			if err != nil {
				panic(err)			
			}
			return "login success"
		}
		return "login failed"
	}
	return "login failed"
}