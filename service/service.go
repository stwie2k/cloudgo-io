package service

import (
	"github.com/mux"
	"github.com/negroni"
	"github.com/render"
	"log"
	"net/http"
	"os"
	"text/template"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	// 返回一个Render实例的指针
	formatter := render.New(render.Options{
		IndentJSON: true,

	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	//mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")

	mx.HandleFunc("/", indexHandlerFunc).Methods("GET")
	mx.HandleFunc("/", Login).Methods("POST")

	mx.HandleFunc("/json", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/unknown", NotImplement)


	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/static/"))))

}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
			formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}
func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	/*使用template.ParseFiles()实现模板的渲染输出
	 *文件路径的根目录以可执行文件为基准*/
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)

}
func NotImplement(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "501 not implemented.", 501)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["username"][0]
	pass := r.Form["password"][0]
	//请求的是登录数据，那么执行登录的逻辑判断
	if len(name) == 0 || len(pass) == 0 {
		log.Fatal("err")
		return
	}
	t := template.Must(template.ParseFiles("templates/table.html"))
	t.Execute(w, map[string]string{
		"Name": name,
		"Pass": pass,
	})
}
