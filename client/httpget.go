
package main

import (
    "fmt"
    "context"
    "time"

    tracelog "github.com/toniz/trace"
    _ "github.com/toniz/trace/zipkin"
)

func main() {
    err := tracelog.SetHttpExport(context.TODO(), "tracing_config.json", "PaymentClient", "v0.1.30")

    res, err := tracelog.HttpGet(context.TODO(), "http://127.0.0.1:8080/hello", "my_client")
    fmt.Println(res)
    fmt.Println(err)

    time.Sleep(10*time.Second)
}

