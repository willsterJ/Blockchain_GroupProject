/*
 * Copyright 2018 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"crypto/x509"
)


func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, string, error) {
	var mspid string
	var err error
	var cert *x509.Certificate

	mspid, err = cid.GetMSPID(stub)
	if err != nil {
		fmt.Printf("Error getting MSP identity: %s\n", err.Error())
		return "", "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %s\n", err.Error())
		return "", "", err
	}

	return mspid, cert.Issuer.CommonName, nil
}

// For now, just hardcode an ACL
// We will support attribute checks in an upgrade

func authenticateBuyerOrg(mspID string, certCN string) bool {
	return ((mspID == "Buyer0OrgMSP") && (certCN == "ca.buyer0org.trade.com")) || ((mspID == "Buyer1OrgMSP") && (certCN == "ca.buyer1org.trade.com"))
}

func authenticateSellerOrg(mspID string, certCN string) bool {
	return ((mspID == "Seller0OrgMSP") && (certCN == "ca.seller0org.trade.com")) || ((mspID == "Selle1rOrgMSP") && (certCN == "ca.seller1org.trade.com"))
}

func authenticateMiddlemanOrg(mspID string, certCN string) bool {
	return (mspID == "MiddlemanOrgMSP") && (certCN == "ca.middlemanorg.trade.com")
}

func authenticateWarehouseOrg(mspID string, certCN string) bool {
	return (mspID == "WarehouseOrgMSP") && (certCN == "ca.warehouseorg.trade.com")
}

func authenticateCarrierOrg(mspID string, certCN string) bool {
	return (mspID == "CarrierOrgMSP") && (certCN == "ca.carrierorg.trade.com")
}

func getCustomAttribute(stub shim.ChaincodeStubInterface, attr string) (string, bool, error) {
	var value string
	var found bool
	var err error

	value, found, err = cid.GetAttributeValue(stub, attr)
	if err != nil {
		fmt.Printf("Error getting MSP identity: %s\n", err.Error())
		return "", found, err
	}

	return value, found, nil
}