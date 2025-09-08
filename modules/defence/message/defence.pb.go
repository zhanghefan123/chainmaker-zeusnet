package message

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DefenceMsgType int32

const (
	DefenceMsgType_MSG_DDOS_ANNOUNCEMENT DefenceMsgType = 0
)

var DefenceMsgType_name = map[int32]string{
	0: "MSG_DDOS_ANNOUNCEMENT",
}

var DefenceMsgType_value = map[string]int32{
	"MSG_DDOS_ANNOUNCEMENT": 0,
}

func (x DefenceMsgType) String() string {
	return proto.EnumName(DefenceMsgType_name, int32(x))
}

func (DefenceMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8ccd92eead7e84a, []int{0}
}

type DefenceMsg struct {
	Type DefenceMsgType `protobuf:"varint,1,opt,name=type,proto3,enum=defence.DefenceMsgType" json:"type,omitempty"`
	Msg  []byte         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *DefenceMsg) Reset()         { *m = DefenceMsg{} }
func (m *DefenceMsg) String() string { return proto.CompactTextString(m) }
func (*DefenceMsg) ProtoMessage()    {}
func (*DefenceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ccd92eead7e84a, []int{0}
}
func (m *DefenceMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefenceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefenceMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefenceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefenceMsg.Merge(m, src)
}
func (m *DefenceMsg) XXX_Size() int {
	return m.Size()
}
func (m *DefenceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DefenceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DefenceMsg proto.InternalMessageInfo

func (m *DefenceMsg) GetType() DefenceMsgType {
	if m != nil {
		return m.Type
	}
	return DefenceMsgType_MSG_DDOS_ANNOUNCEMENT
}

func (m *DefenceMsg) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type DdosAnnouncementMsg struct {
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UnderDdosAttack bool   `protobuf:"varint,2,opt,name=underDdosAttack,proto3" json:"underDdosAttack,omitempty"`
}

func (m *DdosAnnouncementMsg) Reset()         { *m = DdosAnnouncementMsg{} }
func (m *DdosAnnouncementMsg) String() string { return proto.CompactTextString(m) }
func (*DdosAnnouncementMsg) ProtoMessage()    {}
func (*DdosAnnouncementMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ccd92eead7e84a, []int{1}
}
func (m *DdosAnnouncementMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DdosAnnouncementMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DdosAnnouncementMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DdosAnnouncementMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DdosAnnouncementMsg.Merge(m, src)
}
func (m *DdosAnnouncementMsg) XXX_Size() int {
	return m.Size()
}
func (m *DdosAnnouncementMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DdosAnnouncementMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DdosAnnouncementMsg proto.InternalMessageInfo

func (m *DdosAnnouncementMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DdosAnnouncementMsg) GetUnderDdosAttack() bool {
	if m != nil {
		return m.UnderDdosAttack
	}
	return false
}

func init() {
	proto.RegisterEnum("defence.DefenceMsgType", DefenceMsgType_name, DefenceMsgType_value)
	proto.RegisterType((*DefenceMsg)(nil), "defence.DefenceMsg")
	proto.RegisterType((*DdosAnnouncementMsg)(nil), "defence.DdosAnnouncementMsg")
}

func init() { proto.RegisterFile("zeusnet/defence/defence.proto", fileDescriptor_a8ccd92eead7e84a) }

var fileDescriptor_a8ccd92eead7e84a = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xac, 0xca, 0x48, 0xd3,
	0x4f, 0x49, 0x4d, 0x4b, 0xcd, 0x4b, 0x4e, 0x85, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0xec, 0x50, 0xae, 0x92, 0x37, 0x17, 0x97, 0x0b, 0x84, 0xe9, 0x5b, 0x9c, 0x2e, 0xa4, 0xcd, 0xc5,
	0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xae, 0x07, 0xd3, 0x84,
	0x50, 0x12, 0x52, 0x59, 0x90, 0x1a, 0x04, 0x56, 0x24, 0x24, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x2e,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x62, 0x2a, 0xf9, 0x73, 0x09, 0xbb, 0xa4, 0xe4, 0x17,
	0x3b, 0xe6, 0xe5, 0xe5, 0x97, 0xe6, 0x25, 0xa7, 0xe6, 0xa6, 0xe6, 0x95, 0x80, 0x4c, 0xe5, 0xe3,
	0x62, 0xca, 0x4c, 0x01, 0x9b, 0xc9, 0x19, 0xc4, 0x94, 0x99, 0x22, 0xa4, 0xc1, 0xc5, 0x5f, 0x9a,
	0x97, 0x92, 0x5a, 0x04, 0x56, 0x5b, 0x52, 0x92, 0x98, 0x9c, 0x0d, 0x36, 0x84, 0x23, 0x08, 0x5d,
	0x58, 0x4b, 0x9b, 0x8b, 0x0f, 0xd5, 0x6a, 0x21, 0x49, 0x2e, 0x51, 0xdf, 0x60, 0xf7, 0x78, 0x17,
	0x17, 0xff, 0xe0, 0x78, 0x47, 0x3f, 0x3f, 0xff, 0x50, 0x3f, 0x67, 0x57, 0x5f, 0x57, 0xbf, 0x10,
	0x01, 0x06, 0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03,
	0x7b, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x01, 0xf0, 0xea, 0x4c, 0x11, 0x01, 0x00, 0x00,
}

func (m *DefenceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefenceMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefenceMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintDefence(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintDefence(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DdosAnnouncementMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DdosAnnouncementMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DdosAnnouncementMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnderDdosAttack {
		i--
		if m.UnderDdosAttack {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDefence(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDefence(dAtA []byte, offset int, v uint64) int {
	offset -= sovDefence(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DefenceMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovDefence(uint64(m.Type))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovDefence(uint64(l))
	}
	return n
}

func (m *DdosAnnouncementMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDefence(uint64(l))
	}
	if m.UnderDdosAttack {
		n += 2
	}
	return n
}

func sovDefence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDefence(x uint64) (n int) {
	return sovDefence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DefenceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DefenceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefenceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= DefenceMsgType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDefence
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDefence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = append(m.Msg[:0], dAtA[iNdEx:postIndex]...)
			if m.Msg == nil {
				m.Msg = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDefence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDefence
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DdosAnnouncementMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDefence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DdosAnnouncementMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DdosAnnouncementMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDefence
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDefence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderDdosAttack", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UnderDdosAttack = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDefence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDefence
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDefence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDefence
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDefence
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDefence
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDefence
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDefence
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDefence        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDefence          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDefence = fmt.Errorf("proto: unexpected end of group")
)
