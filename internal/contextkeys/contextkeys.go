package contextkeys

type ctxKey string

const (
	UserID ctxKey = "user_id"
	Email  ctxKey = "email"
	Role   ctxKey = "role"
	ReqID  ctxKey = "request_id"
)
