/* 
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Occurrence includes information about analysis occurrences for an image. ``
type Occurrence struct {

	// The name of the occurrence in the form \"projects/{project_id}/occurrences/{occurrence_id}\" @OutputOnly
	Name string `json:"name,omitempty"`

	// The unique url of the image or container for which the occurrence applies. Example: https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
	ResourceUrl string `json:"resourceUrl,omitempty"`

	// An analysis note associated with this image, in the form \"projects/{project_id}/notes/{note_id}\" This field can be used as a filter in list requests.
	NoteName string `json:"noteName,omitempty"`

	// This explicitly denotes which of the occurrence details is specified. This field can be used as a filter in list requests. @OutputOnly
	Kind string `json:"kind,omitempty"`

	// Details of the custom note.
	CustomDetails CustomDetails `json:"customDetails,omitempty"`

	// Details of a security vulnerability note.
	VulnerabilityDetails VulnerabilityDetails `json:"vulnerabilityDetails,omitempty"`

	// Build details for a verifiable build.
	BuildDetails BuildDetails `json:"buildDetails,omitempty"`

	// Describes how this resource derives from the basis in the associated note.
	DerivedImage Derived `json:"derivedImage,omitempty"`

	// Describes the installation of a package on the linked resource.
	Installation Installation `json:"installation,omitempty"`

	// Describes the deployment of an artifact on a runtime.
	Deployment Deployment `json:"deployment,omitempty"`

	// Describes the initial scan status for this resource.
	Discovered Discovered `json:"discovered,omitempty"`

	// A description of actions that can be taken to remedy the note
	Remediation string `json:"remediation,omitempty"`

	// The time this occurrence was created. @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// The time this occurrence was last updated. @OutputOnly
	UpdateTime string `json:"updateTime,omitempty"`

	// The name of the operation that created this note.
	OperationName string `json:"operationName,omitempty"`
}
