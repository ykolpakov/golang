package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8080"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"phrase":"Test phrase1"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req) //отправляет http апрос и возвращает ответ
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //закрытие тела

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}