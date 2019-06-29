/*Maps - карты*/
package main
import "fmt"

func main(){
	webSites := make(map[string]float64)
	webSites["vk.com"] = 0.8
	webSites["google.com"] = 0.99
	fmt.Println(webSites["vk.com"])
}
