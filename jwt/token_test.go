package jwt

import (
	"testing"
)

const secretkey string = "Pa$$W00rDDDDdddd"

func TestToken(t *testing.T) {
	t.Run("Generate and Validate token", func(t *testing.T) {
		token := GenerateToken(secretkey, 1)
		if ValidateToken(secretkey, token) != nil {
			t.Errorf("Token isn't valid")
		}
	})
}

func TestNotValidToken(t *testing.T) {
	t.Run("Check invalid token", func(t *testing.T) {
		token := "1234"
		if ValidateToken(secretkey, token) == nil {
			t.Errorf("Invalid token is valid")
		}
	})

	t.Run("check expired token", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0b3B0YWwuY29tIiwiZXhwIjoxNDI2NDIwODAwLCJodHRwOi8vdG9wdGFsLmNvbS9qd3RfY2xhaW1zL2lzX2FkbWluIjp0cnVlLCJjb21wYW55IjoiVG9wdGFsIiwiYXdlc29tZSI6dHJ1ZX0.yRQYnWzskCZUxPwaQupWkiUzKELZ49eM7oWxAQK_ZXw"
		err := ValidateToken(secretkey, token)
		if err == nil {
			t.Errorf("Expired token is valid")
		}
	})
}
