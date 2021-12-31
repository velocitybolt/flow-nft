package internal

import (
	"encoding/json"

	"engram-flow-api-v1/cmd/blockchainRoutes"
	requests "engram-flow-api-v1/internal/models/requests"
	responses "engram-flow-api-v1/internal/models/responses"
	"engram-flow-api-v1/pkg/blockchain/createAccount"
	"engram-flow-api-v1/pkg/blockchain/flowUtility"
	"engram-flow-api-v1/pkg/health"
	"engram-flow-api-v1/pkg/status"
	"net/http"
	"strconv"

	blockchainModels "engram-flow-api-v1/internal/models/flowBlockchainModels"

	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// add this if in the handler func if you want to see the request body
// requestDump, err2 := httputil.DumpRequest(req, true)
// if err2 != nil {
// 	fmt.Println(err2)
// }
// fmt.Println(string(requestDump))

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, AppEnv)

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(appEnv AppEnv, fn func(http.ResponseWriter, *http.Request, AppEnv)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Terry Pratchett tribute
		w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		// return function with AppEnv
		fn(w, r, appEnv)
	}
}

func HealthcheckHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	check := health.Check{
		AppName: "engram-flow-api",
		Version: appEnv.Version,
		Time:    time.Now().Format("01-02-2006 15:04:05"),
	}
	appEnv.Render.JSON(w, http.StatusOK, check)
}

//=== Contracts ===
func DeployContractHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	//vars := mux.Vars(req)
	//contractName := vars["contractname"]
	//accOneFlowAddr := vars["accOneFlowAddr"]
	//createdAccountInfo := vars["createdAccountInfo"]

	_, accAddr, createdAccountInfo, _, _ := createAccount.InitalizeUser()

	txIdEngramToken, txIdFungibleToken, _, err := blockchainRoutes.DeployInitialContracts(accAddr, createdAccountInfo)

	fungTokenAcctIdStr := flowUtility.GetAddress(txIdFungibleToken)
	//change to return a list of blockchainModels.ContractTransaction, err
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "error in deploying contracts!",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusNotFound,
		}).Error("error in deploying contracts!")
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}

	contracts := []blockchainModels.ContractTransaction{
		{
			TransactionId: txIdEngramToken,
			ContractName:  "engramToken",
		},
		{
			TransactionId: txIdFungibleToken,
			ContractName:  "fungibleToken",
		},
		{
			AccountId:    fungTokenAcctIdStr,
			ContractName: "fungibleTokenAccountID",
		},
	}

	deployContractsResponse := responses.DeployMainContractsResponse{
		ContractsDeployed: contracts,
	}

	appEnv.Render.JSON(w, http.StatusOK, deployContractsResponse)
}

//=== Transfer Engram Token across users ===
func TransferEngramTokenHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var transferRequest requests.TransferEngramCoinsRequest
	err := decoder.Decode(&transferRequest)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed transfer request object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed transfer request object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	receiverPrebalance := 1000
	receiverPostbalance := receiverPrebalance - transferRequest.NumCoinsToTranser

	transferResponse := responses.TransferEngramCoinsResponse{
		OwnerAddress:        transferRequest.OwnerAddress,
		ReceiverAddress:     transferRequest.ReceiverAddress,
		NumCoinsToTranser:   transferRequest.NumCoinsToTranser,
		ReceiverPreBalance:  10000,
		ReceiverPostBalance: receiverPostbalance,
	}

	appEnv.Render.JSON(w, http.StatusCreated, transferResponse)

}

func GetEngramTokenSupplyHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	supplyResponse := responses.EngramCoinSupplyResponse{
		ContractAddress: "hex here",
		Supply:          1000,
	}
	appEnv.Render.JSON(w, http.StatusOK, supplyResponse)
}

//=== Mint an NFT in exchange for Engram Token ===
func MintNftHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var newNftR requests.NewNftRequest
	err := decoder.Decode(&newNftR)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed mint nft object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed mint nft object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	nftAddresses := []string{"address 1", "address 2", "address 3"}

	newNftResponse := responses.NewNftResponse{
		Id:                     newNftR.Id,
		UserName:               newNftR.UserName,
		MainOwnerAddress:       newNftR.MainOwnerAddress,
		UserAddress:            newNftR.UserAddress,
		TransactionId:          "some hex here",
		RemainingEngramBalance: 35345,
		NftAdresses:            nftAddresses,
	}

	appEnv.Render.JSON(w, http.StatusCreated, newNftResponse)
}

