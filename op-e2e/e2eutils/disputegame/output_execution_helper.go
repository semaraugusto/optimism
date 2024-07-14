package disputegame

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/execution"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/outputs"
	"github.com/ethereum-optimism/optimism/op-challenger/metrics"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger"
	// "github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type OutputExecutionGameHelper struct {
	OutputGameHelper
}

func (g *OutputExecutionGameHelper) StartChallenger(
	ctx context.Context,
	l2Node string,
	name string,
	options ...challenger.Option,
) *challenger.Helper {
	opts := []challenger.Option{
		challenger.WithAlphabet(),
		challenger.WithoutLayer2(true),
		challenger.WithExecution(),
		challenger.WithFactoryAddress(g.FactoryAddr),
		challenger.WithGameAddress(g.Addr),
	}
	opts = append(opts, options...)
	c := challenger.NewChallenger(g.T, ctx, g.System, name, opts...)

	g.T.Cleanup(func() {
		_ = c.Close()
	})
	return c
}

func (g *OutputExecutionGameHelper) CreateHonestActor(ctx context.Context) *OutputHonestHelper {
	logger := testlog.Logger(g.T, log.LevelInfo).New("role", "HonestHelper", "game", g.Addr)
	prestateBlock, poststateBlock, err := g.Game.GetBlockRange(ctx)
	g.Require.NoError(err, "Get block range")
	splitDepth := g.SplitDepth(ctx)
	l1Head := g.GetL1Head(ctx)
	prestateProvider := execution.NewTraceProvider(big.NewInt(int64(prestateBlock)), splitDepth)

	correctTrace, err := outputs.NewOutputAlphabetTraceAccessor(logger, metrics.NoopMetrics, prestateProvider, nil, nil, l1Head, splitDepth, prestateBlock, poststateBlock)
	g.Require.NoError(err, "Create trace accessor")
	return NewOutputHonestHelper(g.T, g.Require, &g.OutputGameHelper, g.Game, correctTrace)
}

func (g *OutputExecutionGameHelper) CreateDishonestHelper(ctx context.Context, defender bool) *DishonestHelper {
	return newDishonestHelper(&g.OutputGameHelper, g.CreateHonestActor(ctx), defender)
}
