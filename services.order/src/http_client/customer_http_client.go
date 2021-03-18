package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/suadev/microservices/services.order/src/dto"
	"github.com/suadev/microservices/shared/config"
)

type Client struct {
	hostURL string
}

func NewCustomerClient() *Client {
	return &Client{
		hostURL: config.AppConfig.CustomerServiceEndpoint,
	}
}

func (c *Client) GetBasketItems(customerID string) ([]dto.BasketItemDto, error) {
	resp, err := http.Get(c.hostURL + fmt.Sprintf("/customers/%v/basketItems", customerID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed fetching basket items.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var basketItems []dto.BasketItemDto
	json.Unmarshal(body, &basketItems)
	return basketItems, nil
}
