package main

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	crypto "github.com/tendermint/tendermint/proto/crypto/keys"
	types1 "github.com/tendermint/tendermint/proto/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CheckTxType int32

const (
	CheckTxType_New     CheckTxType = 0
	CheckTxType_Recheck CheckTxType = 1
)

var CheckTxType_name = map[int32]string{
	0: "NEW",
	1: "RECHECK",
}

var CheckTxType_value = map[string]int32{
	"NEW":     0,
	"RECHECK": 1,
}

func (x CheckTxType) String() string {
	return proto.EnumName(CheckTxType_name, int32(x))
}

func (CheckTxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{0}
}

type ResponseOfferSnapshot_Result int32

const (
	ResponseOfferSnapshot_UNKNOWN       ResponseOfferSnapshot_Result = 0
	ResponseOfferSnapshot_ACCEPT        ResponseOfferSnapshot_Result = 1
	ResponseOfferSnapshot_ABORT         ResponseOfferSnapshot_Result = 2
	ResponseOfferSnapshot_REJECT        ResponseOfferSnapshot_Result = 3
	ResponseOfferSnapshot_REJECT_FORMAT ResponseOfferSnapshot_Result = 4
	ResponseOfferSnapshot_REJECT_SENDER ResponseOfferSnapshot_Result = 5
)

var ResponseOfferSnapshot_Result_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACCEPT",
	2: "ABORT",
	3: "REJECT",
	4: "REJECT_FORMAT",
	5: "REJECT_SENDER",
}

var ResponseOfferSnapshot_Result_value = map[string]int32{
	"UNKNOWN":       0,
	"ACCEPT":        1,
	"ABORT":         2,
	"REJECT":        3,
	"REJECT_FORMAT": 4,
	"REJECT_SENDER": 5,
}

func (x ResponseOfferSnapshot_Result) String() string {
	return proto.EnumName(ResponseOfferSnapshot_Result_name, int32(x))
}

func (ResponseOfferSnapshot_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{30, 0}
}

type ResponseApplySnapshotChunk_Result int32

const (
	ResponseApplySnapshotChunk_UNKNOWN         ResponseApplySnapshotChunk_Result = 0
	ResponseApplySnapshotChunk_ACCEPT          ResponseApplySnapshotChunk_Result = 1
	ResponseApplySnapshotChunk_ABORT           ResponseApplySnapshotChunk_Result = 2
	ResponseApplySnapshotChunk_RETRY           ResponseApplySnapshotChunk_Result = 3
	ResponseApplySnapshotChunk_RETRY_SNAPSHOT  ResponseApplySnapshotChunk_Result = 4
	ResponseApplySnapshotChunk_REJECT_SNAPSHOT ResponseApplySnapshotChunk_Result = 5
)

var ResponseApplySnapshotChunk_Result_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACCEPT",
	2: "ABORT",
	3: "RETRY",
	4: "RETRY_SNAPSHOT",
	5: "REJECT_SNAPSHOT",
}

var ResponseApplySnapshotChunk_Result_value = map[string]int32{
	"UNKNOWN":         0,
	"ACCEPT":          1,
	"ABORT":           2,
	"RETRY":           3,
	"RETRY_SNAPSHOT":  4,
	"REJECT_SNAPSHOT": 5,
}

func (x ResponseApplySnapshotChunk_Result) String() string {
	return proto.EnumName(ResponseApplySnapshotChunk_Result_name, int32(x))
}

func (ResponseApplySnapshotChunk_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{32, 0}
}

type Request struct {
	// Types that are valid to be assigned to Value:
	//	*Request_Echo
	//	*Request_Flush
	//	*Request_Info
	//	*Request_SetOption
	//	*Request_InitChain
	//	*Request_Query
	//	*Request_BeginBlock
	//	*Request_CheckTx
	//	*Request_DeliverTx
	//	*Request_EndBlock
	//	*Request_Commit
	//	*Request_ListSnapshots
	//	*Request_OfferSnapshot
	//	*Request_LoadSnapshotChunk
	//	*Request_ApplySnapshotChunk
	Value isRequest_Value `protobuf_oneof:"value"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type isRequest_Value interface {
	isRequest_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Request_Echo struct {
	Echo *RequestEcho `protobuf:"bytes,1,opt,name=echo,proto3,oneof" json:"echo,omitempty"`
}
type Request_Flush struct {
	Flush *RequestFlush `protobuf:"bytes,2,opt,name=flush,proto3,oneof" json:"flush,omitempty"`
}
type Request_Info struct {
	Info *RequestInfo `protobuf:"bytes,3,opt,name=info,proto3,oneof" json:"info,omitempty"`
}
type Request_SetOption struct {
	SetOption *RequestSetOption `protobuf:"bytes,4,opt,name=set_option,json=setOption,proto3,oneof" json:"set_option,omitempty"`
}
type Request_InitChain struct {
	InitChain *RequestInitChain `protobuf:"bytes,5,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_chain,omitempty"`
}
type Request_Query struct {
	Query *RequestQuery `protobuf:"bytes,6,opt,name=query,proto3,oneof" json:"query,omitempty"`
}
type Request_BeginBlock struct {
	BeginBlock *RequestBeginBlock `protobuf:"bytes,7,opt,name=begin_block,json=beginBlock,proto3,oneof" json:"begin_block,omitempty"`
}
type Request_CheckTx struct {
	CheckTx *RequestCheckTx `protobuf:"bytes,8,opt,name=check_tx,json=checkTx,proto3,oneof" json:"check_tx,omitempty"`
}
type Request_DeliverTx struct {
	DeliverTx *RequestDeliverTx `protobuf:"bytes,9,opt,name=deliver_tx,json=deliverTx,proto3,oneof" json:"deliver_tx,omitempty"`
}
type Request_EndBlock struct {
	EndBlock *RequestEndBlock `protobuf:"bytes,10,opt,name=end_block,json=endBlock,proto3,oneof" json:"end_block,omitempty"`
}
type Request_Commit struct {
	Commit *RequestCommit `protobuf:"bytes,11,opt,name=commit,proto3,oneof" json:"commit,omitempty"`
}
type Request_ListSnapshots struct {
	ListSnapshots *RequestListSnapshots `protobuf:"bytes,12,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"list_snapshots,omitempty"`
}
type Request_OfferSnapshot struct {
	OfferSnapshot *RequestOfferSnapshot `protobuf:"bytes,13,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"offer_snapshot,omitempty"`
}
type Request_LoadSnapshotChunk struct {
	LoadSnapshotChunk *RequestLoadSnapshotChunk `protobuf:"bytes,14,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"load_snapshot_chunk,omitempty"`
}
type Request_ApplySnapshotChunk struct {
	ApplySnapshotChunk *RequestApplySnapshotChunk `protobuf:"bytes,15,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"apply_snapshot_chunk,omitempty"`
}

func (*Request_Echo) isRequest_Value()               {}
func (*Request_Flush) isRequest_Value()              {}
func (*Request_Info) isRequest_Value()               {}
func (*Request_SetOption) isRequest_Value()          {}
func (*Request_InitChain) isRequest_Value()          {}
func (*Request_Query) isRequest_Value()              {}
func (*Request_BeginBlock) isRequest_Value()         {}
func (*Request_CheckTx) isRequest_Value()            {}
func (*Request_DeliverTx) isRequest_Value()          {}
func (*Request_EndBlock) isRequest_Value()           {}
func (*Request_Commit) isRequest_Value()             {}
func (*Request_ListSnapshots) isRequest_Value()      {}
func (*Request_OfferSnapshot) isRequest_Value()      {}
func (*Request_LoadSnapshotChunk) isRequest_Value()  {}
func (*Request_ApplySnapshotChunk) isRequest_Value() {}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Request) GetEcho() *RequestEcho {
	if x, ok := m.GetValue().(*Request_Echo); ok {
		return x.Echo
	}
	return nil
}

func (m *Request) GetFlush() *RequestFlush {
	if x, ok := m.GetValue().(*Request_Flush); ok {
		return x.Flush
	}
	return nil
}

func (m *Request) GetInfo() *RequestInfo {
	if x, ok := m.GetValue().(*Request_Info); ok {
		return x.Info
	}
	return nil
}

func (m *Request) GetSetOption() *RequestSetOption {
	if x, ok := m.GetValue().(*Request_SetOption); ok {
		return x.SetOption
	}
	return nil
}

func (m *Request) GetInitChain() *RequestInitChain {
	if x, ok := m.GetValue().(*Request_InitChain); ok {
		return x.InitChain
	}
	return nil
}

func (m *Request) GetQuery() *RequestQuery {
	if x, ok := m.GetValue().(*Request_Query); ok {
		return x.Query
	}
	return nil
}

func (m *Request) GetBeginBlock() *RequestBeginBlock {
	if x, ok := m.GetValue().(*Request_BeginBlock); ok {
		return x.BeginBlock
	}
	return nil
}

func (m *Request) GetCheckTx() *RequestCheckTx {
	if x, ok := m.GetValue().(*Request_CheckTx); ok {
		return x.CheckTx
	}
	return nil
}

func (m *Request) GetDeliverTx() *RequestDeliverTx {
	if x, ok := m.GetValue().(*Request_DeliverTx); ok {
		return x.DeliverTx
	}
	return nil
}

func (m *Request) GetEndBlock() *RequestEndBlock {
	if x, ok := m.GetValue().(*Request_EndBlock); ok {
		return x.EndBlock
	}
	return nil
}

func (m *Request) GetCommit() *RequestCommit {
	if x, ok := m.GetValue().(*Request_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Request) GetListSnapshots() *RequestListSnapshots {
	if x, ok := m.GetValue().(*Request_ListSnapshots); ok {
		return x.ListSnapshots
	}
	return nil
}

func (m *Request) GetOfferSnapshot() *RequestOfferSnapshot {
	if x, ok := m.GetValue().(*Request_OfferSnapshot); ok {
		return x.OfferSnapshot
	}
	return nil
}

func (m *Request) GetLoadSnapshotChunk() *RequestLoadSnapshotChunk {
	if x, ok := m.GetValue().(*Request_LoadSnapshotChunk); ok {
		return x.LoadSnapshotChunk
	}
	return nil
}

func (m *Request) GetApplySnapshotChunk() *RequestApplySnapshotChunk {
	if x, ok := m.GetValue().(*Request_ApplySnapshotChunk); ok {
		return x.ApplySnapshotChunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_Echo)(nil),
		(*Request_Flush)(nil),
		(*Request_Info)(nil),
		(*Request_SetOption)(nil),
		(*Request_InitChain)(nil),
		(*Request_Query)(nil),
		(*Request_BeginBlock)(nil),
		(*Request_CheckTx)(nil),
		(*Request_DeliverTx)(nil),
		(*Request_EndBlock)(nil),
		(*Request_Commit)(nil),
		(*Request_ListSnapshots)(nil),
		(*Request_OfferSnapshot)(nil),
		(*Request_LoadSnapshotChunk)(nil),
		(*Request_ApplySnapshotChunk)(nil),
	}
}

type RequestEcho struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *RequestEcho) Reset()         { *m = RequestEcho{} }
func (m *RequestEcho) String() string { return proto.CompactTextString(m) }
func (*RequestEcho) ProtoMessage()    {}
func (*RequestEcho) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{1}
}
func (m *RequestEcho) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestEcho) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestEcho.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestEcho) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEcho.Merge(m, src)
}
func (m *RequestEcho) XXX_Size() int {
	return m.Size()
}
func (m *RequestEcho) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEcho.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEcho proto.InternalMessageInfo

func (m *RequestEcho) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RequestFlush struct {
}

func (m *RequestFlush) Reset()         { *m = RequestFlush{} }
func (m *RequestFlush) String() string { return proto.CompactTextString(m) }
func (*RequestFlush) ProtoMessage()    {}
func (*RequestFlush) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{2}
}
func (m *RequestFlush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestFlush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestFlush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestFlush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFlush.Merge(m, src)
}
func (m *RequestFlush) XXX_Size() int {
	return m.Size()
}
func (m *RequestFlush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFlush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFlush proto.InternalMessageInfo

type RequestInfo struct {
	Version      string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	BlockVersion uint64 `protobuf:"varint,2,opt,name=block_version,json=blockVersion,proto3" json:"block_version,omitempty"`
	P2PVersion   uint64 `protobuf:"varint,3,opt,name=p2p_version,json=p2pVersion,proto3" json:"p2p_version,omitempty"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{3}
}
func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return m.Size()
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequestInfo) GetBlockVersion() uint64 {
	if m != nil {
		return m.BlockVersion
	}
	return 0
}

func (m *RequestInfo) GetP2PVersion() uint64 {
	if m != nil {
		return m.P2PVersion
	}
	return 0
}

// nondeterministic
type RequestSetOption struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *RequestSetOption) Reset()         { *m = RequestSetOption{} }
func (m *RequestSetOption) String() string { return proto.CompactTextString(m) }
func (*RequestSetOption) ProtoMessage()    {}
func (*RequestSetOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{4}
}
func (m *RequestSetOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestSetOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestSetOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestSetOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSetOption.Merge(m, src)
}
func (m *RequestSetOption) XXX_Size() int {
	return m.Size()
}
func (m *RequestSetOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSetOption.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSetOption proto.InternalMessageInfo

func (m *RequestSetOption) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestSetOption) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestInitChain struct {
	Time            time.Time         `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
	ChainId         string            `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConsensusParams *ConsensusParams  `protobuf:"bytes,3,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params,omitempty"`
	Validators      []ValidatorUpdate `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators"`
	AppStateBytes   []byte            `protobuf:"bytes,5,opt,name=app_state_bytes,json=appStateBytes,proto3" json:"app_state_bytes,omitempty"`
}

func (m *RequestInitChain) Reset()         { *m = RequestInitChain{} }
func (m *RequestInitChain) String() string { return proto.CompactTextString(m) }
func (*RequestInitChain) ProtoMessage()    {}
func (*RequestInitChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{5}
}
func (m *RequestInitChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInitChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInitChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInitChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInitChain.Merge(m, src)
}
func (m *RequestInitChain) XXX_Size() int {
	return m.Size()
}
func (m *RequestInitChain) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInitChain.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInitChain proto.InternalMessageInfo

func (m *RequestInitChain) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *RequestInitChain) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *RequestInitChain) GetConsensusParams() *ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return nil
}

func (m *RequestInitChain) GetValidators() []ValidatorUpdate {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *RequestInitChain) GetAppStateBytes() []byte {
	if m != nil {
		return m.AppStateBytes
	}
	return nil
}

type RequestQuery struct {
	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Prove  bool   `protobuf:"varint,4,opt,name=prove,proto3" json:"prove,omitempty"`
}

func (m *RequestQuery) Reset()         { *m = RequestQuery{} }
func (m *RequestQuery) String() string { return proto.CompactTextString(m) }
func (*RequestQuery) ProtoMessage()    {}
func (*RequestQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{6}
}
func (m *RequestQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestQuery.Merge(m, src)
}
func (m *RequestQuery) XXX_Size() int {
	return m.Size()
}
func (m *RequestQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RequestQuery proto.InternalMessageInfo

func (m *RequestQuery) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RequestQuery) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RequestQuery) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RequestQuery) GetProve() bool {
	if m != nil {
		return m.Prove
	}
	return false
}

type RequestBeginBlock struct {
	Hash                []byte         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Header              types1.Header  `protobuf:"bytes,2,opt,name=header,proto3" json:"header"`
	LastCommitInfo      LastCommitInfo `protobuf:"bytes,3,opt,name=last_commit_info,json=lastCommitInfo,proto3" json:"last_commit_info"`
	ByzantineValidators []Evidence     `protobuf:"bytes,4,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"byzantine_validators"`
}

