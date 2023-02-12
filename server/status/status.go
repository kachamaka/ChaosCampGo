package status

// Status codes for backend response
var (
	OK                  = 200
	AUTHORIZATION_ERROR = 401
	BODY_ERROR          = 400
	METHOD_ERROR        = 405
	LOGIN_ERROR         = 1
	REGISTER_ERROR      = 2
	TOKEN_ERROR         = 3
	GET_USER_ERROR      = 4
	ADD_EVENT_ERROR     = 5
	GET_EVENTS_ERROR    = 6
	DELETE_EVENT_ERROR  = 7
	ADD_REMINDER_ERROR  = 8
)
