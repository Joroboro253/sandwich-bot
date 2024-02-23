package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func SubscribeToPendingTransactions(rpcClient *rpc.Client, ethClient *ethclient.Client) error {
	ch := make(chan common.Hash)
	sub, err := rpcClient.EthSubscribe(context.Background(), ch, "newPendingTransactions")

	if err != nil {
		log.Fatalf("Failed to subscribe to new pending transactions: %v", err)
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
			return err
		case txHash := <-ch:
			analyzeTransaction(txHash, ethClient)
		}
	}
}
