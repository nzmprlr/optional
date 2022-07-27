package optional_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nzmprlr/optional"
)

func Test_Map_Get(t *testing.T) {
	cases := []map[int]int{
		nil,
		map[int]int{},
		map[int]int{1: 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(optional.Map(c).Get(), c) {
				t.Errorf("map Get failed")
			}
		})
	}
}

func Test_Map_Empty_and_Present(t *testing.T) {
	cases := []map[int]int{
		nil,
		map[int]int{},
		map[int]int{1: 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Map(c)

			if e, p := o.Empty(), o.Present(); e == p {
				t.Errorf("map Empty and Present are same")
			}
		})
	}
}

func Test_Map_IfEmpty(t *testing.T) {
	cases := []struct {
		v        map[int]int
		willCall bool
		called   bool
	}{
		{nil, true, false},
		{map[int]int{}, true, false},
		{map[int]int{1: 1}, false, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Map(c.v).IfEmpty(func(m map[int]int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("map IfEmpty failed")
			}
		})
	}
}

func Test_Map_IfPresent(t *testing.T) {
	cases := []struct {
		v        map[int]int
		willCall bool
		called   bool
	}{
		{nil, false, false},
		{map[int]int{}, true, false},
		{map[int]int{1: 1}, true, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Map(c.v).IfPresent(func(m map[int]int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("map IfPresent failed")
			}
		})
	}
}

func Test_Map_If_Else(t *testing.T) {
	cases := []struct {
		v         map[int]int
		condition bool
		e         map[int]int
		result    map[int]int
	}{
		{nil, false, nil, nil},
		{nil, true, nil, nil},
		{nil, false, map[int]int{}, map[int]int{}},
		{map[int]int{1: 1}, false, nil, nil},
		{map[int]int{1: 1}, true, nil, map[int]int{1: 1}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Map(c.v).If(func(m map[int]int) bool {
				return c.condition
			}).Else(c.e)

			if !reflect.DeepEqual(o, c.result) {
				t.Errorf("map If-Else failed")
			}
		})
	}
}
