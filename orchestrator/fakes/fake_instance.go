// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/orchestrator"
)

type FakeInstance struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	IndexStub        func() string
	indexMutex       sync.RWMutex
	indexArgsForCall []struct{}
	indexReturns     struct {
		result1 string
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	IsBackupableStub        func() bool
	isBackupableMutex       sync.RWMutex
	isBackupableArgsForCall []struct{}
	isBackupableReturns     struct {
		result1 bool
	}
	IsPostBackupUnlockableStub        func() bool
	isPostBackupUnlockableMutex       sync.RWMutex
	isPostBackupUnlockableArgsForCall []struct{}
	isPostBackupUnlockableReturns     struct {
		result1 bool
	}
	IsPreBackupLockableStub        func() bool
	isPreBackupLockableMutex       sync.RWMutex
	isPreBackupLockableArgsForCall []struct{}
	isPreBackupLockableReturns     struct {
		result1 bool
	}
	IsRestorableStub        func() bool
	isRestorableMutex       sync.RWMutex
	isRestorableArgsForCall []struct{}
	isRestorableReturns     struct {
		result1 bool
	}
	PreBackupLockStub        func() error
	preBackupLockMutex       sync.RWMutex
	preBackupLockArgsForCall []struct{}
	preBackupLockReturns     struct {
		result1 error
	}
	BackupStub        func() error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct{}
	backupReturns     struct {
		result1 error
	}
	PostBackupUnlockStub        func() error
	postBackupUnlockMutex       sync.RWMutex
	postBackupUnlockArgsForCall []struct{}
	postBackupUnlockReturns     struct {
		result1 error
	}
	RestoreStub        func() error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct{}
	restoreReturns     struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	StreamBackupToRemoteStub        func(io.Reader) error
	streamBackupToRemoteMutex       sync.RWMutex
	streamBackupToRemoteArgsForCall []struct {
		arg1 io.Reader
	}
	streamBackupToRemoteReturns struct {
		result1 error
	}
	BackupSizeStub        func() (string, error)
	backupSizeMutex       sync.RWMutex
	backupSizeArgsForCall []struct{}
	backupSizeReturns     struct {
		result1 string
		result2 error
	}
	BackupChecksumStub        func() (orchestrator.BackupChecksum, error)
	backupChecksumMutex       sync.RWMutex
	backupChecksumArgsForCall []struct{}
	backupChecksumReturns     struct {
		result1 orchestrator.BackupChecksum
		result2 error
	}
	BlobsStub        func() []orchestrator.BackupBlob
	blobsMutex       sync.RWMutex
	blobsArgsForCall []struct{}
	blobsReturns     struct {
		result1 []orchestrator.BackupBlob
	}
	CustomBlobNamesStub        func() []string
	customBlobNamesMutex       sync.RWMutex
	customBlobNamesArgsForCall []struct{}
	customBlobNamesReturns     struct {
		result1 []string
	}
	IsNamedStub        func() bool
	isNamedMutex       sync.RWMutex
	isNamedArgsForCall []struct{}
	isNamedReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstance) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	return fake.nameReturns.result1
}

func (fake *FakeInstance) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeInstance) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) Index() string {
	fake.indexMutex.Lock()
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct{}{})
	fake.recordInvocation("Index", []interface{}{})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub()
	}
	return fake.indexReturns.result1
}

func (fake *FakeInstance) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeInstance) IndexReturns(result1 string) {
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	return fake.iDReturns.result1
}

func (fake *FakeInstance) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeInstance) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) IsBackupable() bool {
	fake.isBackupableMutex.Lock()
	fake.isBackupableArgsForCall = append(fake.isBackupableArgsForCall, struct{}{})
	fake.recordInvocation("IsBackupable", []interface{}{})
	fake.isBackupableMutex.Unlock()
	if fake.IsBackupableStub != nil {
		return fake.IsBackupableStub()
	}
	return fake.isBackupableReturns.result1
}

func (fake *FakeInstance) IsBackupableCallCount() int {
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	return len(fake.isBackupableArgsForCall)
}

