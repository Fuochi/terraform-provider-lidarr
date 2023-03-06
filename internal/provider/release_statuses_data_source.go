package provider

import (
	"context"
	"strconv"

	"github.com/devopsarr/lidarr-go/lidarr"
	"github.com/devopsarr/terraform-provider-lidarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const releaseStatusesDataSourceName = "release_statuses"

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ReleaseStatusesDataSource{}

func NewReleaseStatusesDataSource() datasource.DataSource {
	return &ReleaseStatusesDataSource{}
}

// ReleaseStatusesDataSource defines the releaseStatus implementation.
type ReleaseStatusesDataSource struct {
	client *lidarr.APIClient
}

func (d *ReleaseStatusesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + releaseStatusesDataSourceName
}

func (d *ReleaseStatusesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<!-- subcategory:Profiles -->List all available [Release Status](../data-sources/release_status).",
		Attributes: map[string]schema.Attribute{
			// TODO: remove ID once framework support tests without ID https://www.terraform.io/plugin/framework/acctests#implement-id-attribute
			"id": schema.StringAttribute{
				Computed: true,
			},
			"elements": schema.SetNestedAttribute{
				MarkdownDescription: "Release status list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Release status ID.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Release status name.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ReleaseStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if client := helpers.DataSourceConfigure(ctx, req, resp); client != nil {
		d.client = client
	}
}

func (d *ReleaseStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MetadataProfileElements

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get release status type current value
	response, _, err := d.client.MetadataProfileSchemaApi.GetMetadataprofileSchema(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, releaseStatusesDataSourceName, err))

		return
	}

	statuses := response.GetReleaseStatuses()

	tflog.Trace(ctx, "read "+releaseStatusesDataSourceName)
	// Map response body to resource schema attribute
	releaseTypes := make([]MetadataProfileElement, len(statuses))
	for i, t := range statuses {
		releaseTypes[i].writeRelease(t.ReleaseStatus)
	}

	tfsdk.ValueFrom(ctx, releaseTypes, data.Elements.Type(ctx), &data.Elements)
	// TODO: remove ID once framework support tests without ID https://www.terraform.io/plugin/framework/acctests#implement-id-attribute
	data.ID = types.StringValue(strconv.Itoa(len(statuses)))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
