package helper

// Based on SendGrid lib
import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/sendgrid/rest"
)

const (
	rateLimitRetry = 5    // times
	rateLimitSleep = 1100 // mili-seconds
)

// Options ...
type Options struct {
	Auth     string
	Endpoint string
	Host     string
}

// Client ...
type Client struct {
	rest.Request
}

func (o *Options) baseURL() string {
	return o.Host + o.Endpoint
}

// RequestNew create Request
// @return [Request] a default request object
func RequestNew(options Options) rest.Request {
	requestHeaders := map[string]string{
		"Authorization": options.Auth,
		"Accept":        "application/json",
	}

	return rest.Request{
		BaseURL: options.baseURL(),
		Headers: requestHeaders,
	}
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = rest.DefaultClient

// MakeRequest request synchronously
func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.Send(request)
}

// MakeRequestRetry a synchronous request, but retry in the event of a rate
// limited response.
func MakeRequestRetry(request rest.Request) (*rest.Response, error) {
	retry := 0
	var response *rest.Response
	var err error

	for {
		response, err = MakeRequest(request)
		if err != nil {
			return nil, err
		}
		Logger("", "").Errorln(err)

		if response.StatusCode != http.StatusTooManyRequests {
			return response, nil
		}

		if retry > rateLimitRetry {
			return nil, errors.New("rate limit retry exceeded")
		}
		retry++

		resetTime := time.Now().Add(rateLimitSleep * time.Millisecond)

		reset, ok := response.Headers["X-RateLimit-Reset"]
		if ok && len(reset) > 0 {
			t, err := strconv.Atoi(reset[0])
			if err == nil {
				resetTime = time.Unix(int64(t), 0)
			}
		}
		time.Sleep(resetTime.Sub(time.Now()))
	}
}

// MakeRequestAsync attempts a request asynchronously in a new go
// routine. This function returns two channels: responses
// and errors. This function will retry in the case of a
// rate limit.
func MakeRequestAsync(request rest.Request) (chan *rest.Response, chan error) {
	r := make(chan *rest.Response)
	e := make(chan error)

	go func() {
		response, err := MakeRequest(request)
		if err != nil {
			e <- err
		}
		if response != nil {
			r <- response
		}
	}()

	return r, e
}
