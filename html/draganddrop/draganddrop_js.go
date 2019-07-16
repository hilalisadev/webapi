// Code generated by webidl-bind. DO NOT EDIT.

package draganddrop

import "syscall/js"

// using following types:

// source idl files:
// html.idl

// transform files:
// html.go.md

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// callback: FunctionStringCallback
type FunctionStringCallbackFunc func(data string)

// FunctionStringCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type FunctionStringCallback js.Func

func FunctionStringCallbackToJS(callback FunctionStringCallbackFunc) *FunctionStringCallback {
	if callback == nil {
		return nil
	}
	ret := FunctionStringCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 string // javascript: DOMString data
		)
		_p0 = (args[0]).String()
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func FunctionStringCallbackFromJS(_value js.Value) FunctionStringCallbackFunc {
	return func(data string) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := data
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}
