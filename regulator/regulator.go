/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*******************************************
 * Alguns structs a nível de demo,o ideal seria que somente as informacoes de transacao tivessem 
 * structs mesmo
 *******************************************/

type Beneficiario struct {
	id int `json:"id_beneficiario"`
	nome string `json:"nome_beneficiario"`
	assinatura string `json:"assign_digital"`
}

type Pagador struct {
	id int `json:"id_pagador"`
	cpf float32 `json:"cpf_pagador"`
	nome string `json:"nome_pagador"`
}

type Proposta struct {
    id int `json:"id_proposta"`
	tipo string `json:"tipo"`
	valor float64 `json:"valor"`
	dtCriacao string `json:"dt_criacao"`
}

type Transacao struct{
	id int `json:"id_transacao"`
	proposta Proposta `json:"proposta"`
	beneficiario Beneficiario `json:"beneficiario"`
	pagador Pagador `json:"pagador"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface,function string,args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create ownership table
	err := stub.CreateTable("Transacoes", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Transacao", Type: shim.ColumnDefinition_STRING, Key: true},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed creating Transacoes table, [%v]", err)
	}

     // Set the admin
	// The metadata will contain the certificate of the administrator
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

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub,"init", args)
	}
	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "read" { //deletes an entity from its state
	    return t.getTransaction(stub, args)
	}
	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query: " + function)
}



//read function with high certs
func (t *SimpleChaincode) getTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Retrieving transaction information...")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

    // Recover the role that is allowed to make assignments
	admin, err := stub.GetState("admin")
	if err != nil {
		fmt.Printf("Error getting role [%v] \n", err)
		return nil, errors.New("Failed fetching assigner role")
	}

	callerRole, err := stub.ReadCertAttribute("role")
	if err != nil {
		fmt.Printf("Error reading attribute 'role' [%v] \n", err)
		return nil, fmt.Errorf("Failed fetching caller role. Error was [%v]", err)
	}

    caller := string(callerRole[:])
	regulator := string(admin[:])

	if caller != regulator {
		fmt.Printf("Caller is not assigner - caller %v assigner %v\n", caller, regulator)
		return nil, fmt.Errorf("The caller does not have the rights to invoke assign. Expected role [%v], caller role [%v]", regulator, caller)
	}
    
	fmt.Printf("[getTransaction] Regulator authorized! [%v]" , args[1])

    return nil,nil
}
