
package dependency

import(
    "core/gjson"
)

func Install (addings []Dependency) {
    // foreach dep: go get dep && if gothamProject(dep) Install(dep)
}

func Add (data *GothamJson, addings []Dependency) (data GothamJson) {
    data.Dependencies = append(data.Dependencies, addings...)
    return
}
