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

// ServiceMonitorApplyConfiguration represents an declarative configuration of the ServiceMonitor type for use
// with apply.
type ServiceMonitorApplyConfiguration struct {
	Enabled       *bool             `json:"enabled,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"`
	SampleLimit   *int              `json:"sampleLimit,omitempty"`
	Interval      *string           `json:"interval,omitempty"`
	ScrapeTimeout *string           `json:"scrapeTimeout,omitempty"`
}

// ServiceMonitorApplyConfiguration constructs an declarative configuration of the ServiceMonitor type for use with
// apply.
func ServiceMonitor() *ServiceMonitorApplyConfiguration {
	return &ServiceMonitorApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *ServiceMonitorApplyConfiguration) WithEnabled(value bool) *ServiceMonitorApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ServiceMonitorApplyConfiguration) WithLabels(entries map[string]string) *ServiceMonitorApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithSampleLimit sets the SampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleLimit field is set to the value of the last call.
func (b *ServiceMonitorApplyConfiguration) WithSampleLimit(value int) *ServiceMonitorApplyConfiguration {
	b.SampleLimit = &value
	return b
}

// WithInterval sets the Interval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Interval field is set to the value of the last call.
func (b *ServiceMonitorApplyConfiguration) WithInterval(value string) *ServiceMonitorApplyConfiguration {
	b.Interval = &value
	return b
}

// WithScrapeTimeout sets the ScrapeTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeTimeout field is set to the value of the last call.
func (b *ServiceMonitorApplyConfiguration) WithScrapeTimeout(value string) *ServiceMonitorApplyConfiguration {
	b.ScrapeTimeout = &value
	return b
}
