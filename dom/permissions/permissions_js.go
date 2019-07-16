// Code generated by webidl-bind. DO NOT EDIT.

package permissions

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget
// javascript.Object
// javascript.PromiseFinally

// source idl files:
// permissions.idl
// promises.idl

// transform files:
// permissions.go.md
// promises.go.md

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

// enum: PermissionState
type PermissionState int

const (
	GrantedPermissionState PermissionState = iota
	DeniedPermissionState
	PromptPermissionState
)

var permissionStateToWasmTable = []string{
	"granted", "denied", "prompt",
}

var permissionStateFromWasmTable = map[string]PermissionState{
	"granted": GrantedPermissionState, "denied": DeniedPermissionState, "prompt": PromptPermissionState,
}

// JSValue is converting this enum into a javascript object
func (this *PermissionState) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PermissionState) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(permissionStateToWasmTable) {
		return permissionStateToWasmTable[idx]
	}
	panic("unknown input value")
}

// PermissionStateFromJS is converting a javascript value into
// a PermissionState enum value.
func PermissionStateFromJS(value js.Value) PermissionState {
	key := value.String()
	conv, ok := permissionStateFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: PromiseTemplateOnFulfilled
type PromisePermissionStatusOnFulfilledFunc func(value *PermissionStatus)

// PromisePermissionStatusOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionStatusOnFulfilled js.Func

func PromisePermissionStatusOnFulfilledToJS(callback PromisePermissionStatusOnFulfilledFunc) *PromisePermissionStatusOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionStatusOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *PermissionStatus // javascript: PermissionStatus value
		)
		_p0 = PermissionStatusFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePermissionStatusOnFulfilledFromJS(_value js.Value) PromisePermissionStatusOnFulfilledFunc {
	return func(value *PermissionStatus) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromisePermissionStatusOnRejectedFunc func(reason js.Value)

// PromisePermissionStatusOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionStatusOnRejected js.Func

func PromisePermissionStatusOnRejectedToJS(callback PromisePermissionStatusOnRejectedFunc) *PromisePermissionStatusOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionStatusOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePermissionStatusOnRejectedFromJS(_value js.Value) PromisePermissionStatusOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// class: PermissionStatus
type PermissionStatus struct {
	domcore.EventTarget
}

// PermissionStatusFromJS is casting a js.Wrapper into PermissionStatus.
func PermissionStatusFromJS(value js.Wrapper) *PermissionStatus {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PermissionStatus{}
	ret.Value_JS = input
	return ret
}

// State returning attribute 'state' with
// type PermissionState (idl: PermissionState).
func (_this *PermissionStatus) State() PermissionState {
	var ret PermissionState
	value := _this.Value_JS.Get("state")
	ret = PermissionStateFromJS(value)
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *PermissionStatus) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *PermissionStatus) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

// class: Permissions
type Permissions struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Permissions) JSValue() js.Value {
	return _this.Value_JS
}

// PermissionsFromJS is casting a js.Wrapper into Permissions.
func PermissionsFromJS(value js.Wrapper) *Permissions {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Permissions{}
	ret.Value_JS = input
	return ret
}

func (_this *Permissions) Query(permissionDesc *javascript.Object) (_result *PromisePermissionStatus) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := permissionDesc.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("query", _args[0:_end]...)
	var (
		_converted *PromisePermissionStatus // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStatusFromJS(_returned)
	_result = _converted
	return
}

// class: Promise
type PromisePermissionStatus struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromisePermissionStatus) JSValue() js.Value {
	return _this.Value_JS
}

// PromisePermissionStatusFromJS is casting a js.Wrapper into PromisePermissionStatus.
func PromisePermissionStatusFromJS(value js.Wrapper) *PromisePermissionStatus {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromisePermissionStatus{}
	ret.Value_JS = input
	return ret
}

func (_this *PromisePermissionStatus) Then(onFulfilled *PromisePermissionStatusOnFulfilled, onRejected *PromisePermissionStatusOnRejected) (_result *PromisePermissionStatus) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromisePermissionStatus // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStatusFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionStatus) Catch(onRejected *PromisePermissionStatusOnRejected) (_result *PromisePermissionStatus) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromisePermissionStatus // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStatusFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionStatus) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePermissionStatus) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromisePermissionStatus // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionStatusFromJS(_returned)
	_result = _converted
	return
}
