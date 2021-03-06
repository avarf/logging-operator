package configmap

// ConfigMapContentNamespace is static content for namespace
const ConfigMapContentNamespace = `@include /fluentd/etc/conf.d/*.conf
<match fluent.**>
@type null
</match>
<source>
@type tail
path /var/log/containers/*.log
exclude_path ["/var/log/containers/*kube-system*.log", "/var/log/containers/*monitoring*.log", "/var/log/containers/*logging*.log"]
pos_file /var/log/fluentd-containers.log.pos
time_format %Y-%m-%dT%H:%M:%S.%NZ
tag kubernetes.*
format json
read_from_head false
</source>
<filter kubernetes.**>
@type kubernetes_metadata
verify_ssl false
</filter>
<filter kubernetes.**>
@type parser
key_name log
reserve_time true
reserve_data true
emit_invalid_record_to_error false
format json
<parse>
  @type json
</parse>
</filter>
<match kubernetes.**>
  @type elasticsearch_dynamic
  include_tag_key true
  logstash_format true
  # logstash_prefix kubernetes-${record['kubernetes']['pod_name']}
  logstash_prefix kubernetes-${record['kubernetes']['namespace_name']}
  host "#{ENV['FLUENT_ELASTICSEARCH_HOST']}"
  port "#{ENV['FLUENT_ELASTICSEARCH_PORT']}"
  scheme "#{ENV['FLUENT_ELASTICSEARCH_SCHEME'] || 'http'}"
  ssl_verify "#{ENV['FLUENT_ELASTICSEARCH_SSL_VERIFY'] || 'true'}"
  user "#{ENV['FLUENT_ELASTICSEARCH_USER']}"
  password "#{ENV['FLUENT_ELASTICSEARCH_PASSWORD']}"
  reload_connections false
  reconnect_on_error true
  reload_on_failure true
  <buffer>
	  flush_thread_count 16
	  flush_interval 5s
	  chunk_limit_size 2M
	  queue_limit_length 32
	  retry_max_interval 30
	  retry_forever true
  </buffer>
</match>
`

// ConfigMapContentPod is static content for pod
const ConfigMapContentPod = `@include /fluentd/etc/conf.d/*.conf
<match fluent.**>
@type null
</match>
<source>
@type tail
path /var/log/containers/*.log
exclude_path ["/var/log/containers/*kube-system*.log", "/var/log/containers/*monitoring*.log", "/var/log/containers/*logging*.log"]
pos_file /var/log/fluentd-containers.log.pos
time_format %Y-%m-%dT%H:%M:%S.%NZ
tag kubernetes.*
format json
read_from_head false
</source>
<filter kubernetes.**>
@type kubernetes_metadata
verify_ssl false
</filter>
<filter kubernetes.**>
@type parser
key_name log
reserve_time true
reserve_data true
emit_invalid_record_to_error false
format json
<parse>
  @type json
</parse>
</filter>
<match kubernetes.**>
  @type elasticsearch_dynamic
  include_tag_key true
  logstash_format true
  # logstash_prefix kubernetes-${record['kubernetes']['pod_name']}
  logstash_prefix kubernetes-${record['kubernetes']['namespace_name']}
  host "#{ENV['FLUENT_ELASTICSEARCH_HOST']}"
  port "#{ENV['FLUENT_ELASTICSEARCH_PORT']}"
  scheme "#{ENV['FLUENT_ELASTICSEARCH_SCHEME'] || 'http'}"
  ssl_verify "#{ENV['FLUENT_ELASTICSEARCH_SSL_VERIFY'] || 'true'}"
  user "#{ENV['FLUENT_ELASTICSEARCH_USER']}"
  password "#{ENV['FLUENT_ELASTICSEARCH_PASSWORD']}"
  reload_connections false
  reconnect_on_error true
  reload_on_failure true
  <buffer>
	  flush_thread_count 16
	  flush_interval 5s
	  chunk_limit_size 2M
	  queue_limit_length 32
	  retry_max_interval 30
	  retry_forever true
  </buffer>
</match>
`

const kiabanConfigTLSData = `server.name: kibana
server.host: "0"
monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.ssl.verificationMode: none
elasticsearch.hosts: ['${ELASTICSEARCH_HOST:elasticsearch-master-headless}:${ELASTICSEARCH_PORT:9200}']
elasticsearch.username: ${ELASTICSEARCH_USERNAME}
elasticsearch.password: ${ELASTICSEARCH_PASSWORD}
`

const kiabanConfigData = `server.name: kibana
server.host: "0"
`
