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

/*
Implementação iniciada por Caue Garcia Polimanti e Vitor Diego dos Santos de Sousa
*/

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)
//	"encoding/base64"
// "github.com/op/go-logging"
//var myLogger = logging.MustGetLogger("dojo_mgm")

// BoletoPropostaChaincode - implementacao do chaincode
type BoletoPropostaChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(BoletoPropostaChaincode))
	if err != nil {
		fmt.Printf("Error starting BoletoPropostaChaincode chaincode: %s", err)
	}
}

// Init resets all the things
func (t *BoletoPropostaChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debug("Init Chaincode...")
	fmt.Println("Init Chaincode...")

	// Verificação da quantidad de argumentos recebidas
	// Não estamos recebendo nenhum argumento
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Verifica se a tabela 'Proposta' existe
	fmt.Println("Verificando se a tabela 'Proposta' existe...")
	tbProposta, err := stub.GetTable("Proposta")
	if err != nil {
		fmt.Println("Falha ao executar stub.GetTable para a tabela 'Proposta'. [%v]", err)
	}
	// Se a tabela 'Proposta' já existir
	if tbProposta != nil {	
		err = stub.DeleteTable("Proposta")		// Excluir a tabela
		fmt.Println("Tabela 'Proposta' excluída.")
	}


	fmt.Println("Criando a tabela 'Proposta'...")
	// Criar tabela de Propostas	
	err = stub.CreateTable("Proposta", []*shim.ColumnDefinition{
		// Identificador da proposta (hash)
		&shim.ColumnDefinition{Name: "Id", Type: shim.ColumnDefinition_STRING, Key: true},
		// CPF do Pagador
		&shim.ColumnDefinition{Name: "cpfPagador", Type: shim.ColumnDefinition_BOOL, Key: false},
		// Status de aceite do Pagador da proposta
		&shim.ColumnDefinition{Name: "pagadorAceitou", Type: shim.ColumnDefinition_BOOL, Key: false},
		// Status de aceite do Beneficiario da proposta
		&shim.ColumnDefinition{Name: "beneficiarioAceitou", Type: shim.ColumnDefinition_BOOL, Key: false},
		// Status do Pagamento do Boleto
		&shim.ColumnDefinition{Name: "boletoPago", Type: shim.ColumnDefinition_BOOL, Key: false},
	})
	if err != nil {
		return nil, fmt.Errorf("Falha ao criar a tabela 'Proposta'. [%v]", err)
	} 
	fmt.Println("Tabela 'Proposta' criada com sucesso.")

	fmt.Println("Init Chaincode... Finalizado!")

	return nil, nil
}

// registrarProposta: função Invoke para registrar uma nova proposta, recebendo os seguintes argumentos
// args[0]: Id. Hash que identificará a proposta
// args[1]: cpfPagador. CPF do Pagador
// args[2]: pagadorAceitou. Status de aceite do Pagador da proposta
// args[3]: beneficiarioAceitou. Status de aceite do Beneficiario da proposta
// args[4]: boletoPago. Status do Pagamento do Boleto
//
func (t *BoletoPropostaChaincode) registrarProposta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//myLogger.Debug("registrarProposta...")
	fmt.Println("registrarProposta...")

	// Verifica se a quantidade de argumentos recebidas corresponde a esperada
	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}

	// Obtem os valores dos argumentos e os prepara para salvar na tabela 'Proposta'
	idProposta := args[0]
	cpfPagador := args[1]
	pagadorAceitou, err := strconv.ParseBool(args[2])
	if err != nil {
		return nil, errors.New("Failed decodinf pagadorAceitou")
	}
	beneficiarioAceitou, err := strconv.ParseBool(args[3])
	if err != nil {
		return nil, errors.New("Failed decodinf beneficiarioAceitou")
	}
	boletoPago, err := strconv.ParseBool(args[4])
	if err != nil {
		return nil, errors.New("Failed decodinf boletoPago")
	}

	// [To do] verificar identidade

	// Registra a proposta na tabela 'Proposta'
	//myLogger.Debugf("Criando Proposta Id [%s] para CPF nº [%s]", idProposta, cpfPagador)
	fmt.Println("Criando Proposta Id [%s] para CPF nº [%s]", idProposta, cpfPagador)

	ok, err := stub.InsertRow("Proposta", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: idProposta}},
			&shim.Column{Value: &shim.Column_String_{String_: cpfPagador}},
			&shim.Column{Value: &shim.Column_Bool{Bool: pagadorAceitou}},
			&shim.Column{Value: &shim.Column_Bool{Bool: beneficiarioAceitou}},
			&shim.Column{Value: &shim.Column_Bool{Bool: boletoPago}} },
	})

	if !ok && err == nil {
		return nil, errors.New("Proposta já existente.")
	}

	//myLogger.Debug("Proposta criada!")
	fmt.Println("Proposta criada!")

	return nil, err
}


// Invoke will be called for every transaction.
// Supported functions are the following:
// "init": initialize the chaincode state, used as reset
// "registrarProposta(Id, cpfPagador, pagadorAceitou, 
// beneficiarioAceitou, boletoPago)": para registrar uma nova proposta.
// Only an administrator can call this function.
// "consultarProposta(Id)": para consultar uma Proposta. 
// Only the owner of the specific asset can call this function.
// An asset is any string to identify it. An owner is representated by one of his ECert/TCert.
func (t *BoletoPropostaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debug("Invoke Chaincode...")
	fmt.Println("Invoke Chaincode...")

	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "registrarProposta" {
		// Registrar nova Proposta
		return t.registrarProposta(stub, args)
	} else if function == "consultarProposta" {
		// Consultar uma Proposta existente
		//return t.consultarProposta(stub, args)
	}
	fmt.Println("invoke did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *BoletoPropostaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debug("Query Chaincode...")
	fmt.Println("Query Chaincode...")

	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" { //read a variable
		fmt.Println("hi there " + function) //error
		return nil, nil
	}
	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query: " + function)
}
