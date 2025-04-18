/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ContainerResourceMetricSource = map[string]string{
	"":                         "ContainerResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in the requests and limits, describing a single container in each of the pods of the current scale target(e.g. CPU or memory). The values will be averaged together before being compared to the target. Such metrics are built into Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source. Only one \"target\" type should be set.",
	"name":                     "name is the name of the resource in question.",
	"targetAverageUtilization": "targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.",
	"targetAverageValue":       "targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the \"pods\" metric source type.",
	"container":                "container is the name of the container in the pods of the scaling target.",
}

func (ContainerResourceMetricSource) SwaggerDoc() map[string]string {
	return map_ContainerResourceMetricSource
}

var map_ContainerResourceMetricStatus = map[string]string{
	"":                          "ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"name":                      "name is the name of the resource in question.",
	"currentAverageUtilization": "currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.",
	"currentAverageValue":       "currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the \"pods\" metric source type. It will always be set, regardless of the corresponding metric specification.",
	"container":                 "container is the name of the container in the pods of the scaling taget",
}

func (ContainerResourceMetricStatus) SwaggerDoc() map[string]string {
	return map_ContainerResourceMetricStatus
}

var map_CrossVersionObjectReference = map[string]string{
	"":           "CrossVersionObjectReference contains enough information to let you identify the referred resource.",
	"kind":       "kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
	"name":       "name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	"apiVersion": "apiVersion is the API version of the referent",
}

func (CrossVersionObjectReference) SwaggerDoc() map[string]string {
	return map_CrossVersionObjectReference
}

var map_ExternalMetricSource = map[string]string{
	"":                   "ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
	"metricName":         "metricName is the name of the metric in question.",
	"metricSelector":     "metricSelector is used to identify a specific time series within a given metric.",
	"targetValue":        "targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.",
	"targetAverageValue": "targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.",
}

func (ExternalMetricSource) SwaggerDoc() map[string]string {
	return map_ExternalMetricSource
}

var map_ExternalMetricStatus = map[string]string{
	"":                    "ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.",
	"metricName":          "metricName is the name of a metric used for autoscaling in metric system.",
	"metricSelector":      "metricSelector is used to identify a specific time series within a given metric.",
	"currentValue":        "currentValue is the current value of the metric (as a quantity)",
	"currentAverageValue": "currentAverageValue is the current value of metric averaged over autoscaled pods.",
}

func (ExternalMetricStatus) SwaggerDoc() map[string]string {
	return map_ExternalMetricStatus
}

var map_HorizontalPodAutoscaler = map[string]string{
	"":         "configuration of a horizontal pod autoscaler.",
	"metadata": "Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec defines the behaviour of autoscaler. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.",
	"status":   "status is the current information about the autoscaler.",
}

func (HorizontalPodAutoscaler) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscaler
}

var map_HorizontalPodAutoscalerCondition = map[string]string{
	"":                   "HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.",
	"type":               "type describes the current condition",
	"status":             "status is the status of the condition (True, False, Unknown)",
	"lastTransitionTime": "lastTransitionTime is the last time the condition transitioned from one status to another",
	"reason":             "reason is the reason for the condition's last transition.",
	"message":            "message is a human-readable explanation containing details about the transition",
}

func (HorizontalPodAutoscalerCondition) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerCondition
}

var map_HorizontalPodAutoscalerList = map[string]string{
	"":         "list of horizontal pod autoscaler objects.",
	"metadata": "Standard list metadata.",
	"items":    "items is the list of horizontal pod autoscaler objects.",
}

func (HorizontalPodAutoscalerList) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerList
}

var map_HorizontalPodAutoscalerSpec = map[string]string{
	"":                               "specification of a horizontal pod autoscaler.",
	"scaleTargetRef":                 "reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.",
	"minReplicas":                    "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
	"maxReplicas":                    "maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
	"targetCPUUtilizationPercentage": "targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.",
}

func (HorizontalPodAutoscalerSpec) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerSpec
}

var map_HorizontalPodAutoscalerStatus = map[string]string{
	"":                                "current status of a horizontal pod autoscaler",
	"observedGeneration":              "observedGeneration is the most recent generation observed by this autoscaler.",
	"lastScaleTime":                   "lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.",
	"currentReplicas":                 "currentReplicas is the current number of replicas of pods managed by this autoscaler.",
	"desiredReplicas":                 "desiredReplicas is the  desired number of replicas of pods managed by this autoscaler.",
	"currentCPUUtilizationPercentage": "currentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.",
}

