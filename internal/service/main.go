package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"sandwich-bot/internal/config"
	"sandwich-bot/internal/service/helpers"
	"time"
)

type service struct {
	log       *logan.Entry
	ethClient *ethclient.Client
	rpcClient *rpc.Client
}

func newService(cfg config.Config) *service {
	rpcClient, err := rpc.Dial("wss://goerli.infura.io/ws/v3/76256d7863c8480ba65718f2c4faabf7")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	return &service{
		log:       cfg.Log(),
		ethClient: ethClient,
		rpcClient: rpcClient,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

func (s *service) run() error {
	s.log.Info("Service started")

	s.subscribeToUniswapEvents()
	return nil

}

func (s *service) subscribeToUniswapEvents() {
	fmt.Println("Subscribing to Uniswap events...")
	for {
		if err := helpers.SubscribeToPendingTransactions(s.rpcClient, s.ethClient); err != nil {
			log.Printf("Subscription failed with error: %v. Re-subscribing in 10 seconds...", err)
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}

}
