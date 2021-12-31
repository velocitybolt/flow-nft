package flowAccountInfo

import (
	"github.com/onflow/flow-go-sdk"
)

type FlowAccountInfo struct {
	AccountType       string
	NodeAddress       string
	PublicKey         string
	PrivateKey        string
	AccountAddressHex string
	PublicKeyHex      string
	PrivateKeyHex     string
	SigntureAlgo      string
	HashAlgo          string
	ServiceSigAlgo    string
	OwnerAddress      flow.Address
}

func ServiceAccount() *FlowAccountInfo {
	acc := FlowAccountInfo{}
	acc.AccountType = "serviceAccount"
	acc.NodeAddress = "127.0.0.1:3569"
	acc.PrivateKeyHex = "58d21f786881e01940b87680cca9a35400526be810b0159d3a9a96a8b14c4e9f"
	acc.AccountAddressHex = "f8d6e0586b0a20c7"
	acc.SigntureAlgo = "ECDSA_P256"
	acc.HashAlgo = "SHA3_256"
	acc.ServiceSigAlgo = "ECDSA_P256"
	acc.OwnerAddress = flow.HexToAddress(acc.AccountAddressHex)

	return &acc
}

func NewInternalAccount() *FlowAccountInfo {
	acc := FlowAccountInfo{}
	acc.AccountType = "internalAccount"
	acc.SigntureAlgo = "ECDSA_P256"
	acc.HashAlgo = "SHA3_256"
	acc.ServiceSigAlgo = "ECDSA_P256"

	return &acc
}
