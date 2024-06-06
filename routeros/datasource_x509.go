package routeros

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceX509() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceParseCertificate,
		Schema: map[string]*schema.Schema{
			"data": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "X509 certificate in PEM format.",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"akid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authority": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"common_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"digest_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invalid_after": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invalid_before": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "key_size": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			"key_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "key_usage": {
			// 	Type:     schema.TypeString,
			// 	Computed: true,
			// },
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subject_alt_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceParseCertificate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	re := regexp.MustCompile(`(?m)^\s+`)
	block, _ := pem.Decode(re.ReplaceAll([]byte(d.Get("data").(string)), nil))
	if block == nil {
		return diag.Errorf("Invalid PEM content")
	}

	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("akid", fmt.Sprintf("%x", c.AuthorityKeyId))
	d.Set("authority", c.IsCA)
	d.Set("common_name", c.Subject.CommonName)
	d.Set("digest_algorithm", c.SignatureAlgorithm.String())
	d.Set("fingerprint", fmt.Sprintf("%x", sha256.Sum256(c.Raw)))
	d.Set("invalid_after", c.NotAfter.String())
	d.Set("invalid_before", c.NotBefore.String())
	d.Set("issuer", c.Issuer.String())
	d.Set("key_type", c.PublicKeyAlgorithm.String())
	d.Set("issuer", c.Issuer.String())
	d.Set("serial_number", c.SerialNumber.Text(16))
	d.Set("skid", fmt.Sprintf("%x", c.SubjectKeyId))
	d.Set("signature_algorithm", c.SignatureAlgorithm.String())
	d.Set("subject", c.Subject.String())
	d.Set("subject_alt_name", getSANs(c))
	d.Set("version", c.Version)
	d.Set("pem", string(pem.EncodeToMemory(block)))

	d.SetId(c.SerialNumber.String())

	return nil
}

func getSANs(c *x509.Certificate) string {
	var res []string
	for _, v := range c.DNSNames {
		res = append(res, fmt.Sprintf("DNS:%v", v))
	}
	for _, v := range c.IPAddresses {
		res = append(res, fmt.Sprintf("IP:%v", v))
	}
	for _, v := range c.EmailAddresses {
		res = append(res, fmt.Sprintf("EMAIL:%v", v))
	}
	for _, v := range c.URIs {
		res = append(res, fmt.Sprintf("URI:%v", v))
	}
	return strings.Join(res, ",")
}
