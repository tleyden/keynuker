/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package keynuker

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"

	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	iam "github.com/aws/aws-sdk-go/service/iam"
)

// IAMAPIMock mock
type IAMAPIMock struct {
	mock.Mock
}

func NewIAMAPIMock() *IAMAPIMock {
	return &IAMAPIMock{}
}

// AddClientIDToOpenIDConnectProvider mocked method
func (m *IAMAPIMock) AddClientIDToOpenIDConnectProvider(p0 *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AddClientIDToOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddClientIDToOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddClientIDToOpenIDConnectProviderRequest mocked method
func (m *IAMAPIMock) AddClientIDToOpenIDConnectProviderRequest(p0 *iam.AddClientIDToOpenIDConnectProviderInput) (*request.Request, *iam.AddClientIDToOpenIDConnectProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AddClientIDToOpenIDConnectProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AddClientIDToOpenIDConnectProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddClientIDToOpenIDConnectProviderWithContext mocked method
func (m *IAMAPIMock) AddClientIDToOpenIDConnectProviderWithContext(p0 aws.Context, p1 *iam.AddClientIDToOpenIDConnectProviderInput, p2 ...request.Option) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AddClientIDToOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddClientIDToOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddRoleToInstanceProfile mocked method
func (m *IAMAPIMock) AddRoleToInstanceProfile(p0 *iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AddRoleToInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddRoleToInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddRoleToInstanceProfileRequest mocked method
func (m *IAMAPIMock) AddRoleToInstanceProfileRequest(p0 *iam.AddRoleToInstanceProfileInput) (*request.Request, *iam.AddRoleToInstanceProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AddRoleToInstanceProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AddRoleToInstanceProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddRoleToInstanceProfileWithContext mocked method
func (m *IAMAPIMock) AddRoleToInstanceProfileWithContext(p0 aws.Context, p1 *iam.AddRoleToInstanceProfileInput, p2 ...request.Option) (*iam.AddRoleToInstanceProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AddRoleToInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddRoleToInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddUserToGroup mocked method
func (m *IAMAPIMock) AddUserToGroup(p0 *iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AddUserToGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddUserToGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddUserToGroupRequest mocked method
func (m *IAMAPIMock) AddUserToGroupRequest(p0 *iam.AddUserToGroupInput) (*request.Request, *iam.AddUserToGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AddUserToGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AddUserToGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AddUserToGroupWithContext mocked method
func (m *IAMAPIMock) AddUserToGroupWithContext(p0 aws.Context, p1 *iam.AddUserToGroupInput, p2 ...request.Option) (*iam.AddUserToGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AddUserToGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AddUserToGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachGroupPolicy mocked method
func (m *IAMAPIMock) AttachGroupPolicy(p0 *iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AttachGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachGroupPolicyRequest mocked method
func (m *IAMAPIMock) AttachGroupPolicyRequest(p0 *iam.AttachGroupPolicyInput) (*request.Request, *iam.AttachGroupPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AttachGroupPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AttachGroupPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachGroupPolicyWithContext mocked method
func (m *IAMAPIMock) AttachGroupPolicyWithContext(p0 aws.Context, p1 *iam.AttachGroupPolicyInput, p2 ...request.Option) (*iam.AttachGroupPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AttachGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachRolePolicy mocked method
func (m *IAMAPIMock) AttachRolePolicy(p0 *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AttachRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachRolePolicyRequest mocked method
func (m *IAMAPIMock) AttachRolePolicyRequest(p0 *iam.AttachRolePolicyInput) (*request.Request, *iam.AttachRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AttachRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AttachRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachRolePolicyWithContext mocked method
func (m *IAMAPIMock) AttachRolePolicyWithContext(p0 aws.Context, p1 *iam.AttachRolePolicyInput, p2 ...request.Option) (*iam.AttachRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AttachRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachUserPolicy mocked method
func (m *IAMAPIMock) AttachUserPolicy(p0 *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.AttachUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachUserPolicyRequest mocked method
func (m *IAMAPIMock) AttachUserPolicyRequest(p0 *iam.AttachUserPolicyInput) (*request.Request, *iam.AttachUserPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.AttachUserPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.AttachUserPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// AttachUserPolicyWithContext mocked method
func (m *IAMAPIMock) AttachUserPolicyWithContext(p0 aws.Context, p1 *iam.AttachUserPolicyInput, p2 ...request.Option) (*iam.AttachUserPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.AttachUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.AttachUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ChangePassword mocked method
func (m *IAMAPIMock) ChangePassword(p0 *iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ChangePasswordOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ChangePasswordOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ChangePasswordRequest mocked method
func (m *IAMAPIMock) ChangePasswordRequest(p0 *iam.ChangePasswordInput) (*request.Request, *iam.ChangePasswordOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ChangePasswordOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ChangePasswordOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ChangePasswordWithContext mocked method
func (m *IAMAPIMock) ChangePasswordWithContext(p0 aws.Context, p1 *iam.ChangePasswordInput, p2 ...request.Option) (*iam.ChangePasswordOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ChangePasswordOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ChangePasswordOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccessKey mocked method
func (m *IAMAPIMock) CreateAccessKey(p0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccessKeyRequest mocked method
func (m *IAMAPIMock) CreateAccessKeyRequest(p0 *iam.CreateAccessKeyInput) (*request.Request, *iam.CreateAccessKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateAccessKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateAccessKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccessKeyWithContext mocked method
func (m *IAMAPIMock) CreateAccessKeyWithContext(p0 aws.Context, p1 *iam.CreateAccessKeyInput, p2 ...request.Option) (*iam.CreateAccessKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccountAlias mocked method
func (m *IAMAPIMock) CreateAccountAlias(p0 *iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateAccountAliasOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateAccountAliasOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccountAliasRequest mocked method
func (m *IAMAPIMock) CreateAccountAliasRequest(p0 *iam.CreateAccountAliasInput) (*request.Request, *iam.CreateAccountAliasOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateAccountAliasOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateAccountAliasOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateAccountAliasWithContext mocked method
func (m *IAMAPIMock) CreateAccountAliasWithContext(p0 aws.Context, p1 *iam.CreateAccountAliasInput, p2 ...request.Option) (*iam.CreateAccountAliasOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateAccountAliasOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateAccountAliasOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateGroup mocked method
func (m *IAMAPIMock) CreateGroup(p0 *iam.CreateGroupInput) (*iam.CreateGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateGroupRequest mocked method
func (m *IAMAPIMock) CreateGroupRequest(p0 *iam.CreateGroupInput) (*request.Request, *iam.CreateGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateGroupWithContext mocked method
func (m *IAMAPIMock) CreateGroupWithContext(p0 aws.Context, p1 *iam.CreateGroupInput, p2 ...request.Option) (*iam.CreateGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateInstanceProfile mocked method
func (m *IAMAPIMock) CreateInstanceProfile(p0 *iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateInstanceProfileRequest mocked method
func (m *IAMAPIMock) CreateInstanceProfileRequest(p0 *iam.CreateInstanceProfileInput) (*request.Request, *iam.CreateInstanceProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateInstanceProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateInstanceProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateInstanceProfileWithContext mocked method
func (m *IAMAPIMock) CreateInstanceProfileWithContext(p0 aws.Context, p1 *iam.CreateInstanceProfileInput, p2 ...request.Option) (*iam.CreateInstanceProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateLoginProfile mocked method
func (m *IAMAPIMock) CreateLoginProfile(p0 *iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateLoginProfileRequest mocked method
func (m *IAMAPIMock) CreateLoginProfileRequest(p0 *iam.CreateLoginProfileInput) (*request.Request, *iam.CreateLoginProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateLoginProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateLoginProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateLoginProfileWithContext mocked method
func (m *IAMAPIMock) CreateLoginProfileWithContext(p0 aws.Context, p1 *iam.CreateLoginProfileInput, p2 ...request.Option) (*iam.CreateLoginProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateOpenIDConnectProvider mocked method
func (m *IAMAPIMock) CreateOpenIDConnectProvider(p0 *iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateOpenIDConnectProviderRequest mocked method
func (m *IAMAPIMock) CreateOpenIDConnectProviderRequest(p0 *iam.CreateOpenIDConnectProviderInput) (*request.Request, *iam.CreateOpenIDConnectProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateOpenIDConnectProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateOpenIDConnectProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateOpenIDConnectProviderWithContext mocked method
func (m *IAMAPIMock) CreateOpenIDConnectProviderWithContext(p0 aws.Context, p1 *iam.CreateOpenIDConnectProviderInput, p2 ...request.Option) (*iam.CreateOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicy mocked method
func (m *IAMAPIMock) CreatePolicy(p0 *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreatePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreatePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicyRequest mocked method
func (m *IAMAPIMock) CreatePolicyRequest(p0 *iam.CreatePolicyInput) (*request.Request, *iam.CreatePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreatePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreatePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicyVersion mocked method
func (m *IAMAPIMock) CreatePolicyVersion(p0 *iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreatePolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreatePolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicyVersionRequest mocked method
func (m *IAMAPIMock) CreatePolicyVersionRequest(p0 *iam.CreatePolicyVersionInput) (*request.Request, *iam.CreatePolicyVersionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreatePolicyVersionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreatePolicyVersionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicyVersionWithContext mocked method
func (m *IAMAPIMock) CreatePolicyVersionWithContext(p0 aws.Context, p1 *iam.CreatePolicyVersionInput, p2 ...request.Option) (*iam.CreatePolicyVersionOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreatePolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreatePolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreatePolicyWithContext mocked method
func (m *IAMAPIMock) CreatePolicyWithContext(p0 aws.Context, p1 *iam.CreatePolicyInput, p2 ...request.Option) (*iam.CreatePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreatePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreatePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateRole mocked method
func (m *IAMAPIMock) CreateRole(p0 *iam.CreateRoleInput) (*iam.CreateRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateRoleRequest mocked method
func (m *IAMAPIMock) CreateRoleRequest(p0 *iam.CreateRoleInput) (*request.Request, *iam.CreateRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateRoleWithContext mocked method
func (m *IAMAPIMock) CreateRoleWithContext(p0 aws.Context, p1 *iam.CreateRoleInput, p2 ...request.Option) (*iam.CreateRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSAMLProvider mocked method
func (m *IAMAPIMock) CreateSAMLProvider(p0 *iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSAMLProviderRequest mocked method
func (m *IAMAPIMock) CreateSAMLProviderRequest(p0 *iam.CreateSAMLProviderInput) (*request.Request, *iam.CreateSAMLProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateSAMLProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateSAMLProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateSAMLProviderWithContext mocked method
func (m *IAMAPIMock) CreateSAMLProviderWithContext(p0 aws.Context, p1 *iam.CreateSAMLProviderInput, p2 ...request.Option) (*iam.CreateSAMLProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceLinkedRole mocked method
func (m *IAMAPIMock) CreateServiceLinkedRole(p0 *iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateServiceLinkedRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateServiceLinkedRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceLinkedRoleRequest mocked method
func (m *IAMAPIMock) CreateServiceLinkedRoleRequest(p0 *iam.CreateServiceLinkedRoleInput) (*request.Request, *iam.CreateServiceLinkedRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateServiceLinkedRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateServiceLinkedRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceLinkedRoleWithContext mocked method
func (m *IAMAPIMock) CreateServiceLinkedRoleWithContext(p0 aws.Context, p1 *iam.CreateServiceLinkedRoleInput, p2 ...request.Option) (*iam.CreateServiceLinkedRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateServiceLinkedRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateServiceLinkedRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceSpecificCredential mocked method
func (m *IAMAPIMock) CreateServiceSpecificCredential(p0 *iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceSpecificCredentialRequest mocked method
func (m *IAMAPIMock) CreateServiceSpecificCredentialRequest(p0 *iam.CreateServiceSpecificCredentialInput) (*request.Request, *iam.CreateServiceSpecificCredentialOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateServiceSpecificCredentialOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateServiceSpecificCredentialOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateServiceSpecificCredentialWithContext mocked method
func (m *IAMAPIMock) CreateServiceSpecificCredentialWithContext(p0 aws.Context, p1 *iam.CreateServiceSpecificCredentialInput, p2 ...request.Option) (*iam.CreateServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateUser mocked method
func (m *IAMAPIMock) CreateUser(p0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateUserRequest mocked method
func (m *IAMAPIMock) CreateUserRequest(p0 *iam.CreateUserInput) (*request.Request, *iam.CreateUserOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateUserOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateUserOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateUserWithContext mocked method
func (m *IAMAPIMock) CreateUserWithContext(p0 aws.Context, p1 *iam.CreateUserInput, p2 ...request.Option) (*iam.CreateUserOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVirtualMFADevice mocked method
func (m *IAMAPIMock) CreateVirtualMFADevice(p0 *iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.CreateVirtualMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateVirtualMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVirtualMFADeviceRequest mocked method
func (m *IAMAPIMock) CreateVirtualMFADeviceRequest(p0 *iam.CreateVirtualMFADeviceInput) (*request.Request, *iam.CreateVirtualMFADeviceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.CreateVirtualMFADeviceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.CreateVirtualMFADeviceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// CreateVirtualMFADeviceWithContext mocked method
func (m *IAMAPIMock) CreateVirtualMFADeviceWithContext(p0 aws.Context, p1 *iam.CreateVirtualMFADeviceInput, p2 ...request.Option) (*iam.CreateVirtualMFADeviceOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.CreateVirtualMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.CreateVirtualMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeactivateMFADevice mocked method
func (m *IAMAPIMock) DeactivateMFADevice(p0 *iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeactivateMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeactivateMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeactivateMFADeviceRequest mocked method
func (m *IAMAPIMock) DeactivateMFADeviceRequest(p0 *iam.DeactivateMFADeviceInput) (*request.Request, *iam.DeactivateMFADeviceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeactivateMFADeviceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeactivateMFADeviceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeactivateMFADeviceWithContext mocked method
func (m *IAMAPIMock) DeactivateMFADeviceWithContext(p0 aws.Context, p1 *iam.DeactivateMFADeviceInput, p2 ...request.Option) (*iam.DeactivateMFADeviceOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeactivateMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeactivateMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccessKey mocked method
func (m *IAMAPIMock) DeleteAccessKey(p0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccessKeyRequest mocked method
func (m *IAMAPIMock) DeleteAccessKeyRequest(p0 *iam.DeleteAccessKeyInput) (*request.Request, *iam.DeleteAccessKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteAccessKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteAccessKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccessKeyWithContext mocked method
func (m *IAMAPIMock) DeleteAccessKeyWithContext(p0 aws.Context, p1 *iam.DeleteAccessKeyInput, p2 ...request.Option) (*iam.DeleteAccessKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountAlias mocked method
func (m *IAMAPIMock) DeleteAccountAlias(p0 *iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteAccountAliasOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccountAliasOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountAliasRequest mocked method
func (m *IAMAPIMock) DeleteAccountAliasRequest(p0 *iam.DeleteAccountAliasInput) (*request.Request, *iam.DeleteAccountAliasOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteAccountAliasOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteAccountAliasOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountAliasWithContext mocked method
func (m *IAMAPIMock) DeleteAccountAliasWithContext(p0 aws.Context, p1 *iam.DeleteAccountAliasInput, p2 ...request.Option) (*iam.DeleteAccountAliasOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteAccountAliasOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccountAliasOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountPasswordPolicy mocked method
func (m *IAMAPIMock) DeleteAccountPasswordPolicy(p0 *iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountPasswordPolicyRequest mocked method
func (m *IAMAPIMock) DeleteAccountPasswordPolicyRequest(p0 *iam.DeleteAccountPasswordPolicyInput) (*request.Request, *iam.DeleteAccountPasswordPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteAccountPasswordPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteAccountPasswordPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteAccountPasswordPolicyWithContext mocked method
func (m *IAMAPIMock) DeleteAccountPasswordPolicyWithContext(p0 aws.Context, p1 *iam.DeleteAccountPasswordPolicyInput, p2 ...request.Option) (*iam.DeleteAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroup mocked method
func (m *IAMAPIMock) DeleteGroup(p0 *iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroupPolicy mocked method
func (m *IAMAPIMock) DeleteGroupPolicy(p0 *iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroupPolicyRequest mocked method
func (m *IAMAPIMock) DeleteGroupPolicyRequest(p0 *iam.DeleteGroupPolicyInput) (*request.Request, *iam.DeleteGroupPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteGroupPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteGroupPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroupPolicyWithContext mocked method
func (m *IAMAPIMock) DeleteGroupPolicyWithContext(p0 aws.Context, p1 *iam.DeleteGroupPolicyInput, p2 ...request.Option) (*iam.DeleteGroupPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroupRequest mocked method
func (m *IAMAPIMock) DeleteGroupRequest(p0 *iam.DeleteGroupInput) (*request.Request, *iam.DeleteGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteGroupWithContext mocked method
func (m *IAMAPIMock) DeleteGroupWithContext(p0 aws.Context, p1 *iam.DeleteGroupInput, p2 ...request.Option) (*iam.DeleteGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteInstanceProfile mocked method
func (m *IAMAPIMock) DeleteInstanceProfile(p0 *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteInstanceProfileRequest mocked method
func (m *IAMAPIMock) DeleteInstanceProfileRequest(p0 *iam.DeleteInstanceProfileInput) (*request.Request, *iam.DeleteInstanceProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteInstanceProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteInstanceProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteInstanceProfileWithContext mocked method
func (m *IAMAPIMock) DeleteInstanceProfileWithContext(p0 aws.Context, p1 *iam.DeleteInstanceProfileInput, p2 ...request.Option) (*iam.DeleteInstanceProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteLoginProfile mocked method
func (m *IAMAPIMock) DeleteLoginProfile(p0 *iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteLoginProfileRequest mocked method
func (m *IAMAPIMock) DeleteLoginProfileRequest(p0 *iam.DeleteLoginProfileInput) (*request.Request, *iam.DeleteLoginProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteLoginProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteLoginProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteLoginProfileWithContext mocked method
func (m *IAMAPIMock) DeleteLoginProfileWithContext(p0 aws.Context, p1 *iam.DeleteLoginProfileInput, p2 ...request.Option) (*iam.DeleteLoginProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteOpenIDConnectProvider mocked method
func (m *IAMAPIMock) DeleteOpenIDConnectProvider(p0 *iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteOpenIDConnectProviderRequest mocked method
func (m *IAMAPIMock) DeleteOpenIDConnectProviderRequest(p0 *iam.DeleteOpenIDConnectProviderInput) (*request.Request, *iam.DeleteOpenIDConnectProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteOpenIDConnectProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteOpenIDConnectProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteOpenIDConnectProviderWithContext mocked method
func (m *IAMAPIMock) DeleteOpenIDConnectProviderWithContext(p0 aws.Context, p1 *iam.DeleteOpenIDConnectProviderInput, p2 ...request.Option) (*iam.DeleteOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicy mocked method
func (m *IAMAPIMock) DeletePolicy(p0 *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeletePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeletePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicyRequest mocked method
func (m *IAMAPIMock) DeletePolicyRequest(p0 *iam.DeletePolicyInput) (*request.Request, *iam.DeletePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeletePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeletePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicyVersion mocked method
func (m *IAMAPIMock) DeletePolicyVersion(p0 *iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeletePolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeletePolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicyVersionRequest mocked method
func (m *IAMAPIMock) DeletePolicyVersionRequest(p0 *iam.DeletePolicyVersionInput) (*request.Request, *iam.DeletePolicyVersionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeletePolicyVersionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeletePolicyVersionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicyVersionWithContext mocked method
func (m *IAMAPIMock) DeletePolicyVersionWithContext(p0 aws.Context, p1 *iam.DeletePolicyVersionInput, p2 ...request.Option) (*iam.DeletePolicyVersionOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeletePolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeletePolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeletePolicyWithContext mocked method
func (m *IAMAPIMock) DeletePolicyWithContext(p0 aws.Context, p1 *iam.DeletePolicyInput, p2 ...request.Option) (*iam.DeletePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeletePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeletePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRole mocked method
func (m *IAMAPIMock) DeleteRole(p0 *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRolePolicy mocked method
func (m *IAMAPIMock) DeleteRolePolicy(p0 *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRolePolicyRequest mocked method
func (m *IAMAPIMock) DeleteRolePolicyRequest(p0 *iam.DeleteRolePolicyInput) (*request.Request, *iam.DeleteRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRolePolicyWithContext mocked method
func (m *IAMAPIMock) DeleteRolePolicyWithContext(p0 aws.Context, p1 *iam.DeleteRolePolicyInput, p2 ...request.Option) (*iam.DeleteRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRoleRequest mocked method
func (m *IAMAPIMock) DeleteRoleRequest(p0 *iam.DeleteRoleInput) (*request.Request, *iam.DeleteRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteRoleWithContext mocked method
func (m *IAMAPIMock) DeleteRoleWithContext(p0 aws.Context, p1 *iam.DeleteRoleInput, p2 ...request.Option) (*iam.DeleteRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSAMLProvider mocked method
func (m *IAMAPIMock) DeleteSAMLProvider(p0 *iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSAMLProviderRequest mocked method
func (m *IAMAPIMock) DeleteSAMLProviderRequest(p0 *iam.DeleteSAMLProviderInput) (*request.Request, *iam.DeleteSAMLProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteSAMLProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteSAMLProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSAMLProviderWithContext mocked method
func (m *IAMAPIMock) DeleteSAMLProviderWithContext(p0 aws.Context, p1 *iam.DeleteSAMLProviderInput, p2 ...request.Option) (*iam.DeleteSAMLProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSSHPublicKey mocked method
func (m *IAMAPIMock) DeleteSSHPublicKey(p0 *iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSSHPublicKeyRequest mocked method
func (m *IAMAPIMock) DeleteSSHPublicKeyRequest(p0 *iam.DeleteSSHPublicKeyInput) (*request.Request, *iam.DeleteSSHPublicKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteSSHPublicKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteSSHPublicKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSSHPublicKeyWithContext mocked method
func (m *IAMAPIMock) DeleteSSHPublicKeyWithContext(p0 aws.Context, p1 *iam.DeleteSSHPublicKeyInput, p2 ...request.Option) (*iam.DeleteSSHPublicKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServerCertificate mocked method
func (m *IAMAPIMock) DeleteServerCertificate(p0 *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServerCertificateRequest mocked method
func (m *IAMAPIMock) DeleteServerCertificateRequest(p0 *iam.DeleteServerCertificateInput) (*request.Request, *iam.DeleteServerCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteServerCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteServerCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServerCertificateWithContext mocked method
func (m *IAMAPIMock) DeleteServerCertificateWithContext(p0 aws.Context, p1 *iam.DeleteServerCertificateInput, p2 ...request.Option) (*iam.DeleteServerCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceLinkedRole mocked method
func (m *IAMAPIMock) DeleteServiceLinkedRole(p0 *iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteServiceLinkedRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServiceLinkedRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceLinkedRoleRequest mocked method
func (m *IAMAPIMock) DeleteServiceLinkedRoleRequest(p0 *iam.DeleteServiceLinkedRoleInput) (*request.Request, *iam.DeleteServiceLinkedRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteServiceLinkedRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteServiceLinkedRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceLinkedRoleWithContext mocked method
func (m *IAMAPIMock) DeleteServiceLinkedRoleWithContext(p0 aws.Context, p1 *iam.DeleteServiceLinkedRoleInput, p2 ...request.Option) (*iam.DeleteServiceLinkedRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteServiceLinkedRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServiceLinkedRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceSpecificCredential mocked method
func (m *IAMAPIMock) DeleteServiceSpecificCredential(p0 *iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceSpecificCredentialRequest mocked method
func (m *IAMAPIMock) DeleteServiceSpecificCredentialRequest(p0 *iam.DeleteServiceSpecificCredentialInput) (*request.Request, *iam.DeleteServiceSpecificCredentialOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteServiceSpecificCredentialOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteServiceSpecificCredentialOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteServiceSpecificCredentialWithContext mocked method
func (m *IAMAPIMock) DeleteServiceSpecificCredentialWithContext(p0 aws.Context, p1 *iam.DeleteServiceSpecificCredentialInput, p2 ...request.Option) (*iam.DeleteServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSigningCertificate mocked method
func (m *IAMAPIMock) DeleteSigningCertificate(p0 *iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSigningCertificateRequest mocked method
func (m *IAMAPIMock) DeleteSigningCertificateRequest(p0 *iam.DeleteSigningCertificateInput) (*request.Request, *iam.DeleteSigningCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteSigningCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteSigningCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteSigningCertificateWithContext mocked method
func (m *IAMAPIMock) DeleteSigningCertificateWithContext(p0 aws.Context, p1 *iam.DeleteSigningCertificateInput, p2 ...request.Option) (*iam.DeleteSigningCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUser mocked method
func (m *IAMAPIMock) DeleteUser(p0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUserPolicy mocked method
func (m *IAMAPIMock) DeleteUserPolicy(p0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUserPolicyRequest mocked method
func (m *IAMAPIMock) DeleteUserPolicyRequest(p0 *iam.DeleteUserPolicyInput) (*request.Request, *iam.DeleteUserPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteUserPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteUserPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUserPolicyWithContext mocked method
func (m *IAMAPIMock) DeleteUserPolicyWithContext(p0 aws.Context, p1 *iam.DeleteUserPolicyInput, p2 ...request.Option) (*iam.DeleteUserPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUserRequest mocked method
func (m *IAMAPIMock) DeleteUserRequest(p0 *iam.DeleteUserInput) (*request.Request, *iam.DeleteUserOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteUserOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteUserOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteUserWithContext mocked method
func (m *IAMAPIMock) DeleteUserWithContext(p0 aws.Context, p1 *iam.DeleteUserInput, p2 ...request.Option) (*iam.DeleteUserOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVirtualMFADevice mocked method
func (m *IAMAPIMock) DeleteVirtualMFADevice(p0 *iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DeleteVirtualMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteVirtualMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVirtualMFADeviceRequest mocked method
func (m *IAMAPIMock) DeleteVirtualMFADeviceRequest(p0 *iam.DeleteVirtualMFADeviceInput) (*request.Request, *iam.DeleteVirtualMFADeviceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DeleteVirtualMFADeviceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DeleteVirtualMFADeviceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DeleteVirtualMFADeviceWithContext mocked method
func (m *IAMAPIMock) DeleteVirtualMFADeviceWithContext(p0 aws.Context, p1 *iam.DeleteVirtualMFADeviceInput, p2 ...request.Option) (*iam.DeleteVirtualMFADeviceOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DeleteVirtualMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DeleteVirtualMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachGroupPolicy mocked method
func (m *IAMAPIMock) DetachGroupPolicy(p0 *iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DetachGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachGroupPolicyRequest mocked method
func (m *IAMAPIMock) DetachGroupPolicyRequest(p0 *iam.DetachGroupPolicyInput) (*request.Request, *iam.DetachGroupPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DetachGroupPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DetachGroupPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachGroupPolicyWithContext mocked method
func (m *IAMAPIMock) DetachGroupPolicyWithContext(p0 aws.Context, p1 *iam.DetachGroupPolicyInput, p2 ...request.Option) (*iam.DetachGroupPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DetachGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachRolePolicy mocked method
func (m *IAMAPIMock) DetachRolePolicy(p0 *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DetachRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachRolePolicyRequest mocked method
func (m *IAMAPIMock) DetachRolePolicyRequest(p0 *iam.DetachRolePolicyInput) (*request.Request, *iam.DetachRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DetachRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DetachRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachRolePolicyWithContext mocked method
func (m *IAMAPIMock) DetachRolePolicyWithContext(p0 aws.Context, p1 *iam.DetachRolePolicyInput, p2 ...request.Option) (*iam.DetachRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DetachRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachUserPolicy mocked method
func (m *IAMAPIMock) DetachUserPolicy(p0 *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.DetachUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachUserPolicyRequest mocked method
func (m *IAMAPIMock) DetachUserPolicyRequest(p0 *iam.DetachUserPolicyInput) (*request.Request, *iam.DetachUserPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.DetachUserPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.DetachUserPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// DetachUserPolicyWithContext mocked method
func (m *IAMAPIMock) DetachUserPolicyWithContext(p0 aws.Context, p1 *iam.DetachUserPolicyInput, p2 ...request.Option) (*iam.DetachUserPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.DetachUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.DetachUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableMFADevice mocked method
func (m *IAMAPIMock) EnableMFADevice(p0 *iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.EnableMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.EnableMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableMFADeviceRequest mocked method
func (m *IAMAPIMock) EnableMFADeviceRequest(p0 *iam.EnableMFADeviceInput) (*request.Request, *iam.EnableMFADeviceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.EnableMFADeviceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.EnableMFADeviceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// EnableMFADeviceWithContext mocked method
func (m *IAMAPIMock) EnableMFADeviceWithContext(p0 aws.Context, p1 *iam.EnableMFADeviceInput, p2 ...request.Option) (*iam.EnableMFADeviceOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.EnableMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.EnableMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GenerateCredentialReport mocked method
func (m *IAMAPIMock) GenerateCredentialReport(p0 *iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GenerateCredentialReportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GenerateCredentialReportOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GenerateCredentialReportRequest mocked method
func (m *IAMAPIMock) GenerateCredentialReportRequest(p0 *iam.GenerateCredentialReportInput) (*request.Request, *iam.GenerateCredentialReportOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GenerateCredentialReportOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GenerateCredentialReportOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GenerateCredentialReportWithContext mocked method
func (m *IAMAPIMock) GenerateCredentialReportWithContext(p0 aws.Context, p1 *iam.GenerateCredentialReportInput, p2 ...request.Option) (*iam.GenerateCredentialReportOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GenerateCredentialReportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GenerateCredentialReportOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccessKeyLastUsed mocked method
func (m *IAMAPIMock) GetAccessKeyLastUsed(p0 *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetAccessKeyLastUsedOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccessKeyLastUsedOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccessKeyLastUsedRequest mocked method
func (m *IAMAPIMock) GetAccessKeyLastUsedRequest(p0 *iam.GetAccessKeyLastUsedInput) (*request.Request, *iam.GetAccessKeyLastUsedOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetAccessKeyLastUsedOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetAccessKeyLastUsedOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccessKeyLastUsedWithContext mocked method
func (m *IAMAPIMock) GetAccessKeyLastUsedWithContext(p0 aws.Context, p1 *iam.GetAccessKeyLastUsedInput, p2 ...request.Option) (*iam.GetAccessKeyLastUsedOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetAccessKeyLastUsedOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccessKeyLastUsedOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountAuthorizationDetails mocked method
func (m *IAMAPIMock) GetAccountAuthorizationDetails(p0 *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetAccountAuthorizationDetailsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountAuthorizationDetailsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountAuthorizationDetailsPages mocked method
func (m *IAMAPIMock) GetAccountAuthorizationDetailsPages(p0 *iam.GetAccountAuthorizationDetailsInput, p1 func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// GetAccountAuthorizationDetailsPagesWithContext mocked method
func (m *IAMAPIMock) GetAccountAuthorizationDetailsPagesWithContext(p0 aws.Context, p1 *iam.GetAccountAuthorizationDetailsInput, p2 func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// GetAccountAuthorizationDetailsRequest mocked method
func (m *IAMAPIMock) GetAccountAuthorizationDetailsRequest(p0 *iam.GetAccountAuthorizationDetailsInput) (*request.Request, *iam.GetAccountAuthorizationDetailsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetAccountAuthorizationDetailsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetAccountAuthorizationDetailsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountAuthorizationDetailsWithContext mocked method
func (m *IAMAPIMock) GetAccountAuthorizationDetailsWithContext(p0 aws.Context, p1 *iam.GetAccountAuthorizationDetailsInput, p2 ...request.Option) (*iam.GetAccountAuthorizationDetailsOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetAccountAuthorizationDetailsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountAuthorizationDetailsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountPasswordPolicy mocked method
func (m *IAMAPIMock) GetAccountPasswordPolicy(p0 *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountPasswordPolicyRequest mocked method
func (m *IAMAPIMock) GetAccountPasswordPolicyRequest(p0 *iam.GetAccountPasswordPolicyInput) (*request.Request, *iam.GetAccountPasswordPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetAccountPasswordPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetAccountPasswordPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountPasswordPolicyWithContext mocked method
func (m *IAMAPIMock) GetAccountPasswordPolicyWithContext(p0 aws.Context, p1 *iam.GetAccountPasswordPolicyInput, p2 ...request.Option) (*iam.GetAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountSummary mocked method
func (m *IAMAPIMock) GetAccountSummary(p0 *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetAccountSummaryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountSummaryOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountSummaryRequest mocked method
func (m *IAMAPIMock) GetAccountSummaryRequest(p0 *iam.GetAccountSummaryInput) (*request.Request, *iam.GetAccountSummaryOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetAccountSummaryOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetAccountSummaryOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetAccountSummaryWithContext mocked method
func (m *IAMAPIMock) GetAccountSummaryWithContext(p0 aws.Context, p1 *iam.GetAccountSummaryInput, p2 ...request.Option) (*iam.GetAccountSummaryOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetAccountSummaryOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetAccountSummaryOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForCustomPolicy mocked method
func (m *IAMAPIMock) GetContextKeysForCustomPolicy(p0 *iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {

	ret := m.Called(p0)

	var r0 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForCustomPolicyRequest mocked method
func (m *IAMAPIMock) GetContextKeysForCustomPolicyRequest(p0 *iam.GetContextKeysForCustomPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForCustomPolicyWithContext mocked method
func (m *IAMAPIMock) GetContextKeysForCustomPolicyWithContext(p0 aws.Context, p1 *iam.GetContextKeysForCustomPolicyInput, p2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForPrincipalPolicy mocked method
func (m *IAMAPIMock) GetContextKeysForPrincipalPolicy(p0 *iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error) {

	ret := m.Called(p0)

	var r0 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForPrincipalPolicyRequest mocked method
func (m *IAMAPIMock) GetContextKeysForPrincipalPolicyRequest(p0 *iam.GetContextKeysForPrincipalPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetContextKeysForPrincipalPolicyWithContext mocked method
func (m *IAMAPIMock) GetContextKeysForPrincipalPolicyWithContext(p0 aws.Context, p1 *iam.GetContextKeysForPrincipalPolicyInput, p2 ...request.Option) (*iam.GetContextKeysForPolicyResponse, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetContextKeysForPolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetContextKeysForPolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetCredentialReport mocked method
func (m *IAMAPIMock) GetCredentialReport(p0 *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetCredentialReportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetCredentialReportOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetCredentialReportRequest mocked method
func (m *IAMAPIMock) GetCredentialReportRequest(p0 *iam.GetCredentialReportInput) (*request.Request, *iam.GetCredentialReportOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetCredentialReportOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetCredentialReportOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetCredentialReportWithContext mocked method
func (m *IAMAPIMock) GetCredentialReportWithContext(p0 aws.Context, p1 *iam.GetCredentialReportInput, p2 ...request.Option) (*iam.GetCredentialReportOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetCredentialReportOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetCredentialReportOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroup mocked method
func (m *IAMAPIMock) GetGroup(p0 *iam.GetGroupInput) (*iam.GetGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroupPages mocked method
func (m *IAMAPIMock) GetGroupPages(p0 *iam.GetGroupInput, p1 func(*iam.GetGroupOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// GetGroupPagesWithContext mocked method
func (m *IAMAPIMock) GetGroupPagesWithContext(p0 aws.Context, p1 *iam.GetGroupInput, p2 func(*iam.GetGroupOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// GetGroupPolicy mocked method
func (m *IAMAPIMock) GetGroupPolicy(p0 *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroupPolicyRequest mocked method
func (m *IAMAPIMock) GetGroupPolicyRequest(p0 *iam.GetGroupPolicyInput) (*request.Request, *iam.GetGroupPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetGroupPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetGroupPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroupPolicyWithContext mocked method
func (m *IAMAPIMock) GetGroupPolicyWithContext(p0 aws.Context, p1 *iam.GetGroupPolicyInput, p2 ...request.Option) (*iam.GetGroupPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroupRequest mocked method
func (m *IAMAPIMock) GetGroupRequest(p0 *iam.GetGroupInput) (*request.Request, *iam.GetGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetGroupWithContext mocked method
func (m *IAMAPIMock) GetGroupWithContext(p0 aws.Context, p1 *iam.GetGroupInput, p2 ...request.Option) (*iam.GetGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetInstanceProfile mocked method
func (m *IAMAPIMock) GetInstanceProfile(p0 *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetInstanceProfileRequest mocked method
func (m *IAMAPIMock) GetInstanceProfileRequest(p0 *iam.GetInstanceProfileInput) (*request.Request, *iam.GetInstanceProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetInstanceProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetInstanceProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetInstanceProfileWithContext mocked method
func (m *IAMAPIMock) GetInstanceProfileWithContext(p0 aws.Context, p1 *iam.GetInstanceProfileInput, p2 ...request.Option) (*iam.GetInstanceProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetLoginProfile mocked method
func (m *IAMAPIMock) GetLoginProfile(p0 *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetLoginProfileRequest mocked method
func (m *IAMAPIMock) GetLoginProfileRequest(p0 *iam.GetLoginProfileInput) (*request.Request, *iam.GetLoginProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetLoginProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetLoginProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetLoginProfileWithContext mocked method
func (m *IAMAPIMock) GetLoginProfileWithContext(p0 aws.Context, p1 *iam.GetLoginProfileInput, p2 ...request.Option) (*iam.GetLoginProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetOpenIDConnectProvider mocked method
func (m *IAMAPIMock) GetOpenIDConnectProvider(p0 *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetOpenIDConnectProviderRequest mocked method
func (m *IAMAPIMock) GetOpenIDConnectProviderRequest(p0 *iam.GetOpenIDConnectProviderInput) (*request.Request, *iam.GetOpenIDConnectProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetOpenIDConnectProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetOpenIDConnectProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetOpenIDConnectProviderWithContext mocked method
func (m *IAMAPIMock) GetOpenIDConnectProviderWithContext(p0 aws.Context, p1 *iam.GetOpenIDConnectProviderInput, p2 ...request.Option) (*iam.GetOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicy mocked method
func (m *IAMAPIMock) GetPolicy(p0 *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicyRequest mocked method
func (m *IAMAPIMock) GetPolicyRequest(p0 *iam.GetPolicyInput) (*request.Request, *iam.GetPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicyVersion mocked method
func (m *IAMAPIMock) GetPolicyVersion(p0 *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetPolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetPolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicyVersionRequest mocked method
func (m *IAMAPIMock) GetPolicyVersionRequest(p0 *iam.GetPolicyVersionInput) (*request.Request, *iam.GetPolicyVersionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetPolicyVersionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetPolicyVersionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicyVersionWithContext mocked method
func (m *IAMAPIMock) GetPolicyVersionWithContext(p0 aws.Context, p1 *iam.GetPolicyVersionInput, p2 ...request.Option) (*iam.GetPolicyVersionOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetPolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetPolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetPolicyWithContext mocked method
func (m *IAMAPIMock) GetPolicyWithContext(p0 aws.Context, p1 *iam.GetPolicyInput, p2 ...request.Option) (*iam.GetPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRole mocked method
func (m *IAMAPIMock) GetRole(p0 *iam.GetRoleInput) (*iam.GetRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRolePolicy mocked method
func (m *IAMAPIMock) GetRolePolicy(p0 *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRolePolicyRequest mocked method
func (m *IAMAPIMock) GetRolePolicyRequest(p0 *iam.GetRolePolicyInput) (*request.Request, *iam.GetRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRolePolicyWithContext mocked method
func (m *IAMAPIMock) GetRolePolicyWithContext(p0 aws.Context, p1 *iam.GetRolePolicyInput, p2 ...request.Option) (*iam.GetRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRoleRequest mocked method
func (m *IAMAPIMock) GetRoleRequest(p0 *iam.GetRoleInput) (*request.Request, *iam.GetRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetRoleWithContext mocked method
func (m *IAMAPIMock) GetRoleWithContext(p0 aws.Context, p1 *iam.GetRoleInput, p2 ...request.Option) (*iam.GetRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSAMLProvider mocked method
func (m *IAMAPIMock) GetSAMLProvider(p0 *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSAMLProviderRequest mocked method
func (m *IAMAPIMock) GetSAMLProviderRequest(p0 *iam.GetSAMLProviderInput) (*request.Request, *iam.GetSAMLProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetSAMLProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetSAMLProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSAMLProviderWithContext mocked method
func (m *IAMAPIMock) GetSAMLProviderWithContext(p0 aws.Context, p1 *iam.GetSAMLProviderInput, p2 ...request.Option) (*iam.GetSAMLProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSSHPublicKey mocked method
func (m *IAMAPIMock) GetSSHPublicKey(p0 *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSSHPublicKeyRequest mocked method
func (m *IAMAPIMock) GetSSHPublicKeyRequest(p0 *iam.GetSSHPublicKeyInput) (*request.Request, *iam.GetSSHPublicKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetSSHPublicKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetSSHPublicKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetSSHPublicKeyWithContext mocked method
func (m *IAMAPIMock) GetSSHPublicKeyWithContext(p0 aws.Context, p1 *iam.GetSSHPublicKeyInput, p2 ...request.Option) (*iam.GetSSHPublicKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServerCertificate mocked method
func (m *IAMAPIMock) GetServerCertificate(p0 *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServerCertificateRequest mocked method
func (m *IAMAPIMock) GetServerCertificateRequest(p0 *iam.GetServerCertificateInput) (*request.Request, *iam.GetServerCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetServerCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetServerCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServerCertificateWithContext mocked method
func (m *IAMAPIMock) GetServerCertificateWithContext(p0 aws.Context, p1 *iam.GetServerCertificateInput, p2 ...request.Option) (*iam.GetServerCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServiceLinkedRoleDeletionStatus mocked method
func (m *IAMAPIMock) GetServiceLinkedRoleDeletionStatus(p0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetServiceLinkedRoleDeletionStatusOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetServiceLinkedRoleDeletionStatusOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServiceLinkedRoleDeletionStatusRequest mocked method
func (m *IAMAPIMock) GetServiceLinkedRoleDeletionStatusRequest(p0 *iam.GetServiceLinkedRoleDeletionStatusInput) (*request.Request, *iam.GetServiceLinkedRoleDeletionStatusOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetServiceLinkedRoleDeletionStatusOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetServiceLinkedRoleDeletionStatusOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetServiceLinkedRoleDeletionStatusWithContext mocked method
func (m *IAMAPIMock) GetServiceLinkedRoleDeletionStatusWithContext(p0 aws.Context, p1 *iam.GetServiceLinkedRoleDeletionStatusInput, p2 ...request.Option) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetServiceLinkedRoleDeletionStatusOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetServiceLinkedRoleDeletionStatusOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUser mocked method
func (m *IAMAPIMock) GetUser(p0 *iam.GetUserInput) (*iam.GetUserOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUserPolicy mocked method
func (m *IAMAPIMock) GetUserPolicy(p0 *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.GetUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUserPolicyRequest mocked method
func (m *IAMAPIMock) GetUserPolicyRequest(p0 *iam.GetUserPolicyInput) (*request.Request, *iam.GetUserPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetUserPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetUserPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUserPolicyWithContext mocked method
func (m *IAMAPIMock) GetUserPolicyWithContext(p0 aws.Context, p1 *iam.GetUserPolicyInput, p2 ...request.Option) (*iam.GetUserPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUserRequest mocked method
func (m *IAMAPIMock) GetUserRequest(p0 *iam.GetUserInput) (*request.Request, *iam.GetUserOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.GetUserOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.GetUserOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// GetUserWithContext mocked method
func (m *IAMAPIMock) GetUserWithContext(p0 aws.Context, p1 *iam.GetUserInput, p2 ...request.Option) (*iam.GetUserOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.GetUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.GetUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccessKeys mocked method
func (m *IAMAPIMock) ListAccessKeys(p0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListAccessKeysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAccessKeysOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccessKeysPages mocked method
func (m *IAMAPIMock) ListAccessKeysPages(p0 *iam.ListAccessKeysInput, p1 func(*iam.ListAccessKeysOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAccessKeysPagesWithContext mocked method
func (m *IAMAPIMock) ListAccessKeysPagesWithContext(p0 aws.Context, p1 *iam.ListAccessKeysInput, p2 func(*iam.ListAccessKeysOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAccessKeysRequest mocked method
func (m *IAMAPIMock) ListAccessKeysRequest(p0 *iam.ListAccessKeysInput) (*request.Request, *iam.ListAccessKeysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListAccessKeysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListAccessKeysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccessKeysWithContext mocked method
func (m *IAMAPIMock) ListAccessKeysWithContext(p0 aws.Context, p1 *iam.ListAccessKeysInput, p2 ...request.Option) (*iam.ListAccessKeysOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListAccessKeysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAccessKeysOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccountAliases mocked method
func (m *IAMAPIMock) ListAccountAliases(p0 *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListAccountAliasesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAccountAliasesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccountAliasesPages mocked method
func (m *IAMAPIMock) ListAccountAliasesPages(p0 *iam.ListAccountAliasesInput, p1 func(*iam.ListAccountAliasesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAccountAliasesPagesWithContext mocked method
func (m *IAMAPIMock) ListAccountAliasesPagesWithContext(p0 aws.Context, p1 *iam.ListAccountAliasesInput, p2 func(*iam.ListAccountAliasesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAccountAliasesRequest mocked method
func (m *IAMAPIMock) ListAccountAliasesRequest(p0 *iam.ListAccountAliasesInput) (*request.Request, *iam.ListAccountAliasesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListAccountAliasesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListAccountAliasesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAccountAliasesWithContext mocked method
func (m *IAMAPIMock) ListAccountAliasesWithContext(p0 aws.Context, p1 *iam.ListAccountAliasesInput, p2 ...request.Option) (*iam.ListAccountAliasesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListAccountAliasesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAccountAliasesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedGroupPolicies mocked method
func (m *IAMAPIMock) ListAttachedGroupPolicies(p0 *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListAttachedGroupPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedGroupPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedGroupPoliciesPages mocked method
func (m *IAMAPIMock) ListAttachedGroupPoliciesPages(p0 *iam.ListAttachedGroupPoliciesInput, p1 func(*iam.ListAttachedGroupPoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedGroupPoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListAttachedGroupPoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListAttachedGroupPoliciesInput, p2 func(*iam.ListAttachedGroupPoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedGroupPoliciesRequest mocked method
func (m *IAMAPIMock) ListAttachedGroupPoliciesRequest(p0 *iam.ListAttachedGroupPoliciesInput) (*request.Request, *iam.ListAttachedGroupPoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListAttachedGroupPoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListAttachedGroupPoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedGroupPoliciesWithContext mocked method
func (m *IAMAPIMock) ListAttachedGroupPoliciesWithContext(p0 aws.Context, p1 *iam.ListAttachedGroupPoliciesInput, p2 ...request.Option) (*iam.ListAttachedGroupPoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListAttachedGroupPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedGroupPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedRolePolicies mocked method
func (m *IAMAPIMock) ListAttachedRolePolicies(p0 *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListAttachedRolePoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedRolePoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedRolePoliciesPages mocked method
func (m *IAMAPIMock) ListAttachedRolePoliciesPages(p0 *iam.ListAttachedRolePoliciesInput, p1 func(*iam.ListAttachedRolePoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedRolePoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListAttachedRolePoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListAttachedRolePoliciesInput, p2 func(*iam.ListAttachedRolePoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedRolePoliciesRequest mocked method
func (m *IAMAPIMock) ListAttachedRolePoliciesRequest(p0 *iam.ListAttachedRolePoliciesInput) (*request.Request, *iam.ListAttachedRolePoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListAttachedRolePoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListAttachedRolePoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedRolePoliciesWithContext mocked method
func (m *IAMAPIMock) ListAttachedRolePoliciesWithContext(p0 aws.Context, p1 *iam.ListAttachedRolePoliciesInput, p2 ...request.Option) (*iam.ListAttachedRolePoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListAttachedRolePoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedRolePoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedUserPolicies mocked method
func (m *IAMAPIMock) ListAttachedUserPolicies(p0 *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListAttachedUserPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedUserPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedUserPoliciesPages mocked method
func (m *IAMAPIMock) ListAttachedUserPoliciesPages(p0 *iam.ListAttachedUserPoliciesInput, p1 func(*iam.ListAttachedUserPoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedUserPoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListAttachedUserPoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListAttachedUserPoliciesInput, p2 func(*iam.ListAttachedUserPoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListAttachedUserPoliciesRequest mocked method
func (m *IAMAPIMock) ListAttachedUserPoliciesRequest(p0 *iam.ListAttachedUserPoliciesInput) (*request.Request, *iam.ListAttachedUserPoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListAttachedUserPoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListAttachedUserPoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListAttachedUserPoliciesWithContext mocked method
func (m *IAMAPIMock) ListAttachedUserPoliciesWithContext(p0 aws.Context, p1 *iam.ListAttachedUserPoliciesInput, p2 ...request.Option) (*iam.ListAttachedUserPoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListAttachedUserPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListAttachedUserPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListEntitiesForPolicy mocked method
func (m *IAMAPIMock) ListEntitiesForPolicy(p0 *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListEntitiesForPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListEntitiesForPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListEntitiesForPolicyPages mocked method
func (m *IAMAPIMock) ListEntitiesForPolicyPages(p0 *iam.ListEntitiesForPolicyInput, p1 func(*iam.ListEntitiesForPolicyOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListEntitiesForPolicyPagesWithContext mocked method
func (m *IAMAPIMock) ListEntitiesForPolicyPagesWithContext(p0 aws.Context, p1 *iam.ListEntitiesForPolicyInput, p2 func(*iam.ListEntitiesForPolicyOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListEntitiesForPolicyRequest mocked method
func (m *IAMAPIMock) ListEntitiesForPolicyRequest(p0 *iam.ListEntitiesForPolicyInput) (*request.Request, *iam.ListEntitiesForPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListEntitiesForPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListEntitiesForPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListEntitiesForPolicyWithContext mocked method
func (m *IAMAPIMock) ListEntitiesForPolicyWithContext(p0 aws.Context, p1 *iam.ListEntitiesForPolicyInput, p2 ...request.Option) (*iam.ListEntitiesForPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListEntitiesForPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListEntitiesForPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupPolicies mocked method
func (m *IAMAPIMock) ListGroupPolicies(p0 *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListGroupPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupPoliciesPages mocked method
func (m *IAMAPIMock) ListGroupPoliciesPages(p0 *iam.ListGroupPoliciesInput, p1 func(*iam.ListGroupPoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupPoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListGroupPoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListGroupPoliciesInput, p2 func(*iam.ListGroupPoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupPoliciesRequest mocked method
func (m *IAMAPIMock) ListGroupPoliciesRequest(p0 *iam.ListGroupPoliciesInput) (*request.Request, *iam.ListGroupPoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListGroupPoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListGroupPoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupPoliciesWithContext mocked method
func (m *IAMAPIMock) ListGroupPoliciesWithContext(p0 aws.Context, p1 *iam.ListGroupPoliciesInput, p2 ...request.Option) (*iam.ListGroupPoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListGroupPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroups mocked method
func (m *IAMAPIMock) ListGroups(p0 *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListGroupsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupsForUser mocked method
func (m *IAMAPIMock) ListGroupsForUser(p0 *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListGroupsForUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupsForUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupsForUserPages mocked method
func (m *IAMAPIMock) ListGroupsForUserPages(p0 *iam.ListGroupsForUserInput, p1 func(*iam.ListGroupsForUserOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupsForUserPagesWithContext mocked method
func (m *IAMAPIMock) ListGroupsForUserPagesWithContext(p0 aws.Context, p1 *iam.ListGroupsForUserInput, p2 func(*iam.ListGroupsForUserOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupsForUserRequest mocked method
func (m *IAMAPIMock) ListGroupsForUserRequest(p0 *iam.ListGroupsForUserInput) (*request.Request, *iam.ListGroupsForUserOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListGroupsForUserOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListGroupsForUserOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupsForUserWithContext mocked method
func (m *IAMAPIMock) ListGroupsForUserWithContext(p0 aws.Context, p1 *iam.ListGroupsForUserInput, p2 ...request.Option) (*iam.ListGroupsForUserOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListGroupsForUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupsForUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupsPages mocked method
func (m *IAMAPIMock) ListGroupsPages(p0 *iam.ListGroupsInput, p1 func(*iam.ListGroupsOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupsPagesWithContext mocked method
func (m *IAMAPIMock) ListGroupsPagesWithContext(p0 aws.Context, p1 *iam.ListGroupsInput, p2 func(*iam.ListGroupsOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListGroupsRequest mocked method
func (m *IAMAPIMock) ListGroupsRequest(p0 *iam.ListGroupsInput) (*request.Request, *iam.ListGroupsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListGroupsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListGroupsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListGroupsWithContext mocked method
func (m *IAMAPIMock) ListGroupsWithContext(p0 aws.Context, p1 *iam.ListGroupsInput, p2 ...request.Option) (*iam.ListGroupsOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListGroupsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListGroupsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfiles mocked method
func (m *IAMAPIMock) ListInstanceProfiles(p0 *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListInstanceProfilesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListInstanceProfilesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfilesForRole mocked method
func (m *IAMAPIMock) ListInstanceProfilesForRole(p0 *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListInstanceProfilesForRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListInstanceProfilesForRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfilesForRolePages mocked method
func (m *IAMAPIMock) ListInstanceProfilesForRolePages(p0 *iam.ListInstanceProfilesForRoleInput, p1 func(*iam.ListInstanceProfilesForRoleOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListInstanceProfilesForRolePagesWithContext mocked method
func (m *IAMAPIMock) ListInstanceProfilesForRolePagesWithContext(p0 aws.Context, p1 *iam.ListInstanceProfilesForRoleInput, p2 func(*iam.ListInstanceProfilesForRoleOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListInstanceProfilesForRoleRequest mocked method
func (m *IAMAPIMock) ListInstanceProfilesForRoleRequest(p0 *iam.ListInstanceProfilesForRoleInput) (*request.Request, *iam.ListInstanceProfilesForRoleOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListInstanceProfilesForRoleOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListInstanceProfilesForRoleOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfilesForRoleWithContext mocked method
func (m *IAMAPIMock) ListInstanceProfilesForRoleWithContext(p0 aws.Context, p1 *iam.ListInstanceProfilesForRoleInput, p2 ...request.Option) (*iam.ListInstanceProfilesForRoleOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListInstanceProfilesForRoleOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListInstanceProfilesForRoleOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfilesPages mocked method
func (m *IAMAPIMock) ListInstanceProfilesPages(p0 *iam.ListInstanceProfilesInput, p1 func(*iam.ListInstanceProfilesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListInstanceProfilesPagesWithContext mocked method
func (m *IAMAPIMock) ListInstanceProfilesPagesWithContext(p0 aws.Context, p1 *iam.ListInstanceProfilesInput, p2 func(*iam.ListInstanceProfilesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListInstanceProfilesRequest mocked method
func (m *IAMAPIMock) ListInstanceProfilesRequest(p0 *iam.ListInstanceProfilesInput) (*request.Request, *iam.ListInstanceProfilesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListInstanceProfilesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListInstanceProfilesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListInstanceProfilesWithContext mocked method
func (m *IAMAPIMock) ListInstanceProfilesWithContext(p0 aws.Context, p1 *iam.ListInstanceProfilesInput, p2 ...request.Option) (*iam.ListInstanceProfilesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListInstanceProfilesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListInstanceProfilesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListMFADevices mocked method
func (m *IAMAPIMock) ListMFADevices(p0 *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListMFADevicesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListMFADevicesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListMFADevicesPages mocked method
func (m *IAMAPIMock) ListMFADevicesPages(p0 *iam.ListMFADevicesInput, p1 func(*iam.ListMFADevicesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListMFADevicesPagesWithContext mocked method
func (m *IAMAPIMock) ListMFADevicesPagesWithContext(p0 aws.Context, p1 *iam.ListMFADevicesInput, p2 func(*iam.ListMFADevicesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListMFADevicesRequest mocked method
func (m *IAMAPIMock) ListMFADevicesRequest(p0 *iam.ListMFADevicesInput) (*request.Request, *iam.ListMFADevicesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListMFADevicesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListMFADevicesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListMFADevicesWithContext mocked method
func (m *IAMAPIMock) ListMFADevicesWithContext(p0 aws.Context, p1 *iam.ListMFADevicesInput, p2 ...request.Option) (*iam.ListMFADevicesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListMFADevicesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListMFADevicesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListOpenIDConnectProviders mocked method
func (m *IAMAPIMock) ListOpenIDConnectProviders(p0 *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListOpenIDConnectProvidersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListOpenIDConnectProvidersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListOpenIDConnectProvidersRequest mocked method
func (m *IAMAPIMock) ListOpenIDConnectProvidersRequest(p0 *iam.ListOpenIDConnectProvidersInput) (*request.Request, *iam.ListOpenIDConnectProvidersOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListOpenIDConnectProvidersOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListOpenIDConnectProvidersOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListOpenIDConnectProvidersWithContext mocked method
func (m *IAMAPIMock) ListOpenIDConnectProvidersWithContext(p0 aws.Context, p1 *iam.ListOpenIDConnectProvidersInput, p2 ...request.Option) (*iam.ListOpenIDConnectProvidersOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListOpenIDConnectProvidersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListOpenIDConnectProvidersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPolicies mocked method
func (m *IAMAPIMock) ListPolicies(p0 *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPoliciesPages mocked method
func (m *IAMAPIMock) ListPoliciesPages(p0 *iam.ListPoliciesInput, p1 func(*iam.ListPoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListPoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListPoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListPoliciesInput, p2 func(*iam.ListPoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListPoliciesRequest mocked method
func (m *IAMAPIMock) ListPoliciesRequest(p0 *iam.ListPoliciesInput) (*request.Request, *iam.ListPoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListPoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListPoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPoliciesWithContext mocked method
func (m *IAMAPIMock) ListPoliciesWithContext(p0 aws.Context, p1 *iam.ListPoliciesInput, p2 ...request.Option) (*iam.ListPoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPolicyVersions mocked method
func (m *IAMAPIMock) ListPolicyVersions(p0 *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListPolicyVersionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListPolicyVersionsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPolicyVersionsPages mocked method
func (m *IAMAPIMock) ListPolicyVersionsPages(p0 *iam.ListPolicyVersionsInput, p1 func(*iam.ListPolicyVersionsOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListPolicyVersionsPagesWithContext mocked method
func (m *IAMAPIMock) ListPolicyVersionsPagesWithContext(p0 aws.Context, p1 *iam.ListPolicyVersionsInput, p2 func(*iam.ListPolicyVersionsOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListPolicyVersionsRequest mocked method
func (m *IAMAPIMock) ListPolicyVersionsRequest(p0 *iam.ListPolicyVersionsInput) (*request.Request, *iam.ListPolicyVersionsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListPolicyVersionsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListPolicyVersionsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListPolicyVersionsWithContext mocked method
func (m *IAMAPIMock) ListPolicyVersionsWithContext(p0 aws.Context, p1 *iam.ListPolicyVersionsInput, p2 ...request.Option) (*iam.ListPolicyVersionsOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListPolicyVersionsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListPolicyVersionsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRolePolicies mocked method
func (m *IAMAPIMock) ListRolePolicies(p0 *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListRolePoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListRolePoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRolePoliciesPages mocked method
func (m *IAMAPIMock) ListRolePoliciesPages(p0 *iam.ListRolePoliciesInput, p1 func(*iam.ListRolePoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListRolePoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListRolePoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListRolePoliciesInput, p2 func(*iam.ListRolePoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListRolePoliciesRequest mocked method
func (m *IAMAPIMock) ListRolePoliciesRequest(p0 *iam.ListRolePoliciesInput) (*request.Request, *iam.ListRolePoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListRolePoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListRolePoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRolePoliciesWithContext mocked method
func (m *IAMAPIMock) ListRolePoliciesWithContext(p0 aws.Context, p1 *iam.ListRolePoliciesInput, p2 ...request.Option) (*iam.ListRolePoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListRolePoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListRolePoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRoles mocked method
func (m *IAMAPIMock) ListRoles(p0 *iam.ListRolesInput) (*iam.ListRolesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListRolesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListRolesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRolesPages mocked method
func (m *IAMAPIMock) ListRolesPages(p0 *iam.ListRolesInput, p1 func(*iam.ListRolesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListRolesPagesWithContext mocked method
func (m *IAMAPIMock) ListRolesPagesWithContext(p0 aws.Context, p1 *iam.ListRolesInput, p2 func(*iam.ListRolesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListRolesRequest mocked method
func (m *IAMAPIMock) ListRolesRequest(p0 *iam.ListRolesInput) (*request.Request, *iam.ListRolesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListRolesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListRolesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListRolesWithContext mocked method
func (m *IAMAPIMock) ListRolesWithContext(p0 aws.Context, p1 *iam.ListRolesInput, p2 ...request.Option) (*iam.ListRolesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListRolesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListRolesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSAMLProviders mocked method
func (m *IAMAPIMock) ListSAMLProviders(p0 *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListSAMLProvidersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSAMLProvidersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSAMLProvidersRequest mocked method
func (m *IAMAPIMock) ListSAMLProvidersRequest(p0 *iam.ListSAMLProvidersInput) (*request.Request, *iam.ListSAMLProvidersOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListSAMLProvidersOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListSAMLProvidersOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSAMLProvidersWithContext mocked method
func (m *IAMAPIMock) ListSAMLProvidersWithContext(p0 aws.Context, p1 *iam.ListSAMLProvidersInput, p2 ...request.Option) (*iam.ListSAMLProvidersOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListSAMLProvidersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSAMLProvidersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSSHPublicKeys mocked method
func (m *IAMAPIMock) ListSSHPublicKeys(p0 *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListSSHPublicKeysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSSHPublicKeysOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSSHPublicKeysPages mocked method
func (m *IAMAPIMock) ListSSHPublicKeysPages(p0 *iam.ListSSHPublicKeysInput, p1 func(*iam.ListSSHPublicKeysOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListSSHPublicKeysPagesWithContext mocked method
func (m *IAMAPIMock) ListSSHPublicKeysPagesWithContext(p0 aws.Context, p1 *iam.ListSSHPublicKeysInput, p2 func(*iam.ListSSHPublicKeysOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListSSHPublicKeysRequest mocked method
func (m *IAMAPIMock) ListSSHPublicKeysRequest(p0 *iam.ListSSHPublicKeysInput) (*request.Request, *iam.ListSSHPublicKeysOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListSSHPublicKeysOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListSSHPublicKeysOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSSHPublicKeysWithContext mocked method
func (m *IAMAPIMock) ListSSHPublicKeysWithContext(p0 aws.Context, p1 *iam.ListSSHPublicKeysInput, p2 ...request.Option) (*iam.ListSSHPublicKeysOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListSSHPublicKeysOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSSHPublicKeysOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServerCertificates mocked method
func (m *IAMAPIMock) ListServerCertificates(p0 *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListServerCertificatesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListServerCertificatesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServerCertificatesPages mocked method
func (m *IAMAPIMock) ListServerCertificatesPages(p0 *iam.ListServerCertificatesInput, p1 func(*iam.ListServerCertificatesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListServerCertificatesPagesWithContext mocked method
func (m *IAMAPIMock) ListServerCertificatesPagesWithContext(p0 aws.Context, p1 *iam.ListServerCertificatesInput, p2 func(*iam.ListServerCertificatesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListServerCertificatesRequest mocked method
func (m *IAMAPIMock) ListServerCertificatesRequest(p0 *iam.ListServerCertificatesInput) (*request.Request, *iam.ListServerCertificatesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListServerCertificatesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListServerCertificatesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServerCertificatesWithContext mocked method
func (m *IAMAPIMock) ListServerCertificatesWithContext(p0 aws.Context, p1 *iam.ListServerCertificatesInput, p2 ...request.Option) (*iam.ListServerCertificatesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListServerCertificatesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListServerCertificatesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServiceSpecificCredentials mocked method
func (m *IAMAPIMock) ListServiceSpecificCredentials(p0 *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListServiceSpecificCredentialsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListServiceSpecificCredentialsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServiceSpecificCredentialsRequest mocked method
func (m *IAMAPIMock) ListServiceSpecificCredentialsRequest(p0 *iam.ListServiceSpecificCredentialsInput) (*request.Request, *iam.ListServiceSpecificCredentialsOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListServiceSpecificCredentialsOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListServiceSpecificCredentialsOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListServiceSpecificCredentialsWithContext mocked method
func (m *IAMAPIMock) ListServiceSpecificCredentialsWithContext(p0 aws.Context, p1 *iam.ListServiceSpecificCredentialsInput, p2 ...request.Option) (*iam.ListServiceSpecificCredentialsOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListServiceSpecificCredentialsOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListServiceSpecificCredentialsOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSigningCertificates mocked method
func (m *IAMAPIMock) ListSigningCertificates(p0 *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListSigningCertificatesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSigningCertificatesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSigningCertificatesPages mocked method
func (m *IAMAPIMock) ListSigningCertificatesPages(p0 *iam.ListSigningCertificatesInput, p1 func(*iam.ListSigningCertificatesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListSigningCertificatesPagesWithContext mocked method
func (m *IAMAPIMock) ListSigningCertificatesPagesWithContext(p0 aws.Context, p1 *iam.ListSigningCertificatesInput, p2 func(*iam.ListSigningCertificatesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListSigningCertificatesRequest mocked method
func (m *IAMAPIMock) ListSigningCertificatesRequest(p0 *iam.ListSigningCertificatesInput) (*request.Request, *iam.ListSigningCertificatesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListSigningCertificatesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListSigningCertificatesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListSigningCertificatesWithContext mocked method
func (m *IAMAPIMock) ListSigningCertificatesWithContext(p0 aws.Context, p1 *iam.ListSigningCertificatesInput, p2 ...request.Option) (*iam.ListSigningCertificatesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListSigningCertificatesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListSigningCertificatesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUserPolicies mocked method
func (m *IAMAPIMock) ListUserPolicies(p0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListUserPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListUserPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUserPoliciesPages mocked method
func (m *IAMAPIMock) ListUserPoliciesPages(p0 *iam.ListUserPoliciesInput, p1 func(*iam.ListUserPoliciesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListUserPoliciesPagesWithContext mocked method
func (m *IAMAPIMock) ListUserPoliciesPagesWithContext(p0 aws.Context, p1 *iam.ListUserPoliciesInput, p2 func(*iam.ListUserPoliciesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListUserPoliciesRequest mocked method
func (m *IAMAPIMock) ListUserPoliciesRequest(p0 *iam.ListUserPoliciesInput) (*request.Request, *iam.ListUserPoliciesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListUserPoliciesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListUserPoliciesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUserPoliciesWithContext mocked method
func (m *IAMAPIMock) ListUserPoliciesWithContext(p0 aws.Context, p1 *iam.ListUserPoliciesInput, p2 ...request.Option) (*iam.ListUserPoliciesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListUserPoliciesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListUserPoliciesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUsers mocked method
func (m *IAMAPIMock) ListUsers(p0 *iam.ListUsersInput) (*iam.ListUsersOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListUsersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListUsersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUsersPages mocked method
func (m *IAMAPIMock) ListUsersPages(p0 *iam.ListUsersInput, p1 func(*iam.ListUsersOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListUsersPagesWithContext mocked method
func (m *IAMAPIMock) ListUsersPagesWithContext(p0 aws.Context, p1 *iam.ListUsersInput, p2 func(*iam.ListUsersOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListUsersRequest mocked method
func (m *IAMAPIMock) ListUsersRequest(p0 *iam.ListUsersInput) (*request.Request, *iam.ListUsersOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListUsersOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListUsersOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListUsersWithContext mocked method
func (m *IAMAPIMock) ListUsersWithContext(p0 aws.Context, p1 *iam.ListUsersInput, p2 ...request.Option) (*iam.ListUsersOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListUsersOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListUsersOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListVirtualMFADevices mocked method
func (m *IAMAPIMock) ListVirtualMFADevices(p0 *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ListVirtualMFADevicesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListVirtualMFADevicesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListVirtualMFADevicesPages mocked method
func (m *IAMAPIMock) ListVirtualMFADevicesPages(p0 *iam.ListVirtualMFADevicesInput, p1 func(*iam.ListVirtualMFADevicesOutput, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListVirtualMFADevicesPagesWithContext mocked method
func (m *IAMAPIMock) ListVirtualMFADevicesPagesWithContext(p0 aws.Context, p1 *iam.ListVirtualMFADevicesInput, p2 func(*iam.ListVirtualMFADevicesOutput, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// ListVirtualMFADevicesRequest mocked method
func (m *IAMAPIMock) ListVirtualMFADevicesRequest(p0 *iam.ListVirtualMFADevicesInput) (*request.Request, *iam.ListVirtualMFADevicesOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ListVirtualMFADevicesOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ListVirtualMFADevicesOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ListVirtualMFADevicesWithContext mocked method
func (m *IAMAPIMock) ListVirtualMFADevicesWithContext(p0 aws.Context, p1 *iam.ListVirtualMFADevicesInput, p2 ...request.Option) (*iam.ListVirtualMFADevicesOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ListVirtualMFADevicesOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ListVirtualMFADevicesOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutGroupPolicy mocked method
func (m *IAMAPIMock) PutGroupPolicy(p0 *iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.PutGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutGroupPolicyRequest mocked method
func (m *IAMAPIMock) PutGroupPolicyRequest(p0 *iam.PutGroupPolicyInput) (*request.Request, *iam.PutGroupPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.PutGroupPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.PutGroupPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutGroupPolicyWithContext mocked method
func (m *IAMAPIMock) PutGroupPolicyWithContext(p0 aws.Context, p1 *iam.PutGroupPolicyInput, p2 ...request.Option) (*iam.PutGroupPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.PutGroupPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutGroupPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutRolePolicy mocked method
func (m *IAMAPIMock) PutRolePolicy(p0 *iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.PutRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutRolePolicyRequest mocked method
func (m *IAMAPIMock) PutRolePolicyRequest(p0 *iam.PutRolePolicyInput) (*request.Request, *iam.PutRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.PutRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.PutRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutRolePolicyWithContext mocked method
func (m *IAMAPIMock) PutRolePolicyWithContext(p0 aws.Context, p1 *iam.PutRolePolicyInput, p2 ...request.Option) (*iam.PutRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.PutRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutUserPolicy mocked method
func (m *IAMAPIMock) PutUserPolicy(p0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.PutUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutUserPolicyRequest mocked method
func (m *IAMAPIMock) PutUserPolicyRequest(p0 *iam.PutUserPolicyInput) (*request.Request, *iam.PutUserPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.PutUserPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.PutUserPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// PutUserPolicyWithContext mocked method
func (m *IAMAPIMock) PutUserPolicyWithContext(p0 aws.Context, p1 *iam.PutUserPolicyInput, p2 ...request.Option) (*iam.PutUserPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.PutUserPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.PutUserPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveClientIDFromOpenIDConnectProvider mocked method
func (m *IAMAPIMock) RemoveClientIDFromOpenIDConnectProvider(p0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.RemoveClientIDFromOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveClientIDFromOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveClientIDFromOpenIDConnectProviderRequest mocked method
func (m *IAMAPIMock) RemoveClientIDFromOpenIDConnectProviderRequest(p0 *iam.RemoveClientIDFromOpenIDConnectProviderInput) (*request.Request, *iam.RemoveClientIDFromOpenIDConnectProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.RemoveClientIDFromOpenIDConnectProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.RemoveClientIDFromOpenIDConnectProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveClientIDFromOpenIDConnectProviderWithContext mocked method
func (m *IAMAPIMock) RemoveClientIDFromOpenIDConnectProviderWithContext(p0 aws.Context, p1 *iam.RemoveClientIDFromOpenIDConnectProviderInput, p2 ...request.Option) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.RemoveClientIDFromOpenIDConnectProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveClientIDFromOpenIDConnectProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveRoleFromInstanceProfile mocked method
func (m *IAMAPIMock) RemoveRoleFromInstanceProfile(p0 *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.RemoveRoleFromInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveRoleFromInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveRoleFromInstanceProfileRequest mocked method
func (m *IAMAPIMock) RemoveRoleFromInstanceProfileRequest(p0 *iam.RemoveRoleFromInstanceProfileInput) (*request.Request, *iam.RemoveRoleFromInstanceProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.RemoveRoleFromInstanceProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.RemoveRoleFromInstanceProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveRoleFromInstanceProfileWithContext mocked method
func (m *IAMAPIMock) RemoveRoleFromInstanceProfileWithContext(p0 aws.Context, p1 *iam.RemoveRoleFromInstanceProfileInput, p2 ...request.Option) (*iam.RemoveRoleFromInstanceProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.RemoveRoleFromInstanceProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveRoleFromInstanceProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveUserFromGroup mocked method
func (m *IAMAPIMock) RemoveUserFromGroup(p0 *iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.RemoveUserFromGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveUserFromGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveUserFromGroupRequest mocked method
func (m *IAMAPIMock) RemoveUserFromGroupRequest(p0 *iam.RemoveUserFromGroupInput) (*request.Request, *iam.RemoveUserFromGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.RemoveUserFromGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.RemoveUserFromGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// RemoveUserFromGroupWithContext mocked method
func (m *IAMAPIMock) RemoveUserFromGroupWithContext(p0 aws.Context, p1 *iam.RemoveUserFromGroupInput, p2 ...request.Option) (*iam.RemoveUserFromGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.RemoveUserFromGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.RemoveUserFromGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetServiceSpecificCredential mocked method
func (m *IAMAPIMock) ResetServiceSpecificCredential(p0 *iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ResetServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ResetServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetServiceSpecificCredentialRequest mocked method
func (m *IAMAPIMock) ResetServiceSpecificCredentialRequest(p0 *iam.ResetServiceSpecificCredentialInput) (*request.Request, *iam.ResetServiceSpecificCredentialOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ResetServiceSpecificCredentialOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ResetServiceSpecificCredentialOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResetServiceSpecificCredentialWithContext mocked method
func (m *IAMAPIMock) ResetServiceSpecificCredentialWithContext(p0 aws.Context, p1 *iam.ResetServiceSpecificCredentialInput, p2 ...request.Option) (*iam.ResetServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ResetServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ResetServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResyncMFADevice mocked method
func (m *IAMAPIMock) ResyncMFADevice(p0 *iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.ResyncMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ResyncMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResyncMFADeviceRequest mocked method
func (m *IAMAPIMock) ResyncMFADeviceRequest(p0 *iam.ResyncMFADeviceInput) (*request.Request, *iam.ResyncMFADeviceOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.ResyncMFADeviceOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.ResyncMFADeviceOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// ResyncMFADeviceWithContext mocked method
func (m *IAMAPIMock) ResyncMFADeviceWithContext(p0 aws.Context, p1 *iam.ResyncMFADeviceInput, p2 ...request.Option) (*iam.ResyncMFADeviceOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.ResyncMFADeviceOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.ResyncMFADeviceOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SetDefaultPolicyVersion mocked method
func (m *IAMAPIMock) SetDefaultPolicyVersion(p0 *iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.SetDefaultPolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SetDefaultPolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SetDefaultPolicyVersionRequest mocked method
func (m *IAMAPIMock) SetDefaultPolicyVersionRequest(p0 *iam.SetDefaultPolicyVersionInput) (*request.Request, *iam.SetDefaultPolicyVersionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.SetDefaultPolicyVersionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.SetDefaultPolicyVersionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SetDefaultPolicyVersionWithContext mocked method
func (m *IAMAPIMock) SetDefaultPolicyVersionWithContext(p0 aws.Context, p1 *iam.SetDefaultPolicyVersionInput, p2 ...request.Option) (*iam.SetDefaultPolicyVersionOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.SetDefaultPolicyVersionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SetDefaultPolicyVersionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulateCustomPolicy mocked method
func (m *IAMAPIMock) SimulateCustomPolicy(p0 *iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error) {

	ret := m.Called(p0)

	var r0 *iam.SimulatePolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulateCustomPolicyPages mocked method
func (m *IAMAPIMock) SimulateCustomPolicyPages(p0 *iam.SimulateCustomPolicyInput, p1 func(*iam.SimulatePolicyResponse, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SimulateCustomPolicyPagesWithContext mocked method
func (m *IAMAPIMock) SimulateCustomPolicyPagesWithContext(p0 aws.Context, p1 *iam.SimulateCustomPolicyInput, p2 func(*iam.SimulatePolicyResponse, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SimulateCustomPolicyRequest mocked method
func (m *IAMAPIMock) SimulateCustomPolicyRequest(p0 *iam.SimulateCustomPolicyInput) (*request.Request, *iam.SimulatePolicyResponse) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.SimulatePolicyResponse
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulateCustomPolicyWithContext mocked method
func (m *IAMAPIMock) SimulateCustomPolicyWithContext(p0 aws.Context, p1 *iam.SimulateCustomPolicyInput, p2 ...request.Option) (*iam.SimulatePolicyResponse, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.SimulatePolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulatePrincipalPolicy mocked method
func (m *IAMAPIMock) SimulatePrincipalPolicy(p0 *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {

	ret := m.Called(p0)

	var r0 *iam.SimulatePolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulatePrincipalPolicyPages mocked method
func (m *IAMAPIMock) SimulatePrincipalPolicyPages(p0 *iam.SimulatePrincipalPolicyInput, p1 func(*iam.SimulatePolicyResponse, bool) bool) error {

	ret := m.Called(p0, p1)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SimulatePrincipalPolicyPagesWithContext mocked method
func (m *IAMAPIMock) SimulatePrincipalPolicyPagesWithContext(p0 aws.Context, p1 *iam.SimulatePrincipalPolicyInput, p2 func(*iam.SimulatePolicyResponse, bool) bool, p3 ...request.Option) error {

	ret := m.Called(p0, p1, p2, p3)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// SimulatePrincipalPolicyRequest mocked method
func (m *IAMAPIMock) SimulatePrincipalPolicyRequest(p0 *iam.SimulatePrincipalPolicyInput) (*request.Request, *iam.SimulatePolicyResponse) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.SimulatePolicyResponse
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// SimulatePrincipalPolicyWithContext mocked method
func (m *IAMAPIMock) SimulatePrincipalPolicyWithContext(p0 aws.Context, p1 *iam.SimulatePrincipalPolicyInput, p2 ...request.Option) (*iam.SimulatePolicyResponse, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.SimulatePolicyResponse
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.SimulatePolicyResponse:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccessKey mocked method
func (m *IAMAPIMock) UpdateAccessKey(p0 *iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccessKeyRequest mocked method
func (m *IAMAPIMock) UpdateAccessKeyRequest(p0 *iam.UpdateAccessKeyInput) (*request.Request, *iam.UpdateAccessKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateAccessKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateAccessKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccessKeyWithContext mocked method
func (m *IAMAPIMock) UpdateAccessKeyWithContext(p0 aws.Context, p1 *iam.UpdateAccessKeyInput, p2 ...request.Option) (*iam.UpdateAccessKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateAccessKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAccessKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccountPasswordPolicy mocked method
func (m *IAMAPIMock) UpdateAccountPasswordPolicy(p0 *iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccountPasswordPolicyRequest mocked method
func (m *IAMAPIMock) UpdateAccountPasswordPolicyRequest(p0 *iam.UpdateAccountPasswordPolicyInput) (*request.Request, *iam.UpdateAccountPasswordPolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateAccountPasswordPolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateAccountPasswordPolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAccountPasswordPolicyWithContext mocked method
func (m *IAMAPIMock) UpdateAccountPasswordPolicyWithContext(p0 aws.Context, p1 *iam.UpdateAccountPasswordPolicyInput, p2 ...request.Option) (*iam.UpdateAccountPasswordPolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateAccountPasswordPolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAccountPasswordPolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAssumeRolePolicy mocked method
func (m *IAMAPIMock) UpdateAssumeRolePolicy(p0 *iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateAssumeRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAssumeRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAssumeRolePolicyRequest mocked method
func (m *IAMAPIMock) UpdateAssumeRolePolicyRequest(p0 *iam.UpdateAssumeRolePolicyInput) (*request.Request, *iam.UpdateAssumeRolePolicyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateAssumeRolePolicyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateAssumeRolePolicyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateAssumeRolePolicyWithContext mocked method
func (m *IAMAPIMock) UpdateAssumeRolePolicyWithContext(p0 aws.Context, p1 *iam.UpdateAssumeRolePolicyInput, p2 ...request.Option) (*iam.UpdateAssumeRolePolicyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateAssumeRolePolicyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateAssumeRolePolicyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateGroup mocked method
func (m *IAMAPIMock) UpdateGroup(p0 *iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateGroupRequest mocked method
func (m *IAMAPIMock) UpdateGroupRequest(p0 *iam.UpdateGroupInput) (*request.Request, *iam.UpdateGroupOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateGroupOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateGroupOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateGroupWithContext mocked method
func (m *IAMAPIMock) UpdateGroupWithContext(p0 aws.Context, p1 *iam.UpdateGroupInput, p2 ...request.Option) (*iam.UpdateGroupOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateGroupOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateGroupOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateLoginProfile mocked method
func (m *IAMAPIMock) UpdateLoginProfile(p0 *iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateLoginProfileRequest mocked method
func (m *IAMAPIMock) UpdateLoginProfileRequest(p0 *iam.UpdateLoginProfileInput) (*request.Request, *iam.UpdateLoginProfileOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateLoginProfileOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateLoginProfileOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateLoginProfileWithContext mocked method
func (m *IAMAPIMock) UpdateLoginProfileWithContext(p0 aws.Context, p1 *iam.UpdateLoginProfileInput, p2 ...request.Option) (*iam.UpdateLoginProfileOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateLoginProfileOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateLoginProfileOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateOpenIDConnectProviderThumbprint mocked method
func (m *IAMAPIMock) UpdateOpenIDConnectProviderThumbprint(p0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateOpenIDConnectProviderThumbprintOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateOpenIDConnectProviderThumbprintOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateOpenIDConnectProviderThumbprintRequest mocked method
func (m *IAMAPIMock) UpdateOpenIDConnectProviderThumbprintRequest(p0 *iam.UpdateOpenIDConnectProviderThumbprintInput) (*request.Request, *iam.UpdateOpenIDConnectProviderThumbprintOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateOpenIDConnectProviderThumbprintOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateOpenIDConnectProviderThumbprintOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateOpenIDConnectProviderThumbprintWithContext mocked method
func (m *IAMAPIMock) UpdateOpenIDConnectProviderThumbprintWithContext(p0 aws.Context, p1 *iam.UpdateOpenIDConnectProviderThumbprintInput, p2 ...request.Option) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateOpenIDConnectProviderThumbprintOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateOpenIDConnectProviderThumbprintOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateRoleDescription mocked method
func (m *IAMAPIMock) UpdateRoleDescription(p0 *iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateRoleDescriptionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateRoleDescriptionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateRoleDescriptionRequest mocked method
func (m *IAMAPIMock) UpdateRoleDescriptionRequest(p0 *iam.UpdateRoleDescriptionInput) (*request.Request, *iam.UpdateRoleDescriptionOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateRoleDescriptionOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateRoleDescriptionOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateRoleDescriptionWithContext mocked method
func (m *IAMAPIMock) UpdateRoleDescriptionWithContext(p0 aws.Context, p1 *iam.UpdateRoleDescriptionInput, p2 ...request.Option) (*iam.UpdateRoleDescriptionOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateRoleDescriptionOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateRoleDescriptionOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSAMLProvider mocked method
func (m *IAMAPIMock) UpdateSAMLProvider(p0 *iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSAMLProviderRequest mocked method
func (m *IAMAPIMock) UpdateSAMLProviderRequest(p0 *iam.UpdateSAMLProviderInput) (*request.Request, *iam.UpdateSAMLProviderOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateSAMLProviderOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateSAMLProviderOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSAMLProviderWithContext mocked method
func (m *IAMAPIMock) UpdateSAMLProviderWithContext(p0 aws.Context, p1 *iam.UpdateSAMLProviderInput, p2 ...request.Option) (*iam.UpdateSAMLProviderOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateSAMLProviderOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSAMLProviderOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSSHPublicKey mocked method
func (m *IAMAPIMock) UpdateSSHPublicKey(p0 *iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSSHPublicKeyRequest mocked method
func (m *IAMAPIMock) UpdateSSHPublicKeyRequest(p0 *iam.UpdateSSHPublicKeyInput) (*request.Request, *iam.UpdateSSHPublicKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateSSHPublicKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateSSHPublicKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSSHPublicKeyWithContext mocked method
func (m *IAMAPIMock) UpdateSSHPublicKeyWithContext(p0 aws.Context, p1 *iam.UpdateSSHPublicKeyInput, p2 ...request.Option) (*iam.UpdateSSHPublicKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServerCertificate mocked method
func (m *IAMAPIMock) UpdateServerCertificate(p0 *iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServerCertificateRequest mocked method
func (m *IAMAPIMock) UpdateServerCertificateRequest(p0 *iam.UpdateServerCertificateInput) (*request.Request, *iam.UpdateServerCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateServerCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateServerCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServerCertificateWithContext mocked method
func (m *IAMAPIMock) UpdateServerCertificateWithContext(p0 aws.Context, p1 *iam.UpdateServerCertificateInput, p2 ...request.Option) (*iam.UpdateServerCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServiceSpecificCredential mocked method
func (m *IAMAPIMock) UpdateServiceSpecificCredential(p0 *iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServiceSpecificCredentialRequest mocked method
func (m *IAMAPIMock) UpdateServiceSpecificCredentialRequest(p0 *iam.UpdateServiceSpecificCredentialInput) (*request.Request, *iam.UpdateServiceSpecificCredentialOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateServiceSpecificCredentialOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateServiceSpecificCredentialOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateServiceSpecificCredentialWithContext mocked method
func (m *IAMAPIMock) UpdateServiceSpecificCredentialWithContext(p0 aws.Context, p1 *iam.UpdateServiceSpecificCredentialInput, p2 ...request.Option) (*iam.UpdateServiceSpecificCredentialOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateServiceSpecificCredentialOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateServiceSpecificCredentialOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSigningCertificate mocked method
func (m *IAMAPIMock) UpdateSigningCertificate(p0 *iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSigningCertificateRequest mocked method
func (m *IAMAPIMock) UpdateSigningCertificateRequest(p0 *iam.UpdateSigningCertificateInput) (*request.Request, *iam.UpdateSigningCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateSigningCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateSigningCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateSigningCertificateWithContext mocked method
func (m *IAMAPIMock) UpdateSigningCertificateWithContext(p0 aws.Context, p1 *iam.UpdateSigningCertificateInput, p2 ...request.Option) (*iam.UpdateSigningCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateUser mocked method
func (m *IAMAPIMock) UpdateUser(p0 *iam.UpdateUserInput) (*iam.UpdateUserOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UpdateUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateUserRequest mocked method
func (m *IAMAPIMock) UpdateUserRequest(p0 *iam.UpdateUserInput) (*request.Request, *iam.UpdateUserOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UpdateUserOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UpdateUserOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UpdateUserWithContext mocked method
func (m *IAMAPIMock) UpdateUserWithContext(p0 aws.Context, p1 *iam.UpdateUserInput, p2 ...request.Option) (*iam.UpdateUserOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UpdateUserOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UpdateUserOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSSHPublicKey mocked method
func (m *IAMAPIMock) UploadSSHPublicKey(p0 *iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UploadSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSSHPublicKeyRequest mocked method
func (m *IAMAPIMock) UploadSSHPublicKeyRequest(p0 *iam.UploadSSHPublicKeyInput) (*request.Request, *iam.UploadSSHPublicKeyOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UploadSSHPublicKeyOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UploadSSHPublicKeyOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSSHPublicKeyWithContext mocked method
func (m *IAMAPIMock) UploadSSHPublicKeyWithContext(p0 aws.Context, p1 *iam.UploadSSHPublicKeyInput, p2 ...request.Option) (*iam.UploadSSHPublicKeyOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UploadSSHPublicKeyOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadSSHPublicKeyOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadServerCertificate mocked method
func (m *IAMAPIMock) UploadServerCertificate(p0 *iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UploadServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadServerCertificateRequest mocked method
func (m *IAMAPIMock) UploadServerCertificateRequest(p0 *iam.UploadServerCertificateInput) (*request.Request, *iam.UploadServerCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UploadServerCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UploadServerCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadServerCertificateWithContext mocked method
func (m *IAMAPIMock) UploadServerCertificateWithContext(p0 aws.Context, p1 *iam.UploadServerCertificateInput, p2 ...request.Option) (*iam.UploadServerCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UploadServerCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadServerCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSigningCertificate mocked method
func (m *IAMAPIMock) UploadSigningCertificate(p0 *iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error) {

	ret := m.Called(p0)

	var r0 *iam.UploadSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSigningCertificateRequest mocked method
func (m *IAMAPIMock) UploadSigningCertificateRequest(p0 *iam.UploadSigningCertificateInput) (*request.Request, *iam.UploadSigningCertificateOutput) {

	ret := m.Called(p0)

	var r0 *request.Request
	switch res := ret.Get(0).(type) {
	case nil:
	case *request.Request:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 *iam.UploadSigningCertificateOutput
	switch res := ret.Get(1).(type) {
	case nil:
	case *iam.UploadSigningCertificateOutput:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// UploadSigningCertificateWithContext mocked method
func (m *IAMAPIMock) UploadSigningCertificateWithContext(p0 aws.Context, p1 *iam.UploadSigningCertificateInput, p2 ...request.Option) (*iam.UploadSigningCertificateOutput, error) {

	ret := m.Called(p0, p1, p2)

	var r0 *iam.UploadSigningCertificateOutput
	switch res := ret.Get(0).(type) {
	case nil:
	case *iam.UploadSigningCertificateOutput:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	var r1 error
	switch res := ret.Get(1).(type) {
	case nil:
	case error:
		r1 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0, r1

}

// WaitUntilInstanceProfileExists mocked method
func (m *IAMAPIMock) WaitUntilInstanceProfileExists(p0 *iam.GetInstanceProfileInput) error {

	ret := m.Called(p0)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// WaitUntilInstanceProfileExistsWithContext mocked method
func (m *IAMAPIMock) WaitUntilInstanceProfileExistsWithContext(p0 aws.Context, p1 *iam.GetInstanceProfileInput, p2 ...request.WaiterOption) error {

	ret := m.Called(p0, p1, p2)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// WaitUntilUserExists mocked method
func (m *IAMAPIMock) WaitUntilUserExists(p0 *iam.GetUserInput) error {

	ret := m.Called(p0)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}

// WaitUntilUserExistsWithContext mocked method
func (m *IAMAPIMock) WaitUntilUserExistsWithContext(p0 aws.Context, p1 *iam.GetUserInput, p2 ...request.WaiterOption) error {

	ret := m.Called(p0, p1, p2)

	var r0 error
	switch res := ret.Get(0).(type) {
	case nil:
	case error:
		r0 = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}

	return r0

}
