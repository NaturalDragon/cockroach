// Code generated by protoc-gen-gogo.
// source: cockroach/sql/sqlwire/sql_api.proto
// DO NOT EDIT!

/*
	Package sqlwire is a generated protocol buffer package.

	It is generated from these files:
		cockroach/sql/sqlwire/sql_api.proto

	It has these top-level messages:
		SQLRequestHeader
		SQLResponseHeader
		Datum
		Result
		SQLRequest
		SQLResponse
*/
package sqlwire

import proto "github.com/gogo/protobuf/proto"
import math "math"
import cockroach_proto3 "github.com/cockroachdb/cockroach/proto"
import cockroach_proto2 "github.com/cockroachdb/cockroach/proto"

// discarding unused import gogoproto "gogoproto/gogo.pb"

import io "io"

import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// SQLRequestHeader is supplied with every CmdRequest.
type SQLRequestHeader struct {
	// User is the originating user.
	User string `protobuf:"bytes,5,opt,name=user" json:"user"`
	// Session settings that were returned in the last response that
	// contained them, being reflected back to the server.
	Session []byte `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// The transaction state returned in the previous response being
	// reflected back.
	Txn []byte `protobuf:"bytes,2,opt,name=txn" json:"txn,omitempty"`
	// CmdID is optionally specified for request idempotence
	// (i.e. replay protection).
	CmdID            cockroach_proto3.ClientCmdID `protobuf:"bytes,3,opt,name=cmd_id" json:"cmd_id"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *SQLRequestHeader) Reset()         { *m = SQLRequestHeader{} }
func (m *SQLRequestHeader) String() string { return proto.CompactTextString(m) }
func (*SQLRequestHeader) ProtoMessage()    {}

func (m *SQLRequestHeader) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SQLRequestHeader) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SQLRequestHeader) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

func (m *SQLRequestHeader) GetCmdID() cockroach_proto3.ClientCmdID {
	if m != nil {
		return m.CmdID
	}
	return cockroach_proto3.ClientCmdID{}
}

// SQLResponseHeader is returned with every Cmd response.
type SQLResponseHeader struct {
	// Error is non-nil if an error occurred.
	Error *cockroach_proto2.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// Setting that should be reflected back in all subsequent requests.
	// When not set, future requests should continue to use existing settings.
	Session []byte `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	// Transaction message returned in a response; not to be interpreted by
	// the recipient and reflected in a subsequent request. When not set,
	// the subsequent request should not contain a transaction object.
	Txn              []byte `protobuf:"bytes,3,opt,name=txn" json:"txn,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SQLResponseHeader) Reset()         { *m = SQLResponseHeader{} }
func (m *SQLResponseHeader) String() string { return proto.CompactTextString(m) }
func (*SQLResponseHeader) ProtoMessage()    {}

func (m *SQLResponseHeader) GetError() *cockroach_proto2.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SQLResponseHeader) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *SQLResponseHeader) GetTxn() []byte {
	if m != nil {
		return m.Txn
	}
	return nil
}