func (m *RequestBeginBlock) Reset()         { *m = RequestBeginBlock{} }
func (m *RequestBeginBlock) String() string { return proto.CompactTextString(m) }
func (*RequestBeginBlock) ProtoMessage()    {}
func (*RequestBeginBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{7}
}
func (m *RequestBeginBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestBeginBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestBeginBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestBeginBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBeginBlock.Merge(m, src)
}
func (m *RequestBeginBlock) XXX_Size() int {
	return m.Size()
}
func (m *RequestBeginBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBeginBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBeginBlock proto.InternalMessageInfo

func (m *RequestBeginBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *RequestBeginBlock) GetHeader() types1.Header {
	if m != nil {
		return m.Header
	}
	return types1.Header{}
}

func (m *RequestBeginBlock) GetLastCommitInfo() LastCommitInfo {
	if m != nil {
		return m.LastCommitInfo
	}
	return LastCommitInfo{}
}

func (m *RequestBeginBlock) GetByzantineValidators() []Evidence {
	if m != nil {
		return m.ByzantineValidators
	}
	return nil
}

type RequestCheckTx struct {
	Tx   []byte      `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Type CheckTxType `protobuf:"varint,2,opt,name=type,proto3,enum=tendermint.abci.CheckTxType" json:"type,omitempty"`
}

func (m *RequestCheckTx) Reset()         { *m = RequestCheckTx{} }
func (m *RequestCheckTx) String() string { return proto.CompactTextString(m) }
func (*RequestCheckTx) ProtoMessage()    {}
func (*RequestCheckTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{8}
}
func (m *RequestCheckTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCheckTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCheckTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCheckTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCheckTx.Merge(m, src)
}
func (m *RequestCheckTx) XXX_Size() int {
	return m.Size()
}
func (m *RequestCheckTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCheckTx.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCheckTx proto.InternalMessageInfo

func (m *RequestCheckTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *RequestCheckTx) GetType() CheckTxType {
	if m != nil {
		return m.Type
	}
	return CheckTxType_New
}

type RequestDeliverTx struct {
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *RequestDeliverTx) Reset()         { *m = RequestDeliverTx{} }
func (m *RequestDeliverTx) String() string { return proto.CompactTextString(m) }
func (*RequestDeliverTx) ProtoMessage()    {}
func (*RequestDeliverTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{9}
}
func (m *RequestDeliverTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestDeliverTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestDeliverTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestDeliverTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestDeliverTx.Merge(m, src)
}
func (m *RequestDeliverTx) XXX_Size() int {
	return m.Size()
}
func (m *RequestDeliverTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestDeliverTx.DiscardUnknown(m)
}

var xxx_messageInfo_RequestDeliverTx proto.InternalMessageInfo

func (m *RequestDeliverTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type RequestEndBlock struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *RequestEndBlock) Reset()         { *m = RequestEndBlock{} }
func (m *RequestEndBlock) String() string { return proto.CompactTextString(m) }
func (*RequestEndBlock) ProtoMessage()    {}
func (*RequestEndBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{10}
}
func (m *RequestEndBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestEndBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestEndBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestEndBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEndBlock.Merge(m, src)
}
func (m *RequestEndBlock) XXX_Size() int {
	return m.Size()
}
func (m *RequestEndBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEndBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEndBlock proto.InternalMessageInfo

func (m *RequestEndBlock) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type RequestCommit struct {
}

func (m *RequestCommit) Reset()         { *m = RequestCommit{} }
func (m *RequestCommit) String() string { return proto.CompactTextString(m) }
func (*RequestCommit) ProtoMessage()    {}
func (*RequestCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{11}
}
func (m *RequestCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCommit.Merge(m, src)
}
func (m *RequestCommit) XXX_Size() int {
	return m.Size()
}
func (m *RequestCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCommit.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCommit proto.InternalMessageInfo

// lists available snapshots
type RequestListSnapshots struct {
}

func (m *RequestListSnapshots) Reset()         { *m = RequestListSnapshots{} }
func (m *RequestListSnapshots) String() string { return proto.CompactTextString(m) }
func (*RequestListSnapshots) ProtoMessage()    {}
func (*RequestListSnapshots) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{12}
}
func (m *RequestListSnapshots) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestListSnapshots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestListSnapshots.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestListSnapshots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestListSnapshots.Merge(m, src)
}
func (m *RequestListSnapshots) XXX_Size() int {
	return m.Size()
}
func (m *RequestListSnapshots) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestListSnapshots.DiscardUnknown(m)
}

var xxx_messageInfo_RequestListSnapshots proto.InternalMessageInfo

// offers a snapshot to the application
type RequestOfferSnapshot struct {
	Snapshot *Snapshot `protobuf:"bytes,1,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	AppHash  []byte    `protobuf:"bytes,2,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
}

func (m *RequestOfferSnapshot) Reset()         { *m = RequestOfferSnapshot{} }
func (m *RequestOfferSnapshot) String() string { return proto.CompactTextString(m) }
func (*RequestOfferSnapshot) ProtoMessage()    {}
func (*RequestOfferSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{13}
}
func (m *RequestOfferSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestOfferSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestOfferSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestOfferSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestOfferSnapshot.Merge(m, src)
}
func (m *RequestOfferSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *RequestOfferSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestOfferSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_RequestOfferSnapshot proto.InternalMessageInfo

func (m *RequestOfferSnapshot) GetSnapshot() *Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *RequestOfferSnapshot) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

// loads a snapshot chunk
type RequestLoadSnapshotChunk struct {
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Format uint32 `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Chunk  uint32 `protobuf:"varint,3,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *RequestLoadSnapshotChunk) Reset()         { *m = RequestLoadSnapshotChunk{} }
func (m *RequestLoadSnapshotChunk) String() string { return proto.CompactTextString(m) }
func (*RequestLoadSnapshotChunk) ProtoMessage()    {}
func (*RequestLoadSnapshotChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{14}
}
func (m *RequestLoadSnapshotChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestLoadSnapshotChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestLoadSnapshotChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestLoadSnapshotChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestLoadSnapshotChunk.Merge(m, src)
}
func (m *RequestLoadSnapshotChunk) XXX_Size() int {
	return m.Size()
}
func (m *RequestLoadSnapshotChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestLoadSnapshotChunk.DiscardUnknown(m)
}

var xxx_messageInfo_RequestLoadSnapshotChunk proto.InternalMessageInfo

func (m *RequestLoadSnapshotChunk) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RequestLoadSnapshotChunk) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *RequestLoadSnapshotChunk) GetChunk() uint32 {
	if m != nil {
		return m.Chunk
	}
	return 0
}

// Applies a snapshot chunk
type RequestApplySnapshotChunk struct {
	Index  uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Chunk  []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *RequestApplySnapshotChunk) Reset()         { *m = RequestApplySnapshotChunk{} }
func (m *RequestApplySnapshotChunk) String() string { return proto.CompactTextString(m) }
func (*RequestApplySnapshotChunk) ProtoMessage()    {}
func (*RequestApplySnapshotChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{15}
}
func (m *RequestApplySnapshotChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestApplySnapshotChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestApplySnapshotChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestApplySnapshotChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestApplySnapshotChunk.Merge(m, src)
}
func (m *RequestApplySnapshotChunk) XXX_Size() int {
	return m.Size()
}
func (m *RequestApplySnapshotChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestApplySnapshotChunk.DiscardUnknown(m)
}

var xxx_messageInfo_RequestApplySnapshotChunk proto.InternalMessageInfo

func (m *RequestApplySnapshotChunk) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RequestApplySnapshotChunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *RequestApplySnapshotChunk) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type Response struct {
	// Types that are valid to be assigned to Value:
	//	*Response_Exception
	//	*Response_Echo
	//	*Response_Flush
	//	*Response_Info
	//	*Response_SetOption
	//	*Response_InitChain
	//	*Response_Query
	//	*Response_BeginBlock
	//	*Response_CheckTx
	//	*Response_DeliverTx
	//	*Response_EndBlock
	//	*Response_Commit
	//	*Response_ListSnapshots
	//	*Response_OfferSnapshot
	//	*Response_LoadSnapshotChunk
	//	*Response_ApplySnapshotChunk
	Value isResponse_Value `protobuf_oneof:"value"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{16}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type isResponse_Value interface {
	isResponse_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Response_Exception struct {
	Exception *ResponseException `protobuf:"bytes,1,opt,name=exception,proto3,oneof" json:"exception,omitempty"`
}
type Response_Echo struct {
	Echo *ResponseEcho `protobuf:"bytes,2,opt,name=echo,proto3,oneof" json:"echo,omitempty"`
}
type Response_Flush struct {
	Flush *ResponseFlush `protobuf:"bytes,3,opt,name=flush,proto3,oneof" json:"flush,omitempty"`
}
type Response_Info struct {
	Info *ResponseInfo `protobuf:"bytes,4,opt,name=info,proto3,oneof" json:"info,omitempty"`
}
type Response_SetOption struct {
	SetOption *ResponseSetOption `protobuf:"bytes,5,opt,name=set_option,json=setOption,proto3,oneof" json:"set_option,omitempty"`
}
type Response_InitChain struct {
	InitChain *ResponseInitChain `protobuf:"bytes,6,opt,name=init_chain,json=initChain,proto3,oneof" json:"init_chain,omitempty"`
}
type Response_Query struct {
	Query *ResponseQuery `protobuf:"bytes,7,opt,name=query,proto3,oneof" json:"query,omitempty"`
}
type Response_BeginBlock struct {
	BeginBlock *ResponseBeginBlock `protobuf:"bytes,8,opt,name=begin_block,json=beginBlock,proto3,oneof" json:"begin_block,omitempty"`
}
type Response_CheckTx struct {
	CheckTx *ResponseCheckTx `protobuf:"bytes,9,opt,name=check_tx,json=checkTx,proto3,oneof" json:"check_tx,omitempty"`
}
type Response_DeliverTx struct {
	DeliverTx *ResponseDeliverTx `protobuf:"bytes,10,opt,name=deliver_tx,json=deliverTx,proto3,oneof" json:"deliver_tx,omitempty"`
}
type Response_EndBlock struct {
	EndBlock *ResponseEndBlock `protobuf:"bytes,11,opt,name=end_block,json=endBlock,proto3,oneof" json:"end_block,omitempty"`
}
type Response_Commit struct {
	Commit *ResponseCommit `protobuf:"bytes,12,opt,name=commit,proto3,oneof" json:"commit,omitempty"`
}
type Response_ListSnapshots struct {
	ListSnapshots *ResponseListSnapshots `protobuf:"bytes,13,opt,name=list_snapshots,json=listSnapshots,proto3,oneof" json:"list_snapshots,omitempty"`
}
type Response_OfferSnapshot struct {
	OfferSnapshot *ResponseOfferSnapshot `protobuf:"bytes,14,opt,name=offer_snapshot,json=offerSnapshot,proto3,oneof" json:"offer_snapshot,omitempty"`
}
type Response_LoadSnapshotChunk struct {
	LoadSnapshotChunk *ResponseLoadSnapshotChunk `protobuf:"bytes,15,opt,name=load_snapshot_chunk,json=loadSnapshotChunk,proto3,oneof" json:"load_snapshot_chunk,omitempty"`
}
type Response_ApplySnapshotChunk struct {
	ApplySnapshotChunk *ResponseApplySnapshotChunk `protobuf:"bytes,16,opt,name=apply_snapshot_chunk,json=applySnapshotChunk,proto3,oneof" json:"apply_snapshot_chunk,omitempty"`
}

func (*Response_Exception) isResponse_Value()          {}
func (*Response_Echo) isResponse_Value()               {}
func (*Response_Flush) isResponse_Value()              {}
func (*Response_Info) isResponse_Value()               {}
func (*Response_SetOption) isResponse_Value()          {}
func (*Response_InitChain) isResponse_Value()          {}
func (*Response_Query) isResponse_Value()              {}
func (*Response_BeginBlock) isResponse_Value()         {}
func (*Response_CheckTx) isResponse_Value()            {}
func (*Response_DeliverTx) isResponse_Value()          {}
func (*Response_EndBlock) isResponse_Value()           {}
func (*Response_Commit) isResponse_Value()             {}
func (*Response_ListSnapshots) isResponse_Value()      {}
func (*Response_OfferSnapshot) isResponse_Value()      {}
func (*Response_LoadSnapshotChunk) isResponse_Value()  {}
func (*Response_ApplySnapshotChunk) isResponse_Value() {}

func (m *Response) GetValue() isResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Response) GetException() *ResponseException {
	if x, ok := m.GetValue().(*Response_Exception); ok {
		return x.Exception
	}
	return nil
}

func (m *Response) GetEcho() *ResponseEcho {
	if x, ok := m.GetValue().(*Response_Echo); ok {
		return x.Echo
	}
	return nil
}

func (m *Response) GetFlush() *ResponseFlush {
	if x, ok := m.GetValue().(*Response_Flush); ok {
		return x.Flush
	}
	return nil
}

func (m *Response) GetInfo() *ResponseInfo {
	if x, ok := m.GetValue().(*Response_Info); ok {
		return x.Info
	}
	return nil
}

func (m *Response) GetSetOption() *ResponseSetOption {
	if x, ok := m.GetValue().(*Response_SetOption); ok {
		return x.SetOption
	}
	return nil
}

func (m *Response) GetInitChain() *ResponseInitChain {
	if x, ok := m.GetValue().(*Response_InitChain); ok {
		return x.InitChain
	}
	return nil
}

func (m *Response) GetQuery() *ResponseQuery {
	if x, ok := m.GetValue().(*Response_Query); ok {
		return x.Query
	}
	return nil
}

func (m *Response) GetBeginBlock() *ResponseBeginBlock {
	if x, ok := m.GetValue().(*Response_BeginBlock); ok {
		return x.BeginBlock
	}
	return nil
}

func (m *Response) GetCheckTx() *ResponseCheckTx {
	if x, ok := m.GetValue().(*Response_CheckTx); ok {
		return x.CheckTx
	}
	return nil
}

func (m *Response) GetDeliverTx() *ResponseDeliverTx {
	if x, ok := m.GetValue().(*Response_DeliverTx); ok {
		return x.DeliverTx
	}
	return nil
}

func (m *Response) GetEndBlock() *ResponseEndBlock {
	if x, ok := m.GetValue().(*Response_EndBlock); ok {
		return x.EndBlock
	}
	return nil
}

func (m *Response) GetCommit() *ResponseCommit {
	if x, ok := m.GetValue().(*Response_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Response) GetListSnapshots() *ResponseListSnapshots {
	if x, ok := m.GetValue().(*Response_ListSnapshots); ok {
		return x.ListSnapshots
	}
	return nil
}

func (m *Response) GetOfferSnapshot() *ResponseOfferSnapshot {
	if x, ok := m.GetValue().(*Response_OfferSnapshot); ok {
		return x.OfferSnapshot
	}
	return nil
}

func (m *Response) GetLoadSnapshotChunk() *ResponseLoadSnapshotChunk {
	if x, ok := m.GetValue().(*Response_LoadSnapshotChunk); ok {
		return x.LoadSnapshotChunk
	}
	return nil
}

func (m *Response) GetApplySnapshotChunk() *ResponseApplySnapshotChunk {
	if x, ok := m.GetValue().(*Response_ApplySnapshotChunk); ok {
		return x.ApplySnapshotChunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Exception)(nil),
		(*Response_Echo)(nil),
		(*Response_Flush)(nil),
		(*Response_Info)(nil),
		(*Response_SetOption)(nil),
		(*Response_InitChain)(nil),
		(*Response_Query)(nil),
		(*Response_BeginBlock)(nil),
		(*Response_CheckTx)(nil),
		(*Response_DeliverTx)(nil),
		(*Response_EndBlock)(nil),
		(*Response_Commit)(nil),
		(*Response_ListSnapshots)(nil),
		(*Response_OfferSnapshot)(nil),
		(*Response_LoadSnapshotChunk)(nil),
		(*Response_ApplySnapshotChunk)(nil),
	}
}

// nondeterministic
type ResponseException struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *ResponseException) Reset()         { *m = ResponseException{} }
func (m *ResponseException) String() string { return proto.CompactTextString(m) }
func (*ResponseException) ProtoMessage()    {}
func (*ResponseException) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{17}
}
func (m *ResponseException) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseException) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseException.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseException) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseException.Merge(m, src)
}
func (m *ResponseException) XXX_Size() int {
	return m.Size()
}
func (m *ResponseException) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseException.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseException proto.InternalMessageInfo

func (m *ResponseException) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResponseEcho struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *ResponseEcho) Reset()         { *m = ResponseEcho{} }
func (m *ResponseEcho) String() string { return proto.CompactTextString(m) }
func (*ResponseEcho) ProtoMessage()    {}
func (*ResponseEcho) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{18}
}
func (m *ResponseEcho) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseEcho) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseEcho.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseEcho) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEcho.Merge(m, src)
}
func (m *ResponseEcho) XXX_Size() int {
	return m.Size()
}
func (m *ResponseEcho) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEcho.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEcho proto.InternalMessageInfo

func (m *ResponseEcho) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResponseFlush struct {
}

func (m *ResponseFlush) Reset()         { *m = ResponseFlush{} }
func (m *ResponseFlush) String() string { return proto.CompactTextString(m) }
func (*ResponseFlush) ProtoMessage()    {}
func (*ResponseFlush) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{19}
}
func (m *ResponseFlush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseFlush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseFlush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseFlush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseFlush.Merge(m, src)
}
func (m *ResponseFlush) XXX_Size() int {
	return m.Size()
}
func (m *ResponseFlush) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseFlush.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseFlush proto.InternalMessageInfo

type ResponseInfo struct {
	Data             string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Version          string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	AppVersion       uint64 `protobuf:"varint,3,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	LastBlockHeight  int64  `protobuf:"varint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockAppHash []byte `protobuf:"bytes,5,opt,name=last_block_app_hash,json=lastBlockAppHash,proto3" json:"last_block_app_hash,omitempty"`
}

