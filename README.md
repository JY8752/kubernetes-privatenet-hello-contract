# kubernetes-privatenet-hello-contract
kubernetes上に構築したブロックチェーンにデプロイするコントラクト

```terminal
//deploy
source .env; forge create --rpc-url $RPC_URL --private-key $PRIVATE_KEY src/Hello.sol:Hello
```
