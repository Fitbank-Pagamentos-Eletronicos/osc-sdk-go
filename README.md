# osc-sdk-go

 ## :blue_book: Descrição do Projeto

SDK para integração com a API da OSC.
    
## :rocket: Instalando

Para clonar o projeto e executar essa aplicação, você precisará do [Git](https://git-scm.com) e do [Go](https://golang.org/) instalados em seu computador.
Depois disso execute os seguintes comandos:

```bash
    # Clone esse repositório
    git clone https://github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go.git
    
    # Entre no repositório
    cd osc-sdk-go
    cd src/main
    
    # Execute o projeto
    go run main.go


```

## :file_folder: Estrutura dos diretórios

    📦osc-sdk-go
    ┣ 📂bin
    ┣ 📂Idea
    ┣ 📂pkg
    ┣ 📂src
    ┃ ┣ 📂domains
    ┃ ┣ 📂main
    ┃ ┣ 📂requests
    ┃ ┗ 📂utils
    ┗ 📂test


##  :hammer_and_wrench: Como executar os testes


```bash
    # Entrar no diretório de testes
    cd test
    
    # Executar todos os testes
    go test -v ./...  Executa todos os testes
    
    # Execute apenas um test
    go test -v Address_test.go


```
## :page_with_curl: Métodos

Os métodos estáo disponíveis na pasta `src/requests` e são:
 
- [x] `OAuth` - Criação do token de autenticação para o uso em endpoints. O resultado  é o `AuthSucess`.
- [x] `DocumentRequest` - Envia um documento para análise. O resultado é o `DocumentResponse`.
- [x] `ProposalRequest` - Recolhe e valida dados necessários para a criação de propostas de acordo com os tipos de produtos selecionados. O resultado é o `PipelineProposal`.
- [x] `PubSub` - Obtém o ID do projeto, ID do tópico e outras coisas. O resultado é o `PubSubResponse`.
- [x] `PubsubSubscribe` - Usa os dados retornados do `PubSub` e cria um ouvinte.
- [x] `SimpleSignup` - Realiza o cadastro de usuários(pode retornar erro caso o usuário já exista). O resultado é o `SignupResponse`.
- [x] `SignupMatchRequest` - Faz a inscrição de usuários e retorna os produtos de créditos com maior chance de aprovação. O resultado é o `SignupMatchResponse`. 
- [x]  `OSC` - Realiza a criação de instancias de OSC.

## :dart: Exemplo de uso

### Signup

#### Fluxograma
```mermaid
sequenceDiagram
    participant Client
    participant SDK
    participant Auth
    participant API

    Client->>+SDK: OSC.createInstance(client_id, client_secret)
    SDK-->>-Client: instancia osc
    
    Client->>+SDK: osc.signup(signupObject)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: signup(signupJson, access_token)
        API-->>-SDK: pipelineJson
    SDK-->>-Client: pipeline instance
```
#### Codificação

```Go
package examples

import (
	"fmt"
	"modulo/src/domains"
	"modulo/src/osc"
)

func main() {
	var instance, _ = osc.CreateInstance("", "", "default")
	var data = domains.SimpleSignup{...} 
    
	var pipeline = instance.SimpleSignup(data)
	fmt.Printf("%s", pipeline.Id)
	
}

```

### Signup + Proposal
#### Fluxograma
```mermaid
sequenceDiagram
    participant Client
    participant SDK
    participant Auth
    participant API
    participant PubSub

    Client->>+SDK: OSC.createInstance(client_id, client_secret)
    SDK-->>-Client: instancia osc
    
    Client->>+SDK: osc.setResponseListening(listeningFunction)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: pubsub(access_token)
        API-->>-SDK: pubsubConfig
        par Abre socket
            SDK->>PubSub: subscription(pubsubConfig)
        end
    SDK-->>-Client: pipeline instance
    
    Client->>+SDK: osc.signup(signupObject)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: signup(signupJson, access_token)
        API-->>-SDK: pipelineJson
    SDK-->>-Client: pipeline instance
    API->>PubSub: publica(signupResponse)
    PubSub-->>SDK: subscriptionSocket(signupResponse)
    SDK-->>Client: listeningFunction(signupResponse)
    
    Client->>+SDK: osc.proposal(pipeline_id, proposalObject)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: proposal(pipeline_id, proposalJson, access_token)
        API-->>-SDK: pipelineJson
    SDK-->>-Client: pipeline instance
    API->>PubSub: publica(proposalResponse)
    PubSub-->>SDK: subscriptionSocket(proposalResponse)
    SDK-->>Client: listeningFunction(proposalResponse)
```
#### Codificação

```Go
package examples

import (
	"fmt"
	"golang.org/x/net/ipv6"
	"modulo/src/domains"
	"modulo/src/osc"
)

func main() {
	var instance, _ = osc.CreateInstance("", "", "default")

	instance.SetResponseListening(func(pipeline domains.Pipeline, err boll) {
		switch pipeline.Status {
		case domains.SIGNUP_ANALISIS:
			fmt.Printf("Async %s cadastro em analise", pipeline.Id)
		case domains.SIGNUP_COMPLETED:
			fmt.Printf("Async %s cadastro concluido", pipeline.Id)
			Proposal(pipeline.Id)
		case domains.SIGNUP_DENIED:
			fmt.Printf("Async %s cadastro regeitado", pipeline.Id)
		case domains.PROPOSAL_ANALISIS:
			fmt.Printf("Async %s proposta em analise", pipeline.Id)
		case domains.PROPOSAL_CREATED:
			fmt.Printf("Async %s proposta criada", pipeline.Id)
		case domains.PROPOSAL_DENIED:
			fmt.Printf("Async %s proposta regeitada", pipeline.Id)
		}
	})
	Signup()
}
func Signup() {
    var data = domains.SimpleSignup{...} 
    var pipeline = instance.SimpleSignup(data)
    fmt.Printf("%s", pipeline.Id)
}

func Proposal(pipelineId string) {
    var data = domains.ProposalReq{...} 
    var pipeline = instance.Proposal(pipelineId, data)
    fmt.Printf("%s", pipeline.Id)
}

```
### PubSub

#### Fluxograma
```mermaid
sequenceDiagram
    participant Client
    participant SDK
    participant Auth
    participant API
    participant PubSub

    Client->>+SDK: OSC.createInstance(client_id, client_secret)
    SDK-->>-Client: osc instance

    Client->>+SDK: osc.setResponseListening(listeningFunction)
        opt Not Authorized
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: pubsub(access_token)
        API-->>-SDK: pubsubConfig
        par Open socket
            SDK->>PubSub: subscription(pubsubConfig)
        end
    SDK-->>-Client: pipeline instance
```
#### Codificação
```Go
 package main
 
 import (
    "fmt", 
    "strings"
    "cloud.google.com/go/pubsub"
    "google.golang.org/api/option"
 )

    type OSC struct {
      clientId  string
      clientSecret string
      authorized bool
      api  *API
      auth *Auth
    }
    

    func (osc *OSC) createInstance(clientId string, clientSecret string) *OSC {
      osc.clientId = clientId
      osc.clientSecret = clientSecret
      osc.authorized = false
      osc.api = new(API)
      osc.auth = new(Auth)
      
      return osc
          
    }
    
    func (osc *OSC) setResponseListening(listeningFunction func(message *pubsub.Message)) {
       if !osc.authorized {
          osc.authorized = true
          osc.auth.authorize(osc.clientId, osc.clientSecret, "pubsub")
       }
          
       pubsubConfig := osc.api.pubsub(osc.auth.accessToken)
       ctx := context.Background()
       client, err := pubsub.NewClient(ctx, pubsubConfig.ProjectId, option.WithCredentialsJSON([]byte(pubsubConfig.Credentials)))
       if err != nil {
         fmt.Println(err)
       }
       defer client.Close()
          
       sub := client.Subscription(pubsubConfig.Subscription)
       err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
          msg.Ack()
          listeningFunction(string(msg.Data))
       })
       if err != nil {
         fmt.Println(err)
       }
    }
```
### Fluxo completo

#### Fluxograma
```mermaid
sequenceDiagram
    participant Client
    participant SDK
    participant Auth
    participant API
    participant PubSub

 

    Client->>+SDK: OSC.createInstance(client_id, client_secret)
    SDK-->>-Client: instancia osc

    Client->>+SDK: osc.setResponseListening(listeningFunction)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: pubsub(access_token)
        API-->>-SDK: pubsubConfig
        par Abre socket
            SDK->>PubSub: subscription(pubsubConfig)
        end
    SDK-->>-Client: pipeline instance

    Client->>+SDK: osc.signup(signupObject)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: signup(signupJson, access_token)
        API-->>-SDK: pipelineJson
    SDK-->>-Client: pipeline instance
    API->>PubSub: publica(signupResponse)
    PubSub-->>SDK: subscriptionSocket(signupResponse)
    SDK-->>Client: listeningFunction(signupResponse)

    Client->>+SDK: osc.proposal(pipeline_id, proposalObject)
        opt Não autorizado 
            SDK->>+Auth: auth(client_id, client_secret, scope)
            Auth-->>-SDK: access_token
        end
        SDK->>+API: proposal(pipeline_id, proposalJson, access_token)
        API-->>-SDK: pipelineJson
    SDK-->>-Client: pipeline instance

    API->>PubSub: publica(proposalResponse)
    PubSub-->>SDK: subscriptionSocket(proposalResponse)
    SDK-->>Client: listeningFunction(proposalResponse)

    par Continuara recevbendo atualização de status durante o fluxo
        API->>PubSub: publica(proposalStatusUpdateResponse)
        PubSub-->>SDK: subscriptionSocket(proposalResponse)
        SDK-->>Client: listeningFunction(proposalResponse)
    and Nesta etapa o envio de documentos esta liberado
        Client->>+SDK: osc.document(pipeline_id, documentObject)
            opt Não autorizado 
                SDK->>+Auth: auth(client_id, client_secret, scope)
                Auth-->>-SDK: access_token
            end
            SDK->>+API: proposal(pipeline_id, documentJson, access_token)
            API-->>-SDK: documentResponseJson
        SDK-->>-Client: documentResponse instance
    and Caso alguma proposta retorne que tem contratos para asinatura
        Client->>+SDK: osc.getContracts(customerServiceNumber)
            opt Não autorizado 
                SDK->>+Auth: auth(client_id, client_secret, scope)
                Auth-->>-SDK: access_token
            end
            SDK->>+API: getContracts(customerServiceNumber, access_token)
            API-->>-SDK: contractsResponseJson
        SDK-->>-Client: contractsResponse instance

        Client->>+SDK: osc.SignContracts(customerServiceNumber, contractsObject)
            opt Não autorizado 
                SDK->>+Auth: auth(client_id, client_secret, scope)
                Auth-->>-SDK: access_token
            end
            SDK->>+API: SignContracts(customerServiceNumber, contractsObject, access_token)
            API-->>-SDK: signContractsResponseJson
        SDK-->>-Client: signContractsResponse instance
    end
```
#### Codificação
```Go
package main
 
 import (
    "fmt", 
    "strings"
    "cloud.google.com/go/pubsub"
    "google.golang.org/api/option"
 )

    type OSC struct {
      clientId  string
      clientSecret string
      authorized bool
      api  *API
      auth *Auth
    }
    
     func (osc *OSC) setResponseListening(listeningFunction func(message *pubsub.Message)) {
       if !osc.authorized {
          osc.authorized = true
          osc.auth.authorize(osc.clientId, osc.clientSecret, "pubsub")
       }
     
     func (osc *OSC) Proposal(pipelineId string, proposalObject interface{}) {
         if !osc.authorized {
            osc.auth.auth(osc.clientId, osc.clientSecret, "pubsub")
            osc.authorized = true
         }
     }
     
     func SinupMatchRequest(signupMatch *SignupMatch, auth *Auth) {
         signupMatchJson, _ := json.Marshal(signupMatch)
         signupMatchResponseJson := api.signup(signupMatchJson, auth.accessToken)
         signupMatchResponse := SignupMatchResponse{}
         json.Unmarshal(signupMatchResponseJson, &signupMatchResponse)
         return signupMatchResponse
     }

```

