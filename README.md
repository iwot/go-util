go-util
=======

util for golang (for me)  
StringBuffer only.

INSTALL
-------
`go get github.com/iwot/go-util`

USAGE
-----
    import (
      "github.com/iwot/go-util/util"
    )
    
    func f() {
      sb := util.NewStringBuffer()
      sb.Append("A")
      sb.Append("B")

      var str string = sb.String()
    }
