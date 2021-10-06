/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/bulkexports/v1"
)

func ResourceExportsJobs() *schema.Resource {
	return &schema.Resource{
		CreateContext: createExportsJobs,
		ReadContext:   readExportsJobs,
		DeleteContext: deleteExportsJobs,
		Schema: map[string]*schema.Schema{
			"resource_type":  AsString(SchemaForceNewRequired),
			"end_day":        AsString(SchemaForceNewRequired),
			"friendly_name":  AsString(SchemaForceNewRequired),
			"start_day":      AsString(SchemaForceNewRequired),
			"email":          AsString(SchemaForceNewOptional),
			"webhook_method": AsString(SchemaForceNewOptional),
			"webhook_url":    AsString(SchemaForceNewOptional),
			"job_sid":        AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseExportsJobsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createExportsJobs(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateExportCustomJobParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	resourceType := d.Get("resource_type").(string)

	r, err := m.(*client.Config).Client.BulkexportsV1.CreateExportCustomJob(resourceType, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{resourceType}
	idParts = append(idParts, (*r.JobSid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteExportsJobs(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	jobSid := d.Get("job_sid").(string)

	err := m.(*client.Config).Client.BulkexportsV1.DeleteJob(jobSid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readExportsJobs(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	jobSid := d.Get("job_sid").(string)

	r, err := m.(*client.Config).Client.BulkexportsV1.FetchJob(jobSid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseExportsJobsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected job_sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("job_sid", importParts[0])

	return nil
}
