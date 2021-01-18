package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	jsonresp "github.com/hiromaily/go-gin-wrapper/pkg/server/response/json"
)

// APIJWTer interface
type APIJWTer interface {
	APIJWTIndexPostAction(c *gin.Context)
}

// APIJWTIndexPostAction is JWT End Point [POST]
func (ctl *controller) APIJWTIndexPostAction(c *gin.Context) {
	ctl.logger.Debug("APIJWTIndexPostAction")

	// login
	// check login
	userID, mail, err := ctl.CheckLoginOnAPI(c)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	ti := time.Now().Add(time.Minute * 60).Unix()
	token, err := ctl.jwter.CreateBasicToken(ti, strconv.Itoa(userID), mail)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	ctl.logger.Debug("", zap.String("token: %s", token))

	// Make json for response and return
	jsonresp.ResponseUserJSON(c, ctl.logger, ctl.cors, 0, jsonresp.CreateJWTJson(token))
}
