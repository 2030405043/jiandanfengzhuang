package route

import (
	"github.com/drone/routes"
	"message/controller"
	"net/http"
)

func init() {
	routeMux := routes.New()
	register(routeMux)
	http.Handle("/",routeMux)
}
func register(r *routes.RouteMux) {

	r.Post("/message", controller.AddMessage)
}
