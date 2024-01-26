package functions

import (
	db "AkwardSilents/pkg/service/dbfunctions"
    "encoding/json"
	"fmt"
)

type ResultOverview struct {
    User1 string `json:"user1"`
    User2 string `json:"user2"`
    Id int `json:"id"`
}

func Overview(name string) string {
    if name != "" {
        fmt.Println("wird ausgeführt",name)
        rows, err := db.DB.Query("SELECT user1, user2, m.id FROM Message AS m LEFT JOIN Personchats AS p ON m.id_Personchats = p.id WHERE user1 = ? OR user2 = ? GROUP BY p.id ORDER BY time desc ;", name, name)
        if err != nil {
            panic(err)
        }
        defer rows.Close()

        results := make([]ResultOverview, 0)
        for rows.Next() {
            var user1, user2 string
            var id int
            err := rows.Scan(&user1, &user2, &id)
            if err != nil {
                panic(err)
            }
            fmt.Println(user1, user2, name, id, results)
            results = append(results, ResultOverview{User1: user1, User2: user2, Id:id})
        }
        fmt.Println(results)
        jsonBytes, err := json.Marshal(results)
        if err != nil {
            panic(err)
        }
        return string(jsonBytes)
    }
    return "logout"
}