func (fake *FakeInstance) IsBackupableReturns(result1 bool) {
	fake.IsBackupableStub = nil
	fake.isBackupableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsPostBackupUnlockable() bool {
	fake.isPostBackupUnlockableMutex.Lock()
	fake.isPostBackupUnlockableArgsForCall = append(fake.isPostBackupUnlockableArgsForCall, struct{}{})
	fake.recordInvocation("IsPostBackupUnlockable", []interface{}{})
	fake.isPostBackupUnlockableMutex.Unlock()
	if fake.IsPostBackupUnlockableStub != nil {
		return fake.IsPostBackupUnlockableStub()
	}
	return fake.isPostBackupUnlockableReturns.result1
}

func (fake *FakeInstance) IsPostBackupUnlockableCallCount() int {
	fake.isPostBackupUnlockableMutex.RLock()
	defer fake.isPostBackupUnlockableMutex.RUnlock()
	return len(fake.isPostBackupUnlockableArgsForCall)
}

func (fake *FakeInstance) IsPostBackupUnlockableReturns(result1 bool) {
	fake.IsPostBackupUnlockableStub = nil
	fake.isPostBackupUnlockableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsPreBackupLockable() bool {
	fake.isPreBackupLockableMutex.Lock()
	fake.isPreBackupLockableArgsForCall = append(fake.isPreBackupLockableArgsForCall, struct{}{})
	fake.recordInvocation("IsPreBackupLockable", []interface{}{})
	fake.isPreBackupLockableMutex.Unlock()
	if fake.IsPreBackupLockableStub != nil {
		return fake.IsPreBackupLockableStub()
	}
	return fake.isPreBackupLockableReturns.result1
}

func (fake *FakeInstance) IsPreBackupLockableCallCount() int {
	fake.isPreBackupLockableMutex.RLock()
	defer fake.isPreBackupLockableMutex.RUnlock()
	return len(fake.isPreBackupLockableArgsForCall)
}

func (fake *FakeInstance) IsPreBackupLockableReturns(result1 bool) {
	fake.IsPreBackupLockableStub = nil
	fake.isPreBackupLockableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) IsRestorable() bool {
	fake.isRestorableMutex.Lock()
	fake.isRestorableArgsForCall = append(fake.isRestorableArgsForCall, struct{}{})
	fake.recordInvocation("IsRestorable", []interface{}{})
	fake.isRestorableMutex.Unlock()
	if fake.IsRestorableStub != nil {
		return fake.IsRestorableStub()
	}
	return fake.isRestorableReturns.result1
}

func (fake *FakeInstance) IsRestorableCallCount() int {
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	return len(fake.isRestorableArgsForCall)
}

func (fake *FakeInstance) IsRestorableReturns(result1 bool) {
	fake.IsRestorableStub = nil
	fake.isRestorableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) PreBackupLock() error {
	fake.preBackupLockMutex.Lock()
	fake.preBackupLockArgsForCall = append(fake.preBackupLockArgsForCall, struct{}{})
	fake.recordInvocation("PreBackupLock", []interface{}{})
	fake.preBackupLockMutex.Unlock()
	if fake.PreBackupLockStub != nil {
		return fake.PreBackupLockStub()
	}
	return fake.preBackupLockReturns.result1
}

func (fake *FakeInstance) PreBackupLockCallCount() int {
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	return len(fake.preBackupLockArgsForCall)
}

func (fake *FakeInstance) PreBackupLockReturns(result1 error) {
	fake.PreBackupLockStub = nil
	fake.preBackupLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Backup() error {
	fake.backupMutex.Lock()
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct{}{})
	fake.recordInvocation("Backup", []interface{}{})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub()
	}
	return fake.backupReturns.result1
}

func (fake *FakeInstance) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeInstance) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) PostBackupUnlock() error {
	fake.postBackupUnlockMutex.Lock()
	fake.postBackupUnlockArgsForCall = append(fake.postBackupUnlockArgsForCall, struct{}{})
	fake.recordInvocation("PostBackupUnlock", []interface{}{})
	fake.postBackupUnlockMutex.Unlock()
	if fake.PostBackupUnlockStub != nil {
		return fake.PostBackupUnlockStub()
	}
	return fake.postBackupUnlockReturns.result1
}

func (fake *FakeInstance) PostBackupUnlockCallCount() int {
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	return len(fake.postBackupUnlockArgsForCall)
}

func (fake *FakeInstance) PostBackupUnlockReturns(result1 error) {
	fake.PostBackupUnlockStub = nil
	fake.postBackupUnlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Restore() error {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct{}{})
	fake.recordInvocation("Restore", []interface{}{})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub()
	}
	return fake.restoreReturns.result1
}

