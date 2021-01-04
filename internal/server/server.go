package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xbryan813x/goReact/fetchrecord"
	"github.com/xbryan813x/goReact/notes"
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

func updateStore(note notes.Note) {
	notes.Store = append(notes.Store, note)
	fmt.Println("Updating Store with note =>", notes.Store)
}

func setupRoutes() {
	http.HandleFunc("/to-do-list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "to-do-list")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		switch r.Method {
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
			return
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

			var note notes.Note

			error := json.Unmarshal([]byte(reqBody), &note)

			if error != nil {
				log.Fatal(error)
			}
			updateStore(note)
			match := fetchrecord.FetchRecord(note)

			if !match.IsEmpty() {
				matchPointer := &match
				response, err := json.Marshal(matchPointer)
				if err != nil {
					fmt.Println(err)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(response)
				fmt.Println("Hoopla!!")
			}
			fmt.Println("Sample Record from cleanup => ")
			fmt.Println(match)
			// w.Write([]byte("Received a POST request\n"))
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
