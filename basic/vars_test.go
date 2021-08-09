package basic

import (
	"fmt"
	"testing"
)

func TestVars(t *testing.T) {
	fmt.Println("variable zero values:")
	variableZeroValues()

	fmt.Println("variable define:")
	variableDefine()

	fmt.Println()
	fmt.Println("iota: ")
	iotaDefine()

	fmt.Println("package variable: ")
	packageVariableDefine()
}
