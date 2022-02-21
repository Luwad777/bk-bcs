/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/resource/example"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/util"
	clusterRes "github.com/Tencent/bk-bcs/bcs-services/cluster-resources/proto/cluster-resources"
)

func TestSA(t *testing.T) {
	crh := NewClusterResourcesHandler()
	ctx := context.TODO()

	manifest, _ := example.LoadDemoManifest("rbac/simple_service_account")
	resName := util.GetWithDefault(manifest, "metadata.name", "")

	// Create
	createManifest, _ := util.Map2pbStruct(manifest)
	createReq := genResCreateReq(createManifest)
	err := crh.CreateSA(ctx, &createReq, &clusterRes.CommonResp{})
	assert.Nil(t, err)

	// List
	listReq, listResp := genResListReq(), clusterRes.CommonResp{}
	err = crh.ListSA(ctx, &listReq, &listResp)
	assert.Nil(t, err)

	respData := listResp.Data.AsMap()
	assert.Equal(t, "ServiceAccountList", util.GetWithDefault(respData, "manifest.kind", ""))

	// Update
	_ = util.SetItems(manifest, "metadata.annotations", map[string]interface{}{"tKey": "tVal"})
	updateManifest, _ := util.Map2pbStruct(manifest)
	updateReq := genResUpdateReq(updateManifest, resName.(string))
	err = crh.UpdateSA(ctx, &updateReq, &clusterRes.CommonResp{})
	assert.Nil(t, err)

	// Get
	getReq, getResp := genResGetReq(resName.(string)), clusterRes.CommonResp{}
	err = crh.GetSA(ctx, &getReq, &getResp)
	assert.Nil(t, err)

	respData = getResp.Data.AsMap()
	assert.Equal(t, "ServiceAccount", util.GetWithDefault(respData, "manifest.kind", ""))
	assert.Equal(t, "tVal", util.GetWithDefault(respData, "manifest.metadata.annotations.tKey", ""))

	// Delete
	deleteReq := genResDeleteReq(resName.(string))
	err = crh.DeleteSA(ctx, &deleteReq, &clusterRes.CommonResp{})
	assert.Nil(t, err)
}