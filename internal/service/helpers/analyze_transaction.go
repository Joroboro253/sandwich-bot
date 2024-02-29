package helpers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func analyzeTransaction(txHash common.Hash, ethClient *ethclient.Client, parsedABI abi.ABI) {
	tx, isPending, err := ethClient.TransactionByHash(context.Background(), txHash)
	if err != nil || !isPending {
		// Skip if any errors or if the transaction is no longer pending
		return
	}

	if tx.To() == nil {
		// Skip contract creation transaction
		return
	}

	method, err := parsedABI.MethodById(tx.Data())
	if err == nil {
		fmt.Printf("Detected Uniswap V2 operation: %s  Method: %s\n", tx.Hash().Hex(), method.Name)
	}
}
