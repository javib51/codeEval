package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func bubbleSort(data []int, cycles int) {

	var n int = len(data)
	
	for i := 0 ; i < cycles && i < n; i++ {
		for j := 0; j < n-1; j++ {
			if data[j] > data[j + 1] {
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}	
}

func parser(str string){
	v := strings.FieldsFunc(str, func(r rune) bool {
		switch r {
		case '|', ' ':
			return true
		}
		return false
	})
	var data = []int{}
	for _, i := range v {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		data = append(data,j)
	}

	bubbleSort(data[:len(data)-1],data[len(data)-1])
	
	//fmt.Println(res[:len(res)-1])
	//var res string = ""
	for i := 0; i < len(data)-1; i++ {
		
		if i == len(data)-2 {
			fmt.Print(data[i],"\n")
		} else {
			fmt.Print(data[i]," ")
		} 
	}
	
	//fmt.Println(res)
}

func main () {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}   
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
    
	for scanner.Scan(){
		parser(scanner.Text())
	}
}
