// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using DestinationRule within kubernetes types, where deepcopy-gen is used.
func (in *DestinationRule) DeepCopyInto(out *DestinationRule) {
	p := proto.Clone(in).(*DestinationRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRule. Required by controller-gen.
func (in *DestinationRule) DeepCopy() *DestinationRule {
	if in == nil {
		return nil
	}
	out := new(DestinationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRule. Required by controller-gen.
func (in *DestinationRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	p := proto.Clone(in).(*TrafficPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy_PortTrafficPolicy within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopyInto(out *TrafficPolicy_PortTrafficPolicy) {
	p := proto.Clone(in).(*TrafficPolicy_PortTrafficPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_PortTrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopy() *TrafficPolicy_PortTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy_PortTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_PortTrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy_TunnelSettings within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy_TunnelSettings) DeepCopyInto(out *TrafficPolicy_TunnelSettings) {
	p := proto.Clone(in).(*TrafficPolicy_TunnelSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_TunnelSettings. Required by controller-gen.
func (in *TrafficPolicy_TunnelSettings) DeepCopy() *TrafficPolicy_TunnelSettings {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy_TunnelSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_TunnelSettings. Required by controller-gen.
func (in *TrafficPolicy_TunnelSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy_ProxyProtocol within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy_ProxyProtocol) DeepCopyInto(out *TrafficPolicy_ProxyProtocol) {
	p := proto.Clone(in).(*TrafficPolicy_ProxyProtocol)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_ProxyProtocol. Required by controller-gen.
func (in *TrafficPolicy_ProxyProtocol) DeepCopy() *TrafficPolicy_ProxyProtocol {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy_ProxyProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_ProxyProtocol. Required by controller-gen.
func (in *TrafficPolicy_ProxyProtocol) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Subset within kubernetes types, where deepcopy-gen is used.
func (in *Subset) DeepCopyInto(out *Subset) {
	p := proto.Clone(in).(*Subset)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset. Required by controller-gen.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Subset. Required by controller-gen.
func (in *Subset) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings) DeepCopyInto(out *LoadBalancerSettings) {
	p := proto.Clone(in).(*LoadBalancerSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings. Required by controller-gen.
func (in *LoadBalancerSettings) DeepCopy() *LoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings. Required by controller-gen.
func (in *LoadBalancerSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopy() *LoadBalancerSettings_ConsistentHashLB {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB_RingHash within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB_RingHash) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB_RingHash) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB_RingHash)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_RingHash. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_RingHash) DeepCopy() *LoadBalancerSettings_ConsistentHashLB_RingHash {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB_RingHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_RingHash. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_RingHash) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB_MagLev within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB_MagLev) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB_MagLev) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB_MagLev)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_MagLev. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_MagLev) DeepCopy() *LoadBalancerSettings_ConsistentHashLB_MagLev {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB_MagLev)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_MagLev. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_MagLev) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB_HTTPCookie within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB_HTTPCookie)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_HTTPCookie. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopy() *LoadBalancerSettings_ConsistentHashLB_HTTPCookie {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB_HTTPCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_HTTPCookie. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_DynamicSubsetLB within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_DynamicSubsetLB) DeepCopyInto(out *LoadBalancerSettings_DynamicSubsetLB) {
	p := proto.Clone(in).(*LoadBalancerSettings_DynamicSubsetLB)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_DynamicSubsetLB. Required by controller-gen.
func (in *LoadBalancerSettings_DynamicSubsetLB) DeepCopy() *LoadBalancerSettings_DynamicSubsetLB {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_DynamicSubsetLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_DynamicSubsetLB. Required by controller-gen.
func (in *LoadBalancerSettings_DynamicSubsetLB) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_DynamicSubsetLB_SubsetSelector within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_DynamicSubsetLB_SubsetSelector) DeepCopyInto(out *LoadBalancerSettings_DynamicSubsetLB_SubsetSelector) {
	p := proto.Clone(in).(*LoadBalancerSettings_DynamicSubsetLB_SubsetSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_DynamicSubsetLB_SubsetSelector. Required by controller-gen.
func (in *LoadBalancerSettings_DynamicSubsetLB_SubsetSelector) DeepCopy() *LoadBalancerSettings_DynamicSubsetLB_SubsetSelector {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_DynamicSubsetLB_SubsetSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_DynamicSubsetLB_SubsetSelector. Required by controller-gen.
func (in *LoadBalancerSettings_DynamicSubsetLB_SubsetSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings) DeepCopyInto(out *ConnectionPoolSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopy() *ConnectionPoolSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_TCPSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopyInto(out *ConnectionPoolSettings_TCPSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings_TCPSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopy() *ConnectionPoolSettings_TCPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_TCPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_TCPSettings_TcpKeepalive within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopyInto(out *ConnectionPoolSettings_TCPSettings_TcpKeepalive) {
	p := proto.Clone(in).(*ConnectionPoolSettings_TCPSettings_TcpKeepalive)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings_TcpKeepalive. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopy() *ConnectionPoolSettings_TCPSettings_TcpKeepalive {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_TCPSettings_TcpKeepalive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings_TcpKeepalive. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_HTTPSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopyInto(out *ConnectionPoolSettings_HTTPSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings_HTTPSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_HTTPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopy() *ConnectionPoolSettings_HTTPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_HTTPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_HTTPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using OutlierDetection within kubernetes types, where deepcopy-gen is used.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	p := proto.Clone(in).(*OutlierDetection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection. Required by controller-gen.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection. Required by controller-gen.
func (in *OutlierDetection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ClientTLSSettings within kubernetes types, where deepcopy-gen is used.
func (in *ClientTLSSettings) DeepCopyInto(out *ClientTLSSettings) {
	p := proto.Clone(in).(*ClientTLSSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientTLSSettings. Required by controller-gen.
func (in *ClientTLSSettings) DeepCopy() *ClientTLSSettings {
	if in == nil {
		return nil
	}
	out := new(ClientTLSSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ClientTLSSettings. Required by controller-gen.
func (in *ClientTLSSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting) DeepCopyInto(out *LocalityLoadBalancerSetting) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting. Required by controller-gen.
func (in *LocalityLoadBalancerSetting) DeepCopy() *LocalityLoadBalancerSetting {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting. Required by controller-gen.
func (in *LocalityLoadBalancerSetting) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting_Distribute within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopyInto(out *LocalityLoadBalancerSetting_Distribute) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting_Distribute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Distribute. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopy() *LocalityLoadBalancerSetting_Distribute {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting_Distribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Distribute. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting_Failover within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopyInto(out *LocalityLoadBalancerSetting_Failover) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting_Failover)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Failover. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopy() *LocalityLoadBalancerSetting_Failover {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting_Failover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Failover. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
