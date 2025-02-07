package keycloak

import (
	"authServer/pkg/utilities"
	"errors"
	"fmt"
	"net/http"
)

var (
	Internal    = errors.New("KeycloakClient.Internal error.")
	Conflict409 = errors.New("KeycloakClient.conflict 409 error.")
)

func (client KeycloakClient) catchHttpStatus(resp *http.Response) error {
	if resp != nil {
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusNoContent {
			return nil
		} else {
			keycloakErr := &KeycloakResponseError{}
			utilities.ParseResponseToStruct(resp, keycloakErr)
			msg := fmt.Sprintf("Wrap error: %s", client.getErrorMSG(keycloakErr))

			if resp.StatusCode == http.StatusConflict {
				return errors.Join(Conflict409, errors.New(keycloakErr.ErrorMessage))
			}
			return errors.Join(Internal, errors.New(msg))
		}
	}
	return nil
}

func (client KeycloakClient) getErrorMSG(resp *KeycloakResponseError) string {
	if resp.Error != "" && resp.ErrorDescription != "" {
		return resp.Error + " " + resp.ErrorDescription
	} else {
		return resp.ErrorMessage
	}
}
