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

// Call_SetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioDecoderConfigurationResponse.
func Call_SetAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioDecoderConfiguration) (media.SetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioDecoderConfigurationResponse media.SetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetAudioDecoderConfiguration")
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
