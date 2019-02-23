// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/common/privdata"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type CollectionStore struct {
	RetrieveCollectionStub        func(common.CollectionCriteria) (privdata.Collection, error)
	retrieveCollectionMutex       sync.RWMutex
	retrieveCollectionArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionReturns struct {
		result1 privdata.Collection
		result2 error
	}
	retrieveCollectionReturnsOnCall map[int]struct {
		result1 privdata.Collection
		result2 error
	}
	RetrieveCollectionAccessPolicyStub        func(common.CollectionCriteria) (privdata.CollectionAccessPolicy, error)
	retrieveCollectionAccessPolicyMutex       sync.RWMutex
	retrieveCollectionAccessPolicyArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionAccessPolicyReturns struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	retrieveCollectionAccessPolicyReturnsOnCall map[int]struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}
	RetrieveCollectionConfigPackageStub        func(common.CollectionCriteria) (*common.CollectionConfigPackage, error)
	retrieveCollectionConfigPackageMutex       sync.RWMutex
	retrieveCollectionConfigPackageArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionConfigPackageReturns struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	retrieveCollectionConfigPackageReturnsOnCall map[int]struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}
	RetrieveCollectionPersistenceConfigsStub        func(common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error)
	retrieveCollectionPersistenceConfigsMutex       sync.RWMutex
	retrieveCollectionPersistenceConfigsArgsForCall []struct {
		arg1 common.CollectionCriteria
	}
	retrieveCollectionPersistenceConfigsReturns struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	retrieveCollectionPersistenceConfigsReturnsOnCall map[int]struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}
	RetrieveReadWritePermissionStub        func(common.CollectionCriteria, *pb.SignedProposal, ledger.QueryExecutor) (bool, bool, error)
	retrieveReadWritePermissionMutex       sync.RWMutex
	retrieveReadWritePermissionArgsForCall []struct {
		arg1 common.CollectionCriteria
		arg2 *pb.SignedProposal
		arg3 ledger.QueryExecutor
	}
	retrieveReadWritePermissionReturns struct {
		result1 bool
		result2 bool
		result3 error
	}
	retrieveReadWritePermissionReturnsOnCall map[int]struct {
		result1 bool
		result2 bool
		result3 error
	}
	AccessFilterStub        func(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (privdata.Filter, error)
	accessFilterMutex       sync.RWMutex
	accessFilterArgsForCall []struct {
		channelName            string
		collectionPolicyConfig *common.CollectionPolicyConfig
	}
	accessFilterReturns struct {
		result1 privdata.Filter
		result2 error
	}
	accessFilterReturnsOnCall map[int]struct {
		result1 privdata.Filter
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CollectionStore) RetrieveCollection(arg1 common.CollectionCriteria) (privdata.Collection, error) {
	fake.retrieveCollectionMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionReturnsOnCall[len(fake.retrieveCollectionArgsForCall)]
	fake.retrieveCollectionArgsForCall = append(fake.retrieveCollectionArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollection", []interface{}{arg1})
	fake.retrieveCollectionMutex.Unlock()
	if fake.RetrieveCollectionStub != nil {
		return fake.RetrieveCollectionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveCollectionReturns.result1, fake.retrieveCollectionReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionCallCount() int {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	return len(fake.retrieveCollectionArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	return fake.retrieveCollectionArgsForCall[i].arg1
}

func (fake *CollectionStore) RetrieveCollectionReturns(result1 privdata.Collection, result2 error) {
	fake.RetrieveCollectionStub = nil
	fake.retrieveCollectionReturns = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionReturnsOnCall(i int, result1 privdata.Collection, result2 error) {
	fake.RetrieveCollectionStub = nil
	if fake.retrieveCollectionReturnsOnCall == nil {
		fake.retrieveCollectionReturnsOnCall = make(map[int]struct {
			result1 privdata.Collection
			result2 error
		})
	}
	fake.retrieveCollectionReturnsOnCall[i] = struct {
		result1 privdata.Collection
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicy(arg1 common.CollectionCriteria) (privdata.CollectionAccessPolicy, error) {
	fake.retrieveCollectionAccessPolicyMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionAccessPolicyReturnsOnCall[len(fake.retrieveCollectionAccessPolicyArgsForCall)]
	fake.retrieveCollectionAccessPolicyArgsForCall = append(fake.retrieveCollectionAccessPolicyArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionAccessPolicy", []interface{}{arg1})
	fake.retrieveCollectionAccessPolicyMutex.Unlock()
	if fake.RetrieveCollectionAccessPolicyStub != nil {
		return fake.RetrieveCollectionAccessPolicyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveCollectionAccessPolicyReturns.result1, fake.retrieveCollectionAccessPolicyReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyCallCount() int {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	return len(fake.retrieveCollectionAccessPolicyArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	return fake.retrieveCollectionAccessPolicyArgsForCall[i].arg1
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturns(result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.RetrieveCollectionAccessPolicyStub = nil
	fake.retrieveCollectionAccessPolicyReturns = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionAccessPolicyReturnsOnCall(i int, result1 privdata.CollectionAccessPolicy, result2 error) {
	fake.RetrieveCollectionAccessPolicyStub = nil
	if fake.retrieveCollectionAccessPolicyReturnsOnCall == nil {
		fake.retrieveCollectionAccessPolicyReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionAccessPolicy
			result2 error
		})
	}
	fake.retrieveCollectionAccessPolicyReturnsOnCall[i] = struct {
		result1 privdata.CollectionAccessPolicy
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackage(arg1 common.CollectionCriteria) (*common.CollectionConfigPackage, error) {
	fake.retrieveCollectionConfigPackageMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionConfigPackageReturnsOnCall[len(fake.retrieveCollectionConfigPackageArgsForCall)]
	fake.retrieveCollectionConfigPackageArgsForCall = append(fake.retrieveCollectionConfigPackageArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionConfigPackage", []interface{}{arg1})
	fake.retrieveCollectionConfigPackageMutex.Unlock()
	if fake.RetrieveCollectionConfigPackageStub != nil {
		return fake.RetrieveCollectionConfigPackageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveCollectionConfigPackageReturns.result1, fake.retrieveCollectionConfigPackageReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageCallCount() int {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	return len(fake.retrieveCollectionConfigPackageArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	return fake.retrieveCollectionConfigPackageArgsForCall[i].arg1
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturns(result1 *common.CollectionConfigPackage, result2 error) {
	fake.RetrieveCollectionConfigPackageStub = nil
	fake.retrieveCollectionConfigPackageReturns = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionConfigPackageReturnsOnCall(i int, result1 *common.CollectionConfigPackage, result2 error) {
	fake.RetrieveCollectionConfigPackageStub = nil
	if fake.retrieveCollectionConfigPackageReturnsOnCall == nil {
		fake.retrieveCollectionConfigPackageReturnsOnCall = make(map[int]struct {
			result1 *common.CollectionConfigPackage
			result2 error
		})
	}
	fake.retrieveCollectionConfigPackageReturnsOnCall[i] = struct {
		result1 *common.CollectionConfigPackage
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigs(arg1 common.CollectionCriteria) (privdata.CollectionPersistenceConfigs, error) {
	fake.retrieveCollectionPersistenceConfigsMutex.Lock()
	ret, specificReturn := fake.retrieveCollectionPersistenceConfigsReturnsOnCall[len(fake.retrieveCollectionPersistenceConfigsArgsForCall)]
	fake.retrieveCollectionPersistenceConfigsArgsForCall = append(fake.retrieveCollectionPersistenceConfigsArgsForCall, struct {
		arg1 common.CollectionCriteria
	}{arg1})
	fake.recordInvocation("RetrieveCollectionPersistenceConfigs", []interface{}{arg1})
	fake.retrieveCollectionPersistenceConfigsMutex.Unlock()
	if fake.RetrieveCollectionPersistenceConfigsStub != nil {
		return fake.RetrieveCollectionPersistenceConfigsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveCollectionPersistenceConfigsReturns.result1, fake.retrieveCollectionPersistenceConfigsReturns.result2
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsCallCount() int {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	return len(fake.retrieveCollectionPersistenceConfigsArgsForCall)
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsArgsForCall(i int) common.CollectionCriteria {
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	return fake.retrieveCollectionPersistenceConfigsArgsForCall[i].arg1
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturns(result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	fake.retrieveCollectionPersistenceConfigsReturns = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveCollectionPersistenceConfigsReturnsOnCall(i int, result1 privdata.CollectionPersistenceConfigs, result2 error) {
	fake.RetrieveCollectionPersistenceConfigsStub = nil
	if fake.retrieveCollectionPersistenceConfigsReturnsOnCall == nil {
		fake.retrieveCollectionPersistenceConfigsReturnsOnCall = make(map[int]struct {
			result1 privdata.CollectionPersistenceConfigs
			result2 error
		})
	}
	fake.retrieveCollectionPersistenceConfigsReturnsOnCall[i] = struct {
		result1 privdata.CollectionPersistenceConfigs
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) RetrieveReadWritePermission(arg1 common.CollectionCriteria, arg2 *pb.SignedProposal, arg3 ledger.QueryExecutor) (bool, bool, error) {
	fake.retrieveReadWritePermissionMutex.Lock()
	ret, specificReturn := fake.retrieveReadWritePermissionReturnsOnCall[len(fake.retrieveReadWritePermissionArgsForCall)]
	fake.retrieveReadWritePermissionArgsForCall = append(fake.retrieveReadWritePermissionArgsForCall, struct {
		arg1 common.CollectionCriteria
		arg2 *pb.SignedProposal
		arg3 ledger.QueryExecutor
	}{arg1, arg2, arg3})
	fake.recordInvocation("RetrieveReadWritePermission", []interface{}{arg1, arg2, arg3})
	fake.retrieveReadWritePermissionMutex.Unlock()
	if fake.RetrieveReadWritePermissionStub != nil {
		return fake.RetrieveReadWritePermissionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.retrieveReadWritePermissionReturns.result1, fake.retrieveReadWritePermissionReturns.result2, fake.retrieveReadWritePermissionReturns.result3
}

func (fake *CollectionStore) RetrieveReadWritePermissionCallCount() int {
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	return len(fake.retrieveReadWritePermissionArgsForCall)
}

func (fake *CollectionStore) RetrieveReadWritePermissionArgsForCall(i int) (common.CollectionCriteria, *pb.SignedProposal, ledger.QueryExecutor) {
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	return fake.retrieveReadWritePermissionArgsForCall[i].arg1, fake.retrieveReadWritePermissionArgsForCall[i].arg2, fake.retrieveReadWritePermissionArgsForCall[i].arg3
}

func (fake *CollectionStore) RetrieveReadWritePermissionReturns(result1 bool, result2 bool, result3 error) {
	fake.RetrieveReadWritePermissionStub = nil
	fake.retrieveReadWritePermissionReturns = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *CollectionStore) RetrieveReadWritePermissionReturnsOnCall(i int, result1 bool, result2 bool, result3 error) {
	fake.RetrieveReadWritePermissionStub = nil
	if fake.retrieveReadWritePermissionReturnsOnCall == nil {
		fake.retrieveReadWritePermissionReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 bool
			result3 error
		})
	}
	fake.retrieveReadWritePermissionReturnsOnCall[i] = struct {
		result1 bool
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *CollectionStore) AccessFilter(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (privdata.Filter, error) {
	fake.accessFilterMutex.Lock()
	ret, specificReturn := fake.accessFilterReturnsOnCall[len(fake.accessFilterArgsForCall)]
	fake.accessFilterArgsForCall = append(fake.accessFilterArgsForCall, struct {
		channelName            string
		collectionPolicyConfig *common.CollectionPolicyConfig
	}{channelName, collectionPolicyConfig})
	fake.recordInvocation("AccessFilter", []interface{}{channelName, collectionPolicyConfig})
	fake.accessFilterMutex.Unlock()
	if fake.AccessFilterStub != nil {
		return fake.AccessFilterStub(channelName, collectionPolicyConfig)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.accessFilterReturns.result1, fake.accessFilterReturns.result2
}

func (fake *CollectionStore) AccessFilterCallCount() int {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	return len(fake.accessFilterArgsForCall)
}

func (fake *CollectionStore) AccessFilterArgsForCall(i int) (string, *common.CollectionPolicyConfig) {
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	return fake.accessFilterArgsForCall[i].channelName, fake.accessFilterArgsForCall[i].collectionPolicyConfig
}

func (fake *CollectionStore) AccessFilterReturns(result1 privdata.Filter, result2 error) {
	fake.AccessFilterStub = nil
	fake.accessFilterReturns = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) AccessFilterReturnsOnCall(i int, result1 privdata.Filter, result2 error) {
	fake.AccessFilterStub = nil
	if fake.accessFilterReturnsOnCall == nil {
		fake.accessFilterReturnsOnCall = make(map[int]struct {
			result1 privdata.Filter
			result2 error
		})
	}
	fake.accessFilterReturnsOnCall[i] = struct {
		result1 privdata.Filter
		result2 error
	}{result1, result2}
}

func (fake *CollectionStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.retrieveCollectionMutex.RLock()
	defer fake.retrieveCollectionMutex.RUnlock()
	fake.retrieveCollectionAccessPolicyMutex.RLock()
	defer fake.retrieveCollectionAccessPolicyMutex.RUnlock()
	fake.retrieveCollectionConfigPackageMutex.RLock()
	defer fake.retrieveCollectionConfigPackageMutex.RUnlock()
	fake.retrieveCollectionPersistenceConfigsMutex.RLock()
	defer fake.retrieveCollectionPersistenceConfigsMutex.RUnlock()
	fake.retrieveReadWritePermissionMutex.RLock()
	defer fake.retrieveReadWritePermissionMutex.RUnlock()
	fake.accessFilterMutex.RLock()
	defer fake.accessFilterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CollectionStore) recordInvocation(key string, args []interface{}) {
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
