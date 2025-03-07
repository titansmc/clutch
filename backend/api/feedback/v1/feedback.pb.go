// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: feedback/v1/feedback.proto

package feedbackv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Origin int32

const (
	// currently these are the supported placements for feedback
	Origin_ORIGIN_UNSPECIFIED Origin = 0
	Origin_HEADER             Origin = 1
	Origin_WIZARD             Origin = 2
)

// Enum value maps for Origin.
var (
	Origin_name = map[int32]string{
		0: "ORIGIN_UNSPECIFIED",
		1: "HEADER",
		2: "WIZARD",
	}
	Origin_value = map[string]int32{
		"ORIGIN_UNSPECIFIED": 0,
		"HEADER":             1,
		"WIZARD":             2,
	}
)

func (x Origin) Enum() *Origin {
	p := new(Origin)
	*p = x
	return p
}

func (x Origin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Origin) Descriptor() protoreflect.EnumDescriptor {
	return file_feedback_v1_feedback_proto_enumTypes[0].Descriptor()
}

func (Origin) Type() protoreflect.EnumType {
	return &file_feedback_v1_feedback_proto_enumTypes[0]
}

func (x Origin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Origin.Descriptor instead.
func (Origin) EnumDescriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{0}
}

type EmojiRating int32

const (
	// these are used to compute a feedback score out of 100
	EmojiRating_EMOJI_UNSPECIFIED EmojiRating = 0
	EmojiRating_SAD               EmojiRating = 1
	EmojiRating_NEUTRAL           EmojiRating = 2
	EmojiRating_HAPPY             EmojiRating = 3
)

// Enum value maps for EmojiRating.
var (
	EmojiRating_name = map[int32]string{
		0: "EMOJI_UNSPECIFIED",
		1: "SAD",
		2: "NEUTRAL",
		3: "HAPPY",
	}
	EmojiRating_value = map[string]int32{
		"EMOJI_UNSPECIFIED": 0,
		"SAD":               1,
		"NEUTRAL":           2,
		"HAPPY":             3,
	}
)

func (x EmojiRating) Enum() *EmojiRating {
	p := new(EmojiRating)
	*p = x
	return p
}

func (x EmojiRating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmojiRating) Descriptor() protoreflect.EnumDescriptor {
	return file_feedback_v1_feedback_proto_enumTypes[1].Descriptor()
}

func (EmojiRating) Type() protoreflect.EnumType {
	return &file_feedback_v1_feedback_proto_enumTypes[1]
}

func (x EmojiRating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmojiRating.Descriptor instead.
func (EmojiRating) EnumDescriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{1}
}

type RatingLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*RatingLabel_Emoji
	Type isRatingLabel_Type `protobuf_oneof:"type"`
	// the corresponding text to show to the user (i.e bad/ok/great)
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RatingLabel) Reset() {
	*x = RatingLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingLabel) ProtoMessage() {}

func (x *RatingLabel) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingLabel.ProtoReflect.Descriptor instead.
func (*RatingLabel) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{0}
}

func (m *RatingLabel) GetType() isRatingLabel_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *RatingLabel) GetEmoji() EmojiRating {
	if x, ok := x.GetType().(*RatingLabel_Emoji); ok {
		return x.Emoji
	}
	return EmojiRating_EMOJI_UNSPECIFIED
}

func (x *RatingLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type isRatingLabel_Type interface {
	isRatingLabel_Type()
}

type RatingLabel_Emoji struct {
	// the emoji type from the defined enums (i.e sad/neutral/happy)
	Emoji EmojiRating `protobuf:"varint,1,opt,name=emoji,proto3,enum=clutch.feedback.v1.EmojiRating,oneof"`
}

func (*RatingLabel_Emoji) isRatingLabel_Type() {}

type RatingScale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*RatingScale_Emoji
	Type isRatingScale_Type `protobuf_oneof:"type"`
}

func (x *RatingScale) Reset() {
	*x = RatingScale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingScale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingScale) ProtoMessage() {}

func (x *RatingScale) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingScale.ProtoReflect.Descriptor instead.
func (*RatingScale) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{1}
}

func (m *RatingScale) GetType() isRatingScale_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *RatingScale) GetEmoji() EmojiRating {
	if x, ok := x.GetType().(*RatingScale_Emoji); ok {
		return x.Emoji
	}
	return EmojiRating_EMOJI_UNSPECIFIED
}

type isRatingScale_Type interface {
	isRatingScale_Type()
}

