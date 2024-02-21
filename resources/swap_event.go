package resources

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SwapEvent struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
}
