package routers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gopromo/app/http/routers/middleware"
)

var midd = middleware.Midd{}

type routeType struct {
	prefix  string
	handler map[string]http.HandlerFunc
}

type routeTypes struct {
	prefix   string
	handlers []routeType
}

func Run() {
	http.HandleFunc("/", indexHandler)

	reg := reg{}
	reg.register([]routeTypes{
		common(), //加入普通活动
	})
	log.Println(http.ListenAndServe(":8000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

type reg struct {
}

func (this *reg) addPrefix(prefix string, route map[string]http.HandlerFunc) map[string]http.HandlerFunc {
	var result = make(map[string]http.HandlerFunc)
	for k, v := range route {
		result[strings.Join([]string{prefix, "/", k}, "")] = v
	}
	return result
}

func (this *reg) addRoute(route map[string]http.HandlerFunc) {
	for k, v := range route {
		http.HandleFunc("/"+k, v)
	}
}

func (this *reg) addFunc(prefix string, route []routeType) {
	for _, v := range route {
		this.addRoute(this.addPrefix(prefix, this.addPrefix(v.prefix, v.handler)))
	}
}

func (this *reg) register(routes []routeTypes) {
	for _, v := range routes {
		this.addFunc(v.prefix, v.handlers)
	}
}
