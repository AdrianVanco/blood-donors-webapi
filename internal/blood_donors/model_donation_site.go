/*
 * Blood Donors Api
 *
 * Blood donor registration and management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: vancoa@stuba.sk
 */

package blood_donors

// DonationSite - transfúzna stanica / odberné miesto (agregát darcov)
type DonationSite struct {

	// Unique identifier of the donation site
	Id string `json:"id"`

	// Human readable display name of the donation site
	Name string `json:"name"`

	// Postal address of the donation site
	Address string `json:"address,omitempty"`

	// Registered donors at this donation site
	Donors []Donor `json:"donors,omitempty"`

	// Predefined donation types offered at this site
	PredefinedDonationTypes []DonationType `json:"predefinedDonationTypes,omitempty"`
}
