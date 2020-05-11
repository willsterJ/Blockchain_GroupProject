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

	// QUERY: queryItems (Seller)
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'queryItems', ["{\"selector\":{\"descriptionOfGoods\":\"laptop\"}}"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('InitItem FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
.then((result) => {
	console.log('\n');
	console.log('-------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('queryItems VALUE:', result);
	console.log('-------------------------');
	console.log('\n');

	// INVOKE: updateItem (Seller)
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'updateItem', ["sellerlaptop", "2"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('queryItems FAILED');
	console.log('------------------------');
	console.log('\n');
	process.exit(1);
})
.then(() => {
	console.log('\n');
	console.log('------------------------------');
	console.log('CHAINCODE INVOCATION COMPLETE');
	console.log('updateItem SUCCEEDED');
	console.log('------------------------------');
	console.log('\n');

	// INVOKE: requestAdvertisement (Seller)
	return invokeCC.invokeChaincode(Constants.SELLER_ORG, Constants.CHAINCODE_VERSION, 'requestAdvertisement', ["adv-1", "middleman", "sellerlaptop", "0.08"], 'Seller');
}, (err) => {
	console.log('\n');
	console.log('-----------------------------');
	console.log('CHAINCODE INVOCATION FAILED:', err);
	console.log('requestAdvertisement FAILED');
	console.log('-----------------------------');
	console.log('\n');
	process.exit(1);
})
.then((result) => {
	console.log('\n');
	console.log('-------------------------');
	console.log('CHAINCODE QUERY COMPLETE');
	console.log('requestAdvertisement VALUE:', result);
	console.log('-------------------------');
	console.log('\n');

	ClientUtils.txEventsCleanup();
}, (err) => {
	console.log('\n');
	console.log('------------------------');
	console.log('CHAINCODE QUERY FAILED:', err);
	console.log('getAccountBalance FAILED');
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
