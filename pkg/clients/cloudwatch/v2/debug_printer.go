package v2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"

	"github.com/prometheus-community/yet-another-cloudwatch-exporter/pkg/model"
)

// getMetricDataOutputToString converts a GetMetricDataOutput to a string for debugging purposes.
func getMetricDataOutputToString(resp cloudwatch.GetMetricDataOutput) *strings.Builder {
	sb := new(strings.Builder)
	sb.WriteString("{")
	sb.WriteString("MetricDataResults: [")
	for i, mdr := range resp.MetricDataResults {
		sb.WriteString("{")
		sb.WriteString("Id: " + aws.ToString(mdr.Id) + ", ")
		sb.WriteString("Timestamps: [")
		for _, ts := range mdr.Timestamps {
			sb.WriteString(ts.String() + ", ")
		}
		sb.WriteString("], ")
		sb.WriteString("Values: [")
		for _, val := range mdr.Values {
			sb.WriteString(fmt.Sprintf("%f", val) + ", ")
		}
		sb.WriteString("]")
		sb.WriteString("}, ")
		if i < len(resp.MetricDataResults)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(" }")
	return sb
}

// getMetricDataInputToString converts a GetMetricDataInput to a string for debugging purposes.
func getMetricDataInputToString(input *cloudwatch.GetMetricDataInput) *strings.Builder {
	sb := new(strings.Builder)
	sb.WriteString("{")
	sb.WriteString("StartTime: " + input.StartTime.String() + ", ")
	sb.WriteString("EndTime: " + input.EndTime.String() + ", ")
	sb.WriteString("MetricDataQueries: [")
	for i, mdq := range input.MetricDataQueries {
		sb.WriteString("{")
		sb.WriteString("Id: " + aws.ToString(mdq.Id) + ", ")
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
			sb.WriteString("Stat: " + aws.ToString(mdq.MetricStat.Stat))
			sb.WriteString("}, ")
		}
		sb.WriteString("ReturnData: " + fmt.Sprintf("%t", aws.ToBool(mdq.ReturnData)))
		sb.WriteString("}, ")
		if i < len(input.MetricDataQueries)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	sb.WriteString(" }")
	return sb
}

// listMetricsInputToString converts a ListMetricsInput to a string for debugging purposes.
func listMetricsInputToString(filter *cloudwatch.ListMetricsInput) *strings.Builder {
	sb := new(strings.Builder)
	sb.WriteString("{")
	sb.WriteString("Namespace: " + aws.ToString(filter.Namespace) + ", ")
	sb.WriteString("MetricName: " + aws.ToString(filter.MetricName))
	if filter.RecentlyActive != "" {
		sb.WriteString(", RecentlyActive: " + string(filter.RecentlyActive))
	}
	sb.WriteString(" }")
	return sb
}

// metricToString converts a slice of model.Metric to a string for debugging purposes.
func metricToString(metricsPage []*model.Metric) *strings.Builder {
	sb := new(strings.Builder)
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
	return sb
}
