// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type githubCheckBranchProtectionOptions struct {
	APIURL                       string   `json:"apiUrl,omitempty"`
	Branch                       string   `json:"branch,omitempty"`
	Owner                        string   `json:"owner,omitempty"`
	Repository                   string   `json:"repository,omitempty"`
	RequiredChecks               []string `json:"requiredChecks,omitempty"`
	RequireEnforceAdmins         bool     `json:"requireEnforceAdmins,omitempty"`
	RequiredApprovingReviewCount int      `json:"requiredApprovingReviewCount,omitempty"`
	Token                        string   `json:"token,omitempty"`
}

// GithubCheckBranchProtectionCommand Check branch protection of a GitHub branch
func GithubCheckBranchProtectionCommand() *cobra.Command {
	const STEP_NAME = "githubCheckBranchProtection"

	metadata := githubCheckBranchProtectionMetadata()
	var stepConfig githubCheckBranchProtectionOptions
	var startTime time.Time

	var createGithubCheckBranchProtectionCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Check branch protection of a GitHub branch",
		Long: `This step allows you to check if certain branch protection rules are fulfilled.

It can for example be used to verify if certain status checks are mandatory. This can be helpful to decide if a certain check needs to be performed again after merging a pull request.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.Token)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			githubCheckBranchProtection(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addGithubCheckBranchProtectionFlags(createGithubCheckBranchProtectionCmd, &stepConfig)
	return createGithubCheckBranchProtectionCmd
}

func addGithubCheckBranchProtectionFlags(cmd *cobra.Command, stepConfig *githubCheckBranchProtectionOptions) {
	cmd.Flags().StringVar(&stepConfig.APIURL, "apiUrl", `https://api.github.com`, "Set the GitHub API url.")
	cmd.Flags().StringVar(&stepConfig.Branch, "branch", os.Getenv("PIPER_branch"), "The name of the branch for which the protection settings should be checked.")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Name of the GitHub organization.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Name of the GitHub repository.")
	cmd.Flags().StringSliceVar(&stepConfig.RequiredChecks, "requiredChecks", []string{}, "List of checks which have to be set to 'required' in the GitHub repository configuration.")
	cmd.Flags().BoolVar(&stepConfig.RequireEnforceAdmins, "requireEnforceAdmins", false, "Check if 'Include Administrators' option is set in the GitHub repository configuration.")
	cmd.Flags().IntVar(&stepConfig.RequiredApprovingReviewCount, "requiredApprovingReviewCount", 0, "Check if 'Require pull request reviews before merging' option is set with at least the defined number of reviewers in the GitHub repository configuration.")
	cmd.Flags().StringVar(&stepConfig.Token, "token", os.Getenv("PIPER_token"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line.")

	cmd.MarkFlagRequired("apiUrl")
	cmd.MarkFlagRequired("branch")
	cmd.MarkFlagRequired("owner")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("token")
}

// retrieve step metadata
func githubCheckBranchProtectionMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "githubCheckBranchProtection",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "apiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubApiUrl"}},
					},
					{
						Name:        "branch",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "owner",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubOrg"}},
					},
					{
						Name:        "repository",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubRepo"}},
					},
					{
						Name:        "requiredChecks",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "requireEnforceAdmins",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "requiredApprovingReviewCount",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name: "token",
						ResourceRef: []config.ResourceReference{
							{
								Name: "githubTokenCredentialsId",
								Type: "secret",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{{Name: "githubToken"}},
					},
				},
			},
		},
	}
	return theMetaData
}
