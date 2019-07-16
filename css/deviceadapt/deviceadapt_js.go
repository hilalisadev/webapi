// Code generated by webidl-bind. DO NOT EDIT.

package deviceadapt

import "syscall/js"

import (
	"github.com/gowebapi/webapi/css/cssom"
)

// using following types:
// cssom.CSSRule
// cssom.CSSStyleDeclaration

// source idl files:
// css-device-adapt.idl

// transform files:
// css-device-adapt.go.md

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

// class: CSSViewportRule
type CSSViewportRule struct {
	cssom.CSSRule
}

// CSSViewportRuleFromJS is casting a js.Wrapper into CSSViewportRule.
func CSSViewportRuleFromJS(value js.Wrapper) *CSSViewportRule {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CSSViewportRule{}
	ret.Value_JS = input
	return ret
}

// Style returning attribute 'style' with
// type cssom.CSSStyleDeclaration (idl: CSSStyleDeclaration).
func (_this *CSSViewportRule) Style() *cssom.CSSStyleDeclaration {
	var ret *cssom.CSSStyleDeclaration
	value := _this.Value_JS.Get("style")
	ret = cssom.CSSStyleDeclarationFromJS(value)
	return ret
}
