// Code generated by protoc-gen-go.
// source: append_entries_request.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	append_entries_request.proto
	append_entries_responses.proto
	log_entry.proto
	request_vote_request.proto
	request_vote_responses.proto
	snapshot_recovery_request.proto
	snapshot_recovery_response.proto
	snapshot_request.proto
	snapshot_response.proto

It has these top-level messages:
	AppendEntriesRequest
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AppendEntriesRequest struct {
	Term             *uint64     `protobuf:"varint,1,req" json:"Term,omitempty"`
	PrevLogIndex     *uint64     `protobuf:"varint,2,req" json:"PrevLogIndex,omitempty"`
	PrevLogTerm      *uint64     `protobuf:"varint,3,req" json:"PrevLogTerm,omitempty"`
	CommitIndex      *uint64     `protobuf:"varint,4,req" json:"CommitIndex,omitempty"`
	LeaderName       *string     `protobuf:"bytes,5,req" json:"LeaderName,omitempty"`
	Entries          []*LogEntry `protobuf:"bytes,6,rep" json:"Entries,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *AppendEntriesRequest) Reset()         { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()    {}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil && m.PrevLogIndex != nil {
		return *m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil && m.PrevLogTerm != nil {
		return *m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
}
