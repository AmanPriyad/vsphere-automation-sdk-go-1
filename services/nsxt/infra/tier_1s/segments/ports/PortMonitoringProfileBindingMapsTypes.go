/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: PortMonitoringProfileBindingMaps.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ports

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func portMonitoringProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMonitoringProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMonitoringProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMonitoringProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapListResultBindingType)
}

func portMonitoringProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMonitoringProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMonitoringProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"port_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMonitoringProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"port_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/segments/{segmentId}/ports/{portId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


