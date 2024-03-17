package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (_IQuoterV2 *IQuoterV2Caller) QuoteExactOutputSingleStatic(opts *bind.CallOpts, params IQuoterV2QuoteExactOutputSingleParams) (struct {
	AmountIn                *big.Int
	SqrtPriceX96After       *big.Int
	InitializedTicksCrossed uint32
	GasEstimate             *big.Int
}, error) {
	var out []interface{}
	err := _IQuoterV2.contract.Call(opts, &out, "quoteExactOutputSingle", params)

	outstruct := new(struct {
		AmountIn                *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SqrtPriceX96After = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitializedTicksCrossed = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasEstimate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}
