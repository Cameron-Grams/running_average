package main

import(
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main(){
        var holder []float64
	var value string = "Yes"
	index := 0

	fmt.Println("Enter numbers to average.  Enter No to stop. ")
        reader := bufio.NewReader(os.Stdin)

	out:
	for value != "No" {
		fmt.Println("Enter value: ")
		value, err := reader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
		if value == "No\n"{
			break out
		}

		value = strings.TrimSpace(value)
		num_value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}
		holder = append(holder, num_value)
		index++
        }	
	fmt.Println("Leaving Entering Numbers")
        fmt.Printf("Numbers: %v\n", holder)
}

