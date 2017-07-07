// Code generated by counterfeiter. DO NOT EDIT.
package datadogfakes

import (
	"cred-alert/datadog"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeClient struct {
	PublishSeriesStub        func(logger lager.Logger, series datadog.Series)
	publishSeriesMutex       sync.RWMutex
	publishSeriesArgsForCall []struct {
		logger lager.Logger
		series datadog.Series
	}
	BuildMetricStub        func(metricType string, metricName string, count float32, tags ...string) datadog.Metric
	buildMetricMutex       sync.RWMutex
	buildMetricArgsForCall []struct {
		metricType string
		metricName string
		count      float32
		tags       []string
	}
	buildMetricReturns struct {
		result1 datadog.Metric
	}
	buildMetricReturnsOnCall map[int]struct {
		result1 datadog.Metric
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) PublishSeries(logger lager.Logger, series datadog.Series) {
	fake.publishSeriesMutex.Lock()
	fake.publishSeriesArgsForCall = append(fake.publishSeriesArgsForCall, struct {
		logger lager.Logger
		series datadog.Series
	}{logger, series})
	fake.recordInvocation("PublishSeries", []interface{}{logger, series})
	fake.publishSeriesMutex.Unlock()
	if fake.PublishSeriesStub != nil {
		fake.PublishSeriesStub(logger, series)
	}
}

func (fake *FakeClient) PublishSeriesCallCount() int {
	fake.publishSeriesMutex.RLock()
	defer fake.publishSeriesMutex.RUnlock()
	return len(fake.publishSeriesArgsForCall)
}

func (fake *FakeClient) PublishSeriesArgsForCall(i int) (lager.Logger, datadog.Series) {
	fake.publishSeriesMutex.RLock()
	defer fake.publishSeriesMutex.RUnlock()
	return fake.publishSeriesArgsForCall[i].logger, fake.publishSeriesArgsForCall[i].series
}

func (fake *FakeClient) BuildMetric(metricType string, metricName string, count float32, tags ...string) datadog.Metric {
	fake.buildMetricMutex.Lock()
	ret, specificReturn := fake.buildMetricReturnsOnCall[len(fake.buildMetricArgsForCall)]
	fake.buildMetricArgsForCall = append(fake.buildMetricArgsForCall, struct {
		metricType string
		metricName string
		count      float32
		tags       []string
	}{metricType, metricName, count, tags})
	fake.recordInvocation("BuildMetric", []interface{}{metricType, metricName, count, tags})
	fake.buildMetricMutex.Unlock()
	if fake.BuildMetricStub != nil {
		return fake.BuildMetricStub(metricType, metricName, count, tags...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.buildMetricReturns.result1
}

func (fake *FakeClient) BuildMetricCallCount() int {
	fake.buildMetricMutex.RLock()
	defer fake.buildMetricMutex.RUnlock()
	return len(fake.buildMetricArgsForCall)
}

func (fake *FakeClient) BuildMetricArgsForCall(i int) (string, string, float32, []string) {
	fake.buildMetricMutex.RLock()
	defer fake.buildMetricMutex.RUnlock()
	return fake.buildMetricArgsForCall[i].metricType, fake.buildMetricArgsForCall[i].metricName, fake.buildMetricArgsForCall[i].count, fake.buildMetricArgsForCall[i].tags
}

func (fake *FakeClient) BuildMetricReturns(result1 datadog.Metric) {
	fake.BuildMetricStub = nil
	fake.buildMetricReturns = struct {
		result1 datadog.Metric
	}{result1}
}

func (fake *FakeClient) BuildMetricReturnsOnCall(i int, result1 datadog.Metric) {
	fake.BuildMetricStub = nil
	if fake.buildMetricReturnsOnCall == nil {
		fake.buildMetricReturnsOnCall = make(map[int]struct {
			result1 datadog.Metric
		})
	}
	fake.buildMetricReturnsOnCall[i] = struct {
		result1 datadog.Metric
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publishSeriesMutex.RLock()
	defer fake.publishSeriesMutex.RUnlock()
	fake.buildMetricMutex.RLock()
	defer fake.buildMetricMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ datadog.Client = new(FakeClient)
