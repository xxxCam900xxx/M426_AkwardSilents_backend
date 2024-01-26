package functions

import (
	db "AkwardSilents/pkg/service/dbfunctions"
    "encoding/json"
	"fmt"
)

type Result struct {
    Message string `json:"message"`
    Sender string `json:"sender"`
    Status int `json:"status"`
    Time int `json:"time"`
}

func GetMessage(content map[string]string, name string) string {
    if name != "" {
		if CheckValues(content, []string{"id_Personchat", "Limit"}) {
            id_Personchat := content["id_Personchat"]
            limit := content["Limit"]
            db.InitDB()
            fmt.Println("wird ausgeführt get message",name)
            rows, err := db.DB.Query("SELECT message, status, time, sender FROM Message WHERE id_Personchats = ? ORDER BY time desc Limit ?;", id_Personchat, limit)
            if err != nil {
                panic(err)
            }
            defer rows.Close()

            results := make([]Result, 0)
            for rows.Next() {
                var message, sender string
                var status, time int
                err := rows.Scan(&message, &status, &time, &sender)
                if err != nil {
                    panic(err)
                }
                fmt.Println(message, status, time, sender, "WWWWWWWWWWWWWWWWWW")
                results = append(results, Result{Message: message, Sender: sender, Status: status, Time: time})
            }

            jsonBytes, err := json.Marshal(results)
            fmt.Println(results,string( jsonBytes))
            if err != nil {
                panic(err)
            }
            return string(jsonBytes)
        }
    }
    return "logout"
}
