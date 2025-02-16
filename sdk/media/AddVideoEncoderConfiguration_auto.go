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

// Call_AddVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoEncoderConfigurationResponse.
func Call_AddVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoEncoderConfiguration) (media.AddVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoEncoderConfigurationResponse media.AddVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddVideoEncoderConfiguration")
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
