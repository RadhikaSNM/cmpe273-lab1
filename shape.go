package main 
import ("fmt";"math")

type Shape interface{
	area() float64
	perimeter() float64
}

type Rectangle struct{
	x1,y1,x2,y2 float64
}

type Circle struct{
	x,y,r float64
}

//Implement the methods
func (c *Circle) area() float64{
	return math.Pi*c.r*c.r
}

func (c *Circle) perimeter() float64{
//checking for -ve radius
if (c.r<0){
return -1}else{
	return 2*math.Pi*c.r}
}


func (r *Rectangle) area() float64 {
l := distance(r.x1, r.y1, r.x1, r.y2)
w := distance(r.x1, r.y1, r.x2, r.y1)
return l * w
}

func (r *Rectangle) perimeter() float64{
	l:=distance(r.x1,r.y1,r.x1,r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	//checking for 0 value : sending -1 then.
	if(l==0||w==0){
	return -1}else{
	return 2*(l+w)}
}

func distance(x1, y1, x2, y2 float64) float64{
	a:=x1-x2
	b:=y1-y2
return math.Sqrt(a*a + b*b)
}

func main() {
fmt.Println("Enter the coordinates and radius for a circle ")
var xc float64
var yc float64
var rc float64
var perimeterc float64
var newline float64

fmt.Println("x: y: z:")
_,err:=fmt.Scanf("%f %f %f",&xc,&yc,&rc)
if err != nil {
               fmt.Println("Invalid input!")
			   return
       }
fmt.Scanln(&newline)

fmt.Println("Enter the diagonal coordinates for a rectangle ")
var x1r float64
var y1r float64
var x2r float64
var y2r float64
var perimeterR float64

fmt.Println("x1: y1: x2: y2:")
_, err1 :=fmt.Scanf("%f %f %f %f",&x1r,&y1r,&x2r,&y2r)
if err1 != nil {
               fmt.Println("Invalid input!")
			   return
       }

c :=Circle{xc,yc,rc}
perimeterc=c.perimeter()

if(perimeterc==-1){
fmt.Println("The supplied radius is incorrect.(-ve)")
}else{
fmt.Println("The perimeter of the circle is: ", perimeterc)
}

r:=Rectangle{x1r,y1r,x2r,y2r}
perimeterR=r.perimeter()

if(perimeterR==-1){
fmt.Println("The supplied coordinates do not constitute a rectangle")
}else{
fmt.Println("The perimeter of the rectangle is: ", perimeterR)
}


//

/*
	c :=Circle{0,0,5}
	r:=Rectangle{0,0,10,10}
	
	fmt.Println("The perimeter of the circle is: ", c.perimeter())
	fmt.Println("The perimeter of the rectangle is: ", r.perimeter())
	*/
	
}