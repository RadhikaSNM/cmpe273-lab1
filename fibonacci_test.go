package main

import "testing"

type testpair struct {
input int64
fibonacci int64
}

var tests = []testpair{
{ -131, -1 },
{ 0, 0 },
{ 1, 1 },
{ 6, 8 },
{ 20, 6765 },
}

func TestFibonacci(t *testing.T) {
for _, pair := range tests {
v := Fibonacci(pair.input)
if v != pair.fibonacci {
t.Error(
"For", pair.input,
"expected", pair.fibonacci,
"got", v,
)
}
}
}