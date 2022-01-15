package main

import (
		"bytes"
		"encoding/json"
		"fmt"
		"strconv"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"

	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {	

}

type robot struct {

		robotCODE string `json:"robotcode"`

		Owner string `json:"owner"`

		Model string `json:"model"`

		Make string `json:"make"`

		Color string `json:"color"`

}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "queryrobot" {
			return s.queryrobot(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createrobot" {
		return s.createrobot(APIstub, args)
	} else if function == "queryAllrobots" {
		return s.queryAllrobots(APIstub)
	} else if function == "changeRobotOwner" {
		return s.changeRobotOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryrobot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	robotAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(robotAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	robots := []robot{
		robot{robotCODE: "E4A12RASDFAS4G", Owner: "육군", Model: "대한민국RT", Make: "국방과학연구소", Color: "Blue"},
		robot{robotCODE: "BA5IJ32432FGSD", Owner: "육군", Model: "HWD001", Make: "한화디펜스", Color: "Red"},
		robot{robotCODE: "11CZJ23R13TVC8", Owner: "현대", Model: "현대SA1", Make: "현대", Color: "Yellow"},
		robot{robotCODE: "3VK5CBV967BA23", Owner: "Subin", Model: "Hwangbenhavn", Make: "금손콩", Color: "Freen"},
		robot{robotCODE: "87234HOF984892", Owner: "Goldwaterbean", Model: "마징가Z", Make: "황황황황황", Color: "Black"},
		robot{robotCODE: "5M8409544CN80M", Owner: "Gosuohwan", Model: "또봇", Make: "수수", Color: "Purple"},
		robot{robotCODE: "98MX14GAS2432G", Owner: "엔티렉스", Model: "엔티렉스R31", Make: "엔티렉스", Color: "Lime"},
		robot{robotCODE: "JD932239F78Q2F", Owner: "위드로봇", Model: "위드로봇W1", Make: "위드로봇", Color: "Aqua"},
		robot{robotCODE: "VE84379R342384", Owner: "칩센", Model: "칩센001", Make: "칩센", Color: "Brown"},
		robot{robotCODE: "XM4128X47GFGS4", Owner: "펌테크", Model: "펌테크K3", Make: "펌테크", Color: "Maroon"},
		robot{robotCODE: "D8BOV83NOD82M1", Owner: "뉴티씨", Model: "뉴티씨N11", Make: "뉴티씨", Color: "Peach"},
		robot{robotCODE: "VXCNZ98Z781DCG", Owner: "디지키", Model: "디지키DGK", Make: "디지키", Color: "Cyan"},
		robot{robotCODE: "QQU31HORE123FD", Owner: "한화", Model: "HWD333", Make: "한화디펜스", Color: "Pink"},
		robot{robotCODE: "39VMVSD41JF981", Owner: "엘지", Model: "LGR", Make: "엘지", Color: "Mustard"},
}

	i := 0
	for i < len(robots) {
		fmt.Println("i is ", i)
		robotAsBytes, _ := json.Marshal(robots[i])
		APIstub.PutState("robot"+strconv.Itoa(i), robotAsBytes)
		fmt.Println("Added", robots[i])
		i = i + 1
}

	return shim.Success(nil)
}

func (s *SmartContract) createrobot(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
	return shim.Error("Incorrect number of arguments. Expecting 6")
	}

var robot = robot{robotCODE: args[1], Owner: args[2], Model: args[3], Make: args[4], Color: args[5]}

	robotAsBytes, _ := json.Marshal(robot)
	APIstub.PutState(args[0], robotAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllrobots(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "AAAAAAAAAAAAAA"
	endKey := "99999999999999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
}
	defer resultsIterator.Close()

var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
	queryResponse, err := resultsIterator.Next()

		if err != nil {
		return shim.Error(err.Error())
}
	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
}
	buffer.WriteString("]")

	fmt.Printf("- queryAllrobots:\n%s\n", buffer.String())

		return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeRobotOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
}

	robotAsBytes, _ := APIstub.GetState(args[0])
	robot := robot{}

	json.Unmarshal(robotAsBytes, &robot)
	robot.Owner = args[1]

	robotAsBytes, _ = json.Marshal(robot)
	APIstub.PutState(args[0], robotAsBytes)

		return shim.Success(nil)
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
