package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/johnmcdnl/tapula"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

const (
	protocol        = "http"
	domain          = "tapula-api"
	port            = "3333"
	concurrentUsers = 10
	task1Items      = 5
	task2Items      = 5
	task3Items      = 5
)

func main() {
	userSuite := tapula.NewUserSuite()

	for i := 1; i <= concurrentUsers; i++ {
		userSuite.WithUser(sampleUser(strconv.Itoa(i)))
	}

	tapula.PrintReady()

	userSuite.Execute()
	fmt.Println(userSuite)

	_ = os.MkdirAll("./data", os.ModePerm)
	_ = ioutil.WriteFile("./data/metrics.json", []byte(userSuite.String()), os.ModePerm)
}

func sampleUser(name string) *tapula.User {

	user := tapula.NewUser(name)
	user.NavigateTo(listAllPage())

	for i := 1; i <= task1Items; i++ {
		user.NavigateTo(task1Page())
		user.MaybeNavigateTo(listAllPage())

		for i := 1; i <= task2Items; i++ {
			user.NavigateTo(task2Page())
			user.MaybeNavigateTo(task1Page())

			for i := 1; i <= task3Items; i++ {
				user.NavigateTo(task3Page())
				user.MaybeNavigateTo(task3Page())
			}

		}
	}

	return user
}

func listAllPage() *tapula.Page {
	logrus.Debugln(`listAllPage()`)

	return tapula.NewPage("listAllPage").
		WithThinkTime(randomThink(5*time.Second, 15*time.Second)).
		WithRequest(tapula.NewRequest("listAllReq").
			WithProtocol(protocol).
			WithDomain(domain).
			WithPort(port).
			WithPath(`/articles`).
			WithCookie(newCookie()),
		)
}

func task1Page() *tapula.Page {
	logrus.Debugln(`task1Page()`)

	return tapula.NewPage("task1Page").
		WithThinkTime(randomThink(3*time.Second, 10*time.Second)).
		WithRequest(tapula.NewRequest("task1Req").
			WithProtocol(protocol).
			WithDomain(domain).
			WithPort(port).
			WithPath("/articles/" + "1").
			WithCookie(newCookie()),
		)
}

func task2Page() *tapula.Page {
	logrus.Debugln(`task2Page()`)

	return tapula.NewPage("task2Page").
		WithThinkTime(randomThink(3*time.Second, 10*time.Second)).
		WithRequest(tapula.NewRequest("task2Req").
			WithProtocol(protocol).
			WithDomain(domain).
			WithPort(port).
			WithPath("/articles/" + "2").
			WithCookie(newCookie()),
		)
}

func task3Page() *tapula.Page {
	logrus.Debugln(`task3Page()`)

	return tapula.NewPage("task3Page").
		WithThinkTime(randomThink(3*time.Second, 10*time.Second)).
		WithRequest(tapula.NewRequest("task3Req").
			WithProtocol(protocol).
			WithDomain(domain).
			WithPort(port).
			WithPath("/articles/" + "3").
			WithCookie(newCookie()),
		)
}

func newCookie() *http.Cookie {
	return &http.Cookie{
		Domain: domain,
		Name:   "JSESSIONID",
		Value:  cookieValue(),
	}
}

func cookieValue() string {
	return "THIS_SHOULD_BE_A_REAL_JSESSIONID"
}

func randomThink(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max.Nanoseconds()) + int(min.Nanoseconds())))
}

func uniqueAppend(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
