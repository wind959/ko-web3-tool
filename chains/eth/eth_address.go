package eth

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"github.com/wind959/ko-web3-tool/pkg/constant"
)

type EthereumAccount struct {
	Address    string
	PrivateKey string
}

// GetEthereumAccount 通过助记词获取evm地址和私钥
// EVM 大概包括（Base, BSC,Polygon,Arbitrum ... ）
func GetEthereumAccount(mnemonic string) (*EthereumAccount, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, errors.New("failed to create HD wallet")
	}
	path := hdwallet.MustParseDerivationPath(constant.ETH_DERIVATION_PATH)

	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, errors.New("failed to derive account")
	}
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, errors.New("failed to get private key")
	}
	pubKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &EthereumAccount{
		Address:    pubKey.Hex(),
		PrivateKey: fmt.Sprintf("0x%s", common.Bytes2Hex(crypto.FromECDSA(privateKey))),
	}, nil
}
