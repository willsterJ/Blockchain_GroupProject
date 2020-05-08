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

// Key names
const (
	buyKey = "Buyer"
	buyBalKey = "BuyersAccountBalance"
	selKey = "Seller"
	selBalKey = "SellersAccountBalance"
	midKey = "Middleman"
	midBalKey = "MiddlemansAccountBalance"
	warKey = "Warehouse"
	warBalKey = "WarehousesAccountBalance"
	carKey		= "Carrier"
	carBalKey = "CarriersAccountBalance"
)

// State values
const (
	REQUESTED	= "REQUESTED"
	ISSUED		= "ISSUED"
	ACCEPTED	= "ACCEPTED"
)

// Location values
const (
	SOURCE		= "SOURCE"
	DESTINATION	= "DESTINATION"
)
