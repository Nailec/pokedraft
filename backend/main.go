package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var addr = flag.String("addr", ":8081", "http service address")

var mc *mongo.Client

//func serveHome(w http.ResponseWriter, r *http.Request) {
//	log.Println(r.URL)
//	if r.URL.Path != "/" {
//		http.Error(w, "Not found", http.StatusNotFound)
//		return
//	}
//	if r.Method != "GET" {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	http.ServeFile(w, r, "home.html")
//}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mc, err := mongo.Connect(ctx)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, mc, w, r)
	})
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
