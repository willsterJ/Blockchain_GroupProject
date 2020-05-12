cd $GOPATH/src/trade-finance-logistics/network/

echo -e "\nNow bringing the network down and back up again\n"
echo "Y" | ./trade.sh down
echo "Y" | ./trade.sh down -d true
echo -e "\nThere may be 'orphans' relating to exportingEntityOrg', so doing another sweep of the containers"
CONTAINER_IDS=$(docker ps -aq)
if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" == " " ]; then
  echo "---- No orphan containers were found, and proceeding to ./trade.sh clean ----"
else
  echo -e "\nFound additional containers " $CONTAINER_IDS
  echo -e "\nRemoving the found containers"
  docker rm -f $CONTAINER_IDS
  echo -e "\nInvoking ./trade.sh down for a second time"
  echo "Y" | ./trade.sh down -d true
  echo -e "\nCalling docker ps -a to see if we got rid of all the containers finally\n"
  docker ps -a
  echo -e "\nBrought the network down and there should be no active containers now"
fi

VOLUME_IDS=$(docker volume ls -q)
if [ -z "$VOLUME_IDS" -o "$VOLUME_IDS" == " " ]; then
  echo "---- No orphan volumes were found, and proceeding to ./trade.sh clean ----"
else
  echo -e "\nFound additional volumes " $VOLUME_IDS
  echo -e "\nRemoving found volumes"
  docker volume rm -f $VOLUME_IDS
  echo -e "\nCalling docker volume ls to see if we got rid of all the containers finally\n"
  docker volume ls
  echo -e "\nBrought the network down and there should be no active containers or volumes now"
fi

echo "Y" | ./trade.sh clean
echo "Cleaned/purged the network"
echo "Y" | ./trade.sh up -d true
echo -e "\nLaunched the chaincode and cli containers among others."
echo -e "\nWaiting for all containers to launch..."

sleep 15

cd $GOPATH

docker cp ./scripts/dev_mode_cli_run.sh cli:/opt/gopath/src/chaincodedev/
docker cp ./scripts/dev_mode_test_run.sh cli:/opt/gopath/src/chaincodedev/

docker cp ./scripts/dev_mode_chaincode_run.sh chaincode:/opt/gopath/src/chaincode/

docker exec -ti chaincode bash

