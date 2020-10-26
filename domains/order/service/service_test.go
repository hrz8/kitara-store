package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"github.com/hrz8/kitara-store/models"
)

const (
	url         = "http://127.0.0.1:8090/api/v1/orders"
	prodNomerSt = "35731de0-e646-4379-a0f1-b69e74742e0a"
	prodNomerTg = "0bde6df2-f505-401f-882c-808855c2871d"
)

func TestOrderMultiple(t *testing.T) {
	// task 1
	products1 := make([]models.CreateOrdersProductsPayload, 0)
	products1 = append(products1, models.CreateOrdersProductsPayload{
		ID:  uuid.FromStringOrNil(prodNomerSt),
		Qty: 3,
	})
	products1 = append(products1, models.CreateOrdersProductsPayload{
		ID:  uuid.FromStringOrNil(prodNomerTg),
		Qty: 3,
	})
	pl1, err := json.Marshal(models.CreateOrderPayload{
		Products: products1,
	})
	assert.Nil(t, err)
	body1 := strings.NewReader(string(pl1))
	rsp1 := make(chan *http.Response)

	// task 2
	products2 := make([]models.CreateOrdersProductsPayload, 0)
	products2 = append(products2, models.CreateOrdersProductsPayload{
		ID:  uuid.FromStringOrNil(prodNomerSt),
		Qty: 16,
	})
	pl2, err := json.Marshal(models.CreateOrderPayload{
		Products: products2,
	})
	assert.Nil(t, err)
	body2 := strings.NewReader(string(pl2))
	rsp2 := make(chan *http.Response)

	// run tasks concurrently
	go func(t *testing.T) {
		r1, err := http.Post(url, "application/json", body1)
		assert.Nil(t, err)
		rsp1 <- r1
	}(t)

	go func(t *testing.T) {
		r2, err := http.Post(url, "application/json", body2)
		assert.Nil(t, err)
		rsp2 <- r2
	}(t)

	response1 := <-rsp1
	response2 := <-rsp2

	rd1, err := ioutil.ReadAll(response1.Body)
	rd2, err := ioutil.ReadAll(response2.Body)

	s1 := gjson.Get(string(rd1), "status")
	s2 := gjson.Get(string(rd2), "status")

	// if payload1 is OK payload 2 must be fail
	// if payload1 is fail payload 2 must be OK
	assert.NotEqual(t, s1.Int(), s2.Int())
}
