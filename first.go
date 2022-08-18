/*
package main
import "fmt"
func main() {
	var a, b, n int 
	a, b = 0, 1
	n = 1000

	for a < n {
		fmt.Printf( "%d ", a )
		 a, b = b, a+b
	}
	 fmt.Println()
}
*/
/*
package main
import "fmt"
func main() {
fib(1000)
}
func fib(n int) {
	var a, b int
 	a, b = 0, 1
	n = 1000

	for a < n {
		fmt.Printf( "%d ", a )
		 a, b = b, a+b
	}
	 fmt.Println()
}
*/

/*
package main
import "fmt"

func main() {
	var inputN int
	fmt.Printf("Enter N integer: ")
	fmt.Scanf("%d", &inputN)
	fib(inputN)
}
func fib(n int) {
	var a, b int
 	a, b = 0, 1
	for a < n {
		fmt.Printf( "%d ", a )
		 a, b = b, a+b
	}
	 fmt.Println()
}
*/


package main
import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	var inputN int
	
	if len(os.Args) >= 2 {
		var temp int
		temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		  }
		  inputN = temp
	} else {
		var temp int
		fmt.Printf("%d",len(os.Args))
		fmt.Printf("Enter N integer: ")
		fmt.Scanf("%d", &temp)
		inputN = temp
	}
fib(inputN)
}
func fib(n int) {
	var a, b int
 	a, b = 0, 1
	for a < n {
		fmt.Printf( "%d ", a )
		 a, b = b, a+b
	}
	fmt.Println()
}
