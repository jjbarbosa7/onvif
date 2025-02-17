// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/jjbarbosa7/onvif"
	"github.com/jjbarbosa7/onvif/sdk"
	"github.com/jjbarbosa7/onvif/media"
)

// Call_AddVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoSourceConfigurationResponse.
func Call_AddVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoSourceConfiguration) (media.AddVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoSourceConfigurationResponse media.AddVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddVideoSourceConfiguration")
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}
