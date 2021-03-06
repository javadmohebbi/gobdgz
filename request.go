package gobdgz

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Request struct {
	ID      string                 `json:"id"`
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`

	APIKey string `json:"-"`

	URL        string `json:"-"`
	HttpMethod string `json:"-"`

	httpRequest *http.Request

	Debug bool `json:"-"`
}

func (r *Request) StringJSON() string {
	// if ID is empty, It should set to default
	// value of null
	if r.ID == "" {
		r.ID = "NULL"
	}

	jsonString, err := json.Marshal(r)

	if err != nil {
		return "{}"
	} else {
		return string(jsonString)
	}
}

func (r *Request) SendRequest(response interface{}) error {

	if r.Debug {
		log.Println("Preparing Request...")
	}

	// HTTP request body building :-D
	payload := []byte(r.StringJSON())
	if r.Debug {
		log.Printf("Request body is: \n\t%v\n", string(payload))
	}

	// preparing HTTP request
	req, err := http.NewRequest(r.HttpMethod, r.URL, bytes.NewReader(payload))
	if err != nil {
		log.Println("Error on creating request object. ", err.Error())
		return err
	}

	// setting the request headers
	strAuthorization := base64.StdEncoding.EncodeToString([]byte(r.APIKey + ":"))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Basic "+strAuthorization)
	// req.Header.Set("Authorization", "Basic MzczZmI2N2M1ZjI4OTQwOTI5YmViN2Y5Y2ZhMzUwNGQzMTBlYTcxZmE3ZDkwYjJhMTNiZGE0NDhmODYxZmE5Yzo=")

	// prepare HTTP client
	client := &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	defer client.CloseIdleConnections()

	if r.Debug {
		log.Println("Dispatching the request")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("Error on dispatching request.", err.Error())
		return err
	}
	defer res.Body.Close()

	if r.Debug {
		log.Println("Retrieving and parsing the response")

	}

	respBdy, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if r.Debug {
			log.Println("Can not read response body. ", err)
		}
		return err
	}

	if r.Debug {
		log.Printf("Got %v HTTP status code with description %v", res.StatusCode, res.Status)
	}

	// check if request is OK
	if res.StatusCode >= 200 && res.StatusCode <= 299 {

		// check if HTTP 200 OK
		// has error message

		// unmarshal response for an error
		respBodyCopy := respBdy
		var respErr ResponseError
		err = json.Unmarshal(respBodyCopy, &respErr)
		if err != nil {
			if r.Debug {
				log.Println("Can not unmarshal body. ", err)
			}
		}
		if respErr.Error.Code != 0 {
			errorStr := fmt.Sprintf("HTTP request returns an error: \n\t%v (%v): %v", respErr.Error.Message, respErr.Error.Code, respErr.Error.Data.Details)
			if r.Debug {
				log.Println(errorStr)
			}
			return errors.New(errorStr)
		}

		if r.Debug {
			log.Println(string(respBdy))
		}

		// OK response
		err = json.Unmarshal(respBdy, &response)
		if err != nil {
			if r.Debug {
				log.Println("Can not unmarshal body. ", err)
			}
			return err
		}

		return nil
	}

	// Request has error
	// http other than 2XX code
	return errors.New(fmt.Sprintf("Request error %v code with status %v", res.StatusCode, res.Status))
}
