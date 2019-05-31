package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type serverVersionCmd struct {
	out io.Writer
}

func NewServerVersionCommand(streams genericclioptions.IOStreams) *cobra.Command {
	helloWorldCmd := &serverVersionCmd{
		out: streams.Out,
	}

	cmd := &cobra.Command{
		Use:          "server-version",
		Short:        "Prints Kubernetes server version",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			return helloWorldCmd.run()
		},
	}

	return cmd
}

func (hw *serverVersionCmd) run() error {
	serverVersion, err := getServerVersion()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(hw.out, "Hello from Kubernetes server with version %s!\n", serverVersion)
	if err != nil {
		return err
	}
	return nil
}

func getServerVersion() (string, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return "", err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}

	sv, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}

	return sv.String(), nil
}