type RatingScale_Emoji struct {
	Emoji EmojiRating `protobuf:"varint,1,opt,name=emoji,proto3,enum=clutch.feedback.v1.EmojiRating,oneof"`
}

func (*RatingScale_Emoji) isRatingScale_Type() {}

type GetSurveysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the origin of the feedback entry. multiple origins can be passed in the request to return their specific survey
	Origins []Origin `protobuf:"varint,1,rep,packed,name=origins,proto3,enum=clutch.feedback.v1.Origin" json:"origins,omitempty"` // future: add a user field if rules are implemented for whether a user should see the feedback survey
}

func (x *GetSurveysRequest) Reset() {
	*x = GetSurveysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSurveysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveysRequest) ProtoMessage() {}

func (x *GetSurveysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveysRequest.ProtoReflect.Descriptor instead.
func (*GetSurveysRequest) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{2}
}

func (x *GetSurveysRequest) GetOrigins() []Origin {
	if x != nil {
		return x.Origins
	}
	return nil
}

type Survey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the prompt for the rating options
	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// the prompt for the freeform feedback
	FreeformPrompt string `protobuf:"bytes,2,opt,name=freeform_prompt,json=freeformPrompt,proto3" json:"freeform_prompt,omitempty"`
	// the options to show to the user (i.e. bad/ok/great) for the corresponding rating scale
	RatingLabels []*RatingLabel `protobuf:"bytes,3,rep,name=rating_labels,json=ratingLabels,proto3" json:"rating_labels,omitempty"`
}

func (x *Survey) Reset() {
	*x = Survey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Survey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Survey) ProtoMessage() {}

func (x *Survey) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Survey.ProtoReflect.Descriptor instead.
func (*Survey) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{3}
}

func (x *Survey) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *Survey) GetFreeformPrompt() string {
	if x != nil {
		return x.FreeformPrompt
	}
	return ""
}

func (x *Survey) GetRatingLabels() []*RatingLabel {
	if x != nil {
		return x.RatingLabels
	}
	return nil
}

type GetSurveysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the key will be the feedback origin name
	OriginSurvey map[string]*Survey `protobuf:"bytes,1,rep,name=origin_survey,json=originSurvey,proto3" json:"origin_survey,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // future: add a field to indicate if a user should see the feedback survey if rules are implemented
}

func (x *GetSurveysResponse) Reset() {
	*x = GetSurveysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSurveysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSurveysResponse) ProtoMessage() {}

func (x *GetSurveysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSurveysResponse.ProtoReflect.Descriptor instead.
func (*GetSurveysResponse) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{4}
}

func (x *GetSurveysResponse) GetOriginSurvey() map[string]*Survey {
	if x != nil {
		return x.OriginSurvey
	}
	return nil
}

type FeedbackMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// extra info on the feedback (i.e. the survey question, the rating options, the feedback origin, etc.)
	Origin          Origin  `protobuf:"varint,1,opt,name=origin,proto3,enum=clutch.feedback.v1.Origin" json:"origin,omitempty"`
	Survey          *Survey `protobuf:"bytes,2,opt,name=survey,proto3" json:"survey,omitempty"`
	UserSubmitted   bool    `protobuf:"varint,3,opt,name=user_submitted,json=userSubmitted,proto3" json:"user_submitted,omitempty"`
	UrlSearchParams string  `protobuf:"bytes,4,opt,name=url_search_params,json=urlSearchParams,proto3" json:"url_search_params,omitempty"`
}

func (x *FeedbackMetadata) Reset() {
	*x = FeedbackMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedbackMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedbackMetadata) ProtoMessage() {}

func (x *FeedbackMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedbackMetadata.ProtoReflect.Descriptor instead.
func (*FeedbackMetadata) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{5}
}

func (x *FeedbackMetadata) GetOrigin() Origin {
	if x != nil {
		return x.Origin
	}
	return Origin_ORIGIN_UNSPECIFIED
}

func (x *FeedbackMetadata) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

func (x *FeedbackMetadata) GetUserSubmitted() bool {
	if x != nil {
		return x.UserSubmitted
	}
	return false
}

func (x *FeedbackMetadata) GetUrlSearchParams() string {
	if x != nil {
		return x.UrlSearchParams
	}
	return ""
}

type Feedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// workflow url path, workflow name or area of feedback that the submission is for
	FeedbackType string `protobuf:"bytes,1,opt,name=feedback_type,json=feedbackType,proto3" json:"feedback_type,omitempty"`
	// the text option the user selected (i.e. bad/ok/great)
	RatingLabel string `protobuf:"bytes,2,opt,name=rating_label,json=ratingLabel,proto3" json:"rating_label,omitempty"`
	// the corresponding rating scale selection
	RatingScale *RatingScale `protobuf:"bytes,3,opt,name=rating_scale,json=ratingScale,proto3" json:"rating_scale,omitempty"`
	// (optional) freeform input
	FreeformResponse string `protobuf:"bytes,4,opt,name=freeform_response,json=freeformResponse,proto3" json:"freeform_response,omitempty"`
}

func (x *Feedback) Reset() {
	*x = Feedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feedback) ProtoMessage() {}

func (x *Feedback) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feedback.ProtoReflect.Descriptor instead.
func (*Feedback) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{6}
}

func (x *Feedback) GetFeedbackType() string {
	if x != nil {
		return x.FeedbackType
	}
	return ""
}

func (x *Feedback) GetRatingLabel() string {
	if x != nil {
		return x.RatingLabel
	}
	return ""
}

func (x *Feedback) GetRatingScale() *RatingScale {
	if x != nil {
		return x.RatingScale
	}
	return nil
}

func (x *Feedback) GetFreeformResponse() string {
	if x != nil {
		return x.FreeformResponse
	}
	return ""
}

type SubmitFeedbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// client-genereated unique feedback id, which we will also use to update the feedback (essentially replace with the
	// latest)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user's email
	UserId   string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Feedback *Feedback         `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	Metadata *FeedbackMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SubmitFeedbackRequest) Reset() {
	*x = SubmitFeedbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitFeedbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitFeedbackRequest) ProtoMessage() {}

