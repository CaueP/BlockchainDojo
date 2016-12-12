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
	"os"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("dojo_mgm")

// BoletoPropostaChaincode example simple Chaincode implementation
type BoletoPropostaChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(BoletoPropostaChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *BoletoPropostaChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Init Chaincode...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *BoletoPropostaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Invoke Chaincode...")

	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}
	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *BoletoPropostaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Query Chaincode...")

	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" { //read a variable
		fmt.Println("hi there " + function) //error
		return nil, nil
	}
	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query: " + function)
}
