package app

import "github.com/josepmdc/goboilerplate/conf"

// Services is a struct containing all of the services nedded by the API
type Services struct {
	User UserService
}

// InitServices returns a Services struct with all the needed services initialized
func InitServices(cfg *conf.Config) (*Services, error) {
	services := &Services{}

	if us, err := NewUserService(cfg); err != nil {
		return nil, err
	} else {
		services.User = us
	}

	return services, nil
}
