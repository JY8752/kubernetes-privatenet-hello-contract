package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました。.envファイルを作成してください. err: %v", err)
	}
}

func RpcUrl() string {
	return os.Getenv("RPC_URL")
}

func ContractAddress() string {
	return os.Getenv("CONTRACT_ADDRESS")
}
