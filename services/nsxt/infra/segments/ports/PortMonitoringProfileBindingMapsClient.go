/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PortMonitoringProfileBindingMaps
 * Used by client-side stubs.
 */

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type PortMonitoringProfileBindingMapsClient interface {

    // API will delete Infra Port Monitoring Profile Binding Profile.
    //
    // @param infraSegmentIdParam Infra Segment ID (required)
    // @param infraPortIdParam Infra Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) error

    // API will get Infra Port Monitoring Profile Binding Map.
    //
    // @param infraSegmentIdParam Infra Segment ID (required)
    // @param infraPortIdParam Infra Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) (model.PortMonitoringProfileBindingMap, error)

    // API will list all Infra Port Monitoring Profile Binding Maps in current port id.
    //
    // @param infraSegmentIdParam (required)
    // @param infraPortIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(infraSegmentIdParam string, infraPortIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortMonitoringProfileBindingMapListResult, error)

    // API will create Infra Port Monitoring Profile Binding Map.
    //
    // @param infraSegmentIdParam Infra Segment ID (required)
    // @param infraPortIdParam Infra Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @param portMonitoringProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam model.PortMonitoringProfileBindingMap) error

    // API will update Infra Port Monitoring Profile Binding Map.
    //
    // @param infraSegmentIdParam InfraSegment ID (required)
    // @param infraPortIdParam Infra Port ID (required)
    // @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
    // @param portMonitoringProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam model.PortMonitoringProfileBindingMap) (model.PortMonitoringProfileBindingMap, error)
}