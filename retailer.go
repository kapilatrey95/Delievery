package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"json/encoded"
	"strconv"
	"time"
)

func (t *Simplechaincode) InitRetailer(stub shim.ChaincodeStubInterface,args []string) ([]byte,err) {


	if len(args)!=1{
		return nil,errro.New("supply 1 args")
	}

	retailerCheck,err:=stub.GetState(args[0])
	if retailerCheck!=nil{
		return nil,error.New("retailer already exists")
	}

	newRetailer=Retailer{}

	newRetailer.wareHouseId=args[0]

	finalRetailer,err:=json.Marshal(newRetailer)

	err:=stub.PutState(args[0],finalRetailer)


return finalRetailer,nil

}

func (t *Simplechaincode) RecieveCartonRetailer(stub shim.ChaincodeStubInterface,args []string) ([]byte,err) {

	//args[0]=retailerId
	//args[1]=cartonId
	
	if len(args)!=2{
		return nil,error.New("pass two arguments")
	}

	validateRetailer,err:=stub.GetState(args[0])
	if err!=nil {
		return nil,error.New("error at 48 distributer")
	}
	if validateRetailer==nil{
		return nil,errors.New("wrong distributer Id")
	}

	validateCarton,err:=stub.GetState(args[1])
	if err!=nil {
		return nil,error.New("error at 56 distributer")
	}
	if validateCarton==nil{
		return nil,errors.New("no carton  with this id ")
	}
	

	cartonDemo=Carton{}
	err:=json.Unmarshal(validateCarton,&cartonDemo)
	if err!=nil{
		return nil,error.New("error in 66 warehouse.go")
	}
	if cartonDemo.Recipient!=args[0] {
		return nil,error.New("this carton is not for you sir")
	}
	cartonDemo.Recipient=nil
	cartonDemo.CurrentOwner=args[0]
	cartonDemo.DateOfDeparture=nil
	date=time.Now()
	date.format("20060102150405")

	update:=" recieved from"+cartonDemo.ManufacturerId+" on "+date

	cartonDemo.CartonStatus+=update

	unitStatusArray:=cartonDemo.UnitsInCarton


	for i:=0;i<len(unitStatusArray);i++ {
		tempUnitID:=unitStatusArray[i].UnitId
		tempUnit,err:=stub.GetState(tempUnitID)
		tempUnit.UnitStatus+=" "+update
		tempUnit.TrackRecord=update
		tempUnit.CurrentOwner=args[0]
		tempUnit.Recipient=nil
		tempUnit.State="delievered to "+args[0]
		finalTempUnitAsBytes,err:=json.Marshal(tempUnit)
		err:=stub.PutState(tempUnitID,finalTempUnitAsBytes)


	}

	finalCarton,err:=json.Marshal(cartonDemo)
	if err!=nil{
		return nil,error.New("error in 100 distributer.go")
	}
	err:=stub.PutState(args[1],finalCarton)
	if err!=nil{
		return nil,error.New("error in 102 distributer.go")
	}

	return finalCarton,nil


}


func (t *Simplechaincode) SellUnit(stub shim.ChaincodeStubInterface,args []string) ([]byte,error) {

	//args[0]=unitKey
	//args[1]=retailerId
	//args[2]=buyerName


	validateUnit,err:=stub.GetState(args[0])
	if err!=nil {
		return nil,error.New("error at 121 retailer")
	}
	if validateRetailer==nil{
		return nil,errors.New("wrong unit Id")
	}
	unitDemo=Units{}
	err:=json.Unmarshal(validateUnit,&unitDemo)
	if err!=nil {
		return nil,error.New("at 129 retailer")
	}

	if unitDemo.Consumer!=nil{
		return nil,error.New("unit already sold")
	}

	if unitDemo.CurrentOwner!=retailerId {
		return nil,error.New("this unit doesnt belong to you")
	}

	date=time.Now()
	date.format("20060102150405")
	update="sold to "+args[2]+" on "+date
	unitDemo.Recipient=args[2]
	unitDemo.Consumer=args[2]
	unitDemo.UnitStatus+=update
	unitDemo.GetState="sold to "+args[2]
	unitDemo.TrackRecord="dispached to "+args[2]+" by "+args[1]+" on date "+date

	finalTempUnitAsBytes,err:=json.Marshal(unitDemo)
		err:=stub.PutState(tempUnitID,finalTempUnitAsBytes)

		if err!=nil {
			return nil,error("error in 154")
		}

		return finalTempUnitAsBytes,nil


}
	