func (m *ResponseInfo) Reset()         { *m = ResponseInfo{} }
func (m *ResponseInfo) String() string { return proto.CompactTextString(m) }
func (*ResponseInfo) ProtoMessage()    {}
func (*ResponseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{20}
}
func (m *ResponseInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInfo.Merge(m, src)
}
func (m *ResponseInfo) XXX_Size() int {
	return m.Size()
}
func (m *ResponseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInfo proto.InternalMessageInfo

func (m *ResponseInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ResponseInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ResponseInfo) GetAppVersion() uint64 {
	if m != nil {
		return m.AppVersion
	}
	return 0
}

func (m *ResponseInfo) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *ResponseInfo) GetLastBlockAppHash() []byte {
	if m != nil {
		return m.LastBlockAppHash
	}
	return nil
}

// nondeterministic
type ResponseSetOption struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// bytes data = 2;
	Log  string `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *ResponseSetOption) Reset()         { *m = ResponseSetOption{} }
func (m *ResponseSetOption) String() string { return proto.CompactTextString(m) }
func (*ResponseSetOption) ProtoMessage()    {}
func (*ResponseSetOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{21}
}
func (m *ResponseSetOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseSetOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseSetOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseSetOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSetOption.Merge(m, src)
}
func (m *ResponseSetOption) XXX_Size() int {
	return m.Size()
}
func (m *ResponseSetOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSetOption.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSetOption proto.InternalMessageInfo

func (m *ResponseSetOption) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseSetOption) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseSetOption) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type ResponseInitChain struct {
	ConsensusParams *ConsensusParams  `protobuf:"bytes,1,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params,omitempty"`
	Validators      []ValidatorUpdate `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators"`
}

func (m *ResponseInitChain) Reset()         { *m = ResponseInitChain{} }
func (m *ResponseInitChain) String() string { return proto.CompactTextString(m) }
func (*ResponseInitChain) ProtoMessage()    {}
func (*ResponseInitChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{22}
}
func (m *ResponseInitChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseInitChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseInitChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseInitChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseInitChain.Merge(m, src)
}
func (m *ResponseInitChain) XXX_Size() int {
	return m.Size()
}
func (m *ResponseInitChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseInitChain.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseInitChain proto.InternalMessageInfo

func (m *ResponseInitChain) GetConsensusParams() *ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return nil
}

func (m *ResponseInitChain) GetValidators() []ValidatorUpdate {
	if m != nil {
		return m.Validators
	}
	return nil
}

type ResponseQuery struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// bytes data = 2; // use "value" instead.
	Log       string           `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string           `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Index     int64            `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Key       []byte           `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte           `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	ProofOps  *crypto.ProofOps `protobuf:"bytes,8,opt,name=proof_ops,json=proofOps,proto3" json:"proof_ops,omitempty"`
	Height    int64            `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Codespace string           `protobuf:"bytes,10,opt,name=codespace,proto3" json:"codespace,omitempty"`
}

func (m *ResponseQuery) Reset()         { *m = ResponseQuery{} }
func (m *ResponseQuery) String() string { return proto.CompactTextString(m) }
func (*ResponseQuery) ProtoMessage()    {}
func (*ResponseQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{23}
}
func (m *ResponseQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseQuery.Merge(m, src)
}
func (m *ResponseQuery) XXX_Size() int {
	return m.Size()
}
func (m *ResponseQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseQuery proto.InternalMessageInfo

func (m *ResponseQuery) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseQuery) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseQuery) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseQuery) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ResponseQuery) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ResponseQuery) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ResponseQuery) GetProofOps() *crypto.ProofOps {
	if m != nil {
		return m.ProofOps
	}
	return nil
}

func (m *ResponseQuery) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ResponseQuery) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseBeginBlock struct {
	Events []Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *ResponseBeginBlock) Reset()         { *m = ResponseBeginBlock{} }
func (m *ResponseBeginBlock) String() string { return proto.CompactTextString(m) }
func (*ResponseBeginBlock) ProtoMessage()    {}
func (*ResponseBeginBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{24}
}
func (m *ResponseBeginBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseBeginBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseBeginBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseBeginBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBeginBlock.Merge(m, src)
}
func (m *ResponseBeginBlock) XXX_Size() int {
	return m.Size()
}
func (m *ResponseBeginBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBeginBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBeginBlock proto.InternalMessageInfo

func (m *ResponseBeginBlock) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ResponseCheckTx struct {
	Code      uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted int64   `protobuf:"varint,5,opt,name=gas_wanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed   int64   `protobuf:"varint,6,opt,name=gas_used,proto3" json:"gas_used,omitempty"`
	Events    []Event `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace string  `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
}

func (m *ResponseCheckTx) Reset()         { *m = ResponseCheckTx{} }
func (m *ResponseCheckTx) String() string { return proto.CompactTextString(m) }
func (*ResponseCheckTx) ProtoMessage()    {}
func (*ResponseCheckTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{25}
}
func (m *ResponseCheckTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseCheckTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseCheckTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseCheckTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCheckTx.Merge(m, src)
}
func (m *ResponseCheckTx) XXX_Size() int {
	return m.Size()
}
func (m *ResponseCheckTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCheckTx.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCheckTx proto.InternalMessageInfo

func (m *ResponseCheckTx) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseCheckTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseCheckTx) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseCheckTx) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseCheckTx) GetGasWanted() int64 {
	if m != nil {
		return m.GasWanted
	}
	return 0
}

func (m *ResponseCheckTx) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ResponseCheckTx) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ResponseCheckTx) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseDeliverTx struct {
	Code      uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string  `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted int64   `protobuf:"varint,5,opt,name=gas_wanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed   int64   `protobuf:"varint,6,opt,name=gas_used,proto3" json:"gas_used,omitempty"`
	Events    []Event `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace string  `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
}

func (m *ResponseDeliverTx) Reset()         { *m = ResponseDeliverTx{} }
func (m *ResponseDeliverTx) String() string { return proto.CompactTextString(m) }
func (*ResponseDeliverTx) ProtoMessage()    {}
func (*ResponseDeliverTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{26}
}
func (m *ResponseDeliverTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseDeliverTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseDeliverTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseDeliverTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseDeliverTx.Merge(m, src)
}
func (m *ResponseDeliverTx) XXX_Size() int {
	return m.Size()
}
func (m *ResponseDeliverTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseDeliverTx.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseDeliverTx proto.InternalMessageInfo

func (m *ResponseDeliverTx) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseDeliverTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseDeliverTx) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *ResponseDeliverTx) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ResponseDeliverTx) GetGasWanted() int64 {
	if m != nil {
		return m.GasWanted
	}
	return 0
}

func (m *ResponseDeliverTx) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ResponseDeliverTx) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ResponseDeliverTx) GetCodespace() string {
	if m != nil {
		return m.Codespace
	}
	return ""
}

type ResponseEndBlock struct {
	ValidatorUpdates      []ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates"`
	ConsensusParamUpdates *ConsensusParams  `protobuf:"bytes,2,opt,name=consensus_param_updates,json=consensusParamUpdates,proto3" json:"consensus_param_updates,omitempty"`
	Events                []Event           `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (m *ResponseEndBlock) Reset()         { *m = ResponseEndBlock{} }
func (m *ResponseEndBlock) String() string { return proto.CompactTextString(m) }
func (*ResponseEndBlock) ProtoMessage()    {}
func (*ResponseEndBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{27}
}
func (m *ResponseEndBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseEndBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseEndBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseEndBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEndBlock.Merge(m, src)
}
func (m *ResponseEndBlock) XXX_Size() int {
	return m.Size()
}
func (m *ResponseEndBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEndBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEndBlock proto.InternalMessageInfo

func (m *ResponseEndBlock) GetValidatorUpdates() []ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdates
	}
	return nil
}

func (m *ResponseEndBlock) GetConsensusParamUpdates() *ConsensusParams {
	if m != nil {
		return m.ConsensusParamUpdates
	}
	return nil
}

func (m *ResponseEndBlock) GetEvents() []Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ResponseCommit struct {
	// reserve 1
	Data         []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	RetainHeight int64  `protobuf:"varint,3,opt,name=retain_height,json=retainHeight,proto3" json:"retain_height,omitempty"`
}

func (m *ResponseCommit) Reset()         { *m = ResponseCommit{} }
func (m *ResponseCommit) String() string { return proto.CompactTextString(m) }
func (*ResponseCommit) ProtoMessage()    {}
func (*ResponseCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{28}
}
func (m *ResponseCommit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseCommit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCommit.Merge(m, src)
}
func (m *ResponseCommit) XXX_Size() int {
	return m.Size()
}
func (m *ResponseCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCommit.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCommit proto.InternalMessageInfo

func (m *ResponseCommit) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseCommit) GetRetainHeight() int64 {
	if m != nil {
		return m.RetainHeight
	}
	return 0
}

type ResponseListSnapshots struct {
	Snapshots []*Snapshot `protobuf:"bytes,1,rep,name=snapshots,proto3" json:"snapshots,omitempty"`
}

func (m *ResponseListSnapshots) Reset()         { *m = ResponseListSnapshots{} }
func (m *ResponseListSnapshots) String() string { return proto.CompactTextString(m) }
func (*ResponseListSnapshots) ProtoMessage()    {}
func (*ResponseListSnapshots) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{29}
}
func (m *ResponseListSnapshots) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseListSnapshots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseListSnapshots.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseListSnapshots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseListSnapshots.Merge(m, src)
}
func (m *ResponseListSnapshots) XXX_Size() int {
	return m.Size()
}
func (m *ResponseListSnapshots) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseListSnapshots.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseListSnapshots proto.InternalMessageInfo

func (m *ResponseListSnapshots) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type ResponseOfferSnapshot struct {
	Result ResponseOfferSnapshot_Result `protobuf:"varint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseOfferSnapshot_Result" json:"result,omitempty"`
}

func (m *ResponseOfferSnapshot) Reset()         { *m = ResponseOfferSnapshot{} }
func (m *ResponseOfferSnapshot) String() string { return proto.CompactTextString(m) }
func (*ResponseOfferSnapshot) ProtoMessage()    {}
func (*ResponseOfferSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{30}
}
func (m *ResponseOfferSnapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseOfferSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseOfferSnapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseOfferSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseOfferSnapshot.Merge(m, src)
}
func (m *ResponseOfferSnapshot) XXX_Size() int {
	return m.Size()
}
func (m *ResponseOfferSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseOfferSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseOfferSnapshot proto.InternalMessageInfo

func (m *ResponseOfferSnapshot) GetResult() ResponseOfferSnapshot_Result {
	if m != nil {
		return m.Result
	}
	return ResponseOfferSnapshot_UNKNOWN
}

type ResponseLoadSnapshotChunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *ResponseLoadSnapshotChunk) Reset()         { *m = ResponseLoadSnapshotChunk{} }
func (m *ResponseLoadSnapshotChunk) String() string { return proto.CompactTextString(m) }
func (*ResponseLoadSnapshotChunk) ProtoMessage()    {}
func (*ResponseLoadSnapshotChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{31}
}
func (m *ResponseLoadSnapshotChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseLoadSnapshotChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseLoadSnapshotChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseLoadSnapshotChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseLoadSnapshotChunk.Merge(m, src)
}
func (m *ResponseLoadSnapshotChunk) XXX_Size() int {
	return m.Size()
}
func (m *ResponseLoadSnapshotChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseLoadSnapshotChunk.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseLoadSnapshotChunk proto.InternalMessageInfo

func (m *ResponseLoadSnapshotChunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type ResponseApplySnapshotChunk struct {
	Result        ResponseApplySnapshotChunk_Result `protobuf:"varint,1,opt,name=result,proto3,enum=tendermint.abci.ResponseApplySnapshotChunk_Result" json:"result,omitempty"`
	RefetchChunks []uint32                          `protobuf:"varint,2,rep,packed,name=refetch_chunks,json=refetchChunks,proto3" json:"refetch_chunks,omitempty"`
	RejectSenders []string                          `protobuf:"bytes,3,rep,name=reject_senders,json=rejectSenders,proto3" json:"reject_senders,omitempty"`
}

func (m *ResponseApplySnapshotChunk) Reset()         { *m = ResponseApplySnapshotChunk{} }
func (m *ResponseApplySnapshotChunk) String() string { return proto.CompactTextString(m) }
func (*ResponseApplySnapshotChunk) ProtoMessage()    {}
func (*ResponseApplySnapshotChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{32}
}
func (m *ResponseApplySnapshotChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseApplySnapshotChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseApplySnapshotChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseApplySnapshotChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseApplySnapshotChunk.Merge(m, src)
}
func (m *ResponseApplySnapshotChunk) XXX_Size() int {
	return m.Size()
}
func (m *ResponseApplySnapshotChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseApplySnapshotChunk.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseApplySnapshotChunk proto.InternalMessageInfo

func (m *ResponseApplySnapshotChunk) GetResult() ResponseApplySnapshotChunk_Result {
	if m != nil {
		return m.Result
	}
	return ResponseApplySnapshotChunk_UNKNOWN
}

func (m *ResponseApplySnapshotChunk) GetRefetchChunks() []uint32 {
	if m != nil {
		return m.RefetchChunks
	}
	return nil
}

func (m *ResponseApplySnapshotChunk) GetRejectSenders() []string {
	if m != nil {
		return m.RejectSenders
	}
	return nil
}

// ConsensusParams contains all consensus-relevant parameters
// that can be adjusted by the abci app
type ConsensusParams struct {
	Block     *BlockParams            `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *types1.EvidenceParams  `protobuf:"bytes,2,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *types1.ValidatorParams `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Version   *types1.VersionParams   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ConsensusParams) Reset()         { *m = ConsensusParams{} }
func (m *ConsensusParams) String() string { return proto.CompactTextString(m) }
func (*ConsensusParams) ProtoMessage()    {}
func (*ConsensusParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{33}
}
func (m *ConsensusParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusParams.Merge(m, src)
}
func (m *ConsensusParams) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusParams.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusParams proto.InternalMessageInfo

func (m *ConsensusParams) GetBlock() *BlockParams {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ConsensusParams) GetEvidence() *types1.EvidenceParams {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *ConsensusParams) GetValidator() *types1.ValidatorParams {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *ConsensusParams) GetVersion() *types1.VersionParams {
	if m != nil {
		return m.Version
	}
	return nil
}

// BlockParams contains limits on the block size.
type BlockParams struct {
	// Note: must be greater than 0
	MaxBytes int64 `protobuf:"varint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Note: must be greater or equal to -1
	MaxGas int64 `protobuf:"varint,2,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
}

func (m *BlockParams) Reset()         { *m = BlockParams{} }
func (m *BlockParams) String() string { return proto.CompactTextString(m) }
func (*BlockParams) ProtoMessage()    {}
func (*BlockParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{34}
}
func (m *BlockParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockParams.Merge(m, src)
}
func (m *BlockParams) XXX_Size() int {
	return m.Size()
}
func (m *BlockParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockParams.DiscardUnknown(m)
}

var xxx_messageInfo_BlockParams proto.InternalMessageInfo

func (m *BlockParams) GetMaxBytes() int64 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *BlockParams) GetMaxGas() int64 {
	if m != nil {
		return m.MaxGas
	}
	return 0
}

