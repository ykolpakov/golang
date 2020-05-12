// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "crypto/sha256"
    "encoding/binary"
    "strconv"
)
// Запрос представляет собой HTTP-запрос, полученный сервером или отправленный клиентом.
//Интерфейс ResponseWriter используется обработчиком HTTP для создания ответа HTTP.
func handler(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body) //пакет функции ввода-вывода /чтение данных /из тела

    if err != nil {
        panic(err)
    }

    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    fmt.Printf("Received %s\n", reqBody)

    var data map[string]string

    if err := json.Unmarshal(reqBody, &data); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    fmt.Printf("Parsed %s\n", data)

    data["hash"] = string(getHash(data["phrase"]))

    fmt.Printf("Response %s\n", data)

    retData, _ := json.Marshal(data)

    fmt.Fprintf(w, string(retData))
}

func getHash(str string) string {
    fmt.Printf("getHash: %s\n", str)

    sum := sha256.Sum256([]byte(str))
	data := binary.BigEndian.Uint64(sum[:8])
	fmt.Printf("getHash in uint64: %d (HEX %x)\n", data, data)

	return strconv.FormatUint(data, 10)
}

func main() {
    http.HandleFunc("/", handler) //регистрирует функцию-обработчик для данного шаблона в DefaultServeMux.
    log.Fatal(http.ListenAndServe(":8080", nil))
}
