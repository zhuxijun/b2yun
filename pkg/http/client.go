package http

import (
	root "b2yun/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client http客户端，负责组织http请求服务
type Client struct {
	URL url.URL

	catService    CatService
	brandService  BrandService
	goodsService  GoodsService
	orderService  OrderService
	memberService MemberService

	taskService  root.Tasker
	basicService root.Basicer

	configer root.Configer
	log      *root.Log
}

// NewClient 创建http客户端请求
func NewClient(configer root.Configer, log *root.Log) *Client {

	c := &Client{}
	c.configer = configer
	c.log = log

	c.catService.client = c
	c.brandService.client = c
	c.goodsService.client = c
	c.orderService.client = c
	c.memberService.client = c

	return c

}

// Open 打开客户端
func (c *Client) Open() error {

	config, err := c.configer.GetConfig()
	if err != nil {
		return err
	}

	c.URL.Host = config.HTTP.BaseHost
	c.URL.Path = config.HTTP.BasePath

	return nil
}

// ServiceOption Client创建参数
type ServiceOption struct {
	Basicer root.Basicer

	Cater    root.Cater
	Brander  root.Brander
	Goodser  root.Goodser
	Orderer  root.Orderer
	Memberer root.Memberer

	Tasker root.Tasker
}

// InitService 初始化服务
func (c *Client) InitService(o ServiceOption) {

	c.catService.catService = o.Cater
	c.brandService.brandService = o.Brander
	c.goodsService.goodsService = o.Goodser
	c.orderService.orderService = o.Orderer
	c.memberService.memberService = o.Memberer

	c.taskService = o.Tasker
	c.basicService = o.Basicer
}

// Post 发送POST请求
// url 请求地址（不必传入基础地址,最终地址为 BaseHost + BasePath + url）
// reqStr 请求需要传入的字符串，记录http.body
func (c *Client) Post(url string, reqStr string) error {

	reqBody := bytes.NewBufferString(reqStr)

	reqURL, err := c.getReqURL(url)
	if err != nil {
		return err
	}

	c.log.Logger.Infof("本次请求地址:[%s]", reqURL)
	c.log.Logger.Infof("本次请求参数:[%s]", reqStr)

	contentType := "application/json;charset=UTF-8" //"application/x-www-form-urlencoded"
	resp, err := http.Post(reqURL, contentType, reqBody)
	if err != nil {
		return err
	}

	respStr, err := handleResponse(resp)
	if err != nil {
		c.log.Logger.Errorf("本次请求出现错误,错误信息[%s]", err.Error())
		return err
	}

	c.log.Logger.Infof("本次请求结束,返回结果:[%s]", respStr)
	return nil
}

// Get 发送GET请求
// url 请求地址（不必传入基础地址,最终地址为 BaseHost + BasePath + url）
// reqStr 请求需要传入的字符串，记录http.body
func (c *Client) Get(url string) (string, error) {

	reqURL, err := c.getReqURL(url)
	if err != nil {
		return "", err
	}

	c.log.Logger.Infof("本次请求地址:[%s]", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		return "", err
	}

	respStr, err := handleResponse(resp)
	if err != nil {
		c.log.Logger.Errorf("本次请求出现错误,错误信息[%s]", err.Error())
		return "", err
	}

	c.log.Logger.Infof("本次请求结束,返回结果:[%s]", respStr)
	return respStr, nil
}

// getReqURL 获取请求地址
func (c *Client) getReqURL(path string) (string, error) {
	//判断token是否需要重新获取
	basic, err := c.getToken()
	if err != nil {
		return "", err
	}

	url := c.URL
	url.Path = url.Path + path + "&access_token=" + basic.Token + "&timeStamp=" + strconv.FormatInt(basic.TimeStamp, 10)
	return url.Host + url.Path, nil
}

//获取token
func (c *Client) getToken() (root.Basic, error) {

	basic, err := c.basicService.GetBasic()
	if err != nil {
		return basic, err
	}

	//1、判定token是否已过期(token的时间戳+有效秒数>当前时间戳，表示还未过期，不用往下走重新从服务器获取)
	if basic.TimeStamp+basic.TokenValidTime > time.Now().Unix() {
		return basic, nil
	}

	//2、获取token写入配置文件
	url := c.URL

	path := "/token/index.php?"
	reqURL := url.Host + url.Path + path + "appid=" + basic.AppID + "&secret=" + basic.Secret

	c.log.Logger.Infof("本次请求地址:[%s]", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		return basic, err
	}

	basic2, err := c.handleResponseToken(resp)
	if err != nil {
		//c.log.Logger.Errorf("本次请求出现错误,错误信息[%s]", err.Error())
		return basic2, err
	}

	//在此处将获取到的token和timestamp写入配置文件

	//c.log.Logger.Infof("本次请求结束,返回结果:[%s]", respStr)
	return basic2, nil
}

// handleResponseToken 处理http响应信息, errCode不为0则返回错误
func (c *Client) handleResponseToken(resp *http.Response) (root.Basic, error) {

	defer resp.Body.Close()

	var basic root.Basic

	var respBody CommonResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return basic, err
	}

	if respBody.ErrCode != 0 {
		var customError root.Error
		customError.Code = strconv.Itoa(respBody.ErrCode)
		customError.Message = respBody.ErrMsg
		return basic, &customError
	}

	respStr, err := json.Marshal(respBody)
	if err != nil {
		c.log.Logger.Errorf("本次请求出现错误,错误信息[%s]", err.Error())
		return basic, err
	}
	c.log.Logger.Infof("本次请求结束,返回结果:[%s]", respStr)

	basic.TimeStamp = respBody.TimeStamp
	basic.Token = respBody.AccessToken

	return basic, nil
}

// handleResponse 处理http响应信息, errCode不为0则返回错误
func handleResponse(resp *http.Response) (string, error) {

	defer resp.Body.Close()

	var respBody CommonResponse
	err := json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return "", err
	}

	if respBody.ErrCode != 0 {
		var customError root.Error
		customError.Code = strconv.Itoa(respBody.ErrCode)
		customError.Message = respBody.ErrMsg
		return "", &customError
	}

	respStr, err := json.Marshal(respBody)
	if err != nil {
		return "", err
	}

	return string(respStr), nil
}

// CatService 分类信息服务
func (c *Client) CatService() *CatService {
	return &c.catService
}

// BrandService 品牌信息服务
func (c *Client) BrandService() *BrandService {
	return &c.brandService
}

// GoodsService 商品信息服务
func (c *Client) GoodsService() *GoodsService {
	return &c.goodsService
}

// OrderService 订单信息服务
func (c *Client) OrderService() *OrderService {
	return &c.orderService
}

// MemberService 会员信息服务
func (c *Client) MemberService() *MemberService {
	return &c.memberService
}
