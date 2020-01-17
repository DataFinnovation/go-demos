package oauth2bearer

import (
	"log"
	"net/http"
)

// AuthenticatedClient glues together a (refresable) token
// and an http client
type AuthenticatedClient struct {
	Client      *http.Client
	BearerToken *BearerToken
}

// Do wraps the http.Client Do with authentication
func (c *AuthenticatedClient) Do(req *http.Request) (*http.Response, error) {
	c.BearerToken.SetAuthHeader(req)
	return c.Client.Do(req)
}

// NewAuthenticatedClient returns a new client which is prepared
// to Do authenticated requests per the given token source
func (source *BearerTokenSource) NewAuthenticatedClient() *AuthenticatedClient {
	token, err := source.Token()
	if err != nil {
		log.Panic(err)
	}
	client := http.Client{}
	aClient := AuthenticatedClient{Client: &client, BearerToken: token}
	return &aClient
}

// SetHeaders applies each key-value pair from the map
// as req.Header.Set(key, value)
func SetHeaders(req *http.Request, keyMap map[string]string) {
	for k, v := range keyMap {
		req.Header.Set(k, v)
	}
}

// eof
