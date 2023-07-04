package core

type RequestInterface interface {
	GetURL() string
	GetMethod() string
	GetVersion() string
	GetHeaders() map[string]string
}

// CtyunRequest is the base struct of service requests
type CtyunRequest struct {
	URL     string
	Method  string
	Header  map[string]string
	Version string
}

func (r CtyunRequest) GetURL() string {
	return r.URL
}

func (r CtyunRequest) GetMethod() string {
	return r.Method
}

func (r CtyunRequest) GetVersion() string {
	return r.Version
}

func (r CtyunRequest) GetHeaders() map[string]string {
	return r.Header
}

func (r *CtyunRequest) AddHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}
