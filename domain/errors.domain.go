package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("NOT_FOUND")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

	ErrAlreadyUsed         = errors.New("ALREADY_USED")
	ErrCreate              = errors.New("CREATE_ERROR")
	ErrDecodeImage         = errors.New("DECODE_IMAGE_ERROR")
	ErrDelete              = errors.New("DELETE_ERROR")
	ErrEmailExists         = errors.New("EMAIL_EXISTS")
	ErrEmailNotVerified    = errors.New("EMAIL_NOT_VERIFIED")
	ErrEmailRequired       = errors.New("EMAIL_REQUIRED")
	ErrEncodeImage         = errors.New("ENCODE_IMAGE_ERROR")
	ErrExpired             = errors.New("EXPIRED")
	ErrGqlApiKeyInvalid    = errors.New("GraphQL API Key is wrong")
	ErrInsert              = errors.New("INSERT_ERROR")
	ErrInvalidContentType  = errors.New("INVALID_CONTENT_TYPE")
	ErrUploadingFile       = errors.New("ERROR_UPLOADING_FILE")
	ErrInvalidInterest     = errors.New("INVALID_INTEREST")
	ErrInvalidRole         = errors.New("INVALID_ROLE")
	ErrJWTInvalid          = errors.New("JWT_INVALID")
	ErrList                = errors.New("LIST_ERROR")
	ErrNotLogged           = errors.New("NOT_LOGGED")
	ErrPasswordNotSecure   = errors.New("PASSWORD_NOT_SECURE")
	ErrPasswordNotSet      = errors.New("PASSWORD_NOT_SET")
	ErrRoleInvalid         = errors.New("ROLE_INVALID")
	ErrServerError         = errors.New("SERVER_ERROR")
	ErrSignInError         = errors.New("SIGNIN_ERROR")
	ErrSourceInvalid       = errors.New("SOURCE_INVALID")
	ErrSourceNotRegistered = errors.New("SOURCE_NOT_REGISTERED")
	ErrUpdate              = errors.New("UPDATE_ERROR")
	ErrSetInterest         = errors.New("SET_INTEREST_ERROR")
	ErrUserDisabled        = errors.New("USER_DISABLED")
	ErrUserNotFound        = errors.New("USER_NOT_FOUND")
	ErrWrongPassword       = errors.New("WRONG_PASSWORD")
	ErrUnauthorized        = errors.New("UNAUTHORIZED")
	ErrInvalidFilter       = errors.New("INVALID_FILTER")
)
