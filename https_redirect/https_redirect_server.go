package main
// HTTPS to HTTP redirect Seerver
// Redirects any HTTPS attempts to the non HTTPS version of the hostname
// Ron Egli - Github.com/smugzombie

// Generate a cert to be used: openssl req -new -x509 -keyout server.pem -out server.pem -days 4000 -nodes 
// or use your own

import (
    "net/http"
    "log"
)

func MainServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    var redirect = "<meta http-equiv=\"refresh\" content=\"0; url=http://" + req.Host + "/\" />"
    w.Write([]byte(redirect))
}

func main() {
    http.HandleFunc("/", MainServer)
    err := http.ListenAndServeTLS(":443", "server.pem","server.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
