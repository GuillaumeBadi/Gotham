
package gjson

import(
  "io/ioutil"
  . "gotham/exception"
)

func parseJson (filename string) (*GothamJson, Exception) {
    content, err := io.ReadFile(filename);
    if err != nil {
        return nil, NewException(err)
    }
}
