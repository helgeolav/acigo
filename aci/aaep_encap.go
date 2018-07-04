package aci

import (
	"bytes"
	"fmt"
	"strconv"
)

// Several methods to work on EPGs mapping in an AAEP

// Method that makes an VLAN encap-string from a VLAN ID
func GetVLANEncap(vlan int) string {
	return "vlan-" + strconv.Itoa(vlan)
}

// Method that makes an VXLAN encap-string from a VNID
func GetVXLANEncap(vnid int) string {
	return "vxlan-" + strconv.Itoa(vnid)
}

// AttachableAccessEntityProfileAdd creates an AAEP mapping entry.
func (c *Client) AttachableAccessEntityProfileEncapAdd(aep, tenant, applicationProfile, epg, encap string) error {

	me := "AttachableAccessEntityProfileEncapAdd"
	rn := rnAEP(aep)
	api := "/api/node/mo/uni/infra/" + rn + "/gen-default.json"
	dnE := dnAEPG(tenant, applicationProfile, epg)
	url := c.getURL(api)
	j := fmt.Sprintf(`{"infraRsFuncToEpg":{"attributes":{"tDn":"uni/%s","status":"created,modified", "encap": "%s"}}}`,
		dnE, encap)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}

// delete AAEP mapping entry
func (c *Client) AttachableAccessEntityProfileEncapDel(aep, tenant, applicationProfile, epg, encap string) error {

	me := "AttachableAccessEntityProfileEncapDel"
	rn := rnAEP(aep)
	api := "/api/node/mo/uni/infra/" + rn + "/gen-default.json"
	dnE := dnAEPG(tenant, applicationProfile, epg)
	url := c.getURL(api)
	j := fmt.Sprintf(`{"infraRsFuncToEpg":{"attributes":{"tDn":"uni/%s","status":"deleted", "encap": "%s"}}}`,
		dnE, encap)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}

// TODO: make a list of all encaps
