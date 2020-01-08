package end_to_end

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	baloo "gopkg.in/h2non/baloo.v3"
)

var (
	BodyString string
	doOnce     sync.Once
	Client     *baloo.Client
)

func MakeEmailAddress() string {
	return fmt.Sprintf("%d.%d@tempmail.com", Seed().Intn(10000), Seed().Intn(10000))
}

func Seed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetBody(res *http.Response, req *http.Request) error {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	BodyString = string(bodyBytes)
	return err
}

func MakeClient() {
	doOnce.Do(func() {
		Client = baloo.New("http://localhost:5000")
	})
}

const GeneralResponseSchema = `{
	"title": "Response",
	"type": "object",
	"properties": {
		"data": {
			"type": "object"
		}
	},
	"required": ["data"]
}`

func Init() {
	MakeClient()
}
