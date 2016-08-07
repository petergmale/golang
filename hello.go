// This is a comment
// test: comment
// plus inert comment
// master with branch bb and bc

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
