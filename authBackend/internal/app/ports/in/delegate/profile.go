package delegate

import useCases "authBackend/internal/app/usecases"

type ProfileDelegate struct {
	createProfileUCase useCases.CreateProfileUCase
	getProfileUCase    useCases.GetProfileUCase
}
