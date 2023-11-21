// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create admin user
	// (POST /v1/admins)
	PostV1Admins(ctx echo.Context) error

	// (POST /v1/admins/access-token)
	PostV1AdminUserAccessToken(ctx echo.Context, params PostV1AdminUserAccessTokenParams) error

	// (POST /v1/admins/sign-in)
	PostV1AdminUserSignIn(ctx echo.Context) error
	// Your GET endpoint
	// (GET /v1/admins/users)
	GetV1AdminUsers(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostV1Admins converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1Admins(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1Admins(ctx)
	return err
}

// PostV1AdminUserAccessToken converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1AdminUserAccessToken(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostV1AdminUserAccessTokenParams
	// ------------- Required query parameter "refreshToken" -------------

	err = runtime.BindQueryParameter("form", true, true, "refreshToken", ctx.QueryParams(), &params.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter refreshToken: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1AdminUserAccessToken(ctx, params)
	return err
}

// PostV1AdminUserSignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1AdminUserSignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1AdminUserSignIn(ctx)
	return err
}

// GetV1AdminUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1AdminUsers(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetV1AdminUsers(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v1/admins", wrapper.PostV1Admins)
	router.POST(baseURL+"/v1/admins/access-token", wrapper.PostV1AdminUserAccessToken)
	router.POST(baseURL+"/v1/admins/sign-in", wrapper.PostV1AdminUserSignIn)
	router.GET(baseURL+"/v1/admins/users", wrapper.GetV1AdminUsers)

}