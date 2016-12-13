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
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)
//	"encoding/base64"
// "github.com/op/go-logging"
//var myLogger = logging.MustGetLogger("dojo_mgm")

// BoletoPropostaChaincode - implementacao do chaincode
type BoletoPropostaChaincode struct {
}

// Tipo Proposta para retornar a consulta JSON
type Proposta struct {
    id string `json:"id_proposta"`
	cpfPagador string `json:"cpf_pagador"`
	pagadorAceitou bool `json:"pagador_aceitou"`
	beneficiarioAceitou bool `json:"beneficiario_aceitou"`
	boletoPago bool `json:"boleto_pago"`
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
    
	// Verificação da quantidade de argumentos recebidas
	// Não estamos recebendo nenhum argumento
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Set the admin
	// The metadata will contain the certificate of the administrator
	// adminCert, err := stub.GetCallerMetadata() 
	// if err != nil {
	// 	fmt.Println("Failed getting metadata")
	// 	return nil, errors.New("Failed getting metadata.")
	// }
	// if len(adminCert) == 0 {
	// 	fmt.Printf("Invalid admin certificate. Empty.")
	// 	return nil, errors.New("Invalid admin certificate. Empty.")
	// }

	// fmt.Printf("The administrator is [%x]", adminCert)

	// stub.PutState("admin", adminCert)

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
		&shim.ColumnDefinition{Name: "cpfPagador", Type: shim.ColumnDefinition_STRING, Key: false},
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
	fmt.Println("Criando Proposta Id [" + idProposta + "] para CPF nº ["+ cpfPagador +"]")
	fmt.Printf("pagadorAceitou: " + strconv.FormatBool(pagadorAceitou)) 
	fmt.Printf(" | beneficiarioAceitou: " + strconv.FormatBool(beneficiarioAceitou))
	fmt.Printf(" | boletoPago: " + strconv.FormatBool(boletoPago) + "\n")

	ok, err := stub.InsertRow("Proposta", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: idProposta}},
			&shim.Column{Value: &shim.Column_String_{String_: cpfPagador}},
			&shim.Column{Value: &shim.Column_Bool{Bool: pagadorAceitou}},
			&shim.Column{Value: &shim.Column_Bool{Bool: beneficiarioAceitou}},
			&shim.Column{Value: &shim.Column_Bool{Bool: boletoPago}} },
	})

	if !ok && err == nil {
		// Atualmente está retornando que a Proposta já existe, mas podemos implementar o update da Proposta
		return nil, errors.New("Proposta já existente.")
	}

	//myLogger.Debug("Proposta criada!")
	fmt.Println("Proposta criada!")

	return nil, err
}

// consultarProposta: função Query para consultar uma proposta existente, recebendo os seguintes argumentos
// args[0]: Id. Hash da proposta
//
func (t *BoletoPropostaChaincode) consultarProposta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//myLogger.Debug("consultarProposta...")
	fmt.Println("consultarProposta...")

	var valAsBytes []byte

	// Verifica se a quantidade de argumentos recebidas corresponde a esperada
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Obtem os valores dos argumentos e os prepara para salvar na tabela 'Proposta'
	idProposta := args[0]

	// [To do] verificar identidade

	// Define o valor de coluna a ser buscado
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: idProposta}}
	columns = append(columns, col1)

	// Consulta a proposta na tabela 'Proposta'
	row, err := stub.GetRow("Proposta", columns)
	if err != nil {
		fmt.Println("Erro ao obter Proposta [%s]: [%s]", string(idProposta), err)
		return nil, fmt.Errorf("Erro ao obter Proposta [%s]: [%s]", string(idProposta), err)
	}

	fmt.Println("Query finalizada [% x]", row.Columns[1].GetBytes())

	// objeto Proposta
	var resProposta Proposta
	resProposta.id = row.Columns[0].GetString_()
	resProposta.cpfPagador = row.Columns[1].GetString_()
	resProposta.pagadorAceitou = row.Columns[2].GetBool()
	resProposta.beneficiarioAceitou = row.Columns[3].GetBool()
	resProposta.boletoPago = row.Columns[4].GetBool()

	fmt.Println("Valores da tabela: [%s], [%s], [%b], [%b], [%b]", row.Columns[0].GetString_(), row.Columns[1].GetString_(), row.Columns[2].GetBool(), row.Columns[3].GetBool(), row.Columns[3].GetBool())

	fmt.Println("Proposta: [%s], [%s], [%b], [%b], [%b]", resProposta.id, resProposta.cpfPagador, resProposta.pagadorAceitou, resProposta.beneficiarioAceitou, resProposta.boletoPago)

	valAsBytes, err = json.Marshal(resProposta)
	return valAsBytes, nil
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
	if function == "consultarProposta" { //read a variable
		// Consultar uma Proposta existente
		return t.consultarProposta(stub, args)
	}
	// if function == "consultarTodas"{
	// 	 return t.consultarTodasAsPropostas(stub,args)
	// }
	fmt.Println("query did not find func: " + function) //error

	return nil, errors.New("Received unknown function query: " + function)
}

// func (t *BoletoPropostaChaincode) consultarTodasAsPropostas(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
//     fmt.Println("[consultarTodasAsPropostas] Verificando Autorização de Admin...")
    
// 	    // Recover the role that is allowed to make assignments
// 	admin, err := stub.GetState("admin")
// 	if err != nil {
// 		fmt.Printf("Error getting role [%v] \n", err)
// 		return nil, errors.New("Failed fetching assigner role")
// 	}

// 	callerRole, err := stub.ReadCertAttribute("role")
// 	if err != nil {
// 		fmt.Printf("Error reading attribute 'role' [%v] \n", err)
// 		return nil, fmt.Errorf("Failed fetching caller role. Error was [%v]", err)
// 	}

//     caller := string(callerRole[:])
// 	regulator := string(admin[:])

// 	if caller != regulator {
// 		fmt.Printf("Caller is not admin - caller %v admin %v\n", caller, regulator)
// 		return nil, fmt.Errorf("The caller does not have the rights to invoke assign. Expected role [%v], caller role [%v]", regulator, caller)
// 	}
    
// 	fmt.Printf("[getTransaction] Regulator authorized! [%v]" , args[0])

//      return nil,nil
// }
