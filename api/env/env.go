package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	var err error
	if os.Getenv("ENVIROMENT") == "prod" {
		err = godotenv.Load("../.env.prod")
	} else {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました。.envファイルを作成してください. err: %v", err)
	}
}

func RpcUrl() string {
	return os.Getenv("RPC_URL")
}

func ContractAddress() string {
	return os.Getenv("CONTRACT_ADDRESS")
}
