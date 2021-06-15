/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/rest/ip_messaging/v2"
)

func ResourceServicesChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannels,
		ReadContext:   readServicesChannels,
		UpdateContext: updateServicesChannels,
		DeleteContext: deleteServicesChannels,
		Schema: map[string]*schema.Schema{
			"service_sid":             AsString(SchemaRequired),
			"xtwilio_webhook_enabled": AsString(SchemaOptional),
			"attributes":              AsString(SchemaOptional),
			"created_by":              AsString(SchemaOptional),
			"date_created":            AsString(SchemaOptional),
			"date_updated":            AsString(SchemaOptional),
			"friendly_name":           AsString(SchemaOptional),
			"type":                    AsString(SchemaOptional),
			"unique_name":             AsString(SchemaOptional),
			"account_sid":             AsString(SchemaComputed),
			"links":                   AsString(SchemaComputed),
			"members_count":           AsInt(SchemaComputed),
			"messages_count":          AsInt(SchemaComputed),
			"sid":                     AsString(SchemaComputed),
			"url":                     AsString(SchemaComputed),
		},
	}
}

func createServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateChannel(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteChannel(serviceSid, sid, &params)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchChannel(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateChannel(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannelsWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsWebhooks,
		ReadContext:   readServicesChannelsWebhooks,
		UpdateContext: updateServicesChannelsWebhooks,
		DeleteContext: deleteServicesChannelsWebhooks,
		Schema: map[string]*schema.Schema{
			"service_sid":               AsString(SchemaRequired),
			"channel_sid":               AsString(SchemaRequired),
			"configuration_filters":     AsString(SchemaOptional),
			"configuration_flow_sid":    AsString(SchemaOptional),
			"configuration_method":      AsString(SchemaOptional),
			"configuration_retry_count": AsInt(SchemaOptional),
			"configuration_triggers":    AsString(SchemaOptional),
			"configuration_url":         AsString(SchemaOptional),
			"type":                      AsString(SchemaOptional),
			"account_sid":               AsString(SchemaComputed),
			"configuration":             AsString(SchemaComputed),
			"date_created":              AsString(SchemaComputed),
			"date_updated":              AsString(SchemaComputed),
			"sid":                       AsString(SchemaComputed),
			"url":                       AsString(SchemaComputed),
		},
	}
}

func createServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateChannelWebhook(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteChannelWebhook(serviceSid, channelSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchChannelWebhook(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateChannelWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateChannelWebhook(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentials,
		ReadContext:   readCredentials,
		UpdateContext: updateCredentials,
		DeleteContext: deleteCredentials,
		Schema: map[string]*schema.Schema{
			"api_key":       AsString(SchemaOptional),
			"certificate":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"private_key":   AsString(SchemaOptional),
			"sandbox":       AsBool(SchemaOptional),
			"secret":        AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateCredential(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteCredential(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateCredential(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannelsMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMembers,
		ReadContext:   readServicesChannelsMembers,
		UpdateContext: updateServicesChannelsMembers,
		DeleteContext: deleteServicesChannelsMembers,
		Schema: map[string]*schema.Schema{
			"service_sid":                 AsString(SchemaRequired),
			"channel_sid":                 AsString(SchemaRequired),
			"xtwilio_webhook_enabled":     AsString(SchemaOptional),
			"attributes":                  AsString(SchemaOptional),
			"date_created":                AsString(SchemaOptional),
			"date_updated":                AsString(SchemaOptional),
			"identity":                    AsString(SchemaOptional),
			"last_consumed_message_index": AsInt(SchemaOptional),
			"last_consumption_timestamp":  AsString(SchemaOptional),
			"role_sid":                    AsString(SchemaOptional),
			"account_sid":                 AsString(SchemaComputed),
			"sid":                         AsString(SchemaComputed),
			"url":                         AsString(SchemaComputed),
		},
	}
}

func createServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateMember(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteMember(serviceSid, channelSid, sid, &params)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchMember(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateMember(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannelsMessages() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMessages,
		ReadContext:   readServicesChannelsMessages,
		UpdateContext: updateServicesChannelsMessages,
		DeleteContext: deleteServicesChannelsMessages,
		Schema: map[string]*schema.Schema{
			"service_sid":             AsString(SchemaRequired),
			"channel_sid":             AsString(SchemaRequired),
			"xtwilio_webhook_enabled": AsString(SchemaOptional),
			"attributes":              AsString(SchemaOptional),
			"body":                    AsString(SchemaOptional),
			"date_created":            AsString(SchemaOptional),
			"date_updated":            AsString(SchemaOptional),
			"from":                    AsString(SchemaOptional),
			"last_updated_by":         AsString(SchemaOptional),
			"media_sid":               AsString(SchemaOptional),
			"account_sid":             AsString(SchemaComputed),
			"index":                   AsInt(SchemaComputed),
			"media":                   AsString(SchemaComputed),
			"sid":                     AsString(SchemaComputed),
			"to":                      AsString(SchemaComputed),
			"type":                    AsString(SchemaComputed),
			"url":                     AsString(SchemaComputed),
			"was_edited":              AsBool(SchemaComputed),
		},
	}
}

func createServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateMessage(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteMessage(serviceSid, channelSid, sid, &params)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchMessage(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateMessage(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRoles,
		ReadContext:   readServicesRoles,
		UpdateContext: updateServicesRoles,
		DeleteContext: deleteServicesRoles,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaOptional),
			"permission":    AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"permissions":   AsList(AsString(SchemaComputed), SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateRole(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteRole(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchRole(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateRole(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":                    AsString(SchemaOptional),
			"account_sid":                      AsString(SchemaComputed),
			"consumption_report_interval":      AsInt(SchemaComputed),
			"date_created":                     AsString(SchemaComputed),
			"date_updated":                     AsString(SchemaComputed),
			"default_channel_creator_role_sid": AsString(SchemaComputed),
			"default_channel_role_sid":         AsString(SchemaComputed),
			"default_service_role_sid":         AsString(SchemaComputed),
			"limits":                           AsString(SchemaComputed),
			"links":                            AsString(SchemaComputed),
			"media":                            AsString(SchemaComputed),
			"notifications":                    AsString(SchemaComputed),
			"post_webhook_retry_count":         AsInt(SchemaComputed),
			"post_webhook_url":                 AsString(SchemaComputed),
			"pre_webhook_retry_count":          AsInt(SchemaComputed),
			"pre_webhook_url":                  AsString(SchemaComputed),
			"reachability_enabled":             AsBool(SchemaComputed),
			"read_status_enabled":              AsBool(SchemaComputed),
			"sid":                              AsString(SchemaComputed),
			"typing_indicator_timeout":         AsInt(SchemaComputed),
			"url":                              AsString(SchemaComputed),
			"webhook_filters":                  AsList(AsString(SchemaComputed), SchemaComputed),
			"webhook_method":                   AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesUsers,
		ReadContext:   readServicesUsers,
		UpdateContext: updateServicesUsers,
		DeleteContext: deleteServicesUsers,
		Schema: map[string]*schema.Schema{
			"service_sid":             AsString(SchemaRequired),
			"xtwilio_webhook_enabled": AsString(SchemaOptional),
			"attributes":              AsString(SchemaOptional),
			"friendly_name":           AsString(SchemaOptional),
			"identity":                AsString(SchemaOptional),
			"role_sid":                AsString(SchemaOptional),
			"account_sid":             AsString(SchemaComputed),
			"date_created":            AsString(SchemaComputed),
			"date_updated":            AsString(SchemaComputed),
			"is_notifiable":           AsBool(SchemaComputed),
			"is_online":               AsBool(SchemaComputed),
			"joined_channels_count":   AsInt(SchemaComputed),
			"links":                   AsString(SchemaComputed),
			"sid":                     AsString(SchemaComputed),
			"url":                     AsString(SchemaComputed),
		},
	}
}

func createServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateUser(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteUser(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchUser(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateUser(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
