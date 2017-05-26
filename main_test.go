package main

import (
	"testing"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestApps(t *testing.T) {
	app := createApp()
	e := httptest.New(app, t)

	expectedResponse := map[string]interface{}{
		"message": "Hello",
		"name":    "John",
	}

	e.GET("/").Expect().Status(200).Body().Equal("hello modulars\n")
	e.GET("/hello").Expect().Status(200).JSON().Equal(expectedResponse)

}