func (x *SubmitFeedbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitFeedbackRequest.ProtoReflect.Descriptor instead.
func (*SubmitFeedbackRequest) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitFeedbackRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubmitFeedbackRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubmitFeedbackRequest) GetFeedback() *Feedback {
	if x != nil {
		return x.Feedback
	}
	return nil
}

func (x *SubmitFeedbackRequest) GetMetadata() *FeedbackMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SubmitFeedbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitFeedbackResponse) Reset() {
	*x = SubmitFeedbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_v1_feedback_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitFeedbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitFeedbackResponse) ProtoMessage() {}

func (x *SubmitFeedbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_v1_feedback_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitFeedbackResponse.ProtoReflect.Descriptor instead.
func (*SubmitFeedbackResponse) Descriptor() ([]byte, []int) {
	return file_feedback_v1_feedback_proto_rawDescGZIP(), []int{8}
}

var File_feedback_v1_feedback_proto protoreflect.FileDescriptor

var file_feedback_v1_feedback_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7e, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x43, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x22, 0x5f, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x12, 0x43, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x92, 0x01, 0x0b, 0x08, 0x01, 0x22, 0x07,
	0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73,
	0x22, 0x8f, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72, 0x65, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x72,
	0x65, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x44, 0x0a, 0x0d,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x1a, 0x5b, 0x0a, 0x11, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01,
	0x20, 0x00, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x72, 0x6c, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x08,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x0d, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x4c, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x66, 0x72, 0x65, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x72, 0x65,
	0x65, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe3, 0x01,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x24, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x38, 0x0a,
	0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x49, 0x47, 0x49,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57,
	0x49, 0x5a, 0x41, 0x52, 0x44, 0x10, 0x02, 0x2a, 0x45, 0x0a, 0x0b, 0x45, 0x6d, 0x6f, 0x6a, 0x69,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x4f, 0x4a, 0x49, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x55, 0x54, 0x52, 0x41,
	0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x41, 0x50, 0x50, 0x59, 0x10, 0x03, 0x32, 0xad,
	0x02, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x50, 0x49, 0x12, 0x85,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x12, 0x25, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2f, 0x67, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x3a, 0x01, 0x2a,
	0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x12, 0x95, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x29, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x65, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x01, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66,
	0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x76, 0x31,
	0x3b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_feedback_v1_feedback_proto_rawDescOnce sync.Once
	file_feedback_v1_feedback_proto_rawDescData = file_feedback_v1_feedback_proto_rawDesc
)

func file_feedback_v1_feedback_proto_rawDescGZIP() []byte {
	file_feedback_v1_feedback_proto_rawDescOnce.Do(func() {
		file_feedback_v1_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_feedback_v1_feedback_proto_rawDescData)
	})
	return file_feedback_v1_feedback_proto_rawDescData
}

