package controllers

import (
	"net/http"

	"gitlab.com/syaifuddin.teddy/test-case-majoo/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Teddy's API")
}