type LastCommitInfo struct {
	Round int32      `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Votes []VoteInfo `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes"`
}

func (m *LastCommitInfo) Reset()         { *m = LastCommitInfo{} }
func (m *LastCommitInfo) String() string { return proto.CompactTextString(m) }
func (*LastCommitInfo) ProtoMessage()    {}
func (*LastCommitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{35}
}
func (m *LastCommitInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastCommitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastCommitInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastCommitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastCommitInfo.Merge(m, src)
}
func (m *LastCommitInfo) XXX_Size() int {
	return m.Size()
}
func (m *LastCommitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LastCommitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LastCommitInfo proto.InternalMessageInfo

func (m *LastCommitInfo) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *LastCommitInfo) GetVotes() []VoteInfo {
	if m != nil {
		return m.Votes
	}
	return nil
}

// Event allows application developers to attach additional information to
// ResponseBeginBlock, ResponseEndBlock, ResponseCheckTx and ResponseDeliverTx.
// Later, transactions may be queried using these events.
type Event struct {
	Type       string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []EventAttribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{36}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetAttributes() []EventAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// EventAttribute is a single key-value pair, associated with an event.
type EventAttribute struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Index bool   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *EventAttribute) Reset()         { *m = EventAttribute{} }
func (m *EventAttribute) String() string { return proto.CompactTextString(m) }
func (*EventAttribute) ProtoMessage()    {}
func (*EventAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{37}
}
func (m *EventAttribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAttribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAttribute.Merge(m, src)
}
func (m *EventAttribute) XXX_Size() int {
	return m.Size()
}
func (m *EventAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_EventAttribute proto.InternalMessageInfo

func (m *EventAttribute) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EventAttribute) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EventAttribute) GetIndex() bool {
	if m != nil {
		return m.Index
	}
	return false
}

// TxResult contains results of executing the transaction.
//
// One usage is indexing transaction results.
type TxResult struct {
	Height int64             `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Index  uint32            `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Tx     []byte            `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	Result ResponseDeliverTx `protobuf:"bytes,4,opt,name=result,proto3" json:"result"`
}

func (m *TxResult) Reset()         { *m = TxResult{} }
func (m *TxResult) String() string { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()    {}
func (*TxResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{38}
}
func (m *TxResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResult.Merge(m, src)
}
func (m *TxResult) XXX_Size() int {
	return m.Size()
}
func (m *TxResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResult.DiscardUnknown(m)
}

var xxx_messageInfo_TxResult proto.InternalMessageInfo

func (m *TxResult) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxResult) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxResult) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxResult) GetResult() ResponseDeliverTx {
	if m != nil {
		return m.Result
	}
	return ResponseDeliverTx{}
}

// Validator
type Validator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// PubKey pub_key = 2 [(gogoproto.nullable)=false];
	Power int64 `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{39}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// ValidatorUpdate
type ValidatorUpdate struct {
	PubKey crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key"`
	Power  int64            `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *ValidatorUpdate) Reset()         { *m = ValidatorUpdate{} }
func (m *ValidatorUpdate) String() string { return proto.CompactTextString(m) }
func (*ValidatorUpdate) ProtoMessage()    {}
func (*ValidatorUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{40}
}
func (m *ValidatorUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorUpdate.Merge(m, src)
}
func (m *ValidatorUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorUpdate proto.InternalMessageInfo

func (m *ValidatorUpdate) GetPubKey() crypto.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return crypto.PublicKey{}
}

func (m *ValidatorUpdate) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// VoteInfo
type VoteInfo struct {
	Validator       Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator"`
	SignedLastBlock bool      `protobuf:"varint,2,opt,name=signed_last_block,json=signedLastBlock,proto3" json:"signed_last_block,omitempty"`
}

func (m *VoteInfo) Reset()         { *m = VoteInfo{} }
func (m *VoteInfo) String() string { return proto.CompactTextString(m) }
func (*VoteInfo) ProtoMessage()    {}
func (*VoteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{41}
}
func (m *VoteInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteInfo.Merge(m, src)
}
func (m *VoteInfo) XXX_Size() int {
	return m.Size()
}
func (m *VoteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VoteInfo proto.InternalMessageInfo

func (m *VoteInfo) GetValidator() Validator {
	if m != nil {
		return m.Validator
	}
	return Validator{}
}

func (m *VoteInfo) GetSignedLastBlock() bool {
	if m != nil {
		return m.SignedLastBlock
	}
	return false
}

type Evidence struct {
	Type      string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Validator Validator `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator"`
	Height    int64     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time      time.Time `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	// Total voting power of the validator set in case the ABCI application does
	// not store historical validators.
	// https://github.com/tendermint/tendermint/issues/4581
	TotalVotingPower int64 `protobuf:"varint,5,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{42}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Evidence) GetValidator() Validator {
	if m != nil {
		return m.Validator
	}
	return Validator{}
}

func (m *Evidence) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Evidence) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Evidence) GetTotalVotingPower() int64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

type Snapshot struct {
	Height   uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Format   uint32 `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Chunks   uint32 `protobuf:"varint,3,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Hash     []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_252557cfdd89a31a, []int{43}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Snapshot) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *Snapshot) GetChunks() uint32 {
	if m != nil {
		return m.Chunks
	}
	return 0
}

func (m *Snapshot) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Snapshot) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("tendermint.abci.CheckTxType", CheckTxType_name, CheckTxType_value)
	proto.RegisterEnum("tendermint.abci.ResponseOfferSnapshot_Result", ResponseOfferSnapshot_Result_name, ResponseOfferSnapshot_Result_value)
	proto.RegisterEnum("tendermint.abci.ResponseApplySnapshotChunk_Result", ResponseApplySnapshotChunk_Result_name, ResponseApplySnapshotChunk_Result_value)
	proto.RegisterType((*Request)(nil), "tendermint.abci.Request")
	proto.RegisterType((*RequestEcho)(nil), "tendermint.abci.RequestEcho")
	proto.RegisterType((*RequestFlush)(nil), "tendermint.abci.RequestFlush")
	proto.RegisterType((*RequestInfo)(nil), "tendermint.abci.RequestInfo")
	proto.RegisterType((*RequestSetOption)(nil), "tendermint.abci.RequestSetOption")
	proto.RegisterType((*RequestInitChain)(nil), "tendermint.abci.RequestInitChain")
	proto.RegisterType((*RequestQuery)(nil), "tendermint.abci.RequestQuery")
	proto.RegisterType((*RequestBeginBlock)(nil), "tendermint.abci.RequestBeginBlock")
	proto.RegisterType((*RequestCheckTx)(nil), "tendermint.abci.RequestCheckTx")
	proto.RegisterType((*RequestDeliverTx)(nil), "tendermint.abci.RequestDeliverTx")
	proto.RegisterType((*RequestEndBlock)(nil), "tendermint.abci.RequestEndBlock")
	proto.RegisterType((*RequestCommit)(nil), "tendermint.abci.RequestCommit")
	proto.RegisterType((*RequestListSnapshots)(nil), "tendermint.abci.RequestListSnapshots")
	proto.RegisterType((*RequestOfferSnapshot)(nil), "tendermint.abci.RequestOfferSnapshot")
	proto.RegisterType((*RequestLoadSnapshotChunk)(nil), "tendermint.abci.RequestLoadSnapshotChunk")
	proto.RegisterType((*RequestApplySnapshotChunk)(nil), "tendermint.abci.RequestApplySnapshotChunk")
	proto.RegisterType((*Response)(nil), "tendermint.abci.Response")
	proto.RegisterType((*ResponseException)(nil), "tendermint.abci.ResponseException")
	proto.RegisterType((*ResponseEcho)(nil), "tendermint.abci.ResponseEcho")
	proto.RegisterType((*ResponseFlush)(nil), "tendermint.abci.ResponseFlush")
	proto.RegisterType((*ResponseInfo)(nil), "tendermint.abci.ResponseInfo")
	proto.RegisterType((*ResponseSetOption)(nil), "tendermint.abci.ResponseSetOption")
	proto.RegisterType((*ResponseInitChain)(nil), "tendermint.abci.ResponseInitChain")
	proto.RegisterType((*ResponseQuery)(nil), "tendermint.abci.ResponseQuery")
	proto.RegisterType((*ResponseBeginBlock)(nil), "tendermint.abci.ResponseBeginBlock")
	proto.RegisterType((*ResponseCheckTx)(nil), "tendermint.abci.ResponseCheckTx")
	proto.RegisterType((*ResponseDeliverTx)(nil), "tendermint.abci.ResponseDeliverTx")
	proto.RegisterType((*ResponseEndBlock)(nil), "tendermint.abci.ResponseEndBlock")
	proto.RegisterType((*ResponseCommit)(nil), "tendermint.abci.ResponseCommit")
	proto.RegisterType((*ResponseListSnapshots)(nil), "tendermint.abci.ResponseListSnapshots")
	proto.RegisterType((*ResponseOfferSnapshot)(nil), "tendermint.abci.ResponseOfferSnapshot")
	proto.RegisterType((*ResponseLoadSnapshotChunk)(nil), "tendermint.abci.ResponseLoadSnapshotChunk")
	proto.RegisterType((*ResponseApplySnapshotChunk)(nil), "tendermint.abci.ResponseApplySnapshotChunk")
	proto.RegisterType((*ConsensusParams)(nil), "tendermint.abci.ConsensusParams")
	proto.RegisterType((*BlockParams)(nil), "tendermint.abci.BlockParams")
	proto.RegisterType((*LastCommitInfo)(nil), "tendermint.abci.LastCommitInfo")
	proto.RegisterType((*Event)(nil), "tendermint.abci.Event")
	proto.RegisterType((*EventAttribute)(nil), "tendermint.abci.EventAttribute")
	proto.RegisterType((*TxResult)(nil), "tendermint.abci.TxResult")
	proto.RegisterType((*Validator)(nil), "tendermint.abci.Validator")
	proto.RegisterType((*ValidatorUpdate)(nil), "tendermint.abci.ValidatorUpdate")
	proto.RegisterType((*VoteInfo)(nil), "tendermint.abci.VoteInfo")
	proto.RegisterType((*Evidence)(nil), "tendermint.abci.Evidence")
	proto.RegisterType((*Snapshot)(nil), "tendermint.abci.Snapshot")
}

func init() { proto.RegisterFile("tendermint/abci/types.proto", fileDescriptor_252557cfdd89a31a) }

