// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: ssp.proto

package ssp

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// EnumChoise is possible choise a player can make.
type EnumChoise int32

const (
	// Unknow choise means a play did not make a choise in time.
	EnumChoise_UnknownChoise EnumChoise = 0
	EnumChoise_Stone         EnumChoise = 1
	EnumChoise_Scissors      EnumChoise = 2
	EnumChoise_Paper         EnumChoise = 3
)

// Enum value maps for EnumChoise.
var (
	EnumChoise_name = map[int32]string{
		0: "UnknownChoise",
		1: "Stone",
		2: "Scissors",
		3: "Paper",
	}
	EnumChoise_value = map[string]int32{
		"UnknownChoise": 0,
		"Stone":         1,
		"Scissors":      2,
		"Paper":         3,
	}
)

func (x EnumChoise) Enum() *EnumChoise {
	p := new(EnumChoise)
	*p = x
	return p
}

func (x EnumChoise) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumChoise) Descriptor() protoreflect.EnumDescriptor {
	return file_ssp_proto_enumTypes[0].Descriptor()
}

func (EnumChoise) Type() protoreflect.EnumType {
	return &file_ssp_proto_enumTypes[0]
}

func (x EnumChoise) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumChoise.Descriptor instead.
func (EnumChoise) EnumDescriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{0}
}

// EnumStatus is a player's result of a round or a game.
type EnumStatus int32

const (
	EnumStatus_UnknownStatus EnumStatus = 0
	EnumStatus_Winner        EnumStatus = 1
	EnumStatus_Looser        EnumStatus = 2
	EnumStatus_Draw          EnumStatus = 3
)

// Enum value maps for EnumStatus.
var (
	EnumStatus_name = map[int32]string{
		0: "UnknownStatus",
		1: "Winner",
		2: "Looser",
		3: "Draw",
	}
	EnumStatus_value = map[string]int32{
		"UnknownStatus": 0,
		"Winner":        1,
		"Looser":        2,
		"Draw":          3,
	}
)

func (x EnumStatus) Enum() *EnumStatus {
	p := new(EnumStatus)
	*p = x
	return p
}

func (x EnumStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ssp_proto_enumTypes[1].Descriptor()
}

func (EnumStatus) Type() protoreflect.EnumType {
	return &file_ssp_proto_enumTypes[1]
}

func (x EnumStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumStatus.Descriptor instead.
func (EnumStatus) EnumDescriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{1}
}

// AuthRequest is a player's authentication requst message.
type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is a player name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// AuthResponse is a player's authentication response message.
type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is a player ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ReadyRequest is a player's ready request.
type ReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PlayerId is an ID of a player.
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *ReadyRequest) Reset() {
	*x = ReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyRequest) ProtoMessage() {}

func (x *ReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyRequest.ProtoReflect.Descriptor instead.
func (*ReadyRequest) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{2}
}

func (x *ReadyRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

// ReadyResponse is a player's ready request response.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ChoiseTimeoutSeconds is a timeout for player's choise in seconds.
	ChoiseTimeoutSeconds int32 `protobuf:"varint,1,opt,name=choise_timeout_seconds,json=choiseTimeoutSeconds,proto3" json:"choise_timeout_seconds,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{3}
}

func (x *ReadyResponse) GetChoiseTimeoutSeconds() int32 {
	if x != nil {
		return x.ChoiseTimeoutSeconds
	}
	return 0
}

// Choise is what a player chose.
type Choise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Choise is a player choise.
	Choise EnumChoise `protobuf:"varint,1,opt,name=choise,proto3,enum=ssp.EnumChoise" json:"choise,omitempty"`
	// Player who made a choise.
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *Choise) Reset() {
	*x = Choise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choise) ProtoMessage() {}

func (x *Choise) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choise.ProtoReflect.Descriptor instead.
func (*Choise) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{4}
}

func (x *Choise) GetChoise() EnumChoise {
	if x != nil {
		return x.Choise
	}
	return EnumChoise_UnknownChoise
}

func (x *Choise) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

// Score reports the latest round results and the current results of the game.
type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoundResults []*RoundResult `protobuf:"bytes,1,rep,name=round_results,json=roundResults,proto3" json:"round_results,omitempty"`
	GameResults  []*GameResult  `protobuf:"bytes,2,rep,name=game_results,json=gameResults,proto3" json:"game_results,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{5}
}

func (x *Score) GetRoundResults() []*RoundResult {
	if x != nil {
		return x.RoundResults
	}
	return nil
}

func (x *Score) GetGameResults() []*GameResult {
	if x != nil {
		return x.GameResults
	}
	return nil
}

// RoundResult is the latest round result of the player.
type RoundResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Player identifies the player who made the choise.
	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	// Choise is the player choise.
	Choise EnumChoise `protobuf:"varint,2,opt,name=choise,proto3,enum=ssp.EnumChoise" json:"choise,omitempty"`
	// Status is the result of player choise.
	Status EnumStatus `protobuf:"varint,3,opt,name=status,proto3,enum=ssp.EnumStatus" json:"status,omitempty"`
}

func (x *RoundResult) Reset() {
	*x = RoundResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundResult) ProtoMessage() {}

func (x *RoundResult) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundResult.ProtoReflect.Descriptor instead.
func (*RoundResult) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{6}
}

func (x *RoundResult) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *RoundResult) GetChoise() EnumChoise {
	if x != nil {
		return x.Choise
	}
	return EnumChoise_UnknownChoise
}

func (x *RoundResult) GetStatus() EnumStatus {
	if x != nil {
		return x.Status
	}
	return EnumStatus_UnknownStatus
}

