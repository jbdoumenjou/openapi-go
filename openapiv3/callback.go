package openapiv3

// Callback is a map of possible out-of band callbacks related to the parent operation ([ref]).
// Each value in the map is a Path Item Object that describes a set of requests that may be initiated by the API provider
// and the expected responses.
// The key value used to identify the path item object is an expression,
// evaluated at runtime, that identifies a URL to use for the callback operation.
//
// To describe incoming requests from the API provider independent of another API call, use the webhooks field.
//
// The key that identifies the PathItem Object is a [runtime expression]
// that can be evaluated in the context of a runtime HTTP request/response
// to identify the URL to be used for the callback request.
// A simple example might be $request.body#/url.
// However, using a [runtime expression] the complete HTTP message can be accessed.
// This includes accessing any part of a body that a JSON Pointer [[RFC6901]] can reference.
//
// For example, given the following HTTP request:
//
// POST /subscribe/myevent?queryUrl=https://clientdomain.com/stillrunning HTTP/1.1
// Host: example.org
// Content-Type: application/json
// Content-Length: 187
//
//	{
//	  "failedUrl" : "https://clientdomain.com/failed",
//	  "successUrls" : [
//	    "https://clientdomain.com/fast",
//	    "https://clientdomain.com/medium",
//	    "https://clientdomain.com/slow"
//	  ]
//	}
//
// 201 Created
// Location: https://example.org/subscription/1
//
// The following examples show how the various expressions evaluate,
// assuming the callback operation has a path parameter named eventType and a query parameter named queryUrl.
//
//   - expression: value
//   - $url: https://example.org/subscribe/myevent?queryUrl=https://clientdomain.com/stillrunning
//   - $method: POST
//   - $request.path.eventType: myevent
//   - $request.query.queryUrl: https://clientdomain.com/stillrunning
//   - $request.header.content-Type: application/json
//   - $request.body#/failedUrl: https://clientdomain.com/failed
//   - $request.body#/successUrls/2: https://clientdomain.com/medium
//   - $response.header.Location: https://example.org/subscription/1
//
// Example:
//
//	{
//	 "myCallback": {
//	   "{$request.query.queryUrl}": {
//	     "post": {
//	       "requestBody": {
//	         "description": "Callback payload",
//	         "content": {
//	           "application/json": {
//	             "schema": {
//	               "$ref": "#/components/schemas/SomePayload"
//	             }
//	           }
//	         }
//	       },
//	       "responses": {
//	         "200": {
//	           "description": "callback successfully processed"
//	         }
//	       }
//	     }
//	   }
//	 }
//	}
//
// This object MAY be extended with [Specification Extensions].
//
// [ref]: https://spec.openapis.org/oas/latest.html#callback-object
// [runtime expression]: https://spec.openapis.org/oas/latest.html#runtimeExpression
// [RFC6901]: https://spec.openapis.org/oas/latest.html#bib-RFC6901
// [Specification Extensions]: https://spec.openapis.org/oas/latest.html#specificationExtensions
type Callback map[string]PathItem
