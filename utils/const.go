package utils

const Prefix = "/v1"

// Response

type UrlMapping struct {
	Url    string
	Method string
	Role   string
}

var ProtectUrl = []UrlMapping{
	{
		Url:    Prefix + "/user",
		Method: "GET",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/user",
		Method: "POST",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/user/{id}",
		Method: "GET",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/user/{id}",
		Method: "PUT",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/user/{id}",
		Method: "DELETE",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/category",
		Method: "POST",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/category",
		Method: "PUT",
		Role:   "admin",
	},
	{
		Url:    Prefix + "/category",
		Method: "DELETE",
		Role:   "admin",
	},
}
