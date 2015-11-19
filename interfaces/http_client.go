package interfaces

type HttpClient interface {
	Get(url string, out interface{}) error
}
