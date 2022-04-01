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
	WrapTokenAddress common.Address
	Klay             int
	//WrapTokenSymbol  string
}

type CollectionInfo struct {
	Contractaddress      common.Address
	CollectionName       string
	CollectionSymbol     string
	TradeTransactinCount int
	Klay                 uint64
	KlayFloat            string
	TokenIDs             []string
	FromToBlockNumber    string
}
