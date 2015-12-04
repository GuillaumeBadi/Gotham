
package gjson

import(
  "io/ioutil"
  "encoding/json"
  . "core/exception"
)

func Parse (filename string) (*GothamJson, *Exception) {
    var data GothamJson
    content, readErr := ioutil.ReadFile(filename)
    if readErr != nil {
        return nil, NewException(readErr)
    }
    parseError := json.Unmarshal(content, &data)
    if parseError != nil {
        return nil, NewException(parseError)
    }
    return &data, nil
}
