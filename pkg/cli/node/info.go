package node

import (
	api "github.com/fabiorphp/kongo"
	"github.com/fabiorphp/kongo-cli/pkg/template"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	nodeInfoTmpl = `Configuration:
  Admin:
    AcessLog: {{ .Configuration.AdminAcessLog }}
    ErrorLog: {{ .Configuration.AdminErrorLog }}
    Listen:{{ range .Configuration.AdminListen }} {{ . }}{{ end }}
    Listeners:{{ range .Configuration.AdminListeners }}
      - SSL: {{ .SSL }}, IP: {{ .Ip }}, Protocol: {{ .Protocol }}, Port: {{ .Port }}, HTTP2: {{ .Http2 }}, Listener: {{ .Listener }}{{ end }}
    SSLCertificateDefault: {{ .Configuration.AdminSSLCertificateDefault }}
    SSLCertificateCsrDefault: {{ .Configuration.AdminSSLCertificateCsrDefault }}
    SSLCertificateKeyDefault: {{ .Configuration.AdminSSLCertificateKeyDefault }}
    SSLEnabled: {{ .Configuration.AdminSSLEnabled }}
  Cassandra:
    Consistency: {{ .Configuration.CassandraConsistency }}
    ContactPoints:{{ range .Configuration.CassandraContactPoints }} {{ . }}{{ end }}
    DataCenters:{{ range .Configuration.CassandraDataCenters }} {{ . }}{{ end }}
    Keyspace: {{ .Configuration.CassandraKeyspace }}
    LBPolicy: {{ .Configuration.CassandraLBPolicy }}
    Port: {{ .Configuration.CassandraPort }}
    ReplicationFactor: {{ .Configuration.CassandraReplicationFactor }}
    ReplicationStrategy: {{ .Configuration.CassandraReplicationStrategy }}
    SchemaConsensusTimeout: {{ .Configuration.CassandraSchemaConsensusTimeout }}
    SSL: {{ .Configuration.CassandraSSL }}
    SSLVerify: {{ .Configuration.CassandraSSLVerify }}
    Timeout: {{ .Configuration.CassandraTimeout }}
    Username: {{ .Configuration.CassandraUsername }}
  Client:
    BodyBufferSize: {{ .Configuration.ClientBodyBufferSize }}
    MaxBodySize: {{ .Configuration.ClientMaxBodySize }}
    SSL: {{ .Configuration.ClientSSL }}
    SSLCertificateCsrDefault: {{ .Configuration.ClientSSLCertificateCsrDefault }}
    SSLCertificateDefault: {{ .Configuration.ClientSSLCertificateDefault }}
    SSLCertificateKeyDefault: {{ .Configuration.ClientSSLCertificateKeyDefault }}
  CustomPlugins:{{ range .Configuration.CustomPlugins }} {{ . }}{{ end }}
  Database:
    Name: {{ .Configuration.Database }}
    CacheTTL: {{ .Configuration.DatabaseCacheTTL }}
    UpdateFrequency: {{ .Configuration.DatabaseUpdateFrequency }}
    UpdatePropagation: {{ .Configuration.DatabaseUpdatePropagation }}
  DNS:
    ErrorTTL: {{ .Configuration.DNSErrorTTL }}
    HostsFile: {{ .Configuration.DNSHostsFile }}
    NotFoundTTL: {{ .Configuration.DNSNotFoundTTL }}
    NoSync: {{ .Configuration.DNSNoSync }}
    Order:{{ range .Configuration.DNSOrder }} {{ . }}{{ end }}
    Resolver:{{ range .Configuration.DNSResolver }} {{ . }}{{ end }}
    StaleTTL: {{ .Configuration.DNSStaleTTL }}
  ErrorDefaultType: {{ .Configuration.ErrorDefaultType }}
  KongEnv: {{ .Configuration.KongEnv }}
  LatencyTokens: {{ .Configuration.LatencyTokens }}
  Lua:
    PackageCPath: {{ .Configuration.LuaPackageCPath }}
    PackagePath: {{ .Configuration.LuaPackagePath }}
    SocketPoolSize: {{ .Configuration.LuaSocketPoolSize }}
    SSLVerifyDepth: {{ .Configuration.LuaSSLVerifyDepth }}
  LogLevel: {{ .Configuration.LogLevel }}
  MemoryCacheSize: {{ .Configuration.MemoryCacheSize }}
  Nginx:
    AccessLogs: {{ .Configuration.NginxAccessLogs }}
    AdminAccessLog: {{ .Configuration.NginxAdminAccessLog }}
    Conf: {{ .Configuration.NginxConf }}
    Daemon: {{ .Configuration.NginxDaemon }}
    ErrorLogs: {{ .Configuration.NginxErrorLogs }}
    KongConf: {{ .Configuration.NginxKongConf }}
    Optimizations: {{ .Configuration.NginxOptimizations }}
    PID: {{ .Configuration.NginxPID }}
    WorkerProcesses: {{ .Configuration.NginxWorkerProcesses }}
  Plugins:{{ range $key, $value := .Configuration.Plugins }} {{ $key }}{{ end }}
  Postgres:
    Database: {{ .Configuration.PostgresDatabase }}
    Host: {{ .Configuration.PostgresHost }}
    Port: {{ .Configuration.PostgresPort }}
    SSL: {{ .Configuration.PostgresSSL }}
    Username: {{ .Configuration.PostgresUsername }}
    SSLVerify: {{ .Configuration.PostgresSSLVerify }}
  Prefix: {{ .Configuration.Prefix }}
  Proxy:
    AccessLog: {{ .Configuration.ProxyAccessLog }}
    ErrorLog: {{ .Configuration.ProxyErrorLog }}
    Listen:{{ range .Configuration.ProxyListen }} {{ . }}{{ end }}
    Listeners:{{ range .Configuration.ProxyListeners }}
      - SSL: {{ .SSL }}, IP: {{ .Ip }}, Protocol: {{ .Protocol }}, Port: {{ .Port }}, HTTP2: {{ .Http2 }}, Listener: {{ .Listener }}{{ end }}
    SSLEnabled: {{ .Configuration.ProxySSLEnabled }}
  Real IP:
    Header: {{ .Configuration.RealIpHeader }}
    Recursive: {{ .Configuration.RealIpRecursive }}
  ServerTokens: {{ .Configuration.ServerTokens }}
  SSL:
    Certificate: {{ .Configuration.SSLCertificate }}
    CertificateDefault: {{ .Configuration.SSLCertificateDefault }}
    CertificateKey: {{ .Configuration.SSLCertificateKey }}
    CertificateDefaultKey: {{ .Configuration.SSLCertificateDefaultKey }}
    CertificateCsrDefault: {{ .Configuration.SSLCertificateCsrDefault }}
    Ciphers: {{ .Configuration.SSLCiphers }}
    CipherSuite: {{ .Configuration.SSLCipherSuite }}
  TrustedIps:{{ range .Configuration.TrustedIps }} {{ . }}{{ end }}
  UpstreamKeepAlive: {{ .Configuration.UpstreamKeepAlive }}
Hostname: {{ .Hostname }}
Lua version: {{ .LuaVersion }}
Plugins:
  Available on server:{{ range $key, $value  := .Plugins.AvailableOnServer }} {{ $key }}{{ end }}
  Enabled in cluster:{{ range .Plugins.EnabledInCluster }} {{ . }}{{ end }}
Prng seeds: {{ range $key, $value := .PrngSeeds }}
  {{ $key }}: {{ $value }}{{ end }}
Tagline: {{ .Tagline }}
Timers:
  Pending: {{ .Timers.Pending }}
  Running: {{ .Timers.Running }}
Version: {{ .Version }}`

	// ErrNodeInfo retrieves an error message when client api fails.
	ErrNodeInfo = "Unable to retrieve the node information"
)

// Info retrieves the information about the server node.
func Info(c *cli.Context) error {
	client := c.App.Metadata["client"].(*api.Kongo)
	info, _, err := client.Node.Info()

	if err != nil {
		return errors.Wrap(err, ErrNodeInfo)
	}

	tmpl := template.NewPlain(nodeInfoTmpl)

	return tmpl.Write(c.App.Writer, info)
}
