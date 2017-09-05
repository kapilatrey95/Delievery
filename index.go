package main

import (
	
	"errors"
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"

)

type SimpleChainCode struct{

}


func main() {
	err:=shim.Start(new(Supplychaincode))

	if err!=nil{
		fmt.Printf("errro starting Simple chaincode: %s",err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface,args []string) ([]byte,error){

	return nil,nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface,function string,args []string) ([]byte,error) {
	

	if function=="init" {
		return t.Init(stub,function,args)
	}  else if function=="AddUnit"{
		return t.AddUnit(stub,args)
	} else if function=="MakeCarton" {
		return t.MakeCarton(stub,args)
	} else if function=="SendCarton" {
		return t.SendCarton(stub,args)
	} else if function=="SendCartonWarehouse"{
		return t.SentCartonWarehouse(stub,args)
	} else if function=="RecieveCartonWarehouse"{
		return t.RecieveCartonWarehouse(stub,args)
	} else if function=="SellUnits"{
		return t.SellUnits(stub,args)
	} else if function=="InitManufacturer"{
		return t.RegisterManufacturer(stub,args)
	} else if function=="InitDistributer"{
		return t.RegisterDistributer(stub,args)
	} else if function=="InitRetailer"{
		return t.RegisterRetailer(stub,args)
	} else if function=="InitWarehouse"{
		return t.RegisterWarehouse(stub,args)
	} else {
		return nil,nil
}

}

