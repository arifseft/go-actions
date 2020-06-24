package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arifseft/go-actions/app"
	"github.com/arifseft/go-actions/global/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/name", nil)
	r.ServeHTTP(w, req)

	e := `{"status":200,"message":"Success get name","result": "M Arif Sefrianto"}`

	actual := types.GetNameResponse{}
	expect := types.GetNameResponse{}

	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(e), &expect); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expect.Message, actual.Message)
	assert.Equal(t, expect.Status, actual.Status)
	assert.Equal(t, expect.Result, actual.Result)
}
