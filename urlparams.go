package main

import (
	"errors"
	"net/http"
	"strconv"
)

var emptyURLParam error = errors.New("the given Placeholder was empty or not part of the URL")

func ReadInt64FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *int64) error {

	sValue := r.PathValue(muxPatternPlaceholder)

	if len(sValue) == 0 {
		return emptyURLParam
	}

	iValue, err := strconv.ParseInt(sValue, 10, 64)
	*readInto = iValue

	return err
}

func ErrIsEmptyURLParamErr(err error) bool {
	return err == emptyURLParam
}
