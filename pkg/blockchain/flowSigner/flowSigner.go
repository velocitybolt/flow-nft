package flowSigner

import (
	"context"
	"engram-flow-api-v1/pkg/blockchain/flowAccountInfo"
	"engram-flow-api-v1/pkg/blockchain/flowUtility"
	"engram-flow-api-v1/pkg/blockchain/generateAccountKeys"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
	"google.golang.org/grpc"
)

func SignFlowCreateAccServiceAccount(currentTransaction *flow.Transaction) string {
	// need to refactor code to get rid of this function later...
	ctx := context.Background()
	c, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)
	transactionSigner := flowAccountInfo.ServiceAccount()
	serviceSigAlgo := crypto.StringToSignatureAlgorithm(transactionSigner.ServiceSigAlgo)
	servicePrivKey, err := crypto.DecodePrivateKeyHex(serviceSigAlgo, transactionSigner.PrivateKeyHex)
	generateAccountKeys.Except(err)

	serviceAccount, err := c.GetAccountAtLatestBlock(ctx, transactionSigner.OwnerAddress)
	generateAccountKeys.Except(err)

	serviceAccountKey := serviceAccount.Keys[0]
	serviceSigner := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	refBlockId := flowUtility.GetReferenceBlockId(c)

	currentTransaction.SetProposalKey(
		transactionSigner.OwnerAddress,
		serviceAccountKey.Index,
		serviceAccountKey.SequenceNumber,
	)

	currentTransaction.SetReferenceBlockID(refBlockId)
	currentTransaction.SetPayer(transactionSigner.OwnerAddress)

	err = currentTransaction.SignEnvelope(transactionSigner.OwnerAddress, serviceAccountKey.Index, serviceSigner)
	generateAccountKeys.Except(err)

	err = c.SendTransaction(ctx, *currentTransaction)
	generateAccountKeys.Except(err)

	serviceAccountKey.SequenceNumber++

	return currentTransaction.ID().String()

}

func CreateAndSignFlowTransaction(addr []flow.Address, filePath string, amount cadence.UFix64, transactionSigner *flowAccountInfo.FlowAccountInfo) {
	ctx := context.Background()
	flowClient, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)

	convScript := flowUtility.ConvertScriptTwo(addr, filePath)
	serviceAccount, err := flowClient.GetAccountAtLatestBlock(ctx, transactionSigner.OwnerAddress)
	generateAccountKeys.Except(err)

	serviceSigAlgo := crypto.StringToSignatureAlgorithm(transactionSigner.ServiceSigAlgo)
	servicePrivKey, err := crypto.DecodePrivateKeyHex(serviceSigAlgo, transactionSigner.PrivateKeyHex)
	generateAccountKeys.Except(err)

	serviceAccountKey := serviceAccount.Keys[0]
	serviceSigner := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	refBlockId := flowUtility.GetReferenceBlockId(flowClient)
	convertedAddr := cadence.BytesToAddress(addr[0].Bytes())

	tx := flow.NewTransaction().
		SetGasLimit(100).
		SetScript(convScript).
		SetProposalKey(transactionSigner.OwnerAddress, serviceAccountKey.Index, serviceAccountKey.SequenceNumber).
		SetReferenceBlockID(refBlockId).
		SetPayer(transactionSigner.OwnerAddress).
		AddAuthorizer(transactionSigner.OwnerAddress)

	if len(addr) > 2 {
		// 2 Args or more [scuffed]
		err = tx.AddArgument(convertedAddr)
		generateAccountKeys.Except(err)
		err = tx.AddArgument(amount)
		generateAccountKeys.Except(err)

		fmt.Println("Sending transaction:")
		fmt.Println()
		fmt.Println("----------------")
		fmt.Println("Script:")
		fmt.Println(string(tx.Script))
		fmt.Println("Arguments:")
		fmt.Printf("Reciever: %s\n", convertedAddr)
		fmt.Printf("Amount: %s\n", amount)
		fmt.Println("----------------")
		fmt.Println()
	} else {
		// No Args
		fmt.Println("Sending transaction:")
		fmt.Println()
		fmt.Println("----------------")
		fmt.Println("Script:")
		fmt.Println(string(tx.Script))
		fmt.Println("----------------")
		fmt.Println()
	}

	err = tx.SignEnvelope(transactionSigner.OwnerAddress, serviceAccountKey.Index, serviceSigner)
	generateAccountKeys.Except(err)

	err = flowClient.SendTransaction(ctx, *tx)
	generateAccountKeys.Except(err)

	_ = flowUtility.WaitForSeal(ctx, flowClient, tx.ID())

}

func SubmitFlowTransaction(currentTransaction *flow.Transaction, transactionSigner *flowAccountInfo.FlowAccountInfo) string {
	ctx := context.Background()
	c, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)
	serviceSigAlgo := crypto.StringToSignatureAlgorithm(transactionSigner.ServiceSigAlgo)
	servicePrivKey, err := crypto.DecodePrivateKeyHex(serviceSigAlgo, transactionSigner.PrivateKeyHex)
	generateAccountKeys.Except(err)

	serviceAccount, err := c.GetAccountAtLatestBlock(ctx, transactionSigner.OwnerAddress)
	generateAccountKeys.Except(err)

	serviceAccountKey := serviceAccount.Keys[0]
	serviceSigner := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	refBlockId := flowUtility.GetReferenceBlockId(c)

	currentTransaction.SetProposalKey(
		transactionSigner.OwnerAddress,
		serviceAccountKey.Index,
		serviceAccountKey.SequenceNumber,
	)

	currentTransaction.SetReferenceBlockID(refBlockId)
	currentTransaction.SetPayer(transactionSigner.OwnerAddress)

	err = currentTransaction.SignEnvelope(transactionSigner.OwnerAddress, serviceAccountKey.Index, serviceSigner)
	generateAccountKeys.Except(err)

	err = c.SendTransaction(ctx, *currentTransaction)
	generateAccountKeys.Except(err)

	serviceAccountKey.SequenceNumber++

	return currentTransaction.ID().String()

}
