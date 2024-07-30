package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/pkg/custom"
	_oauth2Exception "github.com/onizukazaza/tarzer-shop-api-tu/pkg/oauth2/exception"
	"golang.org/x/oauth2"
	"net/http"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(pctx echo.Context , next echo.HandlerFunc) error {
	ctx := context.Background()


	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.playerTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := playerGoogleOAuth2.Client(ctx , tokenSource)

	userInfo , err := c.getUserInfo(client)
	if err!= nil {
        return custom.Error(pctx, http.StatusUnauthorized, err)
    }

	if c.oauth2Service.IsThisGuyReallyPlayer(userInfo.ID){
		return custom.Error(pctx , http.StatusUnauthorized, &_oauth2Exception.Uauthorized{})
	}

	pctx.Set("playerID" , userInfo.ID) //set player id handler

	return next(pctx)  //middleware sent to handler use next 
}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context , next echo.HandlerFunc) error {
	ctx := context.Background()


	tokenSource, err := c.getTokenSource(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.adminTokenRefreshing(pctx, tokenSource)
		if err != nil {
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	client := adminGoogleOAuth2.Client(ctx , tokenSource)

	userInfo , err := c.getUserInfo(client)
	if err!= nil {
        return custom.Error(pctx, http.StatusUnauthorized, err)
    }

	if c.oauth2Service.IsThisGuyReallyAdmin(userInfo.ID){
		return custom.Error(pctx , http.StatusUnauthorized, &_oauth2Exception.Uauthorized{})
	}

	pctx.Set("adminID" , userInfo.ID) //set player id handler

	return next(pctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updateToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Uauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updateToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updateToken.RefreshToken)

	return updateToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updateToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &_oauth2Exception.Uauthorized{}
	}

	c.setSameSiteCookie(pctx, accessTokenCookieName, updateToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updateToken.RefreshToken)

	return updateToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(pctx echo.Context) (*oauth2.Token, error) { //get the token source form cookie
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Uauthorized{}
	}

	refreshToken, err := pctx.Cookie(refreshTokenCookieName)
	if err != nil {
		return nil, &_oauth2Exception.Uauthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
