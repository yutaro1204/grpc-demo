package main

import "fmt"

// echo: https://echo.labstack.com/guide

type A struct {
	i int
}

type Data struct {
	number int
}

var onMemoryData *Data

func main() {
	onMemoryData = &Data{
		number: 100,
	}
	var newOnMemoryData Data = *onMemoryData

	onMemoryData.number = 200 // this doesn't affect to newOnMemoryData

	fmt.Printf("onMemoryData value: %p\n", onMemoryData) // the address of the actual value
	fmt.Println(*onMemoryData)                           // the actual value

	fmt.Printf("onMemoryData address: %p\n", &onMemoryData) // the address of the pointer for the actual value

	fmt.Println("##########")

	fmt.Println(newOnMemoryData)                               // the address of the actual value
	fmt.Printf("onMemoryData address: %p\n", &newOnMemoryData) // the address of the pointer for the actual value

	// list := []A{}
	// list = append(list, A{i: 1})
	// list = append(list, A{i: 2})
	// list = append(list, A{i: 3})

	// pList := make([]*A, 0, len(list))
	// for i := range list {
	// 	newOne := list[i]              // recreate item with new address
	// 	pList = append(pList, &newOne) // the address recreated
	// }

	// for _, v := range pList {
	// 	fmt.Println(v.i)
	// 	fmt.Println(&v.i)
	// }

	// // list will be changed, but pList won't be changed
	// list[0].i = 5
	// fmt.Println("Again for list")
	// for _, v := range list {
	// 	fmt.Println(v.i)
	// 	fmt.Println(&v.i)
	// }

	// fmt.Println("Again for pList")
	// for _, v := range pList {
	// 	fmt.Println(v.i)
	// 	fmt.Println(&v.i)
	// }

	// number1 := new(int)
	// fmt.Printf("a2 value: %p\n", number1)
	// fmt.Printf("a2 address: %p\n", &number1)

	// number2 := number1
	// fmt.Printf("a2 value: %p\n", number2)
	// fmt.Printf("a2 address: %p\n", &number2)

	// e := echo.New()
	// echoWebSite := "https://echo.labstack.com/"
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, fmt.Sprintf("Started echo service, bring it on! See %s", echoWebSite))
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
