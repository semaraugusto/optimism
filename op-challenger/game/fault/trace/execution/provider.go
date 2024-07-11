package execution

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/alphabet"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"

	"github.com/ethereum/go-ethereum/common"
)

const (
	L2ClaimBlockNumberLocalIndex = 4
)

var (
	ErrIndexTooLarge = errors.New("index is larger than the maximum index")
)

var _ types.TraceProvider = (*ExecutionTraceProvider)(nil)

// ExecutionTraceProvider is a [TraceProvider] that monotonically increments
// the starting l2 block number as the claim value.
type ExecutionTraceProvider struct {
	alphabet.AlphabetTraceProvider
}

// NewTraceProvider returns a new [AlphabetProvider].
func NewTraceProvider(startingBlockNumber *big.Int, depth types.Depth) *ExecutionTraceProvider {
	a := alphabet.NewTraceProvider(startingBlockNumber, depth)
	return &ExecutionTraceProvider{*a}
}

func (ep *ExecutionTraceProvider) ClaimedBlockNumber(pos types.Position) (uint64, error) {
	return 0, fmt.Errorf("This game doesn't care about Claimed Block Number. Don't use this.")
}

func (ap *ExecutionTraceProvider) GetStepData(ctx context.Context, pos types.Position) ([]byte, []byte, *types.PreimageOracleData, error) {
	return ap.GetStepData(ctx, pos)
}

// Get returns the claim value at the given index in the trace.
func (ap *ExecutionTraceProvider) Get(ctx context.Context, i types.Position) (common.Hash, error) {
	return ap.Get(ctx, i)
}

func (ap *ExecutionTraceProvider) GetL2BlockNumberChallenge(_ context.Context) (*types.InvalidL2BlockNumberChallenge, error) {
	return nil, types.ErrL2BlockNumberValid
}
