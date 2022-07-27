package optional_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nzmprlr/optional"
)

type testInterface interface {
	Test()
}

type testStruct struct{}

func (testStruct) Test() {}

func Test_Error(t *testing.T) {
	cases := []error{
		nil,
		errors.New(""),
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Error(c)

			if e, p := o.Empty(), o.Present(); e == p {
				t.Errorf("interface Empty and Present are same")
			}
		})
	}
}

func Test_Interface_Get(t *testing.T) {
	cases := []any{
		nil,
		testInterface(nil),
		testStruct{},
		&testStruct{},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if optional.Interface(c).Get() != c {
				t.Errorf("interface Get failed")
			}
		})
	}
}

func Test_Interface_Empty_and_Present(t *testing.T) {
	cases := []any{
		nil,
		error(nil),
		testInterface(nil),
		testStruct{},
		&testStruct{},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Interface(c)

			if e, p := o.Empty(), o.Present(); e == p {
				t.Errorf("interface Empty and Present are same")
			}
		})
	}
}

func Test_Interface_IfEmpty(t *testing.T) {
	cases := []struct {
		v        any
		willCall bool
		called   bool
	}{
		{nil, true, false},
		{error(nil), true, false},
		{testInterface(nil), true, false},
		{testStruct{}, false, false},
		{&testStruct{}, false, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Interface(c.v).IfEmpty(func(a any) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("interface IfEmpty failed")
			}
		})
	}
}

func Test_Interface_IfPresent(t *testing.T) {
	cases := []struct {
		v        any
		willCall bool
		called   bool
	}{
		{nil, false, false},
		{error(nil), false, false},
		{testInterface(nil), false, false},
		{testStruct{}, true, false},
		{&testStruct{}, true, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Interface(c.v).IfPresent(func(a any) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("interface IfPresent failed")
			}
		})
	}
}

func Test_Interface_If_Else(t *testing.T) {
	cases := []struct {
		v         any
		condition bool
		e         any
		result    any
	}{
		{nil, false, nil, nil},
		{nil, true, nil, nil},
		{nil, false, testStruct{}, testStruct{}},
		{testStruct{}, false, nil, nil},
		{testStruct{}, true, nil, testStruct{}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Interface(c.v).If(func(a any) bool {
				return c.condition
			}).Else(c.e)

			if o != c.result {
				t.Errorf("interface If-Else failed")
			}
		})
	}
}
