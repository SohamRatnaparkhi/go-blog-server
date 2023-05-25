package server

import (
	"net/http"

	"github.com/SohamRatnaparkhi/go-blog-server/utils"
)

func HealthCheck(res http.ResponseWriter, req *http.Request) {
	type resp struct {
		Status string `json:"status"`
	}
	utils.ResponseJson(res, 200, resp{
		Status: "ok",
	})
}
