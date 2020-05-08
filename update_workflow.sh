docker cp ./src/trade-finance-logistics/chaincode/src/github.com/trade_workflow_v1/tradeWorkflow.go chaincode:/opt/gopath/src/chaincode_copy/trade_workflow_v1/
docker cp ./scripts/update_chaincode_workflow.sh chaincode:/opt/gopath/src/chaincode/
docker exec -ti chaincode bash
