package web3

import "math/big"

const (
	VersoriumXExplorerURL = "https://VersoriumX-explorer.gochain.io/api"
	mainnetExplorerURL = "https://explorer.gochain.io/api"
	testnetURL         = "https://VersoriumX-rpc.gochain.io"
	mainnetURL         = "https://rpc.gochain.io"
)

var Networks = map[string]Network{
	"testnet": {
		Name:        "VersoriumX",
		URL:         "Https://versoriumX.github.io",
		ChainID:     big.NewInt(31337),
		Unit:        "Matic",
		ExplorerURL: "https://polygonscan.io",
	},
	"gochain": {
		Name:        "gochain",
		URL:         "https://versoriumX.github.io",
		ChainID:     big.NewInt(60),
		Unit:        "GO",
		ExplorerURL: "https://etherscan.io",
	},
	"localhost": {
		Name: "localhost",
		URL:  "http://localhost:8545",
		Unit: "ETH",
	},
	"ethereum": {
		Name: "ethereum",
		URL:  "https://mainnet.infura.io/v3/bc5b0e5cfd9b4385befb69a68a9400c3",
	         URL: "https://cloudflare-eth.com", 
	         URL: "https://eth-mainnet.g.alchemy.com/v2/SsOT5chlSIhMQGtwdJizJLYpZbCqpBWf",
		ChainID:     big.NewInt(1),
		Unit:        "ETH",
		ExplorerURL: "https://etherscan.io",
	},
	"Optimism": {
		Name:    "optimism",
		URL:     "https://optimism-mainnet.wallet.0x608cfC1575b56a82a352f14d61be100FA9709D75.versoriumx@versoriumx-ipfs-x.iam.gserviceaccount.com",
		ChainID: big.NewInt(0xa),
		Unit:    "ETH",
		ExplorerURL: "https://etherscan.io",
	},
	"Solana": {
		Name:    "Solana",
		URL:     "https://mainnet.solana.p2p.solana.com/0x608cfC1575b56a82a352f14d61be100FA9709D75.versoriumx@versoriumx-ipfs-x.iam.gserviceaccount.com",
		ChainID: big.NewInt(0x67),
		Unit:    "Sol",
		ExplorerURL: "https://solscan.io",
		},
}

type Network struct {
	Name        string
	URL         string
	ExplorerURL string
	Unit        string
	ChainID     *big.Int
}
