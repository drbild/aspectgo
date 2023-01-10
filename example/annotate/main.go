package main

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/AkihiroSuda/aspectgo/example/annotate/types"
	"github.com/AkihiroSuda/aspectgo/example/annotate/validate"
)

var validHello validate.ValidateFunc = validate.StringInSlice(validate.ValidHelloValues(), false)
var validGoodbye validate.ValidateFunc = validate.StringInSlice(validate.ValidGoodbyeValues(), false)
var validMonthCode validate.ValidateFunc = validate.IntBetween(1, 12)

func sayHello(s string) {
	valid := validHello(s)
	if valid {
		fmt.Println(s + " world")
	} else {
		fmt.Println("hello" + " world")
	}
}

func logFunc(name string, f interface{}) {
	fmt.Println(name)
	if funcInfo, ok := types.FuncLog[reflect.ValueOf(f)]; ok {
		fmt.Println("\tType:", funcInfo.Name)
		switch funcInfo.Name {
		case "StringInSlice":
			values, _ := json.Marshal(funcInfo.Args[0])
			ignore, _ := json.Marshal(funcInfo.Args[1])
			fmt.Println("\tValues:", string(values))
			fmt.Println("\tIgnoreCase:", string(ignore))
		case "IntBetween":
			min, _ := json.Marshal(funcInfo.Args[0])
			max, _ := json.Marshal(funcInfo.Args[1])
			fmt.Println("\tMin", string(min))
			fmt.Println("\tMax", string(max))
		}
	} else {
		fmt.Println("\tNot Found")
	}

}

func main() {
	sayHello("hola")
	sayHello("bonjour")
	sayHello("goodbye")

	logFunc("validHello", validHello)
	logFunc("validGoodbye", validGoodbye)
	logFunc("validMonthCode", validMonthCode)
}
