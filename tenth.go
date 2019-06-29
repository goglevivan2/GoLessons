/*Функции*/
package main
import "fmt"

func main(){
	 var a =2
	 var b = 3
	 a,b = summ(a,b)
	 fmt.Println(a,b)
}

func summ(num_1 int,num_2 int) (int,int){
	var res int
	res = num_1 + num_2
	return res,num_1*num_2
}