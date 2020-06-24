package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arifseft/go-actions/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/name", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/name")

	if assert.NoError(t, controllers.GetName(c)) {
		expect := `{"status":200,"message":"Success get name","result":"M Arif Sefrianto"}`

		res := controllers.GetNameResponse{}
		mock := controllers.GetNameResponse{}

		if err := json.Unmarshal([]byte(rec.Body.String()), &res); err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(expect), &mock); err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mock.Message, res.Message)
		assert.Equal(t, mock.Status, res.Status)
		assert.NotZero(t, res.Result)
	}
}
