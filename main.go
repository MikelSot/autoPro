package main

import (
	"fmt"
	"strconv"
	"time"
)

//import (
//	"net/http"
//	"github.com/labstack/echo/v4"
//)

func main() {
	//database.Migration()
	//busy := make(map[string]string)
	//
	//busy["5"] = "5"
	//busy["10"] = "10"
	//busy["15"] = "15"
	//busy["1"] = "1"
	//busy["18"] = "18"
	//busy["13"] = "13"
	//busy["14"] = "14"
	//busy["4"] = "4"
	//
	// validos := signValid(busy)
	//	fmt.Println()
	//	fmt.Println(validos)
	//for _, v := range validos {
	//	fmt.Println(v)
	//}

	now := time.Now().Format("2006-01-02")
	fmt.Println(now)
}

//
// signValid firltra el N° que estan disponibles para la cita
func signValid(busy map[string]string)  map[int]string{
	numbers := make(map[int]string)
	for i := 1; i <= 20; i++ {
		numbers[i] = strconv.Itoa(i)
	}

	for _, v := range busy {
		for key, _ := range numbers {
			if numbers[key] == v{
				delete(numbers, key)
				delete(busy, v)
			}
		}
	}
	return numbers
}

