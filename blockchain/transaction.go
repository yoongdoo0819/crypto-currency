package blockchain

import (
	"time"

	"github.com/nomadcoders/nomadcoin/utils"
)

const (
	minerResard int = 50
)

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

func makeCoinBaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerResard},
	}
	txOuts := []*TxOut{
		{address, minerResard},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}
