package api

import "net/http"

func (api *API) RedisServerHealthCheck() error {
	url, err := RedisServer(api.Config.Url.RedisServer).HealthCheck()
	if err != nil {
		return err
	}

	return api.Do(http.MethodGet, nil, nil, url, nil)
}
