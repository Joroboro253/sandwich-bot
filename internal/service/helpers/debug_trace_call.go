package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type SwapDirection string

const (
	Buy  SwapDirection = "buy"
	Sell SwapDirection = "sell"
)

type SwapInfo struct {
	TxHash       common.Hash
	TargetPair   common.Address
	MainCurrency common.Address
	TargetToken  common.Address
	Version      uint8
	Token0IsMain bool
	Direction    SwapDirection
}

type CallFrame struct {
	Type  string         `json:"type"`
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Input string         `json:"input"`
	Calls []*CallFrame   `json:"calls,omitempty"`
	// new field
	Logs []string
}

// DebugTraceCall Function for transaction tracing and returning a result
func DebugTraceCall(client *rpc.Client, txHash common.Hash) (*CallFrame, error) {
	var result CallFrame
	ctx := context.Background()

	err := client.CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
		"tracer": "callTracer",
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}
