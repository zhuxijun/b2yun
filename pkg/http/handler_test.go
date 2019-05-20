package http_test

import (
	"net/http/httptest"
	"testing"
	root "b2yun/pkg"
	"b2yun/pkg/http"
)

func TestHandler(t *testing.T) {

	handler := http.NewHandler(root.NewLogStdOut())
	handler.Init(http.Services{})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := server.Client()

	_, err := client.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	_, err = client.Get(server.URL + "/api/goods/")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Get(server.URL + "/api/images/")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Get(server.URL + "/api/users/")
	if err != nil {
		t.Error(err)
	}

	_, err = client.Get(server.URL + "/others/")
	if err != nil {
		t.Error(err)
	}

}
