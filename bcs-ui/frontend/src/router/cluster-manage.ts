import { RouteConfig } from 'vue-router';
// 集群首页
const Cluster = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/index.vue');
// 创建集群
const ClusterCreate = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/create-cluster/create-cluster.vue');
// 表单模式
const CreateFormCluster = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/create-cluster/create-form-cluster.vue');
// ee版本创建集群流程
const CreateFormClusterEE = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/create-cluster/create-form-cluster-ee.vue');
// import模式
const CreateImportCluster = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/create-cluster/create-import-cluster.vue');
// 集群详情
const ClusterDetail = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-detail.vue');
const ClusterNodeOverview = () => import(/* webpackChunkName: 'cluster' */'@/views/cluster/node-overview.vue');
const Node = () => import(/* webpackChunkName: 'cluster' */'@/views/node/node.vue');
const NodeTemplate = () => import(/* webpackChunkName: 'cluster'  */'@/views/node/node-template/node-template.vue');
const EditNodeTemplate = () => import(/* webpackChunkName: 'cluster' */'@/views/node/node-template/edit-node-template.vue');
const AddClusterNode = () => import(/* webpackChunkName: 'cluster' */'@/views/node/add-cluster-node.vue');
const AutoScalerConfig = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler-tencent/autoscaler-config.vue');
const NodePool = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler-tencent/node-pool.vue');
const NodePoolDetail = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler-tencent/node-pool-detail.vue');
const EditNodePool = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler-tencent/edit-node-pool.vue');
const InternalAutoScalerConfig = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler/autoscaler-config.vue');
const InternalNodePool = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler/node-pool.vue');
const InternalNodePoolDetail = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler/node-pool-detail.vue');
const InternalEditNodePool = () => import(/* webpackChunkName: 'cluster' */'@/views/node/cluster-autoscaler/edit-node-pool.vue');

// 集群管理
export default [
  {
    path: 'clusters',
    name: 'clusterMain',
    component: Cluster,
  },
  // 创建集群
  {
    path: 'clusters/create',
    name: 'clusterCreate',
    component: ClusterCreate,
    meta: {
      menuId: 'CLUSTER',
      title: window.i18n.t('新建集群'),
    },
  },
  // 创建集群 - 表单模式
  {
    path: 'clusters/form',
    name: 'createFormCluster',
    component: window.REGION === 'ieod' ? CreateFormCluster : CreateFormClusterEE,
    meta: {
      menuId: 'CLUSTER',
      title: window.i18n.t('自建集群'),
    },
  },
  // 创建集群 - import导入模式
  {
    path: 'clusters/import',
    name: 'createImportCluster',
    component: CreateImportCluster,
    meta: {
      menuId: 'CLUSTER',
      title: window.i18n.t('导入集群'),
    },
  },
  // 集群详情
  {
    path: 'clusters/:clusterId',
    name: 'clusterDetail',
    props: route => ({ ...route.params, ...route.query }),
    component: ClusterDetail,
    meta: {
      menuId: 'CLUSTER',
    },
  },
  // 集群总览
  {
    path: '',
    name: 'clusterOverview',
    redirect: {
      name: 'clusterDetail',
      query: {
        active: 'overview',
      },
    },
  },
  // 节点列表
  {
    path: '',
    name: 'clusterNode',
    redirect: {
      name: 'clusterDetail',
      query: {
        active: 'node',
      },
    },
    meta: {
      menuId: 'OVERVIEW',
    },
  },
  // 集群里的集群信息
  {
    path: '',
    name: 'clusterInfo',
    redirect: {
      name: 'clusterDetail',
      query: {
        active: 'info',
      },
    },
    meta: {
      menuId: 'OVERVIEW',
    },
  },
  // 集群里的具体节点
  {
    path: 'clusters/:clusterId/nodes/:nodeName/detail',
    name: 'clusterNodeOverview',
    component: ClusterNodeOverview,
    meta: {
      menuId: 'NODE',
    },
  },
  {
    path: 'nodes',
    name: 'nodeMain',
    component: Node,
    meta: {
      title: window.i18n.t('节点'),
      hideBack: true,
    },
  },
  {
    path: 'node-template',
    name: 'nodeTemplate',
    component: NodeTemplate,
    meta: {
      menuId: 'NODETEMPLATE',
    },
  },
  {
    path: 'node-template/create',
    name: 'addNodeTemplate',
    component: EditNodeTemplate,
    meta: {
      title: window.i18n.t('新建节点模板'),
      menuId: 'NODETEMPLATE',
    },
  },
  {
    path: 'node-template/:nodeTemplateID',
    name: 'editNodeTemplate',
    props: true,
    component: EditNodeTemplate,
    meta: {
      title: window.i18n.t('编辑节点模板'),
      menuId: 'NODETEMPLATE',
    },
  },
  {
    path: 'clusters/:clusterId/nodes/add',
    name: 'addClusterNode',
    props: true,
    component: AddClusterNode,
    meta: {
      title: window.i18n.t('添加节点'),
      menuId: 'CLUSTER',
    },
  },
  {
    path: 'clusters/:clusterId/autoscaler',
    name: 'autoScalerConfig',
    props: true,
    component: window.REGION === 'ieod' ? InternalAutoScalerConfig : AutoScalerConfig,
    meta: {
      menuId: 'CLUSTER',
    },
  },
  {
    path: 'cluster/:clusterId/nodepools',
    name: 'nodePool',
    props: true,
    component: window.REGION === 'ieod' ? InternalNodePool : NodePool,
    meta: {
      menuId: 'CLUSTER',
    },
  },
  {
    path: 'cluster/:clusterId/nodepools/:nodeGroupID',
    name: 'editNodePool',
    props: true,
    component: window.REGION === 'ieod' ? InternalEditNodePool : EditNodePool,
    meta: {
      menuId: 'CLUSTER',
    },
  },
  {
    path: 'cluster/:clusterId/nodepools/:nodeGroupID/detail',
    name: 'nodePoolDetail',
    props: true,
    component: window.REGION === 'ieod' ? InternalNodePoolDetail : NodePoolDetail,
    meta: {
      menuId: 'CLUSTER',
    },
  },
] as RouteConfig[];
