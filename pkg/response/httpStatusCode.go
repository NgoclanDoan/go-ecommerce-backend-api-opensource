package response

const (
	ErrorCodeSuccess = 20001 // Success
	ErrorCodeParamInvalid = 2003 // Email is invalid
)


// message 
var msg = map[int]string{
	ErrorCodeSuccess: "success",
	ErrorCodeParamInvalid: "Email is invalid",
}
