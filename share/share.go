// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package share

import js "github.com/gowebapi/webapi/core/js"

// using following types:

// source idl files:
// web-share.idl

// transform files:
// web-share.go.md

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

// dictionary: ShareData
type ShareData struct {
	Title string
	Text  string
	Url   string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ShareData) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Title
	out.Set("title", value0)
	value1 := _this.Text
	out.Set("text", value1)
	value2 := _this.Url
	out.Set("url", value2)
	return out
}

// ShareDataFromJS is allocating a new
// ShareData object and copy all values from
// input javascript object
func ShareDataFromJS(value js.Wrapper) *ShareData {
	input := value.JSValue()
	var out ShareData
	var (
		value0 string // javascript: USVString {title Title title}
		value1 string // javascript: USVString {text Text text}
		value2 string // javascript: USVString {url Url url}
	)
	value0 = (input.Get("title")).String()
	out.Title = value0
	value1 = (input.Get("text")).String()
	out.Text = value1
	value2 = (input.Get("url")).String()
	out.Url = value2
	return &out
}
