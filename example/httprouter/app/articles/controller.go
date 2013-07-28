package articles

import (
	"../../../../util"
	"../../libs"
	"net/http"
	"strconv"
)

func TopPage(params map[string]string, w http.ResponseWriter, r *http.Request) (message string, err util.ActionError) {
	//fmt.Fprintf(w, "Accept, %q\n", html.EscapeString(r.URL.Path))
	data := make(map[string]string)
	data["title"] = "タイトル"
	data["message"] = "Hello, world."
	view := libs.Viewer("templates/articles/top.html")
	if view != nil {
		view(w, data)
	} else {
		err = util.NewActionError("404", "View Not Found")
	}
	message = "TopPage:" + r.URL.Path
	return
}

func ArticlePage(params map[string]string, w http.ResponseWriter, r *http.Request) (message string, err util.ActionError) {
	message = "TopPage:" + r.URL.Path
	data := make(map[string]interface{})
	data["title"] = "タイトル"
	_, ok := params["id"]
	if ok {
		id, e := strconv.ParseInt(params["id"], 10, 0)
		if e != nil {
			err = util.NewActionError("404", "id mismatch")
			return
		}
		data["id"] = id
		article, e := Get(id)
		if e != nil {
			err = util.NewActionError("404", "article not found")
			return
		}
		data["article"] = article
	} else {
		err = util.NewActionError("404", "article not found")
		return
	}

	view := libs.Viewer("templates/articles/article.html")
	if view != nil {
		view(w, data)
	} else {
		err = util.NewActionError("404", "view not found")
	}
	return
}
