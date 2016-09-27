package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartResourceLinks chart resource links

swagger:model chartResourceLinks
*/
type ChartResourceLinks struct {

	/* latest

	Required: true
	Min Length: 1
	*/
	Latest *string `json:"latest"`
}

// Validate validates this chart resource links
func (m *ChartResourceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartResourceLinks) validateLatest(formats strfmt.Registry) error {

	if err := validate.Required("latest", "body", m.Latest); err != nil {
		return err
	}

	if err := validate.MinLength("latest", "body", string(*m.Latest), 1); err != nil {
		return err
	}

	return nil
}
