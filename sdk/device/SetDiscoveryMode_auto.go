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

// Call_SetDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDiscoveryModeResponse.
func Call_SetDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.SetDiscoveryMode) (device.SetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDiscoveryModeResponse device.SetDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.SetDiscoveryModeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetDiscoveryMode")
		return reply.Body.SetDiscoveryModeResponse, errors.Annotate(err, "reply")
	}
}
