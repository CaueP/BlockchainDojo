# BlockchainDojo

* Pontos de uso com segurança de acesso no chaincode

#app_internal.go:
* initCryptoClients() => crypto.Init()
* Initialize the clients mapping alice, bob, charlie and dave
* to identities already defined in 'membersrvc.yaml'

* deployInternal(usamos um dos clients já registrados e inicializados) => spec(type,chaincodeId,ctorMsg,Metadata,confidentialityLevel)