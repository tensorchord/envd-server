package server

const (
	ContextLoginName = "login-name"
)

type AuthMiddlewareHeaderRequest struct {
	JWTToken string `header:"Authorization"`
}

type AuthMiddlewareURIRequest struct {
	LoginName string `uri:"login_name" example:"alice"`
}
