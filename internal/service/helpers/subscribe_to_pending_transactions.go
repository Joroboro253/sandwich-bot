package helpers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func SubscribeToPendingTransactions(rpcClient *rpc.Client) {
	ch := make(chan string)
	_, err := rpcClient.Subscribe(context.Background(), "eth", ch, "newPendingTransactions")
	if err != nil {
		log.Fatalf("Failed to subscribe to new pending transactions: %v", err)
	}

	fmt.Println("Subscribed to new pending transactions")
	for {
		select {
		case txHash := <-ch:
			fmt.Println("New pending transaction:", txHash)
			// Здесь вы можете добавить дополнительную логику для фильтрации и анализа транзакций
		}
	}
}
