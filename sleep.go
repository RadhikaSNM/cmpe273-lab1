package main

import ("fmt";"time")

func Sleep (t time.Duration) bool{
	fmt.Println("going to sleep for ",t)
	
	//Checking for 0 or -ve or overflow
	if t<=0{
		fmt.Println("provided time is 0 or -ve value.Returning immediately")
		return false
	}
	
	select {
		//case timeValue:= <- time.After(time.Millisecond*time.Duration(t)):
		case  <- time.After(time.Millisecond*t):
			//fmt.Println(timeValue)
			return true		
	 }
	
}

func main() {
fmt.Println("Enter the time to sleep in millisec")
var input int64
fmt.Scanf("%d",&input)
fmt.Println("the time is: ",time.Now())
Sleep(time.Duration(input))
fmt.Println("the time after the sleep is: ",time.Now())
}