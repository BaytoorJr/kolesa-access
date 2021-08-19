package transport

import (
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/utils/validators"
	"time"
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

func (r *GetCarsByYearRequest) Validate() error {
	if r.Year == 0 {
		return errors.IncorrectRequest.SetDevMessage("year not set")
	}

	if r.Year > time.Now().Year() {
		return errors.IncorrectRequest.SetDevMessage("incorrect year set")
	}

	return nil
}

func (r *GetCarsByAvgPriceRequest) Validate() error {
	if r.Min > r.Max {
		return errors.IncorrectRequest.SetDevMessage("max value cannot be less than min value")
	}

	return nil
}

func (r *GetCarsByPriceRequest) Validate() error {
	if r.Min > r.Max {
		return errors.IncorrectRequest.SetDevMessage("max value cannot be less than min value")
	}

	return nil
}
