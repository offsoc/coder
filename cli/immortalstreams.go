package cli

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/coder/v2/codersdk/workspacesdk"
	"github.com/coder/serpent"
)

func (r *RootCmd) immortalStreams() *serpent.Command {
	cmd := &serpent.Command{
		Use:     "immortal-streams",
		Short:   "Manage immortal streams for workspaces",
		Long:    "List and delete immortal streams that provide resilient connections to workspaces.",
		Aliases: []string{"immortal-stream"},
		Handler: func(inv *serpent.Invocation) error {
			return inv.Command.HelpHandler(inv)
		},
		Children: []*serpent.Command{
			r.immortalStreamsList(),
			r.immortalStreamsDelete(),
		},
	}

	return cmd
}

func (r *RootCmd) immortalStreamsList() *serpent.Command {
	var (
		appearanceConfig codersdk.AppearanceConfig
	)
	client := new(codersdk.Client)
	cmd := &serpent.Command{
		Use:   "list <workspace>",
		Short: "List immortal streams for a workspace",
		Middleware: serpent.Chain(
			serpent.RequireNArgs(1),
			r.InitClient(client),
			initAppearance(client, &appearanceConfig),
		),
		Handler: func(inv *serpent.Invocation) error {
			ctx := inv.Context()

			_, workspaceAgent, _, err := getWorkspaceAndAgent(ctx, inv, client, false, inv.Args[0])
			if err != nil {
				return err
			}

			// Connect to the agent to access its API
			conn, err := workspacesdk.New(client).DialAgent(ctx, workspaceAgent.ID, &workspacesdk.DialAgentOptions{
				Logger: inv.Logger,
			})
			if err != nil {
				return xerrors.Errorf("dial agent: %w", err)
			}
			defer conn.Close()

			if !conn.AwaitReachable(ctx) {
				return xerrors.New("agent not reachable")
			}

			// TODO: Make API call to list immortal streams
			// For now, return placeholder message
			_, _ = fmt.Fprintln(inv.Stdout, "No immortal streams found.")
			return nil
		},
	}

	return cmd
}

func (r *RootCmd) immortalStreamsDelete() *serpent.Command {
	var (
		appearanceConfig codersdk.AppearanceConfig
	)
	client := new(codersdk.Client)
	cmd := &serpent.Command{
		Use:   "delete <workspace> <stream-name>",
		Short: "Delete an immortal stream",
		Middleware: serpent.Chain(
			serpent.RequireNArgs(2),
			r.InitClient(client),
			initAppearance(client, &appearanceConfig),
		),
		Handler: func(inv *serpent.Invocation) error {
			ctx := inv.Context()

			workspaceName := inv.Args[0]
			streamName := inv.Args[1]

			_, workspaceAgent, _, err := getWorkspaceAndAgent(ctx, inv, client, false, workspaceName)
			if err != nil {
				return err
			}

			// Connect to the agent to access its API
			conn, err := workspacesdk.New(client).DialAgent(ctx, workspaceAgent.ID, &workspacesdk.DialAgentOptions{
				Logger: inv.Logger,
			})
			if err != nil {
				return xerrors.Errorf("dial agent: %w", err)
			}
			defer conn.Close()

			if !conn.AwaitReachable(ctx) {
				return xerrors.New("agent not reachable")
			}

			// TODO: Make API call to delete immortal stream
			// For now, return placeholder message
			_, _ = fmt.Fprintf(inv.Stdout, "Immortal Stream %q would be deleted (not implemented yet).\n", streamName)
			return nil
		},
	}

	return cmd
}
