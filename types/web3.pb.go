// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/web3.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// ExtensionOptionsWeb3Tx is an extension option that specifies the typed chain id,
// the fee payer as well as its signature data.
type ExtensionOptionsWeb3Tx struct {
	// typed_data_chain_id is used only in EIP712 Domain and should match
	// Ethereum network ID in a Web3 provider (e.g. Metamask).
	TypedDataChainID uint64 `protobuf:"varint,1,opt,name=typed_data_chain_id,json=typedDataChainId,proto3" json:"typedDataChainID,omitempty"`
	// fee_payer is an account address for the fee payer. It will be validated
	// during EIP712 signature checking.
	FeePayer string `protobuf:"bytes,2,opt,name=fee_payer,json=feePayer,proto3" json:"feePayer,omitempty"`
	// fee_payer_sig is a signature data from the fee paying account,
	// allows to perform fee delegation when using EIP712 Domain.
	FeePayerSig []byte `protobuf:"bytes,3,opt,name=fee_payer_sig,json=feePayerSig,proto3" json:"feePayerSig,omitempty"`
}

func (m *ExtensionOptionsWeb3Tx) Reset()         { *m = ExtensionOptionsWeb3Tx{} }
func (m *ExtensionOptionsWeb3Tx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionsWeb3Tx) ProtoMessage()    {}
func (*ExtensionOptionsWeb3Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb7cd56e3c92bc3, []int{0}
}
func (m *ExtensionOptionsWeb3Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionsWeb3Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionsWeb3Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionsWeb3Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionsWeb3Tx.Merge(m, src)
}
func (m *ExtensionOptionsWeb3Tx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionsWeb3Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionsWeb3Tx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionsWeb3Tx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionOptionsWeb3Tx)(nil), "ethermint.types.v1.ExtensionOptionsWeb3Tx")
}

func init() { proto.RegisterFile("ethermint/types/v1/web3.proto", fileDescriptor_9eb7cd56e3c92bc3) }

var fileDescriptor_9eb7cd56e3c92bc3 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0x87, 0x73, 0xff, 0x7f, 0x11, 0x1b, 0x15, 0x4a, 0xd4, 0x52, 0x0b, 0x5e, 0x8a, 0x53, 0x87,
	0x92, 0x50, 0xb3, 0x09, 0x2e, 0xb1, 0x0a, 0x4e, 0x8a, 0x16, 0x04, 0x97, 0x70, 0xc9, 0xbd, 0x4d,
	0x6f, 0xc8, 0x5d, 0x68, 0xde, 0xd6, 0xf6, 0x1b, 0x38, 0xfa, 0x11, 0xfc, 0x38, 0x8e, 0x1d, 0x9d,
	0x82, 0xa4, 0x5b, 0x77, 0x77, 0xb9, 0x96, 0x96, 0xd0, 0xed, 0xe5, 0x79, 0xde, 0x67, 0xf9, 0x99,
	0xe7, 0x80, 0x43, 0x18, 0x25, 0x42, 0xa2, 0x8b, 0xb3, 0x14, 0x32, 0x77, 0xd2, 0x75, 0xdf, 0x20,
	0xf4, 0x9c, 0x74, 0xa4, 0x50, 0x59, 0xd6, 0x56, 0x3b, 0x2b, 0xed, 0x4c, 0xba, 0xcd, 0x93, 0x58,
	0xc5, 0x6a, 0xa5, 0x5d, 0x7d, 0xad, 0x3f, 0x2f, 0x7e, 0x89, 0x59, 0xbf, 0x9d, 0x22, 0xc8, 0x4c,
	0x28, 0xf9, 0x90, 0xa2, 0x50, 0x32, 0x7b, 0x81, 0xd0, 0xeb, 0x4f, 0x2d, 0x66, 0x1e, 0xeb, 0x98,
	0x07, 0x9c, 0x21, 0x0b, 0xa2, 0x21, 0x13, 0x32, 0x10, 0xbc, 0x41, 0x5a, 0xa4, 0x5d, 0xf1, 0x2f,
	0x8b, 0xdc, 0xae, 0xf5, 0xb5, 0xee, 0x31, 0x64, 0x37, 0x5a, 0xde, 0xf7, 0x96, 0xb9, 0xdd, 0xc4,
	0x1d, 0xd6, 0x51, 0x89, 0x40, 0x48, 0x52, 0x9c, 0x3d, 0xd5, 0x76, 0x1c, 0xb7, 0x3c, 0xb3, 0x3a,
	0x00, 0x08, 0x52, 0x36, 0x83, 0x51, 0xe3, 0x5f, 0x8b, 0xb4, 0xab, 0x7e, 0x7d, 0x99, 0xdb, 0xd6,
	0x00, 0xe0, 0x51, 0xb3, 0x52, 0xbc, 0xbf, 0x61, 0xd6, 0xb5, 0x79, 0xb4, 0x8d, 0x82, 0x4c, 0xc4,
	0x8d, 0xff, 0x2d, 0xd2, 0x3e, 0xf4, 0xcf, 0x96, 0xb9, 0x7d, 0xba, 0x79, 0x7a, 0x16, 0x71, 0xa9,
	0x3d, 0x28, 0xe1, 0xab, 0xca, 0xfb, 0xa7, 0x6d, 0xf8, 0x77, 0x5f, 0x05, 0x25, 0xf3, 0x82, 0x92,
	0x9f, 0x82, 0x92, 0x8f, 0x05, 0x35, 0xe6, 0x0b, 0x6a, 0x7c, 0x2f, 0xa8, 0xf1, 0xda, 0x89, 0x05,
	0x0e, 0xc7, 0xa1, 0x13, 0xa9, 0xc4, 0x8d, 0x58, 0x16, 0x31, 0x2e, 0xd8, 0x40, 0x8d, 0x25, 0x67,
	0x7a, 0x9c, 0x2d, 0x5a, 0xef, 0x1e, 0xee, 0xad, 0x66, 0xf4, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0xcd, 0x1f, 0x51, 0x91, 0x01, 0x00, 0x00,
}

