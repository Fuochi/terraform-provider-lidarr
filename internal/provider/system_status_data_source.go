package provider

import (
	"context"
	"fmt"

	"github.com/devopsarr/lidarr-go/lidarr"
	"github.com/devopsarr/terraform-provider-lidarr/tools"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const systemStatusDataSourceName = "system_status"

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SystemStatusDataSource{}

func NewSystemStatusDataSource() datasource.DataSource {
	return &SystemStatusDataSource{}
}

// SystemStatusDataSource defines the system status implementation.
type SystemStatusDataSource struct {
	client *lidarr.APIClient
}

// SystemStatus describes the system status data model.
type SystemStatus struct {
	RuntimeName            types.String `tfsdk:"runtime_name"`
	AppData                types.String `tfsdk:"app_data"`
	OsName                 types.String `tfsdk:"os_name"`
	PackageVersion         types.String `tfsdk:"package_version"`
	Mode                   types.String `tfsdk:"mode"`
	PackageUpdateMechanism types.String `tfsdk:"package_update_mechanism"`
	PackageAuthor          types.String `tfsdk:"package_author"`
	InstanceName           types.String `tfsdk:"instance_name"`
	AppName                types.String `tfsdk:"app_name"`
	URLBase                types.String `tfsdk:"url_base"`
	SqliteVersion          types.String `tfsdk:"sqlite_version"`
	Branch                 types.String `tfsdk:"branch"`
	StartupPath            types.String `tfsdk:"startup_path"`
	RuntimeVersion         types.String `tfsdk:"runtime_version"`
	StartTime              types.String `tfsdk:"start_time"`
	BuildTime              types.String `tfsdk:"build_time"`
	Version                types.String `tfsdk:"version"`
	Authentication         types.String `tfsdk:"authentication"`
	MigrationVersion       types.Int64  `tfsdk:"migration_version"`
	ID                     types.Int64  `tfsdk:"id"`
	IsDocker               types.Bool   `tfsdk:"is_docker"`
	IsDebug                types.Bool   `tfsdk:"is_debug"`
	IsNetCore              types.Bool   `tfsdk:"is_net_core"`
	IsAdmin                types.Bool   `tfsdk:"is_admin"`
	IsProduction           types.Bool   `tfsdk:"is_production"`
	IsWindows              types.Bool   `tfsdk:"is_windows"`
	IsOsx                  types.Bool   `tfsdk:"is_osx"`
	IsLinux                types.Bool   `tfsdk:"is_linux"`
	IsUserInteractive      types.Bool   `tfsdk:"is_user_interactive"`
}

func (d *SystemStatusDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_status"
}

func (d *SystemStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the delay server.
		MarkdownDescription: "<!-- subcategory:Status -->System Status resource. User must have rights to read `config.xml`.\nFor more information refer to [System Status](https://wiki.servarr.com/lidarr/system#status) documentation.",
		Attributes: map[string]schema.Attribute{
			// TODO: remove ID once framework support tests without ID https://www.terraform.io/plugin/framework/acctests#implement-id-attribute
			"id": schema.Int64Attribute{
				MarkdownDescription: "Delay Profile ID.",
				Computed:            true,
			},
			"is_debug": schema.BoolAttribute{
				MarkdownDescription: "Is debug flag.",
				Computed:            true,
			},
			"is_production": schema.BoolAttribute{
				MarkdownDescription: "Is production flag.",
				Computed:            true,
			},
			"is_admin": schema.BoolAttribute{
				MarkdownDescription: "Is admin flag.",
				Computed:            true,
			},
			"is_user_interactive": schema.BoolAttribute{
				MarkdownDescription: "Is user interactive flag.",
				Computed:            true,
			},
			"is_net_core": schema.BoolAttribute{
				MarkdownDescription: "Is net core flag.",
				Computed:            true,
			},
			"is_linux": schema.BoolAttribute{
				MarkdownDescription: "Is linux flag.",
				Computed:            true,
			},
			"is_osx": schema.BoolAttribute{
				MarkdownDescription: "Is osx flag.",
				Computed:            true,
			},
			"is_windows": schema.BoolAttribute{
				MarkdownDescription: "Is windows flag.",
				Computed:            true,
			},
			"is_docker": schema.BoolAttribute{
				MarkdownDescription: "Is docker flag.",
				Computed:            true,
			},
			"migration_version": schema.Int64Attribute{
				MarkdownDescription: "Migration version.",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Version.",
				Computed:            true,
			},
			"startup_path": schema.StringAttribute{
				MarkdownDescription: "Startup path.",
				Computed:            true,
			},
			"app_data": schema.StringAttribute{
				MarkdownDescription: "App data folder.",
				Computed:            true,
			},
			"os_name": schema.StringAttribute{
				MarkdownDescription: "OS name.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Mode.",
				Computed:            true,
			},
			"branch": schema.StringAttribute{
				MarkdownDescription: "Branch.",
				Computed:            true,
			},
			"authentication": schema.StringAttribute{
				MarkdownDescription: "Authentication.",
				Computed:            true,
			},
			"sqlite_version": schema.StringAttribute{
				MarkdownDescription: "SQLite version.",
				Computed:            true,
			},
			"url_base": schema.StringAttribute{
				MarkdownDescription: "Base URL.",
				Computed:            true,
			},
			"runtime_version": schema.StringAttribute{
				MarkdownDescription: "Runtime version.",
				Computed:            true,
			},
			"runtime_name": schema.StringAttribute{
				MarkdownDescription: "Runtime name.",
				Computed:            true,
			},
			"build_time": schema.StringAttribute{
				MarkdownDescription: "Build time.",
				Computed:            true,
			},
			"start_time": schema.StringAttribute{
				MarkdownDescription: "Start time.",
				Computed:            true,
			},
			"app_name": schema.StringAttribute{
				MarkdownDescription: "App name.",
				Computed:            true,
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: "Instance name.",
				Computed:            true,
			},
			"package_author": schema.StringAttribute{
				MarkdownDescription: "Package author.",
				Computed:            true,
			},
			"package_update_mechanism": schema.StringAttribute{
				MarkdownDescription: "Package update mechanism.",
				Computed:            true,
			},
			"package_version": schema.StringAttribute{
				MarkdownDescription: "Package version.",
				Computed:            true,
			},
		},
	}
}

