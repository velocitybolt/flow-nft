package flowUtility

import (
	"context"
	"engram-flow-api-v1/pkg/blockchain/generateAccountKeys"
	"fmt"
	"time"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

// Some of these functions are from Go-SDK Example code
// https://github.com/onflow/flow-go-sdk/blob/master/examples/examples.go

func GetReferenceBlockId(flowClient *client.Client) flow.Identifier {
	block, err := flowClient.GetLatestBlock(context.Background(), true)
	generateAccountKeys.Except(err)

	return block.ID
}

func GetAddress(txIDHex string) string {
	ctx := context.Background()
	c, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)

	txID := flow.HexToID(txIDHex)
	result, err := c.GetTransactionResult(ctx, txID)
	generateAccountKeys.Except(err)

	var address flow.Address

	if result.Status == flow.TransactionStatusSealed {
		for _, event := range result.Events {
			if event.Type == flow.EventAccountCreated {
				accountCreatedEvent := flow.AccountCreatedEvent(event)
				address = accountCreatedEvent.Address()
			}
		}
	}

	return address.Hex()
}

func GetAddressFlow(txIDHex string) flow.Address {
	ctx := context.Background()
	c, err := client.New("127.0.0.1:3569", grpc.WithInsecure())
	generateAccountKeys.Except(err)

	txID := flow.HexToID(txIDHex)
	result, err := c.GetTransactionResult(ctx, txID)
	generateAccountKeys.Except(err)

	var address flow.Address

	if result.Status == flow.TransactionStatusSealed {
		for _, event := range result.Events {
			if event.Type == flow.EventAccountCreated {
				accountCreatedEvent := flow.AccountCreatedEvent(event)
				address = accountCreatedEvent.Address()
			}
		}
	}
	return address
}

func ConvertScript(scriptAddr flow.Address, filePath string) []byte {
	template := generateAccountKeys.ReadFile(filePath)
	return []byte(fmt.Sprintf(template, scriptAddr))
}

func ConvertScriptTwo(inputList []flow.Address, filePath string) []byte {
	if len(inputList) > 2 {
		template := generateAccountKeys.ReadFile(filePath)
		return []byte(fmt.Sprintf(template, inputList[1], inputList[2]))
	} else {
		template := generateAccountKeys.ReadFile(filePath)
		return []byte(fmt.Sprintf(template, inputList[0], inputList[1]))
	}
}

func WaitForSeal(ctx context.Context, c *client.Client, id flow.Identifier) *flow.TransactionResult {
	result, err := c.GetTransactionResult(ctx, id)
	generateAccountKeys.Except(err)

	fmt.Printf("Waiting for transaction %s to be sealed...\n", id)

	for result.Status != flow.TransactionStatusSealed {
		time.Sleep(time.Second)
		fmt.Print(".")
		result, err = c.GetTransactionResult(ctx, id)
		generateAccountKeys.Except(err)

	}

	fmt.Println()
	fmt.Printf("Transaction %s sealed\n", id)
	return result
}
