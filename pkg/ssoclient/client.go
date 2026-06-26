package ssoclient

import (
	"net/http"
	"time"

	"buf.build/gen/go/zik-zikurrat-sso/sso/connectrpc/go/sso/v1/ssov1connect"
)

func New(baseURL string) ssov1connect.RegistryServiceClient {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	return ssov1connect.NewRegistryServiceClient(httpClient, baseURL)
}
