// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/mint/types
package mint

import (
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

const (
	ModuleName            = types.ModuleName
	DefaultParamspace     = types.DefaultParamspace
	StoreKey              = types.StoreKey
	QuerierRoute          = types.QuerierRoute
	QueryParameters       = types.QueryParameters
	QueryInflation        = types.QueryInflation
	QueryAnnualProvisions = types.QueryAnnualProvisions
)

var (
	// functions aliases
	ParamKeyTable = types.ParamKeyTable
	NewParams     = types.NewParams
	DefaultParams = types.DefaultParams

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	KeyMintDenom           = types.KeyMintDenom
	KeyInflationRateChange = types.KeyInflationRateChange
	KeyInflationMax        = types.KeyInflationMax
	KeyInflationMin        = types.KeyInflationMin
	KeyGoalBonded          = types.KeyGoalBonded
	KeyBlocksPerYear       = types.KeyBlocksPerYear
)

type (
	Params = types.Params
)
