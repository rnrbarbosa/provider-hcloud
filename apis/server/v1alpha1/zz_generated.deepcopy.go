//go:build !ignore_autogenerated

/*
Copyright 2024 Cloudhippie <info@cloudhippie.de>
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInitParameters) DeepCopyInto(out *NetworkInitParameters) {
	*out = *in
	if in.AliasIps != nil {
		in, out := &in.AliasIps, &out.AliasIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInitParameters.
func (in *NetworkInitParameters) DeepCopy() *NetworkInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.AliasIps != nil {
		in, out := &in.AliasIps, &out.AliasIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.AliasIps != nil {
		in, out := &in.AliasIps, &out.AliasIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicNetInitParameters) DeepCopyInto(out *PublicNetInitParameters) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(float64)
		**out = **in
	}
	if in.IPv4Enabled != nil {
		in, out := &in.IPv4Enabled, &out.IPv4Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(float64)
		**out = **in
	}
	if in.IPv6Enabled != nil {
		in, out := &in.IPv6Enabled, &out.IPv6Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicNetInitParameters.
func (in *PublicNetInitParameters) DeepCopy() *PublicNetInitParameters {
	if in == nil {
		return nil
	}
	out := new(PublicNetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicNetObservation) DeepCopyInto(out *PublicNetObservation) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(float64)
		**out = **in
	}
	if in.IPv4Enabled != nil {
		in, out := &in.IPv4Enabled, &out.IPv4Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(float64)
		**out = **in
	}
	if in.IPv6Enabled != nil {
		in, out := &in.IPv6Enabled, &out.IPv6Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicNetObservation.
func (in *PublicNetObservation) DeepCopy() *PublicNetObservation {
	if in == nil {
		return nil
	}
	out := new(PublicNetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicNetParameters) DeepCopyInto(out *PublicNetParameters) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(float64)
		**out = **in
	}
	if in.IPv4Enabled != nil {
		in, out := &in.IPv4Enabled, &out.IPv4Enabled
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(float64)
		**out = **in
	}
	if in.IPv6Enabled != nil {
		in, out := &in.IPv6Enabled, &out.IPv6Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicNetParameters.
func (in *PublicNetParameters) DeepCopy() *PublicNetParameters {
	if in == nil {
		return nil
	}
	out := new(PublicNetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerInitParameters) DeepCopyInto(out *ServerInitParameters) {
	*out = *in
	if in.AllowDeprecatedImages != nil {
		in, out := &in.AllowDeprecatedImages, &out.AllowDeprecatedImages
		*out = new(bool)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DeleteProtection != nil {
		in, out := &in.DeleteProtection, &out.DeleteProtection
		*out = new(bool)
		**out = **in
	}
	if in.FirewallIds != nil {
		in, out := &in.FirewallIds, &out.FirewallIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.IgnoreRemoteFirewallIds != nil {
		in, out := &in.IgnoreRemoteFirewallIds, &out.IgnoreRemoteFirewallIds
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Iso != nil {
		in, out := &in.Iso, &out.Iso
		*out = new(string)
		**out = **in
	}
	if in.KeepDisk != nil {
		in, out := &in.KeepDisk, &out.KeepDisk
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementGroupID != nil {
		in, out := &in.PlacementGroupID, &out.PlacementGroupID
		*out = new(float64)
		**out = **in
	}
	if in.PublicNet != nil {
		in, out := &in.PublicNet, &out.PublicNet
		*out = make([]PublicNetInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RebuildProtection != nil {
		in, out := &in.RebuildProtection, &out.RebuildProtection
		*out = new(bool)
		**out = **in
	}
	if in.Rescue != nil {
		in, out := &in.Rescue, &out.Rescue
		*out = new(string)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerType != nil {
		in, out := &in.ServerType, &out.ServerType
		*out = new(string)
		**out = **in
	}
	if in.ShutdownBeforeDeletion != nil {
		in, out := &in.ShutdownBeforeDeletion, &out.ShutdownBeforeDeletion
		*out = new(bool)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerInitParameters.
func (in *ServerInitParameters) DeepCopy() *ServerInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerObservation) DeepCopyInto(out *ServerObservation) {
	*out = *in
	if in.AllowDeprecatedImages != nil {
		in, out := &in.AllowDeprecatedImages, &out.AllowDeprecatedImages
		*out = new(bool)
		**out = **in
	}
	if in.BackupWindow != nil {
		in, out := &in.BackupWindow, &out.BackupWindow
		*out = new(string)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DeleteProtection != nil {
		in, out := &in.DeleteProtection, &out.DeleteProtection
		*out = new(bool)
		**out = **in
	}
	if in.FirewallIds != nil {
		in, out := &in.FirewallIds, &out.FirewallIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.IPv6Network != nil {
		in, out := &in.IPv6Network, &out.IPv6Network
		*out = new(string)
		**out = **in
	}
	if in.IgnoreRemoteFirewallIds != nil {
		in, out := &in.IgnoreRemoteFirewallIds, &out.IgnoreRemoteFirewallIds
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Iso != nil {
		in, out := &in.Iso, &out.Iso
		*out = new(string)
		**out = **in
	}
	if in.KeepDisk != nil {
		in, out := &in.KeepDisk, &out.KeepDisk
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementGroupID != nil {
		in, out := &in.PlacementGroupID, &out.PlacementGroupID
		*out = new(float64)
		**out = **in
	}
	if in.PrimaryDiskSize != nil {
		in, out := &in.PrimaryDiskSize, &out.PrimaryDiskSize
		*out = new(float64)
		**out = **in
	}
	if in.PublicNet != nil {
		in, out := &in.PublicNet, &out.PublicNet
		*out = make([]PublicNetObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RebuildProtection != nil {
		in, out := &in.RebuildProtection, &out.RebuildProtection
		*out = new(bool)
		**out = **in
	}
	if in.Rescue != nil {
		in, out := &in.Rescue, &out.Rescue
		*out = new(string)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerType != nil {
		in, out := &in.ServerType, &out.ServerType
		*out = new(string)
		**out = **in
	}
	if in.ShutdownBeforeDeletion != nil {
		in, out := &in.ShutdownBeforeDeletion, &out.ShutdownBeforeDeletion
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerObservation.
func (in *ServerObservation) DeepCopy() *ServerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerParameters) DeepCopyInto(out *ServerParameters) {
	*out = *in
	if in.AllowDeprecatedImages != nil {
		in, out := &in.AllowDeprecatedImages, &out.AllowDeprecatedImages
		*out = new(bool)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DeleteProtection != nil {
		in, out := &in.DeleteProtection, &out.DeleteProtection
		*out = new(bool)
		**out = **in
	}
	if in.FirewallIds != nil {
		in, out := &in.FirewallIds, &out.FirewallIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.IgnoreRemoteFirewallIds != nil {
		in, out := &in.IgnoreRemoteFirewallIds, &out.IgnoreRemoteFirewallIds
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Iso != nil {
		in, out := &in.Iso, &out.Iso
		*out = new(string)
		**out = **in
	}
	if in.KeepDisk != nil {
		in, out := &in.KeepDisk, &out.KeepDisk
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementGroupID != nil {
		in, out := &in.PlacementGroupID, &out.PlacementGroupID
		*out = new(float64)
		**out = **in
	}
	if in.PublicNet != nil {
		in, out := &in.PublicNet, &out.PublicNet
		*out = make([]PublicNetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RebuildProtection != nil {
		in, out := &in.RebuildProtection, &out.RebuildProtection
		*out = new(bool)
		**out = **in
	}
	if in.Rescue != nil {
		in, out := &in.Rescue, &out.Rescue
		*out = new(string)
		**out = **in
	}
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerType != nil {
		in, out := &in.ServerType, &out.ServerType
		*out = new(string)
		**out = **in
	}
	if in.ShutdownBeforeDeletion != nil {
		in, out := &in.ShutdownBeforeDeletion, &out.ShutdownBeforeDeletion
		*out = new(bool)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerParameters.
func (in *ServerParameters) DeepCopy() *ServerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}
