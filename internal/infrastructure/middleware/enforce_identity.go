package middleware

import (
	"encoding/base64"
	"encoding/json"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// The Identity header present into the public headers
const headerXRhIdentity = "X-Rh-Identity"

// FIXME Refactor to use the signature: func(c echo.Context) Error
//
//	so that the predicate has information about the http Request
//	context

// Represent t
type IdentityPredicate func(data *identity.Identity) error

// identityConfig Represent the configuration for this middleware
// enforcement.
type identityConfig struct {
	// Skipper function to skip for some request if necessary
	skipper echo_middleware.Skipper
	// Map of predicates to be applied, all the predicates must
	// return true, if any of them fail, the enforcement will
	// return error for the request.
	predicates map[string]IdentityPredicate
}

// NewIdentityConfig creates a new identityConfig for the
// EnforcementIdentity middleware.
// Return an identityConfig structure to configure the
// middleware.
func NewIdentityConfig() *identityConfig {
	return &identityConfig{
		predicates: map[string]IdentityPredicate{},
	}
}

// SetSkipper set a skipper function for the middleware.
// skipper is the function which check by using the current
// request context to check if the current request will be
// processed by this middleware.
// Return the identityConfig updated.
func (ic *identityConfig) SetSkipper(skipper echo_middleware.Skipper) *identityConfig {
	ic.skipper = skipper
	return ic
}

// AddPredicate add a predicate function to check the IdentityEnforcement,
// by allowing reuse the same middleware for different enforcements. We can
// add several functions, but if a key collide, the predicate will be overrided.
// key that will be associated to this predicate, it is used to report to the
// log which predicate failed.
// predicate is the check function to be added.
// Return the identityConfig updated.
func (ic *identityConfig) AddPredicate(key string, predicate IdentityPredicate) *identityConfig {
	if predicate != nil {
		ic.predicates[key] = predicate
	}
	return ic
}

// IdentityAlwaysTrue is a predicate that always return nil
// so everything was ok.
// data is the reference to the identity.Identity data.
// Return nil on success or an error with additional
// information about the predicate failure.
func IdentityAlwaysTrue(data *identity.Identity) error {
	return nil
}

// FIXME Add user enforcement predicate

// FIXME Add cert enforcement predicate

// EnforceIdentityWithConfig instantiate a EnforceIdentity middleware
// for the configuration provided. This middleware depends on
// NewContext middleware. If the request pass the enforcement
// check, then the unmarshalled version of the identity is stored
// for the request context.
// config is the configuration with the skipper and predicates
// to be used for the middleware.
// Return an echo middleware function.
func EnforceIdentityWithConfig(config *identityConfig) func(echo.HandlerFunc) echo.HandlerFunc {
	if config == nil {
		panic("config cannot be nil")
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.skipper != nil && config.skipper(c) {
				return next(c)
			}
			cc, ok := c.(DomainContextInterface)
			if !ok {
				c.Logger().Error("Expected a 'Interface'")
				return echo.ErrInternalServerError
			}
			b64Identity := cc.Request().Header.Get(headerXRhIdentity)
			if b64Identity == "" {
				cc.Logger().Error("%s not present", headerXRhIdentity)
				return echo.ErrUnauthorized
			}
			stringIdentity, err := base64.StdEncoding.DecodeString(b64Identity)
			if err != nil {
				cc.Logger().Error(err)
				return echo.ErrUnauthorized
			}
			var data *identity.Identity = &identity.Identity{}
			if err := json.Unmarshal([]byte(stringIdentity), data); err != nil {
				cc.Logger().Error(err)
				return echo.ErrUnauthorized
			}

			// All the predicates should return true
			for _, predicate := range config.predicates {
				if err := predicate(data); err != nil {
					if err != nil {
						cc.Logger().Error(err)
					}
					return echo.ErrUnauthorized
				}
			}
			cc.SetIdentity(data)
			return next(c)
		}
	}
}
