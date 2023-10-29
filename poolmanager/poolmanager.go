package poolmanager

import (
	"net/rpc"

	"github.com/ethereum/go-ethereum/log"
)

type PoolManager struct {
	Client         *rpc.Client
}

func NewPoolManager(poolManager string) (*PoolManager, error) {
	client, err := rpc.DialHTTP("tcp", poolManager)
	if err != nil {
		log.Error("Connection error: ", err)
	}
	log.Info("PoolManager initialized successfully", "PoolManager", poolManager)
	return &PoolManager{
		Client:         client,
	}, nil
}