type Datum struct {
	Bval    *bool    `protobuf:"varint,1,opt,name=bval" json:"bval,omitempty"`
	Ival    *int64   `protobuf:"varint,2,opt,name=ival" json:"ival,omitempty"`
	Dval    *float64 `protobuf:"fixed64,3,opt,name=dval" json:"dval,omitempty"`
	Blobval []byte   `protobuf:"bytes,4,opt,name=blobval" json:"blobval,omitempty"`
	// Checksum is a CRC-32-IEEE checksum of the value.
	// If this is an integer value, then the value is interpreted as an 8
	// byte, big-endian encoded value. This value is set by the client on
	// updates to do end-to-end integrity verification. If the checksum is
	// incorrect, the update operation will fail. If the client does not
	// wish to use end-to-end checksumming, this value should be nil.
	Checksum         *uint32 `protobuf:"fixed32,9,opt,name=checksum" json:"checksum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Datum) Reset()         { *m = Datum{} }
func (m *Datum) String() string { return proto.CompactTextString(m) }
func (*Datum) ProtoMessage()    {}

func (m *Datum) GetBval() bool {
	if m != nil && m.Bval != nil {
		return *m.Bval
	}
	return false
}

func (m *Datum) GetIval() int64 {
	if m != nil && m.Ival != nil {
		return *m.Ival
	}
	return 0
}

func (m *Datum) GetDval() float64 {
	if m != nil && m.Dval != nil {
		return *m.Dval
	}
	return 0
}

func (m *Datum) GetBlobval() []byte {
	if m != nil {
		return m.Blobval
	}
	return nil
}

func (m *Datum) GetChecksum() uint32 {
	if m != nil && m.Checksum != nil {
		return *m.Checksum
	}
	return 0
}

// A Result is a collection of values representing a row
// in a result view. A column value not present in a row
// has Nil Bytes in the value.
type Result struct {
	Values           []*Datum `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

func (m *Result) GetValues() []*Datum {
	if m != nil {
		return m.Values
	}
	return nil
}

// A SQLRequest to cockroach. A transaction can consist of multiple
// SQLRequests.
type SQLRequest struct {
	// Request header.
	SQLRequestHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	Cmds             []*SQLRequest_Cmd `protobuf:"bytes,2,rep,name=cmds" json:"cmds,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SQLRequest) Reset()         { *m = SQLRequest{} }
func (m *SQLRequest) String() string { return proto.CompactTextString(m) }
func (*SQLRequest) ProtoMessage()    {}

func (m *SQLRequest) GetCmds() []*SQLRequest_Cmd {
	if m != nil {
		return m.Cmds
	}
	return nil
}

// SQL commands/queries to be serially executed by the server.
type SQLRequest_Cmd struct {
	Sql string `protobuf:"bytes,1,opt,name=sql" json:"sql"`
	// parameters are referred to in the above sql command/query using "?".
	Params           []*Datum `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SQLRequest_Cmd) Reset()         { *m = SQLRequest_Cmd{} }
func (m *SQLRequest_Cmd) String() string { return proto.CompactTextString(m) }
func (*SQLRequest_Cmd) ProtoMessage()    {}

func (m *SQLRequest_Cmd) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *SQLRequest_Cmd) GetParams() []*Datum {
	if m != nil {
		return m.Params
	}
	return nil
}

