go-util
=======

util for golang.  

- PriorityQueue (Dijkstra Sample : priorityqueue_test.go)
- PathMatch
- StringBuffer.
- Stack.

INSTALL
-------
`go get github.com/iwot/go-util`

USAGE
-----
### HttpRouter
簡単なルーター。
PathMatchを使用してルーティングします。
使用例はexample/httprouter/main.go
```go
package main

import (
  "./app/articles"
  "fmt"
  "github.com/iwot/go-util/util"
  "html"
  "net/http"
  "os"
)

func main() {

	basePath := "/articles/"
	router := util.NewHttpRouter().SetBasePath(basePath)
	router.SetRoute("", articles.TopPage)
	router.SetRoute("article/<id>", articles.ArticlePage)

	// for default error
	router.SetError("default", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Error, %q\n", html.EscapeString(message))
	})
	// for 404 error
	router.SetError("404", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "404 Error, %q\n", html.EscapeString(message))
	})
	// for other error
	router.SetError("other", func(message string, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "other Error, %q\n", html.EscapeString(message))
	})

	http.Handle(basePath, router)
	fmt.Fprint(os.Stdout, "http://localhost:8000/articles/\n")
	fmt.Fprint(os.Stdout, "http://localhost:8000/articles/article/1\n")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
```

### PriorityQueue
プライオリティ付きのキュー。
ダイクストラのサンプルをテストにつけてます。
```go
q := NewPriorityQueue()
assert.Equal(t, 0, q.Size())

q.Push(&N1{1})
v2 := &N1{4}
q.Push(v2)
v3 := &N1{3}
q.Push(v3)
assert.Equal(t, 3, q.Size())
p := q.Pop().(*N1)
assert.Equal(t, 2, q.Size())
assert.Equal(t, 4, p.value)

assert.True(t, q.Contain(v3))
assert.False(t, q.Contain(v2))
```

### PathMatch
URLのパスのパターンに対して、与えられたパスがマッチするかを判定し、マッチする場合はそのマッチ結果を返す。
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

### StringBuffer
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

### Stack
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