var fileDescriptor_252557cfdd89a31a = []byte{
	// 2761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x5a, 0x4b, 0x77, 0x23, 0xc5,
	0xf5, 0xd7, 0x5b, 0xea, 0x6b, 0x5b, 0x92, 0x6b, 0xcc, 0x20, 0x7a, 0x66, 0xec, 0xa1, 0x38, 0xbc,
	0x06, 0xb0, 0xff, 0x98, 0x03, 0x7f, 0x5e, 0x09, 0x58, 0x42, 0x13, 0x0d, 0x03, 0xf6, 0xa4, 0x6d,
	0x86, 0xbc, 0x98, 0xa6, 0xa4, 0x2e, 0x4b, 0xcd, 0x48, 0xdd, 0x8d, 0xba, 0x64, 0x6c, 0x96, 0x24,
	0xd9, 0x90, 0x0d, 0x39, 0xd9, 0x64, 0xc3, 0x2e, 0x8b, 0x7c, 0x81, 0x9c, 0x93, 0x55, 0xd6, 0x2c,
	0x39, 0x27, 0x9b, 0xac, 0x48, 0x0e, 0x24, 0x9b, 0x7c, 0x81, 0x2c, 0x93, 0x53, 0x8f, 0x7e, 0x49,
	0x6a, 0x4b, 0x86, 0xec, 0xb2, 0xab, 0xc7, 0xbd, 0xb7, 0xfa, 0x96, 0xea, 0xfe, 0xea, 0x77, 0x6f,
	0x09, 0xae, 0x30, 0xea, 0x58, 0x74, 0x3c, 0xb2, 0x1d, 0xb6, 0x43, 0xba, 0x3d, 0x7b, 0x87, 0x9d,
	0x79, 0xd4, 0xdf, 0xf6, 0xc6, 0x2e, 0x73, 0x51, 0x2d, 0x9a, 0xdc, 0xe6, 0x93, 0xfa, 0xb5, 0x98,
	0x74, 0x6f, 0x7c, 0xe6, 0x31, 0x77, 0xc7, 0x1b, 0xbb, 0xee, 0xb1, 0x94, 0xd7, 0xaf, 0xc6, 0xa6,
	0x85, 0x9d, 0xb8, 0xb5, 0xc4, 0xac, 0x52, 0xbe, 0x4f, 0xcf, 0x82, 0xd9, 0x6b, 0x33, 0xba, 0x1e,
	0x19, 0x93, 0x51, 0x30, 0xbd, 0xd5, 0x77, 0xdd, 0xfe, 0x90, 0xee, 0x88, 0x5e, 0x77, 0x72, 0xbc,
	0xc3, 0xec, 0x11, 0xf5, 0x19, 0x19, 0x79, 0x4a, 0xe0, 0xf1, 0x98, 0xbe, 0x92, 0x25, 0x9e, 0xbd,
	0x43, 0x1c, 0xc7, 0x65, 0x84, 0xd9, 0xae, 0x13, 0x58, 0xda, 0xe8, 0xbb, 0x7d, 0x57, 0x34, 0x77,
	0x78, 0x4b, 0x8e, 0xe2, 0x5f, 0x57, 0xa0, 0x6c, 0xd0, 0x0f, 0x27, 0xd4, 0x67, 0x68, 0x17, 0x0a,
	0xb4, 0x37, 0x70, 0x1b, 0xd9, 0xeb, 0xd9, 0x27, 0x56, 0x76, 0xaf, 0x6e, 0x4f, 0xed, 0xc2, 0xb6,
	0x92, 0x6b, 0xf7, 0x06, 0x6e, 0x27, 0x63, 0x08, 0x59, 0xf4, 0x3c, 0x14, 0x8f, 0x87, 0x13, 0x7f,
	0xd0, 0xc8, 0x09, 0xa5, 0x6b, 0x69, 0x4a, 0x37, 0xb9, 0x50, 0x27, 0x63, 0x48, 0x69, 0xbe, 0x94,
	0xed, 0x1c, 0xbb, 0x8d, 0xfc, 0xf9, 0x4b, 0xdd, 0x72, 0x8e, 0xc5, 0x52, 0x5c, 0x16, 0x35, 0x01,
	0x7c, 0xca, 0x4c, 0xd7, 0xe3, 0x5e, 0x35, 0x0a, 0x42, 0xf3, 0xe1, 0x34, 0xcd, 0x43, 0xca, 0x0e,
	0x84, 0x60, 0x27, 0x63, 0x68, 0x7e, 0xd0, 0xe1, 0x36, 0x6c, 0xc7, 0x66, 0x66, 0x6f, 0x40, 0x6c,
	0xa7, 0x51, 0x3c, 0xdf, 0xc6, 0x2d, 0xc7, 0x66, 0x2d, 0x2e, 0xc8, 0x6d, 0xd8, 0x41, 0x87, 0xbb,
	0xfc, 0xe1, 0x84, 0x8e, 0xcf, 0x1a, 0xa5, 0xf3, 0x5d, 0xfe, 0x21, 0x17, 0xe2, 0x2e, 0x0b, 0x69,
	0xd4, 0x86, 0x95, 0x2e, 0xed, 0xdb, 0x8e, 0xd9, 0x1d, 0xba, 0xbd, 0xfb, 0x8d, 0xb2, 0x50, 0xc6,
	0x69, 0xca, 0x4d, 0x2e, 0xda, 0xe4, 0x92, 0x9d, 0x8c, 0x01, 0xdd, 0xb0, 0x87, 0x5e, 0x85, 0x4a,
	0x6f, 0x40, 0x7b, 0xf7, 0x4d, 0x76, 0xda, 0xa8, 0x08, 0x1b, 0x5b, 0x69, 0x36, 0x5a, 0x5c, 0xee,
	0xe8, 0xb4, 0x93, 0x31, 0xca, 0x3d, 0xd9, 0xe4, 0xfe, 0x5b, 0x74, 0x68, 0x9f, 0xd0, 0x31, 0xd7,
	0xd7, 0xce, 0xf7, 0xff, 0x0d, 0x29, 0x29, 0x2c, 0x68, 0x56, 0xd0, 0x41, 0xaf, 0x81, 0x46, 0x1d,
	0x4b, 0xb9, 0x01, 0xc2, 0xc4, 0xf5, 0xd4, 0xb3, 0xe2, 0x58, 0x81, 0x13, 0x15, 0xaa, 0xda, 0xe8,
	0x45, 0x28, 0xf5, 0xdc, 0xd1, 0xc8, 0x66, 0x8d, 0x15, 0xa1, 0xbd, 0x99, 0xea, 0x80, 0x90, 0xea,
	0x64, 0x0c, 0x25, 0x8f, 0xf6, 0xa1, 0x3a, 0xb4, 0x7d, 0x66, 0xfa, 0x0e, 0xf1, 0xfc, 0x81, 0xcb,
	0xfc, 0xc6, 0xaa, 0xb0, 0xf0, 0x68, 0x9a, 0x85, 0xb7, 0x6c, 0x9f, 0x1d, 0x06, 0xc2, 0x9d, 0x8c,
	0xb1, 0x36, 0x8c, 0x0f, 0x70, 0x7b, 0xee, 0xf1, 0x31, 0x1d, 0x87, 0x06, 0x1b, 0x6b, 0xe7, 0xdb,
	0x3b, 0xe0, 0xd2, 0x81, 0x3e, 0xb7, 0xe7, 0xc6, 0x07, 0xd0, 0x4f, 0xe1, 0xd2, 0xd0, 0x25, 0x56,
	0x68, 0xce, 0xec, 0x0d, 0x26, 0xce, 0xfd, 0x46, 0x55, 0x18, 0x7d, 0x32, 0xf5, 0x23, 0x5d, 0x62,
	0x05, 0x26, 0x5a, 0x5c, 0xa1, 0x93, 0x31, 0xd6, 0x87, 0xd3, 0x83, 0xe8, 0x1e, 0x6c, 0x10, 0xcf,
	0x1b, 0x9e, 0x4d, 0x5b, 0xaf, 0x09, 0xeb, 0x37, 0xd2, 0xac, 0xef, 0x71, 0x9d, 0x69, 0xf3, 0x88,
	0xcc, 0x8c, 0x36, 0xcb, 0x50, 0x3c, 0x21, 0xc3, 0x09, 0xc5, 0x8f, 0xc3, 0x4a, 0x2c, 0xd4, 0x51,
	0x03, 0xca, 0x23, 0xea, 0xfb, 0xa4, 0x4f, 0x05, 0x32, 0x68, 0x46, 0xd0, 0xc5, 0x55, 0x58, 0x8d,
	0x87, 0x37, 0x1e, 0x85, 0x8a, 0x3c, 0x70, 0xb9, 0xe2, 0x09, 0x1d, 0xfb, 0x3c, 0x5a, 0x95, 0xa2,
	0xea, 0xa2, 0x47, 0x60, 0x4d, 0x1c, 0x1f, 0x33, 0x98, 0xe7, 0xe8, 0x51, 0x30, 0x56, 0xc5, 0xe0,
	0x5d, 0x25, 0xb4, 0x05, 0x2b, 0xde, 0xae, 0x17, 0x8a, 0xe4, 0x85, 0x08, 0x78, 0xbb, 0x9e, 0x12,
	0xc0, 0x2f, 0x43, 0x7d, 0x3a, 0xda, 0x51, 0x1d, 0xf2, 0xf7, 0xe9, 0x99, 0x5a, 0x8f, 0x37, 0xd1,
	0x86, 0x72, 0x4b, 0xac, 0xa1, 0x19, 0xca, 0xc7, 0xdf, 0xe5, 0x42, 0xe5, 0x30, 0xcc, 0xd1, 0x8b,
	0x50, 0xe0, 0xf0, 0xaa, 0x00, 0x50, 0xdf, 0x96, 0x78, 0xba, 0x1d, 0x60, 0xef, 0xf6, 0x51, 0x80,
	0xbd, 0xcd, 0xca, 0x17, 0x5f, 0x6d, 0x65, 0x3e, 0xfb, 0xeb, 0x56, 0xd6, 0x10, 0x1a, 0xe8, 0x21,
	0x1e, 0x95, 0xc4, 0x76, 0x4c, 0xdb, 0x52, 0xeb, 0x94, 0x45, 0xff, 0x96, 0x85, 0x6e, 0x43, 0xbd,
	0xe7, 0x3a, 0x3e, 0x75, 0xfc, 0x89, 0x6f, 0x4a, 0x6c, 0x57, 0xb0, 0x37, 0x1b, 0x35, 0xad, 0x40,
	0xf0, 0x8e, 0x90, 0x33, 0x6a, 0xbd, 0xe4, 0x00, 0xba, 0x09, 0x70, 0x42, 0x86, 0xb6, 0x45, 0x98,
	0x3b, 0xf6, 0x1b, 0x85, 0xeb, 0xf9, 0xb9, 0x66, 0xee, 0x06, 0x22, 0xef, 0x78, 0x16, 0x61, 0xb4,
	0x59, 0xe0, 0x5f, 0x6b, 0xc4, 0x34, 0xd1, 0x63, 0x50, 0x23, 0x9e, 0x67, 0xfa, 0x8c, 0x30, 0x6a,
	0x76, 0xcf, 0x18, 0xf5, 0x05, 0x18, 0xae, 0x1a, 0x6b, 0xc4, 0xf3, 0x0e, 0xf9, 0x68, 0x93, 0x0f,
	0x62, 0x2b, 0xfc, 0x85, 0x05, 0x9a, 0x21, 0x04, 0x05, 0x8b, 0x30, 0x22, 0x76, 0x68, 0xd5, 0x10,
	0x6d, 0x3e, 0xe6, 0x11, 0x36, 0x50, 0x7e, 0x8b, 0x36, 0xba, 0x0c, 0xa5, 0x01, 0xb5, 0xfb, 0x03,
	0x26, 0x5c, 0xcd, 0x1b, 0xaa, 0xc7, 0x7f, 0x0c, 0x6f, 0xec, 0x9e, 0x50, 0x01, 0xdf, 0x15, 0x43,
	0x76, 0xf0, 0x2f, 0x72, 0xb0, 0x3e, 0x83, 0x7b, 0xdc, 0xee, 0x80, 0xf8, 0x83, 0x60, 0x2d, 0xde,
	0x46, 0x2f, 0x70, 0xbb, 0xc4, 0xa2, 0x63, 0x75, 0xdf, 0x34, 0xe2, 0xbe, 0xcb, 0x4b, 0xb7, 0x23,
	0xe6, 0x95, 0xcf, 0x4a, 0x1a, 0x1d, 0x40, 0x7d, 0x48, 0x7c, 0x66, 0x4a, 0x1c, 0x31, 0x63, 0x77,
	0xcf, 0x2c, 0x7a, 0xbe, 0x45, 0x02, 0xe4, 0xe1, 0xa7, 0x58, 0x19, 0xaa, 0x0e, 0x13, 0xa3, 0xc8,
	0x80, 0x8d, 0xee, 0xd9, 0xc7, 0xc4, 0x61, 0xb6, 0x43, 0xcd, 0x99, 0x9f, 0xe4, 0xa1, 0x19, 0xa3,
	0xed, 0x13, 0xdb, 0xa2, 0x4e, 0x2f, 0xf8, 0x2d, 0x2e, 0x85, 0xca, 0xe1, 0x6f, 0xe5, 0x63, 0x03,
	0xaa, 0x49, 0xe4, 0x46, 0x55, 0xc8, 0xb1, 0x53, 0xb5, 0x01, 0x39, 0x76, 0x8a, 0xfe, 0x0f, 0x0a,
	0xdc, 0x49, 0xe1, 0x7c, 0x75, 0xce, 0xb5, 0xa9, 0xf4, 0x8e, 0xce, 0x3c, 0x6a, 0x08, 0x49, 0x8c,
	0xc3, 0x63, 0x1e, 0xa2, 0xf9, 0xb4, 0x55, 0xfc, 0x24, 0xd4, 0xa6, 0xe0, 0x3a, 0xf6, 0xfb, 0x65,
	0xe3, 0xbf, 0x1f, 0xae, 0xc1, 0x5a, 0x02, 0x9b, 0xf1, 0x65, 0xd8, 0x98, 0x07, 0xb5, 0x78, 0x10,
	0x8e, 0x27, 0x20, 0x13, 0x3d, 0x0f, 0x95, 0x10, 0x6b, 0x65, 0x98, 0xcd, 0xee, 0x55, 0x20, 0x6c,
	0x84, 0xa2, 0x3c, 0xbe, 0xf8, 0x79, 0x15, 0xe7, 0x21, 0x27, 0x3e, 0xbc, 0x4c, 0x3c, 0xaf, 0x43,
	0xfc, 0x01, 0x7e, 0x1f, 0x1a, 0x69, 0x38, 0x3a, 0xe5, 0x46, 0x21, 0x3c, 0x86, 0x97, 0xa1, 0x74,
	0xec, 0x8e, 0x47, 0x84, 0x09, 0x63, 0x6b, 0x86, 0xea, 0xf1, 0xe3, 0x29, 0x31, 0x35, 0x2f, 0x86,
	0x65, 0x07, 0x9b, 0xf0, 0x50, 0x2a, 0x96, 0x72, 0x15, 0xdb, 0xb1, 0xa8, 0xdc, 0xcf, 0x35, 0x43,
	0x76, 0x22, 0x43, 0xf2, 0x63, 0x65, 0x87, 0x2f, 0xeb, 0x0b, 0x5f, 0x85, 0x7d, 0xcd, 0x50, 0x3d,
	0xfc, 0x8f, 0x0a, 0x54, 0x0c, 0xea, 0x7b, 0x3c, 0xd8, 0x51, 0x13, 0x34, 0x7a, 0xda, 0xa3, 0x92,
	0xe5, 0x64, 0x53, 0x59, 0x82, 0x94, 0x6e, 0x07, 0x92, 0xfc, 0x8a, 0x0e, 0xd5, 0xd0, 0x73, 0x8a,
	0xc9, 0xa5, 0x93, 0x32, 0xa5, 0x1e, 0xa7, 0x72, 0x2f, 0x04, 0x54, 0x2e, 0x9f, 0x7a, 0x2b, 0x4b,
	0xad, 0x29, 0x2e, 0xf7, 0x9c, 0xe2, 0x72, 0x85, 0x05, 0x8b, 0x25, 0xc8, 0x5c, 0x2b, 0x41, 0xe6,
	0x8a, 0x0b, 0xdc, 0x4c, 0x61, 0x73, 0xad, 0x04, 0x9b, 0x2b, 0x2d, 0x30, 0x92, 0x42, 0xe7, 0x5e,
	0x08, 0xe8, 0x5c, 0x79, 0x81, 0xdb, 0x53, 0x7c, 0xee, 0x66, 0x92, 0xcf, 0x49, 0x2e, 0xf6, 0x48,
	0xaa, 0x76, 0x2a, 0xa1, 0xfb, 0x5e, 0x8c, 0xd0, 0x69, 0xa9, 0x6c, 0x4a, 0x1a, 0x99, 0xc3, 0xe8,
	0x5a, 0x09, 0x46, 0x07, 0x0b, 0xf6, 0x20, 0x85, 0xd2, 0xbd, 0x1e, 0xa7, 0x74, 0x2b, 0xa9, 0xac,
	0x50, 0x1d, 0x9a, 0x79, 0x9c, 0xee, 0xa5, 0x90, 0xd3, 0xad, 0xa6, 0x92, 0x52, 0xe5, 0xc3, 0x34,
	0xa9, 0x3b, 0x98, 0x21, 0x75, 0x92, 0x84, 0x3d, 0x96, 0x6a, 0x62, 0x01, 0xab, 0x3b, 0x98, 0x61,
	0x75, 0xd5, 0x05, 0x06, 0x17, 0xd0, 0xba, 0x9f, 0xcd, 0xa7, 0x75, 0xe9, 0xc4, 0x4b, 0x7d, 0xe6,
	0x72, 0xbc, 0xce, 0x4c, 0xe1, 0x75, 0x75, 0x61, 0xfe, 0xa9, 0x54, 0xf3, 0x17, 0x27, 0x76, 0x4f,
	0xf2, 0x6b, 0x76, 0x0a, 0x38, 0x38, 0x54, 0xd1, 0xf1, 0xd8, 0x1d, 0x2b, 0xce, 0x24, 0x3b, 0xf8,
	0x09, 0x7e, 0xf1, 0x47, 0x20, 0x71, 0x0e, 0x09, 0x14, 0x57, 0x42, 0x0c, 0x18, 0xf0, 0x1f, 0xb3,
	0x91, 0xae, 0xb8, 0x2b, 0xe3, 0xa4, 0x41, 0x53, 0xa4, 0x21, 0xc6, 0x0d, 0x73, 0x49, 0x6e, 0xb8,
	0x05, 0x2b, 0x1c, 0xea, 0xa7, 0x68, 0x1f, 0xf1, 0x02, 0xda, 0x87, 0x6e, 0xc0, 0xba, 0xb8, 0xcb,
	0x25, 0x83, 0x54, 0xf8, 0x5e, 0x10, 0xd7, 0x54, 0x8d, 0x4f, 0xc8, 0xc3, 0x29, 0x81, 0xfe, 0x19,
	0xb8, 0x14, 0x93, 0x0d, 0xaf, 0x10, 0xc9, 0x75, 0xea, 0xa1, 0xf4, 0x9e, 0xba, 0x4b, 0xde, 0x8e,
	0x36, 0x28, 0xa2, 0x94, 0x08, 0x0a, 0x3d, 0xd7, 0xa2, 0x0a, 0xe0, 0x45, 0x9b, 0xd3, 0xcc, 0xa1,
	0xdb, 0x57, 0x30, 0xce, 0x9b, 0x5c, 0x2a, 0x44, 0x41, 0x4d, 0x82, 0x1c, 0xfe, 0x7d, 0x36, 0xb2,
	0x17, 0xb1, 0xcc, 0x79, 0x84, 0x30, 0xfb, 0xdf, 0x21, 0x84, 0xb9, 0x6f, 0x4b, 0x08, 0xf1, 0xbf,
	0xb2, 0xd1, 0xcf, 0x18, 0x52, 0xbd, 0x6f, 0xe7, 0x76, 0x74, 0x25, 0x16, 0xc5, 0x8f, 0xa2, 0xae,
	0x44, 0xc5, 0xcc, 0x4b, 0x62, 0xeb, 0x93, 0xcc, 0xbc, 0x2c, 0x2f, 0x49, 0xd1, 0x41, 0x2f, 0x82,
	0x26, 0x6a, 0x2b, 0xa6, 0xeb, 0xf9, 0x0a, 0x55, 0xaf, 0xc4, 0x1d, 0x92, 0x25, 0x94, 0xed, 0x3b,
	0x5c, 0xe6, 0xc0, 0xf3, 0x8d, 0x8a, 0xa7, 0x5a, 0xb1, 0xdb, 0x5e, 0x4b, 0x90, 0xce, 0xab, 0xa0,
	0xf1, 0xaf, 0xf7, 0x3d, 0xd2, 0xa3, 0x02, 0x21, 0x35, 0x23, 0x1a, 0xc0, 0xf7, 0x00, 0xcd, 0x62,
	0x34, 0xea, 0x40, 0x89, 0x9e, 0x50, 0x87, 0xf1, 0x9f, 0x86, 0xef, 0xe9, 0xe5, 0x39, 0x8c, 0x8e,
	0x3a, 0xac, 0xd9, 0xe0, 0x3b, 0xf9, 0xcf, 0xaf, 0xb6, 0xea, 0x52, 0xfa, 0x69, 0x77, 0x64, 0x33,
	0x3a, 0xf2, 0xd8, 0x99, 0xa1, 0xf4, 0xf1, 0xcf, 0x73, 0x9c, 0x5e, 0x25, 0xf0, 0x7b, 0xee, 0xde,
	0x06, 0x51, 0x92, 0x8b, 0x51, 0xeb, 0xe5, 0xf6, 0x7b, 0x13, 0xa0, 0x4f, 0x7c, 0xf3, 0x23, 0xe2,
	0x30, 0x6a, 0xa9, 0x4d, 0x8f, 0x8d, 0x20, 0x1d, 0x2a, 0xbc, 0x37, 0xf1, 0xa9, 0x25, 0xb6, 0x3f,
	0x6f, 0x84, 0xfd, 0x98, 0x9f, 0xe5, 0xef, 0xe6, 0x67, 0x72, 0x97, 0x2b, 0xd3, 0xbb, 0xfc, 0xcb,
	0x5c, 0x14, 0x0a, 0x11, 0x13, 0xfd, 0xdf, 0xdb, 0x87, 0x5f, 0x89, 0xbc, 0x33, 0x79, 0x91, 0xa2,
	0x43, 0x58, 0x0f, 0x43, 0xd1, 0x9c, 0x88, 0x10, 0x0d, 0xce, 0xdd, 0xb2, 0xb1, 0x5c, 0x3f, 0x49,
	0x0e, 0xfb, 0xe8, 0x47, 0xf0, 0xe0, 0x14, 0xcc, 0x84, 0xa6, 0x73, 0x4b, 0xa2, 0xcd, 0x03, 0x49,
	0xb4, 0x09, 0x2c, 0x47, 0x7b, 0x95, 0xff, 0x8e, 0xb1, 0x71, 0x8b, 0x67, 0x3c, 0x71, 0x5a, 0x30,
	0xf7, 0xd7, 0x7f, 0x04, 0xd6, 0xc6, 0x94, 0xf1, 0xec, 0x3a, 0x91, 0x53, 0xae, 0xca, 0x41, 0x89,
	0xf4, 0xf8, 0x0e, 0x3c, 0x30, 0x97, 0x1e, 0xa0, 0xff, 0x07, 0x2d, 0x62, 0x16, 0xd9, 0x94, 0xf4,
	0x2c, 0x4c, 0x39, 0x22, 0x59, 0xfc, 0xa7, 0x6c, 0x64, 0x32, 0x99, 0xc4, 0xb4, 0xa1, 0x34, 0xa6,
	0xfe, 0x64, 0x28, 0xd3, 0x8a, 0xea, 0xee, 0x33, 0xcb, 0x11, 0x0b, 0x3e, 0x3a, 0x19, 0x32, 0x43,
	0x29, 0xe3, 0x7b, 0x50, 0x92, 0x23, 0x68, 0x05, 0xca, 0xef, 0xec, 0xdf, 0xde, 0x3f, 0x78, 0x77,
	0xbf, 0x9e, 0x41, 0x00, 0xa5, 0xbd, 0x56, 0xab, 0x7d, 0xe7, 0xa8, 0x9e, 0x45, 0x1a, 0x14, 0xf7,
	0x9a, 0x07, 0xc6, 0x51, 0x3d, 0xc7, 0x87, 0x8d, 0xf6, 0x9b, 0xed, 0xd6, 0x51, 0x3d, 0x8f, 0xd6,
	0x61, 0x4d, 0xb6, 0xcd, 0x9b, 0x07, 0xc6, 0xdb, 0x7b, 0x47, 0xf5, 0x42, 0x6c, 0xe8, 0xb0, 0xbd,
	0xff, 0x46, 0xdb, 0xa8, 0x17, 0xf1, 0xb3, 0x3c, 0x6f, 0x49, 0xa1, 0x22, 0x51, 0x86, 0x92, 0x8d,
	0x65, 0x28, 0xf8, 0xb7, 0x39, 0xd0, 0xd3, 0xf9, 0x05, 0x7a, 0x73, 0xca, 0xf1, 0xdd, 0x0b, 0x90,
	0x93, 0x29, 0xef, 0xd1, 0xa3, 0x50, 0x1d, 0xd3, 0x63, 0xca, 0x7a, 0x03, 0xc9, 0x77, 0xe4, 0xed,
	0xb5, 0x66, 0xac, 0xa9, 0x51, 0xa1, 0xe4, 0x4b, 0xb1, 0x0f, 0x68, 0x8f, 0x99, 0x32, 0x59, 0x92,
	0x87, 0x4e, 0xe3, 0x62, 0x7c, 0xf4, 0x50, 0x0e, 0xe2, 0xf7, 0x2f, 0xb4, 0x97, 0x1a, 0x14, 0x8d,
	0xf6, 0x91, 0xf1, 0xe3, 0x7a, 0x1e, 0x21, 0xa8, 0x8a, 0xa6, 0x79, 0xb8, 0xbf, 0x77, 0xe7, 0xb0,
	0x73, 0xc0, 0xf7, 0xf2, 0x12, 0xd4, 0x82, 0xbd, 0x0c, 0x06, 0x8b, 0xf8, 0xdf, 0x59, 0xa8, 0x4d,
	0x05, 0x08, 0xda, 0x85, 0xa2, 0xe4, 0xcc, 0x69, 0x25, 0x73, 0x11, 0xdf, 0x2a, 0x9a, 0xa4, 0x28,
	0x7a, 0x15, 0x2a, 0x54, 0x15, 0x03, 0xe6, 0x05, 0xa2, 0x2c, 0x62, 0x04, 0xe5, 0x02, 0xa5, 0x1a,
	0x6a, 0xa0, 0xd7, 0x40, 0x0b, 0x23, 0x5d, 0x25, 0x6a, 0x0f, 0xcf, 0xaa, 0x87, 0x18, 0xa1, 0xf4,
	0x23, 0x1d, 0xf4, 0x52, 0x44, 0xbc, 0x0a, 0xb3, 0x4c, 0x5d, 0xa9, 0x4b, 0x01, 0xa5, 0x1c, 0xc8,
	0xe3, 0x16, 0xac, 0xc4, 0xfc, 0x41, 0x57, 0x40, 0x1b, 0x91, 0x53, 0x55, 0x3d, 0x92, 0x65, 0x82,
	0xca, 0x88, 0x9c, 0x8a, 0xc2, 0x11, 0x7a, 0x10, 0xca, 0x7c, 0xb2, 0x4f, 0x24, 0xda, 0xe4, 0x8d,
	0xd2, 0x88, 0x9c, 0xfe, 0x80, 0xf8, 0xf8, 0x3d, 0xa8, 0x26, 0x0b, 0x2c, 0xfc, 0x24, 0x8e, 0xdd,
	0x89, 0x63, 0x09, 0x1b, 0x45, 0x43, 0x76, 0xd0, 0xf3, 0x50, 0x3c, 0x71, 0x25, 0x58, 0xcd, 0x0f,
	0xd9, 0xbb, 0x2e, 0xa3, 0xb1, 0x02, 0x8d, 0x94, 0xc6, 0x1f, 0x43, 0x51, 0x80, 0x0f, 0x07, 0x12,
	0x51, 0x2a, 0x51, 0xa4, 0x93, 0xb7, 0xd1, 0x7b, 0x00, 0x84, 0xb1, 0xb1, 0xdd, 0x9d, 0x44, 0x86,
	0xb7, 0xe6, 0x83, 0xd7, 0x5e, 0x20, 0xd7, 0xbc, 0xaa, 0x50, 0x6c, 0x23, 0x52, 0x8d, 0x21, 0x59,
	0xcc, 0x20, 0xde, 0x87, 0x6a, 0x52, 0x37, 0x5e, 0x8d, 0x5c, 0x9d, 0x53, 0x8d, 0x0c, 0x39, 0x4f,
	0xc8, 0x98, 0xf2, 0xb2, 0x2c, 0x26, 0x3a, 0xf8, 0xd3, 0x2c, 0x54, 0x8e, 0x4e, 0xd5, 0xb1, 0x4e,
	0xa9, 0xc8, 0x44, 0xaa, 0xb9, 0x78, 0xfd, 0x41, 0x96, 0x78, 0xf2, 0x61, 0xe1, 0xe8, 0xf5, 0x30,
	0x70, 0x0b, 0xcb, 0x66, 0x88, 0x41, 0x05, 0x4d, 0x81, 0xd5, 0x2b, 0xa0, 0x85, 0xa7, 0x8a, 0xb3,
	0x77, 0x62, 0x59, 0x63, 0xea, 0xfb, 0xca, 0xb7, 0xa0, 0x2b, 0x0a, 0x7c, 0xee, 0x47, 0xaa, 0xc2,
	0x91, 0x37, 0x64, 0x07, 0x5b, 0x50, 0x9b, 0xba, 0xb6, 0xd0, 0x2b, 0x50, 0xf6, 0x26, 0x5d, 0x33,
	0xd8, 0x9e, 0xa9, 0xe0, 0x09, 0x48, 0xde, 0xa4, 0x3b, 0xb4, 0x7b, 0xb7, 0xe9, 0x59, 0xf0, 0x31,
	0xde, 0xa4, 0x7b, 0x5b, 0xee, 0xa2, 0x5c, 0x25, 0x17, 0x5f, 0xe5, 0x04, 0x2a, 0xc1, 0xa1, 0x40,
	0xdf, 0x8f, 0xc7, 0x49, 0x50, 0xcf, 0x4d, 0xbd, 0x4a, 0x95, 0xf9, 0x58, 0x98, 0xdc, 0x80, 0x75,
	0xdf, 0xee, 0x3b, 0xd4, 0x32, 0xa3, 0xfc, 0x41, 0xac, 0x56, 0x31, 0x6a, 0x72, 0xe2, 0xad, 0x20,
	0x79, 0xc0, 0x5f, 0x65, 0xa1, 0x12, 0x04, 0xec, 0xdc, 0x73, 0x97, 0xf8, 0x98, 0xdc, 0xc5, 0x3f,
	0x26, 0xad, 0x9a, 0x1a, 0xd4, 0xab, 0x0b, 0x17, 0xae, 0x57, 0x3f, 0x0d, 0x88, 0xb9, 0x8c, 0x0c,
	0xcd, 0x13, 0x97, 0xd9, 0x4e, 0xdf, 0x94, 0xbb, 0x29, 0x29, 0x53, 0x5d, 0xcc, 0xdc, 0x15, 0x13,
	0x77, 0xc4, 0xc6, 0x7e, 0x92, 0x85, 0x4a, 0x78, 0xf9, 0x5d, 0xb4, 0xa6, 0x76, 0x19, 0x4a, 0x0a,
	0xdf, 0x65, 0x51, 0x4d, 0xf5, 0xc2, 0xf2, 0x6e, 0x21, 0x56, 0xde, 0xd5, 0xa1, 0x32, 0xa2, 0x8c,
	0x08, 0x06, 0x20, 0x73, 0xb4, 0xb0, 0x7f, 0xe3, 0x25, 0x58, 0x89, 0x95, 0x37, 0x79, 0x68, 0xed,
	0xb7, 0xdf, 0xad, 0x67, 0xf4, 0xf2, 0xa7, 0x9f, 0x5f, 0xcf, 0xef, 0xd3, 0x8f, 0xf8, 0xa1, 0x34,
	0xda, 0xad, 0x4e, 0xbb, 0x75, 0xbb, 0x9e, 0xd5, 0x57, 0x3e, 0xfd, 0xfc, 0x7a, 0xd9, 0xa0, 0xa2,
	0x4a, 0xb2, 0xfb, 0x87, 0x15, 0xa8, 0xed, 0x35, 0x5b, 0xb7, 0xf8, 0xa5, 0x64, 0xf7, 0x88, 0xaa,
	0x1d, 0x15, 0x44, 0x62, 0x7b, 0xee, 0x33, 0xa7, 0x7e, 0x7e, 0xe9, 0x0c, 0xdd, 0x84, 0xa2, 0xc8,
	0x79, 0xd1, 0xf9, 0xef, 0x9e, 0xfa, 0x82, 0x5a, 0x1a, 0xff, 0x18, 0x71, 0x6a, 0xcf, 0x7d, 0x08,
	0xd5, 0xcf, 0x2f, 0xad, 0x21, 0x03, 0xb4, 0x28, 0x69, 0x5d, 0xfc, 0x30, 0xaa, 0x2f, 0x51, 0x6e,
	0x43, 0x1f, 0x82, 0x16, 0xb1, 0xf5, 0xc5, 0x0f, 0x85, 0xfa, 0x12, 0xb8, 0x82, 0xaf, 0x7d, 0xf2,
	0xe7, 0xbf, 0xff, 0x26, 0xf7, 0x20, 0x46, 0xf2, 0xe1, 0xfd, 0xe4, 0xd9, 0x9d, 0x70, 0xee, 0xe5,
	0xec, 0x0d, 0x64, 0x43, 0x39, 0x48, 0x93, 0x16, 0xbd, 0x6c, 0xea, 0x0b, 0x2b, 0x65, 0xf8, 0x8a,
	0x58, 0xec, 0x01, 0x5c, 0x0f, 0x17, 0x53, 0x33, 0x7c, 0xa9, 0x9b, 0x50, 0x94, 0xb9, 0xee, 0xf9,
	0x6f, 0xb8, 0xfa, 0x82, 0x9a, 0x20, 0xa2, 0x50, 0x52, 0xf4, 0x75, 0xc1, 0x53, 0xa6, 0xbe, 0xa8,
	0x2c, 0x86, 0x75, 0xf1, 0xbd, 0x1b, 0xb8, 0x16, 0x7d, 0xaf, 0x98, 0xe0, 0x9f, 0xeb, 0x80, 0x16,
	0x55, 0x11, 0x16, 0xbf, 0x5a, 0xeb, 0x4b, 0x94, 0x42, 0x83, 0xf5, 0x50, 0xf4, 0x63, 0x44, 0x4b,
	0x4c, 0x00, 0x62, 0x19, 0xf1, 0x12, 0x4f, 0xd5, 0xfa, 0x32, 0xe5, 0x4f, 0xbc, 0x29, 0x96, 0x6c,
	0xe0, 0x4b, 0xe1, 0x92, 0xd1, 0x24, 0x77, 0x73, 0x04, 0x95, 0x30, 0x33, 0x5a, 0xf8, 0xb0, 0xac,
	0x2f, 0xae, 0x53, 0xe2, 0xab, 0x62, 0xc1, 0xcb, 0x78, 0x3d, 0x5c, 0x30, 0x98, 0xe2, 0xcb, 0xdd,
	0x83, 0xb5, 0x64, 0xc2, 0xb0, 0xdc, 0x63, 0xb2, 0xbe, 0x64, 0x79, 0x92, 0xdb, 0x4f, 0x66, 0x0f,
	0xcb, 0x3d, 0x2e, 0xeb, 0x4b, 0x56, 0x2b, 0xd1, 0x07, 0xb0, 0x3e, 0xcb, 0xee, 0x97, 0x7f, 0x6b,
	0xd6, 0x2f, 0x50, 0xbf, 0x44, 0x23, 0x40, 0x73, 0xb2, 0x82, 0x0b, 0x3c, 0x3d, 0xeb, 0x17, 0x29,
	0x67, 0x36, 0xdb, 0x5f, 0x7c, 0xbd, 0x99, 0xfd, 0xf2, 0xeb, 0xcd, 0xec, 0xdf, 0xbe, 0xde, 0xcc,
	0x7e, 0xf6, 0xcd, 0x66, 0xe6, 0xcb, 0x6f, 0x36, 0x33, 0x7f, 0xf9, 0x66, 0x33, 0xf3, 0x93, 0xa7,
	0xfa, 0x36, 0x1b, 0x4c, 0xba, 0xdb, 0x3d, 0x77, 0xb4, 0x43, 0x2c, 0x77, 0x6c, 0x13, 0xdf, 0x3d,
	0x66, 0x3b, 0x73, 0xff, 0xd4, 0xd3, 0x2d, 0x89, 0x0b, 0xf1, 0xb9, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x32, 0xc0, 0x40, 0xf4, 0x23, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ABCIApplicationClient is the client API for ABCIApplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ABCIApplicationClient interface {
	Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error)
	Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error)
	Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error)
	SetOption(ctx context.Context, in *RequestSetOption, opts ...grpc.CallOption) (*ResponseSetOption, error)
	DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error)
	CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error)
	Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error)
	Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error)
	InitChain(ctx context.Context, in *RequestInitChain, opts ...grpc.CallOption) (*ResponseInitChain, error)
	BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error)
	EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error)
	ListSnapshots(ctx context.Context, in *RequestListSnapshots, opts ...grpc.CallOption) (*ResponseListSnapshots, error)
	OfferSnapshot(ctx context.Context, in *RequestOfferSnapshot, opts ...grpc.CallOption) (*ResponseOfferSnapshot, error)
	LoadSnapshotChunk(ctx context.Context, in *RequestLoadSnapshotChunk, opts ...grpc.CallOption) (*ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(ctx context.Context, in *RequestApplySnapshotChunk, opts ...grpc.CallOption) (*ResponseApplySnapshotChunk, error)
}

