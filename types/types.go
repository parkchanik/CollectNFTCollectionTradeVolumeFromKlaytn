package types

import (
	"github.com/klaytn/klaytn/common"
)

type TokenInfo struct {
	TransactionHash  common.Hash
	BlockTime        string
	Contractaddress  common.Address
	ContractName     string
	ContractSymbol   string
	TokenID          string
	WrapTokenAddress common.Address
	WrapTokenName    string
	WrapTokenSymbol  string
	KlayValue        string
}

type Info struct {
	BlockNumber      uint64
	WrapTokenName    string
	WrapTokenSymbol  string
	WrapTokenAddress common.Address
	Klay             int
}
