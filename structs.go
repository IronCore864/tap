package main

import "github.com/spf13/viper"

// TerraformTfvars is the struct used for the templating for terraform.tfvars file
type TerraformTfvars struct {
	// account
	AssumeRoleFromAWSAccountID string
	// networking
	VPCID               string
	WorkerSubnetIDs     []string
	PodSubnetIDs        []string
	HarmonisedSubnetIDs []string
	// cluster
	ClusterName  string
	SSHWorkerKey string
	// proxy
	WorkerSubnetProxyHostPort string
	R53HostedZoneDomainName   string
	// nat
	NATEnabled              bool
	NATTagWorkerSubnetKey   string
	NATTagWorkerSubnetValue string
	NATTagPodSubnetKey      string
	NATTagPodSubnetValue    string
}

// ConfigTf is the struct used for the templating for terraform.tfvars file
type ConfigTf struct {
	Region            string
	TFStateBucketName string
	TFStateFileName   string
}

func buildTerraformTfvarsStruct() *TerraformTfvars {
	s := TerraformTfvars{
		viper.GetString("assume_role_from_aws_account_id"),
		viper.GetString("vpc_id"),
		viper.GetStringSlice("worker_subnet_ids"),
		viper.GetStringSlice("pod_subnet_ids"),
		viper.GetStringSlice("harmonised_subnet_ids"),
		viper.GetString("cluster_name"),
		viper.GetString("ssh_worker_key"),
		viper.GetString("worker_subnet_proxy_host_port"),
		viper.GetString("r53_hosted_zone_donamin_name"),
		viper.GetBool("nat_enabled"),
		viper.GetString("nat_tag_worker_subnet_key"),
		viper.GetString("nat_tag_worker_subnet_value"),
		viper.GetString("nat_tag_pod_subnet_key"),
		viper.GetString("nat_tag_pod_subnet_value"),
	}
	return &s
}

func buildConfigTfStruct() *ConfigTf {
	s := ConfigTf{
		viper.GetString("region"),
		viper.GetString("s3_bucket"),
		viper.GetString("s3_key"),
	}
	return &s
}
