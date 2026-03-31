// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package ecloudsdkcomputer

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcomputer/model"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type Client struct {
	apiClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *request.HttpRequest
	allRegions  map[string]string
}

func NewClient(config *config.Config) *Client {
	httpRequest := request.DefaultHttpRequest()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	ecloudsdkcore.InitConfig(config)
	apiClient := ecloudsdkcore.DefaultApiClient(config, httpRequest)
	client := &Client{
		apiClient:   apiClient,
		config:      config,
		httpRequest: httpRequest,
	}
	client.allRegions = client.initRegions()
	client.setEndpoint(config, httpRequest)
	return client
}

const (
    product    string = "Computer"
    version           = "v1"
    sdkVersion        = "1.0.17"
)

func (c *Client) initRegions() map[string]string {
	m := map[string]string{
        "CIDC-RP-25" : "https://api-wuxi-1.cmecloud.cn:8443",
        "CIDC-RP-26" : "https://api-dongguan-1.cmecloud.cn:8443",
        "CIDC-RP-27" : "https://api-yaan-1.cmecloud.cn:8443",
        "CIDC-RP-28" : "https://api-zhengzhou-1.cmecloud.cn:8443",
        "CIDC-RP-29" : "https://api-beijing-2.cmecloud.cn:8443",
        "CIDC-RP-30" : "https://api-zhuzhou-1.cmecloud.cn:8443",
        "CIDC-RP-31" : "https://api-jinan-1.cmecloud.cn:8443",
        "CIDC-RP-32" : "https://api-xian-1.cmecloud.cn:8443",
        "CIDC-RP-33" : "https://api-shanghai-1.cmecloud.cn:8443",
        "CIDC-RP-34" : "https://api-chongqing-1.cmecloud.cn:8443",
        "CIDC-RP-35" : "https://api-ningbo-1.cmecloud.cn:8443",
        "CIDC-RP-36" : "https://api-tianjin-1.cmecloud.cn:8443",
        "CIDC-RP-37" : "https://api-jilin-1.cmecloud.cn:8443",
        "CIDC-RP-38" : "https://api-hubei-1.cmecloud.cn:8443",
        "CIDC-RP-39" : "https://api-jiangxi-1.cmecloud.cn:8443",
        "CIDC-RP-40" : "https://api-gansu-1.cmecloud.cn:8443",
        "CIDC-RP-41" : "https://api-shanxi-1.cmecloud.cn:8443",
        "CIDC-RP-42" : "https://api-liaoning-1.cmecloud.cn:8443",
        "CIDC-RP-43" : "https://api-yunnan-2.cmecloud.cn:8443",
        "CIDC-RP-44" : "https://api-hebei-1.cmecloud.cn:8443",
        "CIDC-RP-45" : "https://api-fujian-1.cmecloud.cn:8443",
        "CIDC-RP-46" : "https://api-guangxi-1.cmecloud.cn:8443",
        "CIDC-RP-47" : "https://api-anhui-1.cmecloud.cn:8443",
        "CIDC-RP-48" : "https://api-huhehaote-1.cmecloud.cn:8443",
        "CIDC-RP-49" : "https://api-guiyang-1.cmecloud.cn:8443",
        "CIDC-RP-53" : "https://api-hainan-1.cmecloud.cn:8443",
        "CIDC-RP-54" : "https://api-xinjiang-1.cmecloud.cn:8443",
        "CIDC-RP-55" : "https://api-heilongjiang-1.cmecloud.cn:8443",
        "CIDC-RP-60" : "https://api-ningxia-1.cmecloud.cn:8443",
        "CIDC-RP-61" : "https://api-qinghai-1.cmecloud.cn:8443",
        "CIDC-RP-62" : "https://api-xizang-1.cmecloud.cn:8443",
        "CIDC-RP-64" : "https://api-wuhan-1.cmecloud.cn:8443",
        "CIDC-RP-68" : "https://api-xizang-2.cmecloud.cn:8443",
        "CIDC-RP-69" : "https://api-qinghai-2.cmecloud.cn:8443",
        "CIDC-RP-71" : "https://api-ningxia-2.cmecloud.cn:8443",
        "CIDC-CORE-00" : "https://ecloud.10086.cn",
	}
	return m
}

func (c *Client) setEndpoint(config *config.Config, httpRequest *request.HttpRequest) {
	if utils.IsUnSet(config.PoolId) {
		httpRequest.Endpoint = utils.DefaultEndpoint
		return
	}
	endpoint := c.allRegions[*config.PoolId]
	if utils.IsUnSet(endpoint) {
		httpRequest.Endpoint = utils.DefaultEndpoint
		return
	}
	httpRequest.Endpoint = endpoint
}

// ShareImageOperation 发起共享镜像
func (c *Client) ShareImageOperation(request *model.ShareImageOperationRequest) (*model.ShareImageOperationResponse, error) {
    return c.ShareImageOperationWithConfig(request, nil)
}

