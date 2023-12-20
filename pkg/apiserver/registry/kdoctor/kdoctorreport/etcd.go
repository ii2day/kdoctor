// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package kdoctorreport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"
	"k8s.io/klog/v2"

	"github.com/kdoctor-io/kdoctor/pkg/apiserver/registry"
	crd "github.com/kdoctor-io/kdoctor/pkg/k8s/apis/kdoctor.io/v1beta1"
	"github.com/kdoctor-io/kdoctor/pkg/k8s/apis/system/v1beta1"
	"github.com/kdoctor-io/kdoctor/pkg/k8s/client/clientset/versioned"
)

const dir = "/report"

func NewREST(clientSet *versioned.Clientset, scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	restOptions, err := optsGetter.GetRESTOptions(v1beta1.Resource("kdoctorreports"))
	if nil != err {
		return nil, err
	}

	dryRunnableStorage, destroyFunc := NewStorage(clientSet, restOptions)
	store := &genericregistry.Store{
		NewFunc:     func() runtime.Object { return &v1beta1.KdoctorReport{} },
		NewListFunc: func() runtime.Object { return &v1beta1.KdoctorReportList{} },
		KeyRootFunc: func(ctx context.Context) string {
			return restOptions.ResourcePrefix
		},
		KeyFunc: func(ctx context.Context, name string) (string, error) {
			return genericregistry.NoNamespaceKeyFunc(ctx, restOptions.ResourcePrefix, name)
		},
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*v1beta1.KdoctorReport).Name, nil
		},
		DefaultQualifiedResource: v1beta1.Resource("kdoctorreports"),
		PredicateFunc:            MatchKdoctorReport,

		CreateStrategy:          strategy,
		UpdateStrategy:          strategy,
		DeleteStrategy:          strategy,
		EnableGarbageCollection: true,

		Storage:     dryRunnableStorage,
		DestroyFunc: destroyFunc,

		// TableConvertor: printers.NewTableGenerator(v1beta1.Resource("kdoctorreports")),
	}

	return &registry.REST{Store: store}, nil
}

func NewStorage(clientSet *versioned.Clientset, restOptions generic.RESTOptions) (genericregistry.DryRunnableStorage, factory.DestroyFunc) {
	dryRunnableStorage := genericregistry.DryRunnableStorage{
		Storage: &kdoctorReportStorage{
			clientSet: clientSet,
		},
		Codec: restOptions.StorageConfig.Codec,
	}

	return dryRunnableStorage, func() {}
}

var _ storage.Interface = &kdoctorReportStorage{}

type kdoctorReportStorage struct {
	clientSet    *versioned.Clientset
	resourceName string
}

func (p kdoctorReportStorage) Versioner() storage.Versioner {
	return storage.APIObjectVersioner{}
}

func (p kdoctorReportStorage) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	return fmt.Errorf("create API not implement")
}

func (p kdoctorReportStorage) Delete(ctx context.Context, key string, out runtime.Object, preconditions *storage.Preconditions, validateDeletion storage.ValidateObjectFunc, cachedExistingObject runtime.Object) error {
	return fmt.Errorf("delete API not implement")
}

func (p kdoctorReportStorage) Watch(ctx context.Context, key string, opts storage.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("watch API not implement")

}

