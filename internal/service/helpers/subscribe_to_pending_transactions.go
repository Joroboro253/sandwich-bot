package helpers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func SubscribeToPendingTransactions(rpcClient *rpc.Client, ethClient *ethclient.Client) {
	ch := make(chan common.Hash)
	sub, err := rpcClient.EthSubscribe(context.Background(), ch, "newPendingTransactions")
	if err != nil {
		log.Fatalf("Failed to subscribe to new pending transactions: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Subscribed to new pending transactions")

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case txHash := <-ch:
			//fmt.Println("New pending transaction:", txHash.Hex())
			analyzeTransaction(txHash, ethClient) // Анализ транзакции
		}
	}
}
