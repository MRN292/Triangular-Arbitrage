package kraken

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"triangular-arbitrage/config"
	"triangular-arbitrage/internal/domain"
)

type Kraken struct {
	cfg        *config.Config
	httpClient *http.Client
}

func NewKraken(cfg *config.Config) *Kraken {
	return &Kraken{

		cfg:        cfg,
		httpClient: &http.Client{},
	}

}

func (k *Kraken) Name() string {
	return "kraken"
}

type krakenPingResponse struct {
	Error  []string `json:"error"`
	Result struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
	} `json:"result"`
}

func (b *Kraken) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.cfg.APIConfig.KrakenPingUrl, nil)
	if err != nil {
		return err
	}

	response, err := b.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("ping %s failed status: %d", b.Name(), response.StatusCode)
	}

	krknPingRes := new(krakenPingResponse)

	if err := json.NewDecoder(response.Body).Decode(krknPingRes); err != nil {
		return err
	}
	if len(krknPingRes.Error) != 0 {
		return fmt.Errorf("Error in ping kraken %v", krknPingRes.Error)
	}

	if krknPingRes.Result.Status != "online" {
		return fmt.Errorf("there are some limitation in pinging, ping result ; %s", krknPingRes.Result.Status)
	}

	return nil
}

func (b *Kraken) GetOrderBook(ctx context.Context, symbols []string) (*domain.OrderBooks, error) {

	output := &domain.OrderBooks{
		Markets: make(map[string]*domain.TopOfBook),
	}

	for _, symbol := range symbols {
		topOfBook, err := b.getTopOfBook(ctx, symbol)
		if err != nil {
			return &domain.OrderBooks{}, err
		}

		output.Markets[symbol] = topOfBook
	}

	return output, nil
}

type KrakenOrderBookDTO struct {
	Error  []any                      `json:"error"`
	Result map[string]KrakenMarketDTO `json:"result"`
}

type KrakenMarketDTO struct {
	Asks [][2]string `json:"asks"`
	Bids [][2]string `json:"bids"`
}

func (b *Kraken) getTopOfBook(ctx context.Context, symbol string) (output *domain.TopOfBook, err error) {
	url, err := url.Parse(b.cfg.APIConfig.KrakenOrderBookUrl)
	if err != nil {
		return output, err
	}

	query := url.Query()
	query.Set("pair", symbol)
	query.Set("count", "1") // getting 1 means top
	url.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return output, err
	}

	response, err := b.httpClient.Do(request)
	if err != nil {
		return output, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return output, fmt.Errorf("orderbook failed for %s code %d", symbol, response.StatusCode)
	}

	orderBookResponse := new(KrakenOrderBookDTO)
	if err := json.NewDecoder(response.Body).Decode(orderBookResponse); err != nil {
		return output, err
	}

	output, err = orderBookResponse.ToDomain()

	return output, nil
}
