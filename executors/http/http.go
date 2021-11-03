package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
)

const (
	headerSeparator = `|`
	headerLength    = 2
	keyIndex        = 0
	valueIndex      = 1
)

const (
	MethodPost   = `post`
	MethodPut    = `put`
	MethodGet    = `get`
	MethodDelete = `delete`
)

func Get(scheduler *pb.Scheduler) (res string, err error) {
	req, err := http.NewRequest(http.MethodGet, scheduler.Url, bytes.NewBuffer([]byte(scheduler.Body)))
	if err != nil {
		return res, err
	}
	for _, header := range scheduler.Headers {
		keyValue := strings.Split(header, headerSeparator)
		if len(keyValue) == headerLength {
			req.Header.Set(keyValue[keyIndex], keyValue[valueIndex])
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode > 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("http status code %d error response %s", resp.StatusCode, string(bodyBytes))
		return res, errors.New(msg)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(bodyBytes), nil
}

func Post(scheduler *pb.Scheduler) (res string, err error) {
	req, err := http.NewRequest(http.MethodPost, scheduler.Url, bytes.NewBuffer([]byte(scheduler.Body)))
	if err != nil {
		return res, err
	}
	for _, header := range scheduler.Headers {
		keyValue := strings.Split(header, headerSeparator)
		if len(keyValue) == headerLength {
			req.Header.Set(keyValue[keyIndex], keyValue[valueIndex])
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode > 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("http status code %d error response %s", resp.StatusCode, string(bodyBytes))
		return res, errors.New(msg)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(bodyBytes), nil
}

func Put(scheduler *pb.Scheduler) (res string, err error) {
	req, err := http.NewRequest(http.MethodPut, scheduler.Url, bytes.NewBuffer([]byte(scheduler.Body)))
	if err != nil {
		return res, err
	}
	for _, header := range scheduler.Headers {
		keyValue := strings.Split(header, headerSeparator)
		if len(keyValue) == headerLength {
			req.Header.Set(keyValue[keyIndex], keyValue[valueIndex])
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode > 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("http status code %d error response %s", resp.StatusCode, string(bodyBytes))
		return res, errors.New(msg)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(bodyBytes), nil
}

func Delete(scheduler *pb.Scheduler) (res string, err error) {
	req, err := http.NewRequest(http.MethodDelete, scheduler.Url, bytes.NewBuffer([]byte(scheduler.Body)))
	if err != nil {
		return res, err
	}
	for _, header := range scheduler.Headers {
		keyValue := strings.Split(header, headerSeparator)
		if len(keyValue) == headerLength {
			req.Header.Set(keyValue[keyIndex], keyValue[valueIndex])
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode > 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		msg := fmt.Sprintf("http status code %d error response %s", resp.StatusCode, string(bodyBytes))
		return res, errors.New(msg)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return string(bodyBytes), nil
}
