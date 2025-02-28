//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomain":          schema_pkg_apis_customdomain_v1alpha1_CustomDomain(ref),
		"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainCondition": schema_pkg_apis_customdomain_v1alpha1_CustomDomainCondition(ref),
		"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainList":      schema_pkg_apis_customdomain_v1alpha1_CustomDomainList(ref),
		"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainSpec":      schema_pkg_apis_customdomain_v1alpha1_CustomDomainSpec(ref),
		"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainStatus":    schema_pkg_apis_customdomain_v1alpha1_CustomDomainStatus(ref),
	}
}

func schema_pkg_apis_customdomain_v1alpha1_CustomDomain(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomDomain is the Schema for the customdomains API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainSpec", "github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_customdomain_v1alpha1_CustomDomainCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomDomainCondition contains details for the current condition of a custom domain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of the condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the status of the condition",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastProbeTime": {
						SchemaProps: spec.SchemaProps{
							Description: "LastProbeTime is the last time we probed the condition.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "LastTransitionTime is the laste time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "Reason is a unique, one-word, CamelCase reason for the condition's last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is a human-readable message indicating details about last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_customdomain_v1alpha1_CustomDomainList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomDomainList contains a list of CustomDomain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomain"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomain", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_customdomain_v1alpha1_CustomDomainSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomDomainSpec defines the desired state of CustomDomain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"domain": {
						SchemaProps: spec.SchemaProps{
							Description: "This field can be used to define the custom domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"certificate": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate points to the custom TLS secret",
							Ref:         ref("k8s.io/api/core/v1.SecretReference"),
						},
					},
					"scope": {
						SchemaProps: spec.SchemaProps{
							Description: "This field determines whether the CustomDomain ingress is internal or external. Defaults to External if empty.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"domain", "certificate"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretReference"},
	}
}

func schema_pkg_apis_customdomain_v1alpha1_CustomDomainStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomDomainStatus defines the observed state of CustomDomain",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "The various conditions for the custom domain",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainCondition"),
									},
								},
							},
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "The overall state of the custom domain",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dnsRecord": {
						SchemaProps: spec.SchemaProps{
							Description: "The DNS record added for the ingress controller",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "The endpoint is a resolvable DNS address for external DNS to point to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"scope": {
						SchemaProps: spec.SchemaProps{
							Description: "The scope dictates whether the ingress controller is internal or external",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "dnsRecord", "endpoint"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-domains-operator/pkg/apis/customdomain/v1alpha1.CustomDomainCondition"},
	}
}
