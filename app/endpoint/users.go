package endpoint

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"net/url"
	"nevissGo/app/serializer"
	"nevissGo/app/service"
	"nevissGo/ent"
	"nevissGo/framework"
	"nevissGo/pkg/jsonhelper"
	"os"
	"sort"
	"strings"
)

var _ framework.Endpoint = &Users{}

type Users struct {
	service *service.Users
}

func NewUsers(service *service.Users) *Users {
	return &Users{
		service: service,
	}
}

func (u *Users) Endpoints(router *framework.Endpoints) {
	router.Register("users/login", u.Login)

	router.Middleware(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authKey := c.Request().Header.Get("Authorization")
			if strings.Contains(authKey, "INIT_DATA:") {
				initData := strings.Replace(authKey, "INIT_DATA:", "", 1)

				isValid, err := validateInitData(initData)
				if err != nil {
					return err
				}

				if !isValid {
					return c.String(403, "Unauthorized")
				}

				parsed, _ := url.ParseQuery(initData)
				tgUser := jsonhelper.Decode[telebot.User]([]byte(parsed.Get("user")))

				user := &ent.User{
					ID:          tgUser.ID,
					DisplayName: tgUser.FirstName,
				}

				err = u.service.GetOrRegister(c.Request().Context(), user)
				if err != nil {
					logrus.WithError(err).Error("couldn't register telegram user")
					return c.String(500, "Couldn't register telegram user")
				}

				c.Set("user", *user)

				return next(c)
			}

			if strings.Contains(authKey, "JWT:") {
				jwt := strings.Replace(authKey, "JWT:", "", 1)
				fmt.Println(jwt)
			}

			return c.String(403, "Unauthorized")
		}
	})
}

func (u *Users) Login(c *framework.Context) error {
	token := generateJWT(c.User)

	return c.Ok(serializer.NewUserWithJwt(c.User, token))
}

func generateJWT(user *ent.User) string {
	claims := jwt.MapClaims{
		"sub": fmt.Sprint(user.ID),
		"channels": []string{
			fmt.Sprintf("personal:#%d", user.ID),
			fmt.Sprintf("personal:#%s", user.GameID),
			"personal:broadcast",
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to sign the token: %v", err)
	}

	return signedToken
}

func validateInitData(inputData string) (bool, error) {
	initData, err := url.ParseQuery(inputData)
	if err != nil {
		logrus.WithError(err).Errorln("couldn't parse web app input data")
		return false, err
	}

	dataCheckString := make([]string, 0, len(initData))
	for k, v := range initData {
		if k == "hash" {
			continue
		}
		if len(v) > 0 {
			dataCheckString = append(dataCheckString, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	sort.Strings(dataCheckString)

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(os.Getenv("TELEGRAM_TOKEN")))

	hHash := hmac.New(sha256.New, secret.Sum(nil))
	hHash.Write([]byte(strings.Join(dataCheckString, "\n")))

	hash := hex.EncodeToString(hHash.Sum(nil))

	if initData.Get("hash") != hash {
		return false, nil
	}

	return true, nil
}
