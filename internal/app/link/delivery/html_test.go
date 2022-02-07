package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/mocks"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/usecase"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeliveryHTML(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/html/form", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := getContext(req, rec)

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	token := &tokenMock{token: LT}
	d := New(u, "", token)

	// Assertions
	if assert.NoError(t, d.HTML(c)) {
		assert.Contains(t, rec.Body.String(), "/api/create")
	}
}
