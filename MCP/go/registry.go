package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_current "github.com/input-api/mcp-server/tools/current"
	tools_count "github.com/input-api/mcp-server/tools/count"
	tools_info "github.com/input-api/mcp-server/tools/info"
	tools_plugin "github.com/input-api/mcp-server/tools/plugin"
	tools_deviceid "github.com/input-api/mcp-server/tools/deviceid"
	tools_create "github.com/input-api/mcp-server/tools/create"
	tools_id "github.com/input-api/mcp-server/tools/id"
	tools_cache "github.com/input-api/mcp-server/tools/cache"
	tools_devicetype "github.com/input-api/mcp-server/tools/devicetype"
	tools_name "github.com/input-api/mcp-server/tools/name"
	tools_api "github.com/input-api/mcp-server/tools/api"
	tools_notification "github.com/input-api/mcp-server/tools/notification"
	tools_config "github.com/input-api/mcp-server/tools/config"
	tools_user "github.com/input-api/mcp-server/tools/user"
	tools_configuration "github.com/input-api/mcp-server/tools/configuration"
	tools_command "github.com/input-api/mcp-server/tools/command"
	tools_device "github.com/input-api/mcp-server/tools/device"
	tools_refresh "github.com/input-api/mcp-server/tools/refresh"
	tools_network "github.com/input-api/mcp-server/tools/network"
	tools_token "github.com/input-api/mcp-server/tools/token"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_current.CreateGet_currentTool(cfg),
		tools_count.CreateGet_countTool(cfg),
		tools_info.CreateGet_infoTool(cfg),
		tools_plugin.CreateGet_plugin_authenticateTool(cfg),
		tools_deviceid.CreateGet_deviceid_commandTool(cfg),
		tools_plugin.CreateGet_plugin_createTool(cfg),
		tools_create.CreateGet_createTool(cfg),
		tools_id.CreateGet_id_devicetype_allTool(cfg),
		tools_cache.CreateGet_cacheTool(cfg),
		tools_deviceid.CreateGet_deviceid_command_pollTool(cfg),
		tools_devicetype.CreateGet_devicetypeTool(cfg),
		tools_name.CreateGet_nameTool(cfg),
		tools_id.CreateGet_id_devicetypeTool(cfg),
		tools_id.CreateGet_id_network_networkidTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_notification.CreateGet_notification_pollTool(cfg),
		tools_config.CreateGet_config_clusterTool(cfg),
		tools_user.CreateGet_userTool(cfg),
		tools_id.CreateGet_idTool(cfg),
		tools_deviceid.CreateGet_deviceid_notificationTool(cfg),
		tools_plugin.CreateGet_pluginTool(cfg),
		tools_deviceid.CreateGet_deviceid_notification_idTool(cfg),
		tools_configuration.CreateGet_configurationTool(cfg),
		tools_command.CreateGet_command_pollTool(cfg),
		tools_id.CreateGet_id_devicetype_devicetypeidTool(cfg),
		tools_device.CreateGet_deviceTool(cfg),
		tools_refresh.CreateGet_refreshTool(cfg),
		tools_deviceid.CreateGet_deviceid_command_commandidTool(cfg),
		tools_deviceid.CreateGet_deviceid_command_commandid_pollTool(cfg),
		tools_network.CreateGet_networkTool(cfg),
		tools_token.CreateGet_tokenTool(cfg),
		tools_deviceid.CreateGet_deviceid_notification_pollTool(cfg),
	}
}
