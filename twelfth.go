/*Откладывание*/
package main
import "fmt"

func main(){
	defer two() // Откладывание
	one()
}

func one(){
	fmt.Println("1")
}

func two(){
	fmt.Println("2")
}