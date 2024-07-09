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

// alphabetPrestateProvider is a stateless [PrestateProvider] that
// uses a pre-determined, fixed pre-state hash.
// type ExecutionPrestateProvider struct{}
type ExecutionPrestateProvider struct {
	prestateBlock uint64
	prestateHash  []byte
	// rollupClient  OutputRollupClient
}

func (ap *ExecutionPrestateProvider) AbsolutePreStateCommitment(_ context.Context) (common.Hash, error) {
	hash := common.BytesToHash(crypto.Keccak256(absolutePrestate))
	hash[0] = mipsevm.VMStatusUnfinished
	return hash, nil
}

func NewExecutionPrestateProvider() *ExecutionPrestateProvider {
	return &ExecutionPrestateProvider{
		// prestateBlock: prestateBlock,
		// prestateHash:  prestateHash,
		// rollupClient:  rollupClient,
	}
}

//
// func (o *ExecutionPrestateProvider) AbsolutePreStateCommitment(ctx context.Context) (common.Hash, error) {
// 	// hash = common.BytesToHash(crypto.Keccak256(o.prestateHash))
// 	// hash[0] = 3
// 	return common.HexToHash("0xDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEF"), nil
// 	// return o.outputAtBlock(ctx, o.prestateBlock)
// }
//
// func (o *ExecutionPrestateProvider) outputAtBlock(context.Context, uint64) (common.Hash, error) {
// 	dst := make([]byte, 32)
// 	binary.LittleEndian.PutUint64(dst, o.prestateBlock+1)
//
// 	return common.BytesToHash(dst), nil
// }
