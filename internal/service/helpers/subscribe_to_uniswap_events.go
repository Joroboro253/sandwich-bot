package helpers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func SubscribeToUniswapEvents(client *ethclient.Client, uniswapV2Address common.Address) {
	fmt.Println(uniswapV2Address)

	// Подписка на логи событий
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{uniswapV2Address},
	}, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	fmt.Println("Subscribed to Uniswap events.")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)
			// Реализация механизма переподключения может быть здесь
			return // или переподключение
		case vLog := <-logs:
			fmt.Printf("Detected Uniswap v2 Swap event: %v\n", vLog)
			// Дальнейшая обработка лога события
		}
	}
}
