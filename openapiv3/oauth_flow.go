package openapiv3

// OAuthFlow is the configuration details for a supported OAuth Flow ([ref]).
//
// Example:
//
//		{
//		  "authorizationUrl": "https://example.com/api/oauth/dialog",
//		  "scopes": {
//		    "write:pets": "modify pets in your account",
//		    "read:pets": "read your pets"
//		  }
//	},
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#oauth-flow-object
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type OAuthFlow struct {
	// REQUIRED.
	// The authorization URL to be used for this flow. This MUST be in the form of a URL.
	// The OAuth2 standard requires the use of TLS.
	AuthorizationURL string `json:"authorizationUrl"`
	// REQUIRED.
	// The token URL to be used for this flow. This MUST be in the form of a URL.
	// The OAuth2 standard requires the use of TLS.
	TokenURL string `json:"tokenUrl"`
	// The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL.
	// The OAuth2 standard requires the use of TLS.
	RefreshURL string `json:"refreshUrl,omitempty"`
	// REQUIRED.
	// The available scopes for the OAuth2 security scheme.
	// A map between the scope name and a short description for it.
	// The map MAY be empty.
	Scopes map[string]string `json:"scopes"`
}
