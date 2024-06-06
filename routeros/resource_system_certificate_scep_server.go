package routeros

import (
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceCertificateScepServer() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/certificate/scep-server"),
		MetaId:           PropId(Id),

		"ca_cert": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the CA certificate to use.",
		},
		"path": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile("^/scep/"), "path must start with /scep/"),
			Description:  "HTTP path starting with `/scep/`.",
		},
		"next_ca_cert": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			Description:      "Name of the next CA certificate or `none`.",
		},
		"days_valid": {
			Type:             schema.TypeInt,
			Optional:         true,
			DiffSuppressFunc: AlwaysPresentNotUserProvided,
			ValidateFunc:     validation.IntAtLeast(0),
			Description:      "The number of days to sign certificates for.",
		},
		"request_lifetime": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: TimeEquall,
			ValidateDiagFunc: ValidationDurationAtLeast(time.Minute * 5),
			Description:      "Request lifetime (5m minimum).",
		},
		KeyDisabled: PropDisabledRw,
	}

	return &schema.Resource{
		Schema:        resSchema,
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
