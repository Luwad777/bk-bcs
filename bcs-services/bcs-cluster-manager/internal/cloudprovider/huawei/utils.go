/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package huawei xxx
package huawei

import (
	"fmt"
	"strconv"
	"strings"

	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider"
)

var (
	cloudName = "huawei"
)

// huaweiCloud taskName
const (
	// importClusterTaskTemplate bk-sops add task template
	importClusterTaskTemplate = "cce-import cluster: %s"
	// deleteClusterTaskTemplate bk-sops add task template
	deleteClusterTaskTemplate = "cce-delete cluster: %s" // nolint
	// addNodeTaskTemplate bk-sops add task template
	addNodeTaskTemplate = "cce-add node: %s" // nolint
	// cleanNodeTaskTemplate bk-sops add task template
	cleanNodeTaskTemplate = "cce-remove node: %s" // nolint
	// createNodeGroupTaskTemplate bk-sops add task template
	createNodeGroupTaskTemplate = "cce-create node group: %s/%s"
	// deleteNodeGroupTaskTemplate bk-sops add task template
	deleteNodeGroupTaskTemplate = "cce-delete node group: %s/%s"
	// updateNodeGroupDesiredNode bk-sops add task template
	updateNodeGroupDesiredNodeTemplate = "cce-update node group desired node: %s/%s"
	// cleanNodeGroupNodesTaskTemplate bk-sops add task template
	cleanNodeGroupNodesTaskTemplate = "cce-remove node group nodes: %s/%s"
	// moveNodesToNodeGroupTaskTemplate bk-sops add task template
	moveNodesToNodeGroupTaskTemplate = "cce-move nodes to node group: %s/%s" // nolint
	// switchNodeGroupAutoScalingTaskTemplate bk-sops add task template
	switchNodeGroupAutoScalingTaskTemplate = "cce-switch node group auto scaling: %s/%s"
	// updateAutoScalingOptionTemplate bk-sops add task template
	updateAutoScalingOptionTemplate = "cce-update auto scaling option: %s"
	// switchAutoScalingOptionStatusTemplate bk-sops add task template
	switchAutoScalingOptionStatusTemplate = "cce-switch auto scaling option status: %s"
)

// tasks
var (
	// import cluster task
	registerClusterKubeConfigStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-RegisterClusterKubeConfigTask", cloudName),
		StepName:   "注册集群kubeConfig认证",
	}
	importClusterNodesStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-ImportClusterNodesTask", cloudName),
		StepName:   "导入集群节点",
	}

	// create nodeGroup task
	createCloudNodeGroupStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-CreateCloudNodeGroupTask", cloudName),
		StepName:   "创建云节点组",
	}
	checkCloudNodeGroupStatusStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-CheckCloudNodeGroupStatusTask", cloudName),
		StepName:   "检测云节点组状态",
	}

	// delete nodeGroup task
	deleteNodeGroupStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-DeleteNodeGroupTask", cloudName),
		StepName:   "删除云节点组",
	}

	// update desired nodes task
	applyInstanceMachinesStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-%s", cloudName, cloudprovider.ApplyInstanceMachinesTask),
		StepName:   "申请节点任务",
	}
	checkClusterNodesStatusStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-CheckClusterNodesStatusTask", cloudName),
		StepName:   "检测节点状态",
	}
	updateDesiredNodesDBInfoStep = cloudprovider.StepInfo{ // nolint
		StepMethod: fmt.Sprintf("%s-UpdateDesiredNodesDBInfoTask", cloudName),
		StepName:   "更新节点数据",
	}

	// clean node in nodeGroup task
	cleanNodeGroupNodesStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-CleanNodeGroupNodesTask", cloudName),
		StepName:   "下架节点组节点",
	}
	checkClusterCleanNodsStep = cloudprovider.StepInfo{
		StepMethod: fmt.Sprintf("%s-CheckClusterCleanNodsTask", cloudName),
		StepName:   "检测下架节点状态",
	}
	checkCleanNodeGroupNodesStatusStep = cloudprovider.StepInfo{ // nolint
		StepMethod: fmt.Sprintf("%s-CheckCleanNodeGroupNodesStatusTask", cloudName),
		StepName:   "检查节点组状态",
	}
	updateCleanNodeGroupNodesDBInfoStep = cloudprovider.StepInfo{ // nolint
		StepMethod: fmt.Sprintf("%s-UpdateCleanNodeGroupNodesDBInfoTask", cloudName),
		StepName:   "更新节点组数据",
	}
)

