/*
 * Blood Donors Api
 *
 * Blood donor registration and management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: vancoa@stuba.sk
 */

package blood_donors

// DonationType - typ odberu (darovanie krvi, darovanie plazmy)
type DonationType struct {

	Value string `json:"value"`

	Code string `json:"code,omitempty"`

	// Link to detailed explanation of the donation type
	Reference string `json:"reference,omitempty"`

	TypicalDurationMinutes int32 `json:"typicalDurationMinutes,omitempty"`
}
