/**************************************************************
 *Exemplo de uso das permissões dentro do chaincode
 *Neste caso somente o admin pode visualizar o valor proposto
 **************************************************************/

package main

import (
	"fmt"
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
	fmt.Println("[init] Init Chaincode")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}	
     // Set the admin
	// The metadata will contain the certificate of the administrator
	fmt.Println("Getting caller metadata")
	adminCert, err := stub.GetCallerMetadata() 
	if err != nil {
		fmt.Println("Failed getting metadata")
		return nil, errors.New("Failed getting metadata.")
	}
	if len(adminCert) == 0 {
		fmt.Printf("Invalid admin certificate. Empty.")
		return nil, errors.New("Invalid admin certificate. Empty.")
	}

	fmt.Printf("The administrator is [%x]", adminCert)

	stub.PutState("admin", adminCert)

	return nil, err
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
	if function == "read" {
		return t.read(stub,args)
	}
    return nil,nil
}

func (t *adminStructState) read(stub shim.ChaincodeStubInterface, args []string)([]byte, error){
    // Recover the role that is allowed to make assignments
	admin, err := stub.GetState("admin")
	if err != nil {
		fmt.Printf("Error getting role [%v] \n", err)
		return nil, errors.New("Failed fetching assigner role")
	}
	regulator := string(admin[:])
	fmt.Printf("[getTransaction] Regulator authorized! [%v]" , regulator)

    return nil,nil
}

