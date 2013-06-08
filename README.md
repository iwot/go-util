go-util
=======

util for golang.  
StringBuffer only.

Goではわざわざbytes.Bufferのラッパーは作るほどでもないですが、Githubへの登録、および`go get`から自分の作ったライブラリのインストールを試す意味でも作ってみました。

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
