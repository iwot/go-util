go-util
=======

util for golang.  

- StringBuffer.
- Stack.

INSTALL
-------
`go get github.com/iwot/go-util`

USAGE
-----
```go
import (
  "github.com/iwot/go-util/util"
)

func f() {
  sb := util.NewStringBuffer()
  sb.Append("A")
  sb.Append("B")

  var str string = sb.String() // "AB"
}

func t() {
  sb := util.NewStringBuffer("AA", "BB")
  sb.Append("A")
  sb.Append("B")

  var str string = sb.String() // "AABBAB"
}
```
