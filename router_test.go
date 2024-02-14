package sesame

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = NewRouter()

func TestGet(t *testing.T) {
	router.Get("/test/get", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	resp, err := s.Client().Get(s.URL + "/test/get")
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestPost(t *testing.T) {
	router.Post("/test/post", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	resp, err := s.Client().Post(s.URL+"/test/post", "text/plain", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestPut(t *testing.T) {
	router.Put("/test/put", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	req, err := http.NewRequest(http.MethodPut, s.URL+"/test/put", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	resp, err := s.Client().Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestPatch(t *testing.T) {
	router.Patch("/test/patch", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	req, err := http.NewRequest(http.MethodPatch, s.URL+"/test/patch", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	resp, err := s.Client().Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestOptions(t *testing.T) {
	router.Options("/test/options", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	req, err := http.NewRequest(http.MethodOptions, s.URL+"/test/options", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	resp, err := s.Client().Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestHead(t *testing.T) {
	router.Get("/test/head", func(ctx Context) error {
		return nil
	})
	s := httptest.NewServer(router.mux)
	req, err := http.NewRequest(http.MethodHead, s.URL+"/test/head", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	resp, err := s.Client().Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Status)
}

func TestRouter(t *testing.T) {
	r := NewRouter()
	go func() {
		r.StartServer(":5000")
	}()
}
