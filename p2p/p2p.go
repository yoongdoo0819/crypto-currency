package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// port :3000 will upgrade the request from :4000
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	openPort := r.URL.Query().Get("openPort")

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}

	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	peer := initPeer(conn, ip, openPort)

	time.Sleep(5 * time.Second)
	peer.inbox <- []byte("Hello from 3000!")
}

func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	peer := initPeer(conn, address, port)
	sendNewestBlock(peer)
}
