package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenTheResponseIsNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code) // проверяем статус ОК

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body) // проверяем строку

}

func TestMainHandlerWhenTheCityIsNotSupported(t *testing.T) {
	totalString := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=4&city=yekaterinburg", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) // проверяем статус BadRequest

	body := responseRecorder.Body.String()
	assert.Equal(t, totalString, body) // сравниваенм текст возвращаемого ответа

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code) // проверяем статус ОК

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, totalCount, len(list)) // сравниваем количество кафе

}
