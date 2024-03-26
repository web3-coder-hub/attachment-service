package commons

import (
	attachmentPb "attachment-service/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	Sucess = 0

	ServerInternal      = 10000
	InvalidParam        = 10001
	GRPCCallFailed      = 10002
	JsonEncodeFailed    = 10003
	JsonDecodeFailed    = 10004
	AsyncTaskSendFailed = 10005
	ReadDBFailed        = 10006
	WriteDBFailed       = 10007
	ReadRedisFailed     = 10008
	WriteRedisFailed    = 10009
	DBRecordExist       = 10010
	DBRecordNoExist     = 10011

	GetLocalcacheFailed = 10012

	CommonFailed           = 20001
	KvConfigKeyNotExist    = 20002
	KvConfigValueWrongType = 20003
)

var (
	SUCCESS = newError(Sucess, "Success", codes.OK)

	ServerInternalError  = newError(ServerInternal, "Server Internal Error", codes.Aborted)
	InvalidParamError    = newError(InvalidParam, "Invalid Param", codes.Aborted)
	GRPCCallFailedError  = newError(GRPCCallFailed, "GRPC Call Failed", codes.Unavailable)
	JsonEncodeError      = newError(JsonEncodeFailed, "Json Encode Failed", codes.Aborted)
	JsonDecodeError      = newError(JsonDecodeFailed, "Json Decode Failed", codes.Aborted)
	AsyncTaskSendError   = newError(AsyncTaskSendFailed, "Async Task Send Failed", codes.Unavailable)
	ReadDBError          = newError(ReadDBFailed, "Read DB Failed", codes.Unavailable)
	WriteDBError         = newError(WriteDBFailed, "Write DB Failed", codes.Unavailable)
	ReadRedisError       = newError(ReadRedisFailed, "Read Redis Failed", codes.Unavailable)
	WriteRedisError      = newError(WriteRedisFailed, "Write Redis Failed", codes.Unavailable)
	DBRecordExistError   = newError(DBRecordExist, "DB Record Exist", codes.Aborted)
	DBRecordNoExistError = newError(DBRecordNoExist, "DB Record Not Exist", codes.Aborted)
	GetLocalcacheError   = newError(GetLocalcacheFailed, "Get Localcache Failed", codes.Aborted)

	CommonFailedError           = newError(CommonFailed, "Call Failed", codes.Aborted)
	KvConfigKeyNotExistError    = newError(KvConfigKeyNotExist, "Config Key Not Exist", codes.Aborted)
	KvConfigValueWrongTypeError = newError(KvConfigValueWrongType, "Config Value Type Is Wrong", codes.Aborted)
)

var errcodes = map[int32]*Error{}

type Error struct {
	Code     int32
	Message  string
	Param    string
	grpcCode codes.Code
}

func newError(code int32, message string, grpcCode codes.Code) *Error {
	if _, ok := errcodes[code]; ok {
		panic(fmt.Sprintf("Error code:%d exist", code))
	}

	e := &Error{
		Code:     code,
		Message:  message,
		grpcCode: grpcCode,
	}
	errcodes[code] = e
	return e
}

func (e *Error) GrpcErr() error {
	pbErr := &attachmentPb.Err{
		Code:    e.Code,
		Message: e.String(),
	}
	s, err := status.New(e.grpcCode, fmt.Sprintf("Code[%d],Msg[%s]", pbErr.Code, pbErr.Message)).WithDetails(pbErr)
	if err != nil {
		panic(err)
	}
	return s.Err()
}

func (e *Error) clone() *Error {
	return &Error{
		Code:     e.Code,
		Message:  e.Message,
		grpcCode: e.grpcCode,
	}
}

func (e *Error) String() string {
	if len(e.Param) == 0 {
		return e.Message
	}
	return fmt.Sprintf("%s:%s", e.Message, e.Param)
}

func (e *Error) With(param string) *Error {
	nE := e.clone()
	nE.Param = param
	return nE
}

func (e *Error) Withf(format string, v ...interface{}) *Error {
	return e.With(fmt.Sprintf(format, v...))
}

func (e *Error) WithErr(err error) *Error {
	return e.With(err.Error())
}
