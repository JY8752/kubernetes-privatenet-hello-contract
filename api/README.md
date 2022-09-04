# Client

## setup

```terminal
go mod init JY8752/kubernetes-privatenet-hello-contract

//etherreum
go get -d github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/rpc@v1.10.21
go get github.com/ethereum/go-ethereum/accounts/keystore@v1.10.21

//abigen
forge clean
forge build --extra-output-files abi

abigen --abi Hello.abi.json --pkg hello --type hello --out hello/hello.go

//dotenv
go get github.com/joho/godotenv
```