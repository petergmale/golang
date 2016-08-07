// This is a comment
// test: comment
// plus inert comment
// branch bb

package main

import (
 "fmt"
 "github.com/petergmale/mystring"
 )

func main() {
      fmt.Printf( "hello, gopher dude!\n" )
      fmt.Printf( mystring.Reverse("hello, gopher dude!\n") )
      fmt.Printf( "\n" )
}
