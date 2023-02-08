package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/comdex-official/comdex/v8/x/liquidity/client/cli"
	"github.com/comdex-official/comdex/v8/x/liquidity/client/rest"
)

var LiquidityProposalHandler = []govclient.ProposalHandler{
	govclient.NewProposalHandler(cli.NewCmdUpdateGenericParamsProposal, rest.UpdateGenericParamsProposalRESTHandler),
	govclient.NewProposalHandler(cli.NewCmdCreateNewLiquidityPairProposal, rest.CreateNewLiquidityPairProposalRESTHandler),
}
