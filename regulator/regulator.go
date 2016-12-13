/**************************************************************
 *Exemplo de uso das permissões dentro do chaincode
 *Neste caso somente o admin pode visualizar o valor proposto
 **************************************************************/

package main

import (
	"fmt"
	"strconv"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type adminStructState struct{
}

func main() {

	err := shim.Start(new(adminStructState))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func(t *adminStructState) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var Aval int
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting integer value for init")
	}

	// Write the state to the ledger
	err = stub.PutState("abc", []byte(strconv.Itoa(Aval))) //making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *adminStructState) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("[adminStructState] Invoke")
	if function == "init" {
      return t.Init(stub,"init",args)
	}
	return nil,nil
}

func (t *adminStructState) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
    return nil,nil
}

