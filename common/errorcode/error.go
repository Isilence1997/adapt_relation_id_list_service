package errorcode

const (
	// ParamsInvalidError 请求参数非法
	ParamsInvalidError = -8820007
	// EmptyInputIDError 请求入参vuid为空
	EmptyInputIDError = -8820008
	// RPCFuncCallError rpc调用失败
	SubsRelRPCFuncCallError = -8820009
	// ReturnCodeError 返回码不等于0
	SubsRelReturnCodeError = -8820010
	// SubsFansRPCFuncCallError rpc调用失败
	SubsFansRPCFuncCallError = -8820011
	// SubsFansReturnCodeError 返回码不等于0
	SubsFansReturnCodeError = -8820012
	// UnknownParamError 参数未知错误
	UnknownParamError = -8820013
	// ParseVuidError 解析vuid失败
	ParseVuidError = -8820014
	// QueryFollowVppsError 调用QueryFollowVpps失败
	QueryFollowVppsError = -8820015
	// CallQueryFansListIdxCountError 调用QueryFansListIdxCount失败
	CallQueryFansListIdxCountError = -8820016
	// ParseIndexError 解析Index失败
	ParseIndexError = -8820017
	// ParseIdxCntError 解析vuid失败
	ParseIdxCntError = -8820018
	// CallQueryFansListError 调用QueryFansList失败
	CallQueryFansListError = -8820019
)
