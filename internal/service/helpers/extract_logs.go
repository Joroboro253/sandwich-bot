package helpers

import (
	"encoding/json"
	"log"
)

func ExtractLogs(callFrame *CallFrame, logs *[]string) {
	if callFrame == nil {
		return
	}

	callJSON, err := json.Marshal(callFrame)
	if err != nil {
		log.Printf("Error marshalling call frame: %v\n", err)
		return
	}

	*logs = append(*logs, string(callJSON))

	for _, call := range callFrame.Calls {
		ExtractLogs(call, logs)
	}
}
