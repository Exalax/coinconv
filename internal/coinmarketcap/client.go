package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//nolint gosec:G101 just a header name.
const apiKeyHeaderName = "X-CMC_PRO_API_KEY"

// httpClient is an HTTP client interface.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	client httpClient
	url    string
	key    string
}

func New(apiURL string, key string) *Client {
	return &Client{
		client: &http.Client{},
		url:    apiURL,
		key:    key,
	}
}

func (c *Client) Convert(ctx context.Context, amount float64, from, to string) (float64, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	queryURL := fmt.Sprintf(
		"%s/v2/tools/price-conversion?symbol=%s&convert=%s&amount=%f",
		c.url,
		url.QueryEscape(from),
		url.QueryEscape(to),
		amount,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrQueryConstructionFailed, err)
	}
	req.Header.Set(apiKeyHeaderName, c.key)

	rsp, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrQueryFailed, err)
	}
	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%w: %s ", ErrQueryFailed, rsp.Status)
	}

	defer func() {
		_ = rsp.Body.Close()
	}()

	dec := json.NewDecoder(rsp.Body)
	var pcResponse priceConversionResponse
	err = dec.Decode(&pcResponse)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrResponseUnmarshalFailed, err)
	}

	return pcResponse.Price(from, to)
}
