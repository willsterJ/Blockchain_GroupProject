set -e

ORG_NAME="buyer0"

fabric-ca-client enroll -u http://admin:adminpw@ca:7054 --mspdir admin

fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=buyer0" -u http://ca:7054

fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}

mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="admin"
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="buyer1"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=buyer1" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="seller0"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=seller0" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="seller1"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=seller1" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts


ORG_NAME="middleman"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=middleman" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="carrier"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=carrier" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts

ORG_NAME="warehouse"
fabric-ca-client register --id.name ${ORG_NAME} --id.secret pwd1 --id.type user \
    --id.attrs "role=warehouse" -u http://ca:7054
fabric-ca-client enroll -u http://${ORG_NAME}:pwd1@ca:7054 \
    --enrollment.attrs "role,email:opt" --mspdir ${ORG_NAME}
mkdir ~/.fabric-ca-client/${ORG_NAME}/admincerts
cp -p ~/.fabric-ca-client/${ORG_NAME}/signcerts/* ~/.fabric-ca-client/${ORG_NAME}/admincerts
