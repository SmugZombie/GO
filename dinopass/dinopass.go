package main

// dinopass.go
// Generates a password based on DinoPass
// Ron Egli - github/smugzombie

import (
        "io/ioutil"
        "log"
        "net/http"
        "fmt"
        "github.com/atotto/clipboard"
)

func main() {
        response, err := http.Get("http://www.dinopass.com/password/strong")
        if err != nil {
                log.Fatal(err)
        } else {
                defer response.Body.Close()
                bodyBytes, err := ioutil.ReadAll(response.Body)
                clipboard.WriteAll(string(bodyBytes));
                fmt.Print(string(bodyBytes))
                _ = err
        }
}
