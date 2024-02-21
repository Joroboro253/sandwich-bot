package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"sandwich-bot/internal/config"
	"sandwich-bot/internal/service/helpers"
)

type service struct {
	log       *logan.Entry
	ethClient *ethclient.Client
}

func newService(cfg config.Config) *service {
	return &service{
		log: cfg.Log(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func (s *service) run() error {
	s.log.Info("Service started")
	ethClient, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/76256d7863c8480ba65718f2c4faabf7")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	rpcClient, err := rpc.Dial("wss://goerli.infura.io/ws/v3/76256d7863c8480ba65718f2c4faabf7")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("We are connected to the Goerli testnet!")

	//contractAddress := common.HexToAddress("0xe592427a0aece92de3edee1f18e0157c05861564")

	fmt.Println("Subscribing to Uniswap events...")
	helpers.SubscribeToPendingTransactions(rpcClient, ethClient)

	return nil

}