type aBCIApplicationClient struct {
	cc *grpc.ClientConn
}

func NewABCIApplicationClient(cc *grpc.ClientConn) ABCIApplicationClient {
	return &aBCIApplicationClient{cc}
}

func (c *aBCIApplicationClient) Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error) {
	out := new(ResponseEcho)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error) {
	out := new(ResponseFlush)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error) {
	out := new(ResponseInfo)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) SetOption(ctx context.Context, in *RequestSetOption, opts ...grpc.CallOption) (*ResponseSetOption, error) {
	out := new(ResponseSetOption)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/SetOption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error) {
	out := new(ResponseDeliverTx)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/DeliverTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error) {
	out := new(ResponseCheckTx)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/CheckTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error) {
	out := new(ResponseQuery)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error) {
	out := new(ResponseCommit)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) InitChain(ctx context.Context, in *RequestInitChain, opts ...grpc.CallOption) (*ResponseInitChain, error) {
	out := new(ResponseInitChain)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/InitChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error) {
	out := new(ResponseBeginBlock)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/BeginBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error) {
	out := new(ResponseEndBlock)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/EndBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) ListSnapshots(ctx context.Context, in *RequestListSnapshots, opts ...grpc.CallOption) (*ResponseListSnapshots, error) {
	out := new(ResponseListSnapshots)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/ListSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) OfferSnapshot(ctx context.Context, in *RequestOfferSnapshot, opts ...grpc.CallOption) (*ResponseOfferSnapshot, error) {
	out := new(ResponseOfferSnapshot)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/OfferSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) LoadSnapshotChunk(ctx context.Context, in *RequestLoadSnapshotChunk, opts ...grpc.CallOption) (*ResponseLoadSnapshotChunk, error) {
	out := new(ResponseLoadSnapshotChunk)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/LoadSnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIApplicationClient) ApplySnapshotChunk(ctx context.Context, in *RequestApplySnapshotChunk, opts ...grpc.CallOption) (*ResponseApplySnapshotChunk, error) {
	out := new(ResponseApplySnapshotChunk)
	err := c.cc.Invoke(ctx, "/tendermint.abci.ABCIApplication/ApplySnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABCIApplicationServer is the server API for ABCIApplication service.
type ABCIApplicationServer interface {
	Echo(context.Context, *RequestEcho) (*ResponseEcho, error)
	Flush(context.Context, *RequestFlush) (*ResponseFlush, error)
	Info(context.Context, *RequestInfo) (*ResponseInfo, error)
	SetOption(context.Context, *RequestSetOption) (*ResponseSetOption, error)
	DeliverTx(context.Context, *RequestDeliverTx) (*ResponseDeliverTx, error)
	CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error)
	Query(context.Context, *RequestQuery) (*ResponseQuery, error)
	Commit(context.Context, *RequestCommit) (*ResponseCommit, error)
	InitChain(context.Context, *RequestInitChain) (*ResponseInitChain, error)
	BeginBlock(context.Context, *RequestBeginBlock) (*ResponseBeginBlock, error)
	EndBlock(context.Context, *RequestEndBlock) (*ResponseEndBlock, error)
	ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error)
	OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error)
	LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error)
}

