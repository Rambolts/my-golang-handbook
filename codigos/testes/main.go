package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	body := []byte(`{
		"additionalProp1": "string",
		"additionalProp2": "string",
		"additionalProp3": "string"
	  }`)
	resp, err := http.Post(
		"https://app.agnet.com.br/bpmsapi/api/Acesso/gerar/eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1lIjoiX3NydldlYlNlcnZpY2VzQUciLCJuYW1laWRlbnRpZmllciI6Il9zcnZXZWJTZXJ2aWNlc0FHIiwiZW1haWxhZGRyZXNzIjoiX3NydldlYlNlcnZpY2VzQUdAYWduZXQuY29tLmJyIiwiQXBsaWNhY2FvSWQiOiJBTkRSQURFX0JQTVMiLCJleHAiOjE3NjE3NjY3Nzh9.WRa5eHLqMW17cOyvoTUj1o54gZklAXK2mGw2mCbWV-U",
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Falha ao fazer requisição: ", err)
	}
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["json"])
}
