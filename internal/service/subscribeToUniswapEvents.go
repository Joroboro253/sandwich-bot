package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func SubscribeToUniswapEvents(client *ethclient.Client) {
	swapSignature := []byte("Swap(address,uint256,uint256,uint256,uint256,address)")
	swapSigHash := crypto.Keccak256Hash(swapSignature)

	// Subscription on logs
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{{swapSigHash}},
	}, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if len(vLog.Topics) == 3 { // Uniswap v2 Swap have 2 indexed parameters + event signature
				if err != nil {
					log.Printf("Error decoding swap event: %s", err)
					continue
				}
				log.Printf("Detected Uniswap v2 Swap event: %v\n", vLog)
			}
		}
	}
}
