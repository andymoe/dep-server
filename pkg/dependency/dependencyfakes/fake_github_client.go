// Code generated by counterfeiter. DO NOT EDIT.
package dependencyfakes

import (
	"sync"

	"github.com/paketo-buildpacks/dep-server/pkg/dependency"
	"github.com/paketo-buildpacks/dep-server/pkg/dependency/internal"
)

type FakeGithubClient struct {
	DownloadReleaseAssetStub        func(string, string, string, string, string) (string, error)
	downloadReleaseAssetMutex       sync.RWMutex
	downloadReleaseAssetArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	downloadReleaseAssetReturns struct {
		result1 string
		result2 error
	}
	downloadReleaseAssetReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DownloadSourceTarballStub        func(string, string, string, string) (string, error)
	downloadSourceTarballMutex       sync.RWMutex
	downloadSourceTarballArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	downloadSourceTarballReturns struct {
		result1 string
		result2 error
	}
	downloadSourceTarballReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetReleaseAssetStub        func(string, string, string, string) ([]byte, error)
	getReleaseAssetMutex       sync.RWMutex
	getReleaseAssetArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	getReleaseAssetReturns struct {
		result1 []byte
		result2 error
	}
	getReleaseAssetReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetReleaseTagsStub        func(string, string) ([]internal.GithubRelease, error)
	getReleaseTagsMutex       sync.RWMutex
	getReleaseTagsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getReleaseTagsReturns struct {
		result1 []internal.GithubRelease
		result2 error
	}
	getReleaseTagsReturnsOnCall map[int]struct {
		result1 []internal.GithubRelease
		result2 error
	}
	GetTagCommitStub        func(string, string, string) (internal.GithubTagCommit, error)
	getTagCommitMutex       sync.RWMutex
	getTagCommitArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getTagCommitReturns struct {
		result1 internal.GithubTagCommit
		result2 error
	}
	getTagCommitReturnsOnCall map[int]struct {
		result1 internal.GithubTagCommit
		result2 error
	}
	GetTagsStub        func(string, string) ([]string, error)
	getTagsMutex       sync.RWMutex
	getTagsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getTagsReturns struct {
		result1 []string
		result2 error
	}
	getTagsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGithubClient) DownloadReleaseAsset(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (string, error) {
	fake.downloadReleaseAssetMutex.Lock()
	ret, specificReturn := fake.downloadReleaseAssetReturnsOnCall[len(fake.downloadReleaseAssetArgsForCall)]
	fake.downloadReleaseAssetArgsForCall = append(fake.downloadReleaseAssetArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("DownloadReleaseAsset", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.downloadReleaseAssetMutex.Unlock()
	if fake.DownloadReleaseAssetStub != nil {
		return fake.DownloadReleaseAssetStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.downloadReleaseAssetReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) DownloadReleaseAssetCallCount() int {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	return len(fake.downloadReleaseAssetArgsForCall)
}

func (fake *FakeGithubClient) DownloadReleaseAssetCalls(stub func(string, string, string, string, string) (string, error)) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = stub
}

func (fake *FakeGithubClient) DownloadReleaseAssetArgsForCall(i int) (string, string, string, string, string) {
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	argsForCall := fake.downloadReleaseAssetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeGithubClient) DownloadReleaseAssetReturns(result1 string, result2 error) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = nil
	fake.downloadReleaseAssetReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) DownloadReleaseAssetReturnsOnCall(i int, result1 string, result2 error) {
	fake.downloadReleaseAssetMutex.Lock()
	defer fake.downloadReleaseAssetMutex.Unlock()
	fake.DownloadReleaseAssetStub = nil
	if fake.downloadReleaseAssetReturnsOnCall == nil {
		fake.downloadReleaseAssetReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.downloadReleaseAssetReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) DownloadSourceTarball(arg1 string, arg2 string, arg3 string, arg4 string) (string, error) {
	fake.downloadSourceTarballMutex.Lock()
	ret, specificReturn := fake.downloadSourceTarballReturnsOnCall[len(fake.downloadSourceTarballArgsForCall)]
	fake.downloadSourceTarballArgsForCall = append(fake.downloadSourceTarballArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DownloadSourceTarball", []interface{}{arg1, arg2, arg3, arg4})
	fake.downloadSourceTarballMutex.Unlock()
	if fake.DownloadSourceTarballStub != nil {
		return fake.DownloadSourceTarballStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.downloadSourceTarballReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) DownloadSourceTarballCallCount() int {
	fake.downloadSourceTarballMutex.RLock()
	defer fake.downloadSourceTarballMutex.RUnlock()
	return len(fake.downloadSourceTarballArgsForCall)
}

func (fake *FakeGithubClient) DownloadSourceTarballCalls(stub func(string, string, string, string) (string, error)) {
	fake.downloadSourceTarballMutex.Lock()
	defer fake.downloadSourceTarballMutex.Unlock()
	fake.DownloadSourceTarballStub = stub
}

func (fake *FakeGithubClient) DownloadSourceTarballArgsForCall(i int) (string, string, string, string) {
	fake.downloadSourceTarballMutex.RLock()
	defer fake.downloadSourceTarballMutex.RUnlock()
	argsForCall := fake.downloadSourceTarballArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGithubClient) DownloadSourceTarballReturns(result1 string, result2 error) {
	fake.downloadSourceTarballMutex.Lock()
	defer fake.downloadSourceTarballMutex.Unlock()
	fake.DownloadSourceTarballStub = nil
	fake.downloadSourceTarballReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) DownloadSourceTarballReturnsOnCall(i int, result1 string, result2 error) {
	fake.downloadSourceTarballMutex.Lock()
	defer fake.downloadSourceTarballMutex.Unlock()
	fake.DownloadSourceTarballStub = nil
	if fake.downloadSourceTarballReturnsOnCall == nil {
		fake.downloadSourceTarballReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.downloadSourceTarballReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetReleaseAsset(arg1 string, arg2 string, arg3 string, arg4 string) ([]byte, error) {
	fake.getReleaseAssetMutex.Lock()
	ret, specificReturn := fake.getReleaseAssetReturnsOnCall[len(fake.getReleaseAssetArgsForCall)]
	fake.getReleaseAssetArgsForCall = append(fake.getReleaseAssetArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetReleaseAsset", []interface{}{arg1, arg2, arg3, arg4})
	fake.getReleaseAssetMutex.Unlock()
	if fake.GetReleaseAssetStub != nil {
		return fake.GetReleaseAssetStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReleaseAssetReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) GetReleaseAssetCallCount() int {
	fake.getReleaseAssetMutex.RLock()
	defer fake.getReleaseAssetMutex.RUnlock()
	return len(fake.getReleaseAssetArgsForCall)
}

func (fake *FakeGithubClient) GetReleaseAssetCalls(stub func(string, string, string, string) ([]byte, error)) {
	fake.getReleaseAssetMutex.Lock()
	defer fake.getReleaseAssetMutex.Unlock()
	fake.GetReleaseAssetStub = stub
}

func (fake *FakeGithubClient) GetReleaseAssetArgsForCall(i int) (string, string, string, string) {
	fake.getReleaseAssetMutex.RLock()
	defer fake.getReleaseAssetMutex.RUnlock()
	argsForCall := fake.getReleaseAssetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGithubClient) GetReleaseAssetReturns(result1 []byte, result2 error) {
	fake.getReleaseAssetMutex.Lock()
	defer fake.getReleaseAssetMutex.Unlock()
	fake.GetReleaseAssetStub = nil
	fake.getReleaseAssetReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetReleaseAssetReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getReleaseAssetMutex.Lock()
	defer fake.getReleaseAssetMutex.Unlock()
	fake.GetReleaseAssetStub = nil
	if fake.getReleaseAssetReturnsOnCall == nil {
		fake.getReleaseAssetReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getReleaseAssetReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetReleaseTags(arg1 string, arg2 string) ([]internal.GithubRelease, error) {
	fake.getReleaseTagsMutex.Lock()
	ret, specificReturn := fake.getReleaseTagsReturnsOnCall[len(fake.getReleaseTagsArgsForCall)]
	fake.getReleaseTagsArgsForCall = append(fake.getReleaseTagsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetReleaseTags", []interface{}{arg1, arg2})
	fake.getReleaseTagsMutex.Unlock()
	if fake.GetReleaseTagsStub != nil {
		return fake.GetReleaseTagsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReleaseTagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) GetReleaseTagsCallCount() int {
	fake.getReleaseTagsMutex.RLock()
	defer fake.getReleaseTagsMutex.RUnlock()
	return len(fake.getReleaseTagsArgsForCall)
}

func (fake *FakeGithubClient) GetReleaseTagsCalls(stub func(string, string) ([]internal.GithubRelease, error)) {
	fake.getReleaseTagsMutex.Lock()
	defer fake.getReleaseTagsMutex.Unlock()
	fake.GetReleaseTagsStub = stub
}

func (fake *FakeGithubClient) GetReleaseTagsArgsForCall(i int) (string, string) {
	fake.getReleaseTagsMutex.RLock()
	defer fake.getReleaseTagsMutex.RUnlock()
	argsForCall := fake.getReleaseTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGithubClient) GetReleaseTagsReturns(result1 []internal.GithubRelease, result2 error) {
	fake.getReleaseTagsMutex.Lock()
	defer fake.getReleaseTagsMutex.Unlock()
	fake.GetReleaseTagsStub = nil
	fake.getReleaseTagsReturns = struct {
		result1 []internal.GithubRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetReleaseTagsReturnsOnCall(i int, result1 []internal.GithubRelease, result2 error) {
	fake.getReleaseTagsMutex.Lock()
	defer fake.getReleaseTagsMutex.Unlock()
	fake.GetReleaseTagsStub = nil
	if fake.getReleaseTagsReturnsOnCall == nil {
		fake.getReleaseTagsReturnsOnCall = make(map[int]struct {
			result1 []internal.GithubRelease
			result2 error
		})
	}
	fake.getReleaseTagsReturnsOnCall[i] = struct {
		result1 []internal.GithubRelease
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetTagCommit(arg1 string, arg2 string, arg3 string) (internal.GithubTagCommit, error) {
	fake.getTagCommitMutex.Lock()
	ret, specificReturn := fake.getTagCommitReturnsOnCall[len(fake.getTagCommitArgsForCall)]
	fake.getTagCommitArgsForCall = append(fake.getTagCommitArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetTagCommit", []interface{}{arg1, arg2, arg3})
	fake.getTagCommitMutex.Unlock()
	if fake.GetTagCommitStub != nil {
		return fake.GetTagCommitStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTagCommitReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) GetTagCommitCallCount() int {
	fake.getTagCommitMutex.RLock()
	defer fake.getTagCommitMutex.RUnlock()
	return len(fake.getTagCommitArgsForCall)
}

func (fake *FakeGithubClient) GetTagCommitCalls(stub func(string, string, string) (internal.GithubTagCommit, error)) {
	fake.getTagCommitMutex.Lock()
	defer fake.getTagCommitMutex.Unlock()
	fake.GetTagCommitStub = stub
}

func (fake *FakeGithubClient) GetTagCommitArgsForCall(i int) (string, string, string) {
	fake.getTagCommitMutex.RLock()
	defer fake.getTagCommitMutex.RUnlock()
	argsForCall := fake.getTagCommitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGithubClient) GetTagCommitReturns(result1 internal.GithubTagCommit, result2 error) {
	fake.getTagCommitMutex.Lock()
	defer fake.getTagCommitMutex.Unlock()
	fake.GetTagCommitStub = nil
	fake.getTagCommitReturns = struct {
		result1 internal.GithubTagCommit
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetTagCommitReturnsOnCall(i int, result1 internal.GithubTagCommit, result2 error) {
	fake.getTagCommitMutex.Lock()
	defer fake.getTagCommitMutex.Unlock()
	fake.GetTagCommitStub = nil
	if fake.getTagCommitReturnsOnCall == nil {
		fake.getTagCommitReturnsOnCall = make(map[int]struct {
			result1 internal.GithubTagCommit
			result2 error
		})
	}
	fake.getTagCommitReturnsOnCall[i] = struct {
		result1 internal.GithubTagCommit
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetTags(arg1 string, arg2 string) ([]string, error) {
	fake.getTagsMutex.Lock()
	ret, specificReturn := fake.getTagsReturnsOnCall[len(fake.getTagsArgsForCall)]
	fake.getTagsArgsForCall = append(fake.getTagsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetTags", []interface{}{arg1, arg2})
	fake.getTagsMutex.Unlock()
	if fake.GetTagsStub != nil {
		return fake.GetTagsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGithubClient) GetTagsCallCount() int {
	fake.getTagsMutex.RLock()
	defer fake.getTagsMutex.RUnlock()
	return len(fake.getTagsArgsForCall)
}

func (fake *FakeGithubClient) GetTagsCalls(stub func(string, string) ([]string, error)) {
	fake.getTagsMutex.Lock()
	defer fake.getTagsMutex.Unlock()
	fake.GetTagsStub = stub
}

func (fake *FakeGithubClient) GetTagsArgsForCall(i int) (string, string) {
	fake.getTagsMutex.RLock()
	defer fake.getTagsMutex.RUnlock()
	argsForCall := fake.getTagsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGithubClient) GetTagsReturns(result1 []string, result2 error) {
	fake.getTagsMutex.Lock()
	defer fake.getTagsMutex.Unlock()
	fake.GetTagsStub = nil
	fake.getTagsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) GetTagsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getTagsMutex.Lock()
	defer fake.getTagsMutex.Unlock()
	fake.GetTagsStub = nil
	if fake.getTagsReturnsOnCall == nil {
		fake.getTagsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getTagsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeGithubClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadReleaseAssetMutex.RLock()
	defer fake.downloadReleaseAssetMutex.RUnlock()
	fake.downloadSourceTarballMutex.RLock()
	defer fake.downloadSourceTarballMutex.RUnlock()
	fake.getReleaseAssetMutex.RLock()
	defer fake.getReleaseAssetMutex.RUnlock()
	fake.getReleaseTagsMutex.RLock()
	defer fake.getReleaseTagsMutex.RUnlock()
	fake.getTagCommitMutex.RLock()
	defer fake.getTagCommitMutex.RUnlock()
	fake.getTagsMutex.RLock()
	defer fake.getTagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGithubClient) recordInvocation(key string, args []interface{}) {
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

var _ dependency.GithubClient = new(FakeGithubClient)
