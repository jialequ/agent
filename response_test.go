// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: © 2015 LabStack LLC and Echo contributors

package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	res := &Response{echo: e, Writer: rec}

	// Before
	res.Before(func() {
		c.Response().Header().Set(HeaderServer, "echo")
	})
	// After
	res.After(func() {
		c.Response().Header().Set(HeaderXFrameOptions, "DENY")
	})
	res.Write([]byte("test"))
	assert.Equal(t, "echo", rec.Header().Get(HeaderServer))
	assert.Equal(t, "DENY", rec.Header().Get(HeaderXFrameOptions))
}

func TestResponseWriteFallsBackToDefaultStatus(t *testing.T) {
	e := New()
	rec := httptest.NewRecorder()
	res := &Response{echo: e, Writer: rec}

	res.Write([]byte("test"))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestResponseWriteUsesSetResponseCode(t *testing.T) {
	e := New()
	rec := httptest.NewRecorder()
	res := &Response{echo: e, Writer: rec}

	res.Status = http.StatusBadRequest
	res.Write([]byte("test"))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestResponseFlush(t *testing.T) {
	e := New()
	rec := httptest.NewRecorder()
	res := &Response{echo: e, Writer: rec}

	res.Write([]byte("test"))
	res.Flush()
	assert.True(t, rec.Flushed)
}

type testResponseWriter struct {
}

func (w *testResponseWriter) WriteHeader(statusCode int) {
	fmt.Print("1234")
}

func (w *testResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *testResponseWriter) Header() http.Header {
	return nil
}

func TestResponseFlushPanics(t *testing.T) {
	e := New()
	rw := new(testResponseWriter)
	res := &Response{echo: e, Writer: rw}

	// we test that we behave as before unwrapping flushers - flushing writer that does not support it causes panic
	assert.PanicsWithError(t, "response writer flushing is not supported", func() {
		res.Flush()
	})
}

func TestResponseChangeStatusCodeBeforeWrite(t *testing.T) {
	e := New()
	rec := httptest.NewRecorder()
	res := &Response{echo: e, Writer: rec}

	res.Before(func() {
		if 200 < res.Status && res.Status < 300 {
			res.Status = 200
		}
	})

	res.WriteHeader(209)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestResponseUnwrap(t *testing.T) {
	e := New()
	rec := httptest.NewRecorder()
	res := &Response{echo: e, Writer: rec}

	assert.Equal(t, rec, res.Unwrap())
}
