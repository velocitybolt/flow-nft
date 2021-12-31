package responses

import (
	blockchainModels "engram-flow-api-v1/internal/models/flowBlockchainModels"
)

type NewAccountResponse struct {
	Id                 int         `json:"id"`
	UserName           string      `json:"userName"`
	AccountId          string      `json:"accountId"`
	TransactionId      string      `json:"TransactionId"`
	CreatedAccountInfo interface{} `json:"CreatedAccountInfo"`
}

type NewNftResponse struct {
	Id                     int      `json:"id"`
	UserName               string   `json:"userName"`
	MainOwnerAddress       string   `json:"mainOwnerAddress"`
	UserAddress            string   `json:"UserAddress"`
	TransactionId          string   `json:"transactionId"`
	RemainingEngramBalance int      `json:"remainingEngramBalance"`
	NftAdresses            []string `json:"nftAddresses"`
}

type NftResponse struct {
	ContractAddress string `json:"contactAddress"`
	AssetAddress    string `json:"assetAddress"`
	NumberMinted    int    `json:"numberMinted"`
	CurrentOwner    string `json:"currentOwner"`
}

type NftsResponse struct {
	Assets []NftResponse `json:assets`
}

type BalanceResponse struct {
	Id            int    `json:"id"`
	UserName      string `json:"username"`
	WalletAddress string `json:"address"`
	Balance       int    `json:"balance"`
}

type EngramCoinSupplyResponse struct {
	ContractAddress string `json:"contractAddress"`
	Supply          int    `json:"supply"`
}

type DeployMainContractsResponse struct {
	ContractsDeployed []blockchainModels.ContractTransaction `json:"contractsDeployed"`
}

type TransferEngramCoinsResponse struct {
	OwnerAddress        string `json:"ownerAddress"`
	ReceiverAddress     string `json:"receiverAddress"`
	NumCoinsToTranser   int    `json:"numCoinsToTransfer"`
	ReceiverPreBalance  int    `json:"receiverPreBalance"`
	ReceiverPostBalance int    `json:"receiverPostBalance"`
}
