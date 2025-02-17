// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"
	"github.com/juju/errors"
	"github.com/jjbarbosa7/onvif"
	"github.com/jjbarbosa7/onvif/sdk"
	"github.com/jjbarbosa7/onvif/event"
)

// Call_Unsubscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a UnsubscribeResponse.
func Call_Unsubscribe(ctx context.Context, dev *onvif.Device, request event.Unsubscribe) (event.UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse event.UnsubscribeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request,0,false); err != nil {
		return reply.Body.UnsubscribeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Unsubscribe")
		return reply.Body.UnsubscribeResponse, errors.Annotate(err, "reply")
	}
}