// UnimplementedABCIApplicationServer can be embedded to have forward compatible implementations.
type UnimplementedABCIApplicationServer struct {
}

func (*UnimplementedABCIApplicationServer) Echo(ctx context.Context, req *RequestEcho) (*ResponseEcho, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedABCIApplicationServer) Flush(ctx context.Context, req *RequestFlush) (*ResponseFlush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (*UnimplementedABCIApplicationServer) Info(ctx context.Context, req *RequestInfo) (*ResponseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedABCIApplicationServer) SetOption(ctx context.Context, req *RequestSetOption) (*ResponseSetOption, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOption not implemented")
}
func (*UnimplementedABCIApplicationServer) DeliverTx(ctx context.Context, req *RequestDeliverTx) (*ResponseDeliverTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTx not implemented")
}
func (*UnimplementedABCIApplicationServer) CheckTx(ctx context.Context, req *RequestCheckTx) (*ResponseCheckTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTx not implemented")
}
func (*UnimplementedABCIApplicationServer) Query(ctx context.Context, req *RequestQuery) (*ResponseQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedABCIApplicationServer) Commit(ctx context.Context, req *RequestCommit) (*ResponseCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedABCIApplicationServer) InitChain(ctx context.Context, req *RequestInitChain) (*ResponseInitChain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitChain not implemented")
}
func (*UnimplementedABCIApplicationServer) BeginBlock(ctx context.Context, req *RequestBeginBlock) (*ResponseBeginBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginBlock not implemented")
}
func (*UnimplementedABCIApplicationServer) EndBlock(ctx context.Context, req *RequestEndBlock) (*ResponseEndBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndBlock not implemented")
}
func (*UnimplementedABCIApplicationServer) ListSnapshots(ctx context.Context, req *RequestListSnapshots) (*ResponseListSnapshots, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (*UnimplementedABCIApplicationServer) OfferSnapshot(ctx context.Context, req *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfferSnapshot not implemented")
}
func (*UnimplementedABCIApplicationServer) LoadSnapshotChunk(ctx context.Context, req *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSnapshotChunk not implemented")
}
func (*UnimplementedABCIApplicationServer) ApplySnapshotChunk(ctx context.Context, req *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySnapshotChunk not implemented")
}

func RegisterABCIApplicationServer(s *grpc.Server, srv ABCIApplicationServer) {
	s.RegisterService(&_ABCIApplication_serviceDesc, srv)
}

func _ABCIApplication_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEcho)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Echo(ctx, req.(*RequestEcho))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFlush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Flush(ctx, req.(*RequestFlush))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Info(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_SetOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).SetOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/SetOption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).SetOption(ctx, req.(*RequestSetOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_DeliverTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeliverTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).DeliverTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/DeliverTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).DeliverTx(ctx, req.(*RequestDeliverTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_CheckTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).CheckTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/CheckTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).CheckTx(ctx, req.(*RequestCheckTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Query(ctx, req.(*RequestQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCommit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).Commit(ctx, req.(*RequestCommit))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_InitChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInitChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).InitChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/InitChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).InitChain(ctx, req.(*RequestInitChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_BeginBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBeginBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).BeginBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/BeginBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).BeginBlock(ctx, req.(*RequestBeginBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_EndBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEndBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).EndBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/EndBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).EndBlock(ctx, req.(*RequestEndBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListSnapshots)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).ListSnapshots(ctx, req.(*RequestListSnapshots))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_OfferSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOfferSnapshot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).OfferSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/OfferSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).OfferSnapshot(ctx, req.(*RequestOfferSnapshot))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_LoadSnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLoadSnapshotChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).LoadSnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/LoadSnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).LoadSnapshotChunk(ctx, req.(*RequestLoadSnapshotChunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIApplication_ApplySnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestApplySnapshotChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIApplicationServer).ApplySnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.abci.ABCIApplication/ApplySnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIApplicationServer).ApplySnapshotChunk(ctx, req.(*RequestApplySnapshotChunk))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABCIApplication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.abci.ABCIApplication",
	HandlerType: (*ABCIApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ABCIApplication_Echo_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _ABCIApplication_Flush_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ABCIApplication_Info_Handler,
		},
		{
			MethodName: "SetOption",
			Handler:    _ABCIApplication_SetOption_Handler,
		},
		{
			MethodName: "DeliverTx",
			Handler:    _ABCIApplication_DeliverTx_Handler,
		},
		{
			MethodName: "CheckTx",
			Handler:    _ABCIApplication_CheckTx_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ABCIApplication_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _ABCIApplication_Commit_Handler,
		},
		{
			MethodName: "InitChain",
			Handler:    _ABCIApplication_InitChain_Handler,
		},
		{
			MethodName: "BeginBlock",
			Handler:    _ABCIApplication_BeginBlock_Handler,
		},
		{
			MethodName: "EndBlock",
			Handler:    _ABCIApplication_EndBlock_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _ABCIApplication_ListSnapshots_Handler,
		},
		{
			MethodName: "OfferSnapshot",
			Handler:    _ABCIApplication_OfferSnapshot_Handler,
		},
		{
			MethodName: "LoadSnapshotChunk",
			Handler:    _ABCIApplication_LoadSnapshotChunk_Handler,
		},
		{
			MethodName: "ApplySnapshotChunk",
			Handler:    _ABCIApplication_ApplySnapshotChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tendermint/abci/types.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Request_Echo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Echo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Echo != nil {
		{
			size, err := m.Echo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Request_Flush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Flush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Flush != nil {
		{
			size, err := m.Flush.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Request_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Request_SetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_SetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SetOption != nil {
		{
			size, err := m.SetOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Request_InitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_InitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitChain != nil {
		{
			size, err := m.InitChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Request_Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Request_BeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_BeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BeginBlock != nil {
		{
			size, err := m.BeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Request_CheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_CheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CheckTx != nil {
		{
			size, err := m.CheckTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Request_DeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_DeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeliverTx != nil {
		{
			size, err := m.DeliverTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *Request_EndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_EndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndBlock != nil {
		{
			size, err := m.EndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *Request_Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *Request_ListSnapshots) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_ListSnapshots) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ListSnapshots != nil {
		{
			size, err := m.ListSnapshots.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *Request_OfferSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_OfferSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OfferSnapshot != nil {
		{
			size, err := m.OfferSnapshot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	return len(dAtA) - i, nil
}
func (m *Request_LoadSnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_LoadSnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LoadSnapshotChunk != nil {
		{
			size, err := m.LoadSnapshotChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	return len(dAtA) - i, nil
}
func (m *Request_ApplySnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request_ApplySnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ApplySnapshotChunk != nil {
		{
			size, err := m.ApplySnapshotChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	return len(dAtA) - i, nil
}
func (m *RequestEcho) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestEcho) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestEcho) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestFlush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestFlush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestFlush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RequestInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.P2PVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.P2PVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.BlockVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockVersion))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestSetOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestSetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestSetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestInitChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppStateBytes) > 0 {
		i -= len(m.AppStateBytes)
		copy(dAtA[i:], m.AppStateBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppStateBytes)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ConsensusParams != nil {
		{
			size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	n17, err17 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err17 != nil {
		return 0, err17
	}
	i -= n17
	i = encodeVarintTypes(dAtA, i, uint64(n17))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RequestQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Prove {
		i--
		if m.Prove {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestBeginBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestBeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestBeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ByzantineValidators) > 0 {
		for iNdEx := len(m.ByzantineValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ByzantineValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.LastCommitInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestCheckTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestDeliverTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestDeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestDeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestEndBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestEndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestEndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RequestCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RequestListSnapshots) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestListSnapshots) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestListSnapshots) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RequestOfferSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestOfferSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestOfferSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Snapshot != nil {
		{
			size, err := m.Snapshot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RequestLoadSnapshotChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestLoadSnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestLoadSnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Chunk != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Chunk))
		i--
		dAtA[i] = 0x18
	}
	if m.Format != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RequestApplySnapshotChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestApplySnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestApplySnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chunk) > 0 {
		i -= len(m.Chunk)
		copy(dAtA[i:], m.Chunk)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chunk)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size := m.Value.Size()
			i -= size
			if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Response_Exception) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Exception) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Exception != nil {
		{
			size, err := m.Exception.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Response_Echo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Echo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Echo != nil {
		{
			size, err := m.Echo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Response_Flush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Flush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Flush != nil {
		{
			size, err := m.Flush.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Response_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Response_SetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_SetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SetOption != nil {
		{
			size, err := m.SetOption.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Response_InitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_InitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InitChain != nil {
		{
			size, err := m.InitChain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Response_Query) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Query) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func (m *Response_BeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_BeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BeginBlock != nil {
		{
			size, err := m.BeginBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Response_CheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_CheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CheckTx != nil {
		{
			size, err := m.CheckTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *Response_DeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_DeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DeliverTx != nil {
		{
			size, err := m.DeliverTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	return len(dAtA) - i, nil
}
func (m *Response_EndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_EndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndBlock != nil {
		{
			size, err := m.EndBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	return len(dAtA) - i, nil
}
func (m *Response_Commit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_Commit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	return len(dAtA) - i, nil
}
func (m *Response_ListSnapshots) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_ListSnapshots) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ListSnapshots != nil {
		{
			size, err := m.ListSnapshots.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	return len(dAtA) - i, nil
}
func (m *Response_OfferSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_OfferSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OfferSnapshot != nil {
		{
			size, err := m.OfferSnapshot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	return len(dAtA) - i, nil
}
func (m *Response_LoadSnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_LoadSnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.LoadSnapshotChunk != nil {
		{
			size, err := m.LoadSnapshotChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x7a
	}
	return len(dAtA) - i, nil
}
func (m *Response_ApplySnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response_ApplySnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ApplySnapshotChunk != nil {
		{
			size, err := m.ApplySnapshotChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	return len(dAtA) - i, nil
}
func (m *ResponseException) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseException) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseException) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseEcho) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseEcho) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseEcho) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseFlush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseFlush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseFlush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ResponseInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastBlockAppHash) > 0 {
		i -= len(m.LastBlockAppHash)
		copy(dAtA[i:], m.LastBlockAppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastBlockAppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.AppVersion != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AppVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseSetOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseSetOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseSetOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseInitChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseInitChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseInitChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ConsensusParams != nil {
		{
			size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x52
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x48
	}
	if m.ProofOps != nil {
		{
			size, err := m.ProofOps.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x32
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseBeginBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseBeginBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseBeginBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseCheckTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseCheckTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseCheckTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.GasUsed != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.GasWanted != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseDeliverTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseDeliverTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseDeliverTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codespace) > 0 {
		i -= len(m.Codespace)
		copy(dAtA[i:], m.Codespace)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Codespace)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.GasUsed != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.GasWanted != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GasWanted))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseEndBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseEndBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseEndBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ConsensusParamUpdates != nil {
		{
			size, err := m.ConsensusParamUpdates.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorUpdates) > 0 {
		for iNdEx := len(m.ValidatorUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseCommit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseCommit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseCommit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RetainHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RetainHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ResponseListSnapshots) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseListSnapshots) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseListSnapshots) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Snapshots) > 0 {
		for iNdEx := len(m.Snapshots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Snapshots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ResponseOfferSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseOfferSnapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseOfferSnapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseLoadSnapshotChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseLoadSnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseLoadSnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chunk) > 0 {
		i -= len(m.Chunk)
		copy(dAtA[i:], m.Chunk)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chunk)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponseApplySnapshotChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseApplySnapshotChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseApplySnapshotChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RejectSenders) > 0 {
		for iNdEx := len(m.RejectSenders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RejectSenders[iNdEx])
			copy(dAtA[i:], m.RejectSenders[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.RejectSenders[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RefetchChunks) > 0 {
		dAtA41 := make([]byte, len(m.RefetchChunks)*10)
		var j40 int
		for _, num := range m.RefetchChunks {
			for num >= 1<<7 {
				dAtA41[j40] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j40++
			}
			dAtA41[j40] = uint8(num)
			j40++
		}
		i -= j40
		copy(dAtA[i:], dAtA41[:j40])
		i = encodeVarintTypes(dAtA, i, uint64(j40))
		i--
		dAtA[i] = 0x12
	}
	if m.Result != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxGas != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxGas))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxBytes != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxBytes))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LastCommitInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastCommitInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastCommitInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAttribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAttribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAttribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index {
		i--
		if m.Index {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VoteInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignedLastBlock {
		i--
		if m.SignedLastBlock {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalVotingPower != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x28
	}
	n49, err49 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err49 != nil {
		return 0, err49
	}
	i -= n49
	i = encodeVarintTypes(dAtA, i, uint64(n49))
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Chunks != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Chunks))
		i--
		dAtA[i] = 0x18
	}
	if m.Format != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Request_Echo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Echo != nil {
		l = m.Echo.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Flush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Flush != nil {
		l = m.Flush.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_SetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SetOption != nil {
		l = m.SetOption.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_InitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitChain != nil {
		l = m.InitChain.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_BeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginBlock != nil {
		l = m.BeginBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_CheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckTx != nil {
		l = m.CheckTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_DeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeliverTx != nil {
		l = m.DeliverTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_EndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndBlock != nil {
		l = m.EndBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_ListSnapshots) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListSnapshots != nil {
		l = m.ListSnapshots.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_OfferSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OfferSnapshot != nil {
		l = m.OfferSnapshot.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_LoadSnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LoadSnapshotChunk != nil {
		l = m.LoadSnapshotChunk.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Request_ApplySnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplySnapshotChunk != nil {
		l = m.ApplySnapshotChunk.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *RequestEcho) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RequestFlush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RequestInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BlockVersion != 0 {
		n += 1 + sovTypes(uint64(m.BlockVersion))
	}
	if m.P2PVersion != 0 {
		n += 1 + sovTypes(uint64(m.P2PVersion))
	}
	return n
}

func (m *RequestSetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RequestInitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ConsensusParams != nil {
		l = m.ConsensusParams.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.AppStateBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RequestQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Prove {
		n += 2
	}
	return n
}

func (m *RequestBeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Header.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.LastCommitInfo.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.ByzantineValidators) > 0 {
		for _, e := range m.ByzantineValidators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *RequestCheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	return n
}

func (m *RequestDeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RequestEndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	return n
}

func (m *RequestCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RequestListSnapshots) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RequestOfferSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Snapshot != nil {
		l = m.Snapshot.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RequestLoadSnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Format != 0 {
		n += 1 + sovTypes(uint64(m.Format))
	}
	if m.Chunk != 0 {
		n += 1 + sovTypes(uint64(m.Chunk))
	}
	return n
}

func (m *RequestApplySnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = len(m.Chunk)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Response_Exception) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Exception != nil {
		l = m.Exception.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Echo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Echo != nil {
		l = m.Echo.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Flush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Flush != nil {
		l = m.Flush.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_SetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SetOption != nil {
		l = m.SetOption.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_InitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitChain != nil {
		l = m.InitChain.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_BeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BeginBlock != nil {
		l = m.BeginBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_CheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckTx != nil {
		l = m.CheckTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_DeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DeliverTx != nil {
		l = m.DeliverTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_EndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndBlock != nil {
		l = m.EndBlock.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_Commit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_ListSnapshots) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListSnapshots != nil {
		l = m.ListSnapshots.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_OfferSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OfferSnapshot != nil {
		l = m.OfferSnapshot.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_LoadSnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LoadSnapshotChunk != nil {
		l = m.LoadSnapshotChunk.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Response_ApplySnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplySnapshotChunk != nil {
		l = m.ApplySnapshotChunk.Size()
		n += 2 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ResponseException) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseEcho) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseFlush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ResponseInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.AppVersion != 0 {
		n += 1 + sovTypes(uint64(m.AppVersion))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.LastBlockHeight))
	}
	l = len(m.LastBlockAppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseSetOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseInitChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConsensusParams != nil {
		l = m.ConsensusParams.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ResponseQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ProofOps != nil {
		l = m.ProofOps.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseBeginBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ResponseCheckTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.GasWanted != 0 {
		n += 1 + sovTypes(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTypes(uint64(m.GasUsed))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseDeliverTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTypes(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.GasWanted != 0 {
		n += 1 + sovTypes(uint64(m.GasWanted))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTypes(uint64(m.GasUsed))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Codespace)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseEndBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorUpdates) > 0 {
		for _, e := range m.ValidatorUpdates {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ConsensusParamUpdates != nil {
		l = m.ConsensusParamUpdates.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ResponseCommit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.RetainHeight != 0 {
		n += 1 + sovTypes(uint64(m.RetainHeight))
	}
	return n
}

func (m *ResponseListSnapshots) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Snapshots) > 0 {
		for _, e := range m.Snapshots {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ResponseOfferSnapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovTypes(uint64(m.Result))
	}
	return n
}

func (m *ResponseLoadSnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chunk)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponseApplySnapshotChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovTypes(uint64(m.Result))
	}
	if len(m.RefetchChunks) > 0 {
		l = 0
		for _, e := range m.RefetchChunks {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	if len(m.RejectSenders) > 0 {
		for _, s := range m.RejectSenders {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ConsensusParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BlockParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxBytes != 0 {
		n += 1 + sovTypes(uint64(m.MaxBytes))
	}
	if m.MaxGas != 0 {
		n += 1 + sovTypes(uint64(m.MaxGas))
	}
	return n
}

func (m *LastCommitInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *EventAttribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Index {
		n += 2
	}
	return n
}

func (m *TxResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Result.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	return n
}

func (m *ValidatorUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PubKey.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	return n
}

func (m *VoteInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Validator.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.SignedLastBlock {
		n += 2
	}
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Validator.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	if m.TotalVotingPower != 0 {
		n += 1 + sovTypes(uint64(m.TotalVotingPower))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	if m.Format != 0 {
		n += 1 + sovTypes(uint64(m.Format))
	}
	if m.Chunks != 0 {
		n += 1 + sovTypes(uint64(m.Chunks))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Echo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestEcho{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Echo{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flush", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestFlush{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Flush{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Info{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestSetOption{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_SetOption{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestInitChain{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_InitChain{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Query{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestBeginBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_BeginBlock{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestCheckTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_CheckTx{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestDeliverTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_DeliverTx{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestEndBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_EndBlock{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestCommit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_Commit{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListSnapshots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestListSnapshots{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_ListSnapshots{v}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferSnapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestOfferSnapshot{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_OfferSnapshot{v}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadSnapshotChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestLoadSnapshotChunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_LoadSnapshotChunk{v}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplySnapshotChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RequestApplySnapshotChunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Request_ApplySnapshotChunk{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestEcho) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestEcho: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestEcho: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestFlush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestFlush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestFlush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockVersion", wireType)
			}
			m.BlockVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field P2PVersion", wireType)
			}
			m.P2PVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.P2PVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestSetOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestSetOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestSetOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestInitChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestInitChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInitChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParams == nil {
				m.ConsensusParams = &ConsensusParams{}
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, ValidatorUpdate{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppStateBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppStateBytes = append(m.AppStateBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.AppStateBytes == nil {
				m.AppStateBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prove", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Prove = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestBeginBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestBeginBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestBeginBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastCommitInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByzantineValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ByzantineValidators = append(m.ByzantineValidators, Evidence{})
			if err := m.ByzantineValidators[len(m.ByzantineValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestCheckTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestCheckTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCheckTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CheckTxType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestDeliverTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestDeliverTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestDeliverTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestEndBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestEndBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestEndBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestListSnapshots) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestListSnapshots: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestListSnapshots: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestOfferSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestOfferSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestOfferSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Snapshot == nil {
				m.Snapshot = &Snapshot{}
			}
			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestLoadSnapshotChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestLoadSnapshotChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestLoadSnapshotChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			m.Chunk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chunk |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestApplySnapshotChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestApplySnapshotChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestApplySnapshotChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunk = append(m.Chunk[:0], dAtA[iNdEx:postIndex]...)
			if m.Chunk == nil {
				m.Chunk = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exception", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseException{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Exception{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Echo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseEcho{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Echo{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flush", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseFlush{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Flush{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Info{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetOption", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseSetOption{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_SetOption{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseInitChain{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_InitChain{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseQuery{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Query{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseBeginBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_BeginBlock{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseCheckTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_CheckTx{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseDeliverTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_DeliverTx{v}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseEndBlock{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_EndBlock{v}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseCommit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_Commit{v}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListSnapshots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseListSnapshots{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_ListSnapshots{v}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferSnapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseOfferSnapshot{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_OfferSnapshot{v}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadSnapshotChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseLoadSnapshotChunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_LoadSnapshotChunk{v}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplySnapshotChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResponseApplySnapshotChunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Response_ApplySnapshotChunk{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseException) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseException: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseException: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseEcho) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseEcho: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseEcho: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseFlush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseFlush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseFlush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppVersion", wireType)
			}
			m.AppVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastBlockAppHash = append(m.LastBlockAppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastBlockAppHash == nil {
				m.LastBlockAppHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseSetOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseSetOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseSetOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseInitChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseInitChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseInitChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParams == nil {
				m.ConsensusParams = &ConsensusParams{}
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, ValidatorUpdate{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofOps == nil {
				m.ProofOps = &crypto.ProofOps{}
			}
			if err := m.ProofOps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseBeginBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseBeginBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseBeginBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseCheckTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseCheckTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseCheckTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseDeliverTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseDeliverTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseDeliverTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
			}
			m.GasWanted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasWanted |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseEndBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseEndBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseEndBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorUpdates = append(m.ValidatorUpdates, ValidatorUpdate{})
			if err := m.ValidatorUpdates[len(m.ValidatorUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParamUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsensusParamUpdates == nil {
				m.ConsensusParamUpdates = &ConsensusParams{}
			}
			if err := m.ConsensusParamUpdates.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseCommit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseCommit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseCommit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetainHeight", wireType)
			}
			m.RetainHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetainHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseListSnapshots) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseListSnapshots: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseListSnapshots: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Snapshots = append(m.Snapshots, &Snapshot{})
			if err := m.Snapshots[len(m.Snapshots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseOfferSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseOfferSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseOfferSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= ResponseOfferSnapshot_Result(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseLoadSnapshotChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseLoadSnapshotChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseLoadSnapshotChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunk = append(m.Chunk[:0], dAtA[iNdEx:postIndex]...)
			if m.Chunk == nil {
				m.Chunk = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseApplySnapshotChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseApplySnapshotChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseApplySnapshotChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= ResponseApplySnapshotChunk_Result(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RefetchChunks = append(m.RefetchChunks, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.RefetchChunks) == 0 {
					m.RefetchChunks = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RefetchChunks = append(m.RefetchChunks, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RefetchChunks", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectSenders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RejectSenders = append(m.RejectSenders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ConsensusParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ConsensusParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &BlockParams{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &types1.EvidenceParams{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &types1.ValidatorParams{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &types1.VersionParams{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *BlockParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BlockParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBytes", wireType)
			}
			m.MaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
			}
			m.MaxGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGas |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LastCommitInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LastCommitInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastCommitInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, VoteInfo{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, EventAttribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *EventAttribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: EventAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Index = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *TxResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TxResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ValidatorUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ValidatorUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *VoteInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: VoteInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedLastBlock", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.SignedLastBlock = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			m.Chunks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chunks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata[:0], dAtA[iNdEx:postIndex]...)
			if m.Metadata == nil {
				m.Metadata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
