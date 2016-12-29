package main

import (
	"fmt"
)

func main(){
    values()
}

func helloWorld(){
    fmt.Println("Hello " + "Go!")
    fmt.Println("Hello", "Go!")
}

func values(){
    fmt.Println("result :", 4 * 3)
    fmt.Println((true && true), (true || false), (!true))
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

func Map(){
    m := make(map[string]int)
    m["onur"] = 23
    m["seray"] = 22

    fmt.Println("map : ", m)
}