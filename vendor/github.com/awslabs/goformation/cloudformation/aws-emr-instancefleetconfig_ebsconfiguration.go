package cloudformation

import (
	"encoding/json"
)

// AWSEMRInstanceFleetConfig_EbsConfiguration AWS CloudFormation Resource (AWS::EMR::InstanceFleetConfig.EbsConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html
type AWSEMRInstanceFleetConfig_EbsConfiguration struct {

	// EbsBlockDeviceConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsblockdeviceconfigs
	EbsBlockDeviceConfigs []AWSEMRInstanceFleetConfig_EbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html#cfn-elasticmapreduce-instancefleetconfig-ebsconfiguration-ebsoptimized
	EbsOptimized *Value `json:"EbsOptimized,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig_EbsConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig.EbsConfiguration"
}

func (r *AWSEMRInstanceFleetConfig_EbsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