// ImportClusterTaskOption 纳管集群
type ImportClusterTaskOption struct {
	Cluster *proto.Cluster
}

// BuildRegisterKubeConfigStep 注册集群kubeConfig
func (ic *ImportClusterTaskOption) BuildRegisterKubeConfigStep(task *proto.Task) {
	registerKubeConfigStep := cloudprovider.InitTaskStep(registerClusterKubeConfigStep)
	registerKubeConfigStep.Params[cloudprovider.ClusterIDKey.String()] = ic.Cluster.ClusterID
	registerKubeConfigStep.Params[cloudprovider.CloudIDKey.String()] = ic.Cluster.Provider

	task.Steps[registerClusterKubeConfigStep.StepMethod] = registerKubeConfigStep
	task.StepSequence = append(task.StepSequence, registerClusterKubeConfigStep.StepMethod)
}

// BuildImportClusterNodesStep 纳管集群节点
func (ic *ImportClusterTaskOption) BuildImportClusterNodesStep(task *proto.Task) {
	importNodesStep := cloudprovider.InitTaskStep(importClusterNodesStep)
	importNodesStep.Params[cloudprovider.ClusterIDKey.String()] = ic.Cluster.ClusterID
	importNodesStep.Params[cloudprovider.CloudIDKey.String()] = ic.Cluster.Provider

	task.Steps[importClusterNodesStep.StepMethod] = importNodesStep
	task.StepSequence = append(task.StepSequence, importClusterNodesStep.StepMethod)
}

// CreateNodeGroupTaskOption 创建节点组
type CreateNodeGroupTaskOption struct {
	Group *proto.NodeGroup
}

// BuildCreateCloudNodeGroupStep 通过云接口创建节点组
func (cn *CreateNodeGroupTaskOption) BuildCreateCloudNodeGroupStep(task *proto.Task) {
	createStep := cloudprovider.InitTaskStep(createCloudNodeGroupStep)

	createStep.Params[cloudprovider.ClusterIDKey.String()] = cn.Group.ClusterID
	createStep.Params[cloudprovider.NodeGroupIDKey.String()] = cn.Group.NodeGroupID
	createStep.Params[cloudprovider.CloudIDKey.String()] = cn.Group.Provider

	task.Steps[createCloudNodeGroupStep.StepMethod] = createStep
	task.StepSequence = append(task.StepSequence, createCloudNodeGroupStep.StepMethod)
}

// BuildCheckCloudNodeGroupStatusStep 检测节点组状态
func (cn *CreateNodeGroupTaskOption) BuildCheckCloudNodeGroupStatusStep(task *proto.Task) {
	checkStep := cloudprovider.InitTaskStep(checkCloudNodeGroupStatusStep)

	checkStep.Params[cloudprovider.ClusterIDKey.String()] = cn.Group.ClusterID
	checkStep.Params[cloudprovider.NodeGroupIDKey.String()] = cn.Group.NodeGroupID
	checkStep.Params[cloudprovider.CloudIDKey.String()] = cn.Group.Provider

	task.Steps[checkCloudNodeGroupStatusStep.StepMethod] = checkStep
	task.StepSequence = append(task.StepSequence, checkCloudNodeGroupStatusStep.StepMethod)
}

// CleanNodeInGroupTaskOption 节点组缩容节点
type CleanNodeInGroupTaskOption struct {
	Group    *proto.NodeGroup
	NodeIPs  []string
	NodeIds  []string
	Operator string
}

// BuildCleanNodeGroupNodesStep 清理节点池节点
func (cn *CleanNodeInGroupTaskOption) BuildCleanNodeGroupNodesStep(task *proto.Task) {
	cleanStep := cloudprovider.InitTaskStep(cleanNodeGroupNodesStep)

	cleanStep.Params[cloudprovider.ClusterIDKey.String()] = cn.Group.ClusterID
	cleanStep.Params[cloudprovider.NodeGroupIDKey.String()] = cn.Group.NodeGroupID
	cleanStep.Params[cloudprovider.CloudIDKey.String()] = cn.Group.Provider
	cleanStep.Params[cloudprovider.NodeIPsKey.String()] = strings.Join(cn.NodeIPs, ",")
	cleanStep.Params[cloudprovider.NodeIDsKey.String()] = strings.Join(cn.NodeIds, ",")

	task.Steps[cleanNodeGroupNodesStep.StepMethod] = cleanStep
	task.StepSequence = append(task.StepSequence, cleanNodeGroupNodesStep.StepMethod)
}

