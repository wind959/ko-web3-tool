package sol

import (
	"errors"
	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
	"github.com/wind959/ko-web3-tool/pkg/constant"
)

type SolanaAccount struct {
	Address    string
	PrivateKey string
}

// GetSolanaAccountPrivateKey 通过mnemonic获取solana私钥
func GetSolanaAccountPrivateKey(mnemonic string) (*SolanaAccount, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}
	seed := bip39.NewSeed(mnemonic, "")
	derivedKey, err := hdwallet.Derived(constant.SOL_DERIVATION_PATH, seed)
	if err != nil {
		return nil, err
	}
	account, err := types.AccountFromSeed(derivedKey.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &SolanaAccount{
		Address:    account.PublicKey.ToBase58(),
		PrivateKey: base58.Encode(account.PrivateKey),
	}, nil
}
