package helper

import "github.com/go-playground/validator/v10"

type response struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data any) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	JsonResponse := response{
		Meta: meta,
		Data: data,
	}

	return JsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
