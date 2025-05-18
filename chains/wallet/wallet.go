package wallet

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/wind959/ko-web3-tool/chains/eth"
	"github.com/wind959/ko-web3-tool/chains/sol"
	"github.com/wind959/ko-web3-tool/chains/trc"
)

type Web3Wallet struct {
	Mnemonic      string // 助记词
	EthAddress    string // ethereum 地址
	EthPrivateKey string // ethereum 私钥
	SolAddress    string // solana 地址
	SolPrivateKey string // solana 私钥
	TrcAddress    string // tron 地址
	TrcPrivateKey string // tron 私钥
}

// CreateWeb3Wallet 生成 web3 钱包
func CreateWeb3Wallet(c int) []Web3Wallet {
	var wallets []Web3Wallet
	if c == 0 {
		c = 1
	}
	for i := 0; i < c; i++ {
		wallet, err := genBatchWallet()
		if err != nil {
			panic(err)
		}
		wallets = append(wallets, *wallet)
	}
	return wallets
}

func genBatchWallet() (*Web3Wallet, error) {
	// 生成助记词
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	ethAcc, _ := eth.GetEthereumAccount(mnemonic)
	solAcc, _ := sol.GetSolanaAccountPrivateKey(mnemonic)
	trcAcc, _ := trc.GetTronAccountPrivateKey(mnemonic)
	return &Web3Wallet{
		Mnemonic:      mnemonic,
		EthAddress:    ethAcc.Address,
		EthPrivateKey: ethAcc.PrivateKey,
		SolAddress:    solAcc.Address,
		SolPrivateKey: solAcc.PrivateKey,
		TrcAddress:    trcAcc.Address,
		TrcPrivateKey: trcAcc.PrivateKey,
	}, nil
}
