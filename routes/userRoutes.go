package routes

import (
	"net/http"

	"github.com/luisc2023/http-handler-golang/controller"
	"github.com/luisc2023/http-handler-golang/middleware"
)

func MainRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("statics"))
	mux.Handle("/statics/", http.StripPrefix("/statics/", fs))

	mux.Handle("GET /{$}", middleware.Middleware(http.HandlerFunc(controller.HomeHandler)))
	mux.Handle("GET /register", middleware.Middleware(http.HandlerFunc(controller.RegisterUsers)))
	mux.Handle("GET /login", middleware.Middleware(http.HandlerFunc(controller.LoginUsers)))
	mux.Handle("GET /update", middleware.Middleware(http.HandlerFunc(controller.UpdateUsers)))
	mux.Handle("GET /delete", middleware.Middleware(http.HandlerFunc(controller.DeleteUsers)))
	mux.Handle("GET /dashboard", middleware.Middleware(http.HandlerFunc(controller.DashboardUsers)))
	mux.Handle("GET /logout", middleware.Middleware(http.HandlerFunc(controller.LogoutUsers)))

	mux.Handle("GET /adminhome", middleware.Middleware(http.HandlerFunc(controller.AdminHome)))
	mux.Handle("GET /adminregister", middleware.Middleware(http.HandlerFunc(controller.AdminRegister)))
	mux.Handle("GET /adminlogin", middleware.Middleware(http.HandlerFunc(controller.AdminLogin)))
	mux.Handle("GET /adminupdate", middleware.Middleware(http.HandlerFunc(controller.AdminUpdate)))
	mux.Handle("GET /admindelete", middleware.Middleware(http.HandlerFunc(controller.AdminDelete)))

	mux.Handle("GET /createproducts", middleware.Middleware(http.HandlerFunc(controller.CreateProducts)))
	mux.Handle("GET /getproduct", middleware.Middleware(http.HandlerFunc(controller.GetProduct)))
	mux.Handle("GET /listproducts", middleware.Middleware(http.HandlerFunc(controller.ListProducts)))
	mux.Handle("GET /updateproduct", middleware.Middleware(http.HandlerFunc(controller.UpdateProduct)))
	mux.Handle("GET /deleteproduct", middleware.Middleware(http.HandlerFunc(controller.DeleteProduct)))

	mux.Handle("GET /createorders", middleware.Middleware(http.HandlerFunc(controller.CreateOrders)))
	mux.Handle("GET /getorder", middleware.Middleware(http.HandlerFunc(controller.GetOrder)))
	mux.Handle("GET /listorders", middleware.Middleware(http.HandlerFunc(controller.ListOrders)))
	mux.Handle("GET /updateorder", middleware.Middleware(http.HandlerFunc(controller.UpdateOrder)))
	mux.Handle("GET /deleteorder", middleware.Middleware(http.HandlerFunc(controller.DeleteOrder)))

	return mux
}
