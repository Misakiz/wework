package api

import (
	"fmt"
	"sync"

	"github.com/Misakiz/wework/config"
)

var CORP_API_TYPE = map[string][]string{
	"GET_ACCESS_TOKEN":  {"/cgi-bin/gettoken", "GET"},
	"USER_CREATE":       {"/cgi-bin/user/create?access_token=ACCESS_TOKEN", "POST"},
	"USER_GET":          {"/cgi-bin/user/get?access_token=ACCESS_TOKEN", "GET"},
	"USER_UPDATE":       {"/cgi-bin/user/update?access_token=ACCESS_TOKEN", "POST"},
	"USER_DELETE":       {"/cgi-bin/user/delete?access_token=ACCESS_TOKEN", "GET"},
	"USER_BATCH_DELETE": {"/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN", "POST"},
	"USER_SIMPLE_LIST ": {"/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN", "GET"},
	"USER_LIST":         {"/cgi-bin/user/list?access_token=ACCESS_TOKEN", "GET"},
	"USERID_TO_OPENID":  {"/cgi-bin/user/convert_to_openid?access_token=ACCESS_TOKEN", "POST"},
	"OPENID_TO_USERID":  {"/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN", "POST"},
	"USER_AUTH_SUCCESS": {"/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN", "GET"},

	"DEPARTMENT_CREATE": {"/cgi-bin/department/create?access_token=ACCESS_TOKEN", "POST"},
	"DEPARTMENT_UPDATE": {"/cgi-bin/department/update?access_token=ACCESS_TOKEN", "POST"},
	"DEPARTMENT_DELETE": {"/cgi-bin/department/delete?access_token=ACCESS_TOKEN", "GET"},
	"DEPARTMENT_LIST":   {"/cgi-bin/department/list?access_token=ACCESS_TOKEN", "GET"},

	"TAG_CREATE":      {"/cgi-bin/tag/create?access_token=ACCESS_TOKEN", "POST"},
	"TAG_UPDATE":      {"/cgi-bin/tag/update?access_token=ACCESS_TOKEN", "POST"},
	"TAG_DELETE":      {"/cgi-bin/tag/delete?access_token=ACCESS_TOKEN", "GET"},
	"TAG_GET_USER":    {"/cgi-bin/tag/get?access_token=ACCESS_TOKEN", "GET"},
	"TAG_ADD_USER":    {"/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN", "POST"},
	"TAG_DELETE_USER": {"/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN", "POST"},
	"TAG_GET_LIST":    {"/cgi-bin/tag/list?access_token=ACCESS_TOKEN", "GET"},

	"BATCH_JOB_GET_RESULT": {"/cgi-bin/batch/getresult?access_token=ACCESS_TOKEN", "GET"},

	"BATCH_INVITE": {"/cgi-bin/batch/invite?access_token=ACCESS_TOKEN", "POST"},

	"AGENT_GET":      {"/cgi-bin/agent/get?access_token=ACCESS_TOKEN", "GET"},
	"AGENT_SET":      {"/cgi-bin/agent/set?access_token=ACCESS_TOKEN", "POST"},
	"AGENT_GET_LIST": {"/cgi-bin/agent/list?access_token=ACCESS_TOKEN", "GET"},

	"MENU_CREATE": {"/cgi-bin/menu/create?access_token=ACCESS_TOKEN", "POST"},
	"MENU_GET":    {"/cgi-bin/menu/get?access_token=ACCESS_TOKEN", "GET"},
	"MENU_DELETE": {"/cgi-bin/menu/delete?access_token=ACCESS_TOKEN", "GET"},

	"MESSAGE_SEND": {"/cgi-bin/message/send?access_token=ACCESS_TOKEN", "POST"},

	"MEDIA_GET": {"/cgi-bin/media/get?access_token=ACCESS_TOKEN", "GET"},

	"GET_USER_INFO_BY_CODE": {"/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN", "GET"},
	"GET_USER_DETAIL":       {"/cgi-bin/user/getuserdetail?access_token=ACCESS_TOKEN", "POST"},

	"GET_TICKET":       {"/cgi-bin/ticket/get?access_token=ACCESS_TOKEN", "GET"},
	"GET_JSAPI_TICKET": {"/cgi-bin/get_jsapi_ticket?access_token=ACCESS_TOKEN", "GET"},

	"GET_CHECKIN_OPTION":  {"/cgi-bin/checkin/getcheckinoption?access_token=ACCESS_TOKEN", "POST"},
	"GET_CHECKIN_DATA":    {"/cgi-bin/checkin/getcheckindata?access_token=ACCESS_TOKEN", "POST"},
	"GET_APPROVAL_DETAIL": {"/cgi-bin/oa/getapprovaldetail?access_token=ACCESS_TOKEN", "POST"},

	"GET_INVOICE_INFO":            {"/cgi-bin/card/invoice/reimburse/getinvoiceinfo?access_token=ACCESS_TOKEN", "POST"},
	"UPDATE_INVOICE_STATUS":       {"/cgi-bin/card/invoice/reimburse/updateinvoicestatus?access_token=ACCESS_TOKEN", "POST"},
	"BATCH_UPDATE_INVOICE_STATUS": {"/cgi-bin/card/invoice/reimburse/updatestatusbatch?access_token=ACCESS_TOKEN", "POST"},
	"BATCH_GET_INVOICE_INFO":      {"/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch?access_token=ACCESS_TOKEN", "POST"},
}

type CorpAPI struct {
	API
	CorpId      string
	Secret      string
	AccessToken string
	mutex       sync.Mutex
}

func NewCorpAPI(corpId, secret string) *CorpAPI {
	c := &CorpAPI{
		CorpId: corpId,
		Secret: secret,
	}
	c.API.AccessAPI = c
	return c
}

func (c *CorpAPI) GetAccessToken() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.AccessToken == "" {
		c.RefreshAccessToken()
	}
	return c.AccessToken
}

func (c *CorpAPI) RefreshAccessToken() {
	response, err := c.HttpCall(CORP_API_TYPE["GET_ACCESS_TOKEN"], map[string]interface{}{
		"corpid":     c.CorpId,
		"corpsecret": c.Secret,
	})
	if err != nil {
		if config.DEBUG {
			fmt.Println(err)
		}
		return
	}
	c.AccessToken, _ = response["access_token"].(string)
}
