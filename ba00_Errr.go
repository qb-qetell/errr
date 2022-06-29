package errr
import (
	"fmt"
)

type Errr struct {
	Iddd string
	Dtll string
	csss error
}
	func Errr_Estb (iddd, dtll string, csss ... error) (*Errr) {
		errr := &Errr {}
		errr.Iddd = iddd
		errr.Dtll = dtll
		if len (csss) > 0 {
			errr.csss = csss [0]
		}
		return errr
	}
	func (objc *Errr) Error () (string) {
		dtll := objc.Dtll
		if objc.csss != nil {
			dtll = fmt.Sprintf ("%s [%s]", dtll, objc.csss.Error ())
		}
		return dtll
	}
