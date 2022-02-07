package delivery

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/mocks"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/usecase"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/config"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
	contextUtils "github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/context"
	"github.com/stretchr/testify/assert"
)

// LT link token
const LT = "1q2w3e4r"

var (
	mockDB = []models.Link{
		{0, "https://google.com", LT},
		{0, "https://googlecom", LT},
	}
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type tokenMock struct {
	token string
}

func (t tokenMock) Generate() string {
	return t.token
}

func TestDeliveryCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	payload := `{"link": "https://google.com"}`
	response := `{"link":"https://shrt.io/1q2w3e4r","stat":"https://shrt.io/html/stat/1q2w3e4r"}` + "\n"

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := getContext(req, rec)
	token := &tokenMock{token: LT}
	d := New(u, "", token)
	mockRepo.EXPECT().Create(c.Request().Context(), &mockDB[0]).Return(nil).Times(1)

	// Assertions
	if assert.NoError(t, d.Create(c)) {
		assert.Equal(t, response, rec.Body.String())
	}
}

func TestDeliveryCreateValidateFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	payload := `{"linkh": "https://googlecom"}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := getContext(req, rec)
	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	token := &tokenMock{token: LT}
	d := New(u, "", token)
	err := echo.NewHTTPError(http.StatusBadRequest, "link's validate failed: link can't be empty")
	mockRepo.EXPECT().Create(nil, nil).Return(nil).Times(0)

	e := d.Create(c)
	// Assertions
	if assert.Error(t, e) {
		assert.Equal(t, err, e)
	}
}

func TestDeliveryCreateFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	payload := `{"link": "https://google.com"}`

	mockRepo := mocks.NewMockRepository(mockCtrl)
	u := usecase.New(mockRepo)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := getContext(req, rec)
	token := &tokenMock{token: LT}
	d := New(u, "", token)
	errExpected := echo.NewHTTPError(http.StatusBadRequest,
		"failed to create link in repo: failed to insert link to db")
	err := errors.New("failed to insert link to db")
	mockRepo.EXPECT().Create(c.Request().Context(), &mockDB[0]).Return(err).Times(1)

	e := d.Create(c)
	// Assertions
	if assert.Error(t, e) {
		assert.Equal(t, errExpected, e)
	}
}

func getContext(req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	// Setup
	e := echo.New()
	fs := os.DirFS("../../../../web/template")
	p, err := template.ParseFS(fs, "*.html", "*/*.html")
	t := &Template{
		templates: template.Must(p, err),
	}
	e.Renderer = t
	e.GET("/:token", nil).Name = "redirect"
	e.GET("/html/stat/:token", nil).Name = "stat"
	c := e.NewContext(req, rec)
	cfg, _ := config.BuildConfig("../../../../configs/config.yml")
	newCtx := contextUtils.SetConfig(c.Request().Context(), *cfg)
	c.SetRequest(c.Request().WithContext(newCtx))

	return c
}
