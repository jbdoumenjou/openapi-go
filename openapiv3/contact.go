package openapiv3

// Contact information for the exposed API ([ref]).
//
// Example:
//
//	{
//	  "name": "API Support",
//	  "url": "https://www.example.com/support",
//	  "email": "support@example.com"
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
// [ref]: https://spec.openapis.org/oas/latest.html#contactObject
type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`
	// The URL pointing to the contact information. This MUST be in the form of a URL.
	URL string `json:"url,omitempty"`
	// The email address of the contact person/organization. This MUST be in the form of an email address.
	// TODO: email validation.
	Email string `json:"email,omitempty"`
}
