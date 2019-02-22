package tapula

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type UserSuite struct {
	*ID
	Users         []*User   `json:"-"`
	ExecutionTime *Duration `json:"execution_time,omitempty"`
	TotalRequests int       `json:"total_requests,omitempty"`
	Metrics       *Metrics  `json:"metrics,omitempty"`
}

func (u *UserSuite) String() string {
	return toPrettyString(u)
}

func NewUserSuite() *UserSuite {
	logrus.Debugln(`NewUser(name string)`)

	return &UserSuite{
		Users:   []*User{},
		Metrics: &Metrics{},
	}
}

func (u *UserSuite) WithUser(user *User) *UserSuite {
	u.Users = append(u.Users, user)
	return u
}

func (u *UserSuite) Execute() {
	logrus.Debugln(`(u *UserSuite) Execute()`)

	for _, user := range u.Users {
		u.TotalRequests += user.TotalRequests
	}
	logrus.Infoln(`(u *UserSuite) Execute()`, "TotalRequests", u.TotalRequests)

	var start = time.Now()

	var wg sync.WaitGroup
	for _, user := range u.Users {
		wg.Add(1)

		thinkTime(&Duration{time.Duration(rand.Intn(20))})

		go func(wg *sync.WaitGroup, u1 *User) {
			defer wg.Done()
			u1.Execute(context.WithValue(context.Background(), "user", u1.Name))
		}(&wg, user)

	}
	wg.Wait()

	end := time.Now()
	u.ExecutionTime = &Duration{end.Sub(start)}

	u.collectMetrics()
}

func (u *UserSuite) collectMetrics() {

	for _, user := range u.Users {
		u.Metrics.Add(user.collectMetrics())
	}

	for _, m := range u.Metrics.Metrics {
		m.calculate()
	}

}
