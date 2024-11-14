package middlewares

//	type Authorizer interface {
//		BasicAuth(ctx context.Context, username, password string) (*models.User, error)
//	}
type Logger interface {
	Error(err error, msg string)
	SendWithFields(fields map[string]interface{})
}
