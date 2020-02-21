assume_role_from_aws_account_id="{{.AssumeRoleFromAWSAccountID}}"

vpc_id="{{.VPCID}}"
worker_subnet_ids=[{{ range  $i, $e := .WorkerSubnetIDs }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]
pod_subnet_ids=[{{ range  $i, $e := .PodSubnetIDs }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]
harmonised_subnet_ids=[{{ range  $i, $e := .HarmonisedSubnetIDs }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]

ssh_worker_key="{{.SSHWorkerKey}}"
cluster_name="{{.ClusterName}}"

worker_subnet_proxy_host_port="{{.WorkerSubnetProxyHostPort}}"

nat_enabled={{.NATEnabled}}
nat_tag_worker_subnet_key="{{.NATTagWorkerSubnetKey}}"
nat_tag_worker_subnet_value="{{.NATTagWorkerSubnetValue}}"
nat_tag_pod_subnet_key="{{.NATTagPodSubnetKey}}"
nat_tag_pod_subnet_value="{{.NATTagPodSubnetValue}}"

r53_hosted_zone_donamin_name="{{.R53HostedZoneDomainName}}"
