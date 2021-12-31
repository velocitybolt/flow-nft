package cdcExecutor

import (
	"context"
	"engram-flow-api-v1/pkg/blockchain/flowAccountInfo"
	"engram-flow-api-v1/pkg/blockchain/flowSigner"

	"engram-flow-api-v1/pkg/blockchain/flowUtility"
	"engram-flow-api-v1/pkg/blockchain/generateAccountKeys"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/templates"
	"google.golang.org/grpc"
)

func ContractCadence(input string, contractName string, mode string) string {
	// for now the service account deploys contracts but will need to change to owners account address when switch to flow testnet/mainnet !!!
	serviceAccountInfo := flowAccountInfo.ServiceAccount()
	var finalCode string

	// Mode 1: Contract where the input is from a file
	// Mode 2: Contract where input is already a string

	if mode == "1" {
		contractCode := generateAccountKeys.ReadFile(input)
		finalCode = contractCode
	}

	if mode == "2" {
		finalCode = input
	}

	contractTx := templates.CreateAccount(nil,
		[]templates.Contract{{
			Name:   contractName,
			Source: finalCode,
		}}, serviceAccountInfo.OwnerAddress)

	txID := flowSigner.SignFlowCreateAccServiceAccount(contractTx)

	return txID

}

func AddContractToAccount(scriptAddr flow.Address, convScript string, contractName string, transactionSigner *flowAccountInfo.FlowAccountInfo) string {

	tx := templates.AddAccountContract(scriptAddr, templates.Contract{
		Name:   contractName,
		Source: convScript,
	})

	txID := flowSigner.SubmitFlowTransaction(tx, transactionSigner)

	return txID
}

func ScriptCadence(args []interface{}, filePath string) cadence.Value {
	// adjust for other types of cadence inputs later
	var addrs []flow.Address
	var convertedArg []cadence.Value
	var res cadence.Value

	if len(args) != 0 {
		for _, arg := range args {
			switch arg := arg.(type) {

			case flow.Address:
				addrs = append(addrs, arg)

			}

		}
	}

	ctx := context.Background()
	c, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)

	if len(args) > 2 {
		// script that imports engramToken contract + fungToken contract + other args

		convScript := flowUtility.ConvertScriptTwo(addrs, filePath)

		var addrSlice [8]byte = addrs[0]

		convertedArgVal := cadence.NewAddress(addrSlice)

		convertedArg = append(convertedArg, convertedArgVal)

		res, err = c.ExecuteScriptAtLatestBlock(ctx, convScript, convertedArg)
		generateAccountKeys.Except(err)

	} else {

		// script that uses just one contract import
		convScript := flowUtility.ConvertScript(addrs[0], filePath)

		res, err = c.ExecuteScriptAtLatestBlock(ctx, convScript, nil)
		generateAccountKeys.Except(err)

	}

	return res

}

func TransactionCadenceArgs(filePath string, args []interface{}, transactionSigner *flowAccountInfo.FlowAccountInfo) {
	// supports two arguments as of now (may need it to support more later on..)
	// 1st arg is always flow address' 2nd is a UFix64
	var addrs []flow.Address
	var intVal float64
	var finalAmount cadence.UFix64

	if len(args) != 0 {
		for _, arg := range args {
			switch arg := arg.(type) {

			case flow.Address:
				addrs = append(addrs, arg)

			case int:
				intVal = float64(arg)
			}
		}
		newVal := fmt.Sprintf("%.2f", intVal)
		amount, err := cadence.NewUFix64(newVal)
		generateAccountKeys.Except(err)
		finalAmount = amount

	}

	flowSigner.CreateAndSignFlowTransaction(addrs, filePath, finalAmount, transactionSigner)

}
