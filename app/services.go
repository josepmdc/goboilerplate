package app

import "github.com/josepmdc/goboilerplate/conf"

// Services is a struct containing all of the services nedded by the API
type Services struct {
	User UserService
	Auth AuthService
}

// InitServices returns a Services struct with all the needed services initialized
func InitServices(cfg *conf.Config) (*Services, error) {
	services := &Services{}

	us, err := NewUserService(cfg)
	if err != nil {
		return nil, err
	}
	services.User = us

	as, err := NewAuthService(services.User)
	if err != nil {
		return nil, err
	}
	services.Auth = as

	return services, nil
}
