package main

import (
	"encoding/json"
	"fmt"

	//  "byte"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}
type QueryBar struct {
	Key    string `json:"Key"`
	Record *Bar
}
type Bar struct {
	BarID              string `json:"id"`
	BarLocation        string `json:"barLocation"`
	BarSerialNumber    string `json:"barSerialNumber"`
	Purity             string `json:"purity"`
	BarRefiner         string `json:"barRefiner"`
	BarHallmarkVerfied string `json:"barHallmarkVerfied"`
	BarWeightInGms     string `json:"barWeightInGms"`
}
type QueryBuy struct {
	Key    string `json:"Key"`
	Record *Buy
}
type Buy struct {
	DSGId           string `json:"id"`
	OrderId         string `json:"orderId"`
	BarCostPostLbma string `json:"barCostPostLbma"`
	TotalToFund     string `json:"totalToFund"`
	TotalWeight     string `json:"totalWeight"`
	BarSerialNumber string `json:"barSerialNumber"`
	BarIds          string `json:"barIds"`
	AccountNo       string `json:"accountNo"`
	UserId          string `json:"userId"`
}
type QuerySell struct {
	Key    string `json:"Key"`
	Record *Sell
}
type Sell struct {
	DSGId       string `json:"id"`
	BarIds      string `json:"barIds"`
	OrderId     string `json:"orderId"`
	TotalWeight string `json:"totalWeight"`
	TotalToFund string `json:"totalToFund"`
	UserId      string `json:"userId"`
}
type QuerySend struct {
	Key    string `json:"Key"`
	Record *Send
}
type Send struct {
	DSGId          string `json:"id"`
	BarIds         string `json:"barIds"`
	OrderId        string `json:"orderId"`
	TotalWeight    string `json:"totalWeight"`
	SenderUserId   string `json:"senderUserId"`
	ReceiverUserId string `json:"receiverUserId"`
}
type QueryTrade struct {
	Key    string `json:"Key"`
	Record *Trade
}
type Trade struct {
	DSGId       string `json:"id"`
	OrderId     string `json:"orderId"`
	TotalWeight string `json:"totalWeight"`
	UserId      string `json:"userId"`
}

func GetUId() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), err
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	fmt.Printf("Hello\n")
	return nil
}
func (s *SmartContract) CreateBar(ctx contractapi.TransactionContextInterface, BarLocation string, BarSerialNumber string, Purity string, BarRefiner string, BarHallmarkVerfied string, BarWeightInGms string) error {

	fmt.Printf("Adding Bar to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
	// }

	//Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "Bar-" + uid
	fmt.Printf("Validating Bar data\n")
	//Validate the Org data
	var bar = Bar{BarID: id,
		BarLocation:        BarLocation,
		BarSerialNumber:    BarSerialNumber,
		Purity:             Purity,
		BarRefiner:         BarRefiner,
		BarHallmarkVerfied: BarHallmarkVerfied,
		BarWeightInGms:     BarWeightInGms,
	}

	//Encrypt and Marshal Org data in order to put in world state
	barAsBytes, _ := json.Marshal(bar)

	return ctx.GetStub().PutState(id, barAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) GetBar(ctx contractapi.TransactionContextInterface, BarSerialNumber string) ([]QueryBar, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Bar-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryBar{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		bar := new(Bar)
		_ = json.Unmarshal(queryResponse.Value, bar)
		if bar.BarSerialNumber == BarSerialNumber {

			queryResult := QueryBar{Key: queryResponse.Key, Record: bar}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) QueryBar(ctx contractapi.TransactionContextInterface, BarSerialNumber string) (*Bar, error) {
	barAsBytes, err := ctx.GetStub().GetState(BarSerialNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if barAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", BarSerialNumber)
	}

	bar := new(Bar)
	_ = json.Unmarshal(barAsBytes, bar)

	return bar, nil
}

func (s *SmartContract) GetBarList(ctx contractapi.TransactionContextInterface) ([]QueryBar, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^Bar-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBar{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		bar := new(Bar)
		_ = json.Unmarshal(queryResponse.Value, bar)
		queryResult := QueryBar{Key: queryResponse.Key, Record: bar}
		results = append(results, queryResult)
	}
	return results, nil
}
func (s *SmartContract) CreateBuy(ctx contractapi.TransactionContextInterface, OrderId string, TotalToFund string, BarCostPostLbma string, TotalWeight string, BarSerialNumber string, AccountNo string, UserId string) error {

	fmt.Printf("Adding Buy to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
	// }

	//Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "DSG-" + uid
	fmt.Printf("Validating Buy data\n")
	//Validate the Org data
	var buy = Buy{DSGId: id,
		OrderId:         OrderId,
		TotalToFund:     TotalToFund,
		BarCostPostLbma: BarCostPostLbma,
		TotalWeight:     TotalWeight,
		BarSerialNumber: BarSerialNumber,
		AccountNo:       AccountNo,
		UserId:          UserId,
	}

	//Encrypt and Marshal Org data in order to put in world state
	buyAsBytes, _ := json.Marshal(buy)

	return ctx.GetStub().PutState(id, buyAsBytes)

}
func (s *SmartContract) QueryBuy(ctx contractapi.TransactionContextInterface, OrderId string) (*Buy, error) {
	buyAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if buyAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	buy := new(Buy)
	_ = json.Unmarshal(buyAsBytes, buy)

	return buy, nil
}
func (s *SmartContract) GetBuy(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		if buy.OrderId == OrderId {

			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetBuyList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryBuy, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryBuy{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		buy := new(Buy)
		_ = json.Unmarshal(queryResponse.Value, buy)
		if buy.OrderId == OrderId {
			queryResult := QueryBuy{Key: queryResponse.Key, Record: buy}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateSell(ctx contractapi.TransactionContextInterface, OrderId string, BarIds string, TotalWeight string, TotalToFund string, UserId string) error {

	fmt.Printf("Adding Sell to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
	// }

	//Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "DSG-" + uid
	fmt.Printf("Validating Sell data\n")
	//Validate the Org data
	var sell = Sell{DSGId: id,
		OrderId:     OrderId,
		BarIds:      BarIds,
		TotalWeight: TotalWeight,
		TotalToFund: TotalToFund,
		UserId:      UserId,
	}

	//Encrypt and Marshal Org data in order to put in world state
	sellAsBytes, _ := json.Marshal(sell)

	return ctx.GetStub().PutState(id, sellAsBytes)

}
func (s *SmartContract) QuerySell(ctx contractapi.TransactionContextInterface, OrderId string) (*Sell, error) {
	sellAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if sellAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	sell := new(Sell)
	_ = json.Unmarshal(sellAsBytes, sell)

	return sell, nil
}
func (s *SmartContract) GetSell(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySell, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QuerySell{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		sell := new(Sell)
		_ = json.Unmarshal(queryResponse.Value, sell)
		if sell.OrderId == OrderId {

			queryResult := QuerySell{Key: queryResponse.Key, Record: sell}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetSellList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySell, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySell{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		sell := new(Sell)
		_ = json.Unmarshal(queryResponse.Value, sell)
		if sell.OrderId == OrderId {
			queryResult := QuerySell{Key: queryResponse.Key, Record: sell}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateSend(ctx contractapi.TransactionContextInterface, OrderId string, BarIds string, TotalWeight string, SenderUserId string, ReceiverUserId string) error {

	fmt.Printf("Adding Send to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
	// }

	//Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "DSG-" + uid
	fmt.Printf("Validating Send data\n")
	//Validate the Org data
	var send = Send{DSGId: id,
		OrderId:        OrderId,
		BarIds:         BarIds,
		TotalWeight:    TotalWeight,
		SenderUserId:   SenderUserId,
		ReceiverUserId: ReceiverUserId,
	}

	//Encrypt and Marshal Org data in order to put in world state
	sendAsBytes, _ := json.Marshal(send)

	return ctx.GetStub().PutState(id, sendAsBytes)

}
func (s *SmartContract) QuerySend(ctx contractapi.TransactionContextInterface, OrderId string) (*Send, error) {
	sendAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if sendAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	send := new(Send)
	_ = json.Unmarshal(sendAsBytes, send)

	return send, nil
}
func (s *SmartContract) GetSend(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySend, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QuerySend{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		send := new(Send)
		_ = json.Unmarshal(queryResponse.Value, send)
		if send.OrderId == OrderId {

			queryResult := QuerySend{Key: queryResponse.Key, Record: send}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetSendList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QuerySend, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QuerySend{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		send := new(Send)
		_ = json.Unmarshal(queryResponse.Value, send)
		if send.OrderId == OrderId {
			queryResult := QuerySend{Key: queryResponse.Key, Record: send}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func (s *SmartContract) CreateTrade(ctx contractapi.TransactionContextInterface, OrderId string, TotalWeight string, UserId string) error {

	fmt.Printf("Adding Trade to the ledger ...\n")
	// if len(args) != 8 {
	// 	return fmt.Errorf("InvalidArgumentError: Incorrect number of arguments. Expecting 8")
	// }

	//Prepare key for the new Org
	uid, err := GetUId()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	id := "DSG-" + uid
	fmt.Printf("Validating Trade data\n")
	//Validate the Org data
	var trade = Trade{DSGId: id,
		OrderId:     OrderId,
		TotalWeight: TotalWeight,
		UserId:      UserId,
	}

	//Encrypt and Marshal Org data in order to put in world state
	tradeAsBytes, _ := json.Marshal(trade)

	return ctx.GetStub().PutState(id, tradeAsBytes)

}
func (s *SmartContract) QueryTrade(ctx contractapi.TransactionContextInterface, OrderId string) (*Trade, error) {
	tradeAsBytes, err := ctx.GetStub().GetState(OrderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if tradeAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", OrderId)
	}

	trade := new(Trade)
	_ = json.Unmarshal(tradeAsBytes, trade)

	return trade, nil
}
func (s *SmartContract) GetTrade(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryTrade, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	result := []QueryTrade{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		trade := new(Trade)
		_ = json.Unmarshal(queryResponse.Value, trade)
		if trade.OrderId == OrderId {

			queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
			result = append(result, queryResult)
		}
	}
	return result, nil
}
func (s *SmartContract) GetTradeList(ctx contractapi.TransactionContextInterface, OrderId string) ([]QueryTrade, error) {
	query := "{\"selector\": {\"_id\": {\"$regex\": \"^DSG-\"} } }"
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []QueryTrade{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		trade := new(Trade)
		_ = json.Unmarshal(queryResponse.Value, trade)
		if trade.OrderId == OrderId {
			queryResult := QueryTrade{Key: queryResponse.Key, Record: trade}
			results = append(results, queryResult)
		}
	}
	return results, nil
}
func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create  chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting  chaincode: %s", err.Error())
	}
}
