package main

import "fmt"

func main(){
	fmt.Println(cock(14, 17))
}

func cock(a,b int)(int){ //в скобках указываем получаемые данные, во вторых скобках тип возвращаемых данных
	return a+b
}