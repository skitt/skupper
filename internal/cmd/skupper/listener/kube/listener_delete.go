package kube

import (
	"context"
	"fmt"
	"time"

	"github.com/skupperproject/skupper/internal/cmd/skupper/utils"
	"github.com/skupperproject/skupper/internal/kube/client"
	skupperv1alpha1 "github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned/typed/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/utils/validator"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	listenerDeleteLong    = "Delete a listener <name>"
	listenerDeleteExample = "skupper listener delete database"
)

type ListenerDelete struct {
	timeout time.Duration
}

type CmdListenerDelete struct {
	client    skupperv1alpha1.SkupperV1alpha1Interface
	CobraCmd  cobra.Command
	flags     ListenerDelete
	namespace string
	name      string
}

func NewCmdListenerDelete() *CmdListenerDelete {

	skupperCmd := CmdListenerDelete{}

	cmd := cobra.Command{
		Use:     "delete <name>",
		Short:   "delete a listener",
		Long:    listenerDeleteLong,
		Example: listenerDeleteExample,
		PreRun:  skupperCmd.NewClient,
		Run: func(cmd *cobra.Command, args []string) {
			utils.HandleErrorList(skupperCmd.ValidateInput(args))
			utils.HandleError(skupperCmd.Run())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return skupperCmd.WaitUntilReady()
		},
	}
	skupperCmd.CobraCmd = cmd
	skupperCmd.AddFlags()

	return &skupperCmd
}

func (cmd *CmdListenerDelete) NewClient(cobraCommand *cobra.Command, args []string) {
	cli, err := client.NewClient(cobraCommand.Flag("namespace").Value.String(), cobraCommand.Flag("context").Value.String(), cobraCommand.Flag("kubeconfig").Value.String())
	utils.HandleError(err)

	cmd.client = cli.GetSkupperClient().SkupperV1alpha1()
	cmd.namespace = cli.Namespace
}

func (cmd *CmdListenerDelete) AddFlags() {
	cmd.CobraCmd.Flags().DurationVarP(&cmd.flags.timeout, "timeout", "t", 60*time.Second, "Raise an error if the operation does not complete in the given period of time.")
}

func (cmd *CmdListenerDelete) ValidateInput(args []string) []error {
	var validationErrors []error
	resourceStringValidator := validator.NewResourceStringValidator()

	// Validate arguments name
	if len(args) < 1 {
		validationErrors = append(validationErrors, fmt.Errorf("listener name must be specified"))
	} else if len(args) > 1 {
		validationErrors = append(validationErrors, fmt.Errorf("only one argument is allowed for this command"))
	} else if args[0] == "" {
		validationErrors = append(validationErrors, fmt.Errorf("listener name must not be empty"))
	} else {
		ok, err := resourceStringValidator.Evaluate(args[0])
		if !ok {
			validationErrors = append(validationErrors, fmt.Errorf("listener name is not valid: %s", err))
		} else {
			cmd.name = args[0]
		}

		if cmd.name != "" {
			// Validate that there is already a listener with this name in the namespace
			listener, err := cmd.client.Listeners(cmd.namespace).Get(context.TODO(), cmd.name, metav1.GetOptions{})
			if err != nil || listener == nil {
				validationErrors = append(validationErrors, fmt.Errorf("listener %s does not exist in namespace %s", cmd.name, cmd.namespace))
			}
		}

		//TBD what is valid timeout
		if cmd.flags.timeout <= 0*time.Minute {
			validationErrors = append(validationErrors, fmt.Errorf("timeout is not valid"))
		}
	}

	return validationErrors
}

func (cmd *CmdListenerDelete) Run() error {
	err := cmd.client.Listeners(cmd.namespace).Delete(context.TODO(), cmd.name, metav1.DeleteOptions{})
	return err
}

func (cmd *CmdListenerDelete) WaitUntilReady() error {
	waitTime := int(cmd.flags.timeout.Seconds())
	err := utils.NewSpinnerWithTimeout("Waiting for deletion to complete...", waitTime, func() error {

		resource, err := cmd.client.Listeners(cmd.namespace).Get(context.TODO(), cmd.name, metav1.GetOptions{})
		if err == nil && resource != nil {
			return fmt.Errorf("error deleting the resource")
		} else {
			return nil
		}
	})

	if err != nil {
		return fmt.Errorf("Listener %q not deleted yet, check the logs for more information %s\n", cmd.name, err)
	}

	fmt.Printf("Listener %q deleted\n", cmd.name)
	return nil
}
