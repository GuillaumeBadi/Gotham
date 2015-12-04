
package gexception

func NewException (e error) *Exception {
    return &Exception{
        Message: e.Error(),
    }
}
