package errhandling

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig()
	if err != nil {
		fmt.Printf("origin error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("\nstack trace: \n%+v\n", err)
	}
	fmt.Println("config:", string(config))

}
