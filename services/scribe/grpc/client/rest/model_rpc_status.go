/*
 * types/v1/service.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package rest

type RpcStatus struct {
	Code    int32                    `json:"code,omitempty"`
	Message string                   `json:"message,omitempty"`
	Details []map[string]interface{} `json:"details,omitempty"`
}
