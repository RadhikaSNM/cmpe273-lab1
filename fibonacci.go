package main

import ("fmt")

func Fibonacci(num int64) int64 {
	//Using uint64 to accomadate largest unsigned number: 
	//Fibonacci series is used only for unsigned integers.
	
	//checking -ve case:
	if(num <0){
	return -1
	} else if(num <=1){
	return num
	}else{
		return ( Fibonacci(num-1)+Fibonacci(num-2))}
}

func main() {
fmt.Println("Enter the INTEGER to calculate the fibonacci series")
var input int64
var output int64
_,err:=fmt.Scanf("%d",&input)
//checking if integer
if err != nil {
               fmt.Println("Invalid input!")
			   return
       }
output=Fibonacci(input)
if (output!=-1){
fmt.Println("The calculated fibonacci value is: ",output)
}else{
fmt.Println("You have supplied a negative number. Fibonacci value cannot be calculated")
}
}