package main

import "fmt"
import "strconv"

//import "./Dick" //импорт локальной либы

func main(){
	num,err := strconv.Atoi("10")  // конверт типа данных (в примере из строки в число), два возвращаемых значения ретурн и ерр, где ерр несет в себе инфу об ошибке, если такова будет, ексцепшенов в го нет
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(num)
}

//a := []string{"kok","world"} // создание списка со строками
//a := make([]string, 5)  //альтернатив способ объявления списка (в примере с 5 пустыми строками)

//for i:=0; i<len(a); i++ { //цикл хуле
//	fmt.Println(a[i])
//}

//for k,v := range a{ //в связке с range получаеться этакий foreach
//	fmt.Println(k)
//	fmt.Println(v)
//}

//fmt.Println(Dick(14, 17))
//func Dick(a,b int)(int){ //в скобках указываем получаемые данные, во вторых скобках тип возвращаемых данных
//	return a+b
//}

//Dick.A = 15  //юзаем переменную из другой либы
//fmt.Println(Dick.A)