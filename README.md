# osc-sdk-go

Branch: Features/TR-6278

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

Para a requisição de autenticação é necessário passar o client_id e client_secret, que são gerados no momento da criação do projeto.
Temos também um método responsável por gerar o token de acesso, que é utilizado para as demais requisições.


## Exemplo de uso
### Autenticação

    package requests
    

    import (
        "encoding/base64"
        "encoding/json"
        "io/ioutil"
        "fmt"
        "modulo/src/domains"
        "net/http"
        "strings"
    )

    var DataBase = domains.Auth{
        Client_id:     "carlos.lima--------957a-f8c494b3328c",
        Client_secret: "47a734df8b08a57cc16023be9c85d64076eb005450f30dea504baa2f2a22a09a",
        Scopes:        []string{"api-external"},
    }

    func OAuth() domains.AuthSucess {
        url := "https://auth.easycredito.com.br/client/auth"
        method := "POST"
        payload := strings.NewReader("grant_type=client_credentials&client_id=" + 
        DataBase.Client_id + "&client_secret=" + DataBase.Client_secret + "&scope=api-external")

       client := &http.Client{}
       req, err := http.NewRequest(method, url, payload)

       if err != nil {
          fmt.Println(err)
          return domains.AuthSucess{}
       }
       req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

       res, err := client.Do(req)
       if err != nil {
          fmt.Println(err)
          return domains.AuthSucess{}
       }

         defer res.Body.Close()
            body, err := ioutil.ReadAll(res.Body)
            if err != nil {
                fmt.Println(err)
                return domains.AuthSucess{}
            }
            var auth domains.AuthSucess
            json.Unmarshal(body, &auth)
            return auth
    }

## Execução 
Para a execução desse função _OAuth_ é necessário importar o pacote _requests_ e chamar a função _OAuth_.
package main. Após a execução dessa função, será retornado um objeto do tipo _AuthSucess_ que contém o token de acesso e o tempo de expiração.

    import (
        "fmt"
        "modulo/src/requests"
    )

    func main() {
        auth := requests.OAuth()
        fmt.Println(auth)
    }

### Execução da função main

    go run main.go
