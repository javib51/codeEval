package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

func countCharacters(line []byte) map[byte]int{
	counts := make(map[byte]int)
	for i:= 0; i < len(line); i++ {
		counts[line[i]]++
	}
	return counts
}

func firstNonRepeatedCharacter(line []byte){
	var element byte
	var counts map[byte]int = countCharacters(line)

	for i:= 0; i < len(line); i++ {
		if counts[line[i]] == 1 {
			element = line[i]
			break
		}		
	}
	
	fmt.Println(string(element))
}

func main () {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}   
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
    
	for scanner.Scan(){
		firstNonRepeatedCharacter(scanner.Bytes())
	}
}

