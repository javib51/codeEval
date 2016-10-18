package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"regexp"
)

func validate ( str string) {

	var validID = regexp.MustCompile(`\w+@\w+.\w+`)
    
	fmt.Println(validID.MatchString(str))

}

func main () {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}   
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
    
	for scanner.Scan(){
		validate(scanner.Text())
	}
}

