package handler

import (
	"net/http"
    "log"
)

var routes = map[string]map[string]http.HandlerFunc{} // ["/messages":["GET":Action1]

func NewRouter(handlers *Handler) {
    Get("/messages", handlers.GetAll)
    Post("/messages", handlers.Create)
    Put("/messages/", handlers.Update)
    Delete("/messages/", handlers.Delete)

    Serve()

}

	// 1) добавляем в массив все маршруты + методы + action
	// 2) http.HandleFunc("/messages", RequestHandler(handlerFunction))
	// 3) в конце проходимся циклом по всем маршрутам и вызываем http.HandleFunc(), указывая RequestHandler('/route', list(Method=>Handler)) как прокси
	// 4) в RequestHandler проходимся по всем методам через switch method

func Get (route string, handler http.HandlerFunc) {
	SetRoute(route, handler, "GET")
}

func Post (route string, handler http.HandlerFunc) {
	SetRoute(route, handler, "POST")
}

func Put (route string, handler http.HandlerFunc) {
	SetRoute(route, handler, "PUT")
}

func Delete (route string, handler http.HandlerFunc) {
	SetRoute(route, handler, "DELETE")
}
func RequestHandler (actions map[string]http.HandlerFunc) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			actions["GET"](w,r)
		case http.MethodPost:
			actions["POST"](w,r)
		case http.MethodPut:
			actions["PUT"](w,r)
		case http.MethodDelete:
			actions["DELETE"](w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func Serve() {
	for route, actions := range routes {
	    http.HandleFunc(route, RequestHandler(actions))
    }
    log.Fatal(http.ListenAndServe(":8180", nil))
}

func SetRoute(route string, handler http.HandlerFunc, method string){
	_, exists := routes[route]
	if exists {
        routes[route][method] = handler
    } else {
    	routes[route] = map[string]http.HandlerFunc{method: handler}
    }
}