// Code generated by webidl-bind. DO NOT EDIT.

package orientation

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget
// javascript.PromiseVoid

// source idl files:
// screen-orientation.idl

// transform files:
// screen-orientation.go.md

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

// enum: OrientationLockType
type OrientationLockType int

const (
	AnyOrientationLockType OrientationLockType = iota
	NaturalOrientationLockType
	LandscapeOrientationLockType
	PortraitOrientationLockType
	PortraitPrimaryOrientationLockType
	PortraitSecondaryOrientationLockType
	LandscapePrimaryOrientationLockType
	LandscapeSecondaryOrientationLockType
)

var orientationLockTypeToWasmTable = []string{
	"any", "natural", "landscape", "portrait", "portrait-primary", "portrait-secondary", "landscape-primary", "landscape-secondary",
}

var orientationLockTypeFromWasmTable = map[string]OrientationLockType{
	"any": AnyOrientationLockType, "natural": NaturalOrientationLockType, "landscape": LandscapeOrientationLockType, "portrait": PortraitOrientationLockType, "portrait-primary": PortraitPrimaryOrientationLockType, "portrait-secondary": PortraitSecondaryOrientationLockType, "landscape-primary": LandscapePrimaryOrientationLockType, "landscape-secondary": LandscapeSecondaryOrientationLockType,
}

// JSValue is converting this enum into a javascript object
func (this *OrientationLockType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this OrientationLockType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(orientationLockTypeToWasmTable) {
		return orientationLockTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// OrientationLockTypeFromJS is converting a javascript value into
// a OrientationLockType enum value.
func OrientationLockTypeFromJS(value js.Value) OrientationLockType {
	key := value.String()
	conv, ok := orientationLockTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: OrientationType
type OrientationType int

const (
	PortraitPrimaryOrientationType OrientationType = iota
	PortraitSecondaryOrientationType
	LandscapePrimaryOrientationType
	LandscapeSecondaryOrientationType
)

var orientationTypeToWasmTable = []string{
	"portrait-primary", "portrait-secondary", "landscape-primary", "landscape-secondary",
}

var orientationTypeFromWasmTable = map[string]OrientationType{
	"portrait-primary": PortraitPrimaryOrientationType, "portrait-secondary": PortraitSecondaryOrientationType, "landscape-primary": LandscapePrimaryOrientationType, "landscape-secondary": LandscapeSecondaryOrientationType,
}

// JSValue is converting this enum into a javascript object
func (this *OrientationType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this OrientationType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(orientationTypeToWasmTable) {
		return orientationTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// OrientationTypeFromJS is converting a javascript value into
// a OrientationType enum value.
func OrientationTypeFromJS(value js.Value) OrientationType {
	key := value.String()
	conv, ok := orientationTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// class: ScreenOrientation
type ScreenOrientation struct {
	domcore.EventTarget
}

// ScreenOrientationFromJS is casting a js.Wrapper into ScreenOrientation.
func ScreenOrientationFromJS(value js.Wrapper) *ScreenOrientation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ScreenOrientation{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type OrientationType (idl: OrientationType).
func (_this *ScreenOrientation) Type() OrientationType {
	var ret OrientationType
	value := _this.Value_JS.Get("type")
	ret = OrientationTypeFromJS(value)
	return ret
}

// Angle returning attribute 'angle' with
// type int (idl: unsigned short).
func (_this *ScreenOrientation) Angle() int {
	var ret int
	value := _this.Value_JS.Get("angle")
	ret = (value).Int()
	return ret
}

// Onchange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *ScreenOrientation) Onchange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnchange setting attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *ScreenOrientation) SetOnchange(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onchange", input)
}

func (_this *ScreenOrientation) Lock(orientation OrientationLockType) (_result *javascript.PromiseVoid) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := orientation.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("lock", _args[0:_end]...)
	var (
		_converted *javascript.PromiseVoid // javascript: PromiseVoid _what_return_name
	)
	_converted = javascript.PromiseVoidFromJS(_returned)
	_result = _converted
	return
}

func (_this *ScreenOrientation) Unlock() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("unlock", _args[0:_end]...)
	return
}
