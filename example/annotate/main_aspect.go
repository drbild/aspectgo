package main

import (
	"reflect"
	"regexp"

	asp "github.com/AkihiroSuda/aspectgo/aspect"
	"github.com/AkihiroSuda/aspectgo/example/annotate/types"
)

// implements interface asp.Aspect
type StringInSliceAspect struct {
}

// Executed on compilation-time
func (a *StringInSliceAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("github.com/AkihiroSuda/aspectgo/example/annotate/validate")
	s := pkg + ".StringInSlice"
	return asp.NewCallPointcutFromRegexp(s)
}

// Executed ONLY on runtime
func (a *StringInSliceAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()
	res := ctx.Call(args)

	// Log the function name and the args used to create the function
	f := reflect.ValueOf(res[0])
	types.FuncLog[f] = types.ValidateFuncInfo{
		"StringInSlice",
		args,
	}

	return res
}

// implements interface asp.Aspect
type IntBetweenAspect struct {
}

func (a *IntBetweenAspect) Pointcut() asp.Pointcut {
	pkg := regexp.QuoteMeta("github.com/AkihiroSuda/aspectgo/example/annotate/validate")
	s := pkg + ".IntBetween"
	return asp.NewCallPointcutFromRegexp(s)
}

// Executed ONLY on runtime
func (a *IntBetweenAspect) Advice(ctx asp.Context) []interface{} {
	args := ctx.Args()
	res := ctx.Call(args)

	// Log the function name and the args used to create the function
	f := reflect.ValueOf(res[0])
	types.FuncLog[f] = types.ValidateFuncInfo{
		"IntBetween",
		args,
	}

	return res
}
