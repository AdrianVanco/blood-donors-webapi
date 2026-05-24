/*
 * Blood Donors Api
 *
 * Blood donor registration and management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: vancoa@stuba.sk
 */

package blood_donors

import (
	"time"
)

type Donor struct {

	// Unique id of the donor record at this donation site
	Id string `json:"id"`

	// Full name of the blood donor
	Name string `json:"name,omitempty"`

	// Unique and immutable registration number of the donor
	DonorId string `json:"donorId"`

	// Sex of the donor: "M" (muž) or "F" (žena)
	Sex string `json:"sex,omitempty"`

	// Blood type of the donor (e.g. A+, 0-, AB+)
	BloodType string `json:"bloodType,omitempty"`

	// Contact e-mail of the donor
	Email string `json:"email,omitempty"`

	// Contact phone number of the donor
	Phone string `json:"phone,omitempty"`

	// Preferred donation type: "blood", "plasma" or "both"
	PreferredDonationType string `json:"preferredDonationType,omitempty"`

	// Preferred donation site (id) of the donor
	PreferredSite string `json:"preferredSite,omitempty"`

	// Whether the donor is generally eligible to donate
	Eligible bool `json:"eligible"`

	// Reason why the donor is not eligible (used when eligible is false)
	EligibilityNote string `json:"eligibilityNote,omitempty"`

	// Timestamp since when the donor is registered
	RegisteredSince time.Time `json:"registeredSince"`

	// Reserved appointment time of the next donation. Ignored on post.
	AppointmentStart time.Time `json:"appointmentStart,omitempty"`

	// History of the donor's donations / appointments
	Donations []Donation `json:"donations,omitempty"`
}
