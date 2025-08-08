package util

// DataWrapper is a generic struct to wrap data in a response, used for swagger documentation
type DataWrapper[T any] struct {
	Data T `json:"data"`
}

// MessageWrapper wraps a message in a response, used for swagger documentation
type MessageWrapper struct {
	Message string `json:"message"`
}

type ListDataWrapper[T any] struct {
	Data struct {
		List T `json:"list"`
	} `json:"data"`
}

func WrapResponse[T any](data T) *DataWrapper[T] {
	return &DataWrapper[T]{Data: data}
}

func WrapListResponse[T any](data T) *ListDataWrapper[T] {
	return &ListDataWrapper[T]{Data: struct {
		List T `json:"list"`
	}{List: data}}
}

func MessageResponse(message string) *MessageWrapper {
	return &MessageWrapper{Message: message}
}
