package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateDeploymentParams creates a new CreateDeploymentParams object
// with the default values initialized.
func NewCreateDeploymentParams() CreateDeploymentParams {
	var (
		replicasDefault int64 = int64(2)
	)
	return CreateDeploymentParams{
		Replicas: &replicasDefault,
	}
}

// CreateDeploymentParams contains all the bound params for the create deployment operation
// typically these are obtained from a http.Request
//
// swagger:parameters createDeployment
type CreateDeploymentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*App ID
	  Required: true
	  In: path
	*/
	AppID int64
	/*
	  Required: true
	  In: formData
	*/
	Description string
	/*
	  Required: true
	  In: formData
	*/
	File runtime.File
	/*
	  In: formData
	  Default: 2
	*/
	Replicas *int64
	/*Team ID
	  Required: true
	  In: path
	*/
	TeamID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateDeploymentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	rAppID, rhkAppID, _ := route.Params.GetOK("app_id")
	if err := o.bindAppID(rAppID, rhkAppID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdDescription, fdhkDescription, _ := fds.GetOK("description")
	if err := o.bindDescription(fdDescription, fdhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else {
		o.File = runtime.File{Data: file, Header: fileHeader}
	}

	fdReplicas, fdhkReplicas, _ := fds.GetOK("replicas")
	if err := o.bindReplicas(fdReplicas, fdhkReplicas, route.Formats); err != nil {
		res = append(res, err)
	}

	rTeamID, rhkTeamID, _ := route.Params.GetOK("team_id")
	if err := o.bindTeamID(rTeamID, rhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateDeploymentParams) bindAppID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("app_id", "path", "int64", raw)
	}
	o.AppID = value

	return nil
}

func (o *CreateDeploymentParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("description", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("description", "formData", raw); err != nil {
		return err
	}

	o.Description = raw

	return nil
}

func (o *CreateDeploymentParams) bindReplicas(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var replicasDefault int64 = int64(2)
		o.Replicas = &replicasDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("replicas", "formData", "int64", raw)
	}
	o.Replicas = &value

	return nil
}

func (o *CreateDeploymentParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("team_id", "path", "int64", raw)
	}
	o.TeamID = value

	return nil
}
