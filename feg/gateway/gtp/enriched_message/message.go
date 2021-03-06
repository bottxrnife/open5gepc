/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package enriched_message

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
	"github.com/wmnsk/go-gtp/gtpv2/message"
)

// MessageWithGrpc wraps Message interface so we can use is as a Message. It adds a field to store
// a GRPC message that should be the translation  of the GPT message
type MessageWithGrpc struct {
	message.Message               // GTP
	grpcMessage     proto.Message // GRPC
}

func NewMessageWithGrpc(gtpMessage message.Message, grpcMessage proto.Message) *MessageWithGrpc {
	return &MessageWithGrpc{
		Message:     gtpMessage,
		grpcMessage: grpcMessage,
	}
}

func (m *MessageWithGrpc) GetGrpcMessage() proto.Message {
	return m.grpcMessage
}

func ExtractGrpcMessageFromGtpMessage(incomingMsg message.Message) (proto.Message, error) {
	// check if it is NewMessageWithGrpc
	var withGrpc *MessageWithGrpc
	switch m := incomingMsg.(type) {
	case *MessageWithGrpc:
		withGrpc = m
	default:
		return nil, fmt.Errorf("incomming message it is not MessageWithGrpc type %+v", incomingMsg)
	}
	grpcMessage := withGrpc.GetGrpcMessage()
	return grpcMessage, nil
}
