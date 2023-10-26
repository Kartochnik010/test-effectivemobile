package service

type RequestsService struct{}

func NewRequestsService() Requests {
	return &RequestsService{}
}

func (r *RequestsService) GetURL(url string) string {
	return ""
}
