package serialization

import (
	"github.com/gogo/protobuf/proto"
)

// MustMarshal 将 proto.Message 进行序列化
func MustMarshal(msg proto.Message) (data []byte) {
	var err error
	defer func() {
		if recover() != nil {
			data, err = proto.Marshal(msg)
			if err != nil {
				panic(err)
			}
		}
	}()
	data, err = proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return
}

// MustUnmarshal 将 proto.Message 进行反序列化
func MustUnmarshal(b []byte, msg proto.Message) {
	if err := proto.Unmarshal(b, msg); err != nil {
		panic(err)
	}
}
