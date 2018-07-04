package aci

import (
	"bytes"
	"fmt"
)

func rnPhysDom(dom string) string {
	return "phys-" + dom
}

func (c *Client) AttachableAccessEntityProfileDomainPhysAdd(aep, physdom string) error {

	me := "AttachableAccessEntityProfileDomainPhysAdd"

	rnE := rnAEP(aep)
	rn := rnPhysDom(physdom)

	api := "/api/node/mo/uni/infra/" + rnE + ".json"

	url := c.getURL(api)

	j := fmt.Sprintf(`{"infraRsDomP":{"attributes":{"tDn":"uni/%s","status":"created"}}}`, rn)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}

func (c *Client) AttachableAccessEntityProfileDomainPhysDel(aep, l2dom string) error {

	me := "AttachableAccessEntityProfileDomainPhysDel"

	rnE := rnAEP(aep)
	rn := rnPhysDom(l2dom)

	api := "/api/node/mo/uni/infra/" + rnE + ".json"

	url := c.getURL(api)

	j := fmt.Sprintf(`{"infraAttEntityP":{"attributes":{"dn":"uni/infra/%s","status":"modified"},"children":[{"infraRsDomP":{"attributes":{"dn":"uni/infra/%s/rsdomP-[uni/%s]","status":"deleted"}}}]}}`,
		rnE, rnE, rn)

	c.debugf("%s: url=%s json=%s", me, url, j)

	body, errPost := c.post(url, contentTypeJSON, bytes.NewBufferString(j))
	if errPost != nil {
		return fmt.Errorf("%s: %v", me, errPost)
	}

	c.debugf("%s: reply: %s", me, string(body))

	return parseJSONError(body)
}
