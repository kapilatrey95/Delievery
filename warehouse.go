package main 

import (

	"fmt"
	"json/encoded"
	"errors"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"

)

func (t *Simplechaincode) InitWarehouse(stub shim.ChaincodeStubInterface,args []string) ([]byte,err) {


	if len(args)!=1{
		return nil,errro.New("supply 1 args")
	}

	warehouseCheck,err:=stub.GetState(args[0])
	if warehouseCheck!=nil{
		return nil,error.New("distributer already exists")
	}

	newWarehouse=Warehouse{}

	newWarehouse.wareHouseId=args[0]

	finalWareHouse,err:=json.Marshal(newWarehouse)

	err:=stub.PutState(args[0],finalWareHouse)


return finalWarehouse,nil

} 

func (t *Simplechaincode) RecieveCartonWarehouse(stub shim.ChaincodeStubInterface,args []string) ([]byte,err) {

	//args[0]=warehouseId
	//args[1]=cartonId
	
	if len(args)!=2{
		return nil,error.New("pass two arguments")
	}

	validateWarehouse,err:=stub.GetState(args[0])
	if err!=nil {
		return nil,error.New("error at 175")
	}
	if validateWarehouse==nil{
		return nil,errors.New("wrong warehouse Id")
	}

	validateCarton,err:=stub.GetState(args[1])
	if err!=nil {
		return nil,error.New("error at 183")
	}
	if validateCarton==nil{
		return nil,errors.New("no carton  with this id ")
	}
	

	cartonDemo=Carton{}
	err:=json.Unmarshal(validateCarton,&cartonDemo)
	if err!=nil{
		return nil,error.New("error in 68 warehouse.go")
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
		return nil,error.New("error in 98 warehouse.go")
	}
	err:=stub.PutState(args[1],finalCarton)
	if err!=nil{
		return nil,error.New("error in 102 warehouse.go")
	}

	return finalCarton,nil


}

func (t *Simplechaincode) SendCartonWarehouse(stub shim.ChaincodeStubInterface,args []string) ([]byte,err) {

	if len(args)!=3{
		return nil,error.New("provide 3 arguments")
	}

	//args[0]=WarehouseId
	//args[1]=cartonId
	//args[2]=DistributerId

	validateWarehouse,err:=stub.GetState(args[0])
	if err!=nil {
		return nil,error.New("error at 175")
	}
	if validateWarehouse==nil{
		return nil,errors.New("wrong warehouse Id")
	}

	validateCarton,err:=stub.GetState(args[1])
	if err!=nil {
		return nil,error.New("error at 183")
	}
	if validateCarton==nil{
		return nil,errors.New(" no carton  with this id")
	}

	validateDistributer,err:=stub.GetState(args[2])
	if err!=nil{
		return nil,error.New("no such distributer")
	}

	cartonDemo=Carton{}
	err:=json.Unmarshal(validateCarton,&cartonDemo)
	if err!=nil {
		return nil,error.New("error at 144 warehouse.go")
	}

	if cartonDemo.CurrentOwner!=args[0]{
		return nil,error.New("your are not owner of this carton")
	}

	date=time.Now()
	date.format("20060102150405")

	cartonDemo.Recipient=args[2]
	update:="carton sent from "+args[1]+" to "+args[2]+date
	cartonDemo.DateOfDeparture=date
	cartonDemo.CartonStatus+=update
	cartonDemo.DateOfDeparture=date


	unitStatusArray:=cartonDemo.UnitsInCarton
	for i:=0;i<len(unitStatusArray);i++ {
		tempUnitID:=unitStatusArray[i].UnitId
		tempUnit,err:=stub.GetState(tempUnitID)
		tempUnit.UnitStatus+=" "+update
		tempUnit.TrackRecord=update
		tempUnit.CurrentOwner=args[0]
		tempUnit.Recipient=args[3]
		tempUnit.State="delieverimg to "+args[3]
		finalTempUnitAsBytes,err:=json.Marshal(tempUnit)
		err:=stub.PutState(tempUnitID,finalTempUnitAsBytes)
		if err!=nil {
			return nil,error.New("at 177 warehouse.go")
		}
	}


	finalCarton,err:=json.Marshal(cartonDemo)
	if err!=nil{
		return nil,error.New("error in 184 warehouse.go")
	}
	err:=stub.PutState(args[1],finalCarton)
	if err!=nil{
		return nil,error.New("error in 188 warehouse.go")
	}

	return finalCarton,nil

}


