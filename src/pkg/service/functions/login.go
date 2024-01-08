package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

func login(content map[string]string) {
	val, ok := content["key"]
	// If the key exists
	if !ok {
		// send back login failed
	}

	if ok {
		err := db.InitDB()
		if err != nil {
			panic(err)
		}
	
		rows, err := db.DB.Query("SELECT * FROM User")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
	
		for rows.Next() {
			var ID int
			err = rows.Scan(&ID)
			if err != nil {
				panic(err)
			}
			fmt.Println(ID)
		}
	}
}