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

// Donation - jeden záznam v histórii darcu (termín / odber)
type Donation struct {

	// Date and time of the donation appointment
	Date time.Time `json:"date,omitempty"`

	DonationType DonationType `json:"donationType,omitempty"`

	// Stav termínu (Rezervácia dokončená, Spôsobilý - čaká na odber, Odber dokončený, Nespôsobilý, Zrušená rezervácia)
	Status string `json:"status,omitempty"`

	// Doplňujúca informačná poznámka k termínu
	Note string `json:"note,omitempty"`
}
