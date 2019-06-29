/* printf */
package main
import "fmt"

func main(){
	var num float64 = 4.3254357
	fmt.Printf("%f \n",num) // всё число
	fmt.Printf("%.2f \n",num)// две цифры после запятой
	fmt.Printf("%T \n",num)// тип переменной

	var isDone bool = true
	fmt.Printf("%t \n",isDone)// просто значение переменной
}
