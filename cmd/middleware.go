package cmd

import (
	"net/http"
	"strings"

	"github.com/versus/gouploadservice/jwt"
	"gopkg.in/macaron.v1"
)

//bearerValidate check Authorization: Bearer header for prometeus metrics
func bearerValidate(ctx *macaron.Context) {
	if strings.Contains(ctx.Req.Header.Get("Authorization"), "Bearer") {
		s := strings.Split(ctx.Req.Header.Get("Authorization"), " ")
		if len(s) == 2 {
			if err := jwt.ValidateToken(secretKey, s[1]); err == nil {
				return
			}
			http.Error(ctx.Resp, "Error: Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(ctx.Resp, "Error: Unauthorized", http.StatusUnauthorized)
		return
	}
	http.Error(ctx.Resp, "No Fish Here!", http.StatusMethodNotAllowed)
	return
}
