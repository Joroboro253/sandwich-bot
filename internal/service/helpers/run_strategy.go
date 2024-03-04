package helpers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

	go subscribeToNewBlocks(ethClient, blockCh)

	go subscribeToPendingTransactions(rpcClient, ethClient, txCh)

	for {
		select {
		case block := <-blockCh:
			fmt.Printf("New block: #%v, BaseFee: %v\n", block.BlockNumber, block.BaseFee)
		case tx := <-txCh:
			fmt.Printf("Pending transaction: %v\n", tx.Hash().Hex())
		}
	}
}
