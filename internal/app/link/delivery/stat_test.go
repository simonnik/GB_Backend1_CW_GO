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

func TestDeliveryStat(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/html/stat/"+LT, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)
	c.SetPath("/html/stat/:token")
	c.SetParamNames("token")
	c.SetParamValues(LT)

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	token := &tokenMock{token: LT}
	d := New(u, "", token)

	mockRepo.EXPECT().FindAllByToken(c.Request().Context(), LT).Return(models.StatList{models.Stat{
		ID:      1,
		Link:    "https://google.com",
		IP:      "0.0.0.0",
		Created: "2022-02-05 17:28:22.716313",
	}}, nil).Times(1)
	// Assertions
	if assert.NoError(t, d.Stat(c)) {
		assert.Contains(t, rec.Body.String(), "https://google.com")
	}
}

func TestDeliveryStatTokenIsEmpty(t *testing.T) {
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
	mockRepo.EXPECT().FindAllByToken(nil, nil).Return(nil, nil).Times(0)

	res := d.Stat(c)
	errExpected := echo.NewHTTPError(http.StatusBadRequest,
		"link id can't be empty")
	// Assertions
	if assert.Error(t, res) {
		assert.Equal(t, errExpected, res)
	}
}

func TestDeliveryStatLinksNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodGet, "/"+LT, strings.NewReader(""))
	//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)
	c.SetPath("/html/stat/:token")
	c.SetParamNames("token")
	c.SetParamValues(LT)

	token := &tokenMock{token: LT}
	d := New(u, "", token)
	e := errors.New("failed to select link from db")
	mockRepo.EXPECT().FindAllByToken(c.Request().Context(), LT).Return(nil, e).Times(1)

	res := d.Stat(c)
	errExpected := echo.NewHTTPError(http.StatusBadRequest,
		"failed to select link from db")
	// Assertions
	if assert.Error(t, res) {
		assert.Equal(t, errExpected, res)
	}
}
