// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// FabricCACSRCAApplyConfiguration represents an declarative configuration of the FabricCACSRCA type for use
// with apply.
type FabricCACSRCAApplyConfiguration struct {
	Expiry     *string `json:"expiry,omitempty"`
	PathLength *int    `json:"pathLength,omitempty"`
}

// FabricCACSRCAApplyConfiguration constructs an declarative configuration of the FabricCACSRCA type for use with
// apply.
func FabricCACSRCA() *FabricCACSRCAApplyConfiguration {
	return &FabricCACSRCAApplyConfiguration{}
}

// WithExpiry sets the Expiry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expiry field is set to the value of the last call.
func (b *FabricCACSRCAApplyConfiguration) WithExpiry(value string) *FabricCACSRCAApplyConfiguration {
	b.Expiry = &value
	return b
}

// WithPathLength sets the PathLength field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathLength field is set to the value of the last call.
func (b *FabricCACSRCAApplyConfiguration) WithPathLength(value int) *FabricCACSRCAApplyConfiguration {
	b.PathLength = &value
	return b
}
