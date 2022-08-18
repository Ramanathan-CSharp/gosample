package main

import ( 
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	var fibresult string
	fibresult = fib(1000)
    fmt.Fprintf(w, fibresult) // installed postman can check/ any browser http://localhost:5000/
    fmt.Println("Endpoint Hit: fib") // at local terminal can see the output 
}

func fibPage(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["inputN"]
	temp, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println(err)
		  }
	var fibresult string
	fibresult = fib(temp)
    fmt.Fprintf(w, fibresult) // installed postman can check/ any browser http://localhost:5000/fibPage/2000
    fmt.Println("Endpoint Hit: fib with querystring") // at local terminal can see the output 
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/fibPage/{inputN}", fibPage)
    log.Fatal(http.ListenAndServe(":5000", myRouter))
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
