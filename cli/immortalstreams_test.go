package cli

import (
	"testing"

	"github.com/coder/serpent"
)

func TestImmortalStreamsCommand(t *testing.T) {
	t.Parallel()

	t.Run("CommandStructure", func(t *testing.T) {
		t.Parallel()

		root := &RootCmd{}
		cmd := root.immortalStreams()

		if cmd.Use != "immortal-streams" {
			t.Fatalf("expected use 'immortal-streams', got %q", cmd.Use)
		}

		if len(cmd.Children) != 2 {
			t.Fatalf("expected 2 subcommands, got %d", len(cmd.Children))
		}

		// Check subcommands
		var listCmd, deleteCmd *serpent.Command
		for _, child := range cmd.Children {
			switch child.Use {
			case "list <workspace>":
				listCmd = child
			case "delete <workspace> <stream-name>":
				deleteCmd = child
			}
		}

		if listCmd == nil {
			t.Fatalf("expected list subcommand")
		}
		if deleteCmd == nil {
			t.Fatalf("expected delete subcommand")
		}
	})

	t.Run("ListCommand", func(t *testing.T) {
		t.Parallel()

		root := &RootCmd{}
		cmd := root.immortalStreamsList()

		if cmd.Use != "list <workspace>" {
			t.Fatalf("expected use 'list <workspace>', got %q", cmd.Use)
		}

		if cmd.Short == "" {
			t.Fatalf("expected non-empty short description")
		}
	})

	t.Run("DeleteCommand", func(t *testing.T) {
		t.Parallel()

		root := &RootCmd{}
		cmd := root.immortalStreamsDelete()

		if cmd.Use != "delete <workspace> <stream-name>" {
			t.Fatalf("expected use 'delete <workspace> <stream-name>', got %q", cmd.Use)
		}

		if cmd.Short == "" {
			t.Fatalf("expected non-empty short description")
		}
	})
}

func TestSSHImmortalFlags(t *testing.T) {
	t.Parallel()

	t.Run("SSHHasImmortalFlags", func(t *testing.T) {
		t.Parallel()

		root := &RootCmd{}
		cmd := root.ssh()

		// Check that immortal flags exist
		var immortalFlag, immortalFallbackFlag *serpent.Option
		for _, opt := range cmd.Options {
			switch opt.Flag {
			case "immortal":
				immortalFlag = &opt
			case "immortal-fallback":
				immortalFallbackFlag = &opt
			}
		}

		if immortalFlag == nil {
			t.Fatalf("expected --immortal flag in SSH command")
		}
		if immortalFallbackFlag == nil {
			t.Fatalf("expected --immortal-fallback flag in SSH command")
		}

		if immortalFlag.Description == "" {
			t.Fatalf("expected non-empty description for --immortal flag")
		}
		if immortalFallbackFlag.Description == "" {
			t.Fatalf("expected non-empty description for --immortal-fallback flag")
		}
	})
}

func TestPortForwardImmortalFlags(t *testing.T) {
	t.Parallel()

	t.Run("PortForwardHasImmortalFlags", func(t *testing.T) {
		t.Parallel()

		root := &RootCmd{}
		cmd := root.portForward()

		// Check that immortal flags exist
		var immortalFlag, immortalFallbackFlag *serpent.Option
		for _, opt := range cmd.Options {
			switch opt.Flag {
			case "immortal":
				immortalFlag = &opt
			case "immortal-fallback":
				immortalFallbackFlag = &opt
			}
		}

		if immortalFlag == nil {
			t.Fatalf("expected --immortal flag in port-forward command")
		}
		if immortalFallbackFlag == nil {
			t.Fatalf("expected --immortal-fallback flag in port-forward command")
		}

		if immortalFlag.Description == "" {
			t.Fatalf("expected non-empty description for --immortal flag")
		}
		if immortalFallbackFlag.Description == "" {
			t.Fatalf("expected non-empty description for --immortal-fallback flag")
		}
	})
}
