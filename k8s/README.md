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
eksctl create cluster \                     
--name eks-private-net \
--nodegroup-name ng-private-net \
--node-type t3.large \
--nodes 3
```

