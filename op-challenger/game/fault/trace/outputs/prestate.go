package outputs

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ types.PrestateProvider = (*OutputPrestateProvider)(nil)

type OutputPrestateProvider struct {
	prestateBlock uint64
	rollupClient  OutputRollupClient
}

func NewPrestateProvider(rollupClient OutputRollupClient, prestateBlock uint64) *OutputPrestateProvider {
	return &OutputPrestateProvider{
		prestateBlock: prestateBlock,
		rollupClient:  rollupClient,
	}
}

func (o *OutputPrestateProvider) AbsolutePreStateCommitment(ctx context.Context) (hash common.Hash, err error) {
	return o.outputAtBlock(ctx, o.prestateBlock)
}

func (o *OutputPrestateProvider) outputAtBlock(ctx context.Context, block uint64) (common.Hash, error) {
	output, err := o.rollupClient.OutputAtBlock(ctx, block)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to fetch output at block %v: %w", block, err)
	}
	return common.Hash(output.OutputRoot), nil
}

type MockPrestateProvider struct {
	prestateBlock uint64
	// rollupClient  OutputRollupClient
}

func NewMockPrestateProvider(prestateBlock uint64) *MockPrestateProvider {
	return &MockPrestateProvider{
		prestateBlock: prestateBlock,
	}
}

func (o *MockPrestateProvider) AbsolutePreStateCommitment(ctx context.Context) (hash common.Hash, err error) {
	// return o.outputAtBlock(ctx, o.prestateBlock)
	return o.outputAtBlock(ctx, o.prestateBlock)
}

func (o *MockPrestateProvider) outputAtBlock(ctx context.Context, block uint64) (common.Hash, error) {
	return common.Hash{}, nil
}
