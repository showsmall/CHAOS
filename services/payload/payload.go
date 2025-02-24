package payload

type Data struct {
	Request     string `json:"request,omitempty"`
	Response    []byte `json:"response,omitempty"`
	HasError    bool   `json:"has_error,omitempty"`
	HasResponse bool   `json:"has_response,omitempty"`
}

type Service interface {
	Set(string, *Data)
	Get(string) (*Data, bool)
	Remove(string)
}