func (HorizontalPodAutoscalerStatus) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerStatus
}

var map_MetricSpec = map[string]string{
	"":                  "MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).",
	"type":              "type is the type of metric source.  It should be one of \"ContainerResource\", \"External\", \"Object\", \"Pods\" or \"Resource\", each mapping to a matching field in the object.",
	"object":            "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
	"pods":              "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
	"resource":          "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"containerResource": "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"external":          "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
}

func (MetricSpec) SwaggerDoc() map[string]string {
	return map_MetricSpec
}

var map_MetricStatus = map[string]string{
	"":                  "MetricStatus describes the last-read state of a single metric.",
	"type":              "type is the type of metric source.  It will be one of \"ContainerResource\", \"External\", \"Object\", \"Pods\" or \"Resource\", each corresponds to a matching field in the object.",
	"object":            "object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).",
	"pods":              "pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.",
	"resource":          "resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"containerResource": "containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"external":          "external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).",
}

func (MetricStatus) SwaggerDoc() map[string]string {
	return map_MetricStatus
}

var map_ObjectMetricSource = map[string]string{
	"":             "ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
	"target":       "target is the described Kubernetes object.",
	"metricName":   "metricName is the name of the metric in question.",
	"targetValue":  "targetValue is the target value of the metric (as a quantity).",
	"selector":     "selector is the string-encoded form of a standard kubernetes label selector for the given metric. When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.",
	"averageValue": "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
}

func (ObjectMetricSource) SwaggerDoc() map[string]string {
	return map_ObjectMetricSource
}

var map_ObjectMetricStatus = map[string]string{
	"":             "ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).",
	"target":       "target is the described Kubernetes object.",
	"metricName":   "metricName is the name of the metric in question.",
	"currentValue": "currentValue is the current value of the metric (as a quantity).",
	"selector":     "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the ObjectMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
	"averageValue": "averageValue is the current value of the average of the metric across all relevant pods (as a quantity)",
}

func (ObjectMetricStatus) SwaggerDoc() map[string]string {
	return map_ObjectMetricStatus
}

var map_PodsMetricSource = map[string]string{
	"":                   "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.",
	"metricName":         "metricName is the name of the metric in question",
	"targetAverageValue": "targetAverageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
	"selector":           "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.",
}

func (PodsMetricSource) SwaggerDoc() map[string]string {
	return map_PodsMetricSource
}

var map_PodsMetricStatus = map[string]string{
	"":                    "PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).",
	"metricName":          "metricName is the name of the metric in question",
	"currentAverageValue": "currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)",
	"selector":            "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set in the PodsMetricSource, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.",
}

func (PodsMetricStatus) SwaggerDoc() map[string]string {
	return map_PodsMetricStatus
}

var map_ResourceMetricSource = map[string]string{
	"":                         "ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.  Only one \"target\" type should be set.",
	"name":                     "name is the name of the resource in question.",
	"targetAverageUtilization": "targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.",
	"targetAverageValue":       "targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the \"pods\" metric source type.",
}

func (ResourceMetricSource) SwaggerDoc() map[string]string {
	return map_ResourceMetricSource
}

var map_ResourceMetricStatus = map[string]string{
	"":                          "ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \"pods\" source.",
	"name":                      "name is the name of the resource in question.",
	"currentAverageUtilization": "currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.",
	"currentAverageValue":       "currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the \"pods\" metric source type. It will always be set, regardless of the corresponding metric specification.",
}

func (ResourceMetricStatus) SwaggerDoc() map[string]string {
	return map_ResourceMetricStatus
}

var map_Scale = map[string]string{
	"":         "Scale represents a scaling request for a resource.",
	"metadata": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.",
	"spec":     "spec defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.",
	"status":   "status is the current status of the scale. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status. Read-only.",
}

func (Scale) SwaggerDoc() map[string]string {
	return map_Scale
}

var map_ScaleSpec = map[string]string{
	"":         "ScaleSpec describes the attributes of a scale subresource.",
	"replicas": "replicas is the desired number of instances for the scaled object.",
}

func (ScaleSpec) SwaggerDoc() map[string]string {
	return map_ScaleSpec
}

var map_ScaleStatus = map[string]string{
	"":         "ScaleStatus represents the current status of a scale subresource.",
	"replicas": "replicas is the actual number of observed instances of the scaled object.",
	"selector": "selector is the label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/",
}

func (ScaleStatus) SwaggerDoc() map[string]string {
	return map_ScaleStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
