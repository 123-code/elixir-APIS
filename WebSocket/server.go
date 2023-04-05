package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"

)
var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

type client struct{
	hub *Hub

Connection *websocket.conn
sender chan []byte
}

func handler(w http.ResponseWriter,r*http.ResponseController){
conn,err := upgrader.Upgrade(w,r,nil)
if(err != nil){
	fmt.Println("Error");
}else{
	fmt.Println("conexion iniciada");
}

}