func (p kdoctorReportStorage) Get(ctx context.Context, key string, opts storage.GetOptions, objPtr runtime.Object) error {
	klog.Infof("Get called with key: %v on resource %v\n", key, p.resourceName)
	kdoctorReport := objPtr.(*v1beta1.KdoctorReport)
	var taskStatus *crd.TaskStatus
	var taskType string
	var creationTimestamp metav1.Time
	_, name, err := NamespaceAndNameFromKey(key, false)
	if nil != err {
		return err
	}

	// TODO (Icarus9913): we need options to specify which CRD that we are looking for.
	netdns, err := p.clientSet.KdoctorV1beta1().Netdnses().Get(ctx, name, metav1.GetOptions{})
	if nil != err {
		if errors.IsNotFound(err) {
			klog.Infof("no NetDNS %s found", name)
		} else {
			return fmt.Errorf("failed to get NetDNS %s, error: %w", name, err)
		}
	} else {
		fmt.Printf("succeed to get NetDNS %s\n", name)
		taskStatus = netdns.Status.DeepCopy()
		creationTimestamp = netdns.CreationTimestamp
		taskType = v1beta1.NetDNSTaskName
		kdoctorReport.Task.Spec.NetDNSTaskSpec = &netdns.Spec
	}

	netReach, err := p.clientSet.KdoctorV1beta1().NetReaches().Get(ctx, name, metav1.GetOptions{})
	if nil != err {
		if errors.IsNotFound(err) {
			klog.Infof("no NetReachHealthy %s found", name)
		} else {
			return fmt.Errorf("failed to get NetReachHealthy %s, error: %w", name, err)
		}
	} else {
		fmt.Printf("succeed to get NetReachHealthy %s\n", name)
		taskStatus = netReach.Status.DeepCopy()
		creationTimestamp = netReach.CreationTimestamp
		taskType = v1beta1.NetReachTaskName
		kdoctorReport.Task.Spec.NetReachTaskSpec = &netReach.Spec
	}

	appHttpHealthy, err := p.clientSet.KdoctorV1beta1().AppHttpHealthies().Get(ctx, name, metav1.GetOptions{})
	if nil != err {
		if errors.IsNotFound(err) {
			klog.Infof("no HttpAppHealthy %s found", name)
		} else {
			return fmt.Errorf("failed to get HttpAppHealthy %s, error: %w", name, err)
		}
	} else {
		fmt.Printf("succeed to get HttpAppHealthy %s\n", name)
		taskStatus = appHttpHealthy.Status.DeepCopy()
		creationTimestamp = appHttpHealthy.CreationTimestamp
		taskType = v1beta1.AppHttpHealthyTaskName
		kdoctorReport.Task.Spec.AppHttpHealthyTaskSpec = &appHttpHealthy.Spec
	}

	if taskStatus == nil {
		return fmt.Errorf("no crd instance %s found", name)
	}
	var status string
	if taskStatus.Finish {
		status = "Finished"
	} else {
		status = "NotFinished"
	}

	var toTalRoundNumber, finishedRoundNumber int64
	if taskStatus.ExpectedRound != nil {
		toTalRoundNumber = *taskStatus.ExpectedRound
	}
	if taskStatus.DoneRound != nil {
		finishedRoundNumber = *taskStatus.DoneRound
	}

	readDir, err := os.ReadDir(dir)
	if nil != err {
		return fmt.Errorf("failed to read directory %s, error: %w", dir, err)
	}
	var fileNameList []string
	for _, item := range readDir {
		if item.IsDir() {
			continue
		}

		if strings.Contains(item.Name(), name) && !strings.Contains(item.Name(), summary) {
			fileNameList = append(fileNameList, item.Name())
		}
	}

	getReports, latestRoundNumber, err := p.getLatestRoundReports(name, fileNameList)
	if nil != err {
		return fmt.Errorf("failed to get latest round reports: %w", err)
	}
	if getReports == nil {
		return fmt.Errorf("no '%s' reports found", name)
	}

	kdoctorReport.CreationTimestamp = creationTimestamp
	kdoctorReport.Report = v1beta1.Reports{LatestRoundReport: getReports}
	kdoctorReport.Status = v1beta1.Status{
		ToTalRoundNumber:    toTalRoundNumber,
		FinishedRoundNumber: finishedRoundNumber,
		Status:              status,
		RoundNumber:         latestRoundNumber,
	}
	kdoctorReport.Task.TaskName = name
	kdoctorReport.Task.TaskType = taskType
	kdoctorReport.Name = strings.ToLower(taskType) + "-" + name
	kdoctorReport.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   v1beta1.GroupName,
		Version: v1beta1.V1betaVersion,
		Kind:    v1beta1.KindKdoctorReport,
	})

	return nil
}

const summary = "summary"

