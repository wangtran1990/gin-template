package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	configs "template/Configs"
	ecode "template/Constants"
	models "template/Models"

	log "github.com/sirupsen/logrus"

	guuid "github.com/google/uuid"
)

// PUBLIC Functions

// UUIDv4 ... Format c01d7cf6-ec3f-47f0-9556-a5d6e9009a43
func UUIDv4() string {
	uuid := guuid.New()
	return fmt.Sprintf("%s", uuid)
}

// LogTraceID ...
func LogTraceID() string {
	curTime := time.Now().Format(ecode.DATETIME_FORMAT_DEFAULT)
	return "TID" + ":" + curTime + ":" + fmt.Sprint(_randomNumber(1, 1e6))
}

// Logger ...
func Logger(requestID, extraData string) *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	if extraData == "" {
		return log.WithField("file", filename).WithField("func", fn).WithField("txn", requestID)
	}
	return log.WithField("file", filename).WithField("func", fn).WithField("txn", requestID).WithField("ext", extraData)
}

// MakeHTTPRequest ...
// Support method GET/POST/PUT/DELETE
// Support all content type
// Retry max 3 times for remote server issue (http response code 5xx)
func MakeHTTPRequest(method, url, contentType string, data interface{}, retry bool) (int, []byte, error) {
	const MAX_RETRY_TIMES = 3
	const RETRY_SLEEP_TIME = 50 // millisecond
	var resHTTPStatusCode int = 400
	var resHTTPBody []byte
	var req *http.Request
	var err error

	if data != nil {
		dataByte, _ := json.Marshal(data)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(dataByte))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		Logger("", "").Errorln(err)
		return resHTTPStatusCode, resHTTPBody, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 1; i <= MAX_RETRY_TIMES; i++ {
		var body []byte
		res, err := client.Do(req)
		if err == nil {
			defer res.Body.Close()
			body, _ = ioutil.ReadAll(res.Body)

			if res.StatusCode < 500 {
				return res.StatusCode, body, nil
			}
			// Only Retry when remote server error
		} else {
			Logger("", "").Errorln(err)
		}

		if !retry || i+1 > MAX_RETRY_TIMES {
			return res.StatusCode, body, errors.New(res.Status)
		}

		time.Sleep(RETRY_SLEEP_TIME * time.Millisecond)
		fmt.Println("RETRY")
	}

	return resHTTPStatusCode, resHTTPBody, errors.New("Max times retry exceeded")
}

// MigrateDataTable ...
func MigrateDataTable() {
	configs.DB.AutoMigrate(&models.User{})
}

// CheckDatabase ...
func CheckDatabase() string {
	var dbTime string
	queryStr := `SELECT SYSDATE()`
	row := configs.DB.Raw(queryStr).Row()
	row.Scan(&dbTime)

	return dbTime
}

/************************************************************/
// PRIVATE Functions                                        */
/************************************************************/

func _randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
