package delivery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/mocks"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/usecase"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeliveryRedirect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/"+LT, strings.NewReader(""))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)
	c.SetPath("/:token")
	c.SetParamNames("token")
	c.SetParamValues(LT)

	token := &tokenMock{token: LT}
	d := New(u, "", token)
	var id int64
	mockRepo.EXPECT().FindByToken(c.Request().Context(), LT).Return(&models.Link{}, nil).Times(1)
	mockRepo.EXPECT().SaveStat(c.Request().Context(), id, c.RealIP()).Return(nil).Times(1)

	// Assertions
	if assert.NoError(t, d.Redirect(c)) {
		assert.Equal(t, http.StatusMovedPermanently, rec.Code)
	}
}

func TestDeliveryRedirectTokenIsEmpty(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/"+LT, strings.NewReader(""))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)

	token := &tokenMock{token: LT}
	d := New(u, "", token)
	mockRepo.EXPECT().FindByToken(nil, nil).Return(&models.Link{}, nil).Times(0)

	res := d.Redirect(c)
	errExpected := echo.NewHTTPError(http.StatusBadRequest,
		"link token can't be empty")
	// Assertions
	if assert.Error(t, res) {
		assert.Equal(t, errExpected, res)
	}
}

func TestDeliveryRedirectTokenNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/"+LT, strings.NewReader(""))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)
	c.SetPath("/:token")
	c.SetParamNames("token")
	c.SetParamValues(LT)

	token := &tokenMock{token: LT}
	d := New(u, "", token)
	var id int64
	e := errors.New("failed to select link from db")
	mockRepo.EXPECT().FindByToken(c.Request().Context(), LT).Return(nil, e).Times(1)
	mockRepo.EXPECT().SaveStat(c.Request().Context(), id, c.RealIP()).Return(nil).Times(0)

	res := d.Redirect(c)
	errExpected := echo.NewHTTPError(http.StatusBadRequest,
		"link not found in repo: failed to select link from db")
	// Assertions
	if assert.Error(t, res) {
		assert.Equal(t, errExpected, res)
	}
}

func TestDeliveryRedirectSaveStatFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/"+LT, strings.NewReader(""))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)
	c.SetPath("/:token")
	c.SetParamNames("token")
	c.SetParamValues(LT)

	token := &tokenMock{token: LT}
	d := New(u, "", token)
	var id int64
	e := errors.New("failed to insert stat to db")
	mockRepo.EXPECT().FindByToken(c.Request().Context(), LT).Return(&models.Link{}, nil).Times(1)
	mockRepo.EXPECT().SaveStat(c.Request().Context(), id, c.RealIP()).Return(e).Times(1)

	res := d.Redirect(c)
	// Assertions
	if assert.NoError(t, res) {
		assert.Equal(t, nil, res)
	}
}
