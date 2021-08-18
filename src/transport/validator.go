package transport

import (
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/validators"
)

func (r *GetCarsByMarkRequest) Validate() error {

	return validators.StringNonEmptyValidate("mark", r.Mark)
}

func (r *GetCarsByMarkAndModelRequest) Validate() error {
	if r.Model == "" {
		return errors.AccessDenied.SetDevMessage("missing car model")
	}

	return validators.StringNonEmptyValidate("mark", r.Mark)
}
