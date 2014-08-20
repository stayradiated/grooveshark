package grooveshark

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

type RequestSignature string

type Request struct {
	Url     string
	Method  string
	Session *Session
	Body    *Body
}

type Body struct {
	Method     *string     `json:"method"`
	Header     *Header     `json:"header"`
	Parameters *Parameters `json:"parameters"`
}

type Parameters map[string]string

type Header struct {
	Client         string           `json:"client"`
	ClientRevision string           `json:"clientRevision"`
	Token          RequestSignature `json:"token"`
	Privacy        int              `json:"privacy"`
	Country        *Country         `json:"country"`
	SessionId      SessionId        `json:"session"`
	Uuid           string           `json:"uuid"`
}

// NewRequest creates a new request
func NewRequest(session *Session, method string, parameters *Parameters) (request *Request) {
	request = &Request{
		Url:     "http://grooveshark.com/more.php?" + method,
		Method:  method,
		Session: session,
	}
	return request
}

// Sign generates a signature for this request
func (r *Request) Sign() {
	// get the session token
	token := r.Session.Token()

	// generate random nonce
	nonceBytes := make([]byte, 3)
	rand.Read(nonceBytes)
	nonce := hex.EncodeToString(nonceBytes)

	// concat values together
	input := strings.Join([]string{
		r.Method, string(token), r.Session.Salt, nonce,
	}, ":")

	fmt.Println(input)

	// hash with sha1
	hash := sha1.New()
	hash.Write([]byte(input))

	// convert to hex
	signature := hex.EncodeToString(hash.Sum(nil))

	// add to header
	header := *(*r.Body).Header
	header.Token = RequestSignature(signature)
}

// Send makes the request
func (r *Request) Send() {
	res, err := http.Post(r.Url, "application/json", &params)
}
