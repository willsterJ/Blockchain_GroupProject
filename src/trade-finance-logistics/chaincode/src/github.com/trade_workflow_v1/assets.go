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

type ItemEntry struct {
	Id                 string  `json:id`
	DescriptionOfGoods string  `json:"descriptionOfGoods"`
	Seller             string  `json:"seller"`
	Middleman          string  `json:"middleman"`
	Warehouse          string  `json:"warehouse"`
	Price              float64 `json:"price"`
	Count              int     `json:"count"`
	FeeToMiddleman	float64 `json:"feeToMiddleman"`
	FeeToWarehouse	float64 `json:"feeToWarehouse"`
	FeeToCarrier	float64 `json:"feeToCarrier"`
	DocType	string	`json:"docType"`
}
type TradeAgreement struct {
	Amount             int    `json:"amount"`
	DescriptionOfGoods string `json:"descriptionOfGoods"`
	Status             string `json:"status"`
	Payment            int    `json:"payment"`
	ItemId 	string `json:"itemId"`
}
type BillOfLading struct {
	Id                 string `json:"id"`
	Seller             string `json:"seller"`
	ItemId 			   string `json:"itemId"`
	Amount             int    `json:"amount"`
	Buyer              string `json:"buyer"`
	Status 			   string `json:"status"`
}
type ContractSellerMiddleman struct {
	Id          string  `json:"id"`
	MiddlemanId string  `json:"middlemanId"`
	ItemId      string  `json:"itemId"`
	Fee         float64 `json:"fee"`
	Status      string  `json:"status"`
}
type ContractSellerWarehouse struct {
	Id          string  `json:"id"`
	WarehouseId string  `json:"warehouseId"`
	ItemId      string  `json:"itemId"`
	Fee         float64 `json:"fee"`
	Status      string  `json:"status"`
}

/* // original
type TradeAgreement struct {
	Amount			int		`json:"amount"`
	DescriptionOfGoods	string		`json:"descriptionOfGoods"`
	Status			string		`json:"status"`
	Payment			int		`json:"payment"`
}

type LetterOfCredit struct {
	Id			string		`json:"id"`
	ExpirationDate		string		`json:"expirationDate"`
	Beneficiary		string		`json:"beneficiary"`
	Amount			int		`json:"amount"`
	Documents		[]string	`json:"documents"`
	Status			string		`json:"status"`
}

type ExportLicense struct {
	Id			string		`json:"id"`
	ExpirationDate		string		`json:"expirationDate"`
	Exporter		string		`json:"exporter"`
	Carrier			string		`json:"carrier"`
	DescriptionOfGoods	string		`json:"descriptionOfGoods"`
	Approver		string		`json:"approver"`
	Status			string		`json:"status"`
}

type BillOfLading struct {
	Id			string		`json:"id"`
	ExpirationDate		string		`json:"expirationDate"`
	Exporter		string		`json:"exporter"`
	Carrier			string		`json:"carrier"`
	DescriptionOfGoods	string		`json:"descriptionOfGoods"`
	Amount			int		`json:"amount"`
	Beneficiary		string		`json:"beneficiary"`
	SourcePort		string		`json:"sourcePort"`
	DestinationPort		string		`json:"destinationPort"`
}
*/
