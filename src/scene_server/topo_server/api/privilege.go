/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"fmt"
	"net/http"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	frtypes "configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/topo_server/core/types"
)

func init() {
	apiInst.initFuncs = append(apiInst.initFuncs, apiInst.initObject)
}

func (cli *topoAPI) initPrivilege() {
	cli.actions = append(cli.actions, action{Method: http.MethodPost, Path: "/privilege/group/detail/{bk_supplier_account}/{group_id}", HandlerFunc: cli.UpdateUserGroupPrivi})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/privilege/group/detail/{bk_supplier_account}/{group_id}", HandlerFunc: cli.GetUserGroupPrivi})
	cli.actions = append(cli.actions, action{Method: http.MethodGet, Path: "/privilege/user/detail/{bk_supplier_account}/{user_name}", HandlerFunc: cli.GetUserPrivi})

}

// UpdateUserGroupPrivi search user goup
func (cli *topoAPI) UpdateUserGroupPrivi(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {

	priviData := &metadata.PrivilegeUserGroup{}

	_, err := priviData.Parse(data)
	if nil != err {
		blog.Errorf("[api-privilege] failed to parse the input data, error info is %s ", err.Error())
		return nil, params.Err.New(common.CCErrCommParamsIsInvalid, err.Error())
	}

	err = cli.core.PermissionOperation().Permission(params).SetUserGroupPermission(params.Header.OwnerID, pathParams("group_id"), priviData)
	return nil, err
}

// GetUserGroupPrivi search user goup
func (cli *topoAPI) GetUserGroupPrivi(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	fmt.Println("SearchObjectBatch")
	return nil, nil
}

// GetUserPrivi search user goup
func (cli *topoAPI) GetUserPrivi(params types.LogicParams, pathParams, queryParams ParamsGetter, data frtypes.MapStr) (interface{}, error) {
	fmt.Println("SearchObjectBatch")
	return nil, nil
}
