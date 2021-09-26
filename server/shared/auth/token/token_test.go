package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"os"
	"testing"
)

func TestJWTTokenVerifier_Verify(t *testing.T) {

	pkFile, err := os.Open("public.key")
	if err != nil {
		panic(err)
	}

	pkBytes, err := io.ReadAll(pkFile)

	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pkBytes)

	jwtTokenVerifier := JWTTokenVerifier{PublicKey: key}
	fmt.Println(jwtTokenVerifier)

	verify, err := jwtTokenVerifier.Verify("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI2NDIxMzksImlhdCI6MTYzMjYzNjEzOSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjE0ZmU3YjM3MDIxYTkzZTRkNmU4ZmU4In0.tUAaH-hHd8Ol4Mofeyu3PN_vGnSfWzyqMXaLf4vtyywJvtQ5idAk3nDJIdIU6MB81wpbDGv5mze7HvSz1CJ7BqwqFsEQ2UCFRSYIAdigI0lHjhkIzzb1mgrdYYhWBM0qPzQT4pMU742vV0zHJ_wckJWfOhrHRd1qiPy3U4NJpXJZ4vUJ3Hs3q_RDNdtA1ZiZp0fGqhZRHsJ2yaq25RvV-QcXBCCvI2WHkaKcEjcgPx2ci-3GiXMQlc5QMuFVB-qsyFEBUSizuzrZTqyl42oxhzOivWbssEc9XQP-aWN0wyxNUv50jTuPJvPQDstEeOVXqre04wwuV6K0nN64rys3Yg")

	if err != nil {
		panic(err)
	}

	fmt.Println(verify)
}
