package keycloak

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	apikitErr "github.com/Zaytsev-Dmitry/apikit/custom_errors"
	generatedApi "userService/api/http"
	"userService/pkg/config_loader"
)

type KeycloakClient struct {
	config *config_loader.AppConfig
	client *gocloak.GoCloak
}

func NewKeycloakClient(config *config_loader.AppConfig) *KeycloakClient {
	return &KeycloakClient{
		config: config,
		client: gocloak.NewClient(config.Keycloak.Host),
	}
}

func (k *KeycloakClient) getToken(ctx context.Context) (*gocloak.JWT, error) {
	return k.client.LoginClient(
		ctx,
		k.config.Keycloak.ClientId,
		k.config.Keycloak.ClientSecret,
		k.config.Keycloak.Realm,
	)
}

func (k *KeycloakClient) Introspect(ctx context.Context, config *config_loader.AppConfig, token string) (*gocloak.IntroSpectTokenResult, error) {
	return k.client.RetrospectToken(
		ctx,
		token,
		config.Keycloak.ClientId,
		config.Keycloak.ClientSecret,
		config.Keycloak.Realm,
	)
}

func (k *KeycloakClient) CreateUser(req generatedApi.CreateAccountRequest) (*gocloak.User, error) {
	ctx := context.Background()

	// Получаем токен
	token, err := k.getToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	// Создаём пользователя
	userID, err := k.client.CreateUser(ctx, token.AccessToken, k.config.Keycloak.Realm, gocloak.User{
		Username: gocloak.StringP(*req.Username),
		Email:    gocloak.StringP(*req.Email),
		Enabled:  gocloak.BoolP(true),
	})
	if err != nil {
		// Если пользователь уже существует (409)
		if apiErr, ok := err.(*gocloak.APIError); ok && apiErr.Code == 409 {
			users, _ := k.client.GetUsers(ctx, token.AccessToken, k.config.Keycloak.Realm, gocloak.GetUsersParams{
				Email:    gocloak.StringP(*req.Email),
				Username: gocloak.StringP(*req.Username),
			})
			if len(users) > 0 {
				return users[0], nil
			}
			return nil, fmt.Errorf("user exists but could not retrieve")
		}
		if apiErr, ok := err.(*gocloak.APIError); ok && apiErr.Code == 403 {
			return nil, apikitErr.ForbiddenError
		}
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	return k.client.GetUserByID(ctx, token.AccessToken, k.config.Keycloak.Realm, userID)
}
