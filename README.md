# osc-sdk-go

Branch: Features/TR-6298 - Escrita de exemplo Signup

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

    ├── osc-sdk-go
    │   ├── Idea
    │   ├── bin
    │   ├── pkg
    │   ├── src
    │   │    ├── domains
    │   │    ├── main
    │   │    ├── requests
    │   │    ├── utils
    ├── test


##  :hammer_and_wrench: Como executar os testes
    
    go test -v ./...  Executa todos os testes
    go test -v Address_test.go  Executa o teste de endereço

## :page_with_curl: Descrição do método

Para a requisição de _Signup_ é utilizado o método POST para fazer requisição para a API da OSC. Essa função recebe uma _struct_
_SigupMatch_ com todos os dados necessários para a requisição. A função retorna uma _string_ com o corpo da função. O corpo da retorna dados do cliente
como "id", "name", "cpf", "dataCriação", e " dataAtualização". Caso ocorra algum erro, a função retorna uma _string_ com a mensagem de erro.


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
package main

import (
    "fmt", 
    "strings"
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

```


