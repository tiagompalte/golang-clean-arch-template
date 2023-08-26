package server

import (
	_ "github.com/tiagompalte/golang-clean-arch-template/api"
	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/routes"
)

// @title						TODO API
// @version						1.0
// @description					TODO API
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url					http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url					http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath					/
// @schemes						https
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
func NewServer(app application.App) error {
	app.Server().RegisterGroupRoutes(routes.CreateRoute(app))
	return app.Server().Start()
}
