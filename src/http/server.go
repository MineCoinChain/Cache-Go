package http

import (
	"../cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	/**调用http包里的方法**/

	//处理缓存
	http.Handle("/cache/", s.cacheHandler())
	//获取当前数据
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
