/*Структуры*/
package main
import "fmt"

func main(){
 bob:=Cats{"Bob",12,12.23}
 fmt.Println("Bob result is",bob.test())

 var vvod int
	fmt.Println("Введите с клавиатуры")
 	fmt.Scan(&vvod)
	fmt.Println("Ввели",vvod)
}

type Cats struct{
	name string
	age int
	happines float64
}

func (cat *Cats) test() float64{
 return float64(cat.age) * cat.happines
}