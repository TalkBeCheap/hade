package framework

import (
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{}
}
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
