package main_test

import (
	"time"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	main "github.com/stiyyagura0901/merge-queue-poc"
)

func TestMain(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	time.Sleep(time.Duration(50) * time.Second)
	c := e.NewContext(req, rec)

	if assert.NoError(t, main.Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
