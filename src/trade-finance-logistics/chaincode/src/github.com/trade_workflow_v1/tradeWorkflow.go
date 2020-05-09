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
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	//"math/rand"
	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// TradeWorkflowChaincode implementation
type TradeWorkflowChaincode struct {
	testMode bool
}

func (t *TradeWorkflowChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initializing Trade Workflow")
	_, args := stub.GetFunctionAndParameters()
	var err error

	// Upgrade Mode 1: leave ledger state as it was
	if len(args) == 0 {
		return shim.Success(nil)
	}

	// Upgrade mode 2: change all the names and account balances
	if len(args) != 10 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 10: {"+
			"Buyer, "+
			"Buyer's Account Balance, "+
			"Seller, "+
			"Seller's Account Balance, "+
			"Middleman, "+
			"Middleman's Account Balance, "+
			"Warehouse, "+
			"Warehouse's Account Balance, "+
			"Carrier"+
			"Carrier's Account Balance, "+
			"}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Type checks
	_, err = strconv.Atoi(string(args[1]))
	if err != nil {
		fmt.Printf("Buyer's account balance must be an integer. Found %s\n", args[1])
		return shim.Error(err.Error())
	}
	_, err = strconv.Atoi(string(args[3]))
	if err != nil {
		fmt.Printf("Seller's account balance must be an integer. Found %s\n", args[3])
		return shim.Error(err.Error())
	}

	fmt.Printf("Buyer: %s\n", args[0])
	fmt.Printf("Buyer's Account Balance: %s\n", args[1])
	fmt.Printf("Seller: %s\n", args[2])
	fmt.Printf("Seller's Account Balance: %s\n", args[3])
	fmt.Printf("Middleman: %s\n", args[4])
	fmt.Printf("Middleman's Account Balance: %s\n", args[5])
	fmt.Printf("Warehouse: %s\n", args[6])
	fmt.Printf("Warehouse's Account Balance: %s\n", args[7])
	fmt.Printf("Carrier: %s\n", args[8])
	fmt.Printf("Carrier's Account Balance: %s\n", args[9])

	// Map participant identities to their roles on the ledger
	roleKeys := []string{buyKey, buyBalKey, selKey, selBalKey, midKey, midBalKey, warKey, warBalKey, carKey, carBalKey}
	for i, roleKey := range roleKeys {
		err = stub.PutState(roleKey, []byte(args[i]))
		if err != nil {
			fmt.Errorf("Error recording key %s: %s\n", roleKey, err.Error())
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func (t *TradeWorkflowChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var creatorOrg, creatorCertIssuer string
	var err error

	fmt.Println("TradeWorkflow Invoke")

	if !t.testMode {
		creatorOrg, creatorCertIssuer, err = getTxCreatorInfo(stub)
		if err != nil {
			fmt.Errorf("Error extracting creator identity info: %s\n", err.Error())
			return shim.Error(err.Error())
		}
		fmt.Printf("TradeWorkflow Invoke by '%s', '%s'\n", creatorOrg, creatorCertIssuer)
	}

	function, args := stub.GetFunctionAndParameters()
	if function == "initItem" {
		// Initialize item object
		return t.initItem(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "queryItems" {
		// query call to couchDB
		return t.queryItems(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "updateItem" {
		// update item
		return t.updateItem(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "requestTrade" {
		// Importer requests a trade
		return t.requestTrade(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "acceptTrade" {
		// Exporter accepts a trade
		return t.acceptTrade(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "acceptShipmentAndIssueBL" {
		// Carrier validates the shipment and issues a B/L
		return t.acceptShipmentAndIssueBL(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "requestPayment" {
		// Exporter's Bank requests a payment
		return t.requestPayment(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "makePayment" {
		// Importer's Bank makes a payment
		return t.makePayment(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "getTradeStatus" {
		// Get status of trade agreement
		return t.getTradeStatus(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "getShipmentLocation" {
		// Get the shipment location
		return t.getShipmentLocation(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "getBillOfLading" {
		// Get the bill of lading
		return t.getBillOfLading(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "getAccountBalance" {
		// Get account balance: Exporter/Importer
		return t.getAccountBalance(stub, creatorOrg, creatorCertIssuer, args)
		/*} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, creatorOrg, creatorCertIssuer, args)*/
	} else if function == "requestAdvertisement" {
		return t.requestAdvertisement(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "acceptAdvertisement" {
		return t.acceptAdvertisement(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "requestStorage" {
		return t.requestStorage(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "acceptStorage" {
		return t.acceptStorage(stub, creatorOrg, creatorCertIssuer, args)
	} else if function == "prepareShipment" {
		return t.prepareShipment(stub, creatorOrg, creatorCertIssuer, args)
	}

	return shim.Error("Invalid invoke function name")
}

// Inititialize an item object
func (t *TradeWorkflowChaincode) initItem(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var itemEntry *ItemEntry
	var itemEntryBytes []byte
	var itemId string
	var itemName string
	var price float64
	var count int
	var err error

	if !t.testMode && !authenticateSellerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}
	// Only in testmode, retrieve the role attribute from org
	role, found, err1 := getCustomAttribute(stub, "role")
	if t.testMode && found && err1 == nil && role != "seller" {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}
	if len(args) != 3 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 3: {Description of goods, Price, Count}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// retrieve input data
	itemName = args[0]
	price, err = strconv.ParseFloat(string(args[1]), 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	count, err = strconv.Atoi(string(args[2]))
	if err != nil {
		return shim.Error(err.Error())
	}
	if t.testMode {
		itemId = string(role + "" + itemName)
	}

	// check if item already exists
	itemEntryBytes, err = stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get item: " + err.Error())
	} else if itemEntryBytes != nil {
		fmt.Println("This item already exists: " + itemId + " , name:" + itemName)
		return shim.Error("This item already exists: " + itemId + " , name:" + itemName)
	}
	// Create item object and marshal to JSON  // TODO change creatorOrg to something unique!
	itemEntry = &ItemEntry{itemId, itemName, creatorOrg, "", "", price, count}
	itemEntryBytes, err = json.Marshal(itemEntry)
	if err != nil {
		return shim.Error("Error marshaling ItemEntry structure")
	}

	// Save item to state
	err = stub.PutState(itemId, itemEntryBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Use query strings to access DB
// ex: "{\"selector\":{\"owner\":\"tom\"}}"
func (t *TradeWorkflowChaincode) queryItems(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// executes passed in query string
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

// Update an item using unique key
func (t *TradeWorkflowChaincode) updateItem(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var itemEntry *ItemEntry // item to be changed
	var itemEntryBytes []byte
	var err error

	// Access control: Only an Seller Org member can invoke this transaction
	if !t.testMode && !authenticateSellerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}
	// Only in testmode, retrieve the role attribute from org
	role, found, err1 := getCustomAttribute(stub, "role")
	if t.testMode && found && err1 == nil && role != "seller" {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}
	if len(args) != 2 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {ItemID, update_count}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	itemId := args[0]
	updateCount, err := strconv.Atoi(args[1])

	itemEntryBytes, err = stub.GetState(itemId)
	if err != nil {
		return shim.Error("Failed to get item:" + err.Error())
	} else if itemEntryBytes == nil {
		return shim.Error("Item does not exist")
	}

	err = json.Unmarshal(itemEntryBytes, &itemEntry)
	if err != nil {
		return shim.Error(err.Error())
	}

	itemEntry.Count += updateCount

	itemEntryBytes, err = json.Marshal(itemEntry)
	if err != nil {
		return shim.Error("Error marshaling ItemEntry structure")
	}
	err = stub.PutState(itemId, itemEntryBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end updateItem (success)")
	return shim.Success(nil)

}

// Request a trade agreement
func (t *TradeWorkflowChaincode) requestTrade(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var tradeKey string
	var tradeAgreement *TradeAgreement
	var tradeAgreementBytes []byte
	var amount int
	var err error

	// ADD TRADELIMIT RETRIEVAL HERE

	// Access control: Only an Buyer Org member can invoke this transaction
	if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Buyer Org. Access denied.")
	}

	if len(args) != 3 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 3: {ID, Amount, Description of Goods}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	amount, err = strconv.Atoi(string(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}

	// ADD TRADE LIMIT CHECK HERE

	tradeAgreement = &TradeAgreement{amount, args[2], REQUESTED, 0}
	tradeAgreementBytes, err = json.Marshal(tradeAgreement)
	if err != nil {
		return shim.Error("Error marshaling trade agreement structure")
	}

	// Write the state to the ledger
	tradeKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(tradeKey, tradeAgreementBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Trade %s request recorded\n", args[0])
	return shim.Success(nil)
}

// Accept a trade agreement
func (t *TradeWorkflowChaincode) acceptTrade(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var tradeKey string
	var tradeAgreement *TradeAgreement
	var tradeAgreementBytes []byte
	var err error

	// Access control: Only an Exporting Entity Org member can invoke this transaction
	if !t.testMode && !authenticateMiddlemanOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Middleman Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	tradeKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	tradeAgreementBytes, err = stub.GetState(tradeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(tradeAgreementBytes) == 0 {
		err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
	if err != nil {
		return shim.Error(err.Error())
	}

	// var queryString string
	// queryString = "{\"selector\":{\"descriptionOfGoods\":\""+tradeAgreement.DescriptionOfGoods+"\"}}"
	// resultsIterator, err := stub.GetQueryResult(queryString)
	// if err != nil {
	// 	return shim.Error("Failed to get query result.")
	// }
	// defer resultsIterator.Close()
	// for resultsIterator.HasNext() {
	// 	queryResponse, err := resultsIterator.Next()
	// 	if err != nil {
	// 		return shim.Error("Failed to do whatever.")
	// 	}
	// 	// t := fmt.Sprintf("%T", queryResponse)
	// 	x := json.Unmarshal(queryResponse.GetValue(), &x)

	// 	// fmt.Println(t)
	// }

	if tradeAgreement.Status == ACCEPTED {
		fmt.Printf("Trade %s already accepted", args[0])
	} else {
		tradeAgreement.Status = ACCEPTED
		tradeAgreementBytes, err = json.Marshal(tradeAgreement)
		if err != nil {
			return shim.Error("Error marshaling trade agreement structure")
		}
		// Write the state to the ledger
		err = stub.PutState(tradeKey, tradeAgreementBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Printf("Trade %s acceptance recorded\n", args[0])

	return shim.Success(nil)
}

/*
// Prepare a shipment; preparation is indicated by setting the location as SOURCE
func (t *TradeWorkflowChaincode) prepareShipment(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var elKey, shipmentLocationKey string
	var shipmentLocationBytes, exportLicenseBytes []byte
	var exportLicense *ExportLicense
	var err error

	// Access control: Only an Exporting Entity Org member can invoke this transaction
	if !t.testMode && !authenticateExportingEntityOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Exporting Entity Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {Trade ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Lookup shipment location from the ledger
	shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	shipmentLocationBytes, err = stub.GetState(shipmentLocationKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(shipmentLocationBytes) != 0 {
		if string(shipmentLocationBytes) == SOURCE {
			fmt.Printf("Shipment for trade %s has already been prepared", args[0])
			return shim.Success(nil)
		} else {
			fmt.Printf("Shipment for trade %s has passed the preparation stage", args[0])
			return shim.Error("Shipment past the preparation stage")
		}
	}

	// Lookup E/L from the ledger
	elKey, err = getELKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	exportLicenseBytes, err = stub.GetState(elKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(exportLicenseBytes, &exportLicense)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Verify that the E/L has already been issued
	if exportLicense.Status != ISSUED {
		fmt.Printf("E/L for trade %s has not been issued", args[0])
		return shim.Error("E/L not issued yet")
	}

	shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(shipmentLocationKey, []byte(SOURCE))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Shipment preparation for trade %s recorded\n", args[0])

	return shim.Success(nil)
}
*/
// Accept a shipment and issue a B/L
func (t *TradeWorkflowChaincode) acceptShipmentAndIssueBL(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	/*	var shipmentLocationKey, blKey, tradeKey string
		var shipmentLocationBytes, tradeAgreementBytes, billOfLadingBytes, exporterBytes, carrierBytes, beneficiaryBytes []byte
		var billOfLading *BillOfLading
		var tradeAgreement *TradeAgreement
		var err error

		// Access control: Only an Carrier Org member can invoke this transaction
		if !t.testMode && !authenticateCarrierOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Carrier Org. Access denied.")
		}

		if len(args) != 5 {
			err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 5: {Trade ID, B/L ID, Expiration Date, Source Port, Destination Port}. Found %d", len(args)))
			return shim.Error(err.Error())
		}

		// Lookup shipment location from the ledger
		shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}

		shipmentLocationBytes, err = stub.GetState(shipmentLocationKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		if len(shipmentLocationBytes) == 0 {
			fmt.Printf("Shipment for trade %s has not been prepared yet", args[0])
			return shim.Error("Shipment not prepared yet")
		}
		if string(shipmentLocationBytes) != SOURCE {
			fmt.Printf("Shipment for trade %s has passed the preparation stage", args[0])
			return shim.Error("Shipment past the preparation stage")
		}

		// Lookup trade agreement from the ledger
		tradeKey, err = getTradeKey(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		tradeAgreementBytes, err = stub.GetState(tradeKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		if len(tradeAgreementBytes) == 0 {
			err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
			return shim.Error(err.Error())
		}

		// Unmarshal the JSON
		err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Lookup exporter
		exporterBytes, err = stub.GetState(expKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Lookup carrier
		carrierBytes, err = stub.GetState(carKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Lookup importer's bank (beneficiary of the title to goods after paymen tis made)
		beneficiaryBytes, err = stub.GetState(ibKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Create and record a B/L
		billOfLading = &BillOfLading{args[1], args[2], string(exporterBytes), string(carrierBytes), tradeAgreement.DescriptionOfGoods,
					     tradeAgreement.Amount, string(beneficiaryBytes), args[3], args[4]}
		billOfLadingBytes, err = json.Marshal(billOfLading)
		if err != nil {
			return shim.Error("Error marshaling bill of lading structure")
		}

		// Write the state to the ledger
		blKey, err = getBLKey(stub, args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		err = stub.PutState(blKey, billOfLadingBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf("Bill of Lading for trade %s recorded\n", args[0])
	*/
	return shim.Success(nil)
}

// Request a payment
func (t *TradeWorkflowChaincode) requestPayment(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var shipmentLocationKey, paymentKey, tradeKey string
	var shipmentLocationBytes, paymentBytes, tradeAgreementBytes []byte
	var tradeAgreement *TradeAgreement
	var err error

	// Access control: Only an Exporting Entity Org member can invoke this transaction
	if !t.testMode && !authenticateMiddlemanOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Middleman Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {Trade ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Lookup trade agreement from the ledger
	tradeKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	tradeAgreementBytes, err = stub.GetState(tradeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(tradeAgreementBytes) == 0 {
		err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Lookup shipment location from the ledger
	shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	shipmentLocationBytes, err = stub.GetState(shipmentLocationKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	// if len(shipmentLocationBytes) == 0 {
	// 	fmt.Printf("Shipment for trade %s has not been prepared yet", args[0])
	// 	return shim.Error("Shipment not prepared yet")
	// }

	// Check if there's already a pending payment request
	paymentKey, err = getPaymentKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	paymentBytes, err = stub.GetState(paymentKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(paymentBytes) != 0 { // The value doesn't matter as this is a temporary key used as a marker
		fmt.Printf("Payment request already pending for trade %s\n", args[0])
	} else {
		// Check what has been paid up to this point
		fmt.Printf("Amount paid thus far for trade %s = %d; total required = %d\n", args[0], tradeAgreement.Payment, tradeAgreement.Amount)
		if tradeAgreement.Amount == tradeAgreement.Payment { // Payment has already been settled
			fmt.Printf("Payment already settled for trade %s\n", args[0])
			return shim.Error("Payment already settled")
		}
		if string(shipmentLocationBytes) == SOURCE && tradeAgreement.Payment != 0 { // Suppress duplicate requests for partial payment
			fmt.Printf("Partial payment already made for trade %s\n", args[0])
			return shim.Error("Partial payment already made")
		}

		// Record request on ledger
		err = stub.PutState(paymentKey, []byte(REQUESTED))
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf("Payment request for trade %s recorded\n", args[0])
	}
	return shim.Success(nil)
}

// Make a payment
func (t *TradeWorkflowChaincode) makePayment(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var paymentKey, tradeKey string
	var paymentAmount, midBal, buyBal, selBal, carBal, warBal float64
	var paymentBytes, tradeAgreementBytes, buyBalBytes, midBalBytes, selBalBytes, carBalBytes, warBalBytes []byte
	var tradeAgreement *TradeAgreement
	var err error

	// Access control: Only an Importer Org member can invoke this transaction
	if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Buyer Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {Trade ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Check if there's already a pending payment request
	paymentKey, err = getPaymentKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	paymentBytes, err = stub.GetState(paymentKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(paymentBytes) == 0 {
		fmt.Printf("No payment request found for trade %s", args[0])
		return shim.Error("No payment request found")
	}

	// Lookup trade agreement from the ledger
	tradeKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	tradeAgreementBytes, err = stub.GetState(tradeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(tradeAgreementBytes) == 0 {
		err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Lookup shipment location from the ledger
	// shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }

	// shipmentLocationBytes, err = stub.GetState(shipmentLocationKey)
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }

	// if len(shipmentLocationBytes) == 0 {
	// 	fmt.Printf("Shipment for trade %s has not been prepared yet", args[0])
	// 	return shim.Error("Shipment not prepared yet")
	// }

	// Lookup account balances
	midBalBytes, err = stub.GetState(midBalKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	midBal, err = strconv.ParseFloat(string(midBalBytes), 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	buyBalBytes, err = stub.GetState(buyBalKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	buyBal, err = strconv.ParseFloat(string(buyBalBytes), 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	selBalBytes, err = stub.GetState(selBalKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	selBal, err = strconv.ParseFloat(string(selBalBytes), 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	warBalBytes, err = stub.GetState(warBalKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	warBal, err = strconv.ParseFloat(string(warBalBytes), 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	carBalBytes, err = stub.GetState(carBalKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	carBal, err = strconv.ParseFloat(string(carBalBytes), 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Record transfer of funds
	// if string(shipmentLocationBytes) == SOURCE {
	// 	paymentAmount = tradeAgreement.Amount/2
	// } else {
	// 	paymentAmount = tradeAgreement.Amount - tradeAgreement.Payment
	// }
	var midRate, selRate, warRate, carRate float64
	midRate = 0.1
	selRate = 0.85
	warRate = 0.025
	carRate = 0.025
	paymentAmount = float64(tradeAgreement.Amount)
	tradeAgreement.Payment += int(paymentAmount)
	midBal += paymentAmount * midRate
	selBal += paymentAmount * selRate
	warBal += paymentAmount * warRate
	carBal += paymentAmount * carRate

	if buyBal < paymentAmount {
		fmt.Printf("Buyer's bank balance %d is insufficient to cover payment amount %d\n", buyBal, paymentAmount)
	}
	buyBal -= paymentAmount

	// Update ledger state
	tradeAgreementBytes, err = json.Marshal(tradeAgreement)
	if err != nil {
		return shim.Error("Error marshaling trade agreement structure")
	}
	err = stub.PutState(tradeKey, tradeAgreementBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(midBalKey, []byte(fmt.Sprintf("%f", midBal)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(buyBalKey, []byte(fmt.Sprintf("%f", buyBal)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(selBalKey, []byte(fmt.Sprintf("%f", selBal)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(warBalKey, []byte(fmt.Sprintf("%f", warBal)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(carBalKey, []byte(fmt.Sprintf("%f", carBal)))
	if err != nil {
		return shim.Error(err.Error())
	}

	// Delete request key from ledger
	err = stub.DelState(paymentKey)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Failed to delete payment request from ledger")
	}

	return shim.Success(nil)
}

// Update shipment location; we will only allow SOURCE and DESTINATION as valid locations for this contract
func (t *TradeWorkflowChaincode) updateShipmentLocation(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var shipmentLocationKey string
	var shipmentLocationBytes []byte
	var err error

	// Access control: Only a Carrier Org member can invoke this transaction
	if !t.testMode && !authenticateCarrierOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Carrier Org. Access denied.")
	}

	if len(args) != 2 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {Trade ID, Location}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Lookup shipment location from the ledger
	shipmentLocationKey, err = getShipmentLocationKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	shipmentLocationBytes, err = stub.GetState(shipmentLocationKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	// if len(shipmentLocationBytes) == 0 {
	// 	fmt.Printf("Shipment for trade %s has not been prepared yet", args[0])
	// 	return shim.Error("Shipment not prepared yet")
	// }
	if string(shipmentLocationBytes) == args[1] {
		fmt.Printf("Shipment for trade %s is already in location %s", args[0], args[1])
	}

	// Write the state to the ledger
	err = stub.PutState(shipmentLocationKey, []byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Shipment location for trade %s recorded\n", args[0])

	return shim.Success(nil)
}

/*// Deletes an entity from state
func (t *TradeWorkflowChaincode) delete(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var key string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: <key name>")
	}

	key = args[0]

	// Delete the key from the state in ledger
	err = stub.DelState(key)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}*/

// Get current state of a trade agreement
func (t *TradeWorkflowChaincode) getTradeStatus(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var tradeKey, jsonResp string
	var tradeAgreement TradeAgreement
	var tradeAgreementBytes []byte
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: <trade ID>")
	}

	// Get the state from the ledger
	tradeKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	tradeAgreementBytes, err = stub.GetState(tradeKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + tradeKey + "\"}"
		return shim.Error(jsonResp)
	}

	if len(tradeAgreementBytes) == 0 {
		jsonResp = "{\"Error\":\"No record found for " + tradeKey + "\"}"
		return shim.Error(jsonResp)
	}

	// Unmarshal the JSON
	err = json.Unmarshal(tradeAgreementBytes, &tradeAgreement)
	if err != nil {
		return shim.Error(err.Error())
	}

	jsonResp = "{\"Status\":\"" + tradeAgreement.Status + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

// Get current location of a shipment
func (t *TradeWorkflowChaincode) getShipmentLocation(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var slKey, jsonResp string
	var shipmentLocationBytes []byte
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: <trade ID>")
	}

	// Get the state from the ledger
	slKey, err = getShipmentLocationKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	shipmentLocationBytes, err = stub.GetState(slKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + slKey + "\"}"
		return shim.Error(jsonResp)
	}

	if len(shipmentLocationBytes) == 0 {
		jsonResp = "{\"Error\":\"No record found for " + slKey + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Location\":\"" + string(shipmentLocationBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

// Get Bill of Lading
func (t *TradeWorkflowChaincode) getBillOfLading(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var blKey, jsonResp string
	var billOfLadingBytes []byte
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: <trade ID>")
	}

	// Get the state from the ledger
	blKey, err = getBLKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	billOfLadingBytes, err = stub.GetState(blKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + blKey + "\"}"
		return shim.Error(jsonResp)
	}

	if len(billOfLadingBytes) == 0 {
		jsonResp = "{\"Error\":\"No record found for " + blKey + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Printf("Query Response:%s\n", string(billOfLadingBytes))
	return shim.Success(billOfLadingBytes)
}

// Get current account balance for a given participant
func (t *TradeWorkflowChaincode) getAccountBalance(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var entity, balanceKey, jsonResp string
	var balanceBytes []byte
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: {Trade ID, Entity}")
	}

	entity = strings.ToLower(args[1])
	if entity == "seller" {
		// Access control: Only an Exporter or Exporting Entity Org member can invoke this transaction
		if !t.testMode && !authenticateSellerOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Seller Org. Access denied.")
		}
		balanceKey = selBalKey
	} else if entity == "buyer" {
		// Access control: Only an Importer Org member can invoke this transaction
		if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Buyer Org. Access denied.")
		}
		balanceKey = buyBalKey
	} else if entity == "middleman" {
		// Access control: Only an Importer Org member can invoke this transaction
		if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Middleman Org. Access denied.")
		}
		balanceKey = midBalKey
	} else if entity == "warehouse" {
		// Access control: Only an Importer Org member can invoke this transaction
		if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Warehouse Org. Access denied.")
		}
		balanceKey = warBalKey
	} else if entity == "carrier" {
		// Access control: Only an Importer Org member can invoke this transaction
		if !t.testMode && !authenticateBuyerOrg(creatorOrg, creatorCertIssuer) {
			return shim.Error("Caller not a member of Carrier Org. Access denied.")
		}
		balanceKey = carBalKey
	} else {
		err = errors.New(fmt.Sprintf("Invalid entity %s; Permissible values: {exporter, importer}", args[1]))
		return shim.Error(err.Error())
	}

	// Get the account balances from the ledger
	balanceBytes, err = stub.GetState(balanceKey)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + balanceKey + "\"}"
		return shim.Error(jsonResp)
	}

	if len(balanceBytes) == 0 {
		jsonResp = "{\"Error\":\"No record found for " + balanceKey + "\"}"
		return shim.Error(jsonResp)
	}
	jsonResp = "{\"Balance\":\"" + string(balanceBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

// Seller requests an advertisement deal with Middleman
func (t *TradeWorkflowChaincode) requestAdvertisement(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var contract *ContractSellerMiddleman
	var contractBytes []byte
	var fee float64
	var contractKey string
	var err error

	// Access control: Only an Buyer Org member can invoke this transaction
	if !t.testMode && !authenticateSellerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4: {Contract ID, Middleman ID, Item ID, Fee}")
	}

	fee, err = strconv.ParseFloat(string(args[3]), 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ADD TRADE LIMIT CHECK HERE

	contract = &ContractSellerMiddleman{args[0], args[1], args[2], fee, REQUESTED}
	contractBytes, err = json.Marshal(contract)
	if err != nil {
		return shim.Error("Error marshaling contract structure")
	}

	// Write the state to the ledger
	contractKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Advertisement Contract %s request recorded\n", args[0])
	return shim.Success(nil)
}

// Middleman accepts advertisement request from Buyer
func (t *TradeWorkflowChaincode) acceptAdvertisement(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var contractKey, itemId string
	var contract *ContractSellerMiddleman
	var itemEntry *ItemEntry
	var contractBytes, itemEntryBytes []byte
	var err error

	// Access control: Only an Exporting Entity Org member can invoke this transaction
	if !t.testMode && !authenticateMiddlemanOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Middleman Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Get the contract state from the ledger
	contractKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	contractBytes, err = stub.GetState(contractKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(contractBytes) == 0 {
		err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(contractBytes, &contract)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if contract had already been accepted.
	if contract.Status == ACCEPTED {
		fmt.Printf("Advertisement Contract %s already accepted", args[0])
	} else {
		// update status to ACCEPTED
		contract.Status = ACCEPTED
		contractBytes, err = json.Marshal(contract)
		if err != nil {
			return shim.Error("Error marshaling advertisement contract structure")
		}
		// Write the state to the ledger
		err = stub.PutState(contractKey, contractBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		// UPDATE ENTRY FIELD
		// Get the item entry in the database using the contract fields
		itemId = contract.ItemId
		itemEntryBytes, err = stub.GetState(itemId)
		if err != nil {
			return shim.Error("Failed to get item: " + err.Error())
		}

		err = json.Unmarshal(itemEntryBytes, &itemEntry)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Update the middleman field of the item entry
		itemEntry.Middleman = contract.MiddlemanId

		itemEntryBytes, err = json.Marshal(itemEntry)
		if err != nil {
			return shim.Error("Error marshaling ItemEntry structure")
		}

		// Save item to state
		err = stub.PutState(itemId, itemEntryBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Printf("Advertisement Contract %s acceptance recorded\n", args[0])

	return shim.Success(nil)
}

// Seller requests storage from Warehouse
func (t *TradeWorkflowChaincode) requestStorage(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var contract *ContractSellerWarehouse
	var contractBytes []byte
	var fee float64
	var contractKey string
	var err error

	// Access control: Only an Buyer Org member can invoke this transaction
	if !t.testMode && !authenticateSellerOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Seller Org. Access denied.")
	}

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4: {Contract ID, Warehouse ID, Item ID, Fee}")
	}
	fee, err = strconv.ParseFloat(string(args[3]), 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ADD TRADE LIMIT CHECK HERE

	contract = &ContractSellerWarehouse{args[0], args[1], args[2], fee, REQUESTED}
	contractBytes, err = json.Marshal(contract)
	if err != nil {
		return shim.Error("Error marshaling contract structure")
	}

	// Write the state to the ledger
	contractKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(contractKey, contractBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Storage Contract %s request recorded\n", args[0])
	return shim.Success(nil)
}

// Warehouse accepts storage request from Seller
func (t *TradeWorkflowChaincode) acceptStorage(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var contractKey, itemId string
	var contract *ContractSellerWarehouse
	var itemEntry *ItemEntry
	var contractBytes, itemEntryBytes []byte
	var err error

	// Access control: Only an Exporting Entity Org member can invoke this transaction
	if !t.testMode && !authenticateWarehouseOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Warehouse Org. Access denied.")
	}

	if len(args) != 1 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 1: {ID}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	// Get the state from the ledger
	contractKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	contractBytes, err = stub.GetState(contractKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if len(contractBytes) == 0 {
		err = errors.New(fmt.Sprintf("No record found for trade ID %s", args[0]))
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal(contractBytes, &contract)
	if err != nil {
		return shim.Error(err.Error())
	}

	if contract.Status == ACCEPTED {
		fmt.Printf("Storage Contract %s already accepted", args[0])
	} else {
		contract.Status = ACCEPTED
		contractBytes, err = json.Marshal(contract)
		if err != nil {
			return shim.Error("Error marshaling trade agreement structure")
		}
		// Write the state to the ledger
		err = stub.PutState(contractKey, contractBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		// UPDATE ENTRY FIELD
		// Get the item entry in the database using the contract fields
		itemId = contract.ItemId
		itemEntryBytes, err = stub.GetState(itemId)
		if err != nil {
			return shim.Error("Failed to get item: " + err.Error())
		}

		err = json.Unmarshal(itemEntryBytes, &itemEntry)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Update the middleman field of the item entry
		itemEntry.Warehouse = contract.WarehouseId

		itemEntryBytes, err = json.Marshal(itemEntry)
		if err != nil {
			return shim.Error("Error marshaling ItemEntry structure")
		}

		// Save item to state
		err = stub.PutState(itemId, itemEntryBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Printf("Storage Contract %s acceptance recorded\n", args[0])

	return shim.Success(nil)
}

func (t *TradeWorkflowChaincode) prepareShipment(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer string, args []string) pb.Response {
	var billOfLading *BillOfLading
	var billOfLadingBytes []byte
	var billOfLadingKey string
	var amount int
	var err error

	if !t.testMode && !authenticateCarrierOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error("Caller not a member of Ccarrier Org. Access denied.")
	}

	if len(args) != 5 {
		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting 5: {ID, Seller ID, Item ID, Amount, Buyer}. Found %d", len(args)))
		return shim.Error(err.Error())
	}

	amount, err = strconv.Atoi(string(args[3]))
	if err != nil {
		return shim.Error(err.Error())
	}

	// ADD TRADE LIMIT CHECK HERE

	billOfLading = &BillOfLading{args[0], args[1], args[2], amount, args[3], PREPARED}
	billOfLadingBytes, err = json.Marshal(billOfLading)
	if err != nil {
		return shim.Error("Error marshaling billOfLading structure")
	}

	// Write the state to the ledger
	billOfLadingKey, err = getTradeKey(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(billOfLadingKey, billOfLadingBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("Preparing shipment %s of item %s for buyer %s", args[0], args[2], args[4])
	return shim.Success(nil)
}

func main() {
	twc := new(TradeWorkflowChaincode)
	twc.testMode = true
	err := shim.Start(twc)
	if err != nil {
		fmt.Printf("Error starting Trade Workflow chaincode: %s", err)
	}
}
