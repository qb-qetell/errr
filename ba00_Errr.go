package errr
import (
	"fmt"
)

type Errr struct {
	Iddd string
	Dtll string
	Csss error
}
	func Errr_Estb (iddd, dtll string, csss ... error) (*Errr) {
		errr := &Errr {}
		errr.Iddd = iddd
		errr.Dtll = dtll
		if len (csss) > 0 {
			errr.Csss = csss [0]
		}
		return errr
	}
	func (objc *Errr) Error () (string) {
		dtll := objc.Dtll
		if objc.Csss != nil {
			dtll = fmt.Sprintf ("%s [%s]", objc.Dtll, objc.Csss.Error ())
		}
		return dtll
	}
