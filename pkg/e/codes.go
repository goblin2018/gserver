package e

var (
	OK            = newErr(200, "OK")
	InvalidParams = newErr(400, "Invalid Params")
	Unauthorized  = newErr(401, "Unauthorized")
	Forbidden     = newErr(403, "Forbidden")

	NotImplemented = newErr(501, "Method not implemented.")
	ServerError    = newErr(500, "Internal server error")

	InvalidToken = newErr(1001, "InvalidToken")

	DBError      = newErr(1002, "DB Error")
	ErrorCaptcha = newErr(1003, "Error Captcha")
)
