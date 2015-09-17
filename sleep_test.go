package main

import "testing"
import "time"

type testpair struct {
value time.Duration
sleepSuccess bool
}

var tests = []testpair{
{ 2000, true },
{ 0, false },
{ -111, false },
}

func TestSleep(t *testing.T) {
for _, pair := range tests {
v := Sleep(pair.value)
if v != pair.sleepSuccess {
t.Error(
"For", pair.value,
"expected", pair.sleepSuccess,
"got", v,
)
}
}
}