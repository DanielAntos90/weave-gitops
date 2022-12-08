// Code generated by counterfeiter. DO NOT EDIT.
package clustersmngrfakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/core/clustersmngr"
	"github.com/weaveworks/weave-gitops/core/clustersmngr/cluster"
	"github.com/weaveworks/weave-gitops/pkg/server/auth"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/discovery"
)

type FakeClustersManager struct {
	GetClustersStub        func() []cluster.Cluster
	getClustersMutex       sync.RWMutex
	getClustersArgsForCall []struct {
	}
	getClustersReturns struct {
		result1 []cluster.Cluster
	}
	getClustersReturnsOnCall map[int]struct {
		result1 []cluster.Cluster
	}
	GetClustersNamespacesStub        func() map[string][]v1.Namespace
	getClustersNamespacesMutex       sync.RWMutex
	getClustersNamespacesArgsForCall []struct {
	}
	getClustersNamespacesReturns struct {
		result1 map[string][]v1.Namespace
	}
	getClustersNamespacesReturnsOnCall map[int]struct {
		result1 map[string][]v1.Namespace
	}
	GetImpersonatedClientStub        func(context.Context, *auth.UserPrincipal) (clustersmngr.Client, error)
	getImpersonatedClientMutex       sync.RWMutex
	getImpersonatedClientArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}
	getImpersonatedClientReturns struct {
		result1 clustersmngr.Client
		result2 error
	}
	getImpersonatedClientReturnsOnCall map[int]struct {
		result1 clustersmngr.Client
		result2 error
	}
	GetImpersonatedClientForClusterStub        func(context.Context, *auth.UserPrincipal, string) (clustersmngr.Client, error)
	getImpersonatedClientForClusterMutex       sync.RWMutex
	getImpersonatedClientForClusterArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
		arg3 string
	}
	getImpersonatedClientForClusterReturns struct {
		result1 clustersmngr.Client
		result2 error
	}
	getImpersonatedClientForClusterReturnsOnCall map[int]struct {
		result1 clustersmngr.Client
		result2 error
	}
	GetImpersonatedDiscoveryClientStub        func(context.Context, *auth.UserPrincipal, string) (discovery.DiscoveryInterface, error)
	getImpersonatedDiscoveryClientMutex       sync.RWMutex
	getImpersonatedDiscoveryClientArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
		arg3 string
	}
	getImpersonatedDiscoveryClientReturns struct {
		result1 discovery.DiscoveryInterface
		result2 error
	}
	getImpersonatedDiscoveryClientReturnsOnCall map[int]struct {
		result1 discovery.DiscoveryInterface
		result2 error
	}
	GetOIDCClientStub        func(context.Context, *auth.UserPrincipal) (clustersmngr.Client, error)
	getOIDCClientMutex       sync.RWMutex
	getOIDCClientArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}
	getOIDCClientReturns struct {
		result1 clustersmngr.Client
		result2 error
	}
	getOIDCClientReturnsOnCall map[int]struct {
		result1 clustersmngr.Client
		result2 error
	}
	GetServerClientStub        func(context.Context) (clustersmngr.Client, error)
	getServerClientMutex       sync.RWMutex
	getServerClientArgsForCall []struct {
		arg1 context.Context
	}
	getServerClientReturns struct {
		result1 clustersmngr.Client
		result2 error
	}
	getServerClientReturnsOnCall map[int]struct {
		result1 clustersmngr.Client
		result2 error
	}
	GetUserNamespacesStub        func(*auth.UserPrincipal) map[string][]v1.Namespace
	getUserNamespacesMutex       sync.RWMutex
	getUserNamespacesArgsForCall []struct {
		arg1 *auth.UserPrincipal
	}
	getUserNamespacesReturns struct {
		result1 map[string][]v1.Namespace
	}
	getUserNamespacesReturnsOnCall map[int]struct {
		result1 map[string][]v1.Namespace
	}
	RemoveWatcherStub        func(*clustersmngr.ClustersWatcher)
	removeWatcherMutex       sync.RWMutex
	removeWatcherArgsForCall []struct {
		arg1 *clustersmngr.ClustersWatcher
	}
	StartStub        func(context.Context)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	SubscribeStub        func() *clustersmngr.ClustersWatcher
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
	}
	subscribeReturns struct {
		result1 *clustersmngr.ClustersWatcher
	}
	subscribeReturnsOnCall map[int]struct {
		result1 *clustersmngr.ClustersWatcher
	}
	UpdateClustersStub        func(context.Context) error
	updateClustersMutex       sync.RWMutex
	updateClustersArgsForCall []struct {
		arg1 context.Context
	}
	updateClustersReturns struct {
		result1 error
	}
	updateClustersReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateNamespacesStub        func(context.Context) error
	updateNamespacesMutex       sync.RWMutex
	updateNamespacesArgsForCall []struct {
		arg1 context.Context
	}
	updateNamespacesReturns struct {
		result1 error
	}
	updateNamespacesReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateUserNamespacesStub        func(context.Context, *auth.UserPrincipal)
	updateUserNamespacesMutex       sync.RWMutex
	updateUserNamespacesArgsForCall []struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClustersManager) GetClusters() []cluster.Cluster {
	fake.getClustersMutex.Lock()
	ret, specificReturn := fake.getClustersReturnsOnCall[len(fake.getClustersArgsForCall)]
	fake.getClustersArgsForCall = append(fake.getClustersArgsForCall, struct {
	}{})
	stub := fake.GetClustersStub
	fakeReturns := fake.getClustersReturns
	fake.recordInvocation("GetClusters", []interface{}{})
	fake.getClustersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) GetClustersCallCount() int {
	fake.getClustersMutex.RLock()
	defer fake.getClustersMutex.RUnlock()
	return len(fake.getClustersArgsForCall)
}

