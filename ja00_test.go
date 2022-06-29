package errr
import (
	"testing"
	"errors"
	"fmt"
)

func Test_ (t *testing.T) {
	errr01 := errors.New ("error 01")
	errr02 := Errr_Estb ("02", "error 02", errr01)
	errr03 := Errr_Estb ("03", "error 03", errr02)
	errr04 := Errr_Estb ("04", "error 04")
	
	fmt.Println ("x1:", errr01.Error ())
	fmt.Println ("x2:", errr02.Error ())
	fmt.Println ("x3:", errr03.Error ())
	fmt.Println ("x4:", errr04.Error ())
	
	fmt.Println ("")
	fmt.Println ("y2:", errr02.Dtll)
	fmt.Println ("y3:", errr03.Dtll)
	fmt.Println ("y4:", errr04.Dtll)
	
	fmt.Println ("")
	fmt.Println ("z2:", errr02.Iddd)
	fmt.Println ("z3:", errr03.Iddd)
	fmt.Println ("z4:", errr04.Iddd)
}
