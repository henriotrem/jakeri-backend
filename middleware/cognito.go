package middleware

import (
	"context"
	"jakeri-backend/models"
	"jakeri-backend/validations"
	"os"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var Region string
var ClientID string
var UserPoolID string
var cognitoClient *cognitoidentityprovider.Client

func init() {

	Region = os.Getenv("COGNITO_REGION")
	ClientID = os.Getenv("COGNITO_CLIENT_ID")
	UserPoolID = os.Getenv("COGNITO_USER_POOL_ID")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(Region))

	if err != nil {
		panic(err)
	}

	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
}

func CognitoRegister(users models.Users) error {

	var err error

	for idx, user := range users {

		input := &cognitoidentityprovider.SignUpInput{
			Username: aws.String(*user.Email),
			Password: aws.String(*user.Password),
			ClientId: aws.String(ClientID),
			UserAttributes: []types.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(*user.Email),
				},
			},
		}

		output, err := cognitoClient.SignUp(context.Background(), input)

		if err != nil {
			return err
		}

		users[idx].ID = output.UserSub
	}
	return err
}

func CognitoDeleteUsers(users models.Users) error {

	var err error

	for _, user := range users {

		input := &cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: aws.String(UserPoolID),
			Username:   aws.String(*user.Email),
		}

		_, err := cognitoClient.AdminDeleteUser(context.Background(), input)

		if err != nil {
			return err
		}
	}
	return err
}

func CognitoDeleteUser(user *models.User) error {

	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(UserPoolID),
		Username:   aws.String(*user.Email),
	}
	_, err := cognitoClient.AdminDeleteUser(context.Background(), input)

	return err
}

func CognitoConfirmRegistration(body validations.AddConfirmationBody) (interface{}, error) {

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(ClientID),
		ConfirmationCode: &body.ConfirmationCode,
		Username:         &body.Username,
	}

	return cognitoClient.ConfirmSignUp(context.Background(), input)
}

func CognitoResendConfirmationCode(body validations.NewConfirmationBody) (interface{}, error) {
	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId: aws.String(ClientID),
		Username: &body.Username,
	}

	return cognitoClient.ResendConfirmationCode(context.Background(), input)
}

func CognitoLogin(body validations.AddSessionBody) (map[string]interface{}, error) {

	var response = make(map[string]interface{})
	var err error
	var csrp *cognitosrp.CognitoSRP
	var initOutput *cognitoidentityprovider.InitiateAuthOutput
	var challengeOutput *cognitoidentityprovider.RespondToAuthChallengeOutput

	csrp, err = cognitosrp.NewCognitoSRP(body.Username, body.Password, UserPoolID, ClientID, nil)

	if err != nil {
		return response, err
	}

	initOutput, err = cognitoClient.InitiateAuth(context.Background(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(ClientID),
		AuthParameters: csrp.GetAuthParams(),
	})

	if err != nil || initOutput.ChallengeName != types.ChallengeNameTypePasswordVerifier {
		return response, err
	}

	challengeResponses, err := csrp.PasswordVerifierChallenge(initOutput.ChallengeParameters, time.Now())

	if err != nil {
		return response, err
	}

	challengeOutput, err = cognitoClient.RespondToAuthChallenge(context.Background(), &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName:      types.ChallengeNameTypePasswordVerifier,
		ChallengeResponses: challengeResponses,
		ClientId:           aws.String(ClientID),
	})

	if err != nil {
		return response, err
	}

	response["ACCESS_TOKEN"] = *challengeOutput.AuthenticationResult.AccessToken
	response["ID_TOKEN"] = *challengeOutput.AuthenticationResult.IdToken
	response["REFRESH_TOKEN"] = *challengeOutput.AuthenticationResult.RefreshToken
	response["EXPIRES_AT"] = time.Now().Unix() + int64(challengeOutput.AuthenticationResult.ExpiresIn)

	return response, err
}

func CognitoRefreshLogin(body validations.UpdateSessionBody) (map[string]interface{}, error) {

	var response = make(map[string]interface{})
	var err error
	var output *cognitoidentityprovider.InitiateAuthOutput

	authParameters := map[string]string{
		"REFRESH_TOKEN": body.RefreshToken,
	}

	output, err = cognitoClient.InitiateAuth(context.Background(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeRefreshTokenAuth,
		ClientId:       aws.String(ClientID),
		AuthParameters: authParameters,
	})

	if err != nil {
		return response, err
	}

	response["ACCESS_TOKEN"] = output.AuthenticationResult.AccessToken
	response["EXPIRES_AT"] = time.Now().Unix() + int64(output.AuthenticationResult.ExpiresIn)

	return response, err
}

func CognitoChangePassword(accessToken *string, body validations.UpdatePasswordBody) (interface{}, error) {

	input := &cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      accessToken,
		PreviousPassword: &body.Previous,
		ProposedPassword: &body.New,
	}

	return cognitoClient.ChangePassword(context.Background(), input)
}

func CognitoForgotPassword(body validations.NewRecoveryBody) (interface{}, error) {

	input := &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: aws.String(ClientID),
		Username: &body.Username,
	}

	return cognitoClient.ForgotPassword(context.Background(), input)
}

func CognitoConfirmForgotPassword(body validations.AddRecoveryBody) (interface{}, error) {

	input := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(ClientID),
		Username:         &body.Username,
		Password:         &body.Password,
		ConfirmationCode: &body.ConfirmationCode,
	}

	return cognitoClient.ConfirmForgotPassword(context.Background(), input)
}
