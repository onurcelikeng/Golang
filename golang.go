package main

import (
	"fmt"
)

func main(){
    loop()
}

func helloWorld(){
    fmt.Println("Hello " + "Go!") 
    fmt.Println("Hello", "Go!")
}

func values(){
    fmt.Println("result :", 4 * 3)
    fmt.Println((true && true), (true || false), (!true))
}

func variables(){
    var str =  "Onur"
    fmt.Println(str)

    var data1, data2 int = 12, 13 //declare multiple variables
    fmt.Println(data1 + data2)

    var value = true //infer the type of init variables.
    fmt.Println(value)

    number := 12 //declaring and initializing
    fmt.Println(number)
}

func loop(){
    
    //classic
    for i := 0; i < 4; i++{
        fmt.Print(i)
    }

    //sing cond.
    i := 0
    for i < 3{
        fmt.Print(i)
        i++ 
    }

    //endless (with break)
    for{
        fmt.Print("onur")
        break
    }
}

func condition(){
    
}

func bubbleSort(){

    var array [10]int =[10]int {4,1,6,3,4,8,4,5,6,3}

    isSwap := true

    for isSwap{
        isSwap = false
        for i := 1; i < len(array); i++{
            if array[i-1] > array[i]{
                array[i], array[i-1] = array[i-1], array[i] //using tupple for swapping
                isSwap = true
            }
        }
    }

    fmt.Println(array)
}

func maps(){
    m := make(map[string]int)
    m["onur"] = 23
    m["seray"] = 22

    fmt.Println("map : ", m)
}