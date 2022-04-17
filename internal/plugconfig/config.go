package plugconfig

// Config is plug of config for test purposes only.
// Use a real config service with secured key storage.
type Config struct {
	CoinMarketCapURL string
	CoinMarketCapKey string
}

// New constructs Config.
func New() Config {
	return Config{
		CoinMarketCapURL: "https://sandbox-api.coinmarketcap.com",
		CoinMarketCapKey: "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
	}
}