var file_feedback_v1_feedback_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_feedback_v1_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_feedback_v1_feedback_proto_goTypes = []interface{}{
	(Origin)(0),                    // 0: clutch.feedback.v1.Origin
	(EmojiRating)(0),               // 1: clutch.feedback.v1.EmojiRating
	(*RatingLabel)(nil),            // 2: clutch.feedback.v1.RatingLabel
	(*RatingScale)(nil),            // 3: clutch.feedback.v1.RatingScale
	(*GetSurveysRequest)(nil),      // 4: clutch.feedback.v1.GetSurveysRequest
	(*Survey)(nil),                 // 5: clutch.feedback.v1.Survey
	(*GetSurveysResponse)(nil),     // 6: clutch.feedback.v1.GetSurveysResponse
	(*FeedbackMetadata)(nil),       // 7: clutch.feedback.v1.FeedbackMetadata
	(*Feedback)(nil),               // 8: clutch.feedback.v1.Feedback
	(*SubmitFeedbackRequest)(nil),  // 9: clutch.feedback.v1.SubmitFeedbackRequest
	(*SubmitFeedbackResponse)(nil), // 10: clutch.feedback.v1.SubmitFeedbackResponse
	nil,                            // 11: clutch.feedback.v1.GetSurveysResponse.OriginSurveyEntry
}
var file_feedback_v1_feedback_proto_depIdxs = []int32{
	1,  // 0: clutch.feedback.v1.RatingLabel.emoji:type_name -> clutch.feedback.v1.EmojiRating
	1,  // 1: clutch.feedback.v1.RatingScale.emoji:type_name -> clutch.feedback.v1.EmojiRating
	0,  // 2: clutch.feedback.v1.GetSurveysRequest.origins:type_name -> clutch.feedback.v1.Origin
	2,  // 3: clutch.feedback.v1.Survey.rating_labels:type_name -> clutch.feedback.v1.RatingLabel
	11, // 4: clutch.feedback.v1.GetSurveysResponse.origin_survey:type_name -> clutch.feedback.v1.GetSurveysResponse.OriginSurveyEntry
	0,  // 5: clutch.feedback.v1.FeedbackMetadata.origin:type_name -> clutch.feedback.v1.Origin
	5,  // 6: clutch.feedback.v1.FeedbackMetadata.survey:type_name -> clutch.feedback.v1.Survey
	3,  // 7: clutch.feedback.v1.Feedback.rating_scale:type_name -> clutch.feedback.v1.RatingScale
	8,  // 8: clutch.feedback.v1.SubmitFeedbackRequest.feedback:type_name -> clutch.feedback.v1.Feedback
	7,  // 9: clutch.feedback.v1.SubmitFeedbackRequest.metadata:type_name -> clutch.feedback.v1.FeedbackMetadata
	5,  // 10: clutch.feedback.v1.GetSurveysResponse.OriginSurveyEntry.value:type_name -> clutch.feedback.v1.Survey
	4,  // 11: clutch.feedback.v1.FeedbackAPI.GetSurveys:input_type -> clutch.feedback.v1.GetSurveysRequest
	9,  // 12: clutch.feedback.v1.FeedbackAPI.SubmitFeedback:input_type -> clutch.feedback.v1.SubmitFeedbackRequest
	6,  // 13: clutch.feedback.v1.FeedbackAPI.GetSurveys:output_type -> clutch.feedback.v1.GetSurveysResponse
	10, // 14: clutch.feedback.v1.FeedbackAPI.SubmitFeedback:output_type -> clutch.feedback.v1.SubmitFeedbackResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_feedback_v1_feedback_proto_init() }
func file_feedback_v1_feedback_proto_init() {
	if File_feedback_v1_feedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feedback_v1_feedback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingLabel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingScale); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSurveysRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Survey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSurveysResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedbackMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feedback); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitFeedbackRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feedback_v1_feedback_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitFeedbackResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_feedback_v1_feedback_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RatingLabel_Emoji)(nil),
	}
	file_feedback_v1_feedback_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RatingScale_Emoji)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feedback_v1_feedback_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feedback_v1_feedback_proto_goTypes,
		DependencyIndexes: file_feedback_v1_feedback_proto_depIdxs,
		EnumInfos:         file_feedback_v1_feedback_proto_enumTypes,
		MessageInfos:      file_feedback_v1_feedback_proto_msgTypes,
	}.Build()
	File_feedback_v1_feedback_proto = out.File
	file_feedback_v1_feedback_proto_rawDesc = nil
	file_feedback_v1_feedback_proto_goTypes = nil
	file_feedback_v1_feedback_proto_depIdxs = nil
}
