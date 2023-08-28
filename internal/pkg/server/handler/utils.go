package handler

import (
	"net/http"
	"strings"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/usecase"
)

var blank = usecase.Blank{}

func extractParamPath(r *http.Request, nrParam int) (string, bool) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) < nrParam {
		return "", false
	}
	return params[nrParam], true
}
