package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_DataProcessor(t *testing.T) {
	data := []struct {
		name     string
		message  []byte
		expected int
	}{
		{"addition", []byte(`1
+
3
4`), 7},
		{"multiplication", []byte(`1
*
3
4`), 12},
		{"substraction", []byte(`1
-
3
4`), -1},
		{"division", []byte(`1
/
4
2`), 2},
		{"division by zero", []byte(`1
/
4
0`), 0},
	}
	in := make(chan []byte, 1)
	out := make(chan Result, 1)
	go DataProcessor(in, out)
	for _, d := range data {
		in <- d.message
		select {
		case res := <-out:
			if res.Value != d.expected {
				t.Errorf("%v test failed, got %v expected %v", d.message, res.Value, d.expected)
			}
			if res.Id != "1" {
				t.Errorf("%v test failed, got %v expected %v", d.message, res.Id, "1")
			}
		}
	}
}

func Test_parser(t *testing.T) {
	data := []struct {
		name     string
		message  []byte
		expected Input
		err      error
	}{
		{"addition", []byte(`1
+
3
4`), Input{Id: "1", Op: "+", Val1: 3, Val2: 4}, nil},
		{"multiplication", []byte(`1
*
3
4`), Input{Id: "1", Op: "*", Val1: 3, Val2: 4}, nil},
		{"substraction", []byte(`1
-
3
4`), Input{Id: "1", Op: "-", Val1: 3, Val2: 4}, nil},
		{"division", []byte(`1
/
4
2`), Input{Id: "1", Op: "/", Val1: 4, Val2: 2}, nil},
		{"bad val1", []byte(`1
/
a
2`), Input{}, nil},
		{"bad val2", []byte(`1
/
3
b`), Input{}, nil},
	}
	for _, d := range data {
		res, err := parser(d.message)
		if err != nil && d.err != nil {
			t.Errorf("%v test failed, got %v expected %v", d.name, err, d.err)
		}
		if res != d.expected {
			t.Errorf("%v test failed, got %v expected %v", d.name, res, d.expected)
		}
	}
}

func Fuzz_parser(f *testing.F) {
	testcases := [][]byte{
		[]byte(`1
+
3
4`),
		[]byte(`1
*
3
4`),
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		out, err := parser(in)
		if err != nil {
			t.Skip("handled error")
		}
		roundTrip := []byte(out.Id + "\n" + out.Op + "\n" + fmt.Sprint(out.Val1) + "\n" + fmt.Sprint(out.Val2))
		out2, err := parser(roundTrip)
		if diff := cmp.Diff(out, out2); diff != "" {
			t.Error(diff)
		}
		if err != nil {
			t.Skip("handled error")
		}
	})
}
