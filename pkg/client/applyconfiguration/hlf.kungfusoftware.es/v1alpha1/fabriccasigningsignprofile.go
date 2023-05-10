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

// FabricCASigningSignProfileApplyConfiguration represents an declarative configuration of the FabricCASigningSignProfile type for use
// with apply.
type FabricCASigningSignProfileApplyConfiguration struct {
	Usage        []string                                                `json:"usage,omitempty"`
	Expiry       *string                                                 `json:"expiry,omitempty"`
	CAConstraint *FabricCASigningSignProfileConstraintApplyConfiguration `json:"caconstraint,omitempty"`
}

// FabricCASigningSignProfileApplyConfiguration constructs an declarative configuration of the FabricCASigningSignProfile type for use with
// apply.
func FabricCASigningSignProfile() *FabricCASigningSignProfileApplyConfiguration {
	return &FabricCASigningSignProfileApplyConfiguration{}
}

// WithUsage adds the given value to the Usage field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Usage field.
func (b *FabricCASigningSignProfileApplyConfiguration) WithUsage(values ...string) *FabricCASigningSignProfileApplyConfiguration {
	for i := range values {
		b.Usage = append(b.Usage, values[i])
	}
	return b
}

// WithExpiry sets the Expiry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expiry field is set to the value of the last call.
func (b *FabricCASigningSignProfileApplyConfiguration) WithExpiry(value string) *FabricCASigningSignProfileApplyConfiguration {
	b.Expiry = &value
	return b
}

// WithCAConstraint sets the CAConstraint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CAConstraint field is set to the value of the last call.
func (b *FabricCASigningSignProfileApplyConfiguration) WithCAConstraint(value *FabricCASigningSignProfileConstraintApplyConfiguration) *FabricCASigningSignProfileApplyConfiguration {
	b.CAConstraint = value
	return b
}
