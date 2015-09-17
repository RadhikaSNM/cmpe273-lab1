package main

import "testing"

type testpair struct {
x float64
y float64
r float64
perimeterC float64
}

var tests = []testpair{
{ -1,-1,-5,-1 },
{ 0, 0,0,0 },
{ 1, 1,1, 6.283185307179586 },
{ 6, 8 , 24, 150.79644737231007},
}

//for rectangle
type testpair1 struct {
x1 float64
y1 float64
x2 float64
y2 float64
perimeterR float64
}

var testsRectanglePerimeter = []testpair1{
{ -1,-2,-3,-4,8},
{ 0, 0,0,0,-1 },
{ 1, 2,1,9,-1 },
{ 2, 6,5,6,-1 },
{ 4, 6, 10,12,24 },
{ 24, 2, 4, 30, 96},
}


func TestPerimeter(t *testing.T) {

//Testing for circle
for _, pair := range tests {
c := Circle{pair.x,pair.y,pair.r}
v := c.perimeter()
if v != pair.perimeterC {
t.Error(
"For", pair.x, pair.y,pair.r,
"expected", pair.perimeterC,
"got", v,
)
}
}


//Testing for rectangle
for _, pair1 := range testsRectanglePerimeter {
r := Rectangle{pair1.x1,pair1.y1,pair1.x2,pair1.y2}
v := r.perimeter()
if v != pair1.perimeterR {
t.Error(
"For", pair1.x1,pair1.y1,pair1.x2,pair1.y2,
"expected", pair1.perimeterR,
"got", v,
)
}
}



}