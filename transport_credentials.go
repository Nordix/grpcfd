// Copyright (c) 2020 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcfd

import (
	context "context"
	"net"

	"google.golang.org/grpc/credentials"
)

type wrapTransportCredentials struct {
	credentials.TransportCredentials
}

// TransportCredentials - transport credentials that will, in addition to applying cred, cause peer.Addr to supply
// the FDSender and FDRecver interfaces
func TransportCredentials(cred credentials.TransportCredentials) credentials.TransportCredentials {
	return &wrapTransportCredentials{
		TransportCredentials: cred,
	}
}

func (c *wrapTransportCredentials) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	rawConn = wrapConn(rawConn)
	if c.TransportCredentials != nil {
		return c.TransportCredentials.ClientHandshake(ctx, authority, rawConn)
	}
	return rawConn, nil, nil
}

func (c *wrapTransportCredentials) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	rawConn = wrapConn(rawConn)
	if c.TransportCredentials != nil {
		return c.TransportCredentials.ServerHandshake(rawConn)
	}
	return rawConn, nil, nil
}

func (c *wrapTransportCredentials) Clone() credentials.TransportCredentials {
	if c.TransportCredentials != nil {
		return &wrapTransportCredentials{
			TransportCredentials: c.TransportCredentials.Clone(),
		}
	}
	return &wrapTransportCredentials{}
}

func (c *wrapTransportCredentials) Info() credentials.ProtocolInfo {
	if c.TransportCredentials == nil {
		return credentials.ProtocolInfo{}
	}
	return c.TransportCredentials.Info()
}

func (c *wrapTransportCredentials) OverrideServerName(s string) error {
	if c.TransportCredentials == nil {
		return nil
	}
	return c.TransportCredentials.OverrideServerName(s)
}