func (p kdoctorReportStorage) GetList(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error {
	readDir, err := os.ReadDir(dir)
	if nil != err {
		return fmt.Errorf("failed to read directory %s, error: %w", dir, err)
	}
	var fileNameList []string
	for _, item := range readDir {
		if item.IsDir() {
			continue
		}
		fileNameList = append(fileNameList, item.Name())
	}

	kdoctorReportList := listObj.(*v1beta1.KdoctorReportList)
	var resList []runtime.Object

	{
		netDNSKdoctorReports, err := p.getNetDNSKdoctorReports(ctx, fileNameList)
		if nil != err {
			return err
		}
		for i := range netDNSKdoctorReports {
			resList = append(resList, netDNSKdoctorReports[i].DeepCopy())
		}
	}

	{
		httpAppHealthyReports, err := p.getHttpAppHealthyReports(ctx, fileNameList)
		if nil != err {
			return err
		}
		for i := range httpAppHealthyReports {
			resList = append(resList, httpAppHealthyReports[i].DeepCopy())
		}
	}

	{
		netReachHealthyReports, err := p.getNetReachHealthyReports(ctx, fileNameList)
		if nil != err {
			return err
		}
		for i := range netReachHealthyReports {
			resList = append(resList, netReachHealthyReports[i].DeepCopy())
		}
	}

	err = meta.SetList(kdoctorReportList, resList)
	if nil != err {
		return err
	}

	kdoctorReportList.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   v1beta1.GroupName,
		Version: v1beta1.V1betaVersion,
		Kind:    v1beta1.KindKdoctorReportList,
	})

	return nil
}

func (p kdoctorReportStorage) GuaranteedUpdate(ctx context.Context, key string, destination runtime.Object, ignoreNotFound bool, preconditions *storage.Preconditions, tryUpdate storage.UpdateFunc, cachedExistingObject runtime.Object) error {
	return fmt.Errorf("GuaranteedUpdate API not implement")
}

func (p kdoctorReportStorage) Count(key string) (int64, error) {
	return 0, fmt.Errorf("Count not supported for key: %s", key)
}

func (p kdoctorReportStorage) getLatestRoundReports(key string, fileNameList []string) (*[]v1beta1.Report, int64, error) {
	var reports []v1beta1.Report
	for _, netDNSFileName := range fileNameList {
		split := strings.Split(netDNSFileName, "_")
		if len(split) != 5 {
			klog.Infof("unrecognized file %s", netDNSFileName)
			continue
		}

		if key == split[1] {
			file, err := os.Open(path.Join(dir, netDNSFileName))
			if nil != err {
				return nil, -1, err
			}
			readAll, err := io.ReadAll(file)
			if nil != err {
				return nil, -1, err
			}

			report := v1beta1.Report{}
			err = json.Unmarshal(readAll, &report)
			if nil != err {
				return nil, -1, err
			}
			reports = append(reports, report)
		}
	}

	var latestRoundNumber int64
	result := func() *[]v1beta1.Report {
		var latestReports []v1beta1.Report

		for _, tmpReport := range reports {
			if tmpReport.RoundNumber > latestRoundNumber {
				latestRoundNumber = tmpReport.RoundNumber
				latestReports = []v1beta1.Report{tmpReport}
			} else if tmpReport.RoundNumber == latestRoundNumber {
				latestReports = append(latestReports, tmpReport)
			} else {
				continue
			}
		}
		if len(latestReports) == 0 {
			return nil
		}

		return &latestReports
	}()

	return result, latestRoundNumber, nil
}

