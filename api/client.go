package api

import (
	"context"

	evmclient "github.com/MetalBlockchain/coreth/plugin/evm/client"
	"github.com/MetalBlockchain/metalgo/api/admin"
	"github.com/MetalBlockchain/metalgo/api/health"
	"github.com/MetalBlockchain/metalgo/api/info"
	"github.com/MetalBlockchain/metalgo/indexer"
	"github.com/MetalBlockchain/metalgo/utils/rpc"
	"github.com/MetalBlockchain/metalgo/vms/avm"
	"github.com/MetalBlockchain/metalgo/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() *platformvm.Client
	XChainAPI() *avm.Client
	XChainWalletAPI() *avm.WalletClient
	CChainAPI() evmclient.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() *info.Client
	HealthAPI() HealthClient
	AdminAPI() *admin.Client
	PChainIndexAPI() *indexer.Client
	CChainIndexAPI() *indexer.Client
	// TODO add methods
}

type HealthClient interface {
	Health(ctx context.Context, tags []string, options ...rpc.Option) (*health.APIReply, error)
	Readiness(ctx context.Context, tags []string, options ...rpc.Option) (*health.APIReply, error)
	Liveness(ctx context.Context, tags []string, options ...rpc.Option) (*health.APIReply, error)
}
