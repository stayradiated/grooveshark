package grooveshark

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/nu7hatch/gouuid"
)

var refreshTime = time.Duration(16 * time.Minute)

type SessionToken string
type SessionId string

type SessionTemplate struct {
	Salt           string
	Client         string
	ClientRevision string
}

var HtmlSharkSession = &SessionTemplate{
	Salt:           "nuggetsOfBaller",
	Client:         "htmlshark",
	ClientRevision: "20130520",
}

type Session struct {
	token       SessionToken
	lastUpdated time.Time

	Salt           string
	Client         string
	ClientRevision string
	Uuid           string
	SessionId      SessionId
	Country        *Country
}

// NewSession creates a new session
func NewSession(template *SessionTemplate) (session *Session) {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	session = &Session{
		Salt:           template.Salt,
		Client:         template.Client,
		ClientRevision: template.ClientRevision,
		Uuid:           id.String(),
	}

	return session
}

// Initiate fetches the necessary data from GrooveShark
func (s *Session) Initiate() {
	s.updateToken()
}

// Token makes sure that the token is still valid, and then returns it
func (s *Session) Token() SessionToken {
	if time.Now().Sub(s.lastUpdated) > refreshTime {
		s.updateToken()
	}
	return s.token
}

var preloadUrl = "http://grooveshark.com/preload.php?getCommunicationToken"
var preloadRegex = regexp.MustCompile("window\\.tokenData = (.*);")

type preloadData struct {
	SessionToken SessionToken `json:"getCommunicationToken"`
	Config       struct {
		Country   Country   `json:"country"`
		SessionId SessionId `json:"sessionID"`
	} `json:"getGSConfig"`
}

type Country struct {
	Id     int    `json:"ID"`
	CC1    int    `json:"CC1"`
	CC2    int    `json:"CC2"`
	CC3    int    `json:"CC3"`
	CC4    int    `json:"CC4"`
	DMA    int    `json:"DMA"`
	Iso    string `json:"iso"`
	Region string `json:"region"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
	IPR    int    `json:"IPR"`
}

func (s *Session) updateToken() {
	// fetch preload data
	res, err := http.Get(preloadUrl)
	if err != nil {
		panic(err)
	}

	// read entire file into memory
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// extract JSON out of body
	jsonData := preloadRegex.FindSubmatch(resBody)[1]

	// unmarshal JSON into a preloadData struct
	data := &preloadData{}
	if err = json.Unmarshal(jsonData, &data); err != nil {
		panic(err)
	}

	s.token = data.SessionToken
	s.Country = &data.Config.Country
	s.SessionId = data.Config.SessionId
	s.lastUpdated = time.Now()
}