// Player decribes a player.
type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is player ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is play name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{7}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GameResult is the current game result of the player.
type GameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Player identifies the player for which the result is for.
	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	// Score is the current score of the player.
	Score int32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	// Status is the current game status of the player.
	Status EnumStatus `protobuf:"varint,3,opt,name=status,proto3,enum=ssp.EnumStatus" json:"status,omitempty"`
	// Rounds is the number of completed rounds in the game.
	Rounds int32 `protobuf:"varint,4,opt,name=rounds,proto3" json:"rounds,omitempty"`
}

func (x *GameResult) Reset() {
	*x = GameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResult) ProtoMessage() {}

func (x *GameResult) ProtoReflect() protoreflect.Message {
	mi := &file_ssp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResult.ProtoReflect.Descriptor instead.
func (*GameResult) Descriptor() ([]byte, []int) {
	return file_ssp_proto_rawDescGZIP(), []int{8}
}

func (x *GameResult) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *GameResult) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GameResult) GetStatus() EnumStatus {
	if x != nil {
		return x.Status
	}
	return EnumStatus_UnknownStatus
}

func (x *GameResult) GetRounds() int32 {
	if x != nil {
		return x.Rounds
	}
	return 0
}

var File_ssp_proto protoreflect.FileDescriptor

var file_ssp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x73, 0x70,
	0x22, 0x21, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x45, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x14, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x6f, 0x69,
	0x73, 0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x35, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x73, 0x70, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b,
	0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0b,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x73,
	0x70, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x73,
	0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x73, 0x70, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x88, 0x01, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x23, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x73, 0x70,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x2a, 0x43, 0x0a, 0x0a, 0x45,
	0x6e, 0x75, 0x6d, 0x43, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x74, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x63, 0x69, 0x73, 0x73,
	0x6f, 0x72, 0x73, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x61, 0x70, 0x65, 0x72, 0x10, 0x03,
	0x2a, 0x41, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4c, 0x6f, 0x6f, 0x73, 0x65, 0x72, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x72, 0x61,
	0x77, 0x10, 0x03, 0x32, 0x8f, 0x01, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x2d, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x05,
	0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x11, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x25,
	0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x0b, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x43, 0x68, 0x6f,
	0x69, 0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x73, 0x73, 0x70, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x76, 0x61, 0x75,
	0x61, 0x2f, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2d, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73,
	0x2d, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x73, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ssp_proto_rawDescOnce sync.Once
	file_ssp_proto_rawDescData = file_ssp_proto_rawDesc
)

func file_ssp_proto_rawDescGZIP() []byte {
	file_ssp_proto_rawDescOnce.Do(func() {
		file_ssp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssp_proto_rawDescData)
	})
	return file_ssp_proto_rawDescData
}

var file_ssp_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ssp_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ssp_proto_goTypes = []interface{}{
	(EnumChoise)(0),       // 0: ssp.EnumChoise
	(EnumStatus)(0),       // 1: ssp.EnumStatus
	(*AuthRequest)(nil),   // 2: ssp.AuthRequest
	(*AuthResponse)(nil),  // 3: ssp.AuthResponse
	(*ReadyRequest)(nil),  // 4: ssp.ReadyRequest
	(*ReadyResponse)(nil), // 5: ssp.ReadyResponse
	(*Choise)(nil),        // 6: ssp.Choise
	(*Score)(nil),         // 7: ssp.Score
	(*RoundResult)(nil),   // 8: ssp.RoundResult
	(*Player)(nil),        // 9: ssp.Player
	(*GameResult)(nil),    // 10: ssp.GameResult
}
var file_ssp_proto_depIdxs = []int32{
	0,  // 0: ssp.Choise.choise:type_name -> ssp.EnumChoise
	8,  // 1: ssp.Score.round_results:type_name -> ssp.RoundResult
	10, // 2: ssp.Score.game_results:type_name -> ssp.GameResult
	9,  // 3: ssp.RoundResult.player:type_name -> ssp.Player
	0,  // 4: ssp.RoundResult.choise:type_name -> ssp.EnumChoise
	1,  // 5: ssp.RoundResult.status:type_name -> ssp.EnumStatus
	9,  // 6: ssp.GameResult.player:type_name -> ssp.Player
	1,  // 7: ssp.GameResult.status:type_name -> ssp.EnumStatus
	2,  // 8: ssp.Gamer.Auth:input_type -> ssp.AuthRequest
	4,  // 9: ssp.Gamer.Ready:input_type -> ssp.ReadyRequest
	6,  // 10: ssp.Gamer.Play:input_type -> ssp.Choise
	3,  // 11: ssp.Gamer.Auth:output_type -> ssp.AuthResponse
	5,  // 12: ssp.Gamer.Ready:output_type -> ssp.ReadyResponse
	7,  // 13: ssp.Gamer.Play:output_type -> ssp.Score
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_ssp_proto_init() }
func file_ssp_proto_init() {
	if File_ssp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_ssp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_ssp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyRequest); i {
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
		file_ssp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
		file_ssp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choise); i {
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
		file_ssp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_ssp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundResult); i {
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
		file_ssp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_ssp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssp_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ssp_proto_goTypes,
		DependencyIndexes: file_ssp_proto_depIdxs,
		EnumInfos:         file_ssp_proto_enumTypes,
		MessageInfos:      file_ssp_proto_msgTypes,
	}.Build()
	File_ssp_proto = out.File
	file_ssp_proto_rawDesc = nil
	file_ssp_proto_goTypes = nil
	file_ssp_proto_depIdxs = nil
}