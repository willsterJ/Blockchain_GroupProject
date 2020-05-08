# create the identitites for the orgs
chmod 777 /opt/trade/createIdentity.sh
/opt/trade/createIdentity.sh

# install and instantiate chaincode
peer chaincode install -p chaincodedev/chaincode/trade_workflow_v1 -n tw -v 0
peer chaincode instantiate -n tw -v 0 -c '{"Args":["init","Buyer", "100000","Seller","200000","Middleman","200000", "Warehouse","50000", "Carrier", "10000"]}' -C tradechannel
sleep 2

#importer requests trade from exporter
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller
peer chaincode invoke -n tw -c '{"Args":["initItem", "laptop", "1200", "20"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["queryItems", "{\"selector\":{\"descriptionOfGoods\":\"laptop\"}}"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["updateItem", "sellerlaptop", "2"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-1", "1", "laptop"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptTrade", "trade-1"]}' -C tradechannel
sleep 2