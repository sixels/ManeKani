package auth

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func StartAuthenticator() error {
	var (
		// TODO: get these values from env var
		supertokensURL    string = "http://localhost:3567"
		supertokensSecret string = "60f98a9e-ce60-48c2-bfa2-8c4f623874af"
		websiteDomain     string = "http://localhost:8082"
		apiDomain         string = "http://localhost:8081"
		apiBasePath       string = "/auth"
		websiteBasePath   string = "/auth"
	)

	supertokensConfig := supertokens.ConnectionInfo{
		ConnectionURI: supertokensURL,
		APIKey:        supertokensSecret,
	}
	supertokensAppInfo := supertokens.AppInfo{
		AppName:         "manekani",
		APIDomain:       apiDomain,
		APIBasePath:     &apiBasePath,
		WebsiteDomain:   websiteDomain,
		WebsiteBasePath: &websiteBasePath,
	}

	if err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokensConfig,
		AppInfo:     supertokensAppInfo,
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	}); err != nil {
		return err
	}

	return nil
}

func (a *Authenticator) GetUserInfo(ctx context.Context, token *oauth2.Token) (*oidc.UserInfo, error) {
	return a.UserInfo(ctx, oauth2.ReuseTokenSource(token, a.TokenSource(ctx, token)))
}
