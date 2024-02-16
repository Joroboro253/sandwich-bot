package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func SubscribeToUniswapEvents(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Couldn`t subscribe to new block headers")
	}

	fmt.Println("Subscription on new transaction...")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Printf("New block: #%v\n", header.Number.String())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatalf("Getting a block failed: %v", err)
			}

			for _, tx := range block.Transactions() {
				fmt.Printf("New transaction: %s\n", tx.Hash().Hex())
				//  transaction analyze
			}
		}
	}
}
