# osc-sdk-go

Branch: Features/TR-6298 - Escrita de exemplo Signup

## Descrição do Projeto

SDK para integração com a API da OSC.
    

## Estrutura dos diretórios

        
        ├── osc-sdk-go
        │   ├── Idea
        │   ├── bin
        │   ├── pkg
        │   ├── src
        │   │   ├── domains
        │   │   ├── main
        │   │   ├── requests
        │   │   ├── utils
        ├── test
        │   └── ...
       

## Como executar os testes
    
    go test -v ./...  Executa todos os testes
    go test -v Address_test.go  Executa o teste de endereço

## Descrição dos métodos

Para a requisição de _Signup_ é utilizado o método POST para fazer requisição para a API da OSC. Essa função recebe uma _struct_
_SigupMatch_ com todos os dados necessários para a requisição. A função retorna uma _string_ com o corpo da função. O corpo da retorna dados do cliente
como "id", "name", "cpf", "dataCriação", e " dataAtualização". Caso ocorra algum erro, a função retorna uma _string_ com a mensagem de erro.

## Exemplo de uso

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
package main
import ("fmt", "strings")

type OSC struct {
  clientId  string
  clientSecret string
  authorized bool
  api  *API
  auth *Auth
}


func (osc *OSC) CreateInstance(clientId string, clientSecret string) *OSC{
    // criar uma instância do objecto do OSC
    
    osc = &OSC{clientId: clientId, clientSecret: clientSecret}
    
    // retorna o valor do objeto instanciado
    return osc
}

```

### Signup + respostas

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras nulla lorem, rhoncus id placerat at, dictum vitae lectus. Etiam tristique pellentesque lorem, eu consequat tellus pulvinar et. Vestibulum diam arcu, eleifend quis vestibulum at, auctor in ligula. Ut ut hendrerit nunc, a facilisis nisl. Nulla sollicitudin interdum venenatis. Etiam at.

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
    
```
#### Codificação
```
!
```


### Signup + Proposal

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras nulla lorem, rhoncus id placerat at, dictum vitae lectus. Etiam tristique pellentesque lorem, eu consequat tellus pulvinar et. Vestibulum diam arcu, eleifend quis vestibulum at, auctor in ligula. Ut ut hendrerit nunc, a facilisis nisl. Nulla sollicitudin interdum venenatis. Etiam at.

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
```
!
```

### Fluxo Completo

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras nulla lorem, rhoncus id placerat at, dictum vitae lectus. Etiam tristique pellentesque lorem, eu consequat tellus pulvinar et. Vestibulum diam arcu, eleifend quis vestibulum at, auctor in ligula. Ut ut hendrerit nunc, a facilisis nisl. Nulla sollicitudin interdum venenatis. Etiam at.

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
```
!
```
# !!!!!
#### Codificação
    package requests

    import (
        "fmt"
        "net/http"
        "io/ioutil"
        "strings"
        "encoding/json"
        "modulo/src/domains"
    )

    func SignupMatchRequest() string {

        url:= "https://demo-api.easycredito.com.br/api/external//v2.1/process/signup"
        method := "POST"
    
        simpleToJson, _ := json.Marshal(data)
    
        payload := strings.NewReader(string(simpleToJson))
    
        client := &http.Client{}
        req, err := http.NewRequest(method, url, payload)
    
        if err != nil {
            fmt.Println(err)
            return  ""
        }
        
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Accept", "application/json")
        req.Header.Add("Authorization", "Bearer " + GetToken()) // Para a autenticação é necessário o token esse token é feito na função GetToken()
    
        res, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
            return ""
        }
        defer res.Body.Close()
    
        body, err := ioutil.ReadAll(res.Body)
    
        if err != nil {
            fmt.Println(err)
            return ""
        }
    
        return string(body)
}

### Função GetToken

    func GetToken () string {
        if auth.Access_token == "" || auth.Expire_at == "" {
        auth = Auth()
        return auth.Access_token

    }

    var expireAt, _ = time.Parse("2006-01-02T15:04:05.000Z", auth.Expire_at)

    if time.Now().After(expireAt) {
        auth = Auth()
        return auth.Access_token
    }

    return auth.Access_token

}


   
    

## Execução 
Para a execução desse função _SignupMatchRequest_ é necessário importar o pacote _requests_ e chamar a função _SignupMatchRequest_.
package main.

    import (
        "fmt"
        "modulo/src/requests"
    )

    func main() {
       requests.SignupMatchRequest()
       
    }

### Execução da função main

    go run main.go

### Retorno da função
        {
            "id": "1k8l4f7bhv8jc33p4ardd847b41f7694041879443333f9d5513",
            "name": "Carlos Alexandre",
            "status": "SIGNUP_ANALISIS",
            "cpf": 36761586011,
            "dateCreated": "2022-12-26T12:22:05.949Z",
            "lastUpdated": "2022-12-26T12:22:05.949Z"
        }
