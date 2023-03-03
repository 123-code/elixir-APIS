package main

import (
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

type Server struct {
connections map [*websocket.Conn]bool
}

func Connections() *Server{
	return &Server{
     connections: make(map[*websocket.Conn]bool),
	}
}

func( s *Server) handleWS(ws *websocket.Conn){
fmt.Println("Incoming connection from:",ws.RemoteAddr());
s.connections[ws] = true;
s.readLoop(ws)
}


func (s*Server) readLoop(ws * websocket.Conn){
	buf := make([]byte,1024);
	for{
		n,err := ws.Read(buf)
		if err != nil{
			if(err == io.EOF){
				break
			}
			fmt.Println("ERROR",err)
			continue
		}
		msg := string(buf[:n])
		fmt.Println(msg);
		ws.Write([]byte("Message"))
	}
}

func main(){
server := gin.Default();
server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		});
})
server.Run(":8080");
};

