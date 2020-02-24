assume_role_from_aws_account_id="{{.assume_role_from_aws_account_id}}"

vpc_id="{{.vpc_id}}"
worker_subnet_ids=[{{ range  $i, $e := .worker_subnet_ids }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]
pod_subnet_ids=[{{ range  $i, $e := .pod_subnet_ids }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]
harmonised_subnet_ids=[{{ range  $i, $e := .harmonised_subnet_ids }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]

cluster_name="{{.cluster_name}}"
ssh_worker_key="{{.ssh_worker_key}}"

worker_subnet_proxy_host_port="{{.worker_subnet_proxy_host_port}}"

nat_enabled={{.nat_enabled}}
nat_tag_worker_subnet_key="{{.nat_tag_worker_subnet_key}}"
nat_tag_worker_subnet_value="{{.nat_tag_worker_subnet_value}}"
nat_tag_pod_subnet_key="{{.nat_tag_pod_subnet_key}}"
nat_tag_pod_subnet_value="{{.nat_tag_pod_subnet_value}}"

r53_hosted_zone_donamin_name="{{.r53_hosted_zone_donamin_name}}"
