package service

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

type ExternalUrlService struct {
	urlS string
}

func NewExternalUrlService(urlS string) *ExternalUrlService {
	return &ExternalUrlService{
		urlS: urlS,
	}
}

func (es *ExternalUrlService) GetLatency() (int, error) {
	var (
		result httpstat.Result
	)
	// Create a new HTTP request
	req, err := http.NewRequest("GET", es.urlS, nil)
	if err != nil {
		return 0, err
	}

	// Create a httpstat powered context
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	if _, err := io.Copy(io.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	_ = time.Now()

	tcp := int(result.TCPConnection / time.Millisecond)

	return tcp, nil
}
