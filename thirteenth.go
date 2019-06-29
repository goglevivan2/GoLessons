/*Указатели*/
package main
import "fmt"

func main(){
 	var x = 0
 	pointer(&x) // передаём адрес переменной
 	fmt.Println(x)
}

func pointer( x *int){ //работа с переменной по её адресу
	*x = 2
}