// BuildCheckClusterCleanNodsStep 检测集群清理节点池节点
func (cn *CleanNodeInGroupTaskOption) BuildCheckClusterCleanNodsStep(task *proto.Task) {
	checkStep := cloudprovider.InitTaskStep(checkClusterCleanNodsStep)

	checkStep.Params[cloudprovider.ClusterIDKey.String()] = cn.Group.ClusterID
	checkStep.Params[cloudprovider.NodeGroupIDKey.String()] = cn.Group.NodeGroupID
	checkStep.Params[cloudprovider.CloudIDKey.String()] = cn.Group.Provider

	checkStep.Params[cloudprovider.NodeIPsKey.String()] = strings.Join(cn.NodeIPs, ",")
	checkStep.Params[cloudprovider.NodeIDsKey.String()] = strings.Join(cn.NodeIds, ",")

	task.Steps[checkClusterCleanNodsStep.StepMethod] = checkStep
	task.StepSequence = append(task.StepSequence, checkClusterCleanNodsStep.StepMethod)
}

// DeleteNodeGroupTaskOption 删除节点组
type DeleteNodeGroupTaskOption struct {
	Group *proto.NodeGroup
}

// BuildDeleteNodeGroupStep 删除云节点组
func (dn *DeleteNodeGroupTaskOption) BuildDeleteNodeGroupStep(task *proto.Task) {
	deleteStep := cloudprovider.InitTaskStep(deleteNodeGroupStep)

	deleteStep.Params[cloudprovider.ClusterIDKey.String()] = dn.Group.ClusterID
	deleteStep.Params[cloudprovider.NodeGroupIDKey.String()] = dn.Group.NodeGroupID
	deleteStep.Params[cloudprovider.CloudIDKey.String()] = dn.Group.Provider

	task.Steps[deleteNodeGroupStep.StepMethod] = deleteStep
	task.StepSequence = append(task.StepSequence, deleteNodeGroupStep.StepMethod)
}

// UpdateDesiredNodesTaskOption 扩容节点组节点
type UpdateDesiredNodesTaskOption struct {
	Group    *proto.NodeGroup
	Desired  uint32
	Operator string
}

// BuildApplyInstanceMachinesStep 申请节点实例
func (ud *UpdateDesiredNodesTaskOption) BuildApplyInstanceMachinesStep(task *proto.Task) {
	applyInstanceStep := cloudprovider.InitTaskStep(applyInstanceMachinesStep)

	applyInstanceStep.Params[cloudprovider.ClusterIDKey.String()] = ud.Group.ClusterID
	applyInstanceStep.Params[cloudprovider.NodeGroupIDKey.String()] = ud.Group.NodeGroupID
	applyInstanceStep.Params[cloudprovider.CloudIDKey.String()] = ud.Group.Provider
	applyInstanceStep.Params[cloudprovider.ScalingNodesNumKey.String()] = strconv.Itoa(int(ud.Desired))
	applyInstanceStep.Params[cloudprovider.OperatorKey.String()] = ud.Operator

	task.Steps[applyInstanceMachinesStep.StepMethod] = applyInstanceStep
	task.StepSequence = append(task.StepSequence, applyInstanceMachinesStep.StepMethod)
}

// BuildCheckClusterNodeStatusStep 检测节点实例状态
func (ud *UpdateDesiredNodesTaskOption) BuildCheckClusterNodeStatusStep(task *proto.Task) {
	checkClusterNodeStatusStep := cloudprovider.InitTaskStep(checkClusterNodesStatusStep)

	checkClusterNodeStatusStep.Params[cloudprovider.ClusterIDKey.String()] = ud.Group.ClusterID
	checkClusterNodeStatusStep.Params[cloudprovider.NodeGroupIDKey.String()] = ud.Group.NodeGroupID
	checkClusterNodeStatusStep.Params[cloudprovider.CloudIDKey.String()] = ud.Group.Provider

	task.Steps[checkClusterNodesStatusStep.StepMethod] = checkClusterNodeStatusStep
	task.StepSequence = append(task.StepSequence, checkClusterNodesStatusStep.StepMethod)
}
