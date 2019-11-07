// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/campaign_experiment_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Possible statuses of a campaign experiment.
type CampaignExperimentStatusEnum_CampaignExperimentStatus int32

const (
	// The status has not been specified.
	CampaignExperimentStatusEnum_UNSPECIFIED CampaignExperimentStatusEnum_CampaignExperimentStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignExperimentStatusEnum_UNKNOWN CampaignExperimentStatusEnum_CampaignExperimentStatus = 1
	// The experiment campaign is being initialized.
	CampaignExperimentStatusEnum_INITIALIZING CampaignExperimentStatusEnum_CampaignExperimentStatus = 2
	// Initialization of the experiment campaign failed.
	CampaignExperimentStatusEnum_INITIALIZATION_FAILED CampaignExperimentStatusEnum_CampaignExperimentStatus = 8
	// The experiment campaign is fully initialized. The experiment is currently
	// running, scheduled to run in the future or has ended based on its
	// end date. An experiment with the status INITIALIZING will be updated to
	// ENABLED when it is fully created.
	CampaignExperimentStatusEnum_ENABLED CampaignExperimentStatusEnum_CampaignExperimentStatus = 3
	// The experiment campaign was graduated to a stand-alone
	// campaign, existing independently of the experiment.
	CampaignExperimentStatusEnum_GRADUATED CampaignExperimentStatusEnum_CampaignExperimentStatus = 4
	// The experiment is removed.
	CampaignExperimentStatusEnum_REMOVED CampaignExperimentStatusEnum_CampaignExperimentStatus = 5
	// The experiment's changes are being applied to the original campaign.
	// The long running operation returned by the promote method can be polled
	// to see the status of the promotion.
	CampaignExperimentStatusEnum_PROMOTING CampaignExperimentStatusEnum_CampaignExperimentStatus = 6
	// Promote of the experiment campaign failed.
	CampaignExperimentStatusEnum_PROMOTION_FAILED CampaignExperimentStatusEnum_CampaignExperimentStatus = 9
	// The changes of the experiment are promoted to their original campaign.
	CampaignExperimentStatusEnum_PROMOTED CampaignExperimentStatusEnum_CampaignExperimentStatus = 7
	// The experiment was ended manually. It did not end based on its end date.
	CampaignExperimentStatusEnum_ENDED_MANUALLY CampaignExperimentStatusEnum_CampaignExperimentStatus = 10
)

var CampaignExperimentStatusEnum_CampaignExperimentStatus_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INITIALIZING",
	8:  "INITIALIZATION_FAILED",
	3:  "ENABLED",
	4:  "GRADUATED",
	5:  "REMOVED",
	6:  "PROMOTING",
	9:  "PROMOTION_FAILED",
	7:  "PROMOTED",
	10: "ENDED_MANUALLY",
}

var CampaignExperimentStatusEnum_CampaignExperimentStatus_value = map[string]int32{
	"UNSPECIFIED":           0,
	"UNKNOWN":               1,
	"INITIALIZING":          2,
	"INITIALIZATION_FAILED": 8,
	"ENABLED":               3,
	"GRADUATED":             4,
	"REMOVED":               5,
	"PROMOTING":             6,
	"PROMOTION_FAILED":      9,
	"PROMOTED":              7,
	"ENDED_MANUALLY":        10,
}

func (x CampaignExperimentStatusEnum_CampaignExperimentStatus) String() string {
	return proto.EnumName(CampaignExperimentStatusEnum_CampaignExperimentStatus_name, int32(x))
}

func (CampaignExperimentStatusEnum_CampaignExperimentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ba51d2891452581, []int{0, 0}
}

// Container for enum describing possible statuses of a campaign experiment.
type CampaignExperimentStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignExperimentStatusEnum) Reset()         { *m = CampaignExperimentStatusEnum{} }
func (m *CampaignExperimentStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignExperimentStatusEnum) ProtoMessage()    {}
func (*CampaignExperimentStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ba51d2891452581, []int{0}
}

func (m *CampaignExperimentStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExperimentStatusEnum.Unmarshal(m, b)
}
func (m *CampaignExperimentStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExperimentStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignExperimentStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExperimentStatusEnum.Merge(m, src)
}
func (m *CampaignExperimentStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignExperimentStatusEnum.Size(m)
}
func (m *CampaignExperimentStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExperimentStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExperimentStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.CampaignExperimentStatusEnum_CampaignExperimentStatus", CampaignExperimentStatusEnum_CampaignExperimentStatus_name, CampaignExperimentStatusEnum_CampaignExperimentStatus_value)
	proto.RegisterType((*CampaignExperimentStatusEnum)(nil), "google.ads.googleads.v1.enums.CampaignExperimentStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/campaign_experiment_status.proto", fileDescriptor_3ba51d2891452581)
}

