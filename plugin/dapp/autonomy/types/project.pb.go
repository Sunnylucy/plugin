// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AutonomyProposalProject struct {
	PropProject *ProposalProject `protobuf:"bytes,1,opt,name=propProject" json:"propProject,omitempty"`
	// 董事会投票结果
	BoardVoteRes *VoteResult `protobuf:"bytes,2,opt,name=boardVoteRes" json:"boardVoteRes,omitempty"`
	// 公示投票
	PubVote *PublicVote `protobuf:"bytes,3,opt,name=pubVote" json:"pubVote,omitempty"`
	// 状态
	Status  int32  `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Height  int64  `protobuf:"varint,6,opt,name=height" json:"height,omitempty"`
	Index   int32  `protobuf:"varint,7,opt,name=index" json:"index,omitempty"`
}

func (m *AutonomyProposalProject) Reset()                    { *m = AutonomyProposalProject{} }
func (m *AutonomyProposalProject) String() string            { return proto.CompactTextString(m) }
func (*AutonomyProposalProject) ProtoMessage()               {}
func (*AutonomyProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *AutonomyProposalProject) GetPropProject() *ProposalProject {
	if m != nil {
		return m.PropProject
	}
	return nil
}

func (m *AutonomyProposalProject) GetBoardVoteRes() *VoteResult {
	if m != nil {
		return m.BoardVoteRes
	}
	return nil
}

func (m *AutonomyProposalProject) GetPubVote() *PublicVote {
	if m != nil {
		return m.PubVote
	}
	return nil
}

func (m *AutonomyProposalProject) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AutonomyProposalProject) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AutonomyProposalProject) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AutonomyProposalProject) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ProposalProject struct {
	// 提案时间
	Year  int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
	// 项目相关
	FirstStage   string `protobuf:"bytes,4,opt,name=firstStage" json:"firstStage,omitempty"`
	LastStage    string `protobuf:"bytes,5,opt,name=lastStage" json:"lastStage,omitempty"`
	Production   string `protobuf:"bytes,6,opt,name=production" json:"production,omitempty"`
	Description  string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Contractor   string `protobuf:"bytes,8,opt,name=contractor" json:"contractor,omitempty"`
	Amount       int64  `protobuf:"varint,9,opt,name=amount" json:"amount,omitempty"`
	AmountDetail string `protobuf:"bytes,10,opt,name=amountDetail" json:"amountDetail,omitempty"`
	// 支付相关
	ToAddr string `protobuf:"bytes,11,opt,name=toAddr" json:"toAddr,omitempty"`
	// 投票相关
	StartBlockHeight    int64 `protobuf:"varint,12,opt,name=startBlockHeight" json:"startBlockHeight,omitempty"`
	EndBlockHeight      int64 `protobuf:"varint,13,opt,name=endBlockHeight" json:"endBlockHeight,omitempty"`
	RealEndBlockHeight  int64 `protobuf:"varint,14,opt,name=realEndBlockHeight" json:"realEndBlockHeight,omitempty"`
	ProjectNeedBlockNum int32 `protobuf:"varint,15,opt,name=projectNeedBlockNum" json:"projectNeedBlockNum,omitempty"`
}

func (m *ProposalProject) Reset()                    { *m = ProposalProject{} }
func (m *ProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ProposalProject) ProtoMessage()               {}
func (*ProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ProposalProject) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ProposalProject) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *ProposalProject) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *ProposalProject) GetFirstStage() string {
	if m != nil {
		return m.FirstStage
	}
	return ""
}

func (m *ProposalProject) GetLastStage() string {
	if m != nil {
		return m.LastStage
	}
	return ""
}

func (m *ProposalProject) GetProduction() string {
	if m != nil {
		return m.Production
	}
	return ""
}

func (m *ProposalProject) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalProject) GetContractor() string {
	if m != nil {
		return m.Contractor
	}
	return ""
}

func (m *ProposalProject) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ProposalProject) GetAmountDetail() string {
	if m != nil {
		return m.AmountDetail
	}
	return ""
}

func (m *ProposalProject) GetToAddr() string {
	if m != nil {
		return m.ToAddr
	}
	return ""
}

func (m *ProposalProject) GetStartBlockHeight() int64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetEndBlockHeight() int64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetRealEndBlockHeight() int64 {
	if m != nil {
		return m.RealEndBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetProjectNeedBlockNum() int32 {
	if m != nil {
		return m.ProjectNeedBlockNum
	}
	return 0
}

type RevokeProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *RevokeProposalProject) Reset()                    { *m = RevokeProposalProject{} }
func (m *RevokeProposalProject) String() string            { return proto.CompactTextString(m) }
func (*RevokeProposalProject) ProtoMessage()               {}
func (*RevokeProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RevokeProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type VoteProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
	Approve    bool   `protobuf:"varint,2,opt,name=approve" json:"approve,omitempty"`
}

func (m *VoteProposalProject) Reset()                    { *m = VoteProposalProject{} }
func (m *VoteProposalProject) String() string            { return proto.CompactTextString(m) }
func (*VoteProposalProject) ProtoMessage()               {}
func (*VoteProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *VoteProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *VoteProposalProject) GetApprove() bool {
	if m != nil {
		return m.Approve
	}
	return false
}

type PubVoteProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
	Oppose     bool   `protobuf:"varint,2,opt,name=oppose" json:"oppose,omitempty"`
}

func (m *PubVoteProposalProject) Reset()                    { *m = PubVoteProposalProject{} }
func (m *PubVoteProposalProject) String() string            { return proto.CompactTextString(m) }
func (*PubVoteProposalProject) ProtoMessage()               {}
func (*PubVoteProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *PubVoteProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

func (m *PubVoteProposalProject) GetOppose() bool {
	if m != nil {
		return m.Oppose
	}
	return false
}

type TerminateProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *TerminateProposalProject) Reset()                    { *m = TerminateProposalProject{} }
func (m *TerminateProposalProject) String() string            { return proto.CompactTextString(m) }
func (*TerminateProposalProject) ProtoMessage()               {}
func (*TerminateProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *TerminateProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

// receipt
type ReceiptProposalProject struct {
	Prev    *AutonomyProposalProject `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *AutonomyProposalProject `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptProposalProject) Reset()                    { *m = ReceiptProposalProject{} }
func (m *ReceiptProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReceiptProposalProject) ProtoMessage()               {}
func (*ReceiptProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ReceiptProposalProject) GetPrev() *AutonomyProposalProject {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptProposalProject) GetCurrent() *AutonomyProposalProject {
	if m != nil {
		return m.Current
	}
	return nil
}

type LocalProposalProject struct {
	PropPrj  *AutonomyProposalProject `protobuf:"bytes,1,opt,name=propPrj" json:"propPrj,omitempty"`
	Comments []string                 `protobuf:"bytes,2,rep,name=comments" json:"comments,omitempty"`
}

func (m *LocalProposalProject) Reset()                    { *m = LocalProposalProject{} }
func (m *LocalProposalProject) String() string            { return proto.CompactTextString(m) }
func (*LocalProposalProject) ProtoMessage()               {}
func (*LocalProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *LocalProposalProject) GetPropPrj() *AutonomyProposalProject {
	if m != nil {
		return m.PropPrj
	}
	return nil
}

func (m *LocalProposalProject) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

// query
type ReplyQueryProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *ReplyQueryProposalProject) Reset()                    { *m = ReplyQueryProposalProject{} }
func (m *ReplyQueryProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReplyQueryProposalProject) ProtoMessage()               {}
func (*ReplyQueryProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ReplyQueryProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type ReplyProposalProject struct {
	PropProjects []*LocalProposalProject `protobuf:"bytes,1,rep,name=propProjects" json:"propProjects,omitempty"`
}

func (m *ReplyProposalProject) Reset()                    { *m = ReplyProposalProject{} }
func (m *ReplyProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReplyProposalProject) ProtoMessage()               {}
func (*ReplyProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ReplyProposalProject) GetPropProjects() []*LocalProposalProject {
	if m != nil {
		return m.PropProjects
	}
	return nil
}

func init() {
	proto.RegisterType((*AutonomyProposalProject)(nil), "types.AutonomyProposalProject")
	proto.RegisterType((*ProposalProject)(nil), "types.ProposalProject")
	proto.RegisterType((*RevokeProposalProject)(nil), "types.RevokeProposalProject")
	proto.RegisterType((*VoteProposalProject)(nil), "types.VoteProposalProject")
	proto.RegisterType((*PubVoteProposalProject)(nil), "types.PubVoteProposalProject")
	proto.RegisterType((*TerminateProposalProject)(nil), "types.TerminateProposalProject")
	proto.RegisterType((*ReceiptProposalProject)(nil), "types.ReceiptProposalProject")
	proto.RegisterType((*LocalProposalProject)(nil), "types.LocalProposalProject")
	proto.RegisterType((*ReplyQueryProposalProject)(nil), "types.ReplyQueryProposalProject")
	proto.RegisterType((*ReplyProposalProject)(nil), "types.ReplyProposalProject")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xd7, 0xa5, 0x5d, 0xaf, 0xdd, 0x0f, 0xbc, 0x51, 0xcc, 0x40, 0x53, 0x95, 0x07, 0x54,
	0x81, 0x54, 0xa1, 0x21, 0xc4, 0x04, 0x0f, 0x68, 0x68, 0x48, 0x20, 0xa1, 0x51, 0x0c, 0x82, 0x67,
	0x37, 0x39, 0xb6, 0x6c, 0x49, 0x6c, 0x39, 0xce, 0x44, 0xff, 0x01, 0xfe, 0x05, 0xfe, 0x54, 0x5e,
	0x51, 0xce, 0x0e, 0x6b, 0xc3, 0x10, 0xf4, 0xcd, 0xdf, 0x77, 0xdf, 0x77, 0x3a, 0x9f, 0xcf, 0x07,
	0x9b, 0xda, 0xa8, 0x0b, 0x8c, 0xec, 0x44, 0x1b, 0x65, 0x15, 0x0b, 0xec, 0x5c, 0x63, 0xb1, 0xbf,
	0x99, 0x46, 0x2a, 0xcb, 0x54, 0xee, 0xd8, 0xf0, 0xc7, 0x1a, 0xdc, 0x39, 0x2e, 0xad, 0xca, 0x55,
	0x36, 0x9f, 0x1a, 0xa5, 0x55, 0x21, 0xd3, 0xa9, 0xf3, 0xb1, 0x23, 0xe8, 0x6b, 0xa3, 0xb4, 0x87,
	0xbc, 0x35, 0x6a, 0x8d, 0xfb, 0x87, 0xc3, 0x09, 0xe5, 0x99, 0x34, 0xc4, 0x62, 0x51, 0xca, 0x9e,
	0xc2, 0x60, 0xa6, 0xa4, 0x89, 0x3f, 0x2b, 0x8b, 0x02, 0x0b, 0xbe, 0x46, 0xd6, 0x5b, 0xde, 0xea,
	0xd9, 0x32, 0xb5, 0x62, 0x49, 0xc6, 0x1e, 0x41, 0x57, 0x97, 0xb3, 0x0a, 0xf1, 0xf6, 0x92, 0x63,
	0x5a, 0xce, 0xd2, 0x24, 0x22, 0x59, 0xad, 0x60, 0x43, 0xe8, 0x14, 0x56, 0xda, 0xb2, 0xe0, 0xeb,
	0xa3, 0xd6, 0x38, 0x10, 0x1e, 0x31, 0x0e, 0x5d, 0x19, 0xc7, 0x06, 0x8b, 0x82, 0x07, 0xa3, 0xd6,
	0xb8, 0x27, 0x6a, 0x58, 0x39, 0xce, 0x31, 0x39, 0x3b, 0xb7, 0xbc, 0x33, 0x6a, 0x8d, 0xdb, 0xc2,
	0x23, 0xb6, 0x07, 0x41, 0x92, 0xc7, 0xf8, 0x8d, 0x77, 0x29, 0x91, 0x03, 0xe1, 0xcf, 0x36, 0x6c,
	0x37, 0x3b, 0xc2, 0x60, 0x7d, 0x8e, 0xd2, 0x50, 0x2b, 0x02, 0x41, 0xe7, 0xca, 0x9d, 0xa9, 0xdc,
	0x9e, 0xd3, 0x25, 0x03, 0xe1, 0x00, 0xdb, 0x81, 0x76, 0x2c, 0xe7, 0x74, 0x8d, 0x40, 0x54, 0x47,
	0x76, 0x00, 0xf0, 0x35, 0x31, 0x85, 0xfd, 0x68, 0xe5, 0x19, 0x52, 0xcd, 0x3d, 0xb1, 0xc0, 0xb0,
	0xfb, 0xd0, 0x4b, 0x65, 0x1d, 0x76, 0x95, 0x5f, 0x13, 0x95, 0x5b, 0x1b, 0x15, 0x97, 0x91, 0x4d,
	0x54, 0x4e, 0xf5, 0xf7, 0xc4, 0x02, 0xc3, 0x46, 0xd0, 0x8f, 0xb1, 0x88, 0x4c, 0xa2, 0x49, 0xd0,
	0x25, 0xc1, 0x22, 0x55, 0x65, 0x88, 0x54, 0x6e, 0x8d, 0x8c, 0xac, 0x32, 0x7c, 0xc3, 0x65, 0xb8,
	0x66, 0xaa, 0xee, 0xc8, 0x4c, 0x95, 0xb9, 0xe5, 0x3d, 0xd7, 0x1d, 0x87, 0x58, 0x08, 0x03, 0x77,
	0x3a, 0x41, 0x2b, 0x93, 0x94, 0x03, 0x39, 0x97, 0xb8, 0xca, 0x6b, 0xd5, 0x71, 0x1c, 0x1b, 0xde,
	0xa7, 0xa8, 0x47, 0xec, 0x21, 0xec, 0x14, 0x56, 0x1a, 0xfb, 0x2a, 0x55, 0xd1, 0xe5, 0x1b, 0xd7,
	0xfb, 0x01, 0x65, 0xff, 0x83, 0x67, 0x0f, 0x60, 0x0b, 0xf3, 0x78, 0x51, 0xb9, 0x49, 0xca, 0x06,
	0xcb, 0x26, 0xc0, 0x0c, 0xca, 0xf4, 0xf5, 0xb2, 0x76, 0x8b, 0xb4, 0x37, 0x44, 0xd8, 0x63, 0xd8,
	0xf5, 0x1f, 0xe1, 0x14, 0xd1, 0x45, 0x4e, 0xcb, 0x8c, 0x6f, 0xd3, 0xcb, 0xdc, 0x14, 0x0a, 0x9f,
	0xc1, 0x6d, 0x81, 0x57, 0xea, 0x12, 0x9b, 0xcf, 0xef, 0x1e, 0x81, 0xa8, 0xb7, 0x27, 0x34, 0x04,
	0xee, 0x11, 0x3c, 0x13, 0xbe, 0x87, 0xdd, 0x6a, 0x34, 0x57, 0xb4, 0xd1, 0xc4, 0x6a, 0x6d, 0xd4,
	0x15, 0xd2, 0x0c, 0x6d, 0x88, 0x1a, 0x86, 0x53, 0x18, 0x4e, 0xdd, 0xb8, 0xaf, 0x9a, 0x73, 0x08,
	0x1d, 0xa5, 0xb5, 0x2a, 0xea, 0x94, 0x1e, 0x85, 0xcf, 0x81, 0x7f, 0x42, 0x93, 0x25, 0xb9, 0x5c,
	0x39, 0x67, 0xf8, 0xbd, 0x05, 0x43, 0x81, 0x11, 0x26, 0xda, 0x36, 0xad, 0x87, 0xb0, 0xae, 0x0d,
	0x5e, 0xf9, 0x1d, 0x71, 0xe0, 0xbf, 0xed, 0x5f, 0x16, 0x8b, 0x20, 0x2d, 0x3b, 0x82, 0x6e, 0x54,
	0x1a, 0x83, 0xb9, 0xf5, 0xfb, 0xe1, 0x5f, 0xb6, 0x5a, 0x1e, 0xa6, 0xb0, 0xf7, 0x4e, 0x45, 0x14,
	0x68, 0x2c, 0xac, 0xae, 0xdb, 0x42, 0x17, 0xff, 0x59, 0x48, 0x2d, 0x67, 0xfb, 0xb0, 0x51, 0xad,
	0x45, 0xcc, 0x6d, 0xb5, 0xac, 0xda, 0xe3, 0x9e, 0xf8, 0x8d, 0xc3, 0x17, 0x70, 0x57, 0xa0, 0x4e,
	0xe7, 0x1f, 0x4a, 0x34, 0xf3, 0x55, 0x7b, 0xf6, 0x05, 0xf6, 0xc8, 0xdc, 0xf4, 0xbd, 0x84, 0xc1,
	0xc2, 0xc2, 0x2c, 0x78, 0x6b, 0xd4, 0x1e, 0xf7, 0x0f, 0xef, 0xf9, 0x7a, 0x6f, 0xba, 0x9d, 0x58,
	0x32, 0xcc, 0x3a, 0xb4, 0xbf, 0x9f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x4e, 0x93, 0x26,
	0xe6, 0x05, 0x00, 0x00,
}
