package coderd

import (
	"net/http"
	"strings"
	"time"

	"cdr.dev/slog"
	"github.com/mark3labs/mcp-go/server"

	"github.com/coder/coder/v2/coderd/httpapi"
	"github.com/coder/coder/v2/coderd/httpmw"
	"github.com/coder/coder/v2/coderd/mcp"
	"github.com/coder/coder/v2/codersdk"
)

// mcpHTTPHandler creates the MCP HTTP transport handler
func (api *API) mcpHTTPHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create MCP server instance for each request
		mcpServer, err := mcp.NewServer(api.Logger.Named("mcp"))
		if err != nil {
			api.Logger.Error(r.Context(), "failed to create MCP server", slog.Error(err))
			httpapi.Write(r.Context(), w, http.StatusInternalServerError, codersdk.Response{
				Message: "MCP server initialization failed",
			})
			return
		}

		authenticatedClient := codersdk.New(api.AccessURL)
		// Extract the original session token from the request
		authenticatedClient.SetSessionToken(httpmw.APITokenFromRequest(r))

		// Register tools with authenticated client
		if err := mcpServer.RegisterTools(authenticatedClient); err != nil {
			api.Logger.Warn(r.Context(), "failed to register MCP tools", slog.Error(err))
		}

		// Handle the MCP request
		mcpServer.ServeHTTP(w, r)
	})
}

// mcpSSEHandler creates the MCP SSE transport handler
func (api *API) mcpSSEHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create MCP server instance
		mcpServer, err := mcp.NewServer(api.Logger.Named("mcp-sse"))
		if err != nil {
			api.Logger.Error(r.Context(), "failed to create MCP server", slog.Error(err))
			httpapi.Write(r.Context(), w, http.StatusInternalServerError, codersdk.Response{
				Message: "MCP server initialization failed",
			})
			return
		}

		authenticatedClient := codersdk.New(api.AccessURL)
		authenticatedClient.SetSessionToken(httpmw.APITokenFromRequest(r))

		// Register tools with authenticated client
		if err := mcpServer.RegisterTools(authenticatedClient); err != nil {
			api.Logger.Warn(r.Context(), "failed to register MCP tools", slog.Error(err))
		}

		// Create SSE server with appropriate configuration
		sseServer := server.NewSSEServer(mcpServer.MCPServer(),
			server.WithBaseURL(api.AccessURL.String()),
			server.WithStaticBasePath("/api/experimental/mcp"),
			server.WithSSEEndpoint("/sse"),
			server.WithMessageEndpoint("/message"),
			server.WithKeepAliveInterval(30*time.Second),
			server.WithKeepAlive(true),
		)

		// Handle the request based on the path
		path := r.URL.Path
		if strings.HasSuffix(path, "/sse") {
			sseServer.SSEHandler().ServeHTTP(w, r)
		} else if strings.HasSuffix(path, "/message") {
			sseServer.MessageHandler().ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}
