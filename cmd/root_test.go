package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRootCmdWithFromToFlag(t *testing.T) {
	// Backup original os.Args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Simulate command-line arguments for the test
	// In this case, we're testing the rootCmd with the --from-to flag set
	os.Args = []string{"azcopy", "copy", "--from-to", "BlobPipe"}

	// Execute the root command
	err := rootCmd.Execute()

	// Assert no error occurred
	assert.NoError(t, err, "expected no error while executing rootCmd with --from-to flag")

	// Additional assertions can be done here
	// For example, verify if the isPipeDownload variable is correctly set
	assert.True(t, isPipeDownload, "expected isPipeDownload to be true when --from-to=BlobPipe is set")
}

// Helper function to reset rootCmd before each test
func resetRootCmd() {
	rootCmd = &cobra.Command{
		Use:   "azcopy",
		Short: "Test AzCopy root command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Mock PersistentPreRunE logic here as needed
			var fromToFlagValue string
			var err error

			if cmd.Flags().Changed("from-to") {
				fromToFlagValue, err = cmd.Flags().GetString("fromto")
				if err != nil {
					return err
				}
			}

			if fromToFlagValue == "BlobPipe" {
				isPipeDownload = true
			}
			return nil
		},
	}
	var azcopyFromTo string
	rootCmd.PersistentFlags().StringVar(&azcopyFromTo, "fromto", "", "Used to specify the source and destination types.")
	rootCmd.PersistentFlags().StringVar(&outputFormatRaw, "output-type", "text", "Format of the command's output.")
	rootCmd.PersistentFlags().StringVar(&outputVerbosityRaw, "output-level", "default", "Define the output verbosity.")
	rootCmd.PersistentFlags().StringVar(&logVerbosityRaw, "log-level", "INFO", "Define the log verbosity for the log file.")
	rootCmd.PersistentFlags().StringVar(&cmdLineExtraSuffixesAAD, trustedSuffixesNameAAD, "", "Specifies additional domain suffixes.")
	rootCmd.PersistentFlags().BoolVar(&azcopySkipVersionCheck, "skip-version-check", false, "Do not perform the version check at startup.")
	rootCmd.PersistentFlags().BoolVar(&cancelFromStdin, "cancel-from-stdin", false, "Used by partner teams to send in `cancel` through stdin.")
	rootCmd.PersistentFlags().BoolVar(&azcopyAwaitContinue, "await-continue", false, "Used when debugging, to tell AzCopy to await `continue` on stdin.")
	rootCmd.PersistentFlags().BoolVar(&azcopyAwaitAllowOpenFiles, "await-open", false, "Used when debugging, to tell AzCopy to await `open` on stdin.")
	rootCmd.PersistentFlags().StringVar(&debugSkipFiles, "debug-skip-files", "", "Used when debugging, to tell AzCopy to cancel the job midway.")
	rootCmd.PersistentFlags().StringVar(&retryStatusCodes, "retry-status-codes", "", "Comma-separated list of HTTP status codes to retry on.")
}
