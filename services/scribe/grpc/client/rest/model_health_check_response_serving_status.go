/*
 * types/v1/service.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package rest

type HealthCheckResponseServingStatus string

// List of HealthCheckResponseServingStatus
const (
	UNKNOWN_HealthCheckResponseServingStatus         HealthCheckResponseServingStatus = "UNKNOWN"
	SERVING_HealthCheckResponseServingStatus         HealthCheckResponseServingStatus = "SERVING"
	NOT_SERVING_HealthCheckResponseServingStatus     HealthCheckResponseServingStatus = "NOT_SERVING"
	SERVICE_UNKNOWN_HealthCheckResponseServingStatus HealthCheckResponseServingStatus = "SERVICE_UNKNOWN"
)
