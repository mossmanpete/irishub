// nolint
package types

import (
	sdk "github.com/irisnet/irishub/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace        sdk.CodespaceType = "distr"
	CodeInvalidInput        CodeType          = 103
	CodeNoDistributionInfo  CodeType          = 104
	CodeDeprecatedOperation CodeType          = 105
)

func ErrNilDelegatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "delegator address is nil")
}
func ErrNilWithdrawAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "withdraw address is nil")
}
func ErrNilValidatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "validator address is nil")
}
func ErrNoDelegationDistInfo(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "no delegation distribution info")
}
func ErrNoValidatorDistInfo(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "no validator distribution info")
}
func ErrDeprecatedOperation(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeDeprecatedOperation, "this operation is deprecated")
}
