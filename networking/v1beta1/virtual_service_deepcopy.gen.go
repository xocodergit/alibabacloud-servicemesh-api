// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using VirtualService within kubernetes types, where deepcopy-gen is used.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	p := proto.Clone(in).(*VirtualService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService. Required by controller-gen.
func (in *VirtualService) DeepCopy() *VirtualService {
	if in == nil {
		return nil
	}
	out := new(VirtualService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService. Required by controller-gen.
func (in *VirtualService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Destination within kubernetes types, where deepcopy-gen is used.
func (in *Destination) DeepCopyInto(out *Destination) {
	p := proto.Clone(in).(*Destination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination. Required by controller-gen.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Destination. Required by controller-gen.
func (in *Destination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRoute within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRoute) DeepCopyInto(out *HTTPRoute) {
	p := proto.Clone(in).(*HTTPRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute. Required by controller-gen.
func (in *HTTPRoute) DeepCopy() *HTTPRoute {
	if in == nil {
		return nil
	}
	out := new(HTTPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute. Required by controller-gen.
func (in *HTTPRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Delegate within kubernetes types, where deepcopy-gen is used.
func (in *Delegate) DeepCopyInto(out *Delegate) {
	p := proto.Clone(in).(*Delegate)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delegate. Required by controller-gen.
func (in *Delegate) DeepCopy() *Delegate {
	if in == nil {
		return nil
	}
	out := new(Delegate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Delegate. Required by controller-gen.
func (in *Delegate) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Headers within kubernetes types, where deepcopy-gen is used.
func (in *Headers) DeepCopyInto(out *Headers) {
	p := proto.Clone(in).(*Headers)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Headers. Required by controller-gen.
func (in *Headers) DeepCopy() *Headers {
	if in == nil {
		return nil
	}
	out := new(Headers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Headers. Required by controller-gen.
func (in *Headers) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Headers_HeaderOperations within kubernetes types, where deepcopy-gen is used.
func (in *Headers_HeaderOperations) DeepCopyInto(out *Headers_HeaderOperations) {
	p := proto.Clone(in).(*Headers_HeaderOperations)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Headers_HeaderOperations. Required by controller-gen.
func (in *Headers_HeaderOperations) DeepCopy() *Headers_HeaderOperations {
	if in == nil {
		return nil
	}
	out := new(Headers_HeaderOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Headers_HeaderOperations. Required by controller-gen.
func (in *Headers_HeaderOperations) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TLSRoute within kubernetes types, where deepcopy-gen is used.
func (in *TLSRoute) DeepCopyInto(out *TLSRoute) {
	p := proto.Clone(in).(*TLSRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSRoute. Required by controller-gen.
func (in *TLSRoute) DeepCopy() *TLSRoute {
	if in == nil {
		return nil
	}
	out := new(TLSRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TLSRoute. Required by controller-gen.
func (in *TLSRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCPRoute within kubernetes types, where deepcopy-gen is used.
func (in *TCPRoute) DeepCopyInto(out *TCPRoute) {
	p := proto.Clone(in).(*TCPRoute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPRoute. Required by controller-gen.
func (in *TCPRoute) DeepCopy() *TCPRoute {
	if in == nil {
		return nil
	}
	out := new(TCPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCPRoute. Required by controller-gen.
func (in *TCPRoute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPMatchRequest within kubernetes types, where deepcopy-gen is used.
func (in *HTTPMatchRequest) DeepCopyInto(out *HTTPMatchRequest) {
	p := proto.Clone(in).(*HTTPMatchRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest. Required by controller-gen.
func (in *HTTPMatchRequest) DeepCopy() *HTTPMatchRequest {
	if in == nil {
		return nil
	}
	out := new(HTTPMatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest. Required by controller-gen.
func (in *HTTPMatchRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRouteDestination within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRouteDestination) DeepCopyInto(out *HTTPRouteDestination) {
	p := proto.Clone(in).(*HTTPRouteDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteDestination. Required by controller-gen.
func (in *HTTPRouteDestination) DeepCopy() *HTTPRouteDestination {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteDestination. Required by controller-gen.
func (in *HTTPRouteDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRouteFallback within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRouteFallback) DeepCopyInto(out *HTTPRouteFallback) {
	p := proto.Clone(in).(*HTTPRouteFallback)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteFallback. Required by controller-gen.
func (in *HTTPRouteFallback) DeepCopy() *HTTPRouteFallback {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteFallback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteFallback. Required by controller-gen.
func (in *HTTPRouteFallback) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteDestination within kubernetes types, where deepcopy-gen is used.
func (in *RouteDestination) DeepCopyInto(out *RouteDestination) {
	p := proto.Clone(in).(*RouteDestination)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopy() *RouteDestination {
	if in == nil {
		return nil
	}
	out := new(RouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteDestination. Required by controller-gen.
func (in *RouteDestination) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using L4MatchAttributes within kubernetes types, where deepcopy-gen is used.
func (in *L4MatchAttributes) DeepCopyInto(out *L4MatchAttributes) {
	p := proto.Clone(in).(*L4MatchAttributes)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4MatchAttributes. Required by controller-gen.
func (in *L4MatchAttributes) DeepCopy() *L4MatchAttributes {
	if in == nil {
		return nil
	}
	out := new(L4MatchAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new L4MatchAttributes. Required by controller-gen.
func (in *L4MatchAttributes) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TLSMatchAttributes within kubernetes types, where deepcopy-gen is used.
func (in *TLSMatchAttributes) DeepCopyInto(out *TLSMatchAttributes) {
	p := proto.Clone(in).(*TLSMatchAttributes)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSMatchAttributes. Required by controller-gen.
func (in *TLSMatchAttributes) DeepCopy() *TLSMatchAttributes {
	if in == nil {
		return nil
	}
	out := new(TLSMatchAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TLSMatchAttributes. Required by controller-gen.
func (in *TLSMatchAttributes) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRedirect within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRedirect) DeepCopyInto(out *HTTPRedirect) {
	p := proto.Clone(in).(*HTTPRedirect)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirect. Required by controller-gen.
func (in *HTTPRedirect) DeepCopy() *HTTPRedirect {
	if in == nil {
		return nil
	}
	out := new(HTTPRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirect. Required by controller-gen.
func (in *HTTPRedirect) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPDirectResponse within kubernetes types, where deepcopy-gen is used.
func (in *HTTPDirectResponse) DeepCopyInto(out *HTTPDirectResponse) {
	p := proto.Clone(in).(*HTTPDirectResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPDirectResponse. Required by controller-gen.
func (in *HTTPDirectResponse) DeepCopy() *HTTPDirectResponse {
	if in == nil {
		return nil
	}
	out := new(HTTPDirectResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPDirectResponse. Required by controller-gen.
func (in *HTTPDirectResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPBody within kubernetes types, where deepcopy-gen is used.
func (in *HTTPBody) DeepCopyInto(out *HTTPBody) {
	p := proto.Clone(in).(*HTTPBody)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBody. Required by controller-gen.
func (in *HTTPBody) DeepCopy() *HTTPBody {
	if in == nil {
		return nil
	}
	out := new(HTTPBody)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBody. Required by controller-gen.
func (in *HTTPBody) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRewrite within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRewrite) DeepCopyInto(out *HTTPRewrite) {
	p := proto.Clone(in).(*HTTPRewrite)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRewrite. Required by controller-gen.
func (in *HTTPRewrite) DeepCopy() *HTTPRewrite {
	if in == nil {
		return nil
	}
	out := new(HTTPRewrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRewrite. Required by controller-gen.
func (in *HTTPRewrite) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using StringMatch within kubernetes types, where deepcopy-gen is used.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	p := proto.Clone(in).(*StringMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch. Required by controller-gen.
func (in *StringMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPRetry within kubernetes types, where deepcopy-gen is used.
func (in *HTTPRetry) DeepCopyInto(out *HTTPRetry) {
	p := proto.Clone(in).(*HTTPRetry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry. Required by controller-gen.
func (in *HTTPRetry) DeepCopy() *HTTPRetry {
	if in == nil {
		return nil
	}
	out := new(HTTPRetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry. Required by controller-gen.
func (in *HTTPRetry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CorsPolicy within kubernetes types, where deepcopy-gen is used.
func (in *CorsPolicy) DeepCopyInto(out *CorsPolicy) {
	p := proto.Clone(in).(*CorsPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy. Required by controller-gen.
func (in *CorsPolicy) DeepCopy() *CorsPolicy {
	if in == nil {
		return nil
	}
	out := new(CorsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy. Required by controller-gen.
func (in *CorsPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection) DeepCopyInto(out *HTTPFaultInjection) {
	p := proto.Clone(in).(*HTTPFaultInjection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopy() *HTTPFaultInjection {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection. Required by controller-gen.
func (in *HTTPFaultInjection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Delay within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Delay) DeepCopyInto(out *HTTPFaultInjection_Delay) {
	p := proto.Clone(in).(*HTTPFaultInjection_Delay)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopy() *HTTPFaultInjection_Delay {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Delay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Delay. Required by controller-gen.
func (in *HTTPFaultInjection_Delay) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPFaultInjection_Abort within kubernetes types, where deepcopy-gen is used.
func (in *HTTPFaultInjection_Abort) DeepCopyInto(out *HTTPFaultInjection_Abort) {
	p := proto.Clone(in).(*HTTPFaultInjection_Abort)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopy() *HTTPFaultInjection_Abort {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection_Abort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection_Abort. Required by controller-gen.
func (in *HTTPFaultInjection_Abort) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PortSelector within kubernetes types, where deepcopy-gen is used.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	p := proto.Clone(in).(*PortSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Percent within kubernetes types, where deepcopy-gen is used.
func (in *Percent) DeepCopyInto(out *Percent) {
	p := proto.Clone(in).(*Percent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percent. Required by controller-gen.
func (in *Percent) DeepCopy() *Percent {
	if in == nil {
		return nil
	}
	out := new(Percent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Percent. Required by controller-gen.
func (in *Percent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
