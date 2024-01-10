package models

type IResponse struct {
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func Response(data interface{}, description string) IResponse {
	return IResponse{
		Data:        data,
		Description: description,
	}
}
