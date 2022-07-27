package optional_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nzmprlr/optional"
)

func Test_Slice_Get(t *testing.T) {
	cases := [][]int{
		nil,
		[]int{},
		[]int{0, 1, -1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(optional.Slice(c).Get(), c) {
				t.Errorf("slice Get failed")
			}
		})
	}
}

func Test_Slice_Empty_and_Present(t *testing.T) {
	cases := [][]int{
		nil,
		[]int{},
		[]int{0, 1, -1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Slice(c)

			if e, p := o.Empty(), o.Present(); e == p {
				t.Errorf("slice Empty and Present are same")
			}
		})
	}
}

func Test_Slice_IfEmpty(t *testing.T) {
	cases := []struct {
		v        []int
		willCall bool
		called   bool
	}{
		{nil, true, false},
		{[]int{}, true, false},
		{[]int{1}, false, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Slice(c.v).IfEmpty(func(m []int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("slice IfEmpty failed")
			}
		})
	}
}

func Test_Slice_IfPresent(t *testing.T) {
	cases := []struct {
		v        []int
		willCall bool
		called   bool
	}{
		{nil, false, false},
		{[]int{}, true, false},
		{[]int{1}, true, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Slice(c.v).IfPresent(func(m []int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("slice IfPresent failed")
			}
		})
	}
}

func Test_Slice_If_Else(t *testing.T) {
	cases := []struct {
		v         []int
		condition bool
		e         []int
		result    []int
	}{
		{nil, false, nil, nil},
		{nil, true, nil, nil},
		{nil, false, []int{}, []int{}},
		{[]int{1}, false, nil, nil},
		{[]int{1}, true, nil, []int{1}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Slice(c.v).If(func(m []int) bool {
				return c.condition
			}).Else(c.e)

			if !reflect.DeepEqual(o, c.result) {
				t.Errorf("slice If-Else failed")
			}
		})
	}
}
