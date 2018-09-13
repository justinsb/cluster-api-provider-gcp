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

package machine_controller_app

import (
	"sigs.k8s.io/cluster-api-provider-gcp/cloud/google"
	"sigs.k8s.io/cluster-api-provider-gcp/cloud/google/machinesetup"
	"sigs.k8s.io/cluster-api-provider-gcp/cmd/gce-controller/machine-controller-app/options"
	machinecontroller "sigs.k8s.io/cluster-api/pkg/controller/machine"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

//"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"

//"sigs.k8s.io/cluster-api/pkg/controller/sharedinformers"

const (
	gceMachineControllerName = "gce-controller"
)

func AddToManager(mgr manager.Manager, server *options.MachineControllerServer) error {
	recorder := mgr.GetRecorder(gceMachineControllerName)

	configWatch, err := machinesetup.NewConfigWatch(server.MachineSetupConfigsPath)
	if err != nil {
		return err
	}

	params := google.MachineActuatorParams{
		V1Alpha1Client:           mgr.GetClient(),
		MachineSetupConfigGetter: configWatch,
		EventRecorder:            recorder,
	}
	actuator, err := google.NewMachineActuator(params)
	if err != nil {
		return err
	}

	return machinecontroller.AddWithActuator(mgr, actuator)
}

/*
func StartMachineController(server *options.MachineControllerServer, recorder record.EventRecorder, shutdown <-chan struct{}) {
	config, err := controller.GetConfig(server.CommonConfig.Kubeconfig)
	if err != nil {
		glog.Fatalf("Could not create Config for talking to the apiserver: %v", err)
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Could not create client for talking to the apiserver: %v", err)
	}

	configWatch, err := machinesetup.NewConfigWatch(server.MachineSetupConfigsPath)
	if err != nil {
		glog.Fatalf("Could not create config watch: %v", err)
	}
	params := google.MachineActuatorParams{
		V1Alpha1Client:           client.ClusterV1alpha1(),
		MachineSetupConfigGetter: configWatch,
		EventRecorder:            recorder,
	}
	actuator, err := google.NewMachineActuator(params)

	if err != nil {
		glog.Fatalf("Could not create Google machine actuator: %v", err)
	}

	si := sharedinformers.NewSharedInformers(config, shutdown)
	// If this doesn't compile, the code generator probably
	// overwrote the customized NewMachineController function.
	c := machine.NewMachineController(config, si, actuator)
	c.RunAsync(shutdown)

	select {}
}

func RunMachineController(server *options.MachineControllerServer) error {
	kubeConfig, err := controller.GetConfig(server.CommonConfig.Kubeconfig)
	if err != nil {
		glog.Errorf("Could not create Config for talking to the apiserver: %v", err)
		return err
	}

	clientSet, err := kubernetes.NewForConfig(
		rest.AddUserAgent(kubeConfig, "machine-controller-manager"),
	)
	if err != nil {
		glog.Errorf("Invalid API configuration for kubeconfig-control: %v", err)
		return err
	}

	recorder, err := createRecorder(clientSet)
	if err != nil {
		glog.Errorf("Could not create event recorder : %v", err)
		return err
	}

	// run function will block and never return.
	run := func(stop <-chan struct{}) {
		StartMachineController(server, recorder, stop)
	}

	leaderElectConfig := config.GetLeaderElectionConfig()
	if !leaderElectConfig.LeaderElect {
		run(make(<-chan (struct{})))
	}

	// Identity used to distinguish between multiple machine controller instances.
	id, err := os.Hostname()
	if err != nil {
		return err
	}

	leaderElectionClient := kubernetes.NewForConfigOrDie(rest.AddUserAgent(kubeConfig, "machine-leader-election"))

	id = id + "-" + string(uuid.NewUUID())
	// Lock required for leader election
	rl, err := resourcelock.New(
		leaderElectConfig.ResourceLock,
		metav1.NamespaceSystem,
		gceMachineControllerName,
		leaderElectionClient.CoreV1(),
		resourcelock.ResourceLockConfig{
			Identity:      id + "-" + gceMachineControllerName,
			EventRecorder: recorder,
		})
	if err != nil {
		return err
	}

	// Try and become the leader and start machine controller loops
	leaderelection.RunOrDie(leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: leaderElectConfig.LeaseDuration.Duration,
		RenewDeadline: leaderElectConfig.RenewDeadline.Duration,
		RetryPeriod:   leaderElectConfig.RetryPeriod.Duration,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				glog.Fatalf("leaderelection lost")
			},
		},
	})
	panic("unreachable")
}

func createRecorder(kubeClient *kubernetes.Clientset) (record.EventRecorder, error) {

	eventsScheme := runtime.NewScheme()
	if err := corev1.AddToScheme(eventsScheme); err != nil {
		return nil, err
	}
	// We also emit events for our own types.
	clusterapiclientsetscheme.AddToScheme(eventsScheme)

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(kubeClient.CoreV1().RESTClient()).Events("")})
	return eventBroadcaster.NewRecorder(eventsScheme, corev1.EventSource{Component: gceMachineControllerName}), nil
}
*/
