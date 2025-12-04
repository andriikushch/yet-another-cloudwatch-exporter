package v2

import (
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"

	"github.com/prometheus-community/yet-another-cloudwatch-exporter/pkg/model"
)

// defaultGrowSize is the initial size to grow the strings.Builder to. The value is arbitrary but should help reduce allocations.
const defaultGrowSize = 512

var debugBuilderPool = sync.Pool{
	New: func() interface{} {
		sb := new(strings.Builder)
		sb.Grow(defaultGrowSize)
		return sb
	},
}

// getMetricDataOutputToString converts a GetMetricDataOutput to a string for debugging purposes.
func getMetricDataOutputToString(resp cloudwatch.GetMetricDataOutput) string {
	sb := debugBuilderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		debugBuilderPool.Put(sb)
	}()

	sb.WriteString("{")
	sb.WriteString("NextToken: " + aws.ToString(resp.NextToken) + ", ")
	sb.WriteString("MetricDataResults: [")
	for i, mdr := range resp.MetricDataResults {
		sb.WriteString("{")
		sb.WriteString("Id: " + aws.ToString(mdr.Id) + ", ")
		sb.WriteString("Label: " + aws.ToString(mdr.Label) + ", ")
		sb.WriteString("StatusCode: " + string(mdr.StatusCode) + ", ")
		sb.WriteString("Timestamps: [")
		for _, ts := range mdr.Timestamps {
			sb.WriteString(ts.String() + ", ")
		}
		sb.WriteString("], ")
		sb.WriteString("Values: [")
		for _, val := range mdr.Values {
			sb.WriteString(fmt.Sprintf("%f", val) + ", ")
		}
		sb.WriteString("], ")
		sb.WriteString("Messages: [")
		for _, msg := range mdr.Messages {
			sb.WriteString("{Code: " + aws.ToString(msg.Code) + ", Value: " + aws.ToString(msg.Value) + "}, ")
		}
		sb.WriteString("]")
		sb.WriteString("}")
		if i < len(resp.MetricDataResults)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(" }")
	return sb.String()
}

// getMetricDataInputToString converts a GetMetricDataInput to a string for debugging purposes.
func getMetricDataInputToString(input *cloudwatch.GetMetricDataInput) string {
	sb := debugBuilderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		debugBuilderPool.Put(sb)
	}()

	sb.WriteString("{")
	sb.WriteString("StartTime: " + input.StartTime.String() + ", ")
	sb.WriteString("EndTime: " + input.EndTime.String() + ", ")
	sb.WriteString("NextToken: " + aws.ToString(input.NextToken) + ", ")
	sb.WriteString("ScanBy: " + string(input.ScanBy) + ", ")
	sb.WriteString("MaxDatapoints: " + fmt.Sprintf("%d", aws.ToInt32(input.MaxDatapoints)) + ", ")
	if input.LabelOptions != nil {
		sb.WriteString("LabelOptions: {Timezone: " + aws.ToString(input.LabelOptions.Timezone) + "}, ")
	}
	sb.WriteString("MetricDataQueries: [")
	for i, mdq := range input.MetricDataQueries {
		sb.WriteString("{")
		sb.WriteString("Id: " + aws.ToString(mdq.Id) + ", ")
		sb.WriteString("Label: " + aws.ToString(mdq.Label) + ", ")
		sb.WriteString("Expression: " + aws.ToString(mdq.Expression) + ", ")
		sb.WriteString("Period: " + fmt.Sprintf("%d", aws.ToInt32(mdq.Period)) + ", ")
		sb.WriteString("AccountId: " + aws.ToString(mdq.AccountId) + ", ")
		if mdq.MetricStat != nil {
			sb.WriteString("MetricStat: {")
			if mdq.MetricStat.Metric != nil {
				sb.WriteString("Metric: {")
				sb.WriteString("Namespace: " + aws.ToString(mdq.MetricStat.Metric.Namespace) + ", ")
				sb.WriteString("MetricName: " + aws.ToString(mdq.MetricStat.Metric.MetricName) + ", ")
				sb.WriteString("Dimensions: [")
				for _, dim := range mdq.MetricStat.Metric.Dimensions {
					sb.WriteString("{")
					sb.WriteString("Name: " + aws.ToString(dim.Name) + ", ")
					sb.WriteString("Value: " + aws.ToString(dim.Value))
					sb.WriteString("}, ")
				}
				sb.WriteString("], ")
				sb.WriteString("}, ")
			}
			sb.WriteString("Period: " + fmt.Sprintf("%d", aws.ToInt32(mdq.MetricStat.Period)) + ", ")
			sb.WriteString("Stat: " + aws.ToString(mdq.MetricStat.Stat) + ", ")
			sb.WriteString("Unit: " + string(mdq.MetricStat.Unit))
			sb.WriteString("}, ")
		}
		sb.WriteString("ReturnData: " + fmt.Sprintf("%t", aws.ToBool(mdq.ReturnData)))
		sb.WriteString("}")
		if i < len(input.MetricDataQueries)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(" }")
	return sb.String()
}

// listMetricsInputToString converts a ListMetricsInput to a string for debugging purposes.
func listMetricsInputToString(filter *cloudwatch.ListMetricsInput) string {
	sb := debugBuilderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		debugBuilderPool.Put(sb)
	}()

	sb.WriteString("{")
	sb.WriteString("Namespace: " + aws.ToString(filter.Namespace) + ", ")
	sb.WriteString("MetricName: " + aws.ToString(filter.MetricName))
	if filter.RecentlyActive != "" {
		sb.WriteString(", RecentlyActive: " + string(filter.RecentlyActive))
	}
	sb.WriteString(", NextToken: " + aws.ToString(filter.NextToken))
	sb.WriteString(", OwningAccount: " + aws.ToString(filter.OwningAccount))
	sb.WriteString(", IncludeLinkedAccounts: " + fmt.Sprintf("%t", aws.ToBool(filter.IncludeLinkedAccounts)))
	if len(filter.Dimensions) > 0 {
		sb.WriteString(", Dimensions: [")
		for _, dim := range filter.Dimensions {
			sb.WriteString("{Name: " + aws.ToString(dim.Name) + ", Value: " + aws.ToString(dim.Value) + "}, ")
		}
		sb.WriteString("]")
	}
	sb.WriteString(" }")
	return sb.String()
}

// metricToString converts a slice of model.Metric to a string for debugging purposes.
func metricToString(metricsPage []*model.Metric) string {
	sb := debugBuilderPool.Get().(*strings.Builder)
	defer func() {
		sb.Reset()
		debugBuilderPool.Put(sb)
	}()

	sb.WriteString("{")
	sb.WriteString("Metrics: [")
	for i, metric := range metricsPage {
		sb.WriteString("{")
		sb.WriteString("MetricName: " + metric.MetricName + ", ")
		sb.WriteString("Namespace: " + metric.Namespace + ", ")
		sb.WriteString("Dimensions: [")
		for _, dim := range metric.Dimensions {
			sb.WriteString("{")
			sb.WriteString("Name: " + dim.Name + ", ")
			sb.WriteString("Value: " + dim.Value)
			sb.WriteString("}, ")
		}
		sb.WriteString("]")
		sb.WriteString("}, ")
		if i < len(metricsPage)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(" }")
	return sb.String()
}