func (m *ExtensionOptionsWeb3Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionsWeb3Tx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionsWeb3Tx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeePayerSig) > 0 {
		i -= len(m.FeePayerSig)
		copy(dAtA[i:], m.FeePayerSig)
		i = encodeVarintWeb3(dAtA, i, uint64(len(m.FeePayerSig)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FeePayer) > 0 {
		i -= len(m.FeePayer)
		copy(dAtA[i:], m.FeePayer)
		i = encodeVarintWeb3(dAtA, i, uint64(len(m.FeePayer)))
		i--
		dAtA[i] = 0x12
	}
	if m.TypedDataChainID != 0 {
		i = encodeVarintWeb3(dAtA, i, uint64(m.TypedDataChainID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWeb3(dAtA []byte, offset int, v uint64) int {
	offset -= sovWeb3(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtensionOptionsWeb3Tx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypedDataChainID != 0 {
		n += 1 + sovWeb3(uint64(m.TypedDataChainID))
	}
	l = len(m.FeePayer)
	if l > 0 {
		n += 1 + l + sovWeb3(uint64(l))
	}
	l = len(m.FeePayerSig)
	if l > 0 {
		n += 1 + l + sovWeb3(uint64(l))
	}
	return n
}

func sovWeb3(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWeb3(x uint64) (n int) {
	return sovWeb3(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtensionOptionsWeb3Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeb3
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
			return fmt.Errorf("proto: ExtensionOptionsWeb3Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionsWeb3Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypedDataChainID", wireType)
			}
			m.TypedDataChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TypedDataChainID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb3
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
				return ErrInvalidLengthWeb3
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWeb3
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeePayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayerSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb3
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
				return ErrInvalidLengthWeb3
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWeb3
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeePayerSig = append(m.FeePayerSig[:0], dAtA[iNdEx:postIndex]...)
			if m.FeePayerSig == nil {
				m.FeePayerSig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWeb3(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWeb3
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
func skipWeb3(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWeb3
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
					return 0, ErrIntOverflowWeb3
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
					return 0, ErrIntOverflowWeb3
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
				return 0, ErrInvalidLengthWeb3
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWeb3
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWeb3
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWeb3        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWeb3          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWeb3 = fmt.Errorf("proto: unexpected end of group")
)