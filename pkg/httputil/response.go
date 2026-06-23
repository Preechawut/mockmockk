// Package httputil provides reusable HTTP response helpers with a consistent
// JSON envelope and typed-error mapping.
package httputil

import (
	"errors"
	"log"
	"net/http"
	"time"

	"mockapi/pkg/apperr"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIDKey = "requestId"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-Id")
		if id == "" {
			id = uuid.NewString()
		}
		c.Set(requestIDKey, id)
		c.Header("X-Request-Id", id)
		c.Next()
	}
}

func requestID(c *gin.Context) string {
	if id := c.GetString(requestIDKey); id != "" {
		return id
	}
	return uuid.NewString()
}

type meta struct {
	RequestID string    `json:"requestId"`
	Timestamp time.Time `json:"timestamp"`
}

type successResponse struct {
	Data any  `json:"data"`
	Meta meta `json:"meta"`
}

type errorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorResponse struct {
	Error errorDetail `json:"error"`
	Meta  meta        `json:"meta"`
}

func newMeta(c *gin.Context) meta {
	return meta{RequestID: requestID(c), Timestamp: time.Now().UTC()}
}

func RespondOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, successResponse{Data: data, Meta: newMeta(c)})
}

func RespondCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, successResponse{Data: data, Meta: newMeta(c)})
}

func RespondAPIError(c *gin.Context, status int, code, message string) {
	c.JSON(status, errorResponse{
		Error: errorDetail{Code: code, Message: message},
		Meta:  newMeta(c),
	})
}

func RespondErr(c *gin.Context, err error) {
	var ae *apperr.Error
	if errors.As(err, &ae) {
		if ae.Status >= http.StatusInternalServerError {
			logErr(c, err)
		}
		RespondAPIError(c, ae.Status, ae.Code, ae.Message)
		return
	}
	logErr(c, err)
	RespondAPIError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "internal error")
}

func logErr(c *gin.Context, err error) {
	log.Printf("[%s] %s %s: %v", requestID(c), c.Request.Method, c.Request.URL.Path, err)
}
