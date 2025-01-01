package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const BODY_READ_THRESHOLD = 512

func ReadJSONBodyFromRequest(r *http.Request) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&output)
	return
}

func AssertJSONFieldType[T any](mp map[string]interface{}, key string) (out T, err error) {

	i, ok := mp[key]

	if ok {
		if v, ok := i.(T); ok {
			out = v

		} else {
			err = errors.New("expected type did not match found type")

		}

	} else {
		err = errors.New("given value seems invalid")

	}

	return

}

func ReadBodyFromRequest(r *http.Request) (body []byte, err error) {

	buffer := make([]byte, BODY_READ_THRESHOLD)
	read, err := r.Body.Read(buffer)

	if read == 0 {
		log.Println("INFO: no body provided")
		return
	}

	for read > 0 {

		body = append(body, buffer[:read]...)

		if err == io.EOF {
			break
		}

		read, err = r.Body.Read(buffer)
	}

	if err == io.EOF {
		err = nil
	}

	return

}