func (p kdoctorReportStorage) getNetDNSKdoctorReports(ctx context.Context, fileNameList []string) ([]*v1beta1.KdoctorReport, error) {
	var resList []*v1beta1.KdoctorReport

	netDNSList, err := p.clientSet.KdoctorV1beta1().Netdnses().List(ctx, metav1.ListOptions{})
	if nil != err {
		return nil, err
	}

	netDNSFileNameList := func() []string {
		var arr []string
		for _, fileName := range fileNameList {
			if strings.HasPrefix(fileName, v1beta1.NetDNSTaskName) {
				if strings.Contains(fileName, summary) {
					continue
				}
				arr = append(arr, fileName)
			}
		}
		sort.Strings(arr)
		return arr
	}()

	for _, netDNS := range netDNSList.Items {
		tmpNetDNS := netDNS.DeepCopy()
		if tmpNetDNS.Status.DoneRound == nil || tmpNetDNS.Status.ExpectedRound == nil {
			klog.Infof("NetDNS %s has no expectedRound or no done round", tmpNetDNS.Name)
			continue
		}

		result, latestRoundNumber, err := p.getLatestRoundReports(tmpNetDNS.Name, netDNSFileNameList)
		if nil != err {
			return nil, err
		}

		// TODO (Icarus9913): redesign this
		var taskStatus string
		if tmpNetDNS.Status.Finish {
			taskStatus = "Finished"
		} else {
			taskStatus = "NotFinished"
		}

		var finishedRoundNumber int64
		if len(tmpNetDNS.Status.History) != 0 {
			finishedRoundNumber = int64(tmpNetDNS.Status.History[0].RoundNumber)
		}

		kdoctorReportStatus := v1beta1.Status{
			ToTalRoundNumber:    *tmpNetDNS.Status.ExpectedRound,
			FinishedRoundNumber: finishedRoundNumber,
			Status:              taskStatus,
			RoundNumber:         latestRoundNumber,
		}

		kdoctorReport := &v1beta1.KdoctorReport{}
		kdoctorReport.Name = strings.ToLower(v1beta1.NetDNSTaskName) + "-" + tmpNetDNS.Name
		kdoctorReport.CreationTimestamp = tmpNetDNS.CreationTimestamp
		kdoctorReport.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1beta1.GroupName,
			Version: v1beta1.V1betaVersion,
			Kind:    v1beta1.KindKdoctorReport,
		})
		kdoctorReport.Status = kdoctorReportStatus
		kdoctorReport.Task.Spec.NetDNSTaskSpec = &netDNS.Spec
		kdoctorReport.Task.TaskName = tmpNetDNS.Name
		kdoctorReport.Task.TaskType = v1beta1.NetDNSTaskName
		kdoctorReport.Report = v1beta1.Reports{LatestRoundReport: result}
		resList = append(resList, kdoctorReport)
	}

	return resList, nil
}

func (p kdoctorReportStorage) getHttpAppHealthyReports(ctx context.Context, fileNameList []string) ([]*v1beta1.KdoctorReport, error) {
	var resList []*v1beta1.KdoctorReport

	httpAppHealthyList, err := p.clientSet.KdoctorV1beta1().AppHttpHealthies().List(ctx, metav1.ListOptions{})
	if nil != err {
		return nil, err
	}

	httpAppHealthyNameList := func() []string {
		var arr []string
		for _, fileName := range fileNameList {
			if strings.Contains(fileName, v1beta1.AppHttpHealthyTaskName) {
				if strings.Contains(fileName, summary) {
					continue
				}
				arr = append(arr, fileName)
			}
		}
		sort.Strings(arr)
		return arr
	}()

	for _, appHttpHealthy := range httpAppHealthyList.Items {
		tmpHttpAppHealthy := appHttpHealthy.DeepCopy()
		if appHttpHealthy.Status.DoneRound == nil || appHttpHealthy.Status.ExpectedRound == nil {
			klog.Infof("HttpAppHealthy %s has no expectedRound or no done round", tmpHttpAppHealthy.Name)
			continue
		}

		result, latestRoundNumber, err := p.getLatestRoundReports(tmpHttpAppHealthy.Name, httpAppHealthyNameList)
		if nil != err {
			return nil, err
		}

		// TODO (Icarus9913): redesign this
		var taskStatus string
		if tmpHttpAppHealthy.Status.Finish {
			taskStatus = "Finished"
		} else {
			taskStatus = "NotFinished"
		}

		var finishedRoundNumber int64
		if len(tmpHttpAppHealthy.Status.History) != 0 {
			finishedRoundNumber = int64(tmpHttpAppHealthy.Status.History[0].RoundNumber)
		}

		kdoctorReportStatus := v1beta1.Status{
			ToTalRoundNumber:    *tmpHttpAppHealthy.Status.ExpectedRound,
			FinishedRoundNumber: finishedRoundNumber,
			Status:              taskStatus,
			RoundNumber:         latestRoundNumber,
		}

		kdoctorReport := &v1beta1.KdoctorReport{}
		kdoctorReport.Name = strings.ToLower(v1beta1.AppHttpHealthyTaskName) + "-" + tmpHttpAppHealthy.Name
		kdoctorReport.CreationTimestamp = tmpHttpAppHealthy.CreationTimestamp
		kdoctorReport.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1beta1.GroupName,
			Version: v1beta1.V1betaVersion,
			Kind:    v1beta1.KindKdoctorReport,
		})
		kdoctorReport.Status = kdoctorReportStatus
		kdoctorReport.Report = v1beta1.Reports{LatestRoundReport: result}
		kdoctorReport.Task.Spec.AppHttpHealthyTaskSpec = &appHttpHealthy.Spec
		kdoctorReport.Task.TaskName = tmpHttpAppHealthy.Name
		kdoctorReport.Task.TaskType = v1beta1.AppHttpHealthyTaskName
		resList = append(resList, kdoctorReport)
	}

	return resList, nil
}

