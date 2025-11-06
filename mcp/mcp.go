package mcp

import (
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

type MCP struct{}

func NewMCP() *MCP {
	return &MCP{}
}

func InitMCP() *goMcp.Server {
	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    "JueJin-MCP",
		Version: "0.0.1",
	}, nil)

	// 添加登录状态工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "login status",
		Description: "获取登录状态",
	}, LoginStatus)
	return server
}
