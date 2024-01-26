package functions

import (
	db "AkwardSilents/pkg/service/dbfunctions"
    "encoding/json"
	"fmt"
)

type UserGetMembers struct {
    Name string `json:"userName"`
}

func GetMembers() string {
        fmt.Println("wird ausgeführt get members")
		rows, err := db.DB.Query("SELECT name FROM User")
        if err != nil {
            panic(err)
        }
        defer rows.Close()

        results := make([]UserGetMembers, 0)
        for rows.Next() {
            var user string
            err := rows.Scan(&user)
            if err != nil {
                panic(err)
            }
            results = append(results, UserGetMembers{Name: user})
        }
        fmt.Println(results)
        jsonBytes, err := json.Marshal(results)
        if err != nil {
            panic(err)
        }
        return string(jsonBytes)
    }
