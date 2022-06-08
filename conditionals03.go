/* RZFeeser | Alta3 Research
   if, else if, and else statements */

package main

import "fmt"
import "math/rand"
import "time"

func main() {
 
    // ensure we produce random values
    rand.Seed(time.Now().UnixNano())
    
              // snippet to choose a random number from a range" 
              // rand.Intn(max - min) + min
              // note growth will only be available in the if then else construct
    if growth := rand.Intn(21 - -100) + -100; growth < 0 {
        fmt.Println("Looks like growth was negative:", growth)
    } else if growth < 10 {
        fmt.Println("Growth was a bit flat:", growth)
    } else {
        fmt.Println("Double digit growth:", growth)
    }
    // note: growth is not available here 
   //  fmt.Println("Here is growth", growth)
} 

