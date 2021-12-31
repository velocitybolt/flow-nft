package createAccount

import (
	"engram-flow-api-v1/pkg/blockchain/flowAccountInfo"
	"engram-flow-api-v1/pkg/blockchain/flowSigner"
	"engram-flow-api-v1/pkg/blockchain/flowUtility"
	"engram-flow-api-v1/pkg/blockchain/generateAccountKeys"
	"log"
	"time"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
)

func InitalizeUser() (string, flow.Address, *flowAccountInfo.FlowAccountInfo, string, error) {

	serviceAccount := flowAccountInfo.ServiceAccount()
	createdAccount := flowAccountInfo.NewInternalAccount()

	publicKey, privateKey := generateAccountKeys.CreateKeys("ECDSA_P256")

	txID := CreateAccount(serviceAccount.NodeAddress, publicKey, serviceAccount.SigntureAlgo, serviceAccount.HashAlgo,
		serviceAccount.AccountAddressHex, serviceAccount.PrivateKeyHex, serviceAccount.ServiceSigAlgo, 100)

	blockTime := 10 * time.Second
	time.Sleep(blockTime)

	addr := flowUtility.GetAddress(txID)

	flowAddr := flowUtility.GetAddressFlow(txID)

	createdAccount.PrivateKeyHex = privateKey
	createdAccount.PublicKeyHex = publicKey
	createdAccount.OwnerAddress = flowAddr

	log.Printf("Account Public Key: %s \n, Account Private Key: %s \n, Account Transaction ID: %s \n, Account Address: %s",
		publicKey, privateKey, txID, addr)

	return addr, flowAddr, createdAccount, txID, nil

}

func CreateAccount(node string, pubKeyHex string, sigAlgoN string,
	hashAlgoN string, servAddHex string, servPrivHex string,
	servSigAlgoN string, gasLimit uint64) string {
	/*
		node -> the IP and port of the node we will connect to, in our case this is the Emulator
		pubKeyHex ->  the hexadecimal representation of our public key
		sigAlgoN -> the name of the signature algorithm used when we generated our keys
		hashAlgoN ->  the name of the hashing algorithm that we want this public key to use when verifying signatures.
		servAddHex -> the service address in hexadecimal format
		servPrivHex -> the service private key in hexadecimal format
		servSigAlgoN -> the name of the signature algorithm that was used when generating the service keys
		gasLimit -> the maximum amount of Flow tokens that will be spent in transaction fees to create this account.
	*/

	// for now the service account creates accounts but will need to change to owners account address when switch to flow testnet/mainnet!!!

	sigAlgo := crypto.StringToSignatureAlgorithm(sigAlgoN)
	pubKey, err := crypto.DecodePublicKeyHex(sigAlgo, pubKeyHex)
	generateAccountKeys.Except(err)
	hashAlgo := crypto.StringToHashAlgorithm(hashAlgoN)
	accountKey := flow.NewAccountKey().
		SetPublicKey(pubKey).
		SetSigAlgo(sigAlgo).
		SetHashAlgo(hashAlgo).
		SetWeight(1500)

	serviceAddress := flow.HexToAddress(servAddHex)
	tx := templates.CreateAccount([]*flow.AccountKey{accountKey}, nil, serviceAddress)

	servAccInfo := flowAccountInfo.ServiceAccount()
	txID := flowSigner.SubmitFlowTransaction(tx, servAccInfo)

	return txID
}
