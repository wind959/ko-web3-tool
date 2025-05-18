package test

import (
	"github.com/wind959/ko-web3-tool/chains/eth"
	"github.com/wind959/ko-web3-tool/chains/sol"
	"github.com/wind959/ko-web3-tool/chains/trc"
	"github.com/wind959/ko-web3-tool/chains/wallet"
	"testing"
)

func TestWallet(t *testing.T) {
	// 通过 助记词获取 私匙
	// 助记词: churn chuckle civil first sea stumble sight cause cactus transfer gaze wing
	//ETH 地址: 0x7b5BbFd603051eA35ba8ef25045c0524bFCc3750
	//ETH 私钥: 0x6b59ebb6d70d81c0ba3f74eceb6e452c3dfc15ec9189c4ff8969c545549d1abc
	//SOL 地址: 7Ayiup9kB3q6NfJTRC8usZt1vQdL1BdFmAwTqLxcNXBR
	//SOL 私钥: 4HsdAtBQ6qDx1PcpFRLwNABowPWMg5j9qAwoPRMNnQXzU5EqgVWAr3CvwCfXJe9qoRZb6qJUahUcMnfyyQuozofH
	//TRC 地址：TQxfDXZtwSHrWccXHPmbJAXgj2nTcAxwms
	//TRC 私钥：0xa96c6427d7ad9f60cf925d29c80bc5175222345ea1501410f6ab6b7bf6f5fd84
	//mnemonic := "churn chuckle civil first sea stumble sight cause cactus transfer gaze wing"
	// =============================
	// 助记词：twelve start embark unveil spike venue sample relax fee cabbage cube reopen
	//ETH 地址: 0x2bAeDcaa0f14b4737FAFcF6d5158b52bDec17B12
	//ETH 私钥: 0x35de43c3d6855030c9102d9c618e0ff416d6d51f3e28f67de59a96c13a8bd116
	//SOL 地址: 2WjMrxTZ6zQssrRnsyviJrJtno6jfXT2qQi1hJuc1RxY
	//SOL 私钥: 2hUzt3JneWf3YxzE5EiLu6DTWQojXwudA8PU8LJp83P5zKAs3guVyxUUfZ5SBMbJ4vg342UpUxqcV9fMkVMPniHS
	//TRC 地址：TJFffGPTHQVWz5v3qv2uyAEmH5brK5bjYN
	//TRC 私钥：0x9f5941fdec730bb30cae7287a4a2d186e357f94ab8afa9cc9c00b6a82b1b50c7
	mnemonic := "hover order culture begin spell dress pink piece fun mouse limit october"

	ethAcc, err := eth.GetEthereumAccount(mnemonic)

	if err != nil {
		t.Error(err)
	}
	t.Log("eth 地址：", ethAcc.Address)
	t.Log("eth 私匙：", ethAcc.PrivateKey)
	solAcc, err := sol.GetSolanaAccountPrivateKey(mnemonic)
	if err != nil {
		t.Error(err)
	}
	t.Log("sol 地址：", solAcc.Address)
	t.Log("sol 私匙：", solAcc.PrivateKey)
	trcAcc, err := trc.GetTronAccountPrivateKey(mnemonic)
	if err != nil {
		t.Error(err)
	}
	t.Log("trc地址：", trcAcc.Address)
	t.Log("trc私匙：", trcAcc.PrivateKey)
}

// 批量生成
func TestWallet2(t *testing.T) {
	count := 50
	web3Wallet := wallet.CreateWeb3Wallet(count)
	for _, wallet := range web3Wallet {
		t.Log("===========================================")
		t.Log("助记词：", wallet.Mnemonic)
		t.Log("ETH 地址：", wallet.EthAddress)
		t.Log("ETH 私匙：", wallet.EthPrivateKey)
		t.Log("SOL 地址：", wallet.SolAddress)
		t.Log("SOL 私匙：", wallet.SolPrivateKey)
		t.Log("TRC 地址：", wallet.TrcAddress)
		t.Log("TRC 私匙：", wallet.TrcPrivateKey)
		t.Log("===========================================")
	}
}
