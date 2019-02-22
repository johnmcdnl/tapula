package tapula

import (
	"encoding/json"
	"fmt"
)

func toString(i interface{}) string {
	return toFlatString(i)
}

func toFlatString(i interface{}) string {
	return handleJson(json.Marshal(i))
}

func toPrettyString(i interface{}) string {
	return handleJson(json.MarshalIndent(i, "", "  "))
}

func handleJson(j []byte, err error) string {
	if err != nil {
		panic(err)
	}
	return string(j)
}

func PrintReady() {
	fmt.Println(`
▄▄▄▄▄ ▄▄▄·  ▄▄▄·▄• ▄▌▄▄▌   ▄▄▄· 
•██  ▐█ ▀█ ▐█ ▄██▪██▌██•  ▐█ ▀█ 
 ▐█.▪▄█▀▀█  ██▀·█▌▐█▌██▪  ▄█▀▀█ 
 ▐█▌·▐█ ▪▐▌▐█▪·•▐█▄█▌▐█▌▐▌▐█ ▪▐▌
 ▀▀▀  ▀  ▀ .▀    ▀▀▀ .▀▀▀  ▀  ▀`)
}
