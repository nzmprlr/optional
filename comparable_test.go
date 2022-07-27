package optional_test

import (
	"fmt"
	"testing"

	"github.com/nzmprlr/optional"
)

func Test_Comparable_Get(t *testing.T) {
	cases := []int{
		0,
		1,
		-1,
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if optional.Comparable(c).Get() != c {
				t.Errorf("comparable Get failed")
			}
		})
	}
}

func Test_Comparable_Empty_and_Present(t *testing.T) {
	cases := []string{
		"",
		"test",
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Comparable(c)

			if e, p := o.Empty(), o.Present(); e == p {
				t.Errorf("comparable Empty and Present are same")
			}
		})
	}
}

func Test_Comparable_IfEmpty(t *testing.T) {
	cases := []struct {
		v        int
		willCall bool
		called   bool
	}{
		{0, true, false},
		{1, false, false},
		{-1, false, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Comparable(c.v).IfEmpty(func(i int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("comparable IfEmpty failed")
			}
		})
	}
}

func Test_Comparable_IfPresent(t *testing.T) {
	cases := []struct {
		v        int
		willCall bool
		called   bool
	}{
		{0, false, false},
		{1, true, false},
		{-1, true, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			optional.Comparable(c.v).IfPresent(func(i int) {
				c.called = true
			})

			if !c.willCall && c.called {
				t.Errorf("comparable IfPresent failed")
			}
		})
	}
}

func Test_Comparable_If_Else(t *testing.T) {
	cases := []struct {
		v         int
		condition bool
		e         int
		result    int
	}{
		{0, false, 0, 0},
		{0, true, 0, 0},
		{0, false, 1, 1},
		{1, false, 0, 0},
		{1, true, 0, 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			o := optional.Comparable(c.v).If(func(i int) bool {
				return c.condition
			}).Else(c.e)

			if o != c.result {
				t.Errorf("interface If-Else failed")
			}
		})
	}
}
