// Code generated by counterfeiter. DO NOT EDIT.
package dependencyfakes

import (
	"sync"

	"github.com/paketo-buildpacks/dep-server/pkg/dependency"
	"github.com/paketo-buildpacks/dep-server/pkg/dependency/internal"
)

type FakeWebClient struct {
	DownloadStub        func(string, string, ...internal.RequestOption) error
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []internal.RequestOption
	}
	downloadReturns struct {
		result1 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string, ...internal.RequestOption) ([]byte, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 []internal.RequestOption
	}
	getReturns struct {
		result1 []byte
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWebClient) Download(arg1 string, arg2 string, arg3 ...internal.RequestOption) error {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []internal.RequestOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("Download", []interface{}{arg1, arg2, arg3})
	fake.downloadMutex.Unlock()
	if fake.DownloadStub != nil {
		return fake.DownloadStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadReturns
	return fakeReturns.result1
}

func (fake *FakeWebClient) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeWebClient) DownloadCalls(stub func(string, string, ...internal.RequestOption) error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = stub
}

func (fake *FakeWebClient) DownloadArgsForCall(i int) (string, string, []internal.RequestOption) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	argsForCall := fake.downloadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeWebClient) DownloadReturns(result1 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebClient) DownloadReturnsOnCall(i int, result1 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebClient) Get(arg1 string, arg2 ...internal.RequestOption) ([]byte, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 []internal.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWebClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeWebClient) GetCalls(stub func(string, ...internal.RequestOption) ([]byte, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeWebClient) GetArgsForCall(i int) (string, []internal.RequestOption) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWebClient) GetReturns(result1 []byte, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeWebClient) GetReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeWebClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWebClient) recordInvocation(key string, args []interface{}) {
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

var _ dependency.WebClient = new(FakeWebClient)