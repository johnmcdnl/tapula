package tapula

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type Request struct {
	*ID
	Method   string         `json:"method,omitempty"`
	Protocol string         `json:"protocol,omitempty"`
	Domain   string         `json:"domain,omitempty"`
	Port     string         `json:"port,omitempty"`
	Path     string         `json:"path,omitempty"`
	Header   http.Header    `json:"header,omitempty"`
	Body     string         `json:"body,omitempty"`
	Cookies  []*http.Cookie `json:"-"`
	req      *http.Request
}

func (r *Request) String() string {
	return toString(r)
}

func NewRequest(name string) *Request {
	logrus.Debugln(`NewRequest(name string)`)

	return &Request{
		ID:       NewID(name),
		Method:   http.MethodGet,
		Protocol: "https",
	}
}

func (r *Request) WithMethod(method string) *Request {
	logrus.Debugln(`(r *Request) WithMethod(method string)`)

	r.Method = method
	return r
}

func (r *Request) WithProtocol(protocol string) *Request {
	logrus.Debugln(`(r *Request) WithProtocol(protocol string)`)

	r.Protocol = protocol
	return r
}

func (r *Request) WithDomain(domain string) *Request {
	logrus.Debugln(`(r *Request) WithDomain(domain string)`)

	r.Domain = domain
	return r
}

func (r *Request) WithPath(path string) *Request {
	logrus.Debugln(`(r *Request) WithPath(path string)`)

	r.Path = path
	return r
}

func (r *Request) WithBody(body string) *Request {
	logrus.Debugln(`(r *Request) WithBody(body string)`)

	r.Body = body
	return r
}

func (r *Request) WithHeader(header http.Header) *Request {
	logrus.Debugln(`(r *Request) WithHeader(header http.Header)`)

	r.Header = header
	return r
}

func (r *Request) WithCookie(cookie *http.Cookie) *Request {
	r.Cookies = append(r.Cookies, cookie)
	return r
}

func (r *Request) build() {
	body := strings.NewReader(r.Body)

	buildURI := func() string {

		var uri string
		uri += r.Protocol
		uri += "://"
		uri += r.Domain

		if r.Port != "" {
			uri += ":"
			uri += r.Port
		}
		uri += r.Path

		return uri

	}

	req, err := http.NewRequest(r.Method, buildURI(), body)
	if err != nil {
		panic(err)
	}
	if r.Header != nil {
		req.Header = r.Header
	}

	for _, c := range r.Cookies {
		req.AddCookie(c)
	}

	r.req = req
}

func (r *Request) Execute(client *http.Client) []byte {
	logrus.Infoln(`(p *Request) Execute() ---`, r)

	r.build()

	resp, err := client.Do(r.req)
	if err != nil {
		panic(err)
	}
	defer ensuredClosed(resp)

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= http.StatusBadRequest {
		panic("bad request" + resp.Status + " " + strconv.Itoa(resp.StatusCode))
	}

	return body
}

func (r *Request) WithPort(port string) *Request {
	r.Port = port
	return r
}
