package controller

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/config"
	"github.com/onizukazaza/tarzer-shop-api-tu/pkg/custom"
	_oauth2Exception "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/exception"
	_oauth2Model "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/model"
	_oauth2Service "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/service"
	"golang.org/x/oauth2"
	_adminModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/model"
	_playerModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/model"
)
type googleOAuth2Controller struct {
	oauth2Service _oauth2Service.OAuth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rft"
	stateCookieName        = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger,
) OAuth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})
	return &googleOAuth2Controller{
		oauth2Service,
		oauth2Conf,
		logger,
	}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	state := c.randomState()

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state)) //return redirect  302
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	state := c.randomState() //check domain

	c.setCookie(pctx, stateCookieName, state)

	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state)) //return redirect to google  302
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error { //state set not in time
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	token, err := playerGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Uauthorized{}) //return 401 unauthorized respone err
	}

	client := playerGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.Uauthorized{}) //return 500 internal
	}

	playerCreatingReq := &_playerModel.PlayerCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingReq); err != nil {
		c.logger.Errorf("Failed to create admin account: %s", err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginRespone{Message: "Login success"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := retry.Do(func() error { //state set not in time
		return c.callbackValidating(pctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	token, err := adminGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2Exception.Uauthorized{}) //return 401 unauthorized respone err
	}

	client := adminGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.Uauthorized{}) //return 500 internal
	}

	adminCreatingReq := &_adminModel.AdminCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(adminCreatingReq); err != nil {
		c.logger.Errorf("Failed to create admin account: %s", err.Error())
		return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)


	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginRespone{Message: "Login success"})
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
	accessToken , err := pctx.Cookie(accessTokenCookieName)
	if err!= nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
        return custom.Error(pctx, http.StatusBadRequest, &_oauth2Exception.Logout{})
    }

	if err := c.revokeToken(accessToken.Value); err!= nil {
        // c.logger.Errorf("Failed to revoke token: %s", err)
        c.logger.Errorf("Error revoking token: %s", err.Error())
        return custom.Error(pctx, http.StatusInternalServerError, &_oauth2Exception.Logout{})
    }

	c.removeSameSiteCookie(pctx , accessTokenCookieName)
	c.removeSameSiteCookie(pctx , refreshTokenCookieName)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LogoutRespone{Message: "Logout success"})

}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s" , c.oauth2Conf.RevokeUrl, accessToken)

	resp, err := http.Post(revokeURL , "application/x-www-form-urlencoded", nil)
	if err!= nil {
        // c.logger.Errorf("Failed to revoke token: %s", err)
		fmt.Println("Error revoking token:", err)
        return err
    }

	defer resp.Body.Close()

	return nil
}
func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true, //fix java script not reading
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true, //fix java script not reading
		MaxAge:   -1,   //time limit to delete cookie
	}
	pctx.SetCookie(cookie)
}


func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true, //fix java script not reading
		SameSite: http.SameSiteStrictMode, //enable same site cookies
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true, //fix java script not reading
		MaxAge:   -1,   //time limit to delete cookie
		SameSite: http.SameSiteStrictMode, //enable same site cookies
		// Secure:   true, //only send cookie over HTTPS
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (c *googleOAuth2Controller) callbackValidating(pctx echo.Context) error {
	state := pctx.QueryParam("state")

	stateFormCookie, err := pctx.Cookie(stateCookieName)
	if err != nil {
		c.logger.Error("State cookie not found: %s", err)
		return &_oauth2Exception.Uauthorized{}
	}

	if state != stateFormCookie.Value {
		c.logger.Error("Invalid state: %s", state)
		return &_oauth2Exception.Uauthorized{}
	}

	c.removeCookie(pctx, stateCookieName) //delete state cookie to prevent CSRF attack

	return nil
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.UserInfo, error) {
	resp, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	userInfoInBytes, err := io.ReadAll(resp.Body) // come respon sent read string
	if err != nil {
		c.logger.Errorf("Failed to read user info response: %s", err.Error())
		return nil, err
	}

	userInfo := new(_oauth2Model.UserInfo)
	if err := json.Unmarshal(userInfoInBytes, userInfo); err != nil {
		c.logger.Errorf("Failed to unmarshal user info: %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}
