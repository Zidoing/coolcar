package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

const privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDWJWUosFTtiTvu
/FIBqqmvhCbls21z0CuIoe6bKiaFq7I+GuHW0+yOCPx//q1/cGPf1psCbgZ3xlqc
uK0mlP84+QkLXdmccpcHxSDsye5aR5dACx60pntoN1LKt0iiNVSeHKxISu9VFZ1r
xhA5wP6dN/Sx09mvCr0Hq0d0LvQkBVsr3kC/tSjO8HkzhnlJAhSwq7GW/nWSfRhS
tphpH//akmMP51EFFFt/OwBWVfXi0bLtcrw2sm1m9UzDaUUig3ZyKbrynXjw+Oyb
f+TmMFmxCrgdwN+uEMSCXYIx77ywQlMAr03yYb1TAryEYWRJS88YTrz3hFIHWZ1K
WqNnNwfvAgMBAAECggEAbN5aDVOH+bL5G0h/0IUY+zguJj0P+f7tAuuASLYnG3Hs
jhJQmkznSFZdOB0by1wyCkwqRhP1idL3ohFpReId1rMU3Ah8Z0MCcVrQoFEQsikP
RenhVGoG9zaZCqIfdQez2yYsmBwwjIsOmwn4mcP6LaeS6Q0Q8CJ/c3SqQPDWAbEg
um0f5BR58dNSS/fUhMnCHZDyq2iEibCE6Q+8wQROyvmrb7FqHET7eHDhUeF8j0fy
XdG6dEQG1wh7N3nkzmQqXXo0zsP9JhSta+EK+KeLL15Le+KL6mcpB9IpsHsiRFmy
MgJ9gqAogiIVb3EjDA0t7ZBX81CuDxFqs9aQbnSsMQKBgQD7aUuZMI/i9hStYNx3
Y3MEtC2wHduSbCS066hudeLWZr+QB5HFWyK2eNR6mdSWyQJWYxzrYYgoZTfyAYdp
D36xtLIK+D3nl4fMfYxWrEz4LojWooLe7vEl+6wE74hW7Ik5jPynmXcOnNo4LYdv
nBfOj27szUDXumsrjGQszblYhQKBgQDaDfru2mUzHMGEfpMTdN636E6cpakyJnm3
AKzIv2XLPhDjhry0yd5IvB0GB7iRDaKpUOhif0ZeD/Wch1seeSQE8kuNDnuRNeYy
E6rRE0BeRpRsGzqAMZmTn5QsNWG0uXk23eT3G6qIF4T74iJUtiz+r1YBR+bdfzN9
hNlXMmKC4wKBgBJjBAVjGNLXfnGjqNwOaOYKK2E6/cj6ocCdj4bsljdp8k6dTxro
7GP9+EQzJeoVL1eUDhRk4KlbpbjMRdP4LyLd/TtyJkzety7Ma8wW89YeySExZ3LY
dMh7XHxL7GO3Dt4non67aYqnDAqZwdL2zoLr2If3BuwAFUBtFxJzxT9lAoGAdZRy
Shez2DLaWuKR1jz/17VJpohhyuwaV8biZMFoV4bZPocp/GaaGShukhDZwkXS/2/k
TzCINjJu43/Nb4otFQm+GYrEsNGalkgqpOC3pd5zLDQmgHe+c27qTEhgrj4REbfn
PiRJ+WaNS/FtR7aQXqkXPEuUPr3XewiG1dESc7cCgYBs9Gxd17JOKN56TV3nvjLc
tUNHmk9xsWeU4S+03tJjAzlNv9p+feHubdLRIFembOXIQVVsTacp+4ZHyXO7Qvsi
13+uqtTLsK5LjOlf50uPdlZEFjk6GO3zK/GzKj6xTsUBzUnz2WFW1WUWV5MC11ve
7tVOLUfDhPw1G/OconsU6Q==
-----END PRIVATE KEY-----`

func TestNewJWTTokenGen(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		panic(err)
	}
	jwtTokenGen := NewJWTTokenGen("coolcar/auth", key)
	fmt.Println(jwtTokenGen.GenerateToken("123", 2000))
}
