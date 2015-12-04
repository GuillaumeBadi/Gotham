
package gjson

import(
    "core/dependency"
)

type GothamJson     struct {
    Name            string          `json:name,omitempty"`
    Version         string          `json:version,omitempty"`
    Source          string          `json:source,omitempty"`
    Licence         string          `json:licence,omitempty"`
    Author          string          `json:author,omitempty"`
    Dependencies    []Dependency    `json:dependencies,omitempty"`
}