func (fake *FakeInstance) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeInstance) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeInstance) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeInstance) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) StreamBackupToRemote(arg1 io.Reader) error {
	fake.streamBackupToRemoteMutex.Lock()
	fake.streamBackupToRemoteArgsForCall = append(fake.streamBackupToRemoteArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("StreamBackupToRemote", []interface{}{arg1})
	fake.streamBackupToRemoteMutex.Unlock()
	if fake.StreamBackupToRemoteStub != nil {
		return fake.StreamBackupToRemoteStub(arg1)
	}
	return fake.streamBackupToRemoteReturns.result1
}

func (fake *FakeInstance) StreamBackupToRemoteCallCount() int {
	fake.streamBackupToRemoteMutex.RLock()
	defer fake.streamBackupToRemoteMutex.RUnlock()
	return len(fake.streamBackupToRemoteArgsForCall)
}

func (fake *FakeInstance) StreamBackupToRemoteArgsForCall(i int) io.Reader {
	fake.streamBackupToRemoteMutex.RLock()
	defer fake.streamBackupToRemoteMutex.RUnlock()
	return fake.streamBackupToRemoteArgsForCall[i].arg1
}

func (fake *FakeInstance) StreamBackupToRemoteReturns(result1 error) {
	fake.StreamBackupToRemoteStub = nil
	fake.streamBackupToRemoteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) BackupSize() (string, error) {
	fake.backupSizeMutex.Lock()
	fake.backupSizeArgsForCall = append(fake.backupSizeArgsForCall, struct{}{})
	fake.recordInvocation("BackupSize", []interface{}{})
	fake.backupSizeMutex.Unlock()
	if fake.BackupSizeStub != nil {
		return fake.BackupSizeStub()
	}
	return fake.backupSizeReturns.result1, fake.backupSizeReturns.result2
}

func (fake *FakeInstance) BackupSizeCallCount() int {
	fake.backupSizeMutex.RLock()
	defer fake.backupSizeMutex.RUnlock()
	return len(fake.backupSizeArgsForCall)
}

func (fake *FakeInstance) BackupSizeReturns(result1 string, result2 error) {
	fake.BackupSizeStub = nil
	fake.backupSizeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) BackupChecksum() (orchestrator.BackupChecksum, error) {
	fake.backupChecksumMutex.Lock()
	fake.backupChecksumArgsForCall = append(fake.backupChecksumArgsForCall, struct{}{})
	fake.recordInvocation("BackupChecksum", []interface{}{})
	fake.backupChecksumMutex.Unlock()
	if fake.BackupChecksumStub != nil {
		return fake.BackupChecksumStub()
	}
	return fake.backupChecksumReturns.result1, fake.backupChecksumReturns.result2
}

func (fake *FakeInstance) BackupChecksumCallCount() int {
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	return len(fake.backupChecksumArgsForCall)
}

func (fake *FakeInstance) BackupChecksumReturns(result1 orchestrator.BackupChecksum, result2 error) {
	fake.BackupChecksumStub = nil
	fake.backupChecksumReturns = struct {
		result1 orchestrator.BackupChecksum
		result2 error
	}{result1, result2}
}

func (fake *FakeInstance) Blobs() []orchestrator.BackupBlob {
	fake.blobsMutex.Lock()
	fake.blobsArgsForCall = append(fake.blobsArgsForCall, struct{}{})
	fake.recordInvocation("Blobs", []interface{}{})
	fake.blobsMutex.Unlock()
	if fake.BlobsStub != nil {
		return fake.BlobsStub()
	}
	return fake.blobsReturns.result1
}

func (fake *FakeInstance) BlobsCallCount() int {
	fake.blobsMutex.RLock()
	defer fake.blobsMutex.RUnlock()
	return len(fake.blobsArgsForCall)
}

func (fake *FakeInstance) BlobsReturns(result1 []orchestrator.BackupBlob) {
	fake.BlobsStub = nil
	fake.blobsReturns = struct {
		result1 []orchestrator.BackupBlob
	}{result1}
}

func (fake *FakeInstance) CustomBlobNames() []string {
	fake.customBlobNamesMutex.Lock()
	fake.customBlobNamesArgsForCall = append(fake.customBlobNamesArgsForCall, struct{}{})
	fake.recordInvocation("CustomBlobNames", []interface{}{})
	fake.customBlobNamesMutex.Unlock()
	if fake.CustomBlobNamesStub != nil {
		return fake.CustomBlobNamesStub()
	}
	return fake.customBlobNamesReturns.result1
}

func (fake *FakeInstance) CustomBlobNamesCallCount() int {
	fake.customBlobNamesMutex.RLock()
	defer fake.customBlobNamesMutex.RUnlock()
	return len(fake.customBlobNamesArgsForCall)
}

func (fake *FakeInstance) CustomBlobNamesReturns(result1 []string) {
	fake.CustomBlobNamesStub = nil
	fake.customBlobNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) IsNamed() bool {
	fake.isNamedMutex.Lock()
	fake.isNamedArgsForCall = append(fake.isNamedArgsForCall, struct{}{})
	fake.recordInvocation("IsNamed", []interface{}{})
	fake.isNamedMutex.Unlock()
	if fake.IsNamedStub != nil {
		return fake.IsNamedStub()
	}
	return fake.isNamedReturns.result1
}

func (fake *FakeInstance) IsNamedCallCount() int {
	fake.isNamedMutex.RLock()
	defer fake.isNamedMutex.RUnlock()
	return len(fake.isNamedArgsForCall)
}

func (fake *FakeInstance) IsNamedReturns(result1 bool) {
	fake.IsNamedStub = nil
	fake.isNamedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	fake.isPostBackupUnlockableMutex.RLock()
	defer fake.isPostBackupUnlockableMutex.RUnlock()
	fake.isPreBackupLockableMutex.RLock()
	defer fake.isPreBackupLockableMutex.RUnlock()
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.streamBackupToRemoteMutex.RLock()
	defer fake.streamBackupToRemoteMutex.RUnlock()
	fake.backupSizeMutex.RLock()
	defer fake.backupSizeMutex.RUnlock()
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	fake.blobsMutex.RLock()
	defer fake.blobsMutex.RUnlock()
	fake.customBlobNamesMutex.RLock()
	defer fake.customBlobNamesMutex.RUnlock()
	fake.isNamedMutex.RLock()
	defer fake.isNamedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeInstance) recordInvocation(key string, args []interface{}) {
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

var _ orchestrator.Instance = new(FakeInstance)
