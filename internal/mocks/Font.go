// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	color "github.com/prasangmisra/maroto/pkg/color"
	consts "github.com/prasangmisra/maroto/pkg/consts"

	mock "github.com/stretchr/testify/mock"
)

// Font is an autogenerated mock type for the Font type
type Font struct {
	mock.Mock
}

// GetColor provides a mock function with given fields:
func (_m *Font) GetColor() color.Color {
	ret := _m.Called()

	var r0 color.Color
	if rf, ok := ret.Get(0).(func() color.Color); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(color.Color)
	}

	return r0
}

// GetFamily provides a mock function with given fields:
func (_m *Font) GetFamily() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetFont provides a mock function with given fields:
func (_m *Font) GetFont() (string, consts.Style, float64) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 consts.Style
	if rf, ok := ret.Get(1).(func() consts.Style); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(consts.Style)
	}

	var r2 float64
	if rf, ok := ret.Get(2).(func() float64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(float64)
	}

	return r0, r1, r2
}

// GetScaleFactor provides a mock function with given fields:
func (_m *Font) GetScaleFactor() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetSize provides a mock function with given fields:
func (_m *Font) GetSize() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetStyle provides a mock function with given fields:
func (_m *Font) GetStyle() consts.Style {
	ret := _m.Called()

	var r0 consts.Style
	if rf, ok := ret.Get(0).(func() consts.Style); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(consts.Style)
	}

	return r0
}

// SetColor provides a mock function with given fields: _a0
func (_m *Font) SetColor(_a0 color.Color) {
	_m.Called(_a0)
}

// SetFamily provides a mock function with given fields: family
func (_m *Font) SetFamily(family string) {
	_m.Called(family)
}

// SetFont provides a mock function with given fields: family, style, size
func (_m *Font) SetFont(family string, style consts.Style, size float64) {
	_m.Called(family, style, size)
}

// SetSize provides a mock function with given fields: size
func (_m *Font) SetSize(size float64) {
	_m.Called(size)
}

// SetStyle provides a mock function with given fields: style
func (_m *Font) SetStyle(style consts.Style) {
	_m.Called(style)
}