//=== Transfer a NFT from one account to another ===
func TransferNftHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	vars := mux.Vars(req)
	address := vars["address"]
	error := error(nil) //make call to the blockchain here
	if error != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "can't find that asset!",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusNotFound,
		}).Error("Can't find asset")
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}

	nftResponse := responses.NftResponse{
		ContractAddress: "hex here",
		AssetAddress:    address,
		NumberMinted:    4,
		CurrentOwner:    "hex here",
	}
	appEnv.Render.JSON(w, http.StatusNotFound, nftResponse)
}

//=== Return the Owner of an NFT ===
func GetNftsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	error := error(nil) //make call to the blockchain here
	if error != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "can't find that asset!",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusNotFound,
		}).Error("Can't find asset")
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}

	assets := []responses.NftResponse{
		responses.NftResponse{
			ContractAddress: "hex here 1",
			AssetAddress:    "hex here 1",
			NumberMinted:    4,
			CurrentOwner:    "hex here 1",
		},
		responses.NftResponse{
			ContractAddress: "hex here 2",
			AssetAddress:    "hex here 2",
			NumberMinted:    4,
			CurrentOwner:    "hex here 2",
		},
		responses.NftResponse{
			ContractAddress: "hex here 3",
			AssetAddress:    "hex here 3",
			NumberMinted:    4,
			CurrentOwner:    "hex here 3",
		},
	}

	nftsResponse := responses.NftsResponse{
		Assets: assets,
	}
	appEnv.Render.JSON(w, http.StatusOK, nftsResponse)
}

//=== Create Account on Flow Blockchain ===
func CreateAccountHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var newWalR requests.NewAccountRequest
	err := decoder.Decode(&newWalR)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed create wallet object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed create wallet object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	_, _, createdAccountInfo, createdAccTxID, err := blockchainRoutes.CreateAccountFlow()
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "error in creating account",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("error in creating account")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	// flow addr and accId is same thing
	convAddr := flowUtility.GetAddress(createdAccTxID)

	newAccountReq := responses.NewAccountResponse{
		UserName:           newWalR.UserName,
		Id:                 newWalR.Id,
		AccountId:          convAddr,
		TransactionId:      createdAccTxID,
		CreatedAccountInfo: createdAccountInfo,
	}

	appEnv.Render.JSON(w, http.StatusOK, newAccountReq)
}

//=== Check Users Engram Coin Balance ===
func GetUserEngramCoinsBalanceHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var balanceRequest requests.BalanceRequest
	err := decoder.Decode(&balanceRequest)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed balnce request object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed balnce request object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	balanceResponse := responses.BalanceResponse{
		Id:            balanceRequest.Id,
		UserName:      balanceRequest.UserName,
		WalletAddress: balanceRequest.WalletAddress,
		Balance:       1000,
	}
	appEnv.Render.JSON(w, http.StatusOK, balanceResponse)
}

func GetUserNftsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var nftsRequest requests.NftsRequest
	err := decoder.Decode(&nftsRequest)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed balnce request object",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusBadRequest,
		}).Error("malformed balnce request object")
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	assets := []responses.NftResponse{
		responses.NftResponse{
			ContractAddress: "hex here 1",
			AssetAddress:    "hex here 1",
			NumberMinted:    4,
			CurrentOwner:    "hex here 1",
		},
		responses.NftResponse{
			ContractAddress: "hex here 2",
			AssetAddress:    "hex here 2",
			NumberMinted:    4,
			CurrentOwner:    "hex here 2",
		},
		responses.NftResponse{
			ContractAddress: "hex here 3",
			AssetAddress:    "hex here 3",
			NumberMinted:    4,
			CurrentOwner:    "hex here 3",
		},
	}
	appEnv.Render.JSON(w, http.StatusOK, assets)
}
