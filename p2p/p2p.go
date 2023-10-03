package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/blockchain"
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
	fmt.Printf("%s wants an upgrade \n", openPort)

	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	initPeer(conn, ip, openPort)
}

func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	fmt.Printf("%s wants to connect to port %s \n", openPort, port)
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	peer := initPeer(conn, address, port)
	sendNewestBlock(peer)
}

func BroadcastNewBlock(b *blockchain.Block) {
	for _, p := range Peers.v {
		notifyNewBlock(b, p)
	}
}
