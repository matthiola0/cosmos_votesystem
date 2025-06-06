// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: votesystem/voting/has_voted_voter.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type HasVotedVoter struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ElectionID   string `protobuf:"bytes,2,opt,name=electionID,proto3" json:"electionID,omitempty"`
	VoterAddress string `protobuf:"bytes,3,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	VotedStatus  bool   `protobuf:"varint,4,opt,name=votedStatus,proto3" json:"votedStatus,omitempty"`
}

func (m *HasVotedVoter) Reset()         { *m = HasVotedVoter{} }
func (m *HasVotedVoter) String() string { return proto.CompactTextString(m) }
func (*HasVotedVoter) ProtoMessage()    {}
func (*HasVotedVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_300dad1222e7aa8c, []int{0}
}
func (m *HasVotedVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HasVotedVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HasVotedVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HasVotedVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasVotedVoter.Merge(m, src)
}
func (m *HasVotedVoter) XXX_Size() int {
	return m.Size()
}
func (m *HasVotedVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_HasVotedVoter.DiscardUnknown(m)
}

var xxx_messageInfo_HasVotedVoter proto.InternalMessageInfo

func (m *HasVotedVoter) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *HasVotedVoter) GetElectionID() string {
	if m != nil {
		return m.ElectionID
	}
	return ""
}

func (m *HasVotedVoter) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *HasVotedVoter) GetVotedStatus() bool {
	if m != nil {
		return m.VotedStatus
	}
	return false
}

func init() {
	proto.RegisterType((*HasVotedVoter)(nil), "votesystem.voting.HasVotedVoter")
}

func init() {
	proto.RegisterFile("votesystem/voting/has_voted_voter.proto", fileDescriptor_300dad1222e7aa8c)
}

var fileDescriptor_300dad1222e7aa8c = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xcb, 0x2f, 0x49,
	0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xd5, 0x2f, 0xcb, 0x2f, 0xc9, 0xcc, 0x4b, 0xd7, 0xcf, 0x48,
	0x2c, 0x8e, 0x07, 0x89, 0xa6, 0x80, 0xc9, 0x22, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x41,
	0x84, 0x42, 0x3d, 0x88, 0x42, 0xa5, 0x6e, 0x46, 0x2e, 0x5e, 0x8f, 0xc4, 0xe2, 0x30, 0x90, 0x5a,
	0x10, 0x51, 0x24, 0x24, 0xc2, 0xc5, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0xe1, 0x08, 0xc9, 0x71, 0x71, 0xa5, 0xe6, 0xa4, 0x26, 0x97, 0x64, 0xe6, 0xe7,
	0x79, 0xba, 0x48, 0x30, 0x81, 0xa5, 0x90, 0x44, 0x84, 0x94, 0xb8, 0x78, 0xc0, 0x36, 0x39, 0xa6,
	0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x83, 0x55, 0xa0, 0x88, 0x09, 0x29, 0x70, 0x71, 0x83,
	0xdd, 0x14, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x11, 0x84, 0x2c,
	0xe4, 0x64, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x92, 0x48, 0x7e,
	0xac, 0x80, 0xf9, 0xb2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x39, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x41, 0xd2, 0xda, 0x07, 0x01, 0x00, 0x00,
}

func (m *HasVotedVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HasVotedVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HasVotedVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotedStatus {
		i--
		if m.VotedStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.VoterAddress) > 0 {
		i -= len(m.VoterAddress)
		copy(dAtA[i:], m.VoterAddress)
		i = encodeVarintHasVotedVoter(dAtA, i, uint64(len(m.VoterAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ElectionID) > 0 {
		i -= len(m.ElectionID)
		copy(dAtA[i:], m.ElectionID)
		i = encodeVarintHasVotedVoter(dAtA, i, uint64(len(m.ElectionID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintHasVotedVoter(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHasVotedVoter(dAtA []byte, offset int, v uint64) int {
	offset -= sovHasVotedVoter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HasVotedVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovHasVotedVoter(uint64(l))
	}
	l = len(m.ElectionID)
	if l > 0 {
		n += 1 + l + sovHasVotedVoter(uint64(l))
	}
	l = len(m.VoterAddress)
	if l > 0 {
		n += 1 + l + sovHasVotedVoter(uint64(l))
	}
	if m.VotedStatus {
		n += 2
	}
	return n
}

func sovHasVotedVoter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHasVotedVoter(x uint64) (n int) {
	return sovHasVotedVoter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HasVotedVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHasVotedVoter
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
			return fmt.Errorf("proto: HasVotedVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HasVotedVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHasVotedVoter
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
				return ErrInvalidLengthHasVotedVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHasVotedVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHasVotedVoter
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
				return ErrInvalidLengthHasVotedVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHasVotedVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ElectionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHasVotedVoter
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
				return ErrInvalidLengthHasVotedVoter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHasVotedVoter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotedStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHasVotedVoter
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
			m.VotedStatus = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHasVotedVoter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHasVotedVoter
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
func skipHasVotedVoter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHasVotedVoter
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
					return 0, ErrIntOverflowHasVotedVoter
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
					return 0, ErrIntOverflowHasVotedVoter
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
				return 0, ErrInvalidLengthHasVotedVoter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHasVotedVoter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHasVotedVoter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHasVotedVoter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHasVotedVoter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHasVotedVoter = fmt.Errorf("proto: unexpected end of group")
)