func (fake *FakeClustersManager) GetClustersCalls(stub func() []cluster.Cluster) {
	fake.getClustersMutex.Lock()
	defer fake.getClustersMutex.Unlock()
	fake.GetClustersStub = stub
}

func (fake *FakeClustersManager) GetClustersReturns(result1 []cluster.Cluster) {
	fake.getClustersMutex.Lock()
	defer fake.getClustersMutex.Unlock()
	fake.GetClustersStub = nil
	fake.getClustersReturns = struct {
		result1 []cluster.Cluster
	}{result1}
}

func (fake *FakeClustersManager) GetClustersReturnsOnCall(i int, result1 []cluster.Cluster) {
	fake.getClustersMutex.Lock()
	defer fake.getClustersMutex.Unlock()
	fake.GetClustersStub = nil
	if fake.getClustersReturnsOnCall == nil {
		fake.getClustersReturnsOnCall = make(map[int]struct {
			result1 []cluster.Cluster
		})
	}
	fake.getClustersReturnsOnCall[i] = struct {
		result1 []cluster.Cluster
	}{result1}
}

func (fake *FakeClustersManager) GetClustersNamespaces() map[string][]v1.Namespace {
	fake.getClustersNamespacesMutex.Lock()
	ret, specificReturn := fake.getClustersNamespacesReturnsOnCall[len(fake.getClustersNamespacesArgsForCall)]
	fake.getClustersNamespacesArgsForCall = append(fake.getClustersNamespacesArgsForCall, struct {
	}{})
	stub := fake.GetClustersNamespacesStub
	fakeReturns := fake.getClustersNamespacesReturns
	fake.recordInvocation("GetClustersNamespaces", []interface{}{})
	fake.getClustersNamespacesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) GetClustersNamespacesCallCount() int {
	fake.getClustersNamespacesMutex.RLock()
	defer fake.getClustersNamespacesMutex.RUnlock()
	return len(fake.getClustersNamespacesArgsForCall)
}

func (fake *FakeClustersManager) GetClustersNamespacesCalls(stub func() map[string][]v1.Namespace) {
	fake.getClustersNamespacesMutex.Lock()
	defer fake.getClustersNamespacesMutex.Unlock()
	fake.GetClustersNamespacesStub = stub
}

