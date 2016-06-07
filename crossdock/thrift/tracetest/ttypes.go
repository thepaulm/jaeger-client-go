// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tracetest

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Transport int64

const (
	Transport_HTTP     Transport = 0
	Transport_TCHANNEL Transport = 1
)

func (p Transport) String() string {
	switch p {
	case Transport_HTTP:
		return "HTTP"
	case Transport_TCHANNEL:
		return "TCHANNEL"
	}
	return "<UNSET>"
}

func TransportFromString(s string) (Transport, error) {
	switch s {
	case "HTTP":
		return Transport_HTTP, nil
	case "TCHANNEL":
		return Transport_TCHANNEL, nil
	}
	return Transport(0), fmt.Errorf("not a valid Transport string")
}

func TransportPtr(v Transport) *Transport { return &v }

func (p Transport) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Transport) UnmarshalText(text []byte) error {
	q, err := TransportFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - Host
//  - Port
//  - Transport
//  - ClientType
//  - Downstream
type Downstream struct {
	Host       string      `thrift:"host,1,required" json:"host"`
	Port       string      `thrift:"port,2,required" json:"port"`
	Transport  Transport   `thrift:"transport,3,required" json:"transport"`
	ClientType string      `thrift:"clientType,4,required" json:"clientType"`
	Downstream *Downstream `thrift:"downstream,5" json:"downstream,omitempty"`
}

func NewDownstream() *Downstream {
	return &Downstream{}
}

func (p *Downstream) GetHost() string {
	return p.Host
}

func (p *Downstream) GetPort() string {
	return p.Port
}

func (p *Downstream) GetTransport() Transport {
	return p.Transport
}

func (p *Downstream) GetClientType() string {
	return p.ClientType
}

var Downstream_Downstream_DEFAULT Downstream

func (p *Downstream) GetDownstream() Downstream {
	if !p.IsSetDownstream() {
		return Downstream_Downstream_DEFAULT
	}
	return *p.Downstream
}
func (p *Downstream) IsSetDownstream() bool {
	return p.Downstream != nil
}

func (p *Downstream) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetHost bool = false
	var issetPort bool = false
	var issetTransport bool = false
	var issetClientType bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetHost = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetPort = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetTransport = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
			issetClientType = true
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetHost {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Host is not set"))
	}
	if !issetPort {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Port is not set"))
	}
	if !issetTransport {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Transport is not set"))
	}
	if !issetClientType {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ClientType is not set"))
	}
	return nil
}

func (p *Downstream) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Host = v
	}
	return nil
}

func (p *Downstream) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Port = v
	}
	return nil
}

func (p *Downstream) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := Transport(v)
		p.Transport = temp
	}
	return nil
}

func (p *Downstream) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.ClientType = v
	}
	return nil
}

func (p *Downstream) readField5(iprot thrift.TProtocol) error {
	p.Downstream = &Downstream{}
	if err := p.Downstream.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Downstream), err)
	}
	return nil
}

func (p *Downstream) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Downstream"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Downstream) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("host", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:host: ", p), err)
	}
	if err := oprot.WriteString(string(p.Host)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.host (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:host: ", p), err)
	}
	return err
}

func (p *Downstream) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("port", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:port: ", p), err)
	}
	if err := oprot.WriteString(string(p.Port)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.port (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:port: ", p), err)
	}
	return err
}

func (p *Downstream) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("transport", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:transport: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Transport)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.transport (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:transport: ", p), err)
	}
	return err
}

func (p *Downstream) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("clientType", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:clientType: ", p), err)
	}
	if err := oprot.WriteString(string(p.ClientType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.clientType (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:clientType: ", p), err)
	}
	return err
}

func (p *Downstream) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetDownstream() {
		if err := oprot.WriteFieldBegin("downstream", thrift.STRUCT, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:downstream: ", p), err)
		}
		if err := p.Downstream.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Downstream), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:downstream: ", p), err)
		}
	}
	return err
}

func (p *Downstream) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Downstream(%+v)", *p)
}

// Attributes:
//  - Sampled
//  - Baggage
//  - Downstream
type StartTraceRequest struct {
	Sampled    bool        `thrift:"sampled,1,required" json:"sampled"`
	Baggage    string      `thrift:"baggage,2,required" json:"baggage"`
	Downstream *Downstream `thrift:"downstream,3,required" json:"downstream"`
}

func NewStartTraceRequest() *StartTraceRequest {
	return &StartTraceRequest{}
}

func (p *StartTraceRequest) GetSampled() bool {
	return p.Sampled
}

func (p *StartTraceRequest) GetBaggage() string {
	return p.Baggage
}

var StartTraceRequest_Downstream_DEFAULT *Downstream

func (p *StartTraceRequest) GetDownstream() *Downstream {
	if !p.IsSetDownstream() {
		return StartTraceRequest_Downstream_DEFAULT
	}
	return p.Downstream
}
func (p *StartTraceRequest) IsSetDownstream() bool {
	return p.Downstream != nil
}

func (p *StartTraceRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetSampled bool = false
	var issetBaggage bool = false
	var issetDownstream bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetSampled = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetBaggage = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetDownstream = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetSampled {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Sampled is not set"))
	}
	if !issetBaggage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Baggage is not set"))
	}
	if !issetDownstream {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Downstream is not set"))
	}
	return nil
}

func (p *StartTraceRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Sampled = v
	}
	return nil
}

func (p *StartTraceRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Baggage = v
	}
	return nil
}

func (p *StartTraceRequest) readField3(iprot thrift.TProtocol) error {
	p.Downstream = &Downstream{}
	if err := p.Downstream.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Downstream), err)
	}
	return nil
}

