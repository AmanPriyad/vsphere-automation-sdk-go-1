/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Dns.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package management_vms

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``ipType`` of method Dns#update.
const Dns_UPDATE_IP_TYPE_PUBLIC = "public"
// Possible value for ``ipType`` of method Dns#update.
const Dns_UPDATE_IP_TYPE_PRIVATE = "private"




func dnsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["management_vm_id"] = bindings.NewStringType()
	fields["ip_type"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["management_vm_id"] = "ManagementVmId"
	fieldNameMap["ip_type"] = "IpType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dnsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TaskBindingType)
}

func dnsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["management_vm_id"] = bindings.NewStringType()
	fields["ip_type"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["management_vm_id"] = "ManagementVmId"
	fieldNameMap["ip_type"] = "IpType"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["management_vm_id"] = bindings.NewStringType()
	paramsTypeMap["ip_type"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["managementVmId"] = bindings.NewStringType()
	paramsTypeMap["ipType"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["management_vm_id"] = "managementVmId"
	pathParams["ip_type"] = "ipType"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"PUT",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/management-vms/{managementVmId}/dns/{ipType}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403})
}


