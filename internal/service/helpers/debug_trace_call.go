package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

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
