package tapula

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func ensuredClosed(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		logrus.Errorln(err)
	}
}
