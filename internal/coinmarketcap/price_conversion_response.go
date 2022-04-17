package coinmarketcap

import "fmt"

// priceConversionResponse is a price conversion query response.
type priceConversionResponse struct {
	Status statusResponse `json:"status"`
	Data   map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

// Price returns price of "from" in "to" units.
func (r priceConversionResponse) Price(from, to string) (float64, error) {
	err := r.Status.Err()
	if err != nil {
		return 0, err
	}

	quote, ok := r.Data[from]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrSymbolNotFound, from)
	}

	price, ok := quote.Quote[to]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrSymbolNotFound, to)
	}

	return price.Price, nil
}
