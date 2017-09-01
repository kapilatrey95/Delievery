package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"json/encoded"
	"strconv"
	"time"
)

func (t *SimpleChainCode) InitManufacturer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//args[0]=manufacturerId
	//args[1]=name
	if len(args) != 2 {
		return nil, errors.New("error:insufficient data")
	}

	manufacturerAsBytes, err := stub.GetState(args[0])

	if manufacturerAsBytes != nil {
		return nil, errors.New("manufacturer already exists with id ", args[0])
	}

	manufacturerAccount := manufacturer{}
	manufacturerAccount.ManufacturerId = args[0]
	manufacturerAccount.Name = args[1]

	manufacturerAsBytes, err := json.Marshal(manufacturerAccount)
	err = stub.PutState(args[0], manufacturerAsbytes)
	if err != nil {
		return nil, errors.New("registration failed please try again")
	}

	return manufacturerAsBytes, nil

}

func (t *Supplychaincode) AddUnit(stub shim.ChaincodeStubInterface, args []string) ([]byte, err) {
	/*args[0]=manufacturerId
	args[1]=productID
	args[2]=unitId
	args[3]=cartonId
	*/

	if len(args) != 4 {
		return nil, error.New("error supply 4 args")
	}

	manufacturerAsbytes, err := stub.GetState(manufacturerId)

	if manufacturerAsbytes == nil {
		return nil, error.New("Error :wrong manufacturerId please register")
	}
	if err != nil {
		return nil, err
	}

	cartonAsBytes, err := stub.GetState(cartonIdIndices)
	if cartonAsBytes == nil {
		return nil, error.New("error:no carton exists please create one")
	}
	if err != nil {
		return nil, error.New("error in getting carton indices")
	}

	newUnit = Units{} //new structure of type Units
	newUnit.ManufacturerId = args[0]
	newUnit.ProductID = args[1]
	newUnit.CartonId = args[3]
	date = time.Now()
	date.format("20060102150405")
	newUnit.UnitStatus = "Created by " + args[0] + " on " + date
	newUnit.DateOfDepartureFromManufacturer = string(date) + date.Zone()
	compositeString = "manufacturerId~unitId"
	compositeKey, err := stub.createCompositeKey(compositeString, []string{args[0], args[2]})
	newUnit.UnitId = args[2]
	newUnit.UnitCompositeKey=compositeKey

	newUnit.TrackRecord=newUnit.UnitStatus

	cartonObject = Carton{}
	err = json.Unmarshal(cartonAsBytes, &cartonObject)
	oldUnits := cartonObject.UnitsInCarton
	oldUnits = append(oldUnits, cartonObject)
	cartonObject.UnitsInCarton = oldUnits

	finalCarton, err := json.Marshal(cartonObject)
	if err != nil {
		return nil, error.New("error at 91 manufacturer.go")

	}

	err := stub.PutState(args[3], finalCarton)
	err := stub.PutState(compositeKey, newUnit)

	if err != nil {
		return nil, error.New("error:errror in 99 manufacturer.go")
	}
	return newUnit, nil

}

func (t *SimpleChainCode) MakeCarton(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, error.New("supply 2 args")
	}

	//args[0]=ManufacturerId
	//args[1]=CartonId

	validateManufacturer, err := stub.GetState(args[0])
	if err != nil {
		return nil, error.New("error at 175")
	}
	if validateManufacturer == nil {
		return nil, errors.New("wrong manufactrurer Id")
	}

	validateCarton, err := stub.GetState(args[1])
	if err != nil {
		return nil, error.New("error at 131")
	}
	if validateCarton != nil {
		return nil, errors.New("carton  with this id already exists")
	}

	newCarton = Carton{}

	newCarton.manufactrurerId = args[0]

	compositeString = "manufacturerId~unitId"
	compositeKey, err := stub.createCompositeKey(compositeString, []string{args[0], args[1]})
	newCarton.CartonId =args[1] 
	newCaron.CartonCompositeKey=compositeKey

	finalCarton, err := json.Marshal(newCarton)

	if err != nil {
		return nil, error.New("error at 148 manufacturer.go")
	}

	err := stub.PutState(compositeKey, newCarton)
	if err != nil {
		return nil, error.New("error:errror in 153 manufacturer.go")
	}
	return newCarton, nil

}

func (t *SimpleChainCode) SendCarton(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, error.New("excepting 3 arguments")
	}

	//args[0]=manufacturerId
	//args[1]=cartonId
	//args[2]=warehouseId

	validateManufacturer, err := stub.GetState(args[0])
	if err != nil {
		return nil, error.New("error at 220 manufacture.go")
	}
	if validateManufacturer == nil {
		return nil, errors.New("wrong manufactrurer Id")
	}

	validateCarton, err := stub.GetState(args[1])
	if err != nil {
		return nil, error.New("error at 230 manufacture.go")
	}
	if validateCarton != nil {
		return nil, errors.New("carton  with this id already exists")
	}

	validateWarehouse, err := stub.GetState(args[2])
	if err != nil {
		return nil, error.New("error at 238 manufacture.go")
	}
	if validateWarehouse == nil {
		return nil, errors.New("warehouse with this doesnt exists")
	}

	cartonDemo = Carton{}
	err := json.Unmarshal(validateCarton, &cartonDemo)
	if err != nil {
		return nil, error.New("error in 200 manufacturer.go")
	}

	date = time.Now()
	date.format("20060102150405")
	update := "delievering to " + args[2] + " by " + args[0] + " on " + date
	cartonDemo.recipient = args[2]

	cartonDemo.cartonStatus+=update
	unitStatusArray:=cartonDemo.UnitsInCarton

	for i:=0;i<len(unitStatusArray);i++ {
		tempUnitID:=unitStatusArray[i].UnitId
		tempUnit,err:=stub.GetState(tempUnitID)
		tempUnit.UnitStatus+=" "+update
		tempUnit.CurrentOwner=args[0]
		tempUnit.Recipient=args[2]
		tempUnit.State="Delievered"
		tempUnit.TrackRecord+=update
		finalTempUnitAsBytes,err:=json.Marshal(tempUnit)
		err:=stub.PutState(tempUnitID,finalTempUnitAsBytes)


	}

	cartonDemo.currentOwner=args[2]
	cartonDemo.dateOfDeparture=date

	CartonAsBytes,err:=json.Marshal(cartonDemo)
	if err!=nil {
		return nil,error.New("error at 260 manufacturer.go")
	}


		err:=stub.PutState(args[1],CartonAsBytes)

		if err!=nil {
			return nil,error.New("error at 286 manufacturee.go")
		}

		return cartonAsBytes,nil
}
