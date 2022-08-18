package main

import ( 
	"fmt"
	"strconv"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	var fibresult string
	fibresult = fib(1000)
    fmt.Fprintf(w, fibresult) // installed postman can check/ any browser http://localhost:5000/
    fmt.Println("Endpoint Hit: fib") // at local terminal can see the output 
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func fib(n int) string {
	var a, b int
	var result string
 	a, b = 0, 1
	for a < n {
		temp := strconv.Itoa(a)
		result = result + " " + temp
		 a, b = b, a+b
	}
	return result
}


func main() {
	handleRequests()
}
