package tsgutils

import (
	"fmt"
)

/*
 other utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
  ternary operator, replace other language: a == b ? c : d
*/
func IIIInterfaceOperator(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func InterfaceToString(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	}
	return ""
}

func FmtPrintln(info interface{}) {
	if info != nil {
		switch t := info.(type) {
		case struct{}:
			fmt.Println(StructToJson(t))
		default:
			fmt.Println(t)
		}
	} else {
		fmt.Println(nil)
	}
}
