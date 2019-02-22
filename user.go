package tapula

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type User struct {
	*ID
	Pages             []*Page          `json:"pages,omitempty"`
	TotalRequests     int              `json:"total_requests,omitempty"`
	EstimatedDuration *Duration        `json:"estimated_duration,omitempty"`
	Client            *http.Client     `json:"-"`
	Stutter           func() *Duration `json:"-"`
}

func (u *User) String() string {
	return toString(u)
}

func NewUser(name string) *User {
	logrus.Debugln(`NewUser(name string)`)

	return &User{
		ID:                NewID(name),
		Pages:             []*Page{},
		Client:            http.DefaultClient,
		EstimatedDuration: &Duration{0},
		Stutter: func() *Duration {
			return &Duration{time.Duration(time.Duration(rand.Intn(10)) * time.Second)}
		},
	}
}

func (u *User) NavigateTo(page *Page) *User {
	logrus.Debugln(`(u *User) NavigateTo(page *Page)`)

	u.Pages = append(u.Pages, page)
	u.TotalRequests = len(u.Pages)
	u.EstimatedDuration = &Duration{u.EstimatedDuration.Duration + page.ThinkTime.Duration + u.Stutter().Duration}

	return u
}

func (u *User) NavigateToEach(pages []*Page) *User {
	logrus.Debugln(`(u *User) NavigateToEach(pages []*Page)`)

	for _, p := range pages {
		u.NavigateTo(p)
	}
	return u
}

var pageReqCount = 0

func (u *User) Execute(ctx context.Context) {
	logrus.Debugln(`(u *User) Execute()`)
	logrus.Infoln(`(u *User) Execute()`, "TotalRequests", u.TotalRequests, "EstimatedDuration", u.EstimatedDuration)

	for _, p := range u.Pages {
		pageReqCount++
		p.Execute(context.WithValue(ctx, "pageReqCount", pageReqCount), u.Client)
		u.doStutter()
	}
}

func (u *User) collectMetrics() []*Metric {
	logrus.Debugln(`(u *User) collectMetrics()`)

	var metrics []*Metric

	for _, p := range u.Pages {

		var foundMetric bool

		for _, m := range metrics {
			if p.Name == m.Endpoint {
				m.Durations = append(m.Durations, p.ExecutionTime)
				foundMetric = true
				continue
			}
		}

		if !foundMetric {
			metrics = append(metrics, &Metric{Endpoint: p.Name, Durations: []*Duration{p.ExecutionTime}})
		}
	}

	for _, m := range metrics {
		m.calculate()
	}

	return metrics

}

func (u *User) doStutter() {
	time.Sleep(u.Stutter().Duration)
}

func (u *User) MaybeNavigateTo(page *Page) {
	if rand.Intn(5)%3 == 0 {
		u.NavigateTo(page)
	}
}
