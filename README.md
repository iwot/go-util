go-util
=======

util for golang.  

- PathMatch
- StringBuffer.
- Stack.

INSTALL
-------
`go get github.com/iwot/go-util`

USAGE
-----
PathMatch
```go
m := NewPathMatch()
pattern := `/member/<id:[0-9]{8}>/<page>`
defaults := make(map[string]string)
err = m.Parse(pattern, defaults)
assert.False(t, ok)

path = `/member/01234567/news`
pathPattern, matches, ok = m.Match(path)

assert.True(t, ok)
assert.Equal(t, pattern, pathPattern)
assert.Equal(t, "01234567", matches["id"])
assert.Equal(t, "news", matches["page"])
```

StringBuffer
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

Stack
```go
stack := new(Stack)
stack.Push("aaa")
stack.Push("bbb")
value, size := stack.Pop()

for elm := range stack.Iter() {
// elm.value
}
```

[MIT License](https://github.com/iwot/go-util/blob/master/LICENSE "MIT License")
