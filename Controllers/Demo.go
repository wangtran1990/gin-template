//Controllers/Demo.go

package controllers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	H "template/Helper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/rest"
)

// HTTPRequestDemo ...
func HTTPRequestDemo(c *gin.Context) {
	requestID := c.GetString("x-request-id")
	var ops H.Options
	ops.Host = "https://google.com"

	request := H.RequestNew(ops)

	request.Method = http.MethodGet

	response, _ := H.MakeRequest(request)
	H.Logger(requestID, "").Infoln("response.StatusCode= ", response.StatusCode)
	H.Logger(requestID, "").Infoln("response.Headers= ", response.Headers)

	c.JSON(http.StatusOK, "")
	return
}

// HTTPRetryRequestDemo ...
func HTTPRetryRequestDemo(c *gin.Context) {
	requestID := c.GetString("x-request-id")

	// Create fake server to reply only http status code = 429
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", strconv.Itoa(int(time.Now().Add(1*time.Second).Unix())))
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer fakeServer.Close()

	// Create options
	var ops H.Options
	ops.Host = fakeServer.URL
	request := H.RequestNew(ops)
	request.Method = http.MethodGet

	// Create custom client objects for some reasons (such as modifies timeout value)
	var custom rest.Client
	custom.HTTPClient = &http.Client{Timeout: time.Millisecond * 10}
	H.DefaultClient = &custom

	response, err := H.MakeRequestRetry(request)
	if err != nil {
		H.Logger(requestID, "").Errorln("msg= ", err)
		return
	}
	H.Logger(requestID, "").Infoln("response.StatusCode= ", response.StatusCode)
	H.Logger(requestID, "").Infoln("response.Headers= ", response.Headers)
	c.JSON(http.StatusOK, "")
	return
}

// HTTPAsyncRequestDemo ...
func HTTPAsyncRequestDemo(c *gin.Context) {

	requestID := c.GetString("x-request-id")
	const maxUrls = 3
	urls := []string{
		"https://duckduckgo.com",
		"https://bing.com",
		"https://google.com",
	}

	var results [maxUrls]chan *rest.Response
	var errors [maxUrls]chan error

	var httpWG sync.WaitGroup
	httpWG.Add(len(urls))

	for idx, url := range urls {
		func(url string) {
			defer httpWG.Done()

			var ops H.Options
			ops.Host = url
			request := H.RequestNew(ops)
			request.Method = http.MethodGet
			results[idx], errors[idx] = H.MakeRequestAsync(request)
		}(url)
	}
	httpWG.Wait() // Wait for all tasks done

	// Process results
	for i := 0; i < len(results); i++ {
		select {
		case val := <-results[i]:
			// Process value here
			H.Logger(requestID, "").Infoln("response.StatusCode= ", val.StatusCode)
		case <-time.After(time.Second * 10):
			H.Logger(requestID, "").Infoln("TIME OUT RESULT")
			select {
			case err := <-errors[i]:
				// Process error here
				H.Logger(requestID, "").Errorln(err)
			case <-time.After(time.Microsecond * 100):
				H.Logger(requestID, "").Infoln("TIME OUT ERROR ALREADY")
			}
		}
	}

	c.JSON(http.StatusOK, "")
	return
}
