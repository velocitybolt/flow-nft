package flowBlockchainModels

type Transaction struct {
	Id string
}

type ContractTransaction struct {
	TransactionId string
	ContractName  string
	AccountId     string
}
