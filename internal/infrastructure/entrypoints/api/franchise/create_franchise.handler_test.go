package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/viictormg/clubHub/mocks"
	"github.com/viictormg/clubHub/mocks/stubs"
)

func TestCreateFranchiseHandler(t *testing.T) {
	t.Run("error malformed body", func(t *testing.T) {
		e := echo.New()

		franchiseUsecaseMock := new(mocks.IFranchiseUsecase)
		req := httptest.NewRequest(http.MethodPost, "/franchise/create", bytes.NewBuffer([]byte("1")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		franchiseHandler := NewFranchiseHandler(franchiseUsecaseMock)

		if assert.NoError(t, franchiseHandler.CreateFranchiseHandler(ctx)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
		}
		franchiseUsecaseMock.AssertExpectations(t)

	})

	t.Run("error validation franchise model", func(t *testing.T) {
		e := echo.New()
		franchiseUsecaseMock := new(mocks.IFranchiseUsecase)
		franchiseModelStub := stubs.GetFranchiseModelBadStub()
		payload, _ := json.Marshal(franchiseModelStub)

		req := httptest.NewRequest(http.MethodPost, "/franchise/create", bytes.NewBuffer(payload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		franchiseHandler := NewFranchiseHandler(franchiseUsecaseMock)

		if assert.NoError(t, franchiseHandler.CreateFranchiseHandler(ctx)) {
			assert.Equal(t, rec.Code, http.StatusBadRequest)
		}
		franchiseUsecaseMock.AssertExpectations(t)

	})

	t.Run("error calling usecase", func(t *testing.T) {
		e := echo.New()
		franchiseUsecaseMock := new(mocks.IFranchiseUsecase)
		franchiseModelStub := stubs.GetFranchiseModelStub()
		payload, _ := json.Marshal(franchiseModelStub)

		franchiseUsecaseMock.On("CreateFranchiseUsecase", franchiseModelStub).Return(nil, errors.New("test"))
		req := httptest.NewRequest(http.MethodPost, "/franchise/create", bytes.NewBuffer(payload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		franchiseHandler := NewFranchiseHandler(franchiseUsecaseMock)

		if assert.NoError(t, franchiseHandler.CreateFranchiseHandler(ctx)) {
			assert.Equal(t, rec.Code, http.StatusConflict)
		}
		franchiseUsecaseMock.AssertExpectations(t)

	})

	t.Run("create franchise usecase successful", func(t *testing.T) {
		e := echo.New()
		franchiseUsecaseMock := new(mocks.IFranchiseUsecase)
		franchiseModelStub := stubs.GetFranchiseModelStub()
		responseCreatedStub := stubs.GetResponseCreatedStub()
		payload, _ := json.Marshal(franchiseModelStub)

		franchiseUsecaseMock.On("CreateFranchiseUsecase", franchiseModelStub).Return(&responseCreatedStub, nil)
		req := httptest.NewRequest(http.MethodPost, "/franchise/create", bytes.NewBuffer(payload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		franchiseHandler := NewFranchiseHandler(franchiseUsecaseMock)

		if assert.NoError(t, franchiseHandler.CreateFranchiseHandler(ctx)) {
			assert.Equal(t, rec.Code, http.StatusCreated)
		}
		franchiseUsecaseMock.AssertExpectations(t)
	})
}
