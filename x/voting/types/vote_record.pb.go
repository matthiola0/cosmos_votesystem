// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: votesystem/voting/vote_record.proto

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

type VoteRecord struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VoterAddress string `protobuf:"bytes,2,opt,name=voterAddress,proto3" json:"voterAddress,omitempty"`
	CandidateID  string `protobuf:"bytes,3,opt,name=candidateID,proto3" json:"candidateID,omitempty"`
	ElectionID   string `protobuf:"bytes,4,opt,name=electionID,proto3" json:"electionID,omitempty"`
}

func (m *VoteRecord) Reset()         { *m = VoteRecord{} }
func (m *VoteRecord) String() string { return proto.CompactTextString(m) }
func (*VoteRecord) ProtoMessage()    {}
func (*VoteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_135e9238d7e32af5, []int{0}
}
func (m *VoteRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRecord.Merge(m, src)
}
func (m *VoteRecord) XXX_Size() int {
	return m.Size()
}
func (m *VoteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRecord proto.InternalMessageInfo

func (m *VoteRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VoteRecord) GetVoterAddress() string {
	if m != nil {
		return m.VoterAddress
	}
	return ""
}

func (m *VoteRecord) GetCandidateID() string {
	if m != nil {
		return m.CandidateID
	}
	return ""
}

func (m *VoteRecord) GetElectionID() string {
	if m != nil {
		return m.ElectionID
	}
	return ""
}

func init() {
	proto.RegisterType((*VoteRecord)(nil), "votesystem.voting.VoteRecord")
}

func init() {
	proto.RegisterFile("votesystem/voting/vote_record.proto", fileDescriptor_135e9238d7e32af5)
}

var fileDescriptor_135e9238d7e32af5 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xcb, 0x2f, 0x49,
	0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xd5, 0x2f, 0xcb, 0x2f, 0xc9, 0xcc, 0x4b, 0x07, 0x51, 0xa9,
	0xf1, 0x45, 0xa9, 0xc9, 0xf9, 0x45, 0x29, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x08,
	0x45, 0x7a, 0x10, 0x45, 0x4a, 0x4d, 0x8c, 0x5c, 0x5c, 0x61, 0xf9, 0x25, 0xa9, 0x41, 0x60, 0x75,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29,
	0x42, 0x4a, 0x5c, 0x3c, 0x20, 0x3d, 0x45, 0x8e, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x28, 0x62, 0x42, 0x0a, 0x5c, 0xdc, 0xc9, 0x89, 0x79, 0x29, 0x99,
	0x29, 0x89, 0x25, 0xa9, 0x9e, 0x2e, 0x12, 0xcc, 0x60, 0x25, 0xc8, 0x42, 0x42, 0x72, 0x5c, 0x5c,
	0xa9, 0x39, 0xa9, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x9e, 0x2e, 0x12, 0x2c, 0x60, 0x05, 0x48, 0x22,
	0x4e, 0xc6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x89, 0xe4, 0xad,
	0x0a, 0x98, 0xc7, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x7e, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0x9a, 0x50, 0x62, 0xfa, 0x00, 0x00, 0x00,
}

func (m *VoteRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ElectionID) > 0 {
		i -= len(m.ElectionID)
		copy(dAtA[i:], m.ElectionID)
		i = encodeVarintVoteRecord(dAtA, i, uint64(len(m.ElectionID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CandidateID) > 0 {
		i -= len(m.CandidateID)
		copy(dAtA[i:], m.CandidateID)
		i = encodeVarintVoteRecord(dAtA, i, uint64(len(m.CandidateID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VoterAddress) > 0 {
		i -= len(m.VoterAddress)
		copy(dAtA[i:], m.VoterAddress)
		i = encodeVarintVoteRecord(dAtA, i, uint64(len(m.VoterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVoteRecord(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVoteRecord(uint64(m.Id))
	}
	l = len(m.VoterAddress)
	if l > 0 {
		n += 1 + l + sovVoteRecord(uint64(l))
	}
	l = len(m.CandidateID)
	if l > 0 {
		n += 1 + l + sovVoteRecord(uint64(l))
	}
	l = len(m.ElectionID)
	if l > 0 {
		n += 1 + l + sovVoteRecord(uint64(l))
	}
	return n
}

func sovVoteRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteRecord(x uint64) (n int) {
	return sovVoteRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteRecord
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
			return fmt.Errorf("proto: VoteRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRecord
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
				return ErrInvalidLengthVoteRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRecord
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
				return ErrInvalidLengthVoteRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CandidateID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteRecord
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
				return ErrInvalidLengthVoteRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ElectionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteRecord
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
func skipVoteRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteRecord
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
					return 0, ErrIntOverflowVoteRecord
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
					return 0, ErrIntOverflowVoteRecord
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
				return 0, ErrInvalidLengthVoteRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteRecord = fmt.Errorf("proto: unexpected end of group")
)
