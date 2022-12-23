package main

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/AkihiroSuda/aspectgo/example/annotate/types"
	"github.com/AkihiroSuda/aspectgo/example/annotate/validate"
)

var validHello validate.ValidateFunc = validate.StringInSlice(validate.ValidValues(), false)

func sayHello(s string) {
	valid := validHello(s)
	if valid {
		fmt.Println(s + " world")
	} else {
		fmt.Println("hello" + " world")
	}
}

func main() {
	sayHello("hola")
	sayHello("bonjour")
	sayHello("goodbye")

	if validHelloInfo, ok := types.FuncLog[reflect.ValueOf(validHello)]; ok {
		fmt.Println("validHello", validHelloInfo.Name)
		fmt.Println("\tType:", validHelloInfo.Name)
		switch validHelloInfo.Name {
		case "StringInSlice":
			values, _ := json.Marshal(validHelloInfo.Args[0])
			ignore, _ := json.Marshal(validHelloInfo.Args[1])
			fmt.Println("\tValues:", string(values))
			fmt.Println("\tIgnoreCase:", string(ignore))
		}
	}

}
