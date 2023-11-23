package http

type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded",
	}
	Created = Status{
		Code:        201,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created",
	}
	NoContent = Status{
		Code:        204,
		Status:      "NO_CONTENT",
		Description: "There is no content to send for this request, but the headers may be useful",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "Invalid argument value passed",
	}
	Unauthorized = Status{
		Code:        401,
		Status:      "UNAUTHORIZED",
		Description: "...",
	}
	Forbidden = Status{
		Code:        403,
		Status:      "FORBIDDEN",
		Description: "...",
	}
	InternalServerError = Status{
		Code:        500,
		Status:      "INTERNAL_SERVER_ERROR",
		Description: "The server encountered an unexpected condition that prevented it from fulfilling the request",
	}
	InvalidAuth = Status{
		Code:        400,
		Status:      "INVALID_AUTH",
		Description: "Authentication/OAuth token is invalid",
	}
	InvalidAuthHeader = Status{
		Code:        400,
		Status:      "INVALID_AUTH_HEADER",
		Description: "Authentication header is invalid",
	}
	AccessDenied = Status{
		Code:        401,
		Status:      "ACCESS_DENIED",
		Description: "Authentication unsuccessful",
	}
	NotFound = Status{
		Code:        404,
		Status:      "NOT_FOUND",
		Description: "Invalid URL",
	}
	RequestConflict = Status{
		Code:        409,
		Status:      "REQUEST_CONFLICT",
		Description: "Requested operation resulted in conflict",
	}
)

// 400	INVALID_BODY	A request body that was not in JSON format was passed.
// 400	UNSUPPORTED_OPERATION	Requested operation not supported.
// 401	ACCESS_DENIED	Authentication unsuccessful.
// 401	NO_AUTH	Authentication not provided.
// 403	NOT_AUTHORIZED	User has not been authorized to perform that action.
// 404	NOT_FOUND	Invalid URL.
// 405	METHOD_NOT_ALLOWED	Method is not allowed for this endpoint.
// 409	REQUEST_CONFLICT	Requested operation resulted in conflict.
// 429	HIT_RATE_LIMIT	Hourly rate limit has been reached for this token. Default rate limits are 2,000 calls per hour.
// 500	EXPANSION_FAILED	Unhandled error occurred during expansion; the request is likely to succeed if you don’t ask for expansions, but contact Eventbrite support if this problem persists.
// 500	INTERNAL_ERROR	Unhandled error occurred in Eventbrite. contact Eventbrite support if this problem persists.
