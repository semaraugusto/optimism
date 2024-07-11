package execution

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var absolutePrestate = common.FromHex("0000000000000000000000000000000000000000000000000000000000000060")
var absolutePrestateInt = new(big.Int).SetBytes(absolutePrestate)

var _ types.PrestateProvider = (*ExecutionPrestateProvider)(nil)

// PrestateProvider provides the alphabet VM prestate
var PrestateProvider = &ExecutionPrestateProvider{}

// ExecutionPrestateProvider is a stateless [PrestateProvider] that
// uses a pre-determined, fixed pre-state hash.
// type ExecutionPrestateProvider struct{}
type ExecutionPrestateProvider struct {
	// prestateBlock uint64
	// prestateHash  []byte
	// rollupClient  OutputRollupClient
}

func (ap *ExecutionPrestateProvider) AbsolutePreStateCommitment(_ context.Context) (common.Hash, error) {
	hash := common.BytesToHash(crypto.Keccak256(absolutePrestate))
	hash[0] = mipsevm.VMStatusUnfinished
	return hash, nil
}

func NewExecutionPrestateProvider() *ExecutionPrestateProvider {
	return &ExecutionPrestateProvider{}
}
