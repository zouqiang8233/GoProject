package main

import (
	"encoding/xml"
)

type ExtendInfo struct {
	BankPassword  string `xml:"bank_password,omitempty"`   // 银行密码
	Sex           string `xml:"sex,omitempty"`             // 性别
	LegalName     string `xml:"legal_name,omitempty"`      // 法人姓名
	LegalCertType string `xml:"legal_cert_type,omitempty"` // 法人证件类型
	LegalCertCode string `xml:"legal_cert_code,omitempty"` // 法人证件号码
	AgentName     string `xml:"agent_name,omitempty"`      // 经办人姓名
	AgentCertType string `xml:"agent_cert_type,omitempty"` // 经办人证件类型
	AgentCertCode string `xml:"agent_cert_code,omitempty"` // 经办人证件号码
	OrgMobile     string `xml:"org_mobile,omitempty"`      // 原手机号
}

type SignReqBody struct {
	ExchNo     string      `xml:"exch_no,omitempty"`      // 交易所编号
	ExchDate   string      `xml:"exch_date,omitempty"`    // 交易所业务日期
	ExchSeq    string      `xml:"exch_seq,omitempty"`     // 交易所流水号
	TradeAcct  string      `xml:"trade_acct,omitempty"`   // 交易账号
	TradeNo    string      `xml:"tran_no,omitempty"`      // 银行业务编号
	Acct       string      `xml:"acct,omitempty"`         // 银行账号
	AcctName   string      `xml:"acct_name,omitempty"`    // 银行账户名
	Currency   string      `xml:"currency,omitempty"`     // 币种
	CardBankNo string      `xml:"card_bank_no,omitempty"` // 银行卡行号
	CardAcct   string      `xml:"card_acct,omitempty"`    // 银行卡号
	CardName   string      `xml:"card_name,omitempty"`    // 银行卡户名
	AcctType   string      `xml:"acct_type,omitempty"`    // 银行账户类型
	CertType   string      `xml:"cert_type,omitempty"`    // 证件类型
	CertCode   string      `xml:"cert_code,omitempty"`    // 证件号码
	ClientName string      `xml:"client_name,omitempty"`  // 客户名称
	Mobile     string      `xml:"mobile,omitempty"`       // 手机号码
	Email      string      `xml:"email,omitempty"`        // 电子邮箱
	ChangeType string      `xml:"change_type,omitempty"`  // 变更类型
	IsForce    string      `xml:"is_force,omitempty"`     // 是否强制
	Exinfo     *ExtendInfo `xml:"extend_info"`
}

type SignReq struct {
	XMLName xml.Name     `xml:"message"`
	Header  *MsgHeader   `xml:"head"`
	Bodyer  *SignReqBody `xml:"body"`
}
