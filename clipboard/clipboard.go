// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package clipboard

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html/datatransfer"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// datatransfer.DataTransfer
// datatransfer.PromiseDataTransfer
// domcore.Event
// domcore.EventTarget
// javascript.PromiseString
// javascript.PromiseVoid

// source idl files:
// clipboard-apis.idl

// transform files:
// clipboard-apis.go.md

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

// dictionary: ClipboardEventInit
type ClipboardEventInit struct {
	Bubbles       bool
	Cancelable    bool
	Composed      bool
	ClipboardData *datatransfer.DataTransfer
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ClipboardEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.ClipboardData.JSValue()
	out.Set("clipboardData", value3)
	return out
}

// ClipboardEventInitFromJS is allocating a new
// ClipboardEventInit object and copy all values from
// input javascript object
func ClipboardEventInitFromJS(value js.Wrapper) *ClipboardEventInit {
	input := value.JSValue()
	var out ClipboardEventInit
	var (
		value0 bool                       // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool                       // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool                       // javascript: boolean {composed Composed composed}
		value3 *datatransfer.DataTransfer // javascript: DataTransfer {clipboardData ClipboardData clipboardData}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	if input.Get("clipboardData").Type() != js.TypeNull && input.Get("clipboardData").Type() != js.TypeUndefined {
		value3 = datatransfer.DataTransferFromJS(input.Get("clipboardData"))
	}
	out.ClipboardData = value3
	return &out
}

// class: Clipboard
type Clipboard struct {
	domcore.EventTarget
}

// ClipboardFromJS is casting a js.Wrapper into Clipboard.
func ClipboardFromJS(value js.Wrapper) *Clipboard {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Clipboard{}
	ret.Value_JS = input
	return ret
}

func (_this *Clipboard) Read() (_result *datatransfer.PromiseDataTransfer) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("read", _args[0:_end]...)
	var (
		_converted *datatransfer.PromiseDataTransfer // javascript: Promise _what_return_name
	)
	_converted = datatransfer.PromiseDataTransferFromJS(_returned)
	_result = _converted
	return
}

func (_this *Clipboard) ReadText() (_result *javascript.PromiseString) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("readText", _args[0:_end]...)
	var (
		_converted *javascript.PromiseString // javascript: Promise _what_return_name
	)
	_converted = javascript.PromiseStringFromJS(_returned)
	_result = _converted
	return
}

func (_this *Clipboard) Write(data *datatransfer.DataTransfer) (_result *javascript.PromiseVoid) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("write", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *Clipboard) WriteText(data string) (_result *javascript.PromiseVoid) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("writeText", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

// class: ClipboardEvent
type ClipboardEvent struct {
	domcore.Event
}

// ClipboardEventFromJS is casting a js.Wrapper into ClipboardEvent.
func ClipboardEventFromJS(value js.Wrapper) *ClipboardEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ClipboardEvent{}
	ret.Value_JS = input
	return ret
}

func NewClipboardEvent(_type string, eventInitDict *ClipboardEventInit) (_result *ClipboardEvent) {
	_klass := js.Global().Get("ClipboardEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ClipboardEvent // javascript: ClipboardEvent _what_return_name
	)
	_converted = ClipboardEventFromJS(_returned)
	_result = _converted
	return
}

// ClipboardData returning attribute 'clipboardData' with
// type datatransfer.DataTransfer (idl: DataTransfer).
func (_this *ClipboardEvent) ClipboardData() *datatransfer.DataTransfer {
	var ret *datatransfer.DataTransfer
	value := _this.Value_JS.Get("clipboardData")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = datatransfer.DataTransferFromJS(value)
	}
	return ret
}
