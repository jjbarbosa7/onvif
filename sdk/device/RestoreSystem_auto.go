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

// Call_RestoreSystem forwards the call to dev.CallMethod() then parses the payload of the reply as a RestoreSystemResponse.
func Call_RestoreSystem(ctx context.Context, dev *onvif.Device, request device.RestoreSystem) (device.RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse device.RestoreSystemResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.RestoreSystemResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RestoreSystem")
		return reply.Body.RestoreSystemResponse, errors.Annotate(err, "reply")
	}
}
