package coinmarketcap

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Convert(t *testing.T) {
	ctx := context.Background()
	apiKey := "123-abc"

	setupHttpClientDo := func(client *mockHttpClient, url string, status int, body string, rspErr error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		require.NoError(t, err)
		req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

		client.On("Do", req).Return(
			&http.Response{
				Status:     http.StatusText(status),
				StatusCode: status,
				Body:       io.NopCloser(bytes.NewReader([]byte(body))),
			},
			rspErr,
		)
	}

	type mocks struct {
		client *mockHttpClient
	}
	type fields struct {
		url string
		key string
	}
	type args struct {
		ctx    context.Context
		amount float64
		from   string
		to     string
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		setup   func() mocks
		want    float64
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "BTC",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=BTC&convert=USD&amount=5.280000",
					http.StatusOK,
					`
					{
						"status": {
							"timestamp": "2022-04-17T09:36:36.502Z",
							"error_code": 0,
							"error_message": null,
							"elapsed": 0,
							"credit_count": 1,
							"notice": null
					   },
						"data": {
							"BTC": {
								"symbol": "BTC",
								"id": "sfga77eblig",
								"name": "c8l868u427p",
								"amount": 5,
								"last_updated": "2022-04-17T09:36:36.502Z",
								"quote": {
									"USD": {
										"price": 6580.25,
										"last_updated": "2022-04-17T09:36:36.502Z"
									}
								}
							}
						}
					}`,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    6580.25,
			wantErr: nil,
		},
		{
			name: "success_low_case",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "btc",
				to:     "usd",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=BTC&convert=USD&amount=5.280000",
					http.StatusOK,
					`
					{
						"status": {
							"timestamp": "2022-04-17T09:36:36.502Z",
							"error_code": 0,
							"error_message": null,
							"elapsed": 0,
							"credit_count": 1,
							"notice": null
					   },
						"data": {
							"BTC": {
								"symbol": "BTC",
								"id": "sfga77eblig",
								"name": "c8l868u427p",
								"amount": 5,
								"last_updated": "2022-04-17T09:36:36.502Z",
								"quote": {
									"USD": {
										"price": 6580.25,
										"last_updated": "2022-04-17T09:36:36.502Z"
									}
								}
							}
						}
					}`,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    6580.25,
			wantErr: nil,
		},
		{
			name: "fail_no_from",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "ETH",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=ETH&convert=USD&amount=5.280000",
					http.StatusOK,
					`
					{
						"status": {
							"timestamp": "2022-04-17T09:36:36.502Z",
							"error_code": 0,
							"error_message": null,
							"elapsed": 0,
							"credit_count": 1,
							"notice": null
					   },
						"data": {
							"BTC": {
								"symbol": "BTC",
								"id": "sfga77eblig",
								"name": "c8l868u427p",
								"amount": 5,
								"last_updated": "2022-04-17T09:36:36.502Z",
								"quote": {
									"USD": {
										"price": 6580.25,
										"last_updated": "2022-04-17T09:36:36.502Z"
									}
								}
							}
						}
					}`,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    0,
			wantErr: ErrSymbolNotFound,
		},
		{
			name: "fail_no_to",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "ETH",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=ETH&convert=USD&amount=5.280000",
					http.StatusOK,
					`
					{
						"status": {
							"timestamp": "2022-04-17T09:36:36.502Z",
							"error_code": 0,
							"error_message": null,
							"elapsed": 0,
							"credit_count": 1,
							"notice": null
					   },
						"data": {
							"BTC": {
								"symbol": "BTC",
								"id": "sfga77eblig",
								"name": "c8l868u427p",
								"amount": 5,
								"last_updated": "2022-04-17T09:36:36.502Z",
								"quote": {
									"USD": {
										"price": 6580.25,
										"last_updated": "2022-04-17T09:36:36.502Z"
									}
								}
							}
						}
					}`,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    0,
			wantErr: ErrSymbolNotFound,
		},
		{
			name: "fail_broken_body",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "ETH",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=ETH&convert=USD&amount=5.280000",
					http.StatusOK,
					`<?xml?>`,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    0,
			wantErr: ErrResponseUnmarshalFailed,
		},
		{
			name: "fail_query_error",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "ETH",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=ETH&convert=USD&amount=5.280000",
					0,
					``,
					errors.New("some error"),
				)

				return mocks{
					client: client,
				}
			},
			want:    0,
			wantErr: ErrQueryFailed,
		},
		{
			name: "fail_unwanted_http_status",
			args: args{
				ctx:    ctx,
				amount: 5.28,
				from:   "BTC",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				client := &mockHttpClient{}
				setupHttpClientDo(
					client,
					"https://example.com/v2/tools/price-conversion?symbol=BTC&convert=USD&amount=5.280000",
					http.StatusInternalServerError,
					``,
					nil,
				)

				return mocks{
					client: client,
				}
			},
			want:    0,
			wantErr: ErrQueryFailed,
		},
		{
			name: "fail_construct_query",
			args: args{
				ctx:    nil,
				amount: 5.28,
				from:   "BTC",
				to:     "USD",
			},
			fields: fields{
				url: "https://example.com",
				key: apiKey,
			},
			setup: func() mocks {
				return mocks{
					client: &mockHttpClient{},
				}
			},
			want:    0,
			wantErr: ErrQueryConstructionFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := tt.setup()
			c := &Client{
				client: mocks.client,
				url:    tt.fields.url,
				key:    tt.fields.key,
			}

			got, err := c.Convert(tt.args.ctx, tt.args.amount, tt.args.from, tt.args.to)

			mocks.client.AssertExpectations(t)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
