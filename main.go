package main

import "fmt"


func add(a, b int) int{
   var answer int = a + b
   
   fmt.Println(a, "+",b, "=", answer )

   return answer
}

func subtract(a, b int)int{
  var answer int = a - b

  fmt.Println(a, "-",b, "=", answer)

  return answer
}

func multiply(a, b int)int{
  var answer int = a * b

  fmt.Println(a, "*",b, "=", answer)

  return answer
}

func divide(a, b int)int{
  var answer int = a / b

  fmt.Println(a, "/",b, "=", answer)

  return answer
}

func main() {
  var a int = 0
  var b int = 75
  
  a = multiply(b,2)
  a = add(a,9)
  a = subtract(a,3)
  a = divide(a,2)
  a = subtract(a,b)
  fmt.Println(a)
}