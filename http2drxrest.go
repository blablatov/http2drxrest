// Web-server-parser модуль для передачи данных в Directum RX, через rest-модуль
package main

import (
	"fmt"
	"log"
	"net/http"
	"restpost"
	"strings"
	"time"
)

type SapData struct {
	Name string `json:"Name"`
	Doc  string `json:"Простой документ"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// this handler is returning component path of URL.
func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// parsing of string
	i := strings.Index(r.URL.Path, ":") // get index of symbol ":"
	sd := &SapData{}
	sd.Name = r.URL.Path[:i]  // get slice of string before symbol ":"
	sd.Doc = r.URL.Path[i+1:] // get slice of string after symbol ":"
	fmt.Println("\nИдентификатор:", strings.TrimPrefix(sd.Name, "/"), "\nТип созданного документа:", sd.Doc)

	cr := make(chan string) // канал функции rest-клиента
	go restpost.PostCreate(sd.Doc, cr)
	// Получение данных из канала cr горутины
	log.Println("\n\nРезультат запроса: Идентификатор: ", <-cr, "\nРезультат запроса: Тип созданного документа: ", <-cr)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs время выполнения запроса\n", secs)
}
