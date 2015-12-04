
package main

import(
    "fmt"
    "core/gjson"
)

func main() {
    file := "./assets/gotham.json"
    json, parseException := gjson.Parse(file)
    if parseException != nil {
        fmt.Printf("Error: %+v\n", parseException)
        return
    }
    fmt.Printf("Json: %+v\n", json)
    writeException := gjson.Write(json, "./output.json")
    if writeException != nil {
        fmt.Printf("Error: %+v\n", writeException)
        return
    }
    return
}
