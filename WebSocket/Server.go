package main 

import("fmt"
"net/http"
"log"
"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
ReadBufferSize:1024,
WriteBufferSize:1024,
}

func homepage(w http.ResponseWriter,r*http.Request){
	fmt.Fprintf(w,"Hola Mundo")
}

// check for incoming messages with loop
func ReadMessages(conn*websocket.Conn){
for{
	messageType,p,err := conn.ReadMessage()
	if(err != nil){
		log.Println(err)
		return
	}
	log.Println(string(p))
// message has to be sent back to the client
if err := conn.WriteMessage(messageType,p); err != nil{
	log.Println(err)
	return
}
return
}
}

func wsEndpoint(w http.ResponseWriter,r*http.Request){
	//fmt.Fprintf(w,"WS endpoint")
	// allow all incoming connections to the wsEndpoint 
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws,err := upgrader.Upgrade(w,r,nil)
	if(err != nil){
		fmt.Println(err)
	}
	log.Println("Client Successfully Connected");
	fmt.Println("Client Successfully Connected");
	ReadMessages(ws)
}

func setuproutes(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main(){
	fmt.Println("Hello World")
	setuproutes()
	log.Fatal(http.ListenAndServe(":8080",nil))
}