package babelcoin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
)

type HttpError struct {
	NestedError  error
	ResponseBody []byte
}

func (e *HttpError) Error() string {
	return e.NestedError.Error()
}

func HttpGetJson(url string, m interface{}) *HttpError {
	bytes, error := HttpDurableGet(url, 10)

	if error != nil {
		return &HttpError{error, bytes}
	}

	if error = json.Unmarshal(bytes, &m); error != nil {
		return &HttpError{error, bytes}
	}

	return nil
}

func HttpDurableGet(url string, times int) ([]byte, error) {
	var body []byte

	for i := 0; i < times; i++ {
		resp, err := http.Get(url)

		if err != nil && i == times {
			return []byte{}, err
		} else if err != nil {
			continue
		}

		body, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil && i == times {
			return []byte{}, err
		} else if err != nil {
			continue
		} else {
			break
		}
	}

	return body, nil
}
