package functions

import (
	"fmt"
	"AkwardSilents/pkg/service/dbfunctions"
	"AkwardSilents/pkg/tools"
)

func Sendmessage(content map[string]string, userName *string) {
    if checkValues(content, []string{"Key", "Name", "Message"}) {
		key, ok := content["Key"]
		sendName, ok := content["Name"]
		message, ok := content["Message"]
        messageStatus := 0

		fmt.Println(key, sendName, message);
        if(tools.sendMessage(sendName, message)){
            messageStatus = 1
        }
        
    }
}


func checkValues(content map[string]string, values []string) bool {
    for _, value := range values {
        if _, ok := content[value]; !ok {
            return false
        }
    }
    return true
}
