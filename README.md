# Airship UI

Airship UI is an [electron](https://www.electronjs.org/) that is designed to allow you to interact with Airship components, find and connect to the kubernetes cluster and use plugins to tie together a singular dashboard to view addons without the need to go to a separate url or application for each.

## Prerequisites

- A working [kubernetes](https://kubernetes.io/) or [airship](https://wiki.openstack.org/wiki/Airship) installation
- [Go 1.13+](https://golang.org/dl/)
- [npm](https://www.npmjs.com/)

## Getting Started

```
git clone https://opendev.org/airship/airshipui
cd airshipui
make build
cd web
npm install
npm install --save-dev electron
npm install electron-json-config
```

## Adding Additional Functionality

Airship UI can be seamlessly integrated with service dashboards and other web-based tools by providing the necessary configuration in
$HOME/.airship/airshipui.json.

To add service dashboards, create a section at the top level of airshipui.json as follows:

```
"clusters": [
        {
            "name": "clusterA",
            "baseFqdn": "svc.cluster.local",
            "namespaces": [
                {
                    "name": "ceph",
                    "dashboards": [
                        {
                            "name": "Ceph",
                            "protocol": "https",
                            "fqdn": "ceph-dash.example.domain",
                            "port": 443,
                            "path": ""
                        }
                    ]
                },
                {
                    "name": "openstack",
                    "dashboards": [
                        {
                            "name": "Horizon",
                            "protocol": "http",
                            "hostname": "horizon",
                            "port": 80,
                            "path": "dashboard/auth/login"
                        }
                    ]
                }
            ]
        }
]
```

For dashboards that are made available through service endpoints in your cluster, the FQDN for the dashboard will be constructed using the format
"hostname.namespace.baseFqdn". In the above example, the configuration for Horizon specifies a service dashboard available at
"http://horizon.openstack.svc.cluster.local:80/dashboard/auth/login"

 Alternatively, you may choose to specify the FQDN directly, as in the above Ceph example. This configuration specifies a Ceph dashboard available at
"https://ceph-dash.example.domain:443/"

If both "hostname" and "fqdn" are provided, "fqdn" will take precedence.

The airshipui.json configuration file can also be used to launch "plugins", or external executables, in the background as Airship UI starts. Any processes
launched by Airship UI will be terminated when Airship UI exits, including any child processes started by the plugins. If the plugin launches a web
dashboard, it can be also be included in the list of service dashboards within Airship UI. The following example demonstrates how to add configuration to
launch and use Octant within Airship UI:

```
"plugins": [
        {
            "name": "Octant",
            "dashboard": {
                "protocol": "http",
                "fqdn": "localhost",
                "port": 7777,
                "path": ""
            },
            "executable": {
                "autoStart": true,
                "filepath": "/usr/local/bin/octant",
                "args": [
                    "--disable-open-browser",
                    "--kubeconfig",
                    "/home/ubuntu/.airship/kubeconfig"
                ]
            }
        }
]
```

To prevent a plugin from launching but retain its configuration for later use, simply set "autoStart" to false.

## Developer's Guide

Instructions on setting up a development environment and more details can be found in the [Developer's Guide](docs/source/developers.md)
