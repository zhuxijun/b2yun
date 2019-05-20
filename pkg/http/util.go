package http

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

// CommonResponse 门店返回值
type CommonResponse struct {
	ErrCode     int    `json:"status"`
	ErrMsg      string `json:"message"`
	TimeStamp   int64  `json:"timestamp"`
	AccessToken string `json:"access_token"`
	Data        string `json:"data"`
	Pages       int    `json:"pages"`
	Next        bool   `json:"next"`
}
