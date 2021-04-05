module github.com/lovethedrake/brigdrake

go 1.15

replace github.com/mholt/caddy => github.com/caddyserver/caddy/v2 v2.3.0

replace github.com/lovethedrake/drakecore => ../drakecore

require (
	github.com/brigadecore/brigade/sdk/v2 v2.0.0-alpha.1
	github.com/carolynvs/magex v0.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/go-github/v18 v18.2.0
	github.com/lovethedrake/drakecore v0.14.0
	github.com/magefile/mage v1.11.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	k8s.io/api v0.19.9
	k8s.io/apimachinery v0.19.9
	k8s.io/client-go v0.19.9
)
