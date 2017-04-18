// This file was generated by counterfeiter
package lockerfakes

import (
	"sync"
)

type FakeGitRepo struct {
	CloneOrPullStub        func() error
	cloneOrPullMutex       sync.RWMutex
	cloneOrPullArgsForCall []struct{}
	cloneOrPullReturns     struct {
		result1 error
	}
	cloneOrPullReturnsOnCall map[int]struct {
		result1 error
	}
	CommitAndPushStub        func(message, user string) error
	commitAndPushMutex       sync.RWMutex
	commitAndPushArgsForCall []struct {
		message string
		user    string
	}
	commitAndPushReturns struct {
		result1 error
	}
	commitAndPushReturnsOnCall map[int]struct {
		result1 error
	}
	DirStub        func() string
	dirMutex       sync.RWMutex
	dirArgsForCall []struct{}
	dirReturns     struct {
		result1 string
	}
	dirReturnsOnCall map[int]struct {
		result1 string
	}
	LatestCommitStub        func(pool string) (committer, date, message string, err error)
	latestCommitMutex       sync.RWMutex
	latestCommitArgsForCall []struct {
		pool string
	}
	latestCommitReturns struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}
	latestCommitReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitRepo) CloneOrPull() error {
	fake.cloneOrPullMutex.Lock()
	ret, specificReturn := fake.cloneOrPullReturnsOnCall[len(fake.cloneOrPullArgsForCall)]
	fake.cloneOrPullArgsForCall = append(fake.cloneOrPullArgsForCall, struct{}{})
	fake.recordInvocation("CloneOrPull", []interface{}{})
	fake.cloneOrPullMutex.Unlock()
	if fake.CloneOrPullStub != nil {
		return fake.CloneOrPullStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloneOrPullReturns.result1
}

func (fake *FakeGitRepo) CloneOrPullCallCount() int {
	fake.cloneOrPullMutex.RLock()
	defer fake.cloneOrPullMutex.RUnlock()
	return len(fake.cloneOrPullArgsForCall)
}

func (fake *FakeGitRepo) CloneOrPullReturns(result1 error) {
	fake.CloneOrPullStub = nil
	fake.cloneOrPullReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) CloneOrPullReturnsOnCall(i int, result1 error) {
	fake.CloneOrPullStub = nil
	if fake.cloneOrPullReturnsOnCall == nil {
		fake.cloneOrPullReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cloneOrPullReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) CommitAndPush(message string, user string) error {
	fake.commitAndPushMutex.Lock()
	ret, specificReturn := fake.commitAndPushReturnsOnCall[len(fake.commitAndPushArgsForCall)]
	fake.commitAndPushArgsForCall = append(fake.commitAndPushArgsForCall, struct {
		message string
		user    string
	}{message, user})
	fake.recordInvocation("CommitAndPush", []interface{}{message, user})
	fake.commitAndPushMutex.Unlock()
	if fake.CommitAndPushStub != nil {
		return fake.CommitAndPushStub(message, user)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.commitAndPushReturns.result1
}

func (fake *FakeGitRepo) CommitAndPushCallCount() int {
	fake.commitAndPushMutex.RLock()
	defer fake.commitAndPushMutex.RUnlock()
	return len(fake.commitAndPushArgsForCall)
}

func (fake *FakeGitRepo) CommitAndPushArgsForCall(i int) (string, string) {
	fake.commitAndPushMutex.RLock()
	defer fake.commitAndPushMutex.RUnlock()
	return fake.commitAndPushArgsForCall[i].message, fake.commitAndPushArgsForCall[i].user
}

func (fake *FakeGitRepo) CommitAndPushReturns(result1 error) {
	fake.CommitAndPushStub = nil
	fake.commitAndPushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) CommitAndPushReturnsOnCall(i int, result1 error) {
	fake.CommitAndPushStub = nil
	if fake.commitAndPushReturnsOnCall == nil {
		fake.commitAndPushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitAndPushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitRepo) Dir() string {
	fake.dirMutex.Lock()
	ret, specificReturn := fake.dirReturnsOnCall[len(fake.dirArgsForCall)]
	fake.dirArgsForCall = append(fake.dirArgsForCall, struct{}{})
	fake.recordInvocation("Dir", []interface{}{})
	fake.dirMutex.Unlock()
	if fake.DirStub != nil {
		return fake.DirStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dirReturns.result1
}

func (fake *FakeGitRepo) DirCallCount() int {
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	return len(fake.dirArgsForCall)
}

func (fake *FakeGitRepo) DirReturns(result1 string) {
	fake.DirStub = nil
	fake.dirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGitRepo) DirReturnsOnCall(i int, result1 string) {
	fake.DirStub = nil
	if fake.dirReturnsOnCall == nil {
		fake.dirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGitRepo) LatestCommit(pool string) (committer, date, message string, err error) {
	fake.latestCommitMutex.Lock()
	ret, specificReturn := fake.latestCommitReturnsOnCall[len(fake.latestCommitArgsForCall)]
	fake.latestCommitArgsForCall = append(fake.latestCommitArgsForCall, struct {
		pool string
	}{pool})
	fake.recordInvocation("LatestCommit", []interface{}{pool})
	fake.latestCommitMutex.Unlock()
	if fake.LatestCommitStub != nil {
		return fake.LatestCommitStub(pool)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.latestCommitReturns.result1, fake.latestCommitReturns.result2, fake.latestCommitReturns.result3, fake.latestCommitReturns.result4
}

func (fake *FakeGitRepo) LatestCommitCallCount() int {
	fake.latestCommitMutex.RLock()
	defer fake.latestCommitMutex.RUnlock()
	return len(fake.latestCommitArgsForCall)
}

func (fake *FakeGitRepo) LatestCommitArgsForCall(i int) string {
	fake.latestCommitMutex.RLock()
	defer fake.latestCommitMutex.RUnlock()
	return fake.latestCommitArgsForCall[i].pool
}

func (fake *FakeGitRepo) LatestCommitReturns(result1 string, result2 string, result3 string, result4 error) {
	fake.LatestCommitStub = nil
	fake.latestCommitReturns = struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeGitRepo) LatestCommitReturnsOnCall(i int, result1 string, result2 string, result3 string, result4 error) {
	fake.LatestCommitStub = nil
	if fake.latestCommitReturnsOnCall == nil {
		fake.latestCommitReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 string
			result4 error
		})
	}
	fake.latestCommitReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeGitRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloneOrPullMutex.RLock()
	defer fake.cloneOrPullMutex.RUnlock()
	fake.commitAndPushMutex.RLock()
	defer fake.commitAndPushMutex.RUnlock()
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	fake.latestCommitMutex.RLock()
	defer fake.latestCommitMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGitRepo) recordInvocation(key string, args []interface{}) {
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
