/* switch
В языке Go в switch - case конструкциях не набо прописывать break
*/
package main
import "fmt"

func main() {
	var age =10
	switch age {
	case 5: fmt.Println("Вам пять лет")
	case 15: fmt.Println("Вам 15 лет")
	case 10: fmt.Println("Вам 10 лет")
	default:fmt.Println("Я не знаю сколко вам лет")

	}
}
