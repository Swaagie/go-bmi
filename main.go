package main

import "os"
import "log"
import "fmt" 
import "math"

/*
 * Prompt user and calculate BMI.
 *
 * @public
 */
func main() {
  fmt.Println("\nCalculate your BMI\n")

  weight := get("weight (kg)")
  length := get("length (cm)")

  fmt.Printf("\nYour BMI: %.2f kg/m² \n\n", weight / math.Pow(length/100, 2));
}

/*
 * Prompt the user for their weight or length.
 *
 * @param {string} s what variable should the user be prompted with
 * @private
 */
func get(s string) (float64) {
  var x float64

  fmt.Print("What is your "+ s +"? ")

  _, err := fmt.Scanf("%f", &x)
  check(err)

  return x
}

/*
 * Log and exit if an error is encountered.
 *
 * @param {error} err
 * @private
 */
func check(err error) {
  if err != nil {
    log.Fatal(err)
    os.Exit(0)
  }
}
