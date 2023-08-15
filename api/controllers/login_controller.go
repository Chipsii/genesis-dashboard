package controllers

import (
	"log"
	"net/http"

	"github.com/KnockOutEZ/rest-api-portfolio/api/auth"
	"github.com/KnockOutEZ/rest-api-portfolio/api/models"
	"github.com/KnockOutEZ/rest-api-portfolio/api/utils/formaterror"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type logindetails struct {
	Token  string    `json:"token"`
	UserId uuid.UUID `json:"id"`
}

func (server *Server) Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	user.Prepare()
	err := user.Validate("login")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	tokenstring, userID, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		return echo.NewHTTPError(http.StatusUnprocessableEntity, formattedError)
	}

	loginDet := logindetails{tokenstring, userID}
	return c.JSON(http.StatusOK, loginDet)
}

func (server *Server) SignIn(email, password string) (string, uuid.UUID, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", uuid.Nil, err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", uuid.Nil, err
	}

	log.Println(user)
	token, Error := auth.CreateToken(user.ID)
	userID := user.ID
	return token, userID, Error
	// return "", 0, err
}