func (d *SystemStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*lidarr.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			tools.UnexpectedDataSourceConfigureType,
			fmt.Sprintf("Expected *lidarr.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *SystemStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "NOT SUPPORTED")
	// TODO: add it back again once system status is present in sdk
	// Get system status current value
	// response, _, err := d.client.SystemApi.GetSystemStatus(ctx).Execute()
	// if err != nil {
	// 	resp.Diagnostics.AddError(tools.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", systemStatusDataSourceName, err))

	// 	return
	// }

	tflog.Trace(ctx, "read "+systemStatusDataSourceName)

	status := SystemStatus{}
	// status.write(response)
	resp.Diagnostics.Append(resp.State.Set(ctx, status)...)
}

// func (s *SystemStatus) write(status *lidarr.SystemResource) {
// 	s.IsDebug = types.BoolValue(status.IsDebug)
// 	s.IsProduction = types.BoolValue(status.IsProduction)
// 	s.IsAdmin = types.BoolValue(status.IsAdmin)
// 	s.IsUserInteractive = types.BoolValue(status.IsUserInteractive)
// 	s.IsNetCore = types.BoolValue(status.IsNetCore)
// 	s.IsLinux = types.BoolValue(status.IsLinux)
// 	s.IsOsx = types.BoolValue(status.IsOsx)
// 	s.IsWindows = types.BoolValue(status.IsWindows)
// 	s.IsDocker = types.BoolValue(status.IsDocker)
// 	s.ID = types.Int64Value(int64(1))
// 	s.MigrationVersion = types.Int64Value(status.MigrationVersion)
// 	s.Version = types.StringValue(status.Version)
// 	s.StartupPath = types.StringValue(status.StartupPath)
// 	s.AppData = types.StringValue(status.AppData)
// 	s.OsName = types.StringValue(status.OsName)
// 	s.Mode = types.StringValue(status.Mode)
// 	s.Branch = types.StringValue(status.Branch)
// 	s.Authentication = types.StringValue(status.Authentication)
// 	s.SqliteVersion = types.StringValue(status.SqliteVersion)
// 	s.URLBase = types.StringValue(status.URLBase)
// 	s.RuntimeVersion = types.StringValue(status.RuntimeVersion)
// 	s.RuntimeName = types.StringValue(status.RuntimeName)
// 	s.BuildTime = types.StringValue(status.BuildTime.String())
// 	s.StartTime = types.StringValue(status.StartTime.String())
// 	s.AppName = types.StringValue(status.AppName)
// 	s.InstanceName = types.StringValue(status.InstanceName)
// 	s.PackageAuthor = types.StringValue(status.PackageAuthor)
// 	s.PackageUpdateMechanism = types.StringValue(status.PackageUpdateMechanism)
// 	s.PackageVersion = types.StringValue(status.PackageVersion)
// }