var fileDescriptor_3ba51d2891452581 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x25, 0x5b, 0xe8, 0x87, 0x5b, 0xc0, 0xb2, 0x40, 0x02, 0xd4, 0x3d, 0xb4, 0x3f, 0xc0, 0x51,
	0xc4, 0xcd, 0x48, 0x48, 0xde, 0xda, 0x5d, 0x59, 0xec, 0x3a, 0xab, 0x6d, 0x12, 0x44, 0x15, 0x29,
	0x32, 0x4d, 0x14, 0x45, 0x6a, 0xec, 0x68, 0x9d, 0xad, 0xf8, 0x3d, 0x1c, 0xf9, 0x29, 0xdc, 0xf9,
	0x13, 0x1c, 0x39, 0x70, 0x46, 0x4e, 0xb2, 0xe1, 0xb4, 0xbd, 0x44, 0x6f, 0xe6, 0xcd, 0x7b, 0x13,
	0xcf, 0x03, 0x1f, 0x4b, 0x63, 0xca, 0xfb, 0xc2, 0x57, 0xb9, 0xf5, 0x7b, 0xe8, 0xd0, 0x43, 0xe0,
	0x17, 0x7a, 0x5b, 0x5b, 0xff, 0x4e, 0xd5, 0x8d, 0xaa, 0x4a, 0x9d, 0x15, 0xdf, 0x9a, 0x62, 0x53,
	0xd5, 0x85, 0x6e, 0x33, 0xdb, 0xaa, 0x76, 0x6b, 0x71, 0xb3, 0x31, 0xad, 0x41, 0xd3, 0x5e, 0x84,
	0x55, 0x6e, 0xf1, 0xa8, 0xc7, 0x0f, 0x01, 0xee, 0xf4, 0xef, 0xce, 0x77, 0xf6, 0x4d, 0xe5, 0x2b,
	0xad, 0x4d, 0xab, 0xda, 0xca, 0xe8, 0x41, 0x7c, 0xf9, 0xd7, 0x03, 0xe7, 0x57, 0xc3, 0x06, 0x3e,
	0x2e, 0xb8, 0xe9, 0xfc, 0xb9, 0xde, 0xd6, 0x97, 0xbf, 0x3c, 0xf0, 0x66, 0xdf, 0x00, 0x7a, 0x09,
	0x4e, 0x63, 0x79, 0xb3, 0xe2, 0x57, 0xe2, 0x5a, 0x70, 0x06, 0x9f, 0xa0, 0x53, 0x70, 0x14, 0xcb,
	0x4f, 0x32, 0xfc, 0x2c, 0xa1, 0x87, 0x20, 0x38, 0x13, 0x52, 0x44, 0x82, 0x2e, 0xc4, 0xad, 0x90,
	0x73, 0x38, 0x41, 0x6f, 0xc1, 0xeb, 0xb1, 0x43, 0x23, 0x11, 0xca, 0xec, 0x9a, 0x8a, 0x05, 0x67,
	0xf0, 0xd8, 0x29, 0xb9, 0xa4, 0x33, 0x57, 0x1c, 0xa0, 0xe7, 0xe0, 0x64, 0xbe, 0xa6, 0x2c, 0xa6,
	0x11, 0x67, 0xf0, 0xa9, 0xe3, 0xd6, 0x7c, 0x19, 0x26, 0x9c, 0xc1, 0x67, 0x8e, 0x5b, 0xad, 0xc3,
	0x65, 0x18, 0x39, 0xcb, 0x43, 0xf4, 0x0a, 0xc0, 0xa1, 0xfc, 0xef, 0x76, 0x82, 0xce, 0xc0, 0x71,
	0xdf, 0xe5, 0x0c, 0x1e, 0x21, 0x04, 0x5e, 0x70, 0xc9, 0x38, 0xcb, 0x96, 0x54, 0xc6, 0x74, 0xb1,
	0xf8, 0x02, 0xc1, 0xec, 0x8f, 0x07, 0x2e, 0xee, 0x4c, 0x8d, 0x1f, 0x3d, 0xde, 0x6c, 0xba, 0xef,
	0xe9, 0x2b, 0x77, 0xbd, 0x95, 0x77, 0x3b, 0x1b, 0xf4, 0xa5, 0xb9, 0x57, 0xba, 0xc4, 0x66, 0x53,
	0xfa, 0x65, 0xa1, 0xbb, 0xdb, 0xee, 0xc2, 0x6c, 0x2a, 0xbb, 0x27, 0xdb, 0x0f, 0xdd, 0xf7, 0xfb,
	0xe4, 0x60, 0x4e, 0xe9, 0x8f, 0xc9, 0x74, 0xde, 0x5b, 0xd1, 0xdc, 0xe2, 0x1e, 0x3a, 0x94, 0x04,
	0xd8, 0xe5, 0x60, 0x7f, 0xee, 0xf8, 0x94, 0xe6, 0x36, 0x1d, 0xf9, 0x34, 0x09, 0xd2, 0x8e, 0xff,
	0x3d, 0xb9, 0xe8, 0x9b, 0x84, 0xd0, 0xdc, 0x12, 0x32, 0x4e, 0x10, 0x92, 0x04, 0x84, 0x74, 0x33,
	0x5f, 0x0f, 0xbb, 0x1f, 0x7b, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x45, 0xad, 0x95, 0x03, 0x73,
	0x02, 0x00, 0x00,
}
