package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {
    fmt.Println(getCWD())
}

func getCWD() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    return dir
}
