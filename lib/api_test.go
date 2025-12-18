package checkusage

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsage(t *testing.T) {
	// Arrange
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"data":{"amount_gib_per_bucket":{"is_applied":false,"quota":10240,"val":98.765432},"num_objects_per_bucket":{"is_applied":false,"quota":10000000,"val":7382}}, "_log_url":"hogehoge"}`)
	})
	testServer := httptest.NewServer(h)
	defer testServer.Close()

	cli := &APIClient{
		url: testServer.URL,
	}

	exp := &Usage{
		quota:  float64(10240),
		amount: float64(98.765432),
	}
	// Act
	act, err := cli.GetUsage("SAKURA_API_ACCESS_TOKEN", "SAKURA_API_ACCESS_TOKEN_SECERT")
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, exp.quota, act.quota, "quota should be expected")
	assert.Equal(t, exp.amount, act.amount, "amount should be expected")
}
