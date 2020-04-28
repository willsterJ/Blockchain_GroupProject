cd ./trade_workflow_v1/
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=tw:0 ./trade_workflow_v1 