func (p *StartTraceRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("StartTraceRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *StartTraceRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sampled", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sampled: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Sampled)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sampled (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sampled: ", p), err)
	}
	return err
}

func (p *StartTraceRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("baggage", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:baggage: ", p), err)
	}
	if err := oprot.WriteString(string(p.Baggage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.baggage (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:baggage: ", p), err)
	}
	return err
}

func (p *StartTraceRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("downstream", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:downstream: ", p), err)
	}
	if err := p.Downstream.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Downstream), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:downstream: ", p), err)
	}
	return err
}

func (p *StartTraceRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StartTraceRequest(%+v)", *p)
}

// Attributes:
//  - Downstream
type JoinTraceRequest struct {
	Downstream *Downstream `thrift:"downstream,1" json:"downstream,omitempty"`
}

func NewJoinTraceRequest() *JoinTraceRequest {
	return &JoinTraceRequest{}
}

var JoinTraceRequest_Downstream_DEFAULT *Downstream

func (p *JoinTraceRequest) GetDownstream() *Downstream {
	if !p.IsSetDownstream() {
		return JoinTraceRequest_Downstream_DEFAULT
	}
	return p.Downstream
}
func (p *JoinTraceRequest) IsSetDownstream() bool {
	return p.Downstream != nil
}

func (p *JoinTraceRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *JoinTraceRequest) readField1(iprot thrift.TProtocol) error {
	p.Downstream = &Downstream{}
	if err := p.Downstream.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Downstream), err)
	}
	return nil
}

func (p *JoinTraceRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("JoinTraceRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *JoinTraceRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetDownstream() {
		if err := oprot.WriteFieldBegin("downstream", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:downstream: ", p), err)
		}
		if err := p.Downstream.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Downstream), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:downstream: ", p), err)
		}
	}
	return err
}

func (p *JoinTraceRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("JoinTraceRequest(%+v)", *p)
}

// Attributes:
//  - TraceId
//  - Sampled
//  - Baggage
type ObservedSpan struct {
	TraceId string `thrift:"traceId,1,required" json:"traceId"`
	Sampled bool   `thrift:"sampled,2,required" json:"sampled"`
	Baggage string `thrift:"baggage,3,required" json:"baggage"`
}

func NewObservedSpan() *ObservedSpan {
	return &ObservedSpan{}
}

func (p *ObservedSpan) GetTraceId() string {
	return p.TraceId
}

func (p *ObservedSpan) GetSampled() bool {
	return p.Sampled
}

func (p *ObservedSpan) GetBaggage() string {
	return p.Baggage
}
func (p *ObservedSpan) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTraceId bool = false
	var issetSampled bool = false
	var issetBaggage bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetTraceId = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetSampled = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetBaggage = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetTraceId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TraceId is not set"))
	}
	if !issetSampled {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Sampled is not set"))
	}
	if !issetBaggage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Baggage is not set"))
	}
	return nil
}

func (p *ObservedSpan) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.TraceId = v
	}
	return nil
}

func (p *ObservedSpan) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Sampled = v
	}
	return nil
}

func (p *ObservedSpan) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Baggage = v
	}
	return nil
}

func (p *ObservedSpan) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ObservedSpan"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ObservedSpan) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("traceId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:traceId: ", p), err)
	}
	if err := oprot.WriteString(string(p.TraceId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.traceId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:traceId: ", p), err)
	}
	return err
}

func (p *ObservedSpan) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sampled", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sampled: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.Sampled)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sampled (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sampled: ", p), err)
	}
	return err
}

func (p *ObservedSpan) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("baggage", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:baggage: ", p), err)
	}
	if err := oprot.WriteString(string(p.Baggage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.baggage (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:baggage: ", p), err)
	}
	return err
}

func (p *ObservedSpan) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ObservedSpan(%+v)", *p)
}

// Attributes:
//  - Span
//  - Downstream
type TraceResponse struct {
	Span       *ObservedSpan  `thrift:"span,1,required" json:"span"`
	Downstream *TraceResponse `thrift:"downstream,2" json:"downstream,omitempty"`
}

func NewTraceResponse() *TraceResponse {
	return &TraceResponse{}
}

var TraceResponse_Span_DEFAULT *ObservedSpan

func (p *TraceResponse) GetSpan() *ObservedSpan {
	if !p.IsSetSpan() {
		return TraceResponse_Span_DEFAULT
	}
	return p.Span
}

var TraceResponse_Downstream_DEFAULT TraceResponse

func (p *TraceResponse) GetDownstream() TraceResponse {
	if !p.IsSetDownstream() {
		return TraceResponse_Downstream_DEFAULT
	}
	return *p.Downstream
}
func (p *TraceResponse) IsSetSpan() bool {
	return p.Span != nil
}

func (p *TraceResponse) IsSetDownstream() bool {
	return p.Downstream != nil
}

func (p *TraceResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetSpan bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetSpan = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetSpan {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Span is not set"))
	}
	return nil
}

func (p *TraceResponse) readField1(iprot thrift.TProtocol) error {
	p.Span = &ObservedSpan{}
	if err := p.Span.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Span), err)
	}
	return nil
}

func (p *TraceResponse) readField2(iprot thrift.TProtocol) error {
	p.Downstream = &TraceResponse{}
	if err := p.Downstream.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Downstream), err)
	}
	return nil
}

func (p *TraceResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TraceResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TraceResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("span", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:span: ", p), err)
	}
	if err := p.Span.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Span), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:span: ", p), err)
	}
	return err
}

func (p *TraceResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDownstream() {
		if err := oprot.WriteFieldBegin("downstream", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:downstream: ", p), err)
		}
		if err := p.Downstream.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Downstream), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:downstream: ", p), err)
		}
	}
	return err
}

func (p *TraceResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TraceResponse(%+v)", *p)
}
