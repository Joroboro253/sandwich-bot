package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

// Subscribe to pending transactions and send them to the channel
func subscribeToPendingTransactions(rpcClient *rpc.Client, ethClient *ethclient.Client, txCh chan<- *types.Transaction) {
	ch := make(chan common.Hash)

	_, err := rpcClient.Subscribe(context.Background(), "eth", ch, "newPendingTransactions")
	if err != nil {
		log.Fatal(err)
	}

	for hash := range ch {
		tx, isPending, err := ethClient.TransactionByHash(context.Background(), hash)
		if err != nil || !isPending {
			continue
		}
		txCh <- tx
	}
}
