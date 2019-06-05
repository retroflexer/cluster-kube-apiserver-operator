// Code generated by go-bindata.
// sources:
// bindata/v3.11.0/kube-apiserver/cm.yaml
// bindata/v3.11.0/kube-apiserver/defaultconfig.yaml
// bindata/v3.11.0/kube-apiserver/kubeconfig-cm.yaml
// bindata/v3.11.0/kube-apiserver/ns.yaml
// bindata/v3.11.0/kube-apiserver/pod-cm.yaml
// bindata/v3.11.0/kube-apiserver/pod.yaml
// bindata/v3.11.0/kube-apiserver/recovery-config.yaml
// bindata/v3.11.0/kube-apiserver/recovery-pod.yaml
// bindata/v3.11.0/kube-apiserver/svc.yaml
// DO NOT EDIT!

package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v3110KubeApiserverCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-apiserver
  name: config
data:
  config.yaml:
`)

func v3110KubeApiserverCmYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverCmYaml, nil
}

func v3110KubeApiserverCmYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverDefaultconfigYaml = []byte(`apiVersion: kubecontrolplane.config.openshift.io/v1
kind: KubeAPIServerConfig
aggregatorConfig:
  proxyClientInfo:
    certFile: /etc/kubernetes/static-pod-certs/secrets/aggregator-client/tls.crt
    keyFile: /etc/kubernetes/static-pod-certs/secrets/aggregator-client/tls.key
apiServerArguments:
  storage-backend:
  - etcd3
  storage-media-type:
  - application/vnd.kubernetes.protobuf
  # switch to direct pod IP routing for aggregated apiservers to avoid service IPs as on source of instability
  enable-aggregator-routing:
  - "true"
  minimal-shutdown-duration:
  - 70s # give SDN some time to converge: 30s for iptable lock contention, 25s for the second try and some seconds for AWS to update ELBs
  http2-max-streams-per-connection:
  - "2000"  # recommended is 1000, but we need to mitigate https://github.com/kubernetes/kubernetes/issues/74412
auditConfig:
  auditFilePath: "/var/log/kube-apiserver/audit.log"
  enabled: true
  logFormat: "json"
  maximumFileSizeMegabytes: 100
  maximumRetainedFiles: 10
  policyConfiguration:
    apiVersion: audit.k8s.io/v1beta1
    kind: Policy
    # Don't generate audit events for all requests in RequestReceived stage.
    omitStages:
    - "RequestReceived"
    rules:
    # Don't log requests for events
    - level: None
      resources:
      - group: ""
        resources: ["events"]
    # Don't log authenticated requests to certain non-resource URL paths.
    - level: None
      userGroups: ["system:authenticated", "system:unauthenticated"]
      nonResourceURLs:
      - "/api*" # Wildcard matching.
      - "/version"
      - "/healthz"
      - "/readyz"
    # A catch-all rule to log all other requests at the Metadata level.
    - level: Metadata
      # Long-running requests like watches that fall under this rule will not
      # generate an audit event in RequestReceived.
      omitStages:
      - "RequestReceived"
authConfig:
  oauthMetadataFile: ""
  requestHeader:
    clientCA: /etc/kubernetes/static-pod-certs/configmaps/aggregator-client-ca/ca-bundle.crt
    clientCommonNames:
    - kube-apiserver-proxy
    - system:kube-apiserver-proxy
    - system:openshift-aggregator
    extraHeaderPrefixes:
    - X-Remote-Extra-
    groupHeaders:
    - X-Remote-Group
    usernameHeaders:
    - X-Remote-User
  webhookTokenAuthenticators: null
consolePublicURL: ""
corsAllowedOrigins:
- //127\.0\.0\.1(:|$)
- //localhost(:|$)
kubeletClientInfo:
  ca: /etc/kubernetes/static-pod-resources/configmaps/kubelet-serving-ca/ca-bundle.crt
  certFile: /etc/kubernetes/static-pod-resources/secrets/kubelet-client/tls.crt
  keyFile: /etc/kubernetes/static-pod-resources/secrets/kubelet-client/tls.key
  port: 10250
projectConfig:
  defaultNodeSelector: ""
servicesNodePortRange: 30000-32767
servicesSubnet: 10.3.0.0/16 # ServiceCIDR
servingInfo:
  bindAddress: 0.0.0.0:6443
  bindNetwork: tcp4
  certFile: /etc/kubernetes/static-pod-certs/secrets/service-network-serving-certkey/tls.crt
  clientCA: /etc/kubernetes/static-pod-certs/configmaps/client-ca/ca-bundle.crt
  keyFile: /etc/kubernetes/static-pod-certs/secrets/service-network-serving-certkey/tls.key
  maxRequestsInFlight: 1200
  namedCertificates: null
  requestTimeoutSeconds: 3600
serviceAccountPublicKeyFiles:
  - /etc/kubernetes/static-pod-resources/configmaps/sa-token-signing-certs
storageConfig:
  ca: /etc/kubernetes/static-pod-resources/configmaps/etcd-serving-ca/ca-bundle.crt
  certFile: /etc/kubernetes/static-pod-resources/secrets/etcd-client/tls.crt
  keyFile: /etc/kubernetes/static-pod-resources/secrets/etcd-client/tls.key
  urls: null
