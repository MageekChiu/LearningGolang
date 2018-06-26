package main

import "fmt"

func main()  {
	str1 := "sfdas实际"

	for _,c := range str1{
		fmt.Println(string(c))
	}
}