# create the identitites for the orgs
/opt/trade/createIdentity.sh

# install and instantiate chaincode
peer chaincode install -p chaincodedev/chaincode/trade_workflow_v1 -n tw -v 0
peer chaincode instantiate -n tw -v 0 -c '{"Args":["init","LumberInc","LumberBank","100000","WoodenToys","ToyBank","200000","UniversalFreight","ForestryDepartment","Lender","5000000"]}' -C tradechannel
sleep 2

#importer requests trade from exporter
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/importer
peer chaincode invoke -n tw -c '{"Args":["requestTrade", "trade-1", "5000", "Wood for Toys"]}' -C tradechannel
sleep 2

#exporter accepts trade
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/exporter
peer chaincode invoke -n tw -c '{"Args":["acceptTrade", "trade-1"]}' -C tradechannel
sleep 2

#importer requests LC
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/importer
peer chaincode invoke -n tw -c '{"Args":["requestLC", "trade-1"]}' -C tradechannel
sleep 2

#importer issues LC
peer chaincode invoke -n tw -c '{"Args":["issueLC", "trade-1", "LC-1", "2020"]}' -C tradechannel
sleep 2

#exporter accepts LC
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/exporter
peer chaincode invoke -n tw -c '{"Args":["acceptLC", "trade-1"]}' -C tradechannel
sleep 2

#exporter requests EL
peer chaincode invoke -n tw -c '{"Args":["requestEL", "trade-1"]}' -C tradechannel
sleep 2

#regulatory authority issues EL
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/regulator
peer chaincode invoke -n tw -c '{"Args":["issueEL", "trade-1", "LC-1", "2021"]}' -C tradechannel
sleep 2

#exporter prepares shipment
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/exporter
peer chaincode invoke -n tw -c '{"Args":["prepareShipment", "trade-1"]}' -C tradechannel
sleep 2

#carrier accepts shipment and bill of lading
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/carrier
peer chaincode invoke -n tw -c '{"Args":["acceptShipmentAndIssueBL", "trade-1","BL-1","2020","new york", "new jersey"]}' -C tradechannel
sleep 2

#exporter requests payment from lender
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/exporter
peer chaincode invoke -n tw -c '{"Args":["requestPayment", "trade-1"]}' -C tradechannel
sleep 2

#lender requests accepted LC from exporter
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/lender
peer chaincode invoke -n tw -c '{"Args":["requestAcceptedLC", "trade-1"]}' -C tradechannel
sleep 2

#lender issues accepted LC to exporter
peer chaincode invoke -n tw -c '{"Args":["issueAcceptedLC", "trade-1","AcceptedLC-1","2020"]}' -C tradechannel
sleep 2

#exporter accepts accepted LC from lender
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/exporter
peer chaincode invoke -n tw -c '{"Args":["acceptAcceptedLC", "trade-1"]}' -C tradechannel
sleep 2

#lender makes payment to exporter
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/lender
peer chaincode invoke -n tw -c '{"Args":["makePayment", "trade-1", "exporter"]}' -C tradechannel
peer chaincode invoke -n tw -c '{"Args":["getAccountBalance", "trade-1", "lender"]}' -C tradechannel
peer chaincode invoke -n tw -c '{"Args":["getAccountBalance", "trade-1", "exporter"]}' -C tradechannel
sleep 2

#lender requests payment from importer in 60 days
peer chaincode invoke -n tw -c '{"Args":["requestPayment", "trade-1"]}' -C tradechannel
sleep 2

#importer makes payment to lender
export CORE_PEER_MSPCONFIGPATH=/root/.fabric-ca-client/importer
peer chaincode invoke -n tw -c '{"Args":["makePayment", "trade-1", "lender"]}' -C tradechannel
peer chaincode invoke -n tw -c '{"Args":["getAccountBalance", "trade-1", "importer"]}' -C tradechannel
peer chaincode invoke -n tw -c '{"Args":["getAccountBalance", "trade-1", "lender"]}' -C tradechannel
