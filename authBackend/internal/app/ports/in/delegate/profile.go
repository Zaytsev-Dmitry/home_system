package delegate

import useCases "authServer/internal/app/usecases"

type ProfileDelegate struct {
	createProfileUCase useCases.CreateProfileUCase
	getProfileUCase    useCases.GetProfileUCase
}
