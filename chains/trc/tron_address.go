package trc

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type TronAccount struct {
	Address    string
	PrivateKey string
}

// GetTronAccountPrivateKey 通过助记词获取 Tron 账户私钥
func GetTronAccountPrivateKey(mnemonic string) (*TronAccount, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}
	// 获取助记词对应的 seed
	seed := bip39.NewSeed(mnemonic, "")

	// BIP32 根节点
	masterKey, _ := bip32.NewMasterKey(seed)
	// 派生路径 m/44'/195'/0'/0/0
	// 注意：bip32 不支持 hardened key，需手动加偏移量
	purpose, _ := masterKey.NewChildKey(44 + bip32.FirstHardenedChild)
	coinType, _ := purpose.NewChildKey(195 + bip32.FirstHardenedChild)
	account, _ := coinType.NewChildKey(0 + bip32.FirstHardenedChild)
	change, _ := account.NewChildKey(0)
	addrIndex, _ := change.NewChildKey(0)
	privateKeyBytes := addrIndex.Key

	// 生成 TRON 地址
	privateKey, _ := cryptoToECDSA(privateKeyBytes)
	pubKey := privateKey.PublicKey
	trxAddress := address.PubkeyToAddress(pubKey)
	return &TronAccount{
		Address:    trxAddress.String(),
		PrivateKey: fmt.Sprintf("0x%x", privateKeyBytes),
	}, nil
}

func cryptoToECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.ToECDSA(d)
	return privateKey, err
}
