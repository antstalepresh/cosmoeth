// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoeth/cosmoeth/storage_proof.proto

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

type StorageProof struct {
	Key   string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Proof []string `protobuf:"bytes,3,rep,name=proof,proto3" json:"proof,omitempty"`
}

func (m *StorageProof) Reset()         { *m = StorageProof{} }
func (m *StorageProof) String() string { return proto.CompactTextString(m) }
func (*StorageProof) ProtoMessage()    {}
func (*StorageProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff55e36f4575f6ea, []int{0}
}
func (m *StorageProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StorageProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StorageProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StorageProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageProof.Merge(m, src)
}
func (m *StorageProof) XXX_Size() int {
	return m.Size()
}
func (m *StorageProof) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageProof.DiscardUnknown(m)
}

var xxx_messageInfo_StorageProof proto.InternalMessageInfo

func (m *StorageProof) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StorageProof) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *StorageProof) GetProof() []string {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*StorageProof)(nil), "cosmoeth.cosmoeth.StorageProof")
}

func init() {
	proto.RegisterFile("cosmoeth/cosmoeth/storage_proof.proto", fileDescriptor_ff55e36f4575f6ea)
}

var fileDescriptor_ff55e36f4575f6ea = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x4f, 0x2d, 0xc9, 0xd0, 0x87, 0x33, 0x8a, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53, 0xe3, 0x0b,
	0x8a, 0xf2, 0xf3, 0xd3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x61, 0xb2, 0x7a, 0x30,
	0x86, 0x92, 0x0f, 0x17, 0x4f, 0x30, 0x44, 0x65, 0x00, 0x48, 0xa1, 0x90, 0x00, 0x17, 0x73, 0x76,
	0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96,
	0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0x40, 0xa2, 0x60, 0x93, 0x25, 0x98, 0x15,
	0x98, 0x41, 0xa2, 0x60, 0x8e, 0x93, 0xf1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x44, 0x49, 0x3a, 0x83, 0x6c, 0x74, 0x2d, 0xc9, 0xd0, 0xaf, 0x40, 0xb8, 0xb1, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0xec, 0x38, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xfe, 0x56,
	0x71, 0xc5, 0x00, 0x00, 0x00,
}

func (m *StorageProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StorageProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		for iNdEx := len(m.Proof) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Proof[iNdEx])
			copy(dAtA[i:], m.Proof[iNdEx])
			i = encodeVarintStorageProof(dAtA, i, uint64(len(m.Proof[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintStorageProof(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintStorageProof(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStorageProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovStorageProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StorageProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovStorageProof(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovStorageProof(uint64(l))
	}
	if len(m.Proof) > 0 {
		for _, s := range m.Proof {
			l = len(s)
			n += 1 + l + sovStorageProof(uint64(l))
		}
	}
	return n
}

func sovStorageProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStorageProof(x uint64) (n int) {
	return sovStorageProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StorageProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorageProof
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
			return fmt.Errorf("proto: StorageProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageProof
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
				return ErrInvalidLengthStorageProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorageProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageProof
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
				return ErrInvalidLengthStorageProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorageProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorageProof
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
				return ErrInvalidLengthStorageProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorageProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorageProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorageProof
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
func skipStorageProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorageProof
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
					return 0, ErrIntOverflowStorageProof
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
					return 0, ErrIntOverflowStorageProof
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
				return 0, ErrInvalidLengthStorageProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStorageProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStorageProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStorageProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorageProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStorageProof = fmt.Errorf("proto: unexpected end of group")
)