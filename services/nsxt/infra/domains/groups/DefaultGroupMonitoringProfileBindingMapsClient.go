
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: GroupMonitoringProfileBindingMaps
 * Functions that implement the generated GroupMonitoringProfileBindingMapsClient interface
 */


package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultGroupMonitoringProfileBindingMapsClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultGroupMonitoringProfileBindingMapsClient(connector client.Connector) *DefaultGroupMonitoringProfileBindingMapsClient {
	interfaceName := "com.vmware.nsx_policy.infra.domains.groups.group_monitoring_profile_binding_maps"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	gIface := DefaultGroupMonitoringProfileBindingMapsClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	gIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	gIface.methodNameToDefMap["delete"] = gIface.deleteMethodDefinition()
	gIface.methodNameToDefMap["get"] = gIface.getMethodDefinition()
	gIface.methodNameToDefMap["list"] = gIface.listMethodDefinition()
	gIface.methodNameToDefMap["patch"] = gIface.patchMethodDefinition()
	gIface.methodNameToDefMap["update"] = gIface.updateMethodDefinition()
	return &gIface
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) Delete(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) error {
	typeConverter := gIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(gIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMapId", groupMonitoringProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.Invoke(gIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) Get(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) (model.GroupMonitoringProfileBindingMap, error) {
	typeConverter := gIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(gIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMapId", groupMonitoringProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMonitoringProfileBindingMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.Invoke(gIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.GroupMonitoringProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupMonitoringProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMonitoringProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.GroupMonitoringProfileBindingMapListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(gIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMonitoringProfileBindingMapListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.Invoke(gIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.GroupMonitoringProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupMonitoringProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMonitoringProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) Patch(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string, groupMonitoringProfileBindingMapParam model.GroupMonitoringProfileBindingMap) error {
	typeConverter := gIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(gIface.interfaceIdentifier, "patch")
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMapId", groupMonitoringProfileBindingMapIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMap", groupMonitoringProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.Invoke(gIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) Update(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string, groupMonitoringProfileBindingMapParam model.GroupMonitoringProfileBindingMap) (model.GroupMonitoringProfileBindingMap, error) {
	typeConverter := gIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(gIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMapId", groupMonitoringProfileBindingMapIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMap", groupMonitoringProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMonitoringProfileBindingMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.Invoke(gIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.GroupMonitoringProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupMonitoringProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMonitoringProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := gIface.connector.GetApiProvider().Invoke(gIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := gIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	gIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := gIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	gIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := gIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	gIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.list method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) patchMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := gIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsPatchInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsPatchOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "patch")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	gIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.patch method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (gIface *DefaultGroupMonitoringProfileBindingMapsClient) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(gIface.interfaceName)
	typeConverter := gIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(groupMonitoringProfileBindingMapsUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	gIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's errors.InvalidRequest error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	gIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultGroupMonitoringProfileBindingMapsClient.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
