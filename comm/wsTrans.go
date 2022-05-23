package comm

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func connWSS() {
	u, _ := url.Parse("ws://localhost:7890/ws")
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial", err)
	}
	defer c.Close()

}
