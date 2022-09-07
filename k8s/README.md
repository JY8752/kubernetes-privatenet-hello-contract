# k8sによるプライベートチェーン構築

## docker

```terminal
//build
docker image build -t private-net .

//run
docker run -it private-net -n private-net -p 8545:8545 -p 8546:8546 -p 30303:30303

//attach
docker exec -it private-net sh
geth attach http://localhost:8545

```

## kind

ローカルにマルチノード構築するために

```terminal
//install
brew install kind

//create
kind create cluster --config kind.yml

//delete
kind delete cluster --name blockchaine-cluster
kubectl config use-context docker-desktop //contextを戻す

//ローカルにkindで構築したマルチノード上にプライベートチェーンを展開する
kubectl apply -f private-net-deploy.yml

//attach
geth attach http://localhost:30045
```

## eks
amazon EKS上にプライベートチェーンを展開する

```terminal
//EKSクラスター作成
eksctl create cluster \                     
--name eks-private-net \
--nodegroup-name ng-private-net \
--node-type t3.large \
--nodes 3

//EKSのnode確認
kubectl get node -o json | jq ".items[]|{name: .metadata.name, ExternalIP: .status.addresses[1].address}"
{
  "name": "ip-192-168-20-197.ap-northeast-1.compute.internal",
  "ExternalIP": "13.114.138.50"
}
{
  "name": "ip-192-168-55-191.ap-northeast-1.compute.internal",
  "ExternalIP": "35.79.222.168"
}
{
  "name": "ip-192-168-87-66.ap-northeast-1.compute.internal",
  "ExternalIP": "3.112.83.0"
}

//apply
kubectl apply -f private-net-deploy.yml

//security group
aws ec2 authorize-security-group-ingress \
    --group-id sg-0b543374e4be3930d \
    --ip-permissions IpProtocol=tcp,FromPort=30045,ToPort=30045,IpRanges='[{CidrIp=0.0.0.0/0,Description="allow geth attach"}]'

aws ec2 authorize-security-group-ingress \
    --group-id sg-0b543374e4be3930d \
    --ip-permissions IpProtocol=tcp,FromPort=30046,ToPort=30046,IpRanges='[{CidrIp=0.0.0.0/0,Description="allow geth attach"}]'

//attach
geth attach http://13.114.138.50:30045

> Welcome to the Geth JavaScript console!

//contract deploy
source .env.prod; forge create --rpc-url $RPC_URL --private-key $PRIVATE_KEY src/Hello.sol:Hello --legacy

//request
ENVIROMENT=prod go run main.go

> Hello World!!

```