type SQLResponse struct {
	SQLResponseHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	// The names of the columns returned in the result set in the order
	// specified (when specified) in the SQL statement. The number of
	// columns must equal the number of Datum in each Result.
	Columns []string `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	// The result set for the last Cmd in the request.
	Results          []*Result `protobuf:"bytes,3,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SQLResponse) Reset()         { *m = SQLResponse{} }
func (m *SQLResponse) String() string { return proto.CompactTextString(m) }
func (*SQLResponse) ProtoMessage()    {}

func (m *SQLResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *SQLResponse) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
}
func (m *SQLRequestHeader) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txn = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CmdID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CmdID.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *SQLResponseHeader) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &cockroach_proto2.Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Session = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txn = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Datum) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bval", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Bval = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ival", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ival = &v
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dval", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			v2 := float64(math.Float64frombits(v))
			m.Dval = &v2
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blobval", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blobval = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 4
			v = uint32(data[iNdEx-4])
			v |= uint32(data[iNdEx-3]) << 8
			v |= uint32(data[iNdEx-2]) << 16
			v |= uint32(data[iNdEx-1]) << 24
			m.Checksum = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Result) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &Datum{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *SQLRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SQLRequestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SQLRequestHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cmds = append(m.Cmds, &SQLRequest_Cmd{})
			if err := m.Cmds[len(m.Cmds)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *SQLRequest_Cmd) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sql = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, &Datum{})
			if err := m.Params[len(m.Params)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *SQLResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SQLResponseHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SQLResponseHeader.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSqlApi(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func skipSqlApi(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSqlApi(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}
func (this *Datum) GetValue() interface{} {
	if this.Bval != nil {
		return this.Bval
	}
	if this.Ival != nil {
		return this.Ival
	}
	if this.Dval != nil {
		return this.Dval
	}
	if this.Blobval != nil {
		return this.Blobval
	}
	if this.Checksum != nil {
		return this.Checksum
	}
	return nil
}

func (this *Datum) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *bool:
		this.Bval = vt
	case *int64:
		this.Ival = vt
	case *float64:
		this.Dval = vt
	case []byte:
		this.Blobval = vt
	case *uint32:
		this.Checksum = vt
	default:
		return false
	}
	return true
}
func (m *SQLRequestHeader) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovSqlApi(uint64(l))
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovSqlApi(uint64(l))
	}
	if m.Txn != nil {
		l = len(m.Txn)
		n += 1 + l + sovSqlApi(uint64(l))
	}
	l = m.CmdID.Size()
	n += 1 + l + sovSqlApi(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SQLResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovSqlApi(uint64(l))
	}
	if m.Session != nil {
		l = len(m.Session)
		n += 1 + l + sovSqlApi(uint64(l))
	}
	if m.Txn != nil {
		l = len(m.Txn)
		n += 1 + l + sovSqlApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Datum) Size() (n int) {
	var l int
	_ = l
	if m.Bval != nil {
		n += 2
	}
	if m.Ival != nil {
		n += 1 + sovSqlApi(uint64(*m.Ival))
	}
	if m.Dval != nil {
		n += 9
	}
	if m.Blobval != nil {
		l = len(m.Blobval)
		n += 1 + l + sovSqlApi(uint64(l))
	}
	if m.Checksum != nil {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Result) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovSqlApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SQLRequest) Size() (n int) {
	var l int
	_ = l
	l = m.SQLRequestHeader.Size()
	n += 1 + l + sovSqlApi(uint64(l))
	if len(m.Cmds) > 0 {
		for _, e := range m.Cmds {
			l = e.Size()
			n += 1 + l + sovSqlApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SQLRequest_Cmd) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sql)
	n += 1 + l + sovSqlApi(uint64(l))
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovSqlApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SQLResponse) Size() (n int) {
	var l int
	_ = l
	l = m.SQLResponseHeader.Size()
	n += 1 + l + sovSqlApi(uint64(l))
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			l = len(s)
			n += 1 + l + sovSqlApi(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovSqlApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSqlApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSqlApi(x uint64) (n int) {
	return sovSqlApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SQLRequestHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SQLRequestHeader) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Session != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSqlApi(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	if m.Txn != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSqlApi(data, i, uint64(len(m.Txn)))
		i += copy(data[i:], m.Txn)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintSqlApi(data, i, uint64(m.CmdID.Size()))
	n1, err := m.CmdID.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x2a
	i++
	i = encodeVarintSqlApi(data, i, uint64(len(m.User)))
	i += copy(data[i:], m.User)
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SQLResponseHeader) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SQLResponseHeader) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSqlApi(data, i, uint64(m.Error.Size()))
		n2, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Session != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSqlApi(data, i, uint64(len(m.Session)))
		i += copy(data[i:], m.Session)
	}
	if m.Txn != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintSqlApi(data, i, uint64(len(m.Txn)))
		i += copy(data[i:], m.Txn)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Datum) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Datum) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Bval != nil {
		data[i] = 0x8
		i++
		if *m.Bval {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.Ival != nil {
		data[i] = 0x10
		i++
		i = encodeVarintSqlApi(data, i, uint64(*m.Ival))
	}
	if m.Dval != nil {
		data[i] = 0x19
		i++
		i = encodeFixed64SqlApi(data, i, uint64(math.Float64bits(*m.Dval)))
	}
	if m.Blobval != nil {
		data[i] = 0x22
		i++
		i = encodeVarintSqlApi(data, i, uint64(len(m.Blobval)))
		i += copy(data[i:], m.Blobval)
	}
	if m.Checksum != nil {
		data[i] = 0x4d
		i++
		i = encodeFixed32SqlApi(data, i, uint32(*m.Checksum))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Result) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Result) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintSqlApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SQLRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SQLRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintSqlApi(data, i, uint64(m.SQLRequestHeader.Size()))
	n3, err := m.SQLRequestHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Cmds) > 0 {
		for _, msg := range m.Cmds {
			data[i] = 0x12
			i++
			i = encodeVarintSqlApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SQLRequest_Cmd) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SQLRequest_Cmd) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintSqlApi(data, i, uint64(len(m.Sql)))
	i += copy(data[i:], m.Sql)
	if len(m.Params) > 0 {
		for _, msg := range m.Params {
			data[i] = 0x12
			i++
			i = encodeVarintSqlApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SQLResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SQLResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintSqlApi(data, i, uint64(m.SQLResponseHeader.Size()))
	n4, err := m.SQLResponseHeader.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			data[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			data[i] = 0x1a
			i++
			i = encodeVarintSqlApi(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64SqlApi(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32SqlApi(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSqlApi(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
