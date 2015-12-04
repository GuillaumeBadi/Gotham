
package gjson

type Dependency     struct {
    Name            string          `json:name`
    Source          string          `json:source`
}

type GothamJson     struct {
    Name            string          `json:name`
    Version         string          `json:version`
    Source          string          `json:source`
    Licence         string          `json:licence`
    Author          string          `json:author`
    Dependencies    []Dependency    `json:dependencies`
}
