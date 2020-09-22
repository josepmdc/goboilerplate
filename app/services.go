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

	if us, err := NewUserService(cfg); err != nil {
		return nil, err
	} else {
		services.User = us
	}

	if as, err := NewAuthService(services.User); err != nil {
		return nil, err
	} else {
		services.Auth = as
	}

	return services, nil
}
