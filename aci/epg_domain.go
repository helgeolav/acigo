package aci

import (
	"bytes"
	"fmt"
)

// Several methods to edit domain associations on an EPG

// Add a physical domain to an EPG
func (c *Client) ApplicationEpgPhysDomainAdd(tenant, applicationProfile, epg, domain string) error {
	me := "ApplicationEpgDomainAdd"
	dnE := dnAEPG(tenant, applicationProfile, epg)
	api := "/api/node/mo/uni/" + dnE + ".json"

	url := c.getURL(api)

	j := fmt.Sprintf(`{"fvRsDomAtt":{"attributes":{"resImedcy":"immediate", "tDn": "uni/phys-%s"}}}`, domain)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}

// Remove a physical domain from an EPG
func (c *Client) ApplicationEpgPhysDomainDel(tenant, applicationProfile, epg, domain string) error {
	me := "ApplicationEpgDomainDel"
	dnE := dnAEPG(tenant, applicationProfile, epg)
	api := "/api/node/mo/uni/" + dnE + ".json"

	url := c.getURL(api)

	j := fmt.Sprintf(`{"fvRsDomAtt":{"attributes":{"resImedcy":"immediate", "tDn": "uni/phys-%s", "status": "deleted"}}}`, domain)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}

// TODO: make a call to list out all domains attached to an EPG
