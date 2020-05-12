# create the identitites for the orgs
chmod 777 /opt/trade/createIdentity.sh
/opt/trade/createIdentity.sh

# install and instantiate chaincode
peer chaincode install -p chaincodedev/chaincode/trade_workflow_v1 -n tw -v 0
peer chaincode instantiate -n tw -v 0 -c '{"Args":["init","Buyer0", "100000", "Buyer1", "10000", "Seller0","200000", "Seller1", "10000", "Middleman","200000", "Warehouse","50000", "Carrier", "10000"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller0
peer chaincode invoke -n tw -c '{"Args":["initItem", "laptop", "1200", "5"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["queryItems", "{\"selector\":{\"descriptionOfGoods\":\"laptop\"}}"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller1
peer chaincode invoke -n  tw -c '{"Args":["initItem", "smartphone", "800", "3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["queryItems", "{\"selector\":{\"descriptionOfGoods\":\"smartphone\"}}"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller0
peer chaincode invoke -n tw -c '{"Args":["updateItem", "seller0laptop", "7"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller1
peer chaincode invoke -n tw -c '{"Args":["updateItem", "seller1smartphone", "7"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller0
peer chaincode invoke -n tw -c '{"Args": ["requestAdvertisement", "adv-1", "middleman", "seller0laptop", "0.08"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller1
peer chaincode invoke -n tw -c '{"Args": ["requestAdvertisement", "adv-2", "middleman", "seller1smartphone", "0.05"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/middleman
peer chaincode invoke -n tw -c '{"Args": ["acceptAdvertisement", "adv-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptAdvertisement", "adv-2"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller0
peer chaincode invoke -n tw -c '{"Args": ["requestStorage", "store-1", "warehouse", "seller0laptop", "0.08"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/seller1
peer chaincode invoke -n tw -c '{"Args": ["requestStorage", "store-2", "warehouse", "seller1smartphone", "0.05"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/warehouse
peer chaincode invoke -n tw -c '{"Args": ["acceptStorage", "store-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptStorage", "store-2"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer0
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-1", "3", "laptop"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-2", "2", "smartphone"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer1
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-3", "1", "laptop"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-4", "3", "smartphone"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/middleman
peer chaincode invoke -n tw -c '{"Args": ["acceptTrade", "trade-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptTrade", "trade-2"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptTrade", "trade-3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["acceptTrade", "trade-4"]}' -C tradechannel
sleep 2

peer chaincode invoke -n tw -c '{"Args": ["requestPayment", "trade-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["requestPayment", "trade-2"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["requestPayment", "trade-3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["requestPayment", "trade-4"]}' -C tradechannel
sleep 2


export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer0
peer chaincode invoke -n tw -c '{"Args": ["makePayment", "trade-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-1","Buyer0OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-1", "Seller0OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-1", "middleman"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-1", "carrier"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-1", "warehouse"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["makePayment", "trade-2"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-2", "Buyer0OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-2", "Seller1OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-2", "middleman"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-2", "carrier"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-2", "warehouse"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer1
peer chaincode invoke -n tw -c '{"Args": ["makePayment", "trade-3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-3", "Buyer1OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-3", "Seller0OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-3", "middleman"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-3", "carrier"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-3", "warehouse"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["makePayment", "trade-4"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-4", "Buyer1OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-4", "Seller1OrgMSP"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-4", "middleman"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-4", "carrier"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getAccountBalance", "trade-4", "warehouse"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/carrier
peer chaincode invoke -n tw -c '{"Args":["prepareShipment","ship-1","seller0","seller0laptop","3","buyer0"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["prepareShipment","ship-2","seller1","seller1smartphone","2","buyer0"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["prepareShipment","ship-3","seller0","seller0laptop","1","buyer1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["prepareShipment","ship-4","seller1","seller1smartphone","3","buyer1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["deliverShipment", "ship-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["deliverShipment", "ship-2"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["deliverShipment", "ship-3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args":["deliverShipment", "ship-4"]}' -C tradechannel

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer0
peer chaincode invoke -n tw -c '{"Args": ["getShipmentStatus","ship-1"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getShipmentStatus","ship-2"]}' -C tradechannel
sleep 2

export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/buyer1
peer chaincode invoke -n tw -c '{"Args": ["getShipmentStatus","ship-3"]}' -C tradechannel
sleep 2
peer chaincode invoke -n tw -c '{"Args": ["getShipmentStatus","ship-4"]}' -C tradechannel
sleep 2