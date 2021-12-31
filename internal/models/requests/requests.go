package requests

type NewAccountRequest struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
}

type NewNftRequest struct {
	Id               int    `json:"id"`
	UserName         string `json:"userName"`
	MainOwnerAddress string `json:"mainOwnerAddress"`
	UserAddress      string `json:"userAddress"`
	NumberToMint     int    `json:"numberToMint"`
}

type BalanceRequest struct {
	Id            int    `json:"id"`
	UserName      string `json:"username"`
	WalletAddress string `json:"address"`
}

type NftsRequest struct {
	Id            int    `json:"id"`
	UserName      string `json:"username"`
	WalletAddress string `json:"address"`
}

type TransferEngramCoinsRequest struct {
	OwnerAddress      string `json:"ownerAddress"`
	ReceiverAddress   string `json:"receiverAddress"`
	NumCoinsToTranser int    `json:"numCoinsToTransfer"`
}
