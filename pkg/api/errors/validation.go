package errors

type ValidationErrorType string

const (
	ValidationErrorTypeNotFound ValidationErrorType = "fieldValueNotFound"
	ValidationErrorRequired     ValidationErrorType = "fieldValueRequired"
	ValidationErrorDuplicate    ValidationErrorType = "fieldValueDuplicate"
	ValidationErrorInvalid      ValidationErrorType = "fieldValueInvalid"
	ValidationErrorNotSupported ValidationErrorType = "fieldValueNotSupported"
)
