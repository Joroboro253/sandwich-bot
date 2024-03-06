package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// Subscribe to new blocks and send them to the channel
func subscribeToNewBlocks(client *ethclient.Client, blockCh chan<- *NewBlock) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for header := range headers {
		blockCh <- &NewBlock{
			BlockNumber: header.Number.Uint64(),
			BaseFee:     header.BaseFee,
		}
	}
}
