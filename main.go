package main

import(
    "net/http"
    "fmt"
    "log"
    "os"
)

func main()  {
    status := true
    var i string
    for status != false{
        fmt.Println("Enter u, h or exit: ")
        fmt.Scanln(&i)
        switch i {
            case "u" : reponseHeaders()
            case "exit","e" : os.Exit(0)
            case "h" : log.Println("Type u for request header, type exit to exit")
            default : log.Println("Type u to request headers, type exit to exit")
        }
    }
}

func reponseHeaders(){
    var (
        x string
        i string
        h string
        )
    fmt.Println("Enter 1 for http, 2 for https")
    fmt.Scanln(&i)
    switch i {
        case "1" : h = "http://"
        case "2" : h = "https://"
        default : log.Println("Type 1 for http, 2 for https")
    }
    fmt.Println("URL: ")
    fmt.Scanln(&x)
    resp, err := http.Head(h + x)
    if err != nil {
        log.Println(err)
    }
    if resp.StatusCode != http.StatusOK {
        log.Println("Server returned not 200 code")
    }

    for k,v := range resp.Header {
        log.Println("key: ", k, "value: ", v)
    }
}

