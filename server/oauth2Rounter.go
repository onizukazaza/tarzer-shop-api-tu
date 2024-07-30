package server
import (
	_oauth2Service "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/service"
	_oauth2Controller "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/controller"
	_playerRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/repository"
	_adminRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/repository"
)
func (s *echoServer) initOAuth2Router()  {
	router := s.app.Group("/v1/oauth2/google")

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service, 
		s.conf.OAuth2, 
		s.app.Logger,
	)
	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/player/login/callback", oauth2Controller.PlayerLoginCallback)
    router.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	router.DELETE("/logout", oauth2Controller.Logout)
 }