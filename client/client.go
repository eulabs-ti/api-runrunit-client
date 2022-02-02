package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/eucatur/go-toolbox/log"

	"github.com/jmoiron/sqlx/types"
	"github.com/parnurzeal/gorequest"
)

func logFile(response gorequest.Response, body []byte, errs []error) {
	var sent string

	if response.Request.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Request.Body)
		sent = buf.String()
	}

	logBody := struct {
		AppKey    string      `json:"X-Eucatur-Api-Id"`
		UserToken string      `json:"X-Eucatur-Api-Key"`
		Method    string      `json:"method"`
		Scheme    string      `json:"scheme"`
		Host      string      `json:"host"`
		URL       string      `json:"url"`
		Sent      string      `json:"sent"`
		Status    int         `json:"status"`
		Received  interface{} `json:"received"`
	}{
		AppKey:    response.Request.Header.Get(headerAppKey),
		UserToken: response.Request.Header.Get(headerUserToken),
		Method:    response.Request.Method,
		Scheme:    response.Request.URL.Scheme,
		Host:      response.Request.URL.Hostname(),
		URL:       response.Request.URL.RequestURI(),
		Sent:      sent,
		Status:    response.StatusCode,
		Received:  types.JSONText(body),
	}

	logBytes, err := json.Marshal(logBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = log.File(time.Now().Format("clients/api-runrunit/2006/01/02/15h.log"), string(logBytes))
	if err != nil {
		fmt.Println(err)
	}

	return
}

func (c *Client) post(targetURL string) (superAgent *gorequest.SuperAgent) {
	superAgent = gorequest.New().Post(c.Host+targetURL).
		AppendHeader(headerAppKey, c.AppKey).
		AppendHeader(headerUserToken, c.UserToken).
		AppendHeader("Content-Type", "application/json")
	return
}

func (c *Client) get(targetURL string) (superAgent *gorequest.SuperAgent) {
	superAgent = gorequest.New().Get(c.Host+targetURL).
		AppendHeader(headerAppKey, c.AppKey).
		AppendHeader(headerUserToken, c.UserToken).
		AppendHeader("Content-Type", "application/json")
	return
}

func (c *Client) put(targetURL string) (superAgent *gorequest.SuperAgent) {
	superAgent = gorequest.New().Put(c.Host+targetURL).
		AppendHeader(headerAppKey, c.AppKey).
		AppendHeader(headerUserToken, c.UserToken).
		AppendHeader("Content-Type", "application/json")
	return
}

func (c *Client) delete(targetURL string) (superAgent *gorequest.SuperAgent) {
	superAgent = gorequest.New().Delete(c.Host+targetURL).
		AppendHeader(headerAppKey, c.AppKey).
		AppendHeader(headerUserToken, c.UserToken).
		AppendHeader("Content-Type", "application/json")
	return
}

// GetOffDays - Doc. Group: Holidays - Met. GET OffDays
func (c *Client) GetOffDays() (offDays []OffDay, err error) {
	response, body, errs := c.get("off_days").EndBytes()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	if response.StatusCode == 200 {
		err = json.Unmarshal(body, &offDays)
		return
	}

	err = errors.New(string(body))
	return
}
