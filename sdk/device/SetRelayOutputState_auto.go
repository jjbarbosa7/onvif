// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/jjbarbosa7/onvif"
	"github.com/jjbarbosa7/onvif/sdk"
	"github.com/jjbarbosa7/onvif/device"
)

// Call_SetRelayOutputState forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRelayOutputStateResponse.
func Call_SetRelayOutputState(ctx context.Context, dev *onvif.Device, request device.SetRelayOutputState) (device.SetRelayOutputStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputStateResponse device.SetRelayOutputStateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRelayOutputState")
		return reply.Body.SetRelayOutputStateResponse, errors.Annotate(err, "reply")
	}
}
