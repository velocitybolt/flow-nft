package blockchainRoutes

import (
	"engram-flow-api-v1/pkg/blockchain/cdcExecutor"
	"engram-flow-api-v1/pkg/blockchain/createAccount"
	"engram-flow-api-v1/pkg/blockchain/flowAccountInfo"
	"engram-flow-api-v1/pkg/blockchain/flowUtility"

	"github.com/onflow/flow-go-sdk"
)

func CreateAccountFlow() (string, flow.Address, *flowAccountInfo.FlowAccountInfo, string, error) {
	accOneFlowAddrStr, accFlowAddr, createdAccountInfo, createdAccTxID, err := createAccount.InitalizeUser()

	return accOneFlowAddrStr, accFlowAddr, createdAccountInfo, createdAccTxID, err
}

func DeployInitialContracts(accOneFlowAddr flow.Address, createdAccountInfo *flowAccountInfo.FlowAccountInfo) (string, string, flow.Address, error) {
	txIdFungibleToken := cdcExecutor.ContractCadence("../../pkg/blockchain/cadence/contracts/FungibleToken.cdc", "FungibleToken", "1")
	cdcExecutor.ContractCadence("../../pkg/blockchain/cadence/contracts/NonFungibleToken.cdc", "NonFungibleToken", "1")
	fungAcctId := flowUtility.GetAddressFlow(txIdFungibleToken)
	convScript := string(flowUtility.ConvertScript(fungAcctId, "../../pkg/blockchain/cadence/contracts/EngramToken.cdc"))
	txIdEngramToken := cdcExecutor.AddContractToAccount(accOneFlowAddr, convScript, "EngramToken", createdAccountInfo)

	return txIdEngramToken, txIdFungibleToken, fungAcctId, nil
}

func SetupFlowAccountVault(accOneFlowAddr flow.Address, fungAccId flow.Address, createdAccountInfo *flowAccountInfo.FlowAccountInfo) error {
	setupAccArgs := []interface{}{fungAccId, accOneFlowAddr}
	cdcExecutor.TransactionCadenceArgs("./cadence/transactions/EngramToken/setup_account.cdc", setupAccArgs, createdAccountInfo)

	return nil
}

func MintEngramToken(accOneFlowAddr flow.Address, amount int, fungAccId flow.Address, createdAccountInfo *flowAccountInfo.FlowAccountInfo) error {
	mintTokenArgs := []interface{}{accOneFlowAddr, amount, fungAccId, accOneFlowAddr}
	cdcExecutor.TransactionCadenceArgs("./cadence/transactions/EngramToken/mint_engram_token.cdc", mintTokenArgs, createdAccountInfo)

	return nil
}

func TransferEngramToken(accOneFlowAddr flow.Address, accTwoFlowAddr flow.Address, amount int, fungAccId flow.Address, createdAccountInfo *flowAccountInfo.FlowAccountInfo) error {
	transferTokenArgs := []interface{}{accTwoFlowAddr, amount, fungAccId, accOneFlowAddr}
	cdcExecutor.TransactionCadenceArgs("./cadence/transactions/EngramToken/transfer_engram_token.cdc", transferTokenArgs, createdAccountInfo)

	return nil

}
