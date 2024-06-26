// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: © 2015 LabStack LLC and Echo contributors

package echo

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupFileFS(t *testing.T) {
	var testCases = []struct {
		name             string
		whenPath         string
		whenFile         string
		whenFS           fs.FS
		givenURL         string
		expectCode       int
		expectStartsWith []byte
	}{
		{
			name:             "ok",
			whenPath:         literal_6815,
			whenFS:           os.DirFS(literal_0963),
			whenFile:         "walle.png",
			givenURL:         "/assets/walle",
			expectCode:       http.StatusOK,
			expectStartsWith: []byte{0x89, 0x50, 0x4e},
		},
		{
			name:             "nok, requesting invalid path",
			whenPath:         literal_6815,
			whenFS:           os.DirFS(literal_0963),
			whenFile:         "walle.png",
			givenURL:         "/assets/walle.png",
			expectCode:       http.StatusNotFound,
			expectStartsWith: []byte(`{"message":"Not Found"}`),
		},
		{
			name:             "nok, serving not existent file from filesystem",
			whenPath:         literal_6815,
			whenFS:           os.DirFS(literal_0963),
			whenFile:         "not-existent.png",
			givenURL:         "/assets/walle",
			expectCode:       http.StatusNotFound,
			expectStartsWith: []byte(`{"message":"Not Found"}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			g := e.Group("/assets")
			g.FileFS(tc.whenPath, tc.whenFile, tc.whenFS)

			req := httptest.NewRequest(http.MethodGet, tc.givenURL, nil)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectCode, rec.Code)

			body := rec.Body.Bytes()
			if len(body) > len(tc.expectStartsWith) {
				body = body[:len(tc.expectStartsWith)]
			}
			assert.Equal(t, tc.expectStartsWith, body)
		})
	}
}

func TestGroupStaticPanic(t *testing.T) {
	var testCases = []struct {
		name        string
		givenRoot   string
		expectError string
	}{
		{
			name:        "panics for ../",
			givenRoot:   "../images",
			expectError: "can not create sub FS, invalid root given, err: sub ../images: invalid name",
		},
		{
			name:        "panics for /",
			givenRoot:   "/images",
			expectError: "can not create sub FS, invalid root given, err: sub /images: invalid name",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := New()
			e.Filesystem = os.DirFS("./")

			g := e.Group("/assets")

			assert.PanicsWithError(t, tc.expectError, func() {
				g.Static("/images", tc.givenRoot)
			})
		})
	}
}

const literal_6815 = "/walle"

const literal_0963 = "_fixture/images"
