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

'use strict';

var Constants = require('./constants.js');
var ClientUtils = require('./clientUtils.js');
var createChannel = require('./create-channel.js');
var joinChannel = require('./join-channel.js');
var installCC = require('./install-chaincode.js');
var instantiateCC = require('./instantiate-chaincode.js');
var invokeCC = require('./invoke-chaincode.js');
var queryCC = require('./query-chaincode.js');

/////////////////////////////////
// INVOKE AND QUERY OPERATIONS //
/////////////////////////////////

// INVOKE: initItem (Seller)
invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'initItem', ['laptop', '1200', '20'], 'Seller')
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('initItem SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return queryCC.queryChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'queryItems', ["{\"selector\":{\"descriptionOfGoods\":\"laptop\"}}"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('InitItem FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// QUERY: queryItems (Seller)
.then((result) => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('queryItems VALUE:', result);
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'updateItem', ["SellerOrgMSPlaptop", "2"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('queryItems FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: updateItem (Seller)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('updateItem SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return queryCC.queryChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'queryItems', ["{\"selector\":{\"descriptionOfGoods\":\"laptop\"}}"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('updateItem FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// QUERY: queryItems (Seller)
.then((result) => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('queryItems VALUE:', result);
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'requestAdvertisement', ["adv-1", "middleman", "SellerOrgMSPlaptop", "0.08"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('queryItems FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: requestAdvertisement (Seller)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('requestAdvertisement SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.MIDDLEMAN_ORG, Constants.CHAINCODE_VERSION, 'acceptAdvertisement', ["adv-1"], 'Middleman');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('requestAdvertisement FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: acceptAdvertisement (Middleman)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('acceptAdvertisement SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'requestStorage', ["store-1", "warehouse", "SellerOrgMSPlaptop", "0.08"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('acceptAdvertisement FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: requestStorage (Seller)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('requestStorage SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.WAREHOUSE_ORG, Constants.CHAINCODE_VERSION, 'acceptStorage', ["store-1"], 'Warehouse');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('requestStorage FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: acceptStorage (Warehouse)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('acceptStorage SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.BUYER_ORG, Constants.CHAINCODE_VERSION, 'requestTrade', ["trade-1", "1", "laptop"], 'Buyer');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('acceptStorage FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: requestTrade (Buyer)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('requestTrade SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.MIDDLEMAN_ORG, Constants.CHAINCODE_VERSION, 'acceptTrade', ["trade-1"], 'Middleman');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('requestTrade FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: acceptTrade (Middleman)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('acceptTrade SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.CARRIER_ORG, Constants.CHAINCODE_VERSION, 'prepareShipment', ["ship-1","SellerOrgMSP", "SellerOrgMSPlaptop", "1", "BuyerOrgMSP"], 'Carrier');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('acceptTrade FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: prepareShipment (Carrier)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('prepareShipment SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return queryCC.queryChaincode(Constants.BUYER_ORG, Constants.CHAINCODE_VERSION, 'getShipmentStatus', ["ship-1"], 'Buyer');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('prepareShipment FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// QUERY: getShipmentStatus (Buyer)
.then((result) => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('getShipmentStatus VALUE:', result);
	console.log('------------------------------');
	console.log('\n');
		
	return invokeCC.invokeChaincode(Constants.CARRIER_ORG, Constants.CHAINCODE_VERSION, 'deliverShipment', ["ship-1"], 'Carrier');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('getShipmentStatus FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// INVOKE: deliverShipment (Carrier)
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('deliverShipment SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');
		
	return queryCC.queryChaincode(Constants.BUYER_ORG, Constants.CHAINCODE_VERSION, 'getShipmentStatus', ["ship-1"], 'Buyer');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('deliverShipment FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
// QUERY: getShipmentStatus (Buyer)
.then((result) => {
	console.log('\n');
	console.log('-------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('getShipmentStatus VALUE:', result);
	console.log('-------------------------');
	console.log('\n');

	ClientUtils.txEventsCleanup();
}, (err) => {
	console.log('\n');
	console.log('------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('getShipmentStatus FAILED');
	console.log('------------------------');
	console.log('\n');
	process.exit(1);
});

process.on('uncaughtException', err => {
	console.error(err);
	joinChannel.joinEventsCleanup();
});

process.on('unhandledRejection', err => {
	console.error(err);
	joinChannel.joinEventsCleanup();
});

process.on('exit', () => {
	joinChannel.joinEventsCleanup();
});