func (fake *FakeClustersManager) GetClustersNamespacesReturns(result1 map[string][]v1.Namespace) {
	fake.getClustersNamespacesMutex.Lock()
	defer fake.getClustersNamespacesMutex.Unlock()
	fake.GetClustersNamespacesStub = nil
	fake.getClustersNamespacesReturns = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeClustersManager) GetClustersNamespacesReturnsOnCall(i int, result1 map[string][]v1.Namespace) {
	fake.getClustersNamespacesMutex.Lock()
	defer fake.getClustersNamespacesMutex.Unlock()
	fake.GetClustersNamespacesStub = nil
	if fake.getClustersNamespacesReturnsOnCall == nil {
		fake.getClustersNamespacesReturnsOnCall = make(map[int]struct {
			result1 map[string][]v1.Namespace
		})
	}
	fake.getClustersNamespacesReturnsOnCall[i] = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeClustersManager) GetImpersonatedClient(arg1 context.Context, arg2 *auth.UserPrincipal) (clustersmngr.Client, error) {
	fake.getImpersonatedClientMutex.Lock()
	ret, specificReturn := fake.getImpersonatedClientReturnsOnCall[len(fake.getImpersonatedClientArgsForCall)]
	fake.getImpersonatedClientArgsForCall = append(fake.getImpersonatedClientArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}{arg1, arg2})
	stub := fake.GetImpersonatedClientStub
	fakeReturns := fake.getImpersonatedClientReturns
	fake.recordInvocation("GetImpersonatedClient", []interface{}{arg1, arg2})
	fake.getImpersonatedClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClustersManager) GetImpersonatedClientCallCount() int {
	fake.getImpersonatedClientMutex.RLock()
	defer fake.getImpersonatedClientMutex.RUnlock()
	return len(fake.getImpersonatedClientArgsForCall)
}

func (fake *FakeClustersManager) GetImpersonatedClientCalls(stub func(context.Context, *auth.UserPrincipal) (clustersmngr.Client, error)) {
	fake.getImpersonatedClientMutex.Lock()
	defer fake.getImpersonatedClientMutex.Unlock()
	fake.GetImpersonatedClientStub = stub
}

