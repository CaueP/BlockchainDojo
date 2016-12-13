/**************************************************************
 *Exemplo de uso das permissões dentro do chaincode
 *Neste caso somente o admin pode visualizar o valor proposto
 **************************************************************/

package main

import (
	"fmt"
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
  fmt.Println("[adminStructState] Init Chaincode...")
  return nil,nil
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

