package go_routerutils

import (
	"errors"
	"net/http"
	"strconv"
)

var emptyURLParam error = errors.New("the given Placeholder was empty or not part of the URL")

func _readURLIntParam(r *http.Request, muxPatternPlaceholder string, base, bits int) (int64, error) {
	val := r.PathValue(muxPatternPlaceholder)
	if len(val) == 0 {
		return 0, emptyURLParam
	}

	out, err := strconv.ParseInt(val, base, bits)
	return out, err
}

func _readURLUIntParam(r *http.Request, muxPatternPlaceholder string, base, bits int) (uint64, error) {
	val := r.PathValue(muxPatternPlaceholder)
	if len(val) == 0 {
		return 0, emptyURLParam
	}

	out, err := strconv.ParseUint(val, base, bits)
	return out, err
}

func ReadInt64FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *int64) (err error) {
	if val, err := _readURLIntParam(r, muxPatternPlaceholder, 10, 64); err == nil {
		*readInto = val
	}

	return
}

func ReadInt32FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *int32) (err error) {
	if val, err := _readURLIntParam(r, muxPatternPlaceholder, 10, 32); err == nil {
		*readInto = int32(val)
	}

	return
}

func ReadInt16FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *int16) (err error) {
	if val, err := _readURLIntParam(r, muxPatternPlaceholder, 10, 16); err == nil {
		*readInto = int16(val)
	}

	return
}

func ReadInt8FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *int8) (err error) {
	if val, err := _readURLIntParam(r, muxPatternPlaceholder, 10, 8); err == nil {
		*readInto = int8(val)
	}

	return
}

func ReadUInt64FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *uint64) (err error) {
	if val, err := _readURLUIntParam(r, muxPatternPlaceholder, 10, 64); err == nil {
		*readInto = val
	}

	return
}

func ReadUInt32FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *uint32) (err error) {
	if val, err := _readURLUIntParam(r, muxPatternPlaceholder, 10, 32); err == nil {
		*readInto = uint32(val)
	}

	return
}

func ReadUInt16FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *uint16) (err error) {
	if val, err := _readURLUIntParam(r, muxPatternPlaceholder, 10, 16); err == nil {
		*readInto = uint16(val)
	}

	return
}

func ReadUInt8FromPathValue(r *http.Request, muxPatternPlaceholder string, readInto *uint8) (err error) {
	if val, err := _readURLUIntParam(r, muxPatternPlaceholder, 10, 8); err == nil {
		*readInto = uint8(val)
	}

	return
}

func ErrIsEmptyURLParamErr(err error) bool {
	return err == emptyURLParam
}