func (fake *FakeClustersManager) GetImpersonatedClientArgsForCall(i int) (context.Context, *auth.UserPrincipal) {
	fake.getImpersonatedClientMutex.RLock()
	defer fake.getImpersonatedClientMutex.RUnlock()
	argsForCall := fake.getImpersonatedClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClustersManager) GetImpersonatedClientReturns(result1 clustersmngr.Client, result2 error) {
	fake.getImpersonatedClientMutex.Lock()
	defer fake.getImpersonatedClientMutex.Unlock()
	fake.GetImpersonatedClientStub = nil
	fake.getImpersonatedClientReturns = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetImpersonatedClientReturnsOnCall(i int, result1 clustersmngr.Client, result2 error) {
	fake.getImpersonatedClientMutex.Lock()
	defer fake.getImpersonatedClientMutex.Unlock()
	fake.GetImpersonatedClientStub = nil
	if fake.getImpersonatedClientReturnsOnCall == nil {
		fake.getImpersonatedClientReturnsOnCall = make(map[int]struct {
			result1 clustersmngr.Client
			result2 error
		})
	}
	fake.getImpersonatedClientReturnsOnCall[i] = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetImpersonatedClientForCluster(arg1 context.Context, arg2 *auth.UserPrincipal, arg3 string) (clustersmngr.Client, error) {
	fake.getImpersonatedClientForClusterMutex.Lock()
	ret, specificReturn := fake.getImpersonatedClientForClusterReturnsOnCall[len(fake.getImpersonatedClientForClusterArgsForCall)]
	fake.getImpersonatedClientForClusterArgsForCall = append(fake.getImpersonatedClientForClusterArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetImpersonatedClientForClusterStub
	fakeReturns := fake.getImpersonatedClientForClusterReturns
	fake.recordInvocation("GetImpersonatedClientForCluster", []interface{}{arg1, arg2, arg3})
	fake.getImpersonatedClientForClusterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClustersManager) GetImpersonatedClientForClusterCallCount() int {
	fake.getImpersonatedClientForClusterMutex.RLock()
	defer fake.getImpersonatedClientForClusterMutex.RUnlock()
	return len(fake.getImpersonatedClientForClusterArgsForCall)
}

func (fake *FakeClustersManager) GetImpersonatedClientForClusterCalls(stub func(context.Context, *auth.UserPrincipal, string) (clustersmngr.Client, error)) {
	fake.getImpersonatedClientForClusterMutex.Lock()
	defer fake.getImpersonatedClientForClusterMutex.Unlock()
	fake.GetImpersonatedClientForClusterStub = stub
}

func (fake *FakeClustersManager) GetImpersonatedClientForClusterArgsForCall(i int) (context.Context, *auth.UserPrincipal, string) {
	fake.getImpersonatedClientForClusterMutex.RLock()
	defer fake.getImpersonatedClientForClusterMutex.RUnlock()
	argsForCall := fake.getImpersonatedClientForClusterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClustersManager) GetImpersonatedClientForClusterReturns(result1 clustersmngr.Client, result2 error) {
	fake.getImpersonatedClientForClusterMutex.Lock()
	defer fake.getImpersonatedClientForClusterMutex.Unlock()
	fake.GetImpersonatedClientForClusterStub = nil
	fake.getImpersonatedClientForClusterReturns = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetImpersonatedClientForClusterReturnsOnCall(i int, result1 clustersmngr.Client, result2 error) {
	fake.getImpersonatedClientForClusterMutex.Lock()
	defer fake.getImpersonatedClientForClusterMutex.Unlock()
	fake.GetImpersonatedClientForClusterStub = nil
	if fake.getImpersonatedClientForClusterReturnsOnCall == nil {
		fake.getImpersonatedClientForClusterReturnsOnCall = make(map[int]struct {
			result1 clustersmngr.Client
			result2 error
		})
	}
	fake.getImpersonatedClientForClusterReturnsOnCall[i] = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClient(arg1 context.Context, arg2 *auth.UserPrincipal, arg3 string) (discovery.DiscoveryInterface, error) {
	fake.getImpersonatedDiscoveryClientMutex.Lock()
	ret, specificReturn := fake.getImpersonatedDiscoveryClientReturnsOnCall[len(fake.getImpersonatedDiscoveryClientArgsForCall)]
	fake.getImpersonatedDiscoveryClientArgsForCall = append(fake.getImpersonatedDiscoveryClientArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetImpersonatedDiscoveryClientStub
	fakeReturns := fake.getImpersonatedDiscoveryClientReturns
	fake.recordInvocation("GetImpersonatedDiscoveryClient", []interface{}{arg1, arg2, arg3})
	fake.getImpersonatedDiscoveryClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClientCallCount() int {
	fake.getImpersonatedDiscoveryClientMutex.RLock()
	defer fake.getImpersonatedDiscoveryClientMutex.RUnlock()
	return len(fake.getImpersonatedDiscoveryClientArgsForCall)
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClientCalls(stub func(context.Context, *auth.UserPrincipal, string) (discovery.DiscoveryInterface, error)) {
	fake.getImpersonatedDiscoveryClientMutex.Lock()
	defer fake.getImpersonatedDiscoveryClientMutex.Unlock()
	fake.GetImpersonatedDiscoveryClientStub = stub
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClientArgsForCall(i int) (context.Context, *auth.UserPrincipal, string) {
	fake.getImpersonatedDiscoveryClientMutex.RLock()
	defer fake.getImpersonatedDiscoveryClientMutex.RUnlock()
	argsForCall := fake.getImpersonatedDiscoveryClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClientReturns(result1 discovery.DiscoveryInterface, result2 error) {
	fake.getImpersonatedDiscoveryClientMutex.Lock()
	defer fake.getImpersonatedDiscoveryClientMutex.Unlock()
	fake.GetImpersonatedDiscoveryClientStub = nil
	fake.getImpersonatedDiscoveryClientReturns = struct {
		result1 discovery.DiscoveryInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetImpersonatedDiscoveryClientReturnsOnCall(i int, result1 discovery.DiscoveryInterface, result2 error) {
	fake.getImpersonatedDiscoveryClientMutex.Lock()
	defer fake.getImpersonatedDiscoveryClientMutex.Unlock()
	fake.GetImpersonatedDiscoveryClientStub = nil
	if fake.getImpersonatedDiscoveryClientReturnsOnCall == nil {
		fake.getImpersonatedDiscoveryClientReturnsOnCall = make(map[int]struct {
			result1 discovery.DiscoveryInterface
			result2 error
		})
	}
	fake.getImpersonatedDiscoveryClientReturnsOnCall[i] = struct {
		result1 discovery.DiscoveryInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetOIDCClient(arg1 context.Context, arg2 *auth.UserPrincipal) (clustersmngr.Client, error) {
	fake.getOIDCClientMutex.Lock()
	ret, specificReturn := fake.getOIDCClientReturnsOnCall[len(fake.getOIDCClientArgsForCall)]
	fake.getOIDCClientArgsForCall = append(fake.getOIDCClientArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}{arg1, arg2})
	stub := fake.GetOIDCClientStub
	fakeReturns := fake.getOIDCClientReturns
	fake.recordInvocation("GetOIDCClient", []interface{}{arg1, arg2})
	fake.getOIDCClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClustersManager) GetOIDCClientCallCount() int {
	fake.getOIDCClientMutex.RLock()
	defer fake.getOIDCClientMutex.RUnlock()
	return len(fake.getOIDCClientArgsForCall)
}

func (fake *FakeClustersManager) GetOIDCClientCalls(stub func(context.Context, *auth.UserPrincipal) (clustersmngr.Client, error)) {
	fake.getOIDCClientMutex.Lock()
	defer fake.getOIDCClientMutex.Unlock()
	fake.GetOIDCClientStub = stub
}

func (fake *FakeClustersManager) GetOIDCClientArgsForCall(i int) (context.Context, *auth.UserPrincipal) {
	fake.getOIDCClientMutex.RLock()
	defer fake.getOIDCClientMutex.RUnlock()
	argsForCall := fake.getOIDCClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClustersManager) GetOIDCClientReturns(result1 clustersmngr.Client, result2 error) {
	fake.getOIDCClientMutex.Lock()
	defer fake.getOIDCClientMutex.Unlock()
	fake.GetOIDCClientStub = nil
	fake.getOIDCClientReturns = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetOIDCClientReturnsOnCall(i int, result1 clustersmngr.Client, result2 error) {
	fake.getOIDCClientMutex.Lock()
	defer fake.getOIDCClientMutex.Unlock()
	fake.GetOIDCClientStub = nil
	if fake.getOIDCClientReturnsOnCall == nil {
		fake.getOIDCClientReturnsOnCall = make(map[int]struct {
			result1 clustersmngr.Client
			result2 error
		})
	}
	fake.getOIDCClientReturnsOnCall[i] = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetServerClient(arg1 context.Context) (clustersmngr.Client, error) {
	fake.getServerClientMutex.Lock()
	ret, specificReturn := fake.getServerClientReturnsOnCall[len(fake.getServerClientArgsForCall)]
	fake.getServerClientArgsForCall = append(fake.getServerClientArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetServerClientStub
	fakeReturns := fake.getServerClientReturns
	fake.recordInvocation("GetServerClient", []interface{}{arg1})
	fake.getServerClientMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClustersManager) GetServerClientCallCount() int {
	fake.getServerClientMutex.RLock()
	defer fake.getServerClientMutex.RUnlock()
	return len(fake.getServerClientArgsForCall)
}

func (fake *FakeClustersManager) GetServerClientCalls(stub func(context.Context) (clustersmngr.Client, error)) {
	fake.getServerClientMutex.Lock()
	defer fake.getServerClientMutex.Unlock()
	fake.GetServerClientStub = stub
}

func (fake *FakeClustersManager) GetServerClientArgsForCall(i int) context.Context {
	fake.getServerClientMutex.RLock()
	defer fake.getServerClientMutex.RUnlock()
	argsForCall := fake.getServerClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) GetServerClientReturns(result1 clustersmngr.Client, result2 error) {
	fake.getServerClientMutex.Lock()
	defer fake.getServerClientMutex.Unlock()
	fake.GetServerClientStub = nil
	fake.getServerClientReturns = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetServerClientReturnsOnCall(i int, result1 clustersmngr.Client, result2 error) {
	fake.getServerClientMutex.Lock()
	defer fake.getServerClientMutex.Unlock()
	fake.GetServerClientStub = nil
	if fake.getServerClientReturnsOnCall == nil {
		fake.getServerClientReturnsOnCall = make(map[int]struct {
			result1 clustersmngr.Client
			result2 error
		})
	}
	fake.getServerClientReturnsOnCall[i] = struct {
		result1 clustersmngr.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeClustersManager) GetUserNamespaces(arg1 *auth.UserPrincipal) map[string][]v1.Namespace {
	fake.getUserNamespacesMutex.Lock()
	ret, specificReturn := fake.getUserNamespacesReturnsOnCall[len(fake.getUserNamespacesArgsForCall)]
	fake.getUserNamespacesArgsForCall = append(fake.getUserNamespacesArgsForCall, struct {
		arg1 *auth.UserPrincipal
	}{arg1})
	stub := fake.GetUserNamespacesStub
	fakeReturns := fake.getUserNamespacesReturns
	fake.recordInvocation("GetUserNamespaces", []interface{}{arg1})
	fake.getUserNamespacesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) GetUserNamespacesCallCount() int {
	fake.getUserNamespacesMutex.RLock()
	defer fake.getUserNamespacesMutex.RUnlock()
	return len(fake.getUserNamespacesArgsForCall)
}

func (fake *FakeClustersManager) GetUserNamespacesCalls(stub func(*auth.UserPrincipal) map[string][]v1.Namespace) {
	fake.getUserNamespacesMutex.Lock()
	defer fake.getUserNamespacesMutex.Unlock()
	fake.GetUserNamespacesStub = stub
}

func (fake *FakeClustersManager) GetUserNamespacesArgsForCall(i int) *auth.UserPrincipal {
	fake.getUserNamespacesMutex.RLock()
	defer fake.getUserNamespacesMutex.RUnlock()
	argsForCall := fake.getUserNamespacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) GetUserNamespacesReturns(result1 map[string][]v1.Namespace) {
	fake.getUserNamespacesMutex.Lock()
	defer fake.getUserNamespacesMutex.Unlock()
	fake.GetUserNamespacesStub = nil
	fake.getUserNamespacesReturns = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeClustersManager) GetUserNamespacesReturnsOnCall(i int, result1 map[string][]v1.Namespace) {
	fake.getUserNamespacesMutex.Lock()
	defer fake.getUserNamespacesMutex.Unlock()
	fake.GetUserNamespacesStub = nil
	if fake.getUserNamespacesReturnsOnCall == nil {
		fake.getUserNamespacesReturnsOnCall = make(map[int]struct {
			result1 map[string][]v1.Namespace
		})
	}
	fake.getUserNamespacesReturnsOnCall[i] = struct {
		result1 map[string][]v1.Namespace
	}{result1}
}

func (fake *FakeClustersManager) RemoveWatcher(arg1 *clustersmngr.ClustersWatcher) {
	fake.removeWatcherMutex.Lock()
	fake.removeWatcherArgsForCall = append(fake.removeWatcherArgsForCall, struct {
		arg1 *clustersmngr.ClustersWatcher
	}{arg1})
	stub := fake.RemoveWatcherStub
	fake.recordInvocation("RemoveWatcher", []interface{}{arg1})
	fake.removeWatcherMutex.Unlock()
	if stub != nil {
		fake.RemoveWatcherStub(arg1)
	}
}

func (fake *FakeClustersManager) RemoveWatcherCallCount() int {
	fake.removeWatcherMutex.RLock()
	defer fake.removeWatcherMutex.RUnlock()
	return len(fake.removeWatcherArgsForCall)
}

func (fake *FakeClustersManager) RemoveWatcherCalls(stub func(*clustersmngr.ClustersWatcher)) {
	fake.removeWatcherMutex.Lock()
	defer fake.removeWatcherMutex.Unlock()
	fake.RemoveWatcherStub = stub
}

func (fake *FakeClustersManager) RemoveWatcherArgsForCall(i int) *clustersmngr.ClustersWatcher {
	fake.removeWatcherMutex.RLock()
	defer fake.removeWatcherMutex.RUnlock()
	argsForCall := fake.removeWatcherArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) Start(arg1 context.Context) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub(arg1)
	}
}

func (fake *FakeClustersManager) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeClustersManager) StartCalls(stub func(context.Context)) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeClustersManager) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) Subscribe() *clustersmngr.ClustersWatcher {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
	}{})
	stub := fake.SubscribeStub
	fakeReturns := fake.subscribeReturns
	fake.recordInvocation("Subscribe", []interface{}{})
	fake.subscribeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeClustersManager) SubscribeCalls(stub func() *clustersmngr.ClustersWatcher) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = stub
}

func (fake *FakeClustersManager) SubscribeReturns(result1 *clustersmngr.ClustersWatcher) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 *clustersmngr.ClustersWatcher
	}{result1}
}

func (fake *FakeClustersManager) SubscribeReturnsOnCall(i int, result1 *clustersmngr.ClustersWatcher) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 *clustersmngr.ClustersWatcher
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 *clustersmngr.ClustersWatcher
	}{result1}
}

func (fake *FakeClustersManager) UpdateClusters(arg1 context.Context) error {
	fake.updateClustersMutex.Lock()
	ret, specificReturn := fake.updateClustersReturnsOnCall[len(fake.updateClustersArgsForCall)]
	fake.updateClustersArgsForCall = append(fake.updateClustersArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.UpdateClustersStub
	fakeReturns := fake.updateClustersReturns
	fake.recordInvocation("UpdateClusters", []interface{}{arg1})
	fake.updateClustersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) UpdateClustersCallCount() int {
	fake.updateClustersMutex.RLock()
	defer fake.updateClustersMutex.RUnlock()
	return len(fake.updateClustersArgsForCall)
}

func (fake *FakeClustersManager) UpdateClustersCalls(stub func(context.Context) error) {
	fake.updateClustersMutex.Lock()
	defer fake.updateClustersMutex.Unlock()
	fake.UpdateClustersStub = stub
}

func (fake *FakeClustersManager) UpdateClustersArgsForCall(i int) context.Context {
	fake.updateClustersMutex.RLock()
	defer fake.updateClustersMutex.RUnlock()
	argsForCall := fake.updateClustersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) UpdateClustersReturns(result1 error) {
	fake.updateClustersMutex.Lock()
	defer fake.updateClustersMutex.Unlock()
	fake.UpdateClustersStub = nil
	fake.updateClustersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClustersManager) UpdateClustersReturnsOnCall(i int, result1 error) {
	fake.updateClustersMutex.Lock()
	defer fake.updateClustersMutex.Unlock()
	fake.UpdateClustersStub = nil
	if fake.updateClustersReturnsOnCall == nil {
		fake.updateClustersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateClustersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClustersManager) UpdateNamespaces(arg1 context.Context) error {
	fake.updateNamespacesMutex.Lock()
	ret, specificReturn := fake.updateNamespacesReturnsOnCall[len(fake.updateNamespacesArgsForCall)]
	fake.updateNamespacesArgsForCall = append(fake.updateNamespacesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.UpdateNamespacesStub
	fakeReturns := fake.updateNamespacesReturns
	fake.recordInvocation("UpdateNamespaces", []interface{}{arg1})
	fake.updateNamespacesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClustersManager) UpdateNamespacesCallCount() int {
	fake.updateNamespacesMutex.RLock()
	defer fake.updateNamespacesMutex.RUnlock()
	return len(fake.updateNamespacesArgsForCall)
}

func (fake *FakeClustersManager) UpdateNamespacesCalls(stub func(context.Context) error) {
	fake.updateNamespacesMutex.Lock()
	defer fake.updateNamespacesMutex.Unlock()
	fake.UpdateNamespacesStub = stub
}

func (fake *FakeClustersManager) UpdateNamespacesArgsForCall(i int) context.Context {
	fake.updateNamespacesMutex.RLock()
	defer fake.updateNamespacesMutex.RUnlock()
	argsForCall := fake.updateNamespacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClustersManager) UpdateNamespacesReturns(result1 error) {
	fake.updateNamespacesMutex.Lock()
	defer fake.updateNamespacesMutex.Unlock()
	fake.UpdateNamespacesStub = nil
	fake.updateNamespacesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClustersManager) UpdateNamespacesReturnsOnCall(i int, result1 error) {
	fake.updateNamespacesMutex.Lock()
	defer fake.updateNamespacesMutex.Unlock()
	fake.UpdateNamespacesStub = nil
	if fake.updateNamespacesReturnsOnCall == nil {
		fake.updateNamespacesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateNamespacesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClustersManager) UpdateUserNamespaces(arg1 context.Context, arg2 *auth.UserPrincipal) {
	fake.updateUserNamespacesMutex.Lock()
	fake.updateUserNamespacesArgsForCall = append(fake.updateUserNamespacesArgsForCall, struct {
		arg1 context.Context
		arg2 *auth.UserPrincipal
	}{arg1, arg2})
	stub := fake.UpdateUserNamespacesStub
	fake.recordInvocation("UpdateUserNamespaces", []interface{}{arg1, arg2})
	fake.updateUserNamespacesMutex.Unlock()
	if stub != nil {
		fake.UpdateUserNamespacesStub(arg1, arg2)
	}
}

func (fake *FakeClustersManager) UpdateUserNamespacesCallCount() int {
	fake.updateUserNamespacesMutex.RLock()
	defer fake.updateUserNamespacesMutex.RUnlock()
	return len(fake.updateUserNamespacesArgsForCall)
}

func (fake *FakeClustersManager) UpdateUserNamespacesCalls(stub func(context.Context, *auth.UserPrincipal)) {
	fake.updateUserNamespacesMutex.Lock()
	defer fake.updateUserNamespacesMutex.Unlock()
	fake.UpdateUserNamespacesStub = stub
}

func (fake *FakeClustersManager) UpdateUserNamespacesArgsForCall(i int) (context.Context, *auth.UserPrincipal) {
	fake.updateUserNamespacesMutex.RLock()
	defer fake.updateUserNamespacesMutex.RUnlock()
	argsForCall := fake.updateUserNamespacesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClustersManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClustersMutex.RLock()
	defer fake.getClustersMutex.RUnlock()
	fake.getClustersNamespacesMutex.RLock()
	defer fake.getClustersNamespacesMutex.RUnlock()
	fake.getImpersonatedClientMutex.RLock()
	defer fake.getImpersonatedClientMutex.RUnlock()
	fake.getImpersonatedClientForClusterMutex.RLock()
	defer fake.getImpersonatedClientForClusterMutex.RUnlock()
	fake.getImpersonatedDiscoveryClientMutex.RLock()
	defer fake.getImpersonatedDiscoveryClientMutex.RUnlock()
	fake.getOIDCClientMutex.RLock()
	defer fake.getOIDCClientMutex.RUnlock()
	fake.getServerClientMutex.RLock()
	defer fake.getServerClientMutex.RUnlock()
	fake.getUserNamespacesMutex.RLock()
	defer fake.getUserNamespacesMutex.RUnlock()
	fake.removeWatcherMutex.RLock()
	defer fake.removeWatcherMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	fake.updateClustersMutex.RLock()
	defer fake.updateClustersMutex.RUnlock()
	fake.updateNamespacesMutex.RLock()
	defer fake.updateNamespacesMutex.RUnlock()
	fake.updateUserNamespacesMutex.RLock()
	defer fake.updateUserNamespacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClustersManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ clustersmngr.ClustersManager = new(FakeClustersManager)
