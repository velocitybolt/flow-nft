package internal

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var routes = Routes{
	Route{"Healthcheck", "GET", "/api/healthcheck", HealthcheckHandler},
	//=== Contracts ===
	Route{"CreateAndDeployMainContract", "POST", "/api/contract/deploy/{contractName:[a-zA-Z]+}", DeployContractHandler},

	//=== Engram Token ===
	Route{"CreateAccount", "POST", "/api/account", CreateAccountHandler},
	Route{"TransferEngramCoins", "POST", "/api/coin/transfer", TransferEngramTokenHandler},
	Route{"GetEngramTokenSupply", "GET", "/api/coin/supply", GetEngramTokenSupplyHandler},

	//=== User's Account ===
	Route{"GetUsersBalance", "GET", "/api/user/balance", GetUserEngramCoinsBalanceHandler},
	Route{"GetUsersNfts", "GET", "/api/user/nfts", GetUserNftsHandler},

	//=== Nfts ===
	Route{"CreateNft", "POST", "/api/nft", MintNftHandler},
	Route{"GetNft", "GET", "/api/nft/{address:0[xX][0-9a-fA-F]+}", MintNftHandler},
	Route{"GetNfts", "GET", "/api/nft", GetNftsHandler},
}
