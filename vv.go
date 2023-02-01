package main

import (
	"flag"
	"fmt"
	"github.com/coreos/pkg/flagutil"
	log "k8s.io/klog"
	"os"
	"strings"
)

type flagSlice []string

func (t *flagSlice) String() string {
	return fmt.Sprintf("%v", *t)
}
func (t *flagSlice) Set(val string) error {
	*t = append(*t, val)
	return nil
}

type CmdLineOpts struct {
	etcdEndpoints             string
	etcdPrefix                string
	etcdKeyfile               string
	etcdCertfile              string
	etcdCAFile                string
	etcdUsername              string
	etcdPassword              string
	version                   bool
	kubeSubnetMgr             bool
	kubeApiUrl                string
	kubeAnnotationPrefix      string
	kubeConfigFile            string
	iface                     flagSlice
	ifaceRegex                flagSlice
	ipMasq                    bool
	ifaceCanReach             string
	subnetFile                string
	publicIP                  string
	publicIPv6                string
	subnetLeaseRenewMargin    int
	healthzIP                 string
	healthzPort               int
	iptablesResyncSeconds     int
	iptablesForwardRules      bool
	netConfPath               string
	setNodeNetworkUnavailable bool
}

var (
	opts         CmdLineOpts
	flannelFlags = flag.NewFlagSet("flannel", flag.ExitOnError)
)

func copyFlag(name string) {
	flannelFlags.Var(flag.Lookup(name).Value, flag.Lookup(name).Name, flag.Lookup(name).Usage)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n", os.Args[0])
	flannelFlags.PrintDefaults()
	os.Exit(0)
}

func init() {
	flannelFlags.StringVar(&opts.etcdEndpoints, "etcd-endpoints", "http://127.0.0.1:4001,http://127.0.0.1:2379", "a comma-delimited list of etcd endpoints")
	flannelFlags.StringVar(&opts.etcdPrefix, "etcd-prefix", "/coreos.com/network", "etcd prefix")
	flannelFlags.StringVar(&opts.etcdKeyfile, "etcd-keyfile", "", "SSL key file used to secure etcd communication")
	flannelFlags.StringVar(&opts.etcdCertfile, "etcd-certfile", "", "SSL certification file used to secure etcd communication")
	flannelFlags.StringVar(&opts.etcdCAFile, "etcd-cafile", "", "SSL Certificate Authority file used to secure etcd communication")
	flannelFlags.StringVar(&opts.etcdUsername, "etcd-username", "", "username for BasicAuth to etcd")
	flannelFlags.StringVar(&opts.etcdPassword, "etcd-password", "", "password for BasicAuth to etcd")
	flannelFlags.Var(&opts.iface, "iface", "interface to use (IP or name) for inter-host communication. Can be specified multiple times to check each option in order. Returns the first match found.")
	flannelFlags.Var(&opts.ifaceRegex, "iface-regex", "regex expression to match the first interface to use (IP or name) for inter-host communication. Can be specified multiple times to check each regex in order. Returns the first match found. Regexes are checked after specific interfaces specified by the iface option have already been checked.")
	flannelFlags.StringVar(&opts.ifaceCanReach, "iface-can-reach", "", "detect interface to use (IP or name) for inter-host communication based on which will be used for provided IP. This is exactly the interface to use of command 'ip route get <ip-address>'")
	flannelFlags.StringVar(&opts.subnetFile, "subnet-file", "/run/flannel/subnet.env", "filename where env variables (subnet, MTU, ... ) will be written to")
	flannelFlags.StringVar(&opts.publicIP, "public-ip", "", "IP accessible by other nodes for inter-host communication")
	flannelFlags.StringVar(&opts.publicIPv6, "public-ipv6", "", "IPv6 accessible by other nodes for inter-host communication")
	flannelFlags.IntVar(&opts.subnetLeaseRenewMargin, "subnet-lease-renew-margin", 60, "subnet lease renewal margin, in minutes, ranging from 1 to 1439")
	flannelFlags.BoolVar(&opts.ipMasq, "ip-masq", false, "setup IP masquerade rule for traffic destined outside of overlay network")
	flannelFlags.BoolVar(&opts.kubeSubnetMgr, "kube-subnet-mgr", false, "contact the Kubernetes API for subnet assignment instead of etcd.")
	flannelFlags.StringVar(&opts.kubeApiUrl, "kube-api-url", "", "Kubernetes API server URL. Does not need to be specified if flannel is running in a pod.")
	flannelFlags.StringVar(&opts.kubeAnnotationPrefix, "kube-annotation-prefix", "flannel.alpha.coreos.com", `Kubernetes annotation prefix. Can contain single slash "/", otherwise it will be appended at the end.`)
	flannelFlags.StringVar(&opts.kubeConfigFile, "kubeconfig-file", "", "kubeconfig file location. Does not need to be specified if flannel is running in a pod.")
	flannelFlags.BoolVar(&opts.version, "version", false, "print version and exit")
	flannelFlags.StringVar(&opts.healthzIP, "healthz-ip", "0.0.0.0", "the IP address for healthz server to listen")
	flannelFlags.IntVar(&opts.healthzPort, "healthz-port", 0, "the port for healthz server to listen(0 to disable)")
	flannelFlags.IntVar(&opts.iptablesResyncSeconds, "iptables-resync", 5, "resync period for iptables rules, in seconds")
	flannelFlags.BoolVar(&opts.iptablesForwardRules, "iptables-forward-rules", true, "add default accept rules to FORWARD chain in iptables")
	flannelFlags.StringVar(&opts.netConfPath, "net-config-path", "/etc/kube-flannel/net-conf.json", "path to the network configuration file")
	flannelFlags.BoolVar(&opts.setNodeNetworkUnavailable, "set-node-network-unavailable", true, "set NodeNetworkUnavailable after ready")

	log.InitFlags(nil)
	err := flag.Set("logtostderr", "true")
	if err != nil {
		log.Error("Can't set the logtostderr flag", err)
		os.Exit(1)
	}

	// Only copy the non file logging options from klog
	copyFlag("v")
	copyFlag("vmodule")
	copyFlag("log_backtrace_at")

	flannelFlags.Usage = usage

	// now parse command line args
	err = flannelFlags.Parse(os.Args[1:])
	if err != nil {
		log.Error("Can't parse flannel flags", err)
		os.Exit(1)
	}
}
func main() {
	//if opts.version {
	//	fmt.Fprintln(os.Stderr, version.Version)
	//	os.Exit(0)
	//}
	err := flagutil.SetFlagsFromEnv(flannelFlags, "FLANNELD")
	if err != nil {
		log.Error("Failed to set flag FLANNELD from env", err)
	}
	//if opts.etcdEndpoints != "" {
	//	fmt.Fprintln(os.Stderr, opts.etcdEndpoints)
	//	os.Exit(0)
	//}
	log.Infoln(strings.Count("../dev", "dev"))
	//ctx, cancel := context.WithCancel(context.Background())
}