func (p kdoctorReportStorage) getNetReachHealthyReports(ctx context.Context, fileNameList []string) ([]*v1beta1.KdoctorReport, error) {
	var resList []*v1beta1.KdoctorReport

	netReachHealthyList, err := p.clientSet.KdoctorV1beta1().NetReaches().List(ctx, metav1.ListOptions{})
	if nil != err {
		return nil, err
	}

	netReachHealthyFileNameList := func() []string {
		var arr []string
		for _, fileName := range fileNameList {
			if strings.HasPrefix(fileName, v1beta1.NetReachTaskName) {
				if strings.Contains(fileName, summary) {
					continue
				}
				arr = append(arr, fileName)
			}
		}
		sort.Strings(arr)
		return arr
	}()

	for _, netReach := range netReachHealthyList.Items {
		tmpNetReachHealthy := netReach.DeepCopy()
		if tmpNetReachHealthy.Status.DoneRound == nil || tmpNetReachHealthy.Status.ExpectedRound == nil {
			klog.Infof("NetReachHealthy %s has no expectedRound or no done round", tmpNetReachHealthy.Name)
			continue
		}

		result, latestRoundNumber, err := p.getLatestRoundReports(tmpNetReachHealthy.Name, netReachHealthyFileNameList)
		if nil != err {
			return nil, err
		}

		// TODO (Icarus9913): redesign this
		var taskStatus string
		if tmpNetReachHealthy.Status.Finish {
			taskStatus = "Finished"
		} else {
			taskStatus = "NotFinished"
		}

		var finishedRoundNumber int64
		if len(tmpNetReachHealthy.Status.History) != 0 {
			finishedRoundNumber = int64(tmpNetReachHealthy.Status.History[0].RoundNumber)
		}

		kdoctorReportStatus := v1beta1.Status{
			ToTalRoundNumber:    *tmpNetReachHealthy.Status.ExpectedRound,
			FinishedRoundNumber: finishedRoundNumber,
			Status:              taskStatus,
			RoundNumber:         latestRoundNumber,
		}

		kdoctorReport := &v1beta1.KdoctorReport{}
		kdoctorReport.Name = strings.ToLower(v1beta1.NetReachTaskName) + "-" + tmpNetReachHealthy.Name
		kdoctorReport.CreationTimestamp = tmpNetReachHealthy.CreationTimestamp
		kdoctorReport.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1beta1.GroupName,
			Version: v1beta1.V1betaVersion,
			Kind:    v1beta1.KindKdoctorReport,
		})
		kdoctorReport.Status = kdoctorReportStatus
		kdoctorReport.Report = v1beta1.Reports{LatestRoundReport: result}
		kdoctorReport.Task.Spec.NetReachTaskSpec = &netReach.Spec
		kdoctorReport.Task.TaskName = tmpNetReachHealthy.Name
		kdoctorReport.Task.TaskType = v1beta1.NetReachTaskName
		resList = append(resList, kdoctorReport)
	}

	return resList, nil
}
