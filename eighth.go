/*Массивы*/
package main
import "fmt"
func main(){
/*	var arr[3] int
	arr[0] = 45
	arr[1] = 97
	arr[2] = 76
	fmt.Println(arr[2]+arr[0])*/
//	Альтернативнйы вариант задания массива
	num:=[3]float64{3.123 , 4.123 , 8.132}
	for i,value := range num{
		fmt.Println(value,i)
	}
}

