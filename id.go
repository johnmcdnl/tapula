package tapula

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/sirupsen/logrus"
)

type ID struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewID(name string) *ID {
	logrus.Debugln(`NewID(name string)`)

	return &ID{
		ID:   strconv.Itoa(rand.Intn(math.MaxInt64)),
		Name: name,
	}
}
