// Code generated by webidl-bind. DO NOT EDIT.

package url

import "syscall/js"

import (
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/html/media"
)

// using following types:
// file.Blob
// media.MediaSource

// source idl files:
// url.idl

// transform files:
// url.go.md

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

// callback: URLSearchParamsForEach
type URLSearchParamsForEachFunc func(currentValue string, currentIndex int, listObj *URLSearchParams)

// URLSearchParamsForEach is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type URLSearchParamsForEach js.Func

func URLSearchParamsForEachToJS(callback URLSearchParamsForEachFunc) *URLSearchParamsForEach {
	if callback == nil {
		return nil
	}
	ret := URLSearchParamsForEach(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 string           // javascript: USVString currentValue
			_p1 int              // javascript: long currentIndex
			_p2 *URLSearchParams // javascript: URLSearchParams listObj
		)
		_p0 = (args[0]).String()
		_p1 = (args[1]).Int()
		_p2 = URLSearchParamsFromJS(args[2])
		callback(_p0, _p1, _p2)
		// returning no return value
		return nil
	}))
	return &ret
}

func URLSearchParamsForEachFromJS(_value js.Value) URLSearchParamsForEachFunc {
	return func(currentValue string, currentIndex int, listObj *URLSearchParams) {
		var (
			_args [3]interface{}
			_end  int
		)
		_p0 := currentValue
		_args[0] = _p0
		_end++
		_p1 := currentIndex
		_args[1] = _p1
		_end++
		_p2 := listObj.JSValue()
		_args[2] = _p2
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: URLSearchParamsEntryIteratorValue
type URLSearchParamsEntryIteratorValue struct {
	Value []js.Value
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *URLSearchParamsEntryIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Value))
	for __idx0, __seq_in0 := range _this.Value {
		__seq_out0 := __seq_in0
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// URLSearchParamsEntryIteratorValueFromJS is allocating a new
// URLSearchParamsEntryIteratorValue object and copy all values from
// input javascript object
func URLSearchParamsEntryIteratorValueFromJS(value js.Wrapper) *URLSearchParamsEntryIteratorValue {
	input := value.JSValue()
	var out URLSearchParamsEntryIteratorValue
	var (
		value0 []js.Value // javascript: sequence<any> {value Value value}
		value1 bool       // javascript: boolean {done Done done}
	)
	__length0 := input.Get("value").Length()
	__array0 := make([]js.Value, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 js.Value
		__seq_in0 := input.Get("value").Index(__idx0)
		__seq_out0 = __seq_in0
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: URLSearchParamsKeyIteratorValue
type URLSearchParamsKeyIteratorValue struct {
	Value string
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *URLSearchParamsKeyIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// URLSearchParamsKeyIteratorValueFromJS is allocating a new
// URLSearchParamsKeyIteratorValue object and copy all values from
// input javascript object
func URLSearchParamsKeyIteratorValueFromJS(value js.Wrapper) *URLSearchParamsKeyIteratorValue {
	input := value.JSValue()
	var out URLSearchParamsKeyIteratorValue
	var (
		value0 string // javascript: USVString {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = (input.Get("value")).String()
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// dictionary: URLSearchParamsValueIteratorValue
type URLSearchParamsValueIteratorValue struct {
	Value string
	Done  bool
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *URLSearchParamsValueIteratorValue) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Value
	out.Set("value", value0)
	value1 := _this.Done
	out.Set("done", value1)
	return out
}

// URLSearchParamsValueIteratorValueFromJS is allocating a new
// URLSearchParamsValueIteratorValue object and copy all values from
// input javascript object
func URLSearchParamsValueIteratorValueFromJS(value js.Wrapper) *URLSearchParamsValueIteratorValue {
	input := value.JSValue()
	var out URLSearchParamsValueIteratorValue
	var (
		value0 string // javascript: USVString {value Value value}
		value1 bool   // javascript: boolean {done Done done}
	)
	value0 = (input.Get("value")).String()
	out.Value = value0
	value1 = (input.Get("done")).Bool()
	out.Done = value1
	return &out
}

// class: URL
type URL struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URL) JSValue() js.Value {
	return _this.Value_JS
}

// URLFromJS is casting a js.Wrapper into URL.
func URLFromJS(value js.Wrapper) *URL {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URL{}
	ret.Value_JS = input
	return ret
}

func CreateObjectURL(blob *file.Blob) (_result string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("createObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func RevokeObjectURL(url string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("revokeObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	_method.Invoke(_args[0:_end]...)
	return
}

func CreateObjectURL2(mediaSource *media.MediaSource) (_result string) {
	_klass := js.Global().Get("URL")
	_method := _klass.Get("createObjectURL")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := mediaSource.JSValue()
	_args[0] = _p0
	_end++
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func NewURL(url string, base *string) (_result *URL) {
	_klass := js.Global().Get("URL")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	if base != nil {
		_p1 := base
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *URL // javascript: URL _what_return_name
	)
	_converted = URLFromJS(_returned)
	_result = _converted
	return
}

// Href returning attribute 'href' with
// type string (idl: USVString).
func (_this *URL) Href() string {
	var ret string
	value := _this.Value_JS.Get("href")
	ret = (value).String()
	return ret
}

// ToString is an alias for Href.
func (_this *URL) ToString() string {
	return _this.Href()
}

// SetHref setting attribute 'href' with
// type string (idl: USVString).
func (_this *URL) SetHref(value string) {
	input := value
	_this.Value_JS.Set("href", input)
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *URL) Origin() string {
	var ret string
	value := _this.Value_JS.Get("origin")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: USVString).
func (_this *URL) Protocol() string {
	var ret string
	value := _this.Value_JS.Get("protocol")
	ret = (value).String()
	return ret
}

// SetProtocol setting attribute 'protocol' with
// type string (idl: USVString).
func (_this *URL) SetProtocol(value string) {
	input := value
	_this.Value_JS.Set("protocol", input)
}

// Username returning attribute 'username' with
// type string (idl: USVString).
func (_this *URL) Username() string {
	var ret string
	value := _this.Value_JS.Get("username")
	ret = (value).String()
	return ret
}

// SetUsername setting attribute 'username' with
// type string (idl: USVString).
func (_this *URL) SetUsername(value string) {
	input := value
	_this.Value_JS.Set("username", input)
}

// Password returning attribute 'password' with
// type string (idl: USVString).
func (_this *URL) Password() string {
	var ret string
	value := _this.Value_JS.Get("password")
	ret = (value).String()
	return ret
}

// SetPassword setting attribute 'password' with
// type string (idl: USVString).
func (_this *URL) SetPassword(value string) {
	input := value
	_this.Value_JS.Set("password", input)
}

// Host returning attribute 'host' with
// type string (idl: USVString).
func (_this *URL) Host() string {
	var ret string
	value := _this.Value_JS.Get("host")
	ret = (value).String()
	return ret
}

// SetHost setting attribute 'host' with
// type string (idl: USVString).
func (_this *URL) SetHost(value string) {
	input := value
	_this.Value_JS.Set("host", input)
}

// Hostname returning attribute 'hostname' with
// type string (idl: USVString).
func (_this *URL) Hostname() string {
	var ret string
	value := _this.Value_JS.Get("hostname")
	ret = (value).String()
	return ret
}

// SetHostname setting attribute 'hostname' with
// type string (idl: USVString).
func (_this *URL) SetHostname(value string) {
	input := value
	_this.Value_JS.Set("hostname", input)
}

// Port returning attribute 'port' with
// type string (idl: USVString).
func (_this *URL) Port() string {
	var ret string
	value := _this.Value_JS.Get("port")
	ret = (value).String()
	return ret
}

// SetPort setting attribute 'port' with
// type string (idl: USVString).
func (_this *URL) SetPort(value string) {
	input := value
	_this.Value_JS.Set("port", input)
}

// Pathname returning attribute 'pathname' with
// type string (idl: USVString).
func (_this *URL) Pathname() string {
	var ret string
	value := _this.Value_JS.Get("pathname")
	ret = (value).String()
	return ret
}

// SetPathname setting attribute 'pathname' with
// type string (idl: USVString).
func (_this *URL) SetPathname(value string) {
	input := value
	_this.Value_JS.Set("pathname", input)
}

// Search returning attribute 'search' with
// type string (idl: USVString).
func (_this *URL) Search() string {
	var ret string
	value := _this.Value_JS.Get("search")
	ret = (value).String()
	return ret
}

// SetSearch setting attribute 'search' with
// type string (idl: USVString).
func (_this *URL) SetSearch(value string) {
	input := value
	_this.Value_JS.Set("search", input)
}

// SearchParams returning attribute 'searchParams' with
// type URLSearchParams (idl: URLSearchParams).
func (_this *URL) SearchParams() *URLSearchParams {
	var ret *URLSearchParams
	value := _this.Value_JS.Get("searchParams")
	ret = URLSearchParamsFromJS(value)
	return ret
}

// Hash returning attribute 'hash' with
// type string (idl: USVString).
func (_this *URL) Hash() string {
	var ret string
	value := _this.Value_JS.Get("hash")
	ret = (value).String()
	return ret
}

// SetHash setting attribute 'hash' with
// type string (idl: USVString).
func (_this *URL) SetHash(value string) {
	input := value
	_this.Value_JS.Set("hash", input)
}

func (_this *URL) ToJSON() (_result string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("toJSON", _args[0:_end]...)
	var (
		_converted string // javascript: USVString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// class: URLSearchParams
type URLSearchParams struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URLSearchParams) JSValue() js.Value {
	return _this.Value_JS
}

// URLSearchParamsFromJS is casting a js.Wrapper into URLSearchParams.
func URLSearchParamsFromJS(value js.Wrapper) *URLSearchParams {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URLSearchParams{}
	ret.Value_JS = input
	return ret
}

func (_this *URLSearchParams) Append(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("append", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Delete(name string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("delete", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Get(name string) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("get", _args[0:_end]...)
	var (
		_converted *string // javascript: USVString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *URLSearchParams) GetAll(name string) (_result []string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getAll", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<USVString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *URLSearchParams) Has(name string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("has", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *URLSearchParams) Set(name string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("set", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Sort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("sort", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) ToString() (_result string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("toString", _args[0:_end]...)
	var (
		_converted string // javascript: USVString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *URLSearchParams) Entries() (_result *URLSearchParamsEntryIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("entries", _args[0:_end]...)
	var (
		_converted *URLSearchParamsEntryIterator // javascript: URLSearchParamsEntryIterator _what_return_name
	)
	_converted = URLSearchParamsEntryIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *URLSearchParams) ForEach(callback *URLSearchParamsForEach, optionalThisForCallbackArgument interface{}) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if callback != nil {
		__callback0 = (*callback).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if optionalThisForCallbackArgument != nil {
		_p1 := optionalThisForCallbackArgument
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("forEach", _args[0:_end]...)
	return
}

func (_this *URLSearchParams) Keys() (_result *URLSearchParamsKeyIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("keys", _args[0:_end]...)
	var (
		_converted *URLSearchParamsKeyIterator // javascript: URLSearchParamsKeyIterator _what_return_name
	)
	_converted = URLSearchParamsKeyIteratorFromJS(_returned)
	_result = _converted
	return
}

func (_this *URLSearchParams) Values() (_result *URLSearchParamsValueIterator) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("values", _args[0:_end]...)
	var (
		_converted *URLSearchParamsValueIterator // javascript: URLSearchParamsValueIterator _what_return_name
	)
	_converted = URLSearchParamsValueIteratorFromJS(_returned)
	_result = _converted
	return
}

// class: URLSearchParamsEntryIterator
type URLSearchParamsEntryIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URLSearchParamsEntryIterator) JSValue() js.Value {
	return _this.Value_JS
}

// URLSearchParamsEntryIteratorFromJS is casting a js.Wrapper into URLSearchParamsEntryIterator.
func URLSearchParamsEntryIteratorFromJS(value js.Wrapper) *URLSearchParamsEntryIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URLSearchParamsEntryIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *URLSearchParamsEntryIterator) Next() (_result *URLSearchParamsEntryIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *URLSearchParamsEntryIteratorValue // javascript: URLSearchParamsEntryIteratorValue _what_return_name
	)
	_converted = URLSearchParamsEntryIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: URLSearchParamsKeyIterator
type URLSearchParamsKeyIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URLSearchParamsKeyIterator) JSValue() js.Value {
	return _this.Value_JS
}

// URLSearchParamsKeyIteratorFromJS is casting a js.Wrapper into URLSearchParamsKeyIterator.
func URLSearchParamsKeyIteratorFromJS(value js.Wrapper) *URLSearchParamsKeyIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URLSearchParamsKeyIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *URLSearchParamsKeyIterator) Next() (_result *URLSearchParamsKeyIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *URLSearchParamsKeyIteratorValue // javascript: URLSearchParamsKeyIteratorValue _what_return_name
	)
	_converted = URLSearchParamsKeyIteratorValueFromJS(_returned)
	_result = _converted
	return
}

// class: URLSearchParamsValueIterator
type URLSearchParamsValueIterator struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *URLSearchParamsValueIterator) JSValue() js.Value {
	return _this.Value_JS
}

// URLSearchParamsValueIteratorFromJS is casting a js.Wrapper into URLSearchParamsValueIterator.
func URLSearchParamsValueIteratorFromJS(value js.Wrapper) *URLSearchParamsValueIterator {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &URLSearchParamsValueIterator{}
	ret.Value_JS = input
	return ret
}

func (_this *URLSearchParamsValueIterator) Next() (_result *URLSearchParamsValueIteratorValue) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("next", _args[0:_end]...)
	var (
		_converted *URLSearchParamsValueIteratorValue // javascript: URLSearchParamsValueIteratorValue _what_return_name
	)
	_converted = URLSearchParamsValueIteratorValueFromJS(_returned)
	_result = _converted
	return
}
