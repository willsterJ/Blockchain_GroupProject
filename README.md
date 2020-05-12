# Prerequisites:
If fabric and fabric-ca folders are present in ./src/github.com/hyperledger/ , cd to that path and delete the folders. Then run 

```
git clone https://github.com/hyperledger/fabric -b "release-1.4" --single-branch

git clone https://github.com/hyperledger/fabric-ca -b "release-1.4" --single-branch

cd ./fabric
make configtxgen cryptogen configtxlator
make docker

cd ..
cd ./fabric-ca
make docker
```

Inside ./src/trade-finance-logistics/middleware/ run the following:
```
npm install??? ADD HERE
```

# Instructions:
#### To run in dev mode:
```
./start_devmode_script.sh
```
Once inside the chaincode container:
```
./dev_mode_chaincode_run.sh
```
Open another terminal (or tab) and run:
```
docker exec -ti cli bash
./dev_mode_cli_run.sh
```

#### To run in production mode:
```
./Middleware_full_production_run.sh
```

# Notes:
We are modifying the file in tradeworkflow_v1 folder in ./src/trade-finance-logistics/chaincode/src/github.com/

#### In devmode:
The bash files to be run inside the chaincode and cli containers are found in ./scripts/
