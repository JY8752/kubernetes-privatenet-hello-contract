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
kind delete cluster
```
