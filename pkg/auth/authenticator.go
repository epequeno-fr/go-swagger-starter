package auth

import (
	"example.ponies.com/api/pkg/config"
	"example.ponies.com/api/swagger_gen/models"
)

func Authenticate(token string, scopes []string) (*models.Principal, error) {
	config.Logger.Warn("Authentication disabled: Uncomment access token introspection in authenticator.go")
	// it, err := am.IntrospectToken(pester.New(), config.Env.AmUrl, config.Env.OAuthClientID,
	// 	config.Env.OAuthClientSecret,
	// 	token)

	// if err != nil {
	// 	config.Logger.WithError(err).Error("failed to introspect token")
	// 	return nil, errors.New(401, "Unauthorized")
	// }

	// if !it.Active {
	// 	return nil, errors.New(401, "Unauthorized")
	// }

	// for _, scope := range scopes {
	// 	if !util.Contains(strings.Split(it.Scope, " "), scope) {
	// 		return nil, errors.New(403, "Unauthorized - Insufficient Scope")
	// 	}
	// }

	// return &models.Principal{
	// 	Sub:       it.Sub,
	// 	Realm:     it.Realm,
	// 	AuthLevel: int64(it.AuthLevel), // go-swagger only allows int64...
	// }, nil

	// Remove this after uncommenting the AM integration above
	return &models.Principal{
		Sub:       "foo",
		Realm:     "/alpha",
		AuthLevel: 0,
	}, nil
}
