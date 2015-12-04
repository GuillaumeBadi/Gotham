
package dependency

type Dependency     struct {
    Name            string          `json:name,omitempty"`
    Source          string          `json:source,omitempty"`
}
