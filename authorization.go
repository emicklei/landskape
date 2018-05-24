package main

import (
	"os"

	restful "github.com/emicklei/go-restful"
)

var apiKey = os.Getenv("API_KEY")

func apiKeyAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	key := req.QueryParameter("api_key")
	if apiKey != key {
		resp.WriteErrorString(401, "401: Not Authorized (invalid or missing api_key query parameter)")
		return
	}
	chain.ProcessFilter(req, resp)
}
