/* Условные операторы */
package main
import "fmt"

func main()  {
	var age = 20
	if age < 5 {
		fmt.Println("Вам пора в детский сад")
	} else if age == 5{
		fmt.Println("Вам пора в  начальную школу")
	}else if (age>5) &&(age <=18){
		var grade = age - 5
		fmt.Println("Вам пора идти в",grade,"класс")
	}else{
		fmt.Println("Вам пора в универ")
	}

}
