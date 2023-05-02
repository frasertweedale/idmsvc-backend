package interactor

import (
	"fmt"

	"github.com/hmsidm/internal/api/header"
	"github.com/hmsidm/internal/api/public"
	api_public "github.com/hmsidm/internal/api/public"
	"github.com/hmsidm/internal/domain/model"
	"github.com/hmsidm/internal/interface/interactor"
	"github.com/lib/pq"
	"github.com/openlyinc/pointy"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

type domainInteractor struct{}

// NewDomainInteractor Create an interactor for the /domain endpoint handler
// Return an initialized instance of interactor.DomainInteractor
func NewDomainInteractor() interactor.DomainInteractor {
	return domainInteractor{}
}

// helperDomainTypeToUint transform public.CreateDomainDomainType to an uint const
// Return the uint representation or model.DomainTypeUndefined if it does not match
// the current types.
func helperDomainTypeToUint(domainType public.DomainType) uint {
	switch domainType {
	case api_public.DomainTypeRhelIdm:
		return model.DomainTypeIpa
	default:
		return model.DomainTypeUndefined
	}
}

// Create translate api request to modle.Domain internal representation.
// params is the x-rh-identity as base64 and the x-rh-insights-request-id value for this request.
// body is the CreateDomain schema received for the POST request.
// Return a reference to a new model.Domain instance and nil error for
// a success transformation, else nil and an error instance.
func (i domainInteractor) Create(params *api_public.CreateDomainParams, body *api_public.CreateDomain) (string, *model.Domain, error) {
	if params == nil {
		return "", nil, fmt.Errorf("'params' cannot be nil")
	}
	if body == nil {
		return "", nil, fmt.Errorf("'body' cannot be nil")
	}

	domain := &model.Domain{}
	// FIXME Pass the context decoded XRHID as argument, so we can use it
	//       directly into the specific service handler and pass the
	//       information to the interactors
	xrhid, err := header.DecodeXRHID(string(params.XRhIdentity))
	if err != nil {
		return "", nil, err
	}
	domain.OrgId = xrhid.Identity.OrgID
	domain.AutoEnrollmentEnabled = pointy.Bool(body.AutoEnrollmentEnabled)
	domain.DomainName = nil
	domain.Title = pointy.String(body.Title)
	domain.Description = pointy.String(body.Description)
	domain.Type = pointy.Uint(helperDomainTypeToUint(api_public.DomainType(body.Type)))
	switch *domain.Type {
	case model.DomainTypeIpa:
		domain.IpaDomain = &model.Ipa{
			RealmName:    pointy.String(""),
			RealmDomains: pq.StringArray{},
			CaCerts:      []model.IpaCert{},
			Servers:      []model.IpaServer{},
		}
	default:
		return "", nil, fmt.Errorf("'Type' is invalid")
	}
	return xrhid.Identity.OrgID, domain, nil
}

// func (i domainInteractor) PartialUpdate(id public.Id, params *api_public.PartialUpdateTodoParams, in *api_public.Todo, out *model.Todo) error {
// 	if id <= 0 {
// 		return fmt.Errorf("'id' should be a positive int64")
// 	}
// 	if params == nil {
// 		return fmt.Errorf("'params' cannot be nil")
// 	}
// 	if in == nil {
// 		return fmt.Errorf("'in' cannot be nil")
// 	}
// 	if out == nil {
// 		return fmt.Errorf("'out' cannot be nil")
// 	}
// 	out.Model.ID = id
// 	if in.Title != nil {
// 		out.Title = pointy.String(*in.Title)
// 	}
// 	if in.Body != nil {
// 		out.Description = pointy.String(*in.Body)
// 	}
// 	return nil
// }

// func (i domainInteractor) FullUpdate(id public.Id, params *api_public.UpdateTodoParams, in *api_public.Todo, out *model.Todo) error {
// 	if id <= 0 {
// 		return fmt.Errorf("'id' should be a positive int64")
// 	}
// 	if params == nil {
// 		return fmt.Errorf("'params' cannot be nil")
// 	}
// 	if in == nil {
// 		return fmt.Errorf("'in' cannot be nil")
// 	}
// 	if out == nil {
// 		return fmt.Errorf("'out' cannot be nil")
// 	}
// 	out.ID = id
// 	out.Title = pointy.String(*in.Title)
// 	out.Description = pointy.String(*in.Body)
// 	return nil
// }

// TODO Document method
func (i domainInteractor) Delete(uuid string, params *api_public.DeleteDomainParams) (string, string, error) {
	if params == nil {
		return "", "", fmt.Errorf("'params' cannot be nil")
	}
	xrhid, err := header.DecodeXRHID(string(params.XRhIdentity))
	if err != nil {
		return "", "", err
	}
	return xrhid.Identity.OrgID, uuid, nil
}

// TODO Document method
func (i domainInteractor) List(params *api_public.ListDomainsParams) (orgId string, offset int, limit int, err error) {
	if params == nil {
		return "", -1, -1, fmt.Errorf("'params' cannot be nil")
	}
	if params.Offset == nil {
		offset = 0
	} else {
		offset = *params.Offset
	}
	if params.Limit == nil {
		limit = 10
	} else {
		limit = *params.Limit
	}
	xrhid, err := header.DecodeXRHID(string(params.XRhIdentity))
	if err != nil {
		return "", -1, -1, err
	}
	return xrhid.Identity.OrgID, offset, limit, nil
}

// TODO Document method
func (i domainInteractor) GetById(uuid string, params *public.ReadDomainParams) (string, string, error) {
	if uuid == "" {
		return "", "", fmt.Errorf("'in' cannot be an empty string")
	}

	xrhid, err := header.DecodeXRHID(string(params.XRhIdentity))
	if err != nil {
		return "", "", err
	}

	return xrhid.Identity.OrgID, uuid, nil
}

// Register translates the API input format into the business
// data models for the PUT /domains/{uuid}/register endpoint.
// params contains the header parameters.
// body contains the input payload.
// Return the orgId and the business model for Ipa information,
// when success translation, else it returns empty string for orgId,
// nil for the Ipa data, and an error filled.
func (i domainInteractor) Register(xrhid *identity.XRHID, params *api_public.RegisterDomainParams, body *public.RegisterDomain) (string, *header.XRHIDMVersion, *model.Domain, error) {
	var err error
	if xrhid == nil {
		return "", nil, nil, fmt.Errorf("'xrhid' is nil")
	}
	if params == nil {
		return "", nil, nil, fmt.Errorf("'params' is nil")
	}
	if body == nil {
		return "", nil, nil, fmt.Errorf("'body' is nil")
	}
	orgId := xrhid.Identity.Internal.OrgID

	// Retrieve the ipa-hcc version information
	clientVersion := header.NewXRHIDMVersionWithHeader(params.XRhIdmVersion)
	if clientVersion == nil {
		return "", nil, nil, fmt.Errorf("'X-Rh-Idm-Version' is invalid")
	}

	// Read the body payload
	domain := &model.Domain{}
	domain.OrgId = orgId
	domain.Title = pointy.String(body.Title)
	domain.Description = pointy.String(body.Description)
	domain.AutoEnrollmentEnabled = pointy.Bool(body.AutoEnrollmentEnabled)
	domain.DomainName = pointy.String(body.DomainName)
	switch body.Type {
	case api_public.RhelIdm:
		domain.Type = pointy.Uint(model.DomainTypeIpa)
		domain.IpaDomain = &model.Ipa{}
		err = i.registerRhelIdm(body, domain.IpaDomain)
	default:
		err = fmt.Errorf("'Type=%s' is invalid", body.Type)
	}
	if err != nil {
		return "", nil, nil, err
	}

	return orgId, clientVersion, domain, nil
}

func (i domainInteractor) registerRhelIdm(body *public.RegisterDomain, domainIpa *model.Ipa) error {
	domainIpa.RealmName = pointy.String(body.RhelIdm.RealmName)

	// Translate realm domains
	i.registerRhelIdmRealmDomains(body, domainIpa)

	// Certificate list
	i.registerRhelIdmCaCerts(body, domainIpa)

	// Server list
	i.registerRhelIdmServers(body, domainIpa)

	return nil
}

func (i domainInteractor) registerRhelIdmRealmDomains(body *public.RegisterDomain, domainIpa *model.Ipa) {
	if body.RhelIdm.RealmDomains == nil {
		domainIpa.RealmDomains = pq.StringArray{}
		return
	}
	domainIpa.RealmDomains = make(pq.StringArray, 0)
	domainIpa.RealmDomains = append(
		domainIpa.RealmDomains,
		body.RhelIdm.RealmDomains...,
	)
}

func (i domainInteractor) registerRhelIdmCaCerts(body *public.RegisterDomainJSONRequestBody, domainIpa *model.Ipa) {
	if body.RhelIdm.CaCerts == nil {
		domainIpa.CaCerts = []model.IpaCert{}
		return
	}
	domainIpa.CaCerts = make([]model.IpaCert, len(body.RhelIdm.CaCerts))
	for idx := range body.RhelIdm.CaCerts {
		i.registerRhelIdmCaCertOne(&domainIpa.CaCerts[idx], &body.RhelIdm.CaCerts[idx])
	}
}

func (i domainInteractor) registerRhelIdmCaCertOne(caCert *model.IpaCert, cert *api_public.DomainIpaCert) {
	caCert.Nickname = cert.Nickname
	caCert.Issuer = cert.Issuer
	caCert.Subject = cert.Subject
	caCert.SerialNumber = cert.SerialNumber
	caCert.NotValidBefore = cert.NotValidBefore
	caCert.NotValidAfter = cert.NotValidAfter
	caCert.Pem = cert.Pem
}

func (i domainInteractor) registerRhelIdmServers(body *public.RegisterDomainJSONRequestBody, domainIpa *model.Ipa) {
	if body.RhelIdm.Servers == nil {
		domainIpa.Servers = []model.IpaServer{}
		return
	}
	// FIXME Set body.Servers as required into the openapi specification
	domainIpa.Servers = make([]model.IpaServer, len(body.RhelIdm.Servers))
	for idx, server := range body.RhelIdm.Servers {
		domainIpa.Servers[idx].FQDN = server.Fqdn
		domainIpa.Servers[idx].RHSMId = server.SubscriptionManagerId
		domainIpa.Servers[idx].PKInitServer = server.PkinitServer
		domainIpa.Servers[idx].CaServer = server.CaServer
		domainIpa.Servers[idx].HCCEnrollmentServer = server.HccEnrollmentServer
		domainIpa.Servers[idx].HCCUpdateServer = server.HccUpdateServer
	}
}
