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

// Call_RemoveMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveMetadataConfigurationResponse.
func Call_RemoveMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveMetadataConfiguration) (media.RemoveMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveMetadataConfigurationResponse media.RemoveMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveMetadataConfiguration")
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Annotate(err, "reply")
	}
}
