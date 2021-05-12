# WelfareSystem

## 操作步骤
1.进入项目根目录，运行go run main.go启动服务

## 模块功能
该模块主要对应着DID Wallet的donate功能，接收用户的捐赠信息，对信息验签后继续校验ClaimID是否合法，在通过的情况下对交易信息存入mysql和区块链中（此处存储使用DIDHub模块）。
