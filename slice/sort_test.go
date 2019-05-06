package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleInt64s() {

	a := []int64{4, 5, 6, 2, 1, 2}

	Int64s(a)

	fmt.Println(a)

	// Output:
	// [1 2 2 4 5 6]
}

type ssCase [2][]interface{}

type s struct {
	Name string
	V    int
}

type m map[string]interface{}

type stu struct {
	Name string
	V    int
	Age  int
}

func TestSort(t *testing.T) {
	var nilResult = []interface{}{}
	var cases = []ssCase{
		ssCase{nilResult, nilResult},
		ssCase{
			[]interface{}{
				s{`3`, 3},
				s{`1`, 1},
				s{`2`, 2},
			},
			[]interface{}{
				s{`1`, 1},
				s{`2`, 2},
				s{`3`, 3},
			},
		},

		ssCase{
			[]interface{}{
				m{`Name`: `3`, `V`: 3},
				m{`Name`: `1`, `V`: 1},
				m{`Name`: `2`, `V`: 2},
			},
			[]interface{}{
				m{`Name`: `1`, `V`: 1},
				m{`Name`: `2`, `V`: 2},
				m{`Name`: `3`, `V`: 3},
			},
		},

		ssCase{
			[]interface{}{
				stu{Name: `3`, V: 3, Age: 2},
				stu{Name: `1`, V: 1, Age: 1},
				stu{Name: `2`, V: 2, Age: 3},
			},
			[]interface{}{
				stu{Name: `1`, V: 1, Age: 1},
				stu{Name: `2`, V: 2, Age: 3},
				stu{Name: `3`, V: 3, Age: 2},
			},
		},
	}
	for _, testCase := range cases {
		got := testCase[0]
		Sort(got, `V`)
		expect := testCase[1]

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("expect: %v, got: %v\n", expect, got)
		}
	}
}
