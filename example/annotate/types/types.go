package types

import (
	"reflect"
	//	"github.com/AkihiroSuda/aspectgo/example/annotate/validate"
)

type ValidateFuncInfo struct {
	Name string
	Args []interface{}
}

type ValidateFuncLog = map[reflect.Value]ValidateFuncInfo

// Global variable that `main_aspect.go` will log the
// function calls into.
// `main.go` will read the logged function calls out of it.
var FuncLog = make(ValidateFuncLog)
