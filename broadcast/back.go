package broadcast

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	reader(ws)
}

func jsonRecv(conn *websocket.Conn) {
	m := &MessageRecv{}

	err := conn.ReadJSON(m)
	if err != nil {
		log.Println(err)
	}
}

func jsonResp() {

}

func reader(conn *websocket.Conn) {
	for {
		// Define MessageRecv Structure
		m := &MessageRecv{}

		// Get MessageRecv
		err := conn.ReadJSON(m)
		if err != nil {
			log.Println(err)
			return
		} else {
			log.Println(m.Data)
		}

		// Send MessageRecv
		r := MessageResp{
			SignalType: "conn_resp",
			Data: DataResp{
				Status: "normal",
				Msg:    "connection_normal",
			},
		}
		log.Println(r)
		err = conn.WriteJSON(r)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
