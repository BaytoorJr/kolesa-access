package http

import (
	"context"
	"encoding/json"
	"git.auto-nomad.kz/auto-nomad/backend/shared-libs/common-lib/errors"
	"github.com/BaytoorJr/kolesa-access/src/transport"
	"net/http"
)

func getCarsByMarkDecoders(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.GetCarsByMarkRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errors.InvalidCharacter.DeveloperMessage = err.Error()
		return nil, errors.InvalidCharacter
	}

	err = request.Validate()
	if err != nil {
		return nil, err
	}

	return request, err
}

func getCarsByMarkAndModelDecoders(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.GetCarsByMarkAndModelRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errors.InvalidCharacter.DeveloperMessage = err.Error()
		return nil, errors.InvalidCharacter
	}

	err = request.Validate()
	if err != nil {
		return nil, err
	}

	return request, nil
}