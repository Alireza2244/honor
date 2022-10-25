// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: honor/dex/pool.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Pool struct {
	Id      uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount  types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	Desc    string     `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Height  int32      `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Creator string     `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fac0207e5bda9a, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Pool) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Pool) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Pool) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Pool)(nil), "honor.dex.Pool")
}

func init() { proto.RegisterFile("honor/dex/pool.proto", fileDescriptor_38fac0207e5bda9a) }

var fileDescriptor_38fac0207e5bda9a = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x18, 0x84, 0xfb, 0xd7, 0x6e, 0x65, 0x23, 0x08, 0x86, 0x45, 0xe2, 0x1e, 0x62, 0xf1, 0x54, 0x10,
	0x12, 0x56, 0x0f, 0xde, 0xd7, 0x17, 0x90, 0x1e, 0xbd, 0xb5, 0x4d, 0x68, 0x03, 0xbb, 0xfd, 0x4b,
	0x13, 0xa5, 0xbe, 0x85, 0x17, 0xdf, 0x69, 0x8f, 0x7b, 0xf4, 0x24, 0xd2, 0xbe, 0x88, 0x6c, 0x5a,
	0x6f, 0x33, 0xf3, 0x0f, 0x1f, 0xfc, 0x43, 0x56, 0x35, 0x36, 0xd8, 0x49, 0xa5, 0x7b, 0xd9, 0x22,
	0xee, 0x44, 0xdb, 0xa1, 0x43, 0xba, 0xf4, 0xa9, 0x50, 0xba, 0x5f, 0xaf, 0x2a, 0xac, 0xd0, 0xa7,
	0xf2, 0xa4, 0xa6, 0xc2, 0x9a, 0x97, 0x68, 0xf7, 0x68, 0x65, 0x91, 0x5b, 0x2d, 0xdf, 0x37, 0x85,
	0x76, 0xf9, 0x46, 0x96, 0x68, 0x9a, 0xe9, 0x7e, 0xf7, 0x05, 0x24, 0x7a, 0x41, 0xdc, 0xd1, 0x4b,
	0x12, 0x1a, 0xc5, 0x20, 0x81, 0x34, 0xca, 0x42, 0xa3, 0xe8, 0x13, 0x89, 0xf3, 0x3d, 0xbe, 0x35,
	0x8e, 0x85, 0x09, 0xa4, 0x17, 0x0f, 0x37, 0x62, 0x22, 0x89, 0x13, 0x49, 0xcc, 0x24, 0xf1, 0x8c,
	0xa6, 0xd9, 0x46, 0x87, 0x9f, 0xdb, 0x20, 0x9b, 0xeb, 0x94, 0x92, 0x48, 0x69, 0x5b, 0xb2, 0xb3,
	0x04, 0xd2, 0x65, 0xe6, 0x35, 0xbd, 0x26, 0x71, 0xad, 0x4d, 0x55, 0x3b, 0x16, 0x25, 0x90, 0x2e,
	0xb2, 0xd9, 0x51, 0x46, 0xce, 0xcb, 0x4e, 0xe7, 0x0e, 0x3b, 0xb6, 0xf0, 0xf5, 0x7f, 0xbb, 0xbd,
	0x3f, 0x0c, 0x1c, 0x8e, 0x03, 0x87, 0xdf, 0x81, 0xc3, 0xe7, 0xc8, 0x83, 0xe3, 0xc8, 0x83, 0xef,
	0x91, 0x07, 0xaf, 0x57, 0xd3, 0x10, 0xbd, 0x9f, 0xc2, 0x7d, 0xb4, 0xda, 0x16, 0xb1, 0xff, 0xe5,
	0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x36, 0xee, 0xc7, 0x9e, 0x24, 0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovPool(uint64(m.Height))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
