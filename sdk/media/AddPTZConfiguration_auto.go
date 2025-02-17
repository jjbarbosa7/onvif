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

// Call_AddPTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddPTZConfigurationResponse.
func Call_AddPTZConfiguration(ctx context.Context, dev *onvif.Device, request media.AddPTZConfiguration) (media.AddPTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddPTZConfigurationResponse media.AddPTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.AddPTZConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddPTZConfiguration")
		return reply.Body.AddPTZConfigurationResponse, errors.Annotate(err, "reply")
	}
}