// ShareImageOperationWithConfig 发起共享镜像
func (c *Client) ShareImageOperationWithConfig(request *model.ShareImageOperationRequest, runtimeConfig *config.RuntimeConfig) (*model.ShareImageOperationResponse, error) {
    params := param.NewParamsBuilder().
        Action("shareImageOperation").
        Uri("/cem_openapi/ccacem/api/v1/image/shareImageOperation").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/shareImageOperation").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ShareImageOperationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchUnbindPolicy 移除云电脑
func (c *Client) BatchUnbindPolicy(request *model.BatchUnbindPolicyRequest) (*model.BatchUnbindPolicyResponse, error) {
    return c.BatchUnbindPolicyWithConfig(request, nil)
}

// BatchUnbindPolicyWithConfig 移除云电脑
func (c *Client) BatchUnbindPolicyWithConfig(request *model.BatchUnbindPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchUnbindPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchUnbindPolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/batchUnbindPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/batchUnbindPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchUnbindPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Revoke 撤回消息
func (c *Client) Revoke(request *model.RevokeRequest) (*model.RevokeResponse, error) {
    return c.RevokeWithConfig(request, nil)
}

// RevokeWithConfig 撤回消息
func (c *Client) RevokeWithConfig(request *model.RevokeRequest, runtimeConfig *config.RuntimeConfig) (*model.RevokeResponse, error) {
    params := param.NewParamsBuilder().
        Action("revoke").
        Uri("/cem_openapi/ccacem/api/v1//message/revoke").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/revoke").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RevokeResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ProductsBw 查询带宽可用区及规格
func (c *Client) ProductsBw() (*model.ProductsBwResponse, error) {
    return c.ProductsBwWithConfig(nil)
}

// ProductsBwWithConfig 查询带宽可用区及规格
func (c *Client) ProductsBwWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ProductsBwResponse, error) {
    params := param.NewParamsBuilder().
        Action("productsBw").
        Uri("/ccabusiorder/api/v1/order/bw/productsBw").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/productsBw").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.ProductsBwResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ListSoftwareUnifyApi 查询一体化服务包软件信息
func (c *Client) ListSoftwareUnifyApi(request *model.ListSoftwareUnifyApiRequest) (*model.ListSoftwareUnifyApiResponse, error) {
    return c.ListSoftwareUnifyApiWithConfig(request, nil)
}

// ListSoftwareUnifyApiWithConfig 查询一体化服务包软件信息
func (c *Client) ListSoftwareUnifyApiWithConfig(request *model.ListSoftwareUnifyApiRequest, runtimeConfig *config.RuntimeConfig) (*model.ListSoftwareUnifyApiResponse, error) {
    params := param.NewParamsBuilder().
        Action("listSoftwareUnifyApi").
        Uri("/ccabusiorder/api/v1/order/unify/softwares").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/unify/softwares").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListSoftwareUnifyApiResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateImage 创建镜像
func (c *Client) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
    return c.CreateImageWithConfig(request, nil)
}

// CreateImageWithConfig 创建镜像
func (c *Client) CreateImageWithConfig(request *model.CreateImageRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("createImage").
        Uri("/cem_openapi/ccacem/api/v1/image/createImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/createImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteImage 删除镜像
func (c *Client) DeleteImage(request *model.DeleteImageRequest) (*model.DeleteImageResponse, error) {
    return c.DeleteImageWithConfig(request, nil)
}

// DeleteImageWithConfig 删除镜像
func (c *Client) DeleteImageWithConfig(request *model.DeleteImageRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteImage").
        Uri("/cem_openapi/ccacem/api/v1/image/deleteImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/deleteImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteOrg 删除组织
func (c *Client) DeleteOrg(request *model.DeleteOrgRequest) (*model.DeleteOrgResponse, error) {
    return c.DeleteOrgWithConfig(request, nil)
}

// DeleteOrgWithConfig 删除组织
func (c *Client) DeleteOrgWithConfig(request *model.DeleteOrgRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteOrgResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteOrg").
        Uri("/cem_openapi/ccacem/api/v1/userManage/deleteOrg").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/deleteOrg").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteOrgResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// PolicyDesc 查询定时任务详情
func (c *Client) PolicyDesc(request *model.PolicyDescRequest) (*model.PolicyDescResponse, error) {
    return c.PolicyDescWithConfig(request, nil)
}

// PolicyDescWithConfig 查询定时任务详情
func (c *Client) PolicyDescWithConfig(request *model.PolicyDescRequest, runtimeConfig *config.RuntimeConfig) (*model.PolicyDescResponse, error) {
    params := param.NewParamsBuilder().
        Action("policyDesc").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/policyDesc").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/policyDesc").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.PolicyDescResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SendDetail 查看消息发送详情
func (c *Client) SendDetail(request *model.SendDetailRequest) (*model.SendDetailResponse, error) {
    return c.SendDetailWithConfig(request, nil)
}

// SendDetailWithConfig 查看消息发送详情
func (c *Client) SendDetailWithConfig(request *model.SendDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.SendDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("sendDetail").
        Uri("/cem_openapi/ccacem/api/v1//message/sendDetail").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/sendDetail").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SendDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Update 消息修改
func (c *Client) Update(request *model.UpdateRequest) (*model.UpdateResponse, error) {
    return c.UpdateWithConfig(request, nil)
}

// UpdateWithConfig 消息修改
func (c *Client) UpdateWithConfig(request *model.UpdateRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateResponse, error) {
    params := param.NewParamsBuilder().
        Action("update").
        Uri("/cem_openapi/ccacem/api/v1//message/update").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/update").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnsubscribeRenew 退订待生效续订云电脑
func (c *Client) UnsubscribeRenew(request *model.UnsubscribeRenewRequest) (*model.UnsubscribeRenewResponse, error) {
    return c.UnsubscribeRenewWithConfig(request, nil)
}

// UnsubscribeRenewWithConfig 退订待生效续订云电脑
func (c *Client) UnsubscribeRenewWithConfig(request *model.UnsubscribeRenewRequest, runtimeConfig *config.RuntimeConfig) (*model.UnsubscribeRenewResponse, error) {
    params := param.NewParamsBuilder().
        Action("unsubscribeRenew").
        Uri("/ccabusiorder/api/v1/order/unsubscribeRenew").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/unsubscribeRenew").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnsubscribeRenewResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddOrg 创建组织
func (c *Client) AddOrg(request *model.AddOrgRequest) (*model.AddOrgResponse, error) {
    return c.AddOrgWithConfig(request, nil)
}

// AddOrgWithConfig 创建组织
func (c *Client) AddOrgWithConfig(request *model.AddOrgRequest, runtimeConfig *config.RuntimeConfig) (*model.AddOrgResponse, error) {
    params := param.NewParamsBuilder().
        Action("addOrg").
        Uri("/cem_openapi/ccacem/api/v1/userManage/createOrg").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/createOrg").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddOrgResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteRule 删除规则
func (c *Client) DeleteRule(request *model.DeleteRuleRequest) (*model.DeleteRuleResponse, error) {
    return c.DeleteRuleWithConfig(request, nil)
}

// DeleteRuleWithConfig 删除规则
func (c *Client) DeleteRuleWithConfig(request *model.DeleteRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/deleteRule").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/deleteRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchCreateImage 批量创建自定义镜像
func (c *Client) BatchCreateImage(request *model.BatchCreateImageRequest) (*model.BatchCreateImageResponse, error) {
    return c.BatchCreateImageWithConfig(request, nil)
}

// BatchCreateImageWithConfig 批量创建自定义镜像
func (c *Client) BatchCreateImageWithConfig(request *model.BatchCreateImageRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchCreateImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchCreateImage").
        Uri("/cem_openapi/ccacem/api/v1/image/batchCreateImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/batchCreateImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchCreateImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// HistoryPage 查询历史会话
func (c *Client) HistoryPage(request *model.HistoryPageRequest) (*model.HistoryPageResponse, error) {
    return c.HistoryPageWithConfig(request, nil)
}

// HistoryPageWithConfig 查询历史会话
func (c *Client) HistoryPageWithConfig(request *model.HistoryPageRequest, runtimeConfig *config.RuntimeConfig) (*model.HistoryPageResponse, error) {
    params := param.NewParamsBuilder().
        Action("historyPage").
        Uri("/cem_openapi/ccacem/api/v1/operation/loginSession/historyPage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/loginSession/historyPage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.HistoryPageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// EditUserPolicy 编辑用户策略
func (c *Client) EditUserPolicy(request *model.EditUserPolicyRequest) (*model.EditUserPolicyResponse, error) {
    return c.EditUserPolicyWithConfig(request, nil)
}

// EditUserPolicyWithConfig 编辑用户策略
func (c *Client) EditUserPolicyWithConfig(request *model.EditUserPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.EditUserPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("editUserPolicy").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/editUserPolicy").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/userPolicy/editUserPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.EditUserPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ReloadImage 切换镜像(自定义/共享)
func (c *Client) ReloadImage(request *model.ReloadImageRequest) (*model.ReloadImageResponse, error) {
    return c.ReloadImageWithConfig(request, nil)
}

// ReloadImageWithConfig 切换镜像(自定义/共享)
func (c *Client) ReloadImageWithConfig(request *model.ReloadImageRequest, runtimeConfig *config.RuntimeConfig) (*model.ReloadImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("reloadImage").
        Uri("/cem_openapi/ccacem/api/v1/image/reloadImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/reloadImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ReloadImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetZoneVpc 查询网关可用区下vpc
func (c *Client) GetZoneVpc(request *model.GetZoneVpcRequest) (*model.GetZoneVpcResponse, error) {
    return c.GetZoneVpcWithConfig(request, nil)
}

// GetZoneVpcWithConfig 查询网关可用区下vpc
func (c *Client) GetZoneVpcWithConfig(request *model.GetZoneVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.GetZoneVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("getZoneVpc").
        Uri("/ccabusiorder/api/v1/order/nat/getZoneVpc").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/getZoneVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetZoneVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SecurityGroupManagementList 查看安全组列表
func (c *Client) SecurityGroupManagementList(request *model.SecurityGroupManagementListRequest) (*model.SecurityGroupManagementListResponse, error) {
    return c.SecurityGroupManagementListWithConfig(request, nil)
}

// SecurityGroupManagementListWithConfig 查看安全组列表
func (c *Client) SecurityGroupManagementListWithConfig(request *model.SecurityGroupManagementListRequest, runtimeConfig *config.RuntimeConfig) (*model.SecurityGroupManagementListResponse, error) {
    params := param.NewParamsBuilder().
        Action("securityGroupManagementList").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/securityGroupManagementList").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/securityGroupManagementList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SecurityGroupManagementListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddMacAddress 单个配置Mac地址白名单
func (c *Client) AddMacAddress(request *model.AddMacAddressRequest) (*model.AddMacAddressResponse, error) {
    return c.AddMacAddressWithConfig(request, nil)
}

// AddMacAddressWithConfig 单个配置Mac地址白名单
func (c *Client) AddMacAddressWithConfig(request *model.AddMacAddressRequest, runtimeConfig *config.RuntimeConfig) (*model.AddMacAddressResponse, error) {
    params := param.NewParamsBuilder().
        Action("addMacAddress").
        Uri("/cem_openapi/ccacem/api/v1/machine/addMacAddress").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/addMacAddress").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddMacAddressResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetDeviceDetail 查询硬终端详情
func (c *Client) GetDeviceDetail(request *model.GetDeviceDetailRequest) (*model.GetDeviceDetailResponse, error) {
    return c.GetDeviceDetailWithConfig(request, nil)
}

// GetDeviceDetailWithConfig 查询硬终端详情
func (c *Client) GetDeviceDetailWithConfig(request *model.GetDeviceDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.GetDeviceDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("getDeviceDetail").
        Uri("/cem_openapi/ccacem/api/v1/device/ccemp/getDeviceDetail").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/device/ccemp/getDeviceDetail").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetDeviceDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnsubscribeNat 退订nat
func (c *Client) UnsubscribeNat(request *model.UnsubscribeNatRequest) (*model.UnsubscribeNatResponse, error) {
    return c.UnsubscribeNatWithConfig(request, nil)
}

// UnsubscribeNatWithConfig 退订nat
func (c *Client) UnsubscribeNatWithConfig(request *model.UnsubscribeNatRequest, runtimeConfig *config.RuntimeConfig) (*model.UnsubscribeNatResponse, error) {
    params := param.NewParamsBuilder().
        Action("unsubscribeNat").
        Uri("/ccabusiorder/api/v1/order/nat/unsubscribeNat").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/unsubscribeNat").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnsubscribeNatResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnbindUserPolicy 用户策略解绑用户
func (c *Client) UnbindUserPolicy(request *model.UnbindUserPolicyRequest) (*model.UnbindUserPolicyResponse, error) {
    return c.UnbindUserPolicyWithConfig(request, nil)
}

// UnbindUserPolicyWithConfig 用户策略解绑用户
func (c *Client) UnbindUserPolicyWithConfig(request *model.UnbindUserPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.UnbindUserPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("unbindUserPolicy").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/unbindUserPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/unbindUserPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnbindUserPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ReloadByPublicImage 切换镜像(公共)
func (c *Client) ReloadByPublicImage(request *model.ReloadByPublicImageRequest) (*model.ReloadByPublicImageResponse, error) {
    return c.ReloadByPublicImageWithConfig(request, nil)
}

// ReloadByPublicImageWithConfig 切换镜像(公共)
func (c *Client) ReloadByPublicImageWithConfig(request *model.ReloadByPublicImageRequest, runtimeConfig *config.RuntimeConfig) (*model.ReloadByPublicImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("reloadByPublicImage").
        Uri("/cem_openapi/ccacem/api/v1/image/reloadByPublicImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/reloadByPublicImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ReloadByPublicImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ListServiceUnifyApi 查询一体化服务包载体及可选服务包信息
func (c *Client) ListServiceUnifyApi() (*model.ListServiceUnifyApiResponse, error) {
    return c.ListServiceUnifyApiWithConfig(nil)
}

// ListServiceUnifyApiWithConfig 查询一体化服务包载体及可选服务包信息
func (c *Client) ListServiceUnifyApiWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ListServiceUnifyApiResponse, error) {
    params := param.NewParamsBuilder().
        Action("listServiceUnifyApi").
        Uri("/ccabusiorder/api/v1/order/unify/services").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/unify/services").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.ListServiceUnifyApiResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ImageSharedRecord2Own 切换共享镜像列表查询(别人共享给自己，且自己已经接收)
func (c *Client) ImageSharedRecord2Own(request *model.ImageSharedRecord2OwnRequest) (*model.ImageSharedRecord2OwnResponse, error) {
    return c.ImageSharedRecord2OwnWithConfig(request, nil)
}

// ImageSharedRecord2OwnWithConfig 切换共享镜像列表查询(别人共享给自己，且自己已经接收)
func (c *Client) ImageSharedRecord2OwnWithConfig(request *model.ImageSharedRecord2OwnRequest, runtimeConfig *config.RuntimeConfig) (*model.ImageSharedRecord2OwnResponse, error) {
    params := param.NewParamsBuilder().
        Action("imageSharedRecord2Own").
        Uri("/cem_openapi/ccacem/api/v1/image/shareImageList").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/image/shareImageList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ImageSharedRecord2OwnResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreatePolicy 创建电脑策略
func (c *Client) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
    return c.CreatePolicyWithConfig(request, nil)
}

// CreatePolicyWithConfig 创建电脑策略
func (c *Client) CreatePolicyWithConfig(request *model.CreatePolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.CreatePolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("createPolicy").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/createPolicy").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/createPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreatePolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ResourceList 查询绑定时下拉的电脑列表
func (c *Client) ResourceList(request *model.ResourceListRequest) (*model.ResourceListResponse, error) {
    return c.ResourceListWithConfig(request, nil)
}

// ResourceListWithConfig 查询绑定时下拉的电脑列表
func (c *Client) ResourceListWithConfig(request *model.ResourceListRequest, runtimeConfig *config.RuntimeConfig) (*model.ResourceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("resourceList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/resourceList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/resourceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ResourceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetShutdownRevertConfig 查询电脑关机还原配置
func (c *Client) GetShutdownRevertConfig(request *model.GetShutdownRevertConfigRequest) (*model.GetShutdownRevertConfigResponse, error) {
    return c.GetShutdownRevertConfigWithConfig(request, nil)
}

// GetShutdownRevertConfigWithConfig 查询电脑关机还原配置
func (c *Client) GetShutdownRevertConfigWithConfig(request *model.GetShutdownRevertConfigRequest, runtimeConfig *config.RuntimeConfig) (*model.GetShutdownRevertConfigResponse, error) {
    params := param.NewParamsBuilder().
        Action("getShutdownRevertConfig").
        Uri("/cem_openapi/ccacem/api/v1/machine/getShutdownRevertConfig").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/getShutdownRevertConfig").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetShutdownRevertConfigResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchRestart 云电脑批量重启接口
func (c *Client) BatchRestart(request *model.BatchRestartRequest) (*model.BatchRestartResponse, error) {
    return c.BatchRestartWithConfig(request, nil)
}

// BatchRestartWithConfig 云电脑批量重启接口
func (c *Client) BatchRestartWithConfig(request *model.BatchRestartRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchRestartResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchRestart").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchRestart").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchRestart").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchRestartResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SubmitCcaOrderGov 订购政企版云电脑
func (c *Client) SubmitCcaOrderGov(request *model.SubmitCcaOrderGovRequest) (*model.SubmitCcaOrderGovResponse, error) {
    return c.SubmitCcaOrderGovWithConfig(request, nil)
}

// SubmitCcaOrderGovWithConfig 订购政企版云电脑
func (c *Client) SubmitCcaOrderGovWithConfig(request *model.SubmitCcaOrderGovRequest, runtimeConfig *config.RuntimeConfig) (*model.SubmitCcaOrderGovResponse, error) {
    params := param.NewParamsBuilder().
        Action("submitCcaOrderGov").
        Uri("/ccabusiorder/api/v1/order/gov/submit").
        GatewayUri("/api/clouddesktopnew/ccabusiorder/api/v1/order/gov/submit").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SubmitCcaOrderGovResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPolicyInfoById 查看电脑策略详情
func (c *Client) GetPolicyInfoById(request *model.GetPolicyInfoByIdRequest) (*model.GetPolicyInfoByIdResponse, error) {
    return c.GetPolicyInfoByIdWithConfig(request, nil)
}

// GetPolicyInfoByIdWithConfig 查看电脑策略详情
func (c *Client) GetPolicyInfoByIdWithConfig(request *model.GetPolicyInfoByIdRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPolicyInfoByIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPolicyInfoById").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyInfoById").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyInfoById").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetPolicyInfoByIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPolicyRelatedMachineList 查询电脑策略已关联电脑
func (c *Client) GetPolicyRelatedMachineList(request *model.GetPolicyRelatedMachineListRequest) (*model.GetPolicyRelatedMachineListResponse, error) {
    return c.GetPolicyRelatedMachineListWithConfig(request, nil)
}

// GetPolicyRelatedMachineListWithConfig 查询电脑策略已关联电脑
func (c *Client) GetPolicyRelatedMachineListWithConfig(request *model.GetPolicyRelatedMachineListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPolicyRelatedMachineListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPolicyRelatedMachineList").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyRelatedMachineList").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyRelatedMachineList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetPolicyRelatedMachineListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// VerifyAccessTicket 获取accessToken接口
func (c *Client) VerifyAccessTicket(request *model.VerifyAccessTicketRequest) (*model.VerifyAccessTicketResponse, error) {
    return c.VerifyAccessTicketWithConfig(request, nil)
}

// VerifyAccessTicketWithConfig 获取accessToken接口
func (c *Client) VerifyAccessTicketWithConfig(request *model.VerifyAccessTicketRequest, runtimeConfig *config.RuntimeConfig) (*model.VerifyAccessTicketResponse, error) {
    params := param.NewParamsBuilder().
        Action("verifyAccessTicket").
        Uri("/cem_openapi/ccacem/api/v1/auth/verifyAccessTicket").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/auth/verifyAccessTicket").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.VerifyAccessTicketResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetEnforcePasswordPeriodMonth 密码强制更新配置
func (c *Client) SetEnforcePasswordPeriodMonth(request *model.SetEnforcePasswordPeriodMonthRequest) (*model.SetEnforcePasswordPeriodMonthResponse, error) {
    return c.SetEnforcePasswordPeriodMonthWithConfig(request, nil)
}

// SetEnforcePasswordPeriodMonthWithConfig 密码强制更新配置
func (c *Client) SetEnforcePasswordPeriodMonthWithConfig(request *model.SetEnforcePasswordPeriodMonthRequest, runtimeConfig *config.RuntimeConfig) (*model.SetEnforcePasswordPeriodMonthResponse, error) {
    params := param.NewParamsBuilder().
        Action("setEnforcePasswordPeriodMonth").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setEnforcePasswordPeriodMonth").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setEnforcePasswordPeriodMonth").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetEnforcePasswordPeriodMonthResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetFirstLoginModifyPassword 首次登录强制修改密码开关查询
func (c *Client) GetFirstLoginModifyPassword() (*model.GetFirstLoginModifyPasswordResponse, error) {
    return c.GetFirstLoginModifyPasswordWithConfig(nil)
}

// GetFirstLoginModifyPasswordWithConfig 首次登录强制修改密码开关查询
func (c *Client) GetFirstLoginModifyPasswordWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetFirstLoginModifyPasswordResponse, error) {
    params := param.NewParamsBuilder().
        Action("getFirstLoginModifyPassword").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getFirstLoginModifyPassword").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getFirstLoginModifyPassword").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetFirstLoginModifyPasswordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetNetwork 查询子网列表
func (c *Client) GetNetwork(request *model.GetNetworkRequest) (*model.GetNetworkResponse, error) {
    return c.GetNetworkWithConfig(request, nil)
}

// GetNetworkWithConfig 查询子网列表
func (c *Client) GetNetworkWithConfig(request *model.GetNetworkRequest, runtimeConfig *config.RuntimeConfig) (*model.GetNetworkResponse, error) {
    params := param.NewParamsBuilder().
        Action("getNetwork").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/getNetwork").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/getNetwork").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetNetworkResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ListSnapshot 查询快照列表
func (c *Client) ListSnapshot(request *model.ListSnapshotRequest) (*model.ListSnapshotResponse, error) {
    return c.ListSnapshotWithConfig(request, nil)
}

// ListSnapshotWithConfig 查询快照列表
func (c *Client) ListSnapshotWithConfig(request *model.ListSnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.ListSnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("listSnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/listSnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/listSnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ListSnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchUnbindDevice 批量移除限制登录硬终端
func (c *Client) BatchUnbindDevice(request *model.BatchUnbindDeviceRequest) (*model.BatchUnbindDeviceResponse, error) {
    return c.BatchUnbindDeviceWithConfig(request, nil)
}

// BatchUnbindDeviceWithConfig 批量移除限制登录硬终端
func (c *Client) BatchUnbindDeviceWithConfig(request *model.BatchUnbindDeviceRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchUnbindDeviceResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchUnbindDevice").
        Uri("/cem_openapi/ccacem/api/v1/device/batchUnbindDevice").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/batchUnbindDevice").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchUnbindDeviceResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StopAutoRenew 关闭自动续订
func (c *Client) StopAutoRenew(request *model.StopAutoRenewRequest) (*model.StopAutoRenewResponse, error) {
    return c.StopAutoRenewWithConfig(request, nil)
}

// StopAutoRenewWithConfig 关闭自动续订
func (c *Client) StopAutoRenewWithConfig(request *model.StopAutoRenewRequest, runtimeConfig *config.RuntimeConfig) (*model.StopAutoRenewResponse, error) {
    params := param.NewParamsBuilder().
        Action("stopAutoRenew").
        Uri("/ccabusiorder/api/v1/order/stopAutoRenew").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/stopAutoRenew").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.StopAutoRenewResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateVpc 编辑vpc
func (c *Client) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
    return c.UpdateVpcWithConfig(request, nil)
}

// UpdateVpcWithConfig 编辑vpc
func (c *Client) UpdateVpcWithConfig(request *model.UpdateVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateVpc").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/updateVpc").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/updateVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SubmitCcaOrder 订购公众版云电脑
func (c *Client) SubmitCcaOrder(request *model.SubmitCcaOrderRequest) (*model.SubmitCcaOrderResponse, error) {
    return c.SubmitCcaOrderWithConfig(request, nil)
}

// SubmitCcaOrderWithConfig 订购公众版云电脑
func (c *Client) SubmitCcaOrderWithConfig(request *model.SubmitCcaOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.SubmitCcaOrderResponse, error) {
    params := param.NewParamsBuilder().
        Action("submitCcaOrder").
        Uri("/ccabusiorder/api/v1/order/submit").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/submit").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SubmitCcaOrderResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetVpcSelect 切换VPC下拉框
func (c *Client) GetVpcSelect(request *model.GetVpcSelectRequest) (*model.GetVpcSelectResponse, error) {
    return c.GetVpcSelectWithConfig(request, nil)
}

// GetVpcSelectWithConfig 切换VPC下拉框
func (c *Client) GetVpcSelectWithConfig(request *model.GetVpcSelectRequest, runtimeConfig *config.RuntimeConfig) (*model.GetVpcSelectResponse, error) {
    params := param.NewParamsBuilder().
        Action("getVpcSelect").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/getVpcSelect").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/getVpcSelect").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetVpcSelectResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// LoginLivingPageDaily 查询日登录活跃度统计
func (c *Client) LoginLivingPageDaily(request *model.LoginLivingPageDailyRequest) (*model.LoginLivingPageDailyResponse, error) {
    return c.LoginLivingPageDailyWithConfig(request, nil)
}

// LoginLivingPageDailyWithConfig 查询日登录活跃度统计
func (c *Client) LoginLivingPageDailyWithConfig(request *model.LoginLivingPageDailyRequest, runtimeConfig *config.RuntimeConfig) (*model.LoginLivingPageDailyResponse, error) {
    params := param.NewParamsBuilder().
        Action("loginLivingPageDaily").
        Uri("/cem_openapi/ccacem/api/v1/operation/loginStatistics/daily").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/loginStatistics/daily").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LoginLivingPageDailyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// RenewComputer 续订云电脑
func (c *Client) RenewComputer(request *model.RenewComputerRequest) (*model.RenewComputerResponse, error) {
    return c.RenewComputerWithConfig(request, nil)
}

// RenewComputerWithConfig 续订云电脑
func (c *Client) RenewComputerWithConfig(request *model.RenewComputerRequest, runtimeConfig *config.RuntimeConfig) (*model.RenewComputerResponse, error) {
    params := param.NewParamsBuilder().
        Action("renewComputer").
        Uri("/ccabusiorder/api/v1/order/renew").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/renew").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RenewComputerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SmsRetry 重发激活链接
func (c *Client) SmsRetry(request *model.SmsRetryRequest) (*model.SmsRetryResponse, error) {
    return c.SmsRetryWithConfig(request, nil)
}

// SmsRetryWithConfig 重发激活链接
func (c *Client) SmsRetryWithConfig(request *model.SmsRetryRequest, runtimeConfig *config.RuntimeConfig) (*model.SmsRetryResponse, error) {
    params := param.NewParamsBuilder().
        Action("smsRetry").
        Uri("/ccabusiorder/api/v1/channel/operate/activeRetry").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/channel/operate/activeRetry").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SmsRetryResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteComputer 退订云电脑
func (c *Client) DeleteComputer(request *model.DeleteComputerRequest) (*model.DeleteComputerResponse, error) {
    return c.DeleteComputerWithConfig(request, nil)
}

// DeleteComputerWithConfig 退订云电脑
func (c *Client) DeleteComputerWithConfig(request *model.DeleteComputerRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteComputerResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteComputer").
        Uri("/ccabusiorder/api/v1/order/unsubscribe").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/unsubscribe").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteComputerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchApplySnapshot 批量还原快照
func (c *Client) BatchApplySnapshot(request *model.BatchApplySnapshotRequest) (*model.BatchApplySnapshotResponse, error) {
    return c.BatchApplySnapshotWithConfig(request, nil)
}

// BatchApplySnapshotWithConfig 批量还原快照
func (c *Client) BatchApplySnapshotWithConfig(request *model.BatchApplySnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchApplySnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchApplySnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/batchApplySnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/batchApplySnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchApplySnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPublicImagePage 查询公共镜像列表
func (c *Client) GetPublicImagePage(request *model.GetPublicImagePageRequest) (*model.GetPublicImagePageResponse, error) {
    return c.GetPublicImagePageWithConfig(request, nil)
}

// GetPublicImagePageWithConfig 查询公共镜像列表
func (c *Client) GetPublicImagePageWithConfig(request *model.GetPublicImagePageRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPublicImagePageResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPublicImagePage").
        Uri("/cem_openapi/ccacem/api/v1/image/publicImagePage").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/image/publicImagePage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetPublicImagePageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateSecurityGroup 编辑安全组
func (c *Client) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
    return c.UpdateSecurityGroupWithConfig(request, nil)
}

// UpdateSecurityGroupWithConfig 编辑安全组
func (c *Client) UpdateSecurityGroupWithConfig(request *model.UpdateSecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateSecurityGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateSecurityGroup").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/updateSecurityGroup").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/updateSecurityGroup").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateSecurityGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchCreateSnapshot 批量创建快照
func (c *Client) BatchCreateSnapshot(request *model.BatchCreateSnapshotRequest) (*model.BatchCreateSnapshotResponse, error) {
    return c.BatchCreateSnapshotWithConfig(request, nil)
}

// BatchCreateSnapshotWithConfig 批量创建快照
func (c *Client) BatchCreateSnapshotWithConfig(request *model.BatchCreateSnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchCreateSnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchCreateSnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/batchCreateSnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/batchCreateSnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchCreateSnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// EditOrg 编辑组织
func (c *Client) EditOrg(request *model.EditOrgRequest) (*model.EditOrgResponse, error) {
    return c.EditOrgWithConfig(request, nil)
}

// EditOrgWithConfig 编辑组织
func (c *Client) EditOrgWithConfig(request *model.EditOrgRequest, runtimeConfig *config.RuntimeConfig) (*model.EditOrgResponse, error) {
    params := param.NewParamsBuilder().
        Action("editOrg").
        Uri("/cem_openapi/ccacem/api/v1/userManage/editOrg").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/editOrg").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.EditOrgResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUserPolicyDetails 查询用户策略详情
func (c *Client) GetUserPolicyDetails(request *model.GetUserPolicyDetailsRequest) (*model.GetUserPolicyDetailsResponse, error) {
    return c.GetUserPolicyDetailsWithConfig(request, nil)
}

// GetUserPolicyDetailsWithConfig 查询用户策略详情
func (c *Client) GetUserPolicyDetailsWithConfig(request *model.GetUserPolicyDetailsRequest, runtimeConfig *config.RuntimeConfig) (*model.GetUserPolicyDetailsResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUserPolicyDetails").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/getUserPolicyDetails").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/getUserPolicyDetails").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetUserPolicyDetailsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SingleUnifyRenew 续订一体化服务包
func (c *Client) SingleUnifyRenew(request *model.SingleUnifyRenewRequest) (*model.SingleUnifyRenewResponse, error) {
    return c.SingleUnifyRenewWithConfig(request, nil)
}

// SingleUnifyRenewWithConfig 续订一体化服务包
func (c *Client) SingleUnifyRenewWithConfig(request *model.SingleUnifyRenewRequest, runtimeConfig *config.RuntimeConfig) (*model.SingleUnifyRenewResponse, error) {
    params := param.NewParamsBuilder().
        Action("singleUnifyRenew").
        Uri("/ccabusiorder/api/v1/order/singleUnifyRenew").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/singleUnifyRenew").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SingleUnifyRenewResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchAvailable 云电脑批量开机接口
func (c *Client) BatchAvailable(request *model.BatchAvailableRequest) (*model.BatchAvailableResponse, error) {
    return c.BatchAvailableWithConfig(request, nil)
}

// BatchAvailableWithConfig 云电脑批量开机接口
func (c *Client) BatchAvailableWithConfig(request *model.BatchAvailableRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchAvailableResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchAvailable").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchAvailable").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchAvailable").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchAvailableResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OrderChange 变更云电脑
func (c *Client) OrderChange(request *model.OrderChangeRequest) (*model.OrderChangeResponse, error) {
    return c.OrderChangeWithConfig(request, nil)
}

// OrderChangeWithConfig 变更云电脑
func (c *Client) OrderChangeWithConfig(request *model.OrderChangeRequest, runtimeConfig *config.RuntimeConfig) (*model.OrderChangeResponse, error) {
    params := param.NewParamsBuilder().
        Action("orderChange").
        Uri("/ccabusiorder/api/v1/order/change").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/change").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.OrderChangeResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnifyInstanceQuery 查询一体化服务包订单
func (c *Client) UnifyInstanceQuery(request *model.UnifyInstanceQueryRequest) (*model.UnifyInstanceQueryResponse, error) {
    return c.UnifyInstanceQueryWithConfig(request, nil)
}

// UnifyInstanceQueryWithConfig 查询一体化服务包订单
func (c *Client) UnifyInstanceQueryWithConfig(request *model.UnifyInstanceQueryRequest, runtimeConfig *config.RuntimeConfig) (*model.UnifyInstanceQueryResponse, error) {
    params := param.NewParamsBuilder().
        Action("unifyInstanceQuery").
        Uri("/ccabusiorder/api/v1/channel/unify/unifyInstanceQuery").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/channel/unify/unifyInstanceQuery").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnifyInstanceQueryResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ModifyDeliveryInfo 修改一体化服务包收货地址
func (c *Client) ModifyDeliveryInfo(request *model.ModifyDeliveryInfoRequest) (*model.ModifyDeliveryInfoResponse, error) {
    return c.ModifyDeliveryInfoWithConfig(request, nil)
}

// ModifyDeliveryInfoWithConfig 修改一体化服务包收货地址
func (c *Client) ModifyDeliveryInfoWithConfig(request *model.ModifyDeliveryInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.ModifyDeliveryInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("modifyDeliveryInfo").
        Uri("/ccabusiorder/api/v1/channel/unify/modifyDeliveryInfo").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/channel/unify/modifyDeliveryInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ModifyDeliveryInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ConfirmDoSign 签收一体化服务包订单
func (c *Client) ConfirmDoSign(request *model.ConfirmDoSignRequest) (*model.ConfirmDoSignResponse, error) {
    return c.ConfirmDoSignWithConfig(request, nil)
}

// ConfirmDoSignWithConfig 签收一体化服务包订单
func (c *Client) ConfirmDoSignWithConfig(request *model.ConfirmDoSignRequest, runtimeConfig *config.RuntimeConfig) (*model.ConfirmDoSignResponse, error) {
    params := param.NewParamsBuilder().
        Action("confirmDoSign").
        Uri("/ccabusiorder/api/v1/channel/unify/confirmDoSign").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/channel/unify/confirmDoSign").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ConfirmDoSignResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetInstanceListCem 查询云电脑订单列表
func (c *Client) GetInstanceListCem(request *model.GetInstanceListCemRequest) (*model.GetInstanceListCemResponse, error) {
    return c.GetInstanceListCemWithConfig(request, nil)
}

// GetInstanceListCemWithConfig 查询云电脑订单列表
func (c *Client) GetInstanceListCemWithConfig(request *model.GetInstanceListCemRequest, runtimeConfig *config.RuntimeConfig) (*model.GetInstanceListCemResponse, error) {
    params := param.NewParamsBuilder().
        Action("getInstanceListCem").
        Uri("/ccabusiorder/api/v1/order/instanceList").
        GatewayUri("/api/clouddesktopnew/ccabusiorder/api/v1/order/instanceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetInstanceListCemResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetNatList 查询nat列表
func (c *Client) GetNatList(request *model.GetNatListRequest) (*model.GetNatListResponse, error) {
    return c.GetNatListWithConfig(request, nil)
}

// GetNatListWithConfig 查询nat列表
func (c *Client) GetNatListWithConfig(request *model.GetNatListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetNatListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getNatList").
        Uri("/ccabusiorder/api/v1/order/nat/getNatList").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/getNatList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetNatListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetShutdownRecordList 查询定时关机任务执行记录
func (c *Client) GetShutdownRecordList(request *model.GetShutdownRecordListRequest) (*model.GetShutdownRecordListResponse, error) {
    return c.GetShutdownRecordListWithConfig(request, nil)
}

// GetShutdownRecordListWithConfig 查询定时关机任务执行记录
func (c *Client) GetShutdownRecordListWithConfig(request *model.GetShutdownRecordListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetShutdownRecordListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getShutdownRecordList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getShutdownRecordList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getShutdownRecordList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetShutdownRecordListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UserLoginInfo 查询用户日志
func (c *Client) UserLoginInfo(request *model.UserLoginInfoRequest) (*model.UserLoginInfoResponse, error) {
    return c.UserLoginInfoWithConfig(request, nil)
}

// UserLoginInfoWithConfig 查询用户日志
func (c *Client) UserLoginInfoWithConfig(request *model.UserLoginInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.UserLoginInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("userLoginInfo").
        Uri("/cem_openapi/ccacem/api/v1/operation/userLoginInfo/page").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/operation/userLoginInfo/page").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UserLoginInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UserExdpolicyPermission 查询指定资源池/可用区/网络工作区和指定协议支持的电脑策略功能配置
func (c *Client) UserExdpolicyPermission(request *model.UserExdpolicyPermissionRequest) (*model.UserExdpolicyPermissionResponse, error) {
    return c.UserExdpolicyPermissionWithConfig(request, nil)
}

// UserExdpolicyPermissionWithConfig 查询指定资源池/可用区/网络工作区和指定协议支持的电脑策略功能配置
func (c *Client) UserExdpolicyPermissionWithConfig(request *model.UserExdpolicyPermissionRequest, runtimeConfig *config.RuntimeConfig) (*model.UserExdpolicyPermissionResponse, error) {
    params := param.NewParamsBuilder().
        Action("userExdpolicyPermission").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/userExdpolicyPermission").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/userExdpolicyPermission").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UserExdpolicyPermissionResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StartAutoRenew 开启自动续订
func (c *Client) StartAutoRenew(request *model.StartAutoRenewRequest) (*model.StartAutoRenewResponse, error) {
    return c.StartAutoRenewWithConfig(request, nil)
}

// StartAutoRenewWithConfig 开启自动续订
func (c *Client) StartAutoRenewWithConfig(request *model.StartAutoRenewRequest, runtimeConfig *config.RuntimeConfig) (*model.StartAutoRenewResponse, error) {
    params := param.NewParamsBuilder().
        Action("startAutoRenew").
        Uri("/ccabusiorder/api/v1/order/startAutoRenew").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/startAutoRenew").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.StartAutoRenewResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// PolicyList 查询定时任务列表
func (c *Client) PolicyList(request *model.PolicyListRequest) (*model.PolicyListResponse, error) {
    return c.PolicyListWithConfig(request, nil)
}

// PolicyListWithConfig 查询定时任务列表
func (c *Client) PolicyListWithConfig(request *model.PolicyListRequest, runtimeConfig *config.RuntimeConfig) (*model.PolicyListResponse, error) {
    params := param.NewParamsBuilder().
        Action("policyList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/policyList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/policyList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.PolicyListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteVpc 删除vpc
func (c *Client) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
    return c.DeleteVpcWithConfig(request, nil)
}

// DeleteVpcWithConfig 删除vpc
func (c *Client) DeleteVpcWithConfig(request *model.DeleteVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteVpc").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/deleteVpc").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/deleteVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetUserPolicyStatus 修改用户策略生效开关
func (c *Client) SetUserPolicyStatus(request *model.SetUserPolicyStatusRequest) (*model.SetUserPolicyStatusResponse, error) {
    return c.SetUserPolicyStatusWithConfig(request, nil)
}

// SetUserPolicyStatusWithConfig 修改用户策略生效开关
func (c *Client) SetUserPolicyStatusWithConfig(request *model.SetUserPolicyStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.SetUserPolicyStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("setUserPolicyStatus").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/setUserPolicyStatus").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/setUserPolicyStatus").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetUserPolicyStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUnbindMobileFlag 用户解绑手机号开关查询
func (c *Client) GetUnbindMobileFlag() (*model.GetUnbindMobileFlagResponse, error) {
    return c.GetUnbindMobileFlagWithConfig(nil)
}

// GetUnbindMobileFlagWithConfig 用户解绑手机号开关查询
func (c *Client) GetUnbindMobileFlagWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetUnbindMobileFlagResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUnbindMobileFlag").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getUnbindMobileFlag").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getUnbindMobileFlag").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetUnbindMobileFlagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeletePolicy 删除定时任务
func (c *Client) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
    return c.DeletePolicyWithConfig(request, nil)
}

// DeletePolicyWithConfig 删除定时任务
func (c *Client) DeletePolicyWithConfig(request *model.DeletePolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.DeletePolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("deletePolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/deletePolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/deletePolicy").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DeletePolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUserPolicyList 查询用户策略列表信息
func (c *Client) GetUserPolicyList(request *model.GetUserPolicyListRequest) (*model.GetUserPolicyListResponse, error) {
    return c.GetUserPolicyListWithConfig(request, nil)
}

// GetUserPolicyListWithConfig 查询用户策略列表信息
func (c *Client) GetUserPolicyListWithConfig(request *model.GetUserPolicyListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetUserPolicyListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUserPolicyList").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/getUserPolicyList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/getUserPolicyList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetUserPolicyListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// RenewNat 续订NAT
func (c *Client) RenewNat(request *model.RenewNatRequest) (*model.RenewNatResponse, error) {
    return c.RenewNatWithConfig(request, nil)
}

// RenewNatWithConfig 续订NAT
func (c *Client) RenewNatWithConfig(request *model.RenewNatRequest, runtimeConfig *config.RuntimeConfig) (*model.RenewNatResponse, error) {
    params := param.NewParamsBuilder().
        Action("renewNat").
        Uri("/ccabusiorder/api/v1/order/nat/renewNat").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/renewNat").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RenewNatResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateSecurityGroup 创建安全组
func (c *Client) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
    return c.CreateSecurityGroupWithConfig(request, nil)
}

// CreateSecurityGroupWithConfig 创建安全组
func (c *Client) CreateSecurityGroupWithConfig(request *model.CreateSecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateSecurityGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("createSecurityGroup").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/createSecurityGroup").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/createSecurityGroup").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateSecurityGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ImageRecordSharedByOthers 查询镜像的共享列表(别人共享给自己的)
func (c *Client) ImageRecordSharedByOthers(request *model.ImageRecordSharedByOthersRequest) (*model.ImageRecordSharedByOthersResponse, error) {
    return c.ImageRecordSharedByOthersWithConfig(request, nil)
}

// ImageRecordSharedByOthersWithConfig 查询镜像的共享列表(别人共享给自己的)
func (c *Client) ImageRecordSharedByOthersWithConfig(request *model.ImageRecordSharedByOthersRequest, runtimeConfig *config.RuntimeConfig) (*model.ImageRecordSharedByOthersResponse, error) {
    params := param.NewParamsBuilder().
        Action("ImageRecordSharedByOthers").
        Uri("/cem_openapi/ccacem/api/v1/image/imageRecordSharedByOthers").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/imageRecordSharedByOthers").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ImageRecordSharedByOthersResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetFirstLoginModifyPassword 首次登录强制修改密码配置
func (c *Client) SetFirstLoginModifyPassword(request *model.SetFirstLoginModifyPasswordRequest) (*model.SetFirstLoginModifyPasswordResponse, error) {
    return c.SetFirstLoginModifyPasswordWithConfig(request, nil)
}

// SetFirstLoginModifyPasswordWithConfig 首次登录强制修改密码配置
func (c *Client) SetFirstLoginModifyPasswordWithConfig(request *model.SetFirstLoginModifyPasswordRequest, runtimeConfig *config.RuntimeConfig) (*model.SetFirstLoginModifyPasswordResponse, error) {
    params := param.NewParamsBuilder().
        Action("setFirstLoginModifyPassword").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setFirstLoginModifyPassword").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setFirstLoginModifyPassword").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetFirstLoginModifyPasswordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ProductsNat 查询网关可用区及规格
func (c *Client) ProductsNat() (*model.ProductsNatResponse, error) {
    return c.ProductsNatWithConfig(nil)
}

// ProductsNatWithConfig 查询网关可用区及规格
func (c *Client) ProductsNatWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ProductsNatResponse, error) {
    params := param.NewParamsBuilder().
        Action("productsNat").
        Uri("/ccabusiorder/api/v1/order/nat/productsNat").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/productsNat").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.ProductsNatResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetAvailableRecordList 查询定时开机任务执行记录
func (c *Client) GetAvailableRecordList(request *model.GetAvailableRecordListRequest) (*model.GetAvailableRecordListResponse, error) {
    return c.GetAvailableRecordListWithConfig(request, nil)
}

// GetAvailableRecordListWithConfig 查询定时开机任务执行记录
func (c *Client) GetAvailableRecordListWithConfig(request *model.GetAvailableRecordListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetAvailableRecordListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getAvailableRecordList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getAvailableRecordList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getAvailableRecordList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetAvailableRecordListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// LoginLivingPageMonthly 查询月登录活跃度统计
func (c *Client) LoginLivingPageMonthly(request *model.LoginLivingPageMonthlyRequest) (*model.LoginLivingPageMonthlyResponse, error) {
    return c.LoginLivingPageMonthlyWithConfig(request, nil)
}

// LoginLivingPageMonthlyWithConfig 查询月登录活跃度统计
func (c *Client) LoginLivingPageMonthlyWithConfig(request *model.LoginLivingPageMonthlyRequest, runtimeConfig *config.RuntimeConfig) (*model.LoginLivingPageMonthlyResponse, error) {
    params := param.NewParamsBuilder().
        Action("loginLivingPageMonthly").
        Uri("/cem_openapi/ccacem/api/v1/operation/loginStatistics/monthly").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/loginStatistics/monthly").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LoginLivingPageMonthlyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetRestartRecordList 查询定时重启任务执行记录
func (c *Client) GetRestartRecordList(request *model.GetRestartRecordListRequest) (*model.GetRestartRecordListResponse, error) {
    return c.GetRestartRecordListWithConfig(request, nil)
}

// GetRestartRecordListWithConfig 查询定时重启任务执行记录
func (c *Client) GetRestartRecordListWithConfig(request *model.GetRestartRecordListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetRestartRecordListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getRestartRecordList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getRestartRecordList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getRestartRecordList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetRestartRecordListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchBindPolicy 定时任务关联电脑
func (c *Client) BatchBindPolicy(request *model.BatchBindPolicyRequest) (*model.BatchBindPolicyResponse, error) {
    return c.BatchBindPolicyWithConfig(request, nil)
}

// BatchBindPolicyWithConfig 定时任务关联电脑
func (c *Client) BatchBindPolicyWithConfig(request *model.BatchBindPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchBindPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchBindPolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/batchBindPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/batchBindPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchBindPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetCloudCompanyTag 查询企业ID
func (c *Client) GetCloudCompanyTag() (*model.GetCloudCompanyTagResponse, error) {
    return c.GetCloudCompanyTagWithConfig(nil)
}

// GetCloudCompanyTagWithConfig 查询企业ID
func (c *Client) GetCloudCompanyTagWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetCloudCompanyTagResponse, error) {
    params := param.NewParamsBuilder().
        Action("getCloudCompanyTag").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getCloudCompanyTag").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getCloudCompanyTag").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.GetCloudCompanyTagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CancelOrder 取消一体化服务包订单
func (c *Client) CancelOrder(request *model.CancelOrderRequest) (*model.CancelOrderResponse, error) {
    return c.CancelOrderWithConfig(request, nil)
}

// CancelOrderWithConfig 取消一体化服务包订单
func (c *Client) CancelOrderWithConfig(request *model.CancelOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.CancelOrderResponse, error) {
    params := param.NewParamsBuilder().
        Action("cancelOrder").
        Uri("/ccabusiorder/api/v1/channel/unify/cancelOrder").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/channel/unify/cancelOrder").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CancelOrderResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BindMachine 电脑策略关联电脑
func (c *Client) BindMachine(request *model.BindMachineRequest) (*model.BindMachineResponse, error) {
    return c.BindMachineWithConfig(request, nil)
}

// BindMachineWithConfig 电脑策略关联电脑
func (c *Client) BindMachineWithConfig(request *model.BindMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.BindMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindMachine").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/bindMachine").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/bindMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BindMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetUserModifyPassword 用户密码修改设置
func (c *Client) SetUserModifyPassword(request *model.SetUserModifyPasswordRequest) (*model.SetUserModifyPasswordResponse, error) {
    return c.SetUserModifyPasswordWithConfig(request, nil)
}

// SetUserModifyPasswordWithConfig 用户密码修改设置
func (c *Client) SetUserModifyPasswordWithConfig(request *model.SetUserModifyPasswordRequest, runtimeConfig *config.RuntimeConfig) (*model.SetUserModifyPasswordResponse, error) {
    params := param.NewParamsBuilder().
        Action("setUserModifyPassword").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setUserModifyPassword").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setUserModifyPassword").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetUserModifyPasswordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetDeviceTrustNum 设备可信认证数量查询
func (c *Client) GetDeviceTrustNum() (*model.GetDeviceTrustNumResponse, error) {
    return c.GetDeviceTrustNumWithConfig(nil)
}

// GetDeviceTrustNumWithConfig 设备可信认证数量查询
func (c *Client) GetDeviceTrustNumWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetDeviceTrustNumResponse, error) {
    params := param.NewParamsBuilder().
        Action("getDeviceTrustNum").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getDeviceTrustNum").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getDeviceTrustNum").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.GetDeviceTrustNumResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ConnectLivingPageMonthly 查询月连接活跃度统计
func (c *Client) ConnectLivingPageMonthly(request *model.ConnectLivingPageMonthlyRequest) (*model.ConnectLivingPageMonthlyResponse, error) {
    return c.ConnectLivingPageMonthlyWithConfig(request, nil)
}

// ConnectLivingPageMonthlyWithConfig 查询月连接活跃度统计
func (c *Client) ConnectLivingPageMonthlyWithConfig(request *model.ConnectLivingPageMonthlyRequest, runtimeConfig *config.RuntimeConfig) (*model.ConnectLivingPageMonthlyResponse, error) {
    params := param.NewParamsBuilder().
        Action("connectLivingPageMonthly").
        Uri("/cem_openapi/ccacem/api/v1/operation/connectStatistics/monthly").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/connectStatistics/monthly").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ConnectLivingPageMonthlyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchUserReActiveNotify 批量重发激活通知
func (c *Client) BatchUserReActiveNotify(request *model.BatchUserReActiveNotifyRequest) (*model.BatchUserReActiveNotifyResponse, error) {
    return c.BatchUserReActiveNotifyWithConfig(request, nil)
}

// BatchUserReActiveNotifyWithConfig 批量重发激活通知
func (c *Client) BatchUserReActiveNotifyWithConfig(request *model.BatchUserReActiveNotifyRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchUserReActiveNotifyResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchUserReActiveNotify").
        Uri("/cem_openapi/ccacem/api/v1/userManage/batchUserReActiveNotify").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/batchUserReActiveNotify").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchUserReActiveNotifyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StartAutoRenewNat NAT开启或关闭自动续订
func (c *Client) StartAutoRenewNat(request *model.StartAutoRenewNatRequest) (*model.StartAutoRenewNatResponse, error) {
    return c.StartAutoRenewNatWithConfig(request, nil)
}

// StartAutoRenewNatWithConfig NAT开启或关闭自动续订
func (c *Client) StartAutoRenewNatWithConfig(request *model.StartAutoRenewNatRequest, runtimeConfig *config.RuntimeConfig) (*model.StartAutoRenewNatResponse, error) {
    params := param.NewParamsBuilder().
        Action("startAutoRenewNat").
        Uri("/ccabusiorder/api/v1/order/nat/startOrStopAutoRenewNat").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/startOrStopAutoRenewNat").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.StartAutoRenewNatResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ConnectLivingPageDaily 查询日连接活跃度统计
func (c *Client) ConnectLivingPageDaily(request *model.ConnectLivingPageDailyRequest) (*model.ConnectLivingPageDailyResponse, error) {
    return c.ConnectLivingPageDailyWithConfig(request, nil)
}

// ConnectLivingPageDailyWithConfig 查询日连接活跃度统计
func (c *Client) ConnectLivingPageDailyWithConfig(request *model.ConnectLivingPageDailyRequest, runtimeConfig *config.RuntimeConfig) (*model.ConnectLivingPageDailyResponse, error) {
    params := param.NewParamsBuilder().
        Action("connectLivingPageDaily").
        Uri("/cem_openapi/ccacem/api/v1/operation/connectStatistics/daily").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/connectStatistics/daily").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ConnectLivingPageDailyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ApplySnapshot 还原快照
func (c *Client) ApplySnapshot(request *model.ApplySnapshotRequest) (*model.ApplySnapshotResponse, error) {
    return c.ApplySnapshotWithConfig(request, nil)
}

// ApplySnapshotWithConfig 还原快照
func (c *Client) ApplySnapshotWithConfig(request *model.ApplySnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.ApplySnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("applySnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/applySnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/applySnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ApplySnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPasswordUpdatePeriodMonths 密码强制更新时间查询
func (c *Client) GetPasswordUpdatePeriodMonths() (*model.GetPasswordUpdatePeriodMonthsResponse, error) {
    return c.GetPasswordUpdatePeriodMonthsWithConfig(nil)
}

// GetPasswordUpdatePeriodMonthsWithConfig 密码强制更新时间查询
func (c *Client) GetPasswordUpdatePeriodMonthsWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetPasswordUpdatePeriodMonthsResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPasswordUpdatePeriodMonths").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getPasswordUpdatePeriodMonths").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getPasswordUpdatePeriodMonths").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.GetPasswordUpdatePeriodMonthsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// LockMachine 单个锁定云电脑
func (c *Client) LockMachine(request *model.LockMachineRequest) (*model.LockMachineResponse, error) {
    return c.LockMachineWithConfig(request, nil)
}

// LockMachineWithConfig 单个锁定云电脑
func (c *Client) LockMachineWithConfig(request *model.LockMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.LockMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("lockMachine").
        Uri("/cem_openapi/ccacem/api/v1/machine/lockMachine").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/lockMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LockMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPolicyList 查看电脑策略列表
func (c *Client) GetPolicyList(request *model.GetPolicyListRequest) (*model.GetPolicyListResponse, error) {
    return c.GetPolicyListWithConfig(request, nil)
}

// GetPolicyListWithConfig 查看电脑策略列表
func (c *Client) GetPolicyListWithConfig(request *model.GetPolicyListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPolicyListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPolicyList").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyList").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/getPolicyList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetPolicyListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchBindDevice 批量添加限制登录硬终端
func (c *Client) BatchBindDevice(request *model.BatchBindDeviceRequest) (*model.BatchBindDeviceResponse, error) {
    return c.BatchBindDeviceWithConfig(request, nil)
}

// BatchBindDeviceWithConfig 批量添加限制登录硬终端
func (c *Client) BatchBindDeviceWithConfig(request *model.BatchBindDeviceRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchBindDeviceResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchBindDevice").
        Uri("/cem_openapi/ccacem/api/v1/device/batchBindDevice").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/batchBindDevice").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchBindDeviceResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ListImage 查询镜像列表
func (c *Client) ListImage(request *model.ListImageRequest) (*model.ListImageResponse, error) {
    return c.ListImageWithConfig(request, nil)
}

// ListImageWithConfig 查询镜像列表
func (c *Client) ListImageWithConfig(request *model.ListImageRequest, runtimeConfig *config.RuntimeConfig) (*model.ListImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("listImage").
        Uri("/cem_openapi/ccacem/api/v1/image/listImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/listImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ListImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPolicyDetailByUserId 查询用户已关联策略
func (c *Client) GetPolicyDetailByUserId(request *model.GetPolicyDetailByUserIdRequest) (*model.GetPolicyDetailByUserIdResponse, error) {
    return c.GetPolicyDetailByUserIdWithConfig(request, nil)
}

// GetPolicyDetailByUserIdWithConfig 查询用户已关联策略
func (c *Client) GetPolicyDetailByUserIdWithConfig(request *model.GetPolicyDetailByUserIdRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPolicyDetailByUserIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPolicyDetailByUserId").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/getPolicyDetailByUserId").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/getPolicyDetailByUserId").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetPolicyDetailByUserIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SubmitBw 订购带宽
func (c *Client) SubmitBw(request *model.SubmitBwRequest) (*model.SubmitBwResponse, error) {
    return c.SubmitBwWithConfig(request, nil)
}

// SubmitBwWithConfig 订购带宽
func (c *Client) SubmitBwWithConfig(request *model.SubmitBwRequest, runtimeConfig *config.RuntimeConfig) (*model.SubmitBwResponse, error) {
    params := param.NewParamsBuilder().
        Action("submitBw").
        Uri("/ccabusiorder/api/v1/order/bw/submitBw").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/submitBw").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SubmitBwResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetTaskRecordList 获取任务执行明细
func (c *Client) GetTaskRecordList(request *model.GetTaskRecordListRequest) (*model.GetTaskRecordListResponse, error) {
    return c.GetTaskRecordListWithConfig(request, nil)
}

// GetTaskRecordListWithConfig 获取任务执行明细
func (c *Client) GetTaskRecordListWithConfig(request *model.GetTaskRecordListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetTaskRecordListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getTaskRecordList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getTaskRecordList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getTaskRecordList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetTaskRecordListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetApplyImageRecord 查询镜像还原记录
func (c *Client) GetApplyImageRecord(request *model.GetApplyImageRecordRequest) (*model.GetApplyImageRecordResponse, error) {
    return c.GetApplyImageRecordWithConfig(request, nil)
}

// GetApplyImageRecordWithConfig 查询镜像还原记录
func (c *Client) GetApplyImageRecordWithConfig(request *model.GetApplyImageRecordRequest, runtimeConfig *config.RuntimeConfig) (*model.GetApplyImageRecordResponse, error) {
    params := param.NewParamsBuilder().
        Action("getApplyImageRecord").
        Uri("/cem_openapi/ccacem/api/v1/image/getApplyImageRecord").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/image/getApplyImageRecord").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetApplyImageRecordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetIPByUserName 根据账号查询IP信息
func (c *Client) GetIPByUserName(request *model.GetIPByUserNameRequest) (*model.GetIPByUserNameResponse, error) {
    return c.GetIPByUserNameWithConfig(request, nil)
}

// GetIPByUserNameWithConfig 根据账号查询IP信息
func (c *Client) GetIPByUserNameWithConfig(request *model.GetIPByUserNameRequest, runtimeConfig *config.RuntimeConfig) (*model.GetIPByUserNameResponse, error) {
    params := param.NewParamsBuilder().
        Action("getIPByUserName").
        Uri("/cem_openapi/ccacem/api/v1/userManage/getIPByUserName").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/getIPByUserName").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetIPByUserNameResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ContactInfo 企业联系人信息查询
func (c *Client) ContactInfo() (*model.ContactInfoResponse, error) {
    return c.ContactInfoWithConfig(nil)
}

// ContactInfoWithConfig 企业联系人信息查询
func (c *Client) ContactInfoWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ContactInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("contactInfo").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/contactInfo").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/contactInfo").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.ContactInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// MoveOrg 移动组织
func (c *Client) MoveOrg(request *model.MoveOrgRequest) (*model.MoveOrgResponse, error) {
    return c.MoveOrgWithConfig(request, nil)
}

// MoveOrgWithConfig 移动组织
func (c *Client) MoveOrgWithConfig(request *model.MoveOrgRequest, runtimeConfig *config.RuntimeConfig) (*model.MoveOrgResponse, error) {
    params := param.NewParamsBuilder().
        Action("moveOrg").
        Uri("/cem_openapi/ccacem/api/v1/userManage/moveOrg").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/moveOrg").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.MoveOrgResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetTwoFactorAuth 查询双因子认证登录配置
func (c *Client) GetTwoFactorAuth() (*model.GetTwoFactorAuthResponse, error) {
    return c.GetTwoFactorAuthWithConfig(nil)
}

// GetTwoFactorAuthWithConfig 查询双因子认证登录配置
func (c *Client) GetTwoFactorAuthWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetTwoFactorAuthResponse, error) {
    params := param.NewParamsBuilder().
        Action("getTwoFactorAuth").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getTwoFactorAuth").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getTwoFactorAuth").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetTwoFactorAuthResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeviceUnBindUser 解绑用户
func (c *Client) DeviceUnBindUser(request *model.DeviceUnBindUserRequest) (*model.DeviceUnBindUserResponse, error) {
    return c.DeviceUnBindUserWithConfig(request, nil)
}

// DeviceUnBindUserWithConfig 解绑用户
func (c *Client) DeviceUnBindUserWithConfig(request *model.DeviceUnBindUserRequest, runtimeConfig *config.RuntimeConfig) (*model.DeviceUnBindUserResponse, error) {
    params := param.NewParamsBuilder().
        Action("deviceUnBindUser").
        Uri("/cem_openapi/ccacem/api/v1/device/deviceUnBindUser").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/device/deviceUnBindUser").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeviceUnBindUserResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// LoginUsageStatisticsPage 查询登录使用统计
func (c *Client) LoginUsageStatisticsPage(request *model.LoginUsageStatisticsPageRequest) (*model.LoginUsageStatisticsPageResponse, error) {
    return c.LoginUsageStatisticsPageWithConfig(request, nil)
}

// LoginUsageStatisticsPageWithConfig 查询登录使用统计
func (c *Client) LoginUsageStatisticsPageWithConfig(request *model.LoginUsageStatisticsPageRequest, runtimeConfig *config.RuntimeConfig) (*model.LoginUsageStatisticsPageResponse, error) {
    params := param.NewParamsBuilder().
        Action("loginUsageStatisticsPage").
        Uri("/cem_openapi/ccacem/api/v1/operation/loginUsageStatistics/page").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/loginUsageStatistics/page").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LoginUsageStatisticsPageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetShowOrderDetails 客户端订单信息展示控制
func (c *Client) SetShowOrderDetails(request *model.SetShowOrderDetailsRequest) (*model.SetShowOrderDetailsResponse, error) {
    return c.SetShowOrderDetailsWithConfig(request, nil)
}

// SetShowOrderDetailsWithConfig 客户端订单信息展示控制
func (c *Client) SetShowOrderDetailsWithConfig(request *model.SetShowOrderDetailsRequest, runtimeConfig *config.RuntimeConfig) (*model.SetShowOrderDetailsResponse, error) {
    params := param.NewParamsBuilder().
        Action("setShowOrderDetails").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setShowOrderDetails").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setShowOrderDetails").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetShowOrderDetailsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateSNatRule 创建SNAT规则
func (c *Client) CreateSNatRule(request *model.CreateSNatRuleRequest) (*model.CreateSNatRuleResponse, error) {
    return c.CreateSNatRuleWithConfig(request, nil)
}

// CreateSNatRuleWithConfig 创建SNAT规则
func (c *Client) CreateSNatRuleWithConfig(request *model.CreateSNatRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateSNatRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("createSNatRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/createSnatRule").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/createSnatRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateSNatRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetUnbindMobileFlag 用户解绑手机号配置
func (c *Client) SetUnbindMobileFlag(request *model.SetUnbindMobileFlagRequest) (*model.SetUnbindMobileFlagResponse, error) {
    return c.SetUnbindMobileFlagWithConfig(request, nil)
}

// SetUnbindMobileFlagWithConfig 用户解绑手机号配置
func (c *Client) SetUnbindMobileFlagWithConfig(request *model.SetUnbindMobileFlagRequest, runtimeConfig *config.RuntimeConfig) (*model.SetUnbindMobileFlagResponse, error) {
    params := param.NewParamsBuilder().
        Action("setUnbindMobileFlag").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setUnbindMobileFlag").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setUnbindMobileFlag").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetUnbindMobileFlagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddMacAddressBatch 批量新增云电脑mac地址白名单
func (c *Client) AddMacAddressBatch(request *model.AddMacAddressBatchRequest) (*model.AddMacAddressBatchResponse, error) {
    return c.AddMacAddressBatchWithConfig(request, nil)
}

// AddMacAddressBatchWithConfig 批量新增云电脑mac地址白名单
func (c *Client) AddMacAddressBatchWithConfig(request *model.AddMacAddressBatchRequest, runtimeConfig *config.RuntimeConfig) (*model.AddMacAddressBatchResponse, error) {
    params := param.NewParamsBuilder().
        Action("addMacAddressBatch").
        Uri("/cem_openapi/ccacem/api/v1/machine/addMacAddressBatch").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/addMacAddressBatch").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddMacAddressBatchResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchShutdown 云电脑批量关机接口
func (c *Client) BatchShutdown(request *model.BatchShutdownRequest) (*model.BatchShutdownResponse, error) {
    return c.BatchShutdownWithConfig(request, nil)
}

// BatchShutdownWithConfig 云电脑批量关机接口
func (c *Client) BatchShutdownWithConfig(request *model.BatchShutdownRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchShutdownResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchShutdown").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchShutdown").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchShutdown").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchShutdownResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// HistoryConnectPage 查询历史连接会话
func (c *Client) HistoryConnectPage(request *model.HistoryConnectPageRequest) (*model.HistoryConnectPageResponse, error) {
    return c.HistoryConnectPageWithConfig(request, nil)
}

// HistoryConnectPageWithConfig 查询历史连接会话
func (c *Client) HistoryConnectPageWithConfig(request *model.HistoryConnectPageRequest, runtimeConfig *config.RuntimeConfig) (*model.HistoryConnectPageResponse, error) {
    params := param.NewParamsBuilder().
        Action("historyConnectPage").
        Uri("/cem_openapi/ccacem/api/v1/operation/connectSession/historyPage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/connectSession/historyPage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.HistoryConnectPageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetIsBindUserInfoList 查询硬终端绑定用户列表信息
func (c *Client) GetIsBindUserInfoList(request *model.GetIsBindUserInfoListRequest) (*model.GetIsBindUserInfoListResponse, error) {
    return c.GetIsBindUserInfoListWithConfig(request, nil)
}

// GetIsBindUserInfoListWithConfig 查询硬终端绑定用户列表信息
func (c *Client) GetIsBindUserInfoListWithConfig(request *model.GetIsBindUserInfoListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetIsBindUserInfoListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getIsBindUserInfoList").
        Uri("/cem_openapi/ccacem/api/v1/device/getIsBindUserInfoList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/getIsBindUserInfoList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetIsBindUserInfoListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchDeleteUser 删除用户
func (c *Client) BatchDeleteUser(request *model.BatchDeleteUserRequest) (*model.BatchDeleteUserResponse, error) {
    return c.BatchDeleteUserWithConfig(request, nil)
}

// BatchDeleteUserWithConfig 删除用户
func (c *Client) BatchDeleteUserWithConfig(request *model.BatchDeleteUserRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchDeleteUserResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchDeleteUser").
        Uri("/cem_openapi/ccacem/api/v1/userManage/batchDeleteUser").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/batchDeleteUser").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchDeleteUserResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateShutdownRevertConfig 修改电脑关机还原配置
func (c *Client) UpdateShutdownRevertConfig(request *model.UpdateShutdownRevertConfigRequest) (*model.UpdateShutdownRevertConfigResponse, error) {
    return c.UpdateShutdownRevertConfigWithConfig(request, nil)
}

// UpdateShutdownRevertConfigWithConfig 修改电脑关机还原配置
func (c *Client) UpdateShutdownRevertConfigWithConfig(request *model.UpdateShutdownRevertConfigRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateShutdownRevertConfigResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateShutdownRevertConfig").
        Uri("/cem_openapi/ccacem/api/v1/machine/updateShutdownRevertConfig").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/updateShutdownRevertConfig").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateShutdownRevertConfigResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Regions 查询一体化服务包下单省市区信息
func (c *Client) Regions(request *model.RegionsRequest) (*model.RegionsResponse, error) {
    return c.RegionsWithConfig(request, nil)
}

// RegionsWithConfig 查询一体化服务包下单省市区信息
func (c *Client) RegionsWithConfig(request *model.RegionsRequest, runtimeConfig *config.RuntimeConfig) (*model.RegionsResponse, error) {
    params := param.NewParamsBuilder().
        Action("regions").
        Uri("/ccabusiorder/api/v1/order/common/regions").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/common/regions").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.RegionsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddPolicy 创建定时任务
func (c *Client) AddPolicy(request *model.AddPolicyRequest) (*model.AddPolicyResponse, error) {
    return c.AddPolicyWithConfig(request, nil)
}

// AddPolicyWithConfig 创建定时任务
func (c *Client) AddPolicyWithConfig(request *model.AddPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.AddPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("addPolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/addPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/addPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetApplySnapshotRecord 查询快照还原记录
func (c *Client) GetApplySnapshotRecord(request *model.GetApplySnapshotRecordRequest) (*model.GetApplySnapshotRecordResponse, error) {
    return c.GetApplySnapshotRecordWithConfig(request, nil)
}

// GetApplySnapshotRecordWithConfig 查询快照还原记录
func (c *Client) GetApplySnapshotRecordWithConfig(request *model.GetApplySnapshotRecordRequest, runtimeConfig *config.RuntimeConfig) (*model.GetApplySnapshotRecordResponse, error) {
    params := param.NewParamsBuilder().
        Action("getApplySnapshotRecord").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/getApplySnapshotRecord").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/getApplySnapshotRecord").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetApplySnapshotRecordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetBindDeviceList 查看已限制登录硬终端
func (c *Client) GetBindDeviceList(request *model.GetBindDeviceListRequest) (*model.GetBindDeviceListResponse, error) {
    return c.GetBindDeviceListWithConfig(request, nil)
}

// GetBindDeviceListWithConfig 查看已限制登录硬终端
func (c *Client) GetBindDeviceListWithConfig(request *model.GetBindDeviceListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetBindDeviceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getBindDeviceList").
        Uri("/cem_openapi/ccacem/api/v1/device/ccemp/getBindDeviceList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/ccemp/getBindDeviceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetBindDeviceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateSnapshot 创建快照
func (c *Client) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
    return c.CreateSnapshotWithConfig(request, nil)
}

// CreateSnapshotWithConfig 创建快照
func (c *Client) CreateSnapshotWithConfig(request *model.CreateSnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateSnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("createSnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/createSnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/createSnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateSnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetShowOrderDetails 客户端订单信息展示开关查询
func (c *Client) GetShowOrderDetails() (*model.GetShowOrderDetailsResponse, error) {
    return c.GetShowOrderDetailsWithConfig(nil)
}

// GetShowOrderDetailsWithConfig 客户端订单信息展示开关查询
func (c *Client) GetShowOrderDetailsWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetShowOrderDetailsResponse, error) {
    params := param.NewParamsBuilder().
        Action("getShowOrderDetails").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getShowOrderDetails").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getShowOrderDetails").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetShowOrderDetailsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUserModifyPassword 用户密码修改开关查询
func (c *Client) GetUserModifyPassword() (*model.GetUserModifyPasswordResponse, error) {
    return c.GetUserModifyPasswordWithConfig(nil)
}

// GetUserModifyPasswordWithConfig 用户密码修改开关查询
func (c *Client) GetUserModifyPasswordWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetUserModifyPasswordResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUserModifyPassword").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/getUserModifyPassword").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/getUserModifyPassword").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetUserModifyPasswordResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ImageRecordSharedWithOthers 查询镜像的共享记录(自己共享给别人的)
func (c *Client) ImageRecordSharedWithOthers(request *model.ImageRecordSharedWithOthersRequest) (*model.ImageRecordSharedWithOthersResponse, error) {
    return c.ImageRecordSharedWithOthersWithConfig(request, nil)
}

// ImageRecordSharedWithOthersWithConfig 查询镜像的共享记录(自己共享给别人的)
func (c *Client) ImageRecordSharedWithOthersWithConfig(request *model.ImageRecordSharedWithOthersRequest, runtimeConfig *config.RuntimeConfig) (*model.ImageRecordSharedWithOthersResponse, error) {
    params := param.NewParamsBuilder().
        Action("ImageRecordSharedWithOthers").
        Uri("/cem_openapi/ccacem/api/v1/image/imageRecordSharedWithOthers").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/imageRecordSharedWithOthers").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ImageRecordSharedWithOthersResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetVmList 获取桌面列表接口
func (c *Client) GetVmList(request *model.GetVmListRequest) (*model.GetVmListResponse, error) {
    return c.GetVmListWithConfig(request, nil)
}

// GetVmListWithConfig 获取桌面列表接口
func (c *Client) GetVmListWithConfig(request *model.GetVmListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetVmListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getVmList").
        Uri("/cem_openapi/ccacem/api/v1/resource/getVmList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource/getVmList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetVmListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetMultiDiskInfo 查询多数据盘信息
func (c *Client) GetMultiDiskInfo(request *model.GetMultiDiskInfoRequest) (*model.GetMultiDiskInfoResponse, error) {
    return c.GetMultiDiskInfoWithConfig(request, nil)
}

// GetMultiDiskInfoWithConfig 查询多数据盘信息
func (c *Client) GetMultiDiskInfoWithConfig(request *model.GetMultiDiskInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.GetMultiDiskInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("getMultiDiskInfo").
        Uri("/ccabusiorder/api/v1/order/getMultiDiskInfo").
        GatewayUri("/api/clouddesktopnew/ccabusiorder/api/v1/order/getMultiDiskInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetMultiDiskInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchReloadByImage 批量切换镜像
func (c *Client) BatchReloadByImage(request *model.BatchReloadByImageRequest) (*model.BatchReloadByImageResponse, error) {
    return c.BatchReloadByImageWithConfig(request, nil)
}

// BatchReloadByImageWithConfig 批量切换镜像
func (c *Client) BatchReloadByImageWithConfig(request *model.BatchReloadByImageRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchReloadByImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchReloadByImage").
        Uri("/cem_openapi/ccacem/api/v1/image/batchReloadByImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/batchReloadByImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchReloadByImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// MultiDiskChange 新增数据盘
func (c *Client) MultiDiskChange(request *model.MultiDiskChangeRequest) (*model.MultiDiskChangeResponse, error) {
    return c.MultiDiskChangeWithConfig(request, nil)
}

// MultiDiskChangeWithConfig 新增数据盘
func (c *Client) MultiDiskChangeWithConfig(request *model.MultiDiskChangeRequest, runtimeConfig *config.RuntimeConfig) (*model.MultiDiskChangeResponse, error) {
    params := param.NewParamsBuilder().
        Action("multiDiskChange").
        Uri("/ccabusiorder/api/v1/order/multiDiskAdd").
        GatewayUri("/api/Computer_api/ccabusiorder/api/v1/order/multiDiskAdd").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.MultiDiskChangeResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUserDetailByPolicyId 用户策略关联用户状态列表查询
func (c *Client) GetUserDetailByPolicyId(request *model.GetUserDetailByPolicyIdRequest) (*model.GetUserDetailByPolicyIdResponse, error) {
    return c.GetUserDetailByPolicyIdWithConfig(request, nil)
}

// GetUserDetailByPolicyIdWithConfig 用户策略关联用户状态列表查询
func (c *Client) GetUserDetailByPolicyIdWithConfig(request *model.GetUserDetailByPolicyIdRequest, runtimeConfig *config.RuntimeConfig) (*model.GetUserDetailByPolicyIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUserDetailByPolicyId").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/getUserDetailByPolicyId").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/getUserDetailByPolicyId").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetUserDetailByPolicyIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Add 新增消息
func (c *Client) Add(request *model.AddRequest) (*model.AddResponse, error) {
    return c.AddWithConfig(request, nil)
}

// AddWithConfig 新增消息
func (c *Client) AddWithConfig(request *model.AddRequest, runtimeConfig *config.RuntimeConfig) (*model.AddResponse, error) {
    params := param.NewParamsBuilder().
        Action("add").
        Uri("/cem_openapi/ccacem/api/v1//message/add").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/message/add").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetClientInfo 登录认证接口
func (c *Client) GetClientInfo(request *model.GetClientInfoRequest) (*model.GetClientInfoResponse, error) {
    return c.GetClientInfoWithConfig(request, nil)
}

// GetClientInfoWithConfig 登录认证接口
func (c *Client) GetClientInfoWithConfig(request *model.GetClientInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.GetClientInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("getClientInfo").
        Uri("/cem_openapi/ccacem/api/v1/auth/login").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/auth/login").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetClientInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// LockDeviceById 锁定终端
func (c *Client) LockDeviceById(request *model.LockDeviceByIdRequest) (*model.LockDeviceByIdResponse, error) {
    return c.LockDeviceByIdWithConfig(request, nil)
}

// LockDeviceByIdWithConfig 锁定终端
func (c *Client) LockDeviceByIdWithConfig(request *model.LockDeviceByIdRequest, runtimeConfig *config.RuntimeConfig) (*model.LockDeviceByIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("lockDeviceById").
        Uri("/cem_openapi/ccacem/api/v1/device/lockDeviceById").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/lockDeviceById").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LockDeviceByIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ConnectUsageStatisticsPage 查询连接使用统计
func (c *Client) ConnectUsageStatisticsPage(request *model.ConnectUsageStatisticsPageRequest) (*model.ConnectUsageStatisticsPageResponse, error) {
    return c.ConnectUsageStatisticsPageWithConfig(request, nil)
}

// ConnectUsageStatisticsPageWithConfig 查询连接使用统计
func (c *Client) ConnectUsageStatisticsPageWithConfig(request *model.ConnectUsageStatisticsPageRequest, runtimeConfig *config.RuntimeConfig) (*model.ConnectUsageStatisticsPageResponse, error) {
    params := param.NewParamsBuilder().
        Action("connectUsageStatisticsPage").
        Uri("/cem_openapi/ccacem/api/v1/operation/connectUsageStatistics/page").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/connectUsageStatistics/page").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ConnectUsageStatisticsPageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateRule 编辑规则
func (c *Client) UpdateRule(request *model.UpdateRuleRequest) (*model.UpdateRuleResponse, error) {
    return c.UpdateRuleWithConfig(request, nil)
}

// UpdateRuleWithConfig 编辑规则
func (c *Client) UpdateRuleWithConfig(request *model.UpdateRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/updateRule").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/updateRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BindInstanceList 查询定时任务已关联电脑
func (c *Client) BindInstanceList(request *model.BindInstanceListRequest) (*model.BindInstanceListResponse, error) {
    return c.BindInstanceListWithConfig(request, nil)
}

// BindInstanceListWithConfig 查询定时任务已关联电脑
func (c *Client) BindInstanceListWithConfig(request *model.BindInstanceListRequest, runtimeConfig *config.RuntimeConfig) (*model.BindInstanceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindInstanceList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/bindInstanceList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/bindInstanceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BindInstanceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Detail 查看消息详情
func (c *Client) Detail(request *model.DetailRequest) (*model.DetailResponse, error) {
    return c.DetailWithConfig(request, nil)
}

// DetailWithConfig 查看消息详情
func (c *Client) DetailWithConfig(request *model.DetailRequest, runtimeConfig *config.RuntimeConfig) (*model.DetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("detail").
        Uri("/cem_openapi/ccacem/api/v1//message/detail").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/detail").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Logout 登出接口
func (c *Client) Logout(request *model.LogoutRequest) (*model.LogoutResponse, error) {
    return c.LogoutWithConfig(request, nil)
}

// LogoutWithConfig 登出接口
func (c *Client) LogoutWithConfig(request *model.LogoutRequest, runtimeConfig *config.RuntimeConfig) (*model.LogoutResponse, error) {
    params := param.NewParamsBuilder().
        Action("logout").
        Uri("/cem_openapi/ccacem/api/v1/auth/logout").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/auth/logout").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LogoutResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteNetwork 删除子网
func (c *Client) DeleteNetwork(request *model.DeleteNetworkRequest) (*model.DeleteNetworkResponse, error) {
    return c.DeleteNetworkWithConfig(request, nil)
}

// DeleteNetworkWithConfig 删除子网
func (c *Client) DeleteNetworkWithConfig(request *model.DeleteNetworkRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteNetworkResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteNetwork").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/deleteNetwork").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/deleteNetwork").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteNetworkResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetResourceDiskInfo 获取电脑数据盘信息
func (c *Client) GetResourceDiskInfo(request *model.GetResourceDiskInfoRequest) (*model.GetResourceDiskInfoResponse, error) {
    return c.GetResourceDiskInfoWithConfig(request, nil)
}

// GetResourceDiskInfoWithConfig 获取电脑数据盘信息
func (c *Client) GetResourceDiskInfoWithConfig(request *model.GetResourceDiskInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.GetResourceDiskInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("getResourceDiskInfo").
        Uri("/cem_openapi/ccacem/api/v1/machine/getResourceDiskInfo").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/getResourceDiskInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetResourceDiskInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SdkCipher sdk认证信息获取接口
func (c *Client) SdkCipher() (*model.SdkCipherResponse, error) {
    return c.SdkCipherWithConfig(nil)
}

// SdkCipherWithConfig sdk认证信息获取接口
func (c *Client) SdkCipherWithConfig(runtimeConfig *config.RuntimeConfig) (*model.SdkCipherResponse, error) {
    params := param.NewParamsBuilder().
        Action("sdkCipher").
        Uri("/cem_openapi/ccacem/api/v1/auth/sdkCipher").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/auth/sdkCipher").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Build()
    returnValue := &model.SdkCipherResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BindResourceRecords 根据资源实例获取分配记录
func (c *Client) BindResourceRecords(request *model.BindResourceRecordsRequest) (*model.BindResourceRecordsResponse, error) {
    return c.BindResourceRecordsWithConfig(request, nil)
}

// BindResourceRecordsWithConfig 根据资源实例获取分配记录
func (c *Client) BindResourceRecordsWithConfig(request *model.BindResourceRecordsRequest, runtimeConfig *config.RuntimeConfig) (*model.BindResourceRecordsResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindResourceRecords").
        Uri("/cem_openapi/ccacem/api/v1/machine/bindResourceRecords").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/bindResourceRecords").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BindResourceRecordsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ChangeVpc 切换VPC
func (c *Client) ChangeVpc(request *model.ChangeVpcRequest) (*model.ChangeVpcResponse, error) {
    return c.ChangeVpcWithConfig(request, nil)
}

// ChangeVpcWithConfig 切换VPC
func (c *Client) ChangeVpcWithConfig(request *model.ChangeVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.ChangeVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("changeVpc").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/changeVpc").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/changeVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ChangeVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ProductsGov 查询政企版规格可用区
func (c *Client) ProductsGov() (*model.ProductsGovResponse, error) {
    return c.ProductsGovWithConfig(nil)
}

// ProductsGovWithConfig 查询政企版规格可用区
func (c *Client) ProductsGovWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ProductsGovResponse, error) {
    params := param.NewParamsBuilder().
        Action("productsGov").
        Uri("/ccabusiorder/api/v1/order/gov/products").
        GatewayUri("/api/clouddesktopnew/ccabusiorder/api/v1/order/gov/products").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.ProductsGovResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchReload 云电脑批量重装接口
func (c *Client) BatchReload(request *model.BatchReloadRequest) (*model.BatchReloadResponse, error) {
    return c.BatchReloadWithConfig(request, nil)
}

// BatchReloadWithConfig 云电脑批量重装接口
func (c *Client) BatchReloadWithConfig(request *model.BatchReloadRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchReloadResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchReload").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchReload").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchReload").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchReloadResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// RenameResourceBatch 批量修改云电脑名称
func (c *Client) RenameResourceBatch(request *model.RenameResourceBatchRequest) (*model.RenameResourceBatchResponse, error) {
    return c.RenameResourceBatchWithConfig(request, nil)
}

// RenameResourceBatchWithConfig 批量修改云电脑名称
func (c *Client) RenameResourceBatchWithConfig(request *model.RenameResourceBatchRequest, runtimeConfig *config.RuntimeConfig) (*model.RenameResourceBatchResponse, error) {
    params := param.NewParamsBuilder().
        Action("renameResourceBatch").
        Uri("/cem_openapi/ccacem/api/v1/machine/renameResourceBatch").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/renameResourceBatch").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RenameResourceBatchResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnBindMachine 从电脑策略移除已关联电脑
func (c *Client) UnBindMachine(request *model.UnBindMachineRequest) (*model.UnBindMachineResponse, error) {
    return c.UnBindMachineWithConfig(request, nil)
}

// UnBindMachineWithConfig 从电脑策略移除已关联电脑
func (c *Client) UnBindMachineWithConfig(request *model.UnBindMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.UnBindMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("unBindMachine").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/unBindMachine").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/unBindMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnBindMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CopySecurityGroup 复制安全组
func (c *Client) CopySecurityGroup(request *model.CopySecurityGroupRequest) (*model.CopySecurityGroupResponse, error) {
    return c.CopySecurityGroupWithConfig(request, nil)
}

// CopySecurityGroupWithConfig 复制安全组
func (c *Client) CopySecurityGroupWithConfig(request *model.CopySecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.CopySecurityGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("copySecurityGroup").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/copySecurityGroup").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/copySecurityGroup").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CopySecurityGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// NatChange NAT变更
func (c *Client) NatChange(request *model.NatChangeRequest) (*model.NatChangeResponse, error) {
    return c.NatChangeWithConfig(request, nil)
}

// NatChangeWithConfig NAT变更
func (c *Client) NatChangeWithConfig(request *model.NatChangeRequest, runtimeConfig *config.RuntimeConfig) (*model.NatChangeResponse, error) {
    params := param.NewParamsBuilder().
        Action("natChange").
        Uri("/ccabusiorder/api/v1/order/nat/natChange").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/natChange").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.NatChangeResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetSecurityGroupRuleInfo 获取安全组对应的规则信息
func (c *Client) GetSecurityGroupRuleInfo(request *model.GetSecurityGroupRuleInfoRequest) (*model.GetSecurityGroupRuleInfoResponse, error) {
    return c.GetSecurityGroupRuleInfoWithConfig(request, nil)
}

// GetSecurityGroupRuleInfoWithConfig 获取安全组对应的规则信息
func (c *Client) GetSecurityGroupRuleInfoWithConfig(request *model.GetSecurityGroupRuleInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.GetSecurityGroupRuleInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("getSecurityGroupRuleInfo").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/getSecurityGroupRuleInfo").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/getSecurityGroupRuleInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetSecurityGroupRuleInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// List 查看消息列表
func (c *Client) List(request *model.ListRequest) (*model.ListResponse, error) {
    return c.ListWithConfig(request, nil)
}

// ListWithConfig 查看消息列表
func (c *Client) ListWithConfig(request *model.ListRequest, runtimeConfig *config.RuntimeConfig) (*model.ListResponse, error) {
    params := param.NewParamsBuilder().
        Action("list").
        Uri("/cem_openapi/ccacem/api/v1//message/list").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/list").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteComputerPolicy 删除电脑策略
func (c *Client) DeleteComputerPolicy(request *model.DeleteComputerPolicyRequest) (*model.DeleteComputerPolicyResponse, error) {
    return c.DeleteComputerPolicyWithConfig(request, nil)
}

// DeleteComputerPolicyWithConfig 删除电脑策略
func (c *Client) DeleteComputerPolicyWithConfig(request *model.DeleteComputerPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteComputerPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteComputerPolicy").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/deletePolicy").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/deletePolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteComputerPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetTwoFactorAuth 配置双因子认证登录
func (c *Client) SetTwoFactorAuth(request *model.SetTwoFactorAuthRequest) (*model.SetTwoFactorAuthResponse, error) {
    return c.SetTwoFactorAuthWithConfig(request, nil)
}

// SetTwoFactorAuthWithConfig 配置双因子认证登录
func (c *Client) SetTwoFactorAuthWithConfig(request *model.SetTwoFactorAuthRequest, runtimeConfig *config.RuntimeConfig) (*model.SetTwoFactorAuthResponse, error) {
    params := param.NewParamsBuilder().
        Action("setTwoFactorAuth").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/setTwoFactorAuth").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/setTwoFactorAuth").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SetTwoFactorAuthResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnlockMachine 单个解锁云电脑
func (c *Client) UnlockMachine(request *model.UnlockMachineRequest) (*model.UnlockMachineResponse, error) {
    return c.UnlockMachineWithConfig(request, nil)
}

// UnlockMachineWithConfig 单个解锁云电脑
func (c *Client) UnlockMachineWithConfig(request *model.UnlockMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.UnlockMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("unlockMachine").
        Uri("/cem_openapi/ccacem/api/v1/machine/unlockMachine").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/unlockMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnlockMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetBwInstanceList 查询带宽列表
func (c *Client) GetBwInstanceList(request *model.GetBwInstanceListRequest) (*model.GetBwInstanceListResponse, error) {
    return c.GetBwInstanceListWithConfig(request, nil)
}

// GetBwInstanceListWithConfig 查询带宽列表
func (c *Client) GetBwInstanceListWithConfig(request *model.GetBwInstanceListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetBwInstanceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getBwInstanceList").
        Uri("/ccabusiorder/api/v1/order/bw/getBwInstanceList").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/getBwInstanceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetBwInstanceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ActiveConnectPage 查询当前连接会话
func (c *Client) ActiveConnectPage(request *model.ActiveConnectPageRequest) (*model.ActiveConnectPageResponse, error) {
    return c.ActiveConnectPageWithConfig(request, nil)
}

// ActiveConnectPageWithConfig 查询当前连接会话
func (c *Client) ActiveConnectPageWithConfig(request *model.ActiveConnectPageRequest, runtimeConfig *config.RuntimeConfig) (*model.ActiveConnectPageResponse, error) {
    params := param.NewParamsBuilder().
        Action("activeConnectPage").
        Uri("/cem_openapi/ccacem/api/v1/operation/connectSession/activePage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/connectSession/activePage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ActiveConnectPageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddVpc 创建vpc
func (c *Client) AddVpc(request *model.AddVpcRequest) (*model.AddVpcResponse, error) {
    return c.AddVpcWithConfig(request, nil)
}

// AddVpcWithConfig 创建vpc
func (c *Client) AddVpcWithConfig(request *model.AddVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.AddVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("addVpc").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/addVpc").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/addVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchUnlockMachine 批量解锁云电脑
func (c *Client) BatchUnlockMachine(request *model.BatchUnlockMachineRequest) (*model.BatchUnlockMachineResponse, error) {
    return c.BatchUnlockMachineWithConfig(request, nil)
}

// BatchUnlockMachineWithConfig 批量解锁云电脑
func (c *Client) BatchUnlockMachineWithConfig(request *model.BatchUnlockMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchUnlockMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchUnlockMachine").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchUnlockMachine").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchUnlockMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchUnlockMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ModifyDeviceTrustNum 设备可信认证数量配置
func (c *Client) ModifyDeviceTrustNum(request *model.ModifyDeviceTrustNumRequest) (*model.ModifyDeviceTrustNumResponse, error) {
    return c.ModifyDeviceTrustNumWithConfig(request, nil)
}

// ModifyDeviceTrustNumWithConfig 设备可信认证数量配置
func (c *Client) ModifyDeviceTrustNumWithConfig(request *model.ModifyDeviceTrustNumRequest, runtimeConfig *config.RuntimeConfig) (*model.ModifyDeviceTrustNumResponse, error) {
    params := param.NewParamsBuilder().
        Action("modifyDeviceTrustNum").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/modifyDeviceTrustNum").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/modifyDeviceTrustNum").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ModifyDeviceTrustNumResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddNetwork 创建子网
func (c *Client) AddNetwork(request *model.AddNetworkRequest) (*model.AddNetworkResponse, error) {
    return c.AddNetworkWithConfig(request, nil)
}

// AddNetworkWithConfig 创建子网
func (c *Client) AddNetworkWithConfig(request *model.AddNetworkRequest, runtimeConfig *config.RuntimeConfig) (*model.AddNetworkResponse, error) {
    params := param.NewParamsBuilder().
        Action("addNetwork").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/addNetwork").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/addNetwork").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddNetworkResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateComputerPolicy 编辑电脑策略
func (c *Client) UpdateComputerPolicy(request *model.UpdateComputerPolicyRequest) (*model.UpdateComputerPolicyResponse, error) {
    return c.UpdateComputerPolicyWithConfig(request, nil)
}

// UpdateComputerPolicyWithConfig 编辑电脑策略
func (c *Client) UpdateComputerPolicyWithConfig(request *model.UpdateComputerPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateComputerPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateComputerPolicy").
        Uri("/cem_openapi/ccacem/api/v1/computerPolicy/updatePolicy").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/computerPolicy/updatePolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateComputerPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetVpc 查询vpc列表
func (c *Client) GetVpc(request *model.GetVpcRequest) (*model.GetVpcResponse, error) {
    return c.GetVpcWithConfig(request, nil)
}

// GetVpcWithConfig 查询vpc列表
func (c *Client) GetVpcWithConfig(request *model.GetVpcRequest, runtimeConfig *config.RuntimeConfig) (*model.GetVpcResponse, error) {
    params := param.NewParamsBuilder().
        Action("getVpc").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/getVpc").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/getVpc").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetVpcResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetNetworkSelect 切换子网下拉框 
func (c *Client) GetNetworkSelect(request *model.GetNetworkSelectRequest) (*model.GetNetworkSelectResponse, error) {
    return c.GetNetworkSelectWithConfig(request, nil)
}

// GetNetworkSelectWithConfig 切换子网下拉框 
func (c *Client) GetNetworkSelectWithConfig(request *model.GetNetworkSelectRequest, runtimeConfig *config.RuntimeConfig) (*model.GetNetworkSelectResponse, error) {
    params := param.NewParamsBuilder().
        Action("getNetworkSelect").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/getNetworkSelect").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/getNetworkSelect").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetNetworkSelectResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetMachineVNC 查询云电脑VNC
func (c *Client) GetMachineVNC(request *model.GetMachineVNCRequest) (*model.GetMachineVNCResponse, error) {
    return c.GetMachineVNCWithConfig(request, nil)
}

// GetMachineVNCWithConfig 查询云电脑VNC
func (c *Client) GetMachineVNCWithConfig(request *model.GetMachineVNCRequest, runtimeConfig *config.RuntimeConfig) (*model.GetMachineVNCResponse, error) {
    params := param.NewParamsBuilder().
        Action("getMachineVNC").
        Uri("/cem_openapi/ccacem/api/v1/machine/getMachineVNC").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/machine/getMachineVNC").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetMachineVNCResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateSNatRule 编辑SNAT规则
func (c *Client) UpdateSNatRule(request *model.UpdateSNatRuleRequest) (*model.UpdateSNatRuleResponse, error) {
    return c.UpdateSNatRuleWithConfig(request, nil)
}

// UpdateSNatRuleWithConfig 编辑SNAT规则
func (c *Client) UpdateSNatRuleWithConfig(request *model.UpdateSNatRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateSNatRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateSNatRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/updateSNatRule").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/updateSNatRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateSNatRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StartOrStopAutoRenewBw 带宽开启或关闭自动续订
func (c *Client) StartOrStopAutoRenewBw(request *model.StartOrStopAutoRenewBwRequest) (*model.StartOrStopAutoRenewBwResponse, error) {
    return c.StartOrStopAutoRenewBwWithConfig(request, nil)
}

// StartOrStopAutoRenewBwWithConfig 带宽开启或关闭自动续订
func (c *Client) StartOrStopAutoRenewBwWithConfig(request *model.StartOrStopAutoRenewBwRequest, runtimeConfig *config.RuntimeConfig) (*model.StartOrStopAutoRenewBwResponse, error) {
    params := param.NewParamsBuilder().
        Action("startOrStopAutoRenewBw").
        Uri("/ccabusiorder/api/v1/order/bw/startOrStopAutoRenewBw").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/startOrStopAutoRenewBw").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.StartOrStopAutoRenewBwResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateRule 创建规则
func (c *Client) CreateRule(request *model.CreateRuleRequest) (*model.CreateRuleResponse, error) {
    return c.CreateRuleWithConfig(request, nil)
}

// CreateRuleWithConfig 创建规则
func (c *Client) CreateRuleWithConfig(request *model.CreateRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("createRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/createRule").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/createRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OperateMachineByAvailable 云电脑开机接口
func (c *Client) OperateMachineByAvailable(request *model.OperateMachineByAvailableRequest) (*model.OperateMachineByAvailableResponse, error) {
    return c.OperateMachineByAvailableWithConfig(request, nil)
}

// OperateMachineByAvailableWithConfig 云电脑开机接口
func (c *Client) OperateMachineByAvailableWithConfig(request *model.OperateMachineByAvailableRequest, runtimeConfig *config.RuntimeConfig) (*model.OperateMachineByAvailableResponse, error) {
    params := param.NewParamsBuilder().
        Action("operateMachineByAvailable").
        Uri("/cem_openapi/ccacem/api/v1/machine/available").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/available").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.OperateMachineByAvailableResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetResourceList 查询云电脑列表
func (c *Client) GetResourceList(request *model.GetResourceListRequest) (*model.GetResourceListResponse, error) {
    return c.GetResourceListWithConfig(request, nil)
}

// GetResourceListWithConfig 查询云电脑列表
func (c *Client) GetResourceListWithConfig(request *model.GetResourceListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetResourceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getResourceList").
        Uri("/cem_openapi/ccacem/api/v1/machine/getList").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/machine/getList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetResourceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetCloudbandwidth 创建SNAT规则-获取未关联规则的带宽
func (c *Client) GetCloudbandwidth(request *model.GetCloudbandwidthRequest) (*model.GetCloudbandwidthResponse, error) {
    return c.GetCloudbandwidthWithConfig(request, nil)
}

// GetCloudbandwidthWithConfig 创建SNAT规则-获取未关联规则的带宽
func (c *Client) GetCloudbandwidthWithConfig(request *model.GetCloudbandwidthRequest, runtimeConfig *config.RuntimeConfig) (*model.GetCloudbandwidthResponse, error) {
    params := param.NewParamsBuilder().
        Action("getCloudbandwidth").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/getCloudbandwidth").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/getCloudbandwidth").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetCloudbandwidthResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UnLockDeviceById 解锁终端
func (c *Client) UnLockDeviceById(request *model.UnLockDeviceByIdRequest) (*model.UnLockDeviceByIdResponse, error) {
    return c.UnLockDeviceByIdWithConfig(request, nil)
}

// UnLockDeviceByIdWithConfig 解锁终端
func (c *Client) UnLockDeviceByIdWithConfig(request *model.UnLockDeviceByIdRequest, runtimeConfig *config.RuntimeConfig) (*model.UnLockDeviceByIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("unLockDeviceById").
        Uri("/cem_openapi/ccacem/api/v1/device/unLockDeviceById").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/device/unLockDeviceById").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UnLockDeviceByIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OperateMachineByReload 云电脑重装接口
func (c *Client) OperateMachineByReload(request *model.OperateMachineByReloadRequest) (*model.OperateMachineByReloadResponse, error) {
    return c.OperateMachineByReloadWithConfig(request, nil)
}

// OperateMachineByReloadWithConfig 云电脑重装接口
func (c *Client) OperateMachineByReloadWithConfig(request *model.OperateMachineByReloadRequest, runtimeConfig *config.RuntimeConfig) (*model.OperateMachineByReloadResponse, error) {
    params := param.NewParamsBuilder().
        Action("operateMachineByReload").
        Uri("/cem_openapi/ccacem/api/v1/machine/reload").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/reload").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.OperateMachineByReloadResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteSecurityGroup 删除安全组
func (c *Client) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
    return c.DeleteSecurityGroupWithConfig(request, nil)
}

// DeleteSecurityGroupWithConfig 删除安全组
func (c *Client) DeleteSecurityGroupWithConfig(request *model.DeleteSecurityGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteSecurityGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteSecurityGroup").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-securityGroup/deleteSecurityGroup").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-securityGroup/deleteSecurityGroup").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteSecurityGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BindUserPolicy 用户策略关联用户
func (c *Client) BindUserPolicy(request *model.BindUserPolicyRequest) (*model.BindUserPolicyResponse, error) {
    return c.BindUserPolicyWithConfig(request, nil)
}

// BindUserPolicyWithConfig 用户策略关联用户
func (c *Client) BindUserPolicyWithConfig(request *model.BindUserPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.BindUserPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindUserPolicy").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/bindUserPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/bindUserPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BindUserPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CemRenameMachine 云电脑重命名接口
func (c *Client) CemRenameMachine(request *model.CemRenameMachineRequest) (*model.CemRenameMachineResponse, error) {
    return c.CemRenameMachineWithConfig(request, nil)
}

// CemRenameMachineWithConfig 云电脑重命名接口
func (c *Client) CemRenameMachineWithConfig(request *model.CemRenameMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.CemRenameMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("cemRenameMachine").
        Uri("/cem_openapi/ccacem/api/v1/machine/rename").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/rename").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.CemRenameMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ActivePage 查询当前会话
func (c *Client) ActivePage(request *model.ActivePageRequest) (*model.ActivePageResponse, error) {
    return c.ActivePageWithConfig(request, nil)
}

// ActivePageWithConfig 查询当前会话
func (c *Client) ActivePageWithConfig(request *model.ActivePageRequest, runtimeConfig *config.RuntimeConfig) (*model.ActivePageResponse, error) {
    params := param.NewParamsBuilder().
        Action("activePage").
        Uri("/cem_openapi/ccacem/api/v1/operation/loginSession/activePage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/operation/loginSession/activePage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ActivePageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OperateMachineByRestart 云电脑重启接口
func (c *Client) OperateMachineByRestart(request *model.OperateMachineByRestartRequest) (*model.OperateMachineByRestartResponse, error) {
    return c.OperateMachineByRestartWithConfig(request, nil)
}

// OperateMachineByRestartWithConfig 云电脑重启接口
func (c *Client) OperateMachineByRestartWithConfig(request *model.OperateMachineByRestartRequest, runtimeConfig *config.RuntimeConfig) (*model.OperateMachineByRestartResponse, error) {
    params := param.NewParamsBuilder().
        Action("operateMachineByRestart").
        Uri("/cem_openapi/ccacem/api/v1/machine/restart").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/restart").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.OperateMachineByRestartResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AcceptShareImageOp 接收共享镜像
func (c *Client) AcceptShareImageOp(request *model.AcceptShareImageOpRequest) (*model.AcceptShareImageOpResponse, error) {
    return c.AcceptShareImageOpWithConfig(request, nil)
}

// AcceptShareImageOpWithConfig 接收共享镜像
func (c *Client) AcceptShareImageOpWithConfig(request *model.AcceptShareImageOpRequest, runtimeConfig *config.RuntimeConfig) (*model.AcceptShareImageOpResponse, error) {
    params := param.NewParamsBuilder().
        Action("acceptShareImageOp").
        Uri("/cem_openapi/ccacem/api/v1/image/acceptShareImageOp").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/acceptShareImageOp").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AcceptShareImageOpResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OperateMachineByShutdown 云电脑关机接口
func (c *Client) OperateMachineByShutdown(request *model.OperateMachineByShutdownRequest) (*model.OperateMachineByShutdownResponse, error) {
    return c.OperateMachineByShutdownWithConfig(request, nil)
}

// OperateMachineByShutdownWithConfig 云电脑关机接口
func (c *Client) OperateMachineByShutdownWithConfig(request *model.OperateMachineByShutdownRequest, runtimeConfig *config.RuntimeConfig) (*model.OperateMachineByShutdownResponse, error) {
    params := param.NewParamsBuilder().
        Action("operateMachineByShutdown").
        Uri("/cem_openapi/ccacem/api/v1/machine/shutdown").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/shutdown").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.OperateMachineByShutdownResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetCcempDeviceList 查询硬终端列表
func (c *Client) GetCcempDeviceList(request *model.GetCcempDeviceListRequest) (*model.GetCcempDeviceListResponse, error) {
    return c.GetCcempDeviceListWithConfig(request, nil)
}

// GetCcempDeviceListWithConfig 查询硬终端列表
func (c *Client) GetCcempDeviceListWithConfig(request *model.GetCcempDeviceListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetCcempDeviceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getCcempDeviceList").
        Uri("/cem_openapi/ccacem/api/v1/device/ccemp/getDeviceList").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/device/ccemp/getDeviceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetCcempDeviceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ModifyUserInfo 修改用户信息
func (c *Client) ModifyUserInfo(request *model.ModifyUserInfoRequest) (*model.ModifyUserInfoResponse, error) {
    return c.ModifyUserInfoWithConfig(request, nil)
}

// ModifyUserInfoWithConfig 修改用户信息
func (c *Client) ModifyUserInfoWithConfig(request *model.ModifyUserInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.ModifyUserInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("modifyUserInfo").
        Uri("/cem_openapi/ccacem/api/v1/userManage/modifyUserInfo").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/modifyUserInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ModifyUserInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CreateUserList 创建用户
func (c *Client) CreateUserList(request *model.CreateUserListRequest) (*model.CreateUserListResponse, error) {
    return c.CreateUserListWithConfig(request, nil)
}

// CreateUserListWithConfig 创建用户
func (c *Client) CreateUserListWithConfig(request *model.CreateUserListRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateUserListResponse, error) {
    params := param.NewParamsBuilder().
        Action("createUserList").
        Uri("/cem_openapi/ccacem/api/v1/userManage/createUserList").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/userManage/createUserList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateUserListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Delete 消息删除
func (c *Client) Delete(request *model.DeleteRequest) (*model.DeleteResponse, error) {
    return c.DeleteWithConfig(request, nil)
}

// DeleteWithConfig 消息删除
func (c *Client) DeleteWithConfig(request *model.DeleteRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteResponse, error) {
    params := param.NewParamsBuilder().
        Action("delete").
        Uri("/cem_openapi/ccacem/api/v1//message/delete").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/message/delete").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchLockMachine 批量锁定云电脑
func (c *Client) BatchLockMachine(request *model.BatchLockMachineRequest) (*model.BatchLockMachineResponse, error) {
    return c.BatchLockMachineWithConfig(request, nil)
}

// BatchLockMachineWithConfig 批量锁定云电脑
func (c *Client) BatchLockMachineWithConfig(request *model.BatchLockMachineRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchLockMachineResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchLockMachine").
        Uri("/cem_openapi/ccacem/api/v1/machine/batchLockMachine").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/batchLockMachine").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchLockMachineResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Allocation 为用户分配云电脑
func (c *Client) Allocation(request *model.AllocationRequest) (*model.AllocationResponse, error) {
    return c.AllocationWithConfig(request, nil)
}

// AllocationWithConfig 为用户分配云电脑
func (c *Client) AllocationWithConfig(request *model.AllocationRequest, runtimeConfig *config.RuntimeConfig) (*model.AllocationResponse, error) {
    params := param.NewParamsBuilder().
        Action("allocation").
        Uri("/cem_openapi/ccacem/api/v1/resource/allocation").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource/allocation").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AllocationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchResetPwd 批量重置密码
func (c *Client) BatchResetPwd(request *model.BatchResetPwdRequest) (*model.BatchResetPwdResponse, error) {
    return c.BatchResetPwdWithConfig(request, nil)
}

// BatchResetPwdWithConfig 批量重置密码
func (c *Client) BatchResetPwdWithConfig(request *model.BatchResetPwdRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchResetPwdResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchResetPwd").
        Uri("/cem_openapi/ccacem/api/v1/userManage/batchResetPwd").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/batchResetPwd").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchResetPwdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeviceBindUser 绑定用户
func (c *Client) DeviceBindUser(request *model.DeviceBindUserRequest) (*model.DeviceBindUserResponse, error) {
    return c.DeviceBindUserWithConfig(request, nil)
}

// DeviceBindUserWithConfig 绑定用户
func (c *Client) DeviceBindUserWithConfig(request *model.DeviceBindUserRequest, runtimeConfig *config.RuntimeConfig) (*model.DeviceBindUserResponse, error) {
    params := param.NewParamsBuilder().
        Action("deviceBindUser").
        Uri("/cem_openapi/ccacem/api/v1/device/deviceBindUser").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/device/deviceBindUser").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeviceBindUserResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Reallocate 为云电脑重新分配用户
func (c *Client) Reallocate(request *model.ReallocateRequest) (*model.ReallocateResponse, error) {
    return c.ReallocateWithConfig(request, nil)
}

// ReallocateWithConfig 为云电脑重新分配用户
func (c *Client) ReallocateWithConfig(request *model.ReallocateRequest, runtimeConfig *config.RuntimeConfig) (*model.ReallocateResponse, error) {
    params := param.NewParamsBuilder().
        Action("reallocate").
        Uri("/cem_openapi/ccacem/api/v1/resource/reallocate").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource/reallocate").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ReallocateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BatchMoveUser 移动用户
func (c *Client) BatchMoveUser(request *model.BatchMoveUserRequest) (*model.BatchMoveUserResponse, error) {
    return c.BatchMoveUserWithConfig(request, nil)
}

// BatchMoveUserWithConfig 移动用户
func (c *Client) BatchMoveUserWithConfig(request *model.BatchMoveUserRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchMoveUserResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchMoveUser").
        Uri("/cem_openapi/ccacem/api/v1/userManage/batchMoveUser").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/batchMoveUser").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BatchMoveUserResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SubmitNat 订购nat
func (c *Client) SubmitNat(request *model.SubmitNatRequest) (*model.SubmitNatResponse, error) {
    return c.SubmitNatWithConfig(request, nil)
}

// SubmitNatWithConfig 订购nat
func (c *Client) SubmitNatWithConfig(request *model.SubmitNatRequest, runtimeConfig *config.RuntimeConfig) (*model.SubmitNatResponse, error) {
    params := param.NewParamsBuilder().
        Action("submitNat").
        Uri("/ccabusiorder/api/v1/order/nat/submitNat").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/nat/submitNat").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.SubmitNatResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// EditCompanyTag 编辑企业ID
func (c *Client) EditCompanyTag(request *model.EditCompanyTagRequest) (*model.EditCompanyTagResponse, error) {
    return c.EditCompanyTagWithConfig(request, nil)
}

// EditCompanyTagWithConfig 编辑企业ID
func (c *Client) EditCompanyTagWithConfig(request *model.EditCompanyTagRequest, runtimeConfig *config.RuntimeConfig) (*model.EditCompanyTagResponse, error) {
    params := param.NewParamsBuilder().
        Action("editCompanyTag").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/editCompanyTag").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/tenant/config/editCompanyTag").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.EditCompanyTagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetMachineInfoDetail 查看电脑详情
func (c *Client) GetMachineInfoDetail(request *model.GetMachineInfoDetailRequest) (*model.GetMachineInfoDetailResponse, error) {
    return c.GetMachineInfoDetailWithConfig(request, nil)
}

// GetMachineInfoDetailWithConfig 查看电脑详情
func (c *Client) GetMachineInfoDetailWithConfig(request *model.GetMachineInfoDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.GetMachineInfoDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("getMachineInfoDetail").
        Uri("/cem_openapi/ccacem/api/v1/machine/machineInfoDetail").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/machineInfoDetail").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetMachineInfoDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdatePolicy 编辑定时任务
func (c *Client) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
    return c.UpdatePolicyWithConfig(request, nil)
}

// UpdatePolicyWithConfig 编辑定时任务
func (c *Client) UpdatePolicyWithConfig(request *model.UpdatePolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdatePolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("updatePolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/updatePolicy").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/resource-policy/updatePolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdatePolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetTaskRecordListByInstanceId 根据资源实例获取定时任务执行记录列表
func (c *Client) GetTaskRecordListByInstanceId(request *model.GetTaskRecordListByInstanceIdRequest) (*model.GetTaskRecordListByInstanceIdResponse, error) {
    return c.GetTaskRecordListByInstanceIdWithConfig(request, nil)
}

// GetTaskRecordListByInstanceIdWithConfig 根据资源实例获取定时任务执行记录列表
func (c *Client) GetTaskRecordListByInstanceIdWithConfig(request *model.GetTaskRecordListByInstanceIdRequest, runtimeConfig *config.RuntimeConfig) (*model.GetTaskRecordListByInstanceIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("getTaskRecordListByInstanceId").
        Uri("/cem_openapi/ccacem/api/v1/machine/getTaskRecordListByInstanceId").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/machine/getTaskRecordListByInstanceId").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetTaskRecordListByInstanceIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CancelInstanceList 退订带宽列表
func (c *Client) CancelInstanceList(request *model.CancelInstanceListRequest) (*model.CancelInstanceListResponse, error) {
    return c.CancelInstanceListWithConfig(request, nil)
}

// CancelInstanceListWithConfig 退订带宽列表
func (c *Client) CancelInstanceListWithConfig(request *model.CancelInstanceListRequest, runtimeConfig *config.RuntimeConfig) (*model.CancelInstanceListResponse, error) {
    params := param.NewParamsBuilder().
        Action("cancelInstanceList").
        Uri("/ccabusiorder/api/v1/order/bw/cancelInstanceList").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/cancelInstanceList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CancelInstanceListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ModifyContactInfo 企业联系人信息配置
func (c *Client) ModifyContactInfo(request *model.ModifyContactInfoRequest) (*model.ModifyContactInfoResponse, error) {
    return c.ModifyContactInfoWithConfig(request, nil)
}

// ModifyContactInfoWithConfig 企业联系人信息配置
func (c *Client) ModifyContactInfoWithConfig(request *model.ModifyContactInfoRequest, runtimeConfig *config.RuntimeConfig) (*model.ModifyContactInfoResponse, error) {
    params := param.NewParamsBuilder().
        Action("modifyContactInfo").
        Uri("/cem_openapi/ccacem/api/v1/tenant/config/modifyContactInfo").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/tenant/config/modifyContactInfo").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ModifyContactInfoResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// Page 查询管理员日志
func (c *Client) Page(request *model.PageRequest) (*model.PageResponse, error) {
    return c.PageWithConfig(request, nil)
}

// PageWithConfig 查询管理员日志
func (c *Client) PageWithConfig(request *model.PageRequest, runtimeConfig *config.RuntimeConfig) (*model.PageResponse, error) {
    params := param.NewParamsBuilder().
        Action("page").
        Uri("/cem_openapi/ccacem/api/v1/manager-operate-log/list").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/manager-operate-log/list").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.PageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetUserInfoList 查询用户列表信息
func (c *Client) GetUserInfoList(request *model.GetUserInfoListRequest) (*model.GetUserInfoListResponse, error) {
    return c.GetUserInfoListWithConfig(request, nil)
}

// GetUserInfoListWithConfig 查询用户列表信息
func (c *Client) GetUserInfoListWithConfig(request *model.GetUserInfoListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetUserInfoListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getUserInfoList").
        Uri("/cem_openapi/ccacem/api/v1/userManage/getUserInfoList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/getUserInfoList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetUserInfoListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// BwChange 带宽变更接口
func (c *Client) BwChange(request *model.BwChangeRequest) (*model.BwChangeResponse, error) {
    return c.BwChangeWithConfig(request, nil)
}

// BwChangeWithConfig 带宽变更接口
func (c *Client) BwChangeWithConfig(request *model.BwChangeRequest, runtimeConfig *config.RuntimeConfig) (*model.BwChangeResponse, error) {
    params := param.NewParamsBuilder().
        Action("bwChange").
        Uri("/ccabusiorder/api/v1/order/bw/bwChange").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/bwChange").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BwChangeResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetOrgList 查询组织列表
func (c *Client) GetOrgList(request *model.GetOrgListRequest) (*model.GetOrgListResponse, error) {
    return c.GetOrgListWithConfig(request, nil)
}

// GetOrgListWithConfig 查询组织列表
func (c *Client) GetOrgListWithConfig(request *model.GetOrgListRequest, runtimeConfig *config.RuntimeConfig) (*model.GetOrgListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getOrgList").
        Uri("/cem_openapi/ccacem/api/v1/userManage/getOrgList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userManage/getOrgList").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetOrgListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetImage 查询指定电脑可用镜像列表
func (c *Client) GetImage(request *model.GetImageRequest) (*model.GetImageResponse, error) {
    return c.GetImageWithConfig(request, nil)
}

// GetImageWithConfig 查询指定电脑可用镜像列表
func (c *Client) GetImageWithConfig(request *model.GetImageRequest, runtimeConfig *config.RuntimeConfig) (*model.GetImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("getImage").
        Uri("/cem_openapi/ccacem/api/v1/image/getMachineImageUsage").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/image/getMachineImageUsage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPolicyInfoByInstanceId 查询电脑关联的定时任务
func (c *Client) GetPolicyInfoByInstanceId(request *model.GetPolicyInfoByInstanceIdRequest) (*model.GetPolicyInfoByInstanceIdResponse, error) {
    return c.GetPolicyInfoByInstanceIdWithConfig(request, nil)
}

// GetPolicyInfoByInstanceIdWithConfig 查询电脑关联的定时任务
func (c *Client) GetPolicyInfoByInstanceIdWithConfig(request *model.GetPolicyInfoByInstanceIdRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPolicyInfoByInstanceIdResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPolicyInfoByInstanceId").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getPolicyInfoByInstanceId").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getPolicyInfoByInstanceId").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetPolicyInfoByInstanceIdResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CopyPolicy 复制定时任务
func (c *Client) CopyPolicy(request *model.CopyPolicyRequest) (*model.CopyPolicyResponse, error) {
    return c.CopyPolicyWithConfig(request, nil)
}

// CopyPolicyWithConfig 复制定时任务
func (c *Client) CopyPolicyWithConfig(request *model.CopyPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.CopyPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("copyPolicy").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/copyPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/copyPolicy").
        Protocol("http").
        ContentType("").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.CopyPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteSnapshot 删除快照
func (c *Client) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
    return c.DeleteSnapshotWithConfig(request, nil)
}

// DeleteSnapshotWithConfig 删除快照
func (c *Client) DeleteSnapshotWithConfig(request *model.DeleteSnapshotRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteSnapshotResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteSnapshot").
        Uri("/cem_openapi/ccacem/api/v1/snapshot/deleteSnapshot").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/snapshot/deleteSnapshot").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteSnapshotResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// RenewBw 续订带宽
func (c *Client) RenewBw(request *model.RenewBwRequest) (*model.RenewBwResponse, error) {
    return c.RenewBwWithConfig(request, nil)
}

// RenewBwWithConfig 续订带宽
func (c *Client) RenewBwWithConfig(request *model.RenewBwRequest, runtimeConfig *config.RuntimeConfig) (*model.RenewBwResponse, error) {
    params := param.NewParamsBuilder().
        Action("renewBw").
        Uri("/ccabusiorder/api/v1/order/bw/renewBw").
        GatewayUri("/api/query/clouddesktopnew/ccabusiorder/api/v1/order/bw/renewBw").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RenewBwResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ListSNatRule 查询SNAT规则列表
func (c *Client) ListSNatRule(request *model.ListSNatRuleRequest) (*model.ListSNatRuleResponse, error) {
    return c.ListSNatRuleWithConfig(request, nil)
}

// ListSNatRuleWithConfig 查询SNAT规则列表
func (c *Client) ListSNatRuleWithConfig(request *model.ListSNatRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.ListSNatRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("listSNatRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/listSNatRule").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/listSNatRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ListSNatRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetSubnet 创建SNAT规则-获取未关联规则的子网
func (c *Client) GetSubnet(request *model.GetSubnetRequest) (*model.GetSubnetResponse, error) {
    return c.GetSubnetWithConfig(request, nil)
}

// GetSubnetWithConfig 创建SNAT规则-获取未关联规则的子网
func (c *Client) GetSubnetWithConfig(request *model.GetSubnetRequest, runtimeConfig *config.RuntimeConfig) (*model.GetSubnetResponse, error) {
    params := param.NewParamsBuilder().
        Action("getSubnet").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/getSubnet").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/getSubnet").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.GetSubnetResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateNetwork 编辑子网
func (c *Client) UpdateNetwork(request *model.UpdateNetworkRequest) (*model.UpdateNetworkResponse, error) {
    return c.UpdateNetworkWithConfig(request, nil)
}

// UpdateNetworkWithConfig 编辑子网
func (c *Client) UpdateNetworkWithConfig(request *model.UpdateNetworkRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateNetworkResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateNetwork").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-vpc/updateNetwork").
        GatewayUri("/api/query/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-vpc/updateNetwork").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UpdateNetworkResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteSNatRule 删除SNAT规则
func (c *Client) DeleteSNatRule(request *model.DeleteSNatRuleRequest) (*model.DeleteSNatRuleResponse, error) {
    return c.DeleteSNatRuleWithConfig(request, nil)
}

// DeleteSNatRuleWithConfig 删除SNAT规则
func (c *Client) DeleteSNatRuleWithConfig(request *model.DeleteSNatRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteSNatRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteSNatRule").
        Uri("/cem_openapi/ccacem/api/v1/ecloud-nat/deleteSNatRule").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/ecloud-nat/deleteSNatRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.DeleteSNatRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// GetPoolList 获取可以创建定时任务的资源池列表
func (c *Client) GetPoolList() (*model.GetPoolListResponse, error) {
    return c.GetPoolListWithConfig(nil)
}

// GetPoolListWithConfig 获取可以创建定时任务的资源池列表
func (c *Client) GetPoolListWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetPoolListResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPoolList").
        Uri("/cem_openapi/ccacem/api/v1/resource-policy/getPoolList").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/resource-policy/getPoolList").
        Protocol("http").
        ContentType("").
        Method("GET").
        Build()
    returnValue := &model.GetPoolListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddUserPolicy 创建用户策略
func (c *Client) AddUserPolicy(request *model.AddUserPolicyRequest) (*model.AddUserPolicyResponse, error) {
    return c.AddUserPolicyWithConfig(request, nil)
}

// AddUserPolicyWithConfig 创建用户策略
func (c *Client) AddUserPolicyWithConfig(request *model.AddUserPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.AddUserPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("addUserPolicy").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/addUserPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/addUserPolicy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddUserPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UserLoginStatistics 查询操作失败统计
func (c *Client) UserLoginStatistics(request *model.UserLoginStatisticsRequest) (*model.UserLoginStatisticsResponse, error) {
    return c.UserLoginStatisticsWithConfig(request, nil)
}

// UserLoginStatisticsWithConfig 查询操作失败统计
func (c *Client) UserLoginStatisticsWithConfig(request *model.UserLoginStatisticsRequest, runtimeConfig *config.RuntimeConfig) (*model.UserLoginStatisticsResponse, error) {
    params := param.NewParamsBuilder().
        Action("userLoginStatistics").
        Uri("/cem_openapi/ccacem/api/v1/operation/userLoginStatistics/all").
        GatewayUri("/api/clouddesktopnew/cem_openapi/ccacem/api/v1/operation/userLoginStatistics/all").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.UserLoginStatisticsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteUserPolicy 删除用户策略
func (c *Client) DeleteUserPolicy(request *model.DeleteUserPolicyRequest) (*model.DeleteUserPolicyResponse, error) {
    return c.DeleteUserPolicyWithConfig(request, nil)
}

// DeleteUserPolicyWithConfig 删除用户策略
func (c *Client) DeleteUserPolicyWithConfig(request *model.DeleteUserPolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteUserPolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteUserPolicy").
        Uri("/cem_openapi/ccacem/api/v1/userPolicy/deleteUserPolicy").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/userPolicy/deleteUserPolicy").
        Protocol("http").
        ContentType("").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteUserPolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// RetryAddImage 重试镜像
func (c *Client) RetryAddImage(request *model.RetryAddImageRequest) (*model.RetryAddImageResponse, error) {
    return c.RetryAddImageWithConfig(request, nil)
}

// RetryAddImageWithConfig 重试镜像
func (c *Client) RetryAddImageWithConfig(request *model.RetryAddImageRequest, runtimeConfig *config.RuntimeConfig) (*model.RetryAddImageResponse, error) {
    params := param.NewParamsBuilder().
        Action("retryAddImage").
        Uri("/cem_openapi/ccacem/api/v1/image/retryAddImage").
        GatewayUri("/api/Computer_api/cem_openapi/ccacem/api/v1/image/retryAddImage").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.RetryAddImageResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

