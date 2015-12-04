
package gjson

import(
    "io/ioutil"
    "encoding/json"
    . "core/exception"
)

func Write (data *GothamJson, filename string) *Exception {
    output, marshalException := json.MarshalIndent(data, "", "  ")
    if marshalException != nil {
        return NewException(marshalException)
    }
    err := ioutil.WriteFile(filename, []byte(output), 0644)
    if err != nil {
        return NewException(err)
    }
    return nil
}
