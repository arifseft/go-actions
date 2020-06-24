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

func TestGetBiodata(t *testing.T) {
	r := gin.Default()
	app := new(app.Application)
	app.CreateApp(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/biodata", nil)
	r.ServeHTTP(w, req)

	actual := types.GetBiodataResponse{}

	if err := json.Unmarshal([]byte(w.Body.String()), &actual); err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Success get biodata", actual.Message)
	assert.Equal(t, http.StatusOK, actual.Status)
	assert.NotEmpty(t, actual.Result)
}
