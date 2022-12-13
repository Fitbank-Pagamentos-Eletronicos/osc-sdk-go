package requests


import (
   "fmt"
   "net/http"
   "io/ioutil"
   "strings"
   "encoding/json"
    "modulo/src/domains"
)

var dataDocument = domains.Document {
    Tipo: domains.IDENTITY_BACK,
    MimeType: domains.IMAGE_JPEG,
    Name: "44983829865_CNH_20102022_CNH Aberta.jpg",
    Base64: "9j/4AAQSkZJRgABAQAAAQABAAD/7QDWUGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAA",
}

func DocumentRequest(ID string) string {
    url := "https://demo-api.easycredito.com.br/api/external//v2/process/document/" + ID
    method := "PUT"

    fmt.Println("URL: ",url)
    simpleDocumentJson, _ := json.Marshal(dataDocument)

    payload := strings.NewReader(string(simpleDocumentJson))

    client := &http.Client {}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
    fmt.Println(err)
        return ""
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", "Bearer " + GetToken())

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

    fmt.Println(string(body))

    return string(body)


}
