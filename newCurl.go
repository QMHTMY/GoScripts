package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

func init(){
    if len(os.Args) != 2 {
        log.Fatal("Usage: " + os.Args[0] + " url")
    }
}

func main(){
    resp, err := http.Get(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    io.Copy(os.Stdout, resp.Body)
    if err := resp.Body.Close(); err != nil {
        fmt.Println(err)
    }
    fmt.Println("")
}
