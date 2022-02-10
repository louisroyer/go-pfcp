// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestPFDManagementRequest(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewPFDManagementRequest(seq,
				ie.NewApplicationIDsPFDs(
					ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
					ie.NewPFDContext(
						ie.NewPFDContents("aa", "bb", "cc", "dd", "ee", []string{"11", "22"}, []string{"33", "44"}, []string{"55", "66"}),
					),
				),
			),
			Serialized: []byte{
				0x20, 0x03, 0x00, 0x69, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3a, 0x00, 0x61,
				// ApplicationID
				0x00, 0x18, 0x00, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6d, 0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2f,
				// PFDContext
				0x00, 0x3b, 0x00, 0x38,
				0x00, 0x3d, 0x00, 0x34,
				0xff, 0x00, // Flags & Spare octet
				0x00, 0x02, 0x61, 0x61, // FD
				0x00, 0x02, 0x62, 0x62, // URL
				0x00, 0x02, 0x63, 0x63, // DN
				0x00, 0x02, 0x64, 0x64, // CP
				0x00, 0x02, 0x65, 0x65, // DNP
				0x00, 0x08, 0x00, 0x02, 0x31, 0x31, 0x00, 0x02, 0x32, 0x32, // AFD
				0x00, 0x08, 0x00, 0x02, 0x33, 0x33, 0x00, 0x02, 0x34, 0x34, // AURL
				0x00, 0x08, 0x00, 0x02, 0x35, 0x35, 0x00, 0x02, 0x36, 0x36, // ADNP
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewPFDManagementRequest(seq,
				ie.NewApplicationIDsPFDs(
					ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
					ie.NewPFDContext(
						ie.NewPFDContents("aa", "bb", "cc", "dd", "ee", []string{"11", "22"}, []string{"33", "44"}, []string{"55", "66"}),
					),
				),
				ie.NewApplicationIDsPFDs(
					ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
					ie.NewPFDContext(
						ie.NewPFDContents("aa", "bb", "cc", "dd", "ee", []string{"11", "22"}, []string{"33", "44"}, []string{"55", "66"}),
					),
				),
			),
			Serialized: []byte{
				0x20, 0x03, 0x00, 0xce, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x3a, 0x00, 0x61,
				// ApplicationID
				0x00, 0x18, 0x00, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6d, 0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2f,
				// PFDContext
				0x00, 0x3b, 0x00, 0x38,
				0x00, 0x3d, 0x00, 0x34,
				0xff, 0x00, // Flags & Spare octet
				0x00, 0x02, 0x61, 0x61, // FD
				0x00, 0x02, 0x62, 0x62, // URL
				0x00, 0x02, 0x63, 0x63, // DN
				0x00, 0x02, 0x64, 0x64, // CP
				0x00, 0x02, 0x65, 0x65, // DNP
				0x00, 0x08, 0x00, 0x02, 0x31, 0x31, 0x00, 0x02, 0x32, 0x32, // AFD
				0x00, 0x08, 0x00, 0x02, 0x33, 0x33, 0x00, 0x02, 0x34, 0x34, // AURL
				0x00, 0x08, 0x00, 0x02, 0x35, 0x35, 0x00, 0x02, 0x36, 0x36, // ADNP
				0x00, 0x3a, 0x00, 0x61,
				// ApplicationID
				0x00, 0x18, 0x00, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6d, 0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x66, 0x63, 0x70, 0x2f,
				// PFDContext
				0x00, 0x3b, 0x00, 0x38,
				0x00, 0x3d, 0x00, 0x34,
				0xff, 0x00, // Flags & Spare octet
				0x00, 0x02, 0x61, 0x61, // FD
				0x00, 0x02, 0x62, 0x62, // URL
				0x00, 0x02, 0x63, 0x63, // DN
				0x00, 0x02, 0x64, 0x64, // CP
				0x00, 0x02, 0x65, 0x65, // DNP
				0x00, 0x08, 0x00, 0x02, 0x31, 0x31, 0x00, 0x02, 0x32, 0x32, // AFD
				0x00, 0x08, 0x00, 0x02, 0x33, 0x33, 0x00, 0x02, 0x34, 0x34, // AURL
				0x00, 0x08, 0x00, 0x02, 0x35, 0x35, 0x00, 0x02, 0x36, 0x36, // ADNP
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParsePFDManagementRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
