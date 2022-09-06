#!bin/sh

geth --datadir /var/share/ethereum --nodiscover --maxpeers 0 \
init /var/share/ethereum/genesis.json \
&& \
geth --datadir /var/share/ethereum --networkid 15 \
--nodiscover --maxpeers 0 --mine --miner.threads 1 \
--http --http.addr "0.0.0.0" --http.corsdomain "*" \
--http.vhosts "*" --http.api "eth,web3,personal,net" \
--ipcpath /tmp/geth.ipc --ws --ws.addr "0.0.0.0" \
--ws.api "eth,web3,personal,net" --ws.origins "*" \
--unlock 0,1,2,3,4 --password /var/share/ethereum/password --allow-insecure-unlock
