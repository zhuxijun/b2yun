package http

import (
	root "b2yun/pkg"
)

// // StringPair 签名值对
// type StringPair map[string]string

// const signTagName = "sign"
// const jsonTagName = "json"

// // GetSign 获取签名
// func (p StringPair) GetSign(token string) string {

// 	linkStr := p.getLinkStr()

// 	md5Str := linkStr + token

// 	cryptor := crypto.NewCrypto()

// 	sign, _ := cryptor.MD5Crypto().Salt(md5Str)

// 	return sign

// }

// func (p StringPair) getLinkStr() string {

// 	var linkStr string

// 	keys := make([]string, len(p))

// 	for key := range p {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)

// 	for _, key := range keys {

// 		if key == "" || p[key] == "" {
// 			continue
// 		}

// 		linkStr = linkStr + key + "=" + p[key] + "&"
// 	}

// 	linkStr = strings.TrimRight(linkStr, "&")

// 	return linkStr

// }

// // StructToSignPair 结构体转字符对
// func StructToSignPair(i interface{}) StringPair {

// 	pair := make(StringPair)

// 	t := reflect.TypeOf(i)

// 	v := reflect.ValueOf(i)

// 	for i := 0; i < t.NumField(); i++ {

// 		name := t.Field(i).Name

// 		jsonTag := t.Field(i).Tag.Get(jsonTagName)

// 		signTag := t.Field(i).Tag.Get(signTagName)

// 		if jsonTag != "" {
// 			name = jsonTag
// 		}

// 		if signTag == "" || signTag == "1" {
// 			value := v.Field(i)
// 			if value.Interface() != reflect.Zero(value.Type()).Interface() {
// 				switch value.Interface().(type) {
// 				case root.JSONTime:
// 					pair[name] = value.Interface().(root.JSONTime).Marshal()
// 				default:
// 					pair[name] = fmt.Sprintf("%v", value)
// 				}

// 			}

// 		}

// 	}
// 	return pair
// }

// CommonResponse POST类型或GET无返回数据类型的通用返回值
type CommonResponse struct {
	ErrCode     int    `json:"status"`
	ErrMsg      string `json:"message"`
	TimeStamp   int64  `json:"timestamp"`
	AccessToken string `json:"access_token"`
	Pages       string `json:"pages"`
	Next        bool   `json:"next"`
}

// CommonResponseOrder GET返回的订单详情
type CommonResponseOrder struct {
	ErrCode     int             `json:"status"`
	ErrMsg      string          `json:"message"`
	TimeStamp   int64           `json:"timestamp"`
	AccessToken string          `json:"access_token"`
	Data        []root.ReqOrder `json:"data"`
	Pages       string          `json:"pages"`
	Next        bool            `json:"next"`
}

// CommonResponseMemberLevel GET返回的会员等级
type CommonResponseMemberLevel struct {
	ErrCode     int                   `json:"status"`
	ErrMsg      string                `json:"message"`
	TimeStamp   int64                 `json:"timestamp"`
	AccessToken string                `json:"access_token"`
	Data        []root.ReqMemberLevel `json:"data"`
	Pages       string                `json:"pages"`
	Next        bool                  `json:"next"`
}

// CommonResponseMemberInfo GET返回的会员信息
type CommonResponseMemberInfo struct {
	ErrCode     int                  `json:"status"`
	ErrMsg      string               `json:"message"`
	TimeStamp   int64                `json:"timestamp"`
	AccessToken string               `json:"access_token"`
	Data        []root.ReqMemberInfo `json:"data"`
	Pages       string               `json:"pages"`
	Next        bool                 `json:"next"`
}
