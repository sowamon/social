package models

type IResponse struct {
	Description string
	Data        interface{}
}

func Response(data interface{}, description string) IResponse {
	return IResponse{
		Data:        data,
		Description: description,
	}
}
