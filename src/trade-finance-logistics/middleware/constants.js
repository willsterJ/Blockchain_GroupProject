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

var os = require('os');
var path = require('path');

var tempdir = "../network/client-certs";
//path.join(os.tmpdir(), 'hfc');

// Frame the endorsement policy
var FOUR_ORG_MEMBERS_AND_ADMIN = [{
	role: {
		name: 'member',
		mspId: 'ExporterOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'ImporterOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'CarrierOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'RegulatorOrgMSP'
	}
}, {
	role: {
		name: 'admin',
		mspId: 'TradeOrdererMSP'
	}
}];

var FIVE_ORG_MEMBERS_AND_ADMIN = [{
	role: {
		name: 'member',
		mspId: 'SellerOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'BuyerOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'MiddlemanOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'CarrierOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'WarehouseOrgMSP'
	}
}, {
	role: {
		name: 'admin',
		mspId: 'TradeOrdererMSP'
	}
}];

var SEVEN_ORG_MEMBERS_AND_ADMIN = [{
	role: {
		name: 'member',
		mspId: 'Seller0OrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'Seller1OrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'Buyer0OrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'Buyer1OrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'MiddlemanOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'CarrierOrgMSP'
	}
}, {
	role: {
		name: 'member',
		mspId: 'WarehouseOrgMSP'
	}
}, {
	role: {
		name: 'admin',
		mspId: 'TradeOrdererMSP'
	}
}];

var ONE_OF_FOUR_ORG_MEMBER = {
	identities: FOUR_ORG_MEMBERS_AND_ADMIN,
	policy: {
		'1-of': [{ 'signed-by': 0 }, { 'signed-by': 1 }, { 'signed-by': 2 }, { 'signed-by': 3 }]
	}
};

var ALL_FOUR_ORG_MEMBERS = {
	identities: FOUR_ORG_MEMBERS_AND_ADMIN,
	policy: {
		'4-of': [{ 'signed-by': 0 }, { 'signed-by': 1 }, { 'signed-by': 2 }, { 'signed-by': 3 }]
	}
};

var ALL_FIVE_ORG_MEMBERS = {
	identities: FIVE_ORG_MEMBERS_AND_ADMIN,
	policy: {
		'5-of': [{ 'signed-by': 0 }, { 'signed-by': 1 }, { 'signed-by': 2 }, { 'signed-by': 3 }, { 'signed-by': 4 }]
	}
};

var ALL_ORGS_EXCEPT_REGULATOR = {
	identities: FOUR_ORG_MEMBERS_AND_ADMIN,
	policy: {
		'3-of': [{ 'signed-by': 0 }, { 'signed-by': 1 }, { 'signed-by': 2 }]
	}
};

var ALL_SEVEN_ORG_MEMBERS = {
	identities: SEVEN_ORG_MEMBERS_AND_ADMIN,
	policy: {
		'7-of': [{ 'signed-by': 0 }, { 'signed-by': 1 }, { 'signed-by': 2 }, { 'signed-by': 3 }, { 'signed-by': 4 }, { 'signed-by': 5 }, { 'signed-by': 6 }]
	}
};

var ACCEPT_ALL = {
	identities: [],
	policy: {
		'0-of': []
	}
};

var chaincodeLocation = '../chaincode';

var networkId = 'trade-network';

var networkConfig = './config.json';

var networkLocation = '../network';

var channelConfig = 'channel-artifacts/channel.tx';

// var BUYER_ORG = 'buyerorg';
// var SELLER_ORG = 'sellerorg';
var BUYER0_ORG = 'buyer0org';
var BUYER1_ORG = 'buyer1org';
var SELLER0_ORG = 'seller0org';
var SELLER1_ORG = 'seller1org';
var MIDDLEMAN_ORG = 'middlemanorg';
var CARRIER_ORG = 'carrierorg';
var WAREHOUSE_ORG = 'warehouseorg';

var CHANNEL_NAME = 'tradechannel';
var CHAINCODE_PATH = 'github.com/trade_workflow_v1'; // Unsure if we should change?
var CHAINCODE_ID = 'tradecc';
var CHAINCODE_VERSION = 'v1'; // here as well
var CHAINCODE_UPGRADE_PATH = 'github.com/trade_workflow_v1';
var CHAINCODE_UPGRADE_VERSION = 'v1';

var TRANSACTION_ENDORSEMENT_POLICY = ALL_SEVEN_ORG_MEMBERS; //Need to change eventually

module.exports = {
	tempdir: tempdir,
	chaincodeLocation: chaincodeLocation,
	networkId: networkId,
	networkConfig: networkConfig,
	networkLocation: networkLocation,
	channelConfig: channelConfig,
	BUYER0_ORG: BUYER0_ORG,
	BUYER1_ORG: BUYER1_ORG,
	SELLER0_ORG: SELLER0_ORG,
	SELLER1_ORG: SELLER1_ORG,
	MIDDLEMAN_ORG: MIDDLEMAN_ORG,
	CARRIER_ORG: CARRIER_ORG,
	WAREHOUSE_ORG: WAREHOUSE_ORG,
	CHANNEL_NAME: CHANNEL_NAME,
	CHAINCODE_PATH: CHAINCODE_PATH,
	CHAINCODE_ID: CHAINCODE_ID,
	CHAINCODE_VERSION: CHAINCODE_VERSION,
	CHAINCODE_UPGRADE_PATH: CHAINCODE_UPGRADE_PATH,
	CHAINCODE_UPGRADE_VERSION: CHAINCODE_UPGRADE_VERSION,
	ALL_FOUR_ORG_MEMBERS: ALL_FOUR_ORG_MEMBERS,
	ALL_FIVE_ORG_MEMBERS: ALL_FIVE_ORG_MEMBERS,
	ALL_SEVEN_ORG_MEMBERS: ALL_SEVEN_ORG_MEMBERS,
	TRANSACTION_ENDORSEMENT_POLICY: TRANSACTION_ENDORSEMENT_POLICY
};
