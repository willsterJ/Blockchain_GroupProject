docker cp ./src/trade-finance-logistics/chaincode/src/github.com/trade_workflow_v1/. chaincode:/opt/gopath/src/chaincode_copy/trade_workflow_v1/
docker cp ./scripts/update_chaincode_workflow.sh chaincode:/opt/gopath/src/chaincode/
docker cp ./scripts/dev_mode_cli_run.sh cli:/opt/gopath/src/chaincodedev/
docker cp ./scripts/dev_mode_test_run.sh cli:/opt/gopath/src/chaincodedev/
docker exec -ti chaincode bash
