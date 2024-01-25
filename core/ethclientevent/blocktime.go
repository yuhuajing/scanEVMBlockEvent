package ethclientevent

import (
	"context"
	"main/common/config"
	"math/big"
)

func ChainBlockTime(blockNumber uint64) (uint64, error) {
	block, err := config.Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		return 0, err
	}
	return block.Time(), nil
}
