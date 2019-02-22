package tapula

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Page struct {
	*ID
	Request       *Request  `json:"request"`
	ThinkTime     Duration  `json:"think_time,omitempty"`
	ExecutionTime *Duration `json:"execution_time,omitempty"`
}

func (p *Page) String() string {
	return toString(p)
}

func NewPage(name string) *Page {
	logrus.Debugln(`NewPage(name string)`)

	return &Page{
		ID:        NewID(name),
		ThinkTime: Duration{1 * time.Second},
	}
}

func (p *Page) WithThinkTime(d time.Duration) *Page {
	logrus.Debugln(`(p *Page) WithThinkTime(d time.Duration)`)

	p.ThinkTime = Duration{d}
	return p
}

func (p *Page) WithRequest(req *Request) *Page {
	logrus.Debugln(`(p *Page) WithRequest(req *Request)`)

	p.Request = req
	return p
}

func (p *Page) Execute(ctx context.Context, client *http.Client) {
	logrus.Infoln("(p *Page) Execute() --- ", p, ctx)

	start := time.Now()
	p.Request.Execute(client)
	end := time.Now()

	p.ExecutionTime = &Duration{end.Sub(start)}
	logrus.Infoln("(p *Page) Execute() --- ", p.String())
	p.think()
}

func (p *Page) think() {
	thinkTime(&p.ThinkTime)
}
