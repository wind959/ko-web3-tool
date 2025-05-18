## 自用web3工具类

```
ko-web3-tool
├── go.mod          # 模块定义
├── internal        # 内部实现（禁止外部引用）
│   ├── crypto      # 加密安全模块
│   └── utils       # 基础工具
├── chains          # 区块链实现层
│   ├── eth         # 以太坊实现
│   ├── sol         # Solana实现
│   └── tron        # TRON实现
├── pkg             # 对外暴露接口
│   ├── types       # 公共数据类型
│   └── api         # 开发者接口
│   └── wallet      # 钱包
└── configs         # 配置管理
```