userAgentMatchingConfig:
  defaultRejectionMessage: ""
  deniedClients: null
  requiredClients: null
`)

func v3110KubeApiserverDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverDefaultconfigYaml, nil
}

func v3110KubeApiserverDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverKubeconfigCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-apiserver-cert-syncer-kubeconfig
  namespace: openshift-kube-apiserver
data:
  kubeconfig: |
    apiVersion: v1
    clusters:
      - cluster:
          certificate-authority: /etc/kubernetes/static-pod-resources/configmaps/kube-apiserver-server-ca/ca-bundle.crt
          server: https://localhost:6443
        name: loopback
    contexts:
      - context:
          cluster: loopback
          user: kube-apiserver-cert-syncer
        name: kube-apiserver-cert-syncer
    current-context: kube-apiserver-cert-syncer
    kind: Config
    preferences: {}
    users:
      - name: kube-apiserver-cert-syncer
        user:
          client-certificate: /etc/kubernetes/static-pod-resources/secrets/kube-apiserver-cert-syncer-client-cert-key/tls.crt
          client-key: /etc/kubernetes/static-pod-resources/secrets/kube-apiserver-cert-syncer-client-cert-key/tls.key
`)

func v3110KubeApiserverKubeconfigCmYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverKubeconfigCmYaml, nil
}

func v3110KubeApiserverKubeconfigCmYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverKubeconfigCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/kubeconfig-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-kube-apiserver
  annotations:
    openshift.io/node-selector: ""
  labels:
    openshift.io/run-level: "0"
`)

func v3110KubeApiserverNsYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverNsYaml, nil
}

func v3110KubeApiserverNsYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverPodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-apiserver
  name: kube-apiserver-pod
data:
  pod.yaml:
  forceRedeploymentReason:
  version:
`)

func v3110KubeApiserverPodCmYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverPodCmYaml, nil
}

func v3110KubeApiserverPodCmYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverPodCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  namespace: openshift-kube-apiserver
  name: kube-apiserver
  labels:
    app: openshift-kube-apiserver
    apiserver: "true"
    revision: "REVISION"
spec:
  initContainers:
    - name: setup
      terminationMessagePolicy: FallbackToLogsOnError
      image: ${IMAGE}
      imagePullPolicy: IfNotPresent
      volumeMounts:
        - mountPath: /var/log/kube-apiserver
          name: audit-dir
      command: ['/usr/bin/timeout', '105', '/bin/bash', '-ec'] # a bit more than 60s for graceful termination + 35s for minimum-termination-duration, 5s extra cri-o's graceful termination period
      args:
      - |
        echo -n "Fixing audit permissions."
        chmod 0700 /var/log/kube-apiserver
        echo -n "Waiting for port :6443 to be released."
        while [ -n "$(lsof -ni :6443)" ]; do
          echo -n "."
          sleep 1
        done
  containers:
  - name: kube-apiserver-REVISION
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["hypershift", "openshift-kube-apiserver"]
    args:
    - --config=/etc/kubernetes/static-pod-resources/configmaps/config/config.yaml
    resources:
      requests:
        memory: 1Gi
        cpu: 150m
    ports:
    - containerPort: 6443
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
    - mountPath: /etc/kubernetes/static-pod-certs
      name: cert-dir
    - mountPath: /var/log/kube-apiserver
      name: audit-dir
    livenessProbe:
      httpGet:
        scheme: HTTPS
        port: 6443
        path: healthz
      initialDelaySeconds: 45
      timeoutSeconds: 10
    readinessProbe:
      httpGet:
        scheme: HTTPS
        port: 6443
        path: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 10
    env:
      - name: POD_NAME
        valueFrom:
          fieldRef:
            fieldPath: metadata.name
      - name: STATIC_POD_VERSION # Avoid using 'REVISION' here otherwise it will be substituted
        value: REVISION
  - name: kube-apiserver-cert-syncer-REVISION
    env:
    - name: POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: POD_NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
    image: ${OPERATOR_IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["cluster-kube-apiserver-operator", "cert-syncer"]
    args:
      - --kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/kube-apiserver-cert-syncer-kubeconfig/kubeconfig
      - --namespace=$(POD_NAMESPACE)
      - --destination-dir=/etc/kubernetes/static-pod-certs
    resources:
      requests:
        memory: 50Mi
        cpu: 10m
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
    - mountPath: /etc/kubernetes/static-pod-certs
      name: cert-dir
  terminationGracePeriodSeconds: 135 # bit more than 70s (minimal termination period) + 60s (apiserver graceful termination)
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/kube-apiserver-pod-REVISION
    name: resource-dir
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/kube-apiserver-certs
    name: cert-dir
  - hostPath:
      path: /var/log/kube-apiserver
    name: audit-dir
`)

func v3110KubeApiserverPodYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverPodYaml, nil
}

func v3110KubeApiserverPodYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverRecoveryConfigYaml = []byte(`apiVersion: kubecontrolplane.config.openshift.io/v1
kind: KubeAPIServerConfig
apiServerArguments:
  storage-backend:
  - etcd3
  storage-media-type:
  - application/vnd.kubernetes.protobuf
servingInfo:
  bindAddress: 127.0.0.1:7443
  bindNetwork: tcp4
  certFile: /etc/kubernetes/static-pod-resources/serving-ca.crt
  keyFile: /etc/kubernetes/static-pod-resources/serving-ca.key
  clientCA: /etc/kubernetes/static-pod-resources/serving-ca.crt
storageConfig:
  keyFile: /etc/kubernetes/static-pod-resources/etcd-client.key
  certFile: /etc/kubernetes/static-pod-resources/etcd-client.crt
  ca: /etc/kubernetes/static-pod-resources/etcd-serving-ca-bundle.crt
  urls:
  - "https://localhost:2379"

# Make hypershift happy.
# (Everything bellow this line is just to provide some certs file
# because hypershift tries to read those even if you don't want to set them up.)
authConfig:
  oauthMetadataFile: ""
  requestHeader:
    clientCA: /etc/kubernetes/static-pod-resources/serving-ca.crt
serviceAccountPublicKeyFiles:
- /etc/kubernetes/static-pod-resources/serving-ca.crt
kubeletClientInfo:
  ca: /etc/kubernetes/static-pod-resources/serving-ca.crt
  certFile: /etc/kubernetes/static-pod-resources/serving-ca.crt
  keyFile: /etc/kubernetes/static-pod-resources/serving-ca.key
aggregatorConfig:
  proxyClientInfo:
    certFile: /etc/kubernetes/static-pod-resources/serving-ca.crt
    keyFile: /etc/kubernetes/static-pod-resources/serving-ca.key
`)

func v3110KubeApiserverRecoveryConfigYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverRecoveryConfigYaml, nil
}

func v3110KubeApiserverRecoveryConfigYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverRecoveryConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/recovery-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverRecoveryPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  namespace: openshift-kube-apiserver
  name: kube-apiserver-recovery
  labels:
    revision: "recovery"
spec:
  containers:
  - name: kube-apiserver-recovery
    image: "{{ .KubeApiserverImage }}"
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["hypershift", "openshift-kube-apiserver"]
    args:
    - --config=/etc/kubernetes/static-pod-resources/config.yaml
    resources:
      requests:
        memory: 1Gi
        cpu: 150m
    ports:
    - containerPort: 7443
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
  terminationGracePeriodSeconds: 0
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
  - hostPath:
      path: "{{ .ResourceDir }}"
    name: resource-dir
`)

func v3110KubeApiserverRecoveryPodYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverRecoveryPodYaml, nil
}

func v3110KubeApiserverRecoveryPodYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverRecoveryPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/recovery-pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-kube-apiserver
  name: apiserver
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  type: ClusterIP
  selector:
    apiserver: "true"
  ports:
  - name: https
    port: 443
    targetPort: 6443
`)

func v3110KubeApiserverSvcYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverSvcYaml, nil
}

func v3110KubeApiserverSvcYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v3.11.0/kube-apiserver/cm.yaml":              v3110KubeApiserverCmYaml,
	"v3.11.0/kube-apiserver/defaultconfig.yaml":   v3110KubeApiserverDefaultconfigYaml,
	"v3.11.0/kube-apiserver/kubeconfig-cm.yaml":   v3110KubeApiserverKubeconfigCmYaml,
	"v3.11.0/kube-apiserver/ns.yaml":              v3110KubeApiserverNsYaml,
	"v3.11.0/kube-apiserver/pod-cm.yaml":          v3110KubeApiserverPodCmYaml,
	"v3.11.0/kube-apiserver/pod.yaml":             v3110KubeApiserverPodYaml,
	"v3.11.0/kube-apiserver/recovery-config.yaml": v3110KubeApiserverRecoveryConfigYaml,
	"v3.11.0/kube-apiserver/recovery-pod.yaml":    v3110KubeApiserverRecoveryPodYaml,
	"v3.11.0/kube-apiserver/svc.yaml":             v3110KubeApiserverSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v3.11.0": {nil, map[string]*bintree{
		"kube-apiserver": {nil, map[string]*bintree{
			"cm.yaml":              {v3110KubeApiserverCmYaml, map[string]*bintree{}},
			"defaultconfig.yaml":   {v3110KubeApiserverDefaultconfigYaml, map[string]*bintree{}},
			"kubeconfig-cm.yaml":   {v3110KubeApiserverKubeconfigCmYaml, map[string]*bintree{}},
			"ns.yaml":              {v3110KubeApiserverNsYaml, map[string]*bintree{}},
			"pod-cm.yaml":          {v3110KubeApiserverPodCmYaml, map[string]*bintree{}},
			"pod.yaml":             {v3110KubeApiserverPodYaml, map[string]*bintree{}},
			"recovery-config.yaml": {v3110KubeApiserverRecoveryConfigYaml, map[string]*bintree{}},
			"recovery-pod.yaml":    {v3110KubeApiserverRecoveryPodYaml, map[string]*bintree{}},
			"svc.yaml":             {v3110KubeApiserverSvcYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
