/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/util/logs"
	"sigs.k8s.io/cluster-api-provider-gcp/cmd/gce-controller/cluster-controller-app"
	"sigs.k8s.io/cluster-api-provider-gcp/cmd/gce-controller/machine-controller-app"
	machineoptions "sigs.k8s.io/cluster-api-provider-gcp/cmd/gce-controller/machine-controller-app/options"
	clusterapis "sigs.k8s.io/cluster-api/pkg/apis"
	clusterapiconfig "sigs.k8s.io/cluster-api/pkg/controller/config"
	controllerruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {

	fs := pflag.CommandLine
	var controllerType, machineSetupConfigsPath string
	fs.StringVar(&controllerType, "controller", controllerType, "specify whether this should run the machine or cluster controller")
	fs.StringVar(&machineSetupConfigsPath, "machinesetup", machineSetupConfigsPath, "path to machine setup configs file")
	clusterapiconfig.ControllerConfig.AddFlags(pflag.CommandLine)
	// the following line exists to make glog happy, for more information, see: https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})

	// Map go flags to pflag
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	pflag.Parse()

	logs.InitLogs()
	defer logs.FlushLogs()

	// Get a config to talk to the apiserver
	cfg, err := controllerruntimeconfig.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Registering Components.")

	// Setup Scheme for all resources
	if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatal(err)
	}

	// Setup all Controllers
	//if err := controller.AddToManager(mgr); err != nil {
	//	log.Fatal(err)
	//}

	if controllerType == "machine" {
		options := machineoptions.NewMachineControllerServer(machineSetupConfigsPath)
		if err := machine_controller_app.AddToManager(mgr, options); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register machine controller. Err: %v\n", err)
			os.Exit(1)
		}
	} else if controllerType == "cluster" {
		//		clusterServer := clusteroptions.NewClusterControllerServer()
		if err := cluster_controller_app.AddToManager(mgr); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register cluster controller. Err: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Failed to start controller, `controller` flag must be either `machine` or `cluster` but was %v.\n", controllerType)
		os.Exit(1)
	}

	log.Printf("Starting the Cmd.")

	// Start the Cmd
	log.Fatal(mgr.Start(signals.SetupSignalHandler()))

}
