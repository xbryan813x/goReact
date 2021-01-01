package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xbryan813x/goReact/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	fmt.Println(r.Method)
	fmt.Fprintf(w, "Simple Server")

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/to-do-list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "to-do-list")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		// body := fetchrecord.FetchRecord()
		// fmt.Println("This is the body => \n", body)

		switch r.Method {
		case "GET":
			for k, v := range r.URL.Query() {
				fmt.Printf("%s: %s\n", k, v)
			}
			w.Write([]byte("Received a GET request\n"))
		case "POST":
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}

			// var note Note

			error := json.Unmarshal([]byte(reqBody), &note)

			if error != nil {
				log.Fatal(error)
			}
			fmt.Printf("%s\n", reqBody)
			fmt.Printf("%+v\n", note)
			w.Write([]byte("Received a POST request\n"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		}

	})

	http.HandleFunc("/ws", serveWs)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
