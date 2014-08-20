package grooveshark

type SessionId string
type Uuid string

type Country struct {
}

type Header struct {
	Client         string
	ClientRevision string
	Country        Country
	Privacy        int
	Session        SessionId
	uuid           Uuid
}

type Request struct {
	Url        string
	Parameters map[string]string
	Method     string
	Header     Header
}

func NewRequest(session *Session) (request *Request) {
	return request
}
