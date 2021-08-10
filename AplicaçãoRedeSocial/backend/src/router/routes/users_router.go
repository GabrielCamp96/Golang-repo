package router

import (
	"api/src/controller"
	"net/http"
)

var userRoutes = []route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controller.CreateUser,
		isAuth:   false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controller.GetUsers,
		isAuth:   false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controller.GetUserById,
		isAuth:   false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controller.UpdateUser,
		isAuth:   false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controller.DeleteUser,
		isAuth:   false,
	},
}
