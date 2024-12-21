package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Declare a struct for message
type Message struct {
	Content string `json:"content"`
	Topic   string `json:"topic"`
}

type Client struct{}

var messages = make(map[string][]Message)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/add", AddMessageHandler).Methods("POST")
	r.HandleFunc("/get", GetMessageHandler).Methods("GET")
	http.ListenAndServe("127.0.0.1:8080", r)
	println("Hello, World!")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if messages[topic] == nil {
		w.Write([]byte("[]"))
	} else {
		j, _ := json.Marshal(messages[topic])
		w.Write(j)
	}
}

func AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	content := r.FormValue("content")
	topic := r.FormValue("topic")
	if messages[topic] == nil {
		messages[topic] = make([]Message, 0)
	}
	messages[topic] = append(messages[topic], Message{content, topic})
	w.Write([]byte("Message added successfully"))
}
