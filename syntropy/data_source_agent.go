package syntropy

import (
	"context"
	"github.com/SyntropyNet/syntropy-sdk-go/syntropy"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.DataSourceType = agentDataSourceType{}
var _ tfsdk.DataSource = agentDataSource{}

type agentDataSourceType struct{}

func (d agentDataSourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Datasource retrieves Syntropy agent data",
		Attributes: map[string]tfsdk.Attribute{
			"skip": {
				Type:        types.Int64Type,
				Optional:    true,
				Description: "Offset used for pagination",
			},
			"take": {
				Type:        types.Int64Type,
				Optional:    true,
				Description: "Number of items to take",
			},
			"search": {
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Agent name pattern. This will be used to filter out agent names that doesn't have specified patter",
			},
			"filter": {
				Description: "Syntropy agent filter",
				Optional:    true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Filter by agent ID",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.Int64Type,
						},
					},
					"tag_id": {
						Description: "Filter by agent tag ID",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.Int64Type,
						},
					},
					"provider_id": {
						Description: "Filter by agent provider ID",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.Int64Type,
						},
					},
					"type": {
						Description: "Filter by agent type",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"version": {
						Description: "Filter by agent version",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"tag_name": {
						Description: "Filter by agent tag name",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"status": {
						Description: "Filter by agent status",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"location_country": {
						Description: "Filter by agent location country",
						Optional:    true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"modified_at_from": {
						Description: "Filter by agent modified at from date",
						Optional:    true,
						Type:        types.StringType,
					},
					"modified_at_to": {
						Description: "Filter by agent modified at to date",
						Optional:    true,
						Type:        types.StringType,
					},
					"name": {
						Description: "Filter by agent modified at to date",
						Optional:    true,
						Type:        types.StringType,
					},
				}),
			},
			"agents": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Description: "Agent ID",
						Computed:    true,
						Type:        types.Int64Type,
					},
					"name": {
						Description: "Agent name",
						Computed:    true,
						Type:        types.StringType,
					},
					"public_ipv4": {
						Description: "Agent public IP",
						Type:        types.StringType,
						Computed:    true,
					},
					"status": {
						Description: "Agent status",
						Type:        types.StringType,
						Computed:    true,
					},
					"is_online": {
						Description: "Agent online status",
						Type:        types.BoolType,
						Computed:    true,
					},
					"version": {
						Description: "Agent version",
						Type:        types.StringType,
						Computed:    true,
					},
					"location_country": {
						Description: "Agent location country code",
						Type:        types.StringType,
						Computed:    true,
					},
					"location_city": {
						Description: "Agent city location",
						Type:        types.StringType,
						Computed:    true,
					},
					"device_id": {
						Description: "Agent device id",
						Type:        types.StringType,
						Computed:    true,
					},
					"is_virtual": {
						Description: "Is agent virtual",
						Type:        types.BoolType,
						Computed:    true,
					},
					"type": {
						Description: "Agent type",
						Type:        types.StringType,
						Computed:    true,
					},
					"modified_at": {
						Description: "Agent modified date",
						Type:        types.StringType,
						Computed:    true,
					},
					"tags": {
						Description: "Agent tags",
						Computed:    true,
						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
							"id": {
								Description: "Agent tag id",
								Type:        types.Int64Type,
								Computed:    true,
							},
							"name": {
								Description: "Agent tag name",
								Type:        types.StringType,
								Computed:    true,
							},
						}),
					},
					"provider": {
						Description: "Agent provider details",
						Computed:    true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"id": {
								Description: "Agent provider id",
								Type:        types.Int64Type,
								Computed:    true,
							},
							"name": {
								Description: "Agent provider name",
								Type:        types.StringType,
								Computed:    true,
							},
						}),
					},
				}),
			},
		},
	}, nil
}

func (d agentDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)
	return agentDataSource{
		provider: provider,
	}, diags
}

type agentDataSource struct {
	provider provider
}

func (d agentDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data AgentDataSource
	ctx = d.provider.createAuthContext(ctx)
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	agentFilter := &syntropy.V1AgentFilter{}

	if data.Filter != nil {
		filter, err := flattenAgentFilter(*data.Filter)
		if err != nil {
			resp.Diagnostics.AddError("Error while parsing agent filter data", err.Error())
			return
		}
		agentFilter = filter
	}

	skip := int32(data.Skip.Value)
	take := int32(data.Take.Value)
	aResp, _, err := d.provider.client.AgentsApi.V1NetworkAgentsSearch(ctx).V1NetworkAgentsSearchRequest(syntropy.V1NetworkAgentsSearchRequest{
		Filter: agentFilter,
		Order:  nil,
		Skip:   &skip,
		Take:   &take,
		Search: &data.Search.Value,
	}).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error while getting Syntropy agent", err.Error())
		return
	}

	for _, agent := range aResp.Data {
		data.Agents = append(data.Agents, Agent{
			ID:              int64(agent.AgentId),
			Name:            agent.AgentName,
			PublicIPv4:      agent.AgentPublicIpv4,
			Status:          NullableAgentStatusToString(agent.AgentStatus),
			IsOnline:        agent.AgentIsOnline,
			Version:         agent.AgentVersion,
			LocationCountry: NullableStringToString(agent.AgentLocationCountry),
			LocationCity:    NullableStringToString(agent.AgentLocationCity),
			DeviceID:        agent.AgentDeviceId,
			IsVirtual:       agent.AgentIsVirtual,
			Type:            string(agent.AgentType),
			ModifiedAt:      agent.AgentModifiedAt.String(),
			Tags:            convertAgentTagsToTfValue(agent.AgentTags),
			AgentProvider: AgentProvider{
				ID:   int64(agent.AgentProvider.AgentProviderId),
				Name: agent.AgentProvider.AgentProviderName,
			},
		})
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func flattenAgentFilter(in AgentFilter) (*syntropy.V1AgentFilter, error) {
	out := &syntropy.V1AgentFilter{
		AgentName: in.Name,
	}

	if in.ID != nil {
		out.AgentId = int64ArrayToInt32Array(*in.ID)
	}

	if in.ProviderID != nil {
		out.AgentProviderId = int64ArrayToInt32Array(*in.ProviderID)
	}

	if in.TagID != nil {
		out.AgentTagId = int64ArrayToInt32Array(*in.TagID)
	}

	if in.Type != nil {
		out.AgentType = stringArrayToAgentTypeArray(*in.Type)
	}

	if in.Version != nil {
		out.AgentVersion = *in.Version
	}

	if in.TagName != nil {
		out.AgentTagName = *in.TagName
	}

	if in.LocationCountry != nil {
		out.AgentLocationCountry = *in.LocationCountry
	}

	if in.Status != nil {
		out.AgentStatus = stringArrayToAgentStatusArray(*in.Status)
	}

	if in.ModifiedAtFrom != nil {
		modifiedAtFrom, err := tfValueToDateP(*in.ModifiedAtFrom)
		if err != nil {
			return nil, err
		}
		out.AgentModifiedAtFrom = modifiedAtFrom
	}

	if in.ModifiedAtTo != nil {
		modifiedAtTo, err := tfValueToDateP(*in.ModifiedAtTo)
		if err != nil {
			return nil, err
		}
		out.AgentModifiedAtTo = modifiedAtTo
	}
	return out, nil
}
