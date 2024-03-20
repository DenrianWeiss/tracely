package constants

type OpCodeEnum int

const (
	OpEnumCall OpCodeEnum = iota
	OpEnumDelegateCall
	OpEnumStaticCall
	OpEnumReturn
	OpEnumCreate
	OpEnumRevert
)

const (
	OpCodeCall         = "CALL"
	OpCodeDelegateCall = "DELEGATECALL"
	OpCodeStaticCall   = "STATICCALL"
	OpCodeReturn       = "RETURN"
	OpCodeCreate       = "CREATE"
	OpCodeRevert       = "REVERT"
)

var OpCodeFocus = map[string]bool{
	OpCodeCall:         true,
	OpCodeDelegateCall: true,
	OpCodeStaticCall:   true,
	OpCodeReturn:       true,
	OpCodeCreate:       true,
	OpCodeRevert:       true,
}

var OpCodeEnumToString = map[OpCodeEnum]string{
	OpEnumCall:         OpCodeCall,
	OpEnumDelegateCall: OpCodeDelegateCall,
	OpEnumStaticCall:   OpCodeStaticCall,
	OpEnumReturn:       OpCodeReturn,
	OpEnumCreate:       OpCodeCreate,
	OpEnumRevert:       OpCodeRevert,
}

var OpCodeStringToEnum = map[string]OpCodeEnum{
	OpCodeCall:         OpEnumCall,
	OpCodeDelegateCall: OpEnumDelegateCall,
	OpCodeStaticCall:   OpEnumStaticCall,
	OpCodeReturn:       OpEnumReturn,
	OpCodeCreate:       OpEnumCreate,
	OpCodeRevert:       OpEnumRevert,
}
