package blocktime

import (
	"context"
	"log"
	"main/common/config"
	"math/big"
)

func GetBlockTime(Blocknumber uint64) uint64 {
	blocks, err := config.Client.BlockByNumber(context.Background(), big.NewInt(int64(Blocknumber)))
	if err != nil {
		log.Fatalf("err in GetBlockTime: %v", err)
	}
	config.BlockWithTimestamp[Blocknumber] = blocks.Time()
	return blocks.Time()
}
