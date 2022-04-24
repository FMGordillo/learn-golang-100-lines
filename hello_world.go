// Required for a standalone executable
package main

// implements formatted I/O
import "fmt"

// 1. Variables must be declare before being used
// NOTE: We can use 'const' instead
var a int
// Group declaration
var (
  b bool
  c float32
  d string
)

func main() {
  // 1. Variables
  a = 42
  b, c = true, 32.0
  d = "string"
  fmt.Println(a, b, c, d)

  // 2. We can declare and assign in the same line
  // "short variable declaration"
  a1:= 42
  b1, c1 := true, 32.0
  d1 := "string"
  fmt.Println(a1, b1, c1, d1)
}
