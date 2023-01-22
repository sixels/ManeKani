package auth

import (
	"context"
	"log"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/services/users"
)

func StartAuthenticator(users *users.UsersRepository) error {
	var (
		// TODO: get these values from env var
		supertokensURL    string = "http://0.0.0.0:3567"
		supertokensSecret string = "60f98a9e-ce60-48c2-bfa2-8c4f623874af"
		websiteDomain     string = "http://0.0.0.0:8082"
		apiDomain         string = "http://0.0.0.0:8081"
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
			emailpassword.Init(&epmodels.TypeInput{
				SignUpFeature: &epmodels.TypeInputSignUp{
					FormFields: []epmodels.TypeInputFormField{
						{ID: "username", Validate: func(value interface{}) *string {
							username := value.(string)

							if !ValidateUsername(username) {
								message := "Username is invalid"
								return &message
							}
							if isAvailable, err := users.UsernameAvailable(context.Background(), username); !isAvailable || err != nil {
								var message string
								if err != nil {
									log.Println(err)
									message = "Could not check username availability. Try again later"
								} else {
									message = "Username already taken"
								}
								return &message
							}

							return nil
						}},
					},
				},
				Override: &epmodels.OverrideStruct{
					APIs: func(originalImplementation epmodels.APIInterface) epmodels.APIInterface {
						signUpPOST := *originalImplementation.SignUpPOST
						(*originalImplementation.SignUpPOST) = func(formFields []epmodels.TypeFormField, options epmodels.APIOptions, userContext supertokens.UserContext) (epmodels.SignUpPOSTResponse, error) {
							ctx := options.Req.Context()

							// Override sign up
							response, err := signUpPOST(formFields, options, userContext)
							if err != nil {
								return epmodels.SignUpPOSTResponse{}, err
							}
							if response.OK != nil {
								created, err := signUpHook(ctx, users, response.OK.User, formFields)
								if err != nil {
									return response, err
								}
								log.Printf("created user %s\n", created.Email)
							}
							return response, nil
						}

						return originalImplementation
					},
				},
			}),
			session.Init(nil),
			dashboard.Init(dashboardmodels.TypeInput{
				ApiKey: supertokensSecret,
			}),
		},
	}); err != nil {
		return err
	}

	return nil
}

func signUpHook(
	ctx context.Context,
	users *users.UsersRepository,
	userObject epmodels.User,
	formFields []epmodels.TypeFormField,
) (*user.User, error) {
	var username string
	for _, field := range formFields {
		if field.ID == "username" {
			username = field.Value
			break
		}
	}

	created, err := users.CreateUser(ctx, user.CreateUserRequest{
		Id:       userObject.ID,
		Email:    userObject.Email,
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return created, nil
}
