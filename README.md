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

rhoncus id placerat at, dictum vitae lectus. Etiam tristique pellentesque lorem, eu
consequat tellus pulvinar et. Vestibulum diam arcu, eleifend quis vestibulum at, auctor
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras nulla lorem,
in ligula. Ut ut hendrerit nunc, a facilisis nisl. Nulla sollicitudin interdum venenatis.
Etiam at.

### Fluxograma
    
```mermaid 
    sequenceDiagram
       participant Client
       participant SDK
       participant Auth
       participant API
       
       
       Client ->> +SDK: OSC.createInstance(client_id, client_secret)
       SDK ->> -Client: intancia osc  
  
  ```