package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/onizukazaza/tarzer-shop-api-tu/config"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
	"golang.org/x/net/context"
	_oauth2Service "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/service"
	_oauth2Controller "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/controller"
	_playerRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/repository"
	_adminRepository "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/repository"

	
)

type echoServer struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})
	return server
}

func (s *echoServer) Start() {
	corsMiddleware :=  getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.TimeOut)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeOutMiddleware)

	authorizingMiddleware := s.getAuthorizingMiddleware()

	s.app.GET("/v1/health", s.healthCheck)

    s.initOAuth2Router()
	s.initItemManagingRouter(authorizingMiddleware) //api route set midleware 
	s.initItemShopRouter()
	
	quitCh := make(chan os.Signal, 1) //stop server fully
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening()

}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}

}

func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal) { //quitechanel interrup "Ctrl+C"
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("Shutting down server....")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK Success")
}
// func getLoggerMiddleware() echo.MiddlewareFunc {
// 	return middleware.Logger()
// }

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout: timeout *time.Second,
	})
}
func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc{
return middleware.CORSWithConfig(middleware.CORSConfig{
	Skipper: middleware.DefaultSkipper,
	AllowOrigins: allowOrigins,
	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
})
}

func getBodyLimitMiddleware(bodyLimit string)echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}


func (s *echoServer) getAuthorizingMiddleware() *authorizingMiddleware  {

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)
	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository)
	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service, 
		s.conf.OAuth2, 
		s.app.Logger,
	)

	return &authorizingMiddleware{
		oauth2Controller: oauth2Controller, 
		oauth2Conf: s.conf.OAuth2, 
		logger:  s.app.Logger,
	}
 }