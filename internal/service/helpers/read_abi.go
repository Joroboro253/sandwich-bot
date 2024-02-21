package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"os"
)

func ReadAbi(filePath string) abi.ABI {
	// Abi file reading
	fmt.Println("Reading ABI from:", filePath)

	abiFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error during reading ABI file: %v", err)
	}

	// Json deserialization
	var abiData struct {
		Abi json.RawMessage `json:"abi"`
	}
	if err := json.Unmarshal(abiFile, &abiData); err != nil {
		log.Fatalf("Failed to unmarshal ABI: %v", err)
	}

	contractABI, err := abi.JSON(bytes.NewReader(abiData.Abi))
	if err != nil {
		log.Fatalf("Invalid ABI: %v", err)
	}
	fmt.Println("ABI successfully read")

	return contractABI
}
