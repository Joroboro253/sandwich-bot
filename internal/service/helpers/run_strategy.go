package helpers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

type Event struct {
	Block     *NewBlock
	PendingTx *types.Transaction
}

type NewBlock struct {
	BlockNumber uint64
	BaseFee     *big.Int
}

func RunSandwichStrategy(rpcClient *rpc.Client, ethClient *ethclient.Client) {
	blockCh := make(chan *NewBlock)
	txCh := make(chan *types.Transaction)
	log.Println("RunSandwichStrategy started")
	go subscribeToNewBlocks(ethClient, blockCh)
	go subscribeToPendingTransactions(rpcClient, ethClient, txCh)

	var currentBlock *NewBlock

	for {
		select {
		case block := <-blockCh:
			fmt.Printf("New block: #%v, BaseFee: %v\n", block.BlockNumber, block.BaseFee)
			currentBlock = block

		case tx := <-txCh:
			if currentBlock != nil {
				frame, err := DebugTraceCall(rpcClient, tx.Hash())
				if err != nil {
					continue
				}

				var logs []string
				ExtractLogs(frame, &logs)
				for _, log := range logs {
					fmt.Printf("Log: %s\n", log)
				}
			}
		}
	}
}
