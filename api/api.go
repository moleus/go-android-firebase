package api

import (
	"bytes"
	"encoding/json"
	. "github.com/BRUHItsABunny/go-android-firebase/constants"
	"net/http"
)

type NotifyInstallationRequestBody struct {
	FID         string `json:"fid"`
	AppID       string `json:"appId"`
	AuthVersion string `json:"authVersion"`
	SDKVersion  string `json:"sdkVersion"`
}

type FireBaseInstallationResponse struct {
	Name         string            `json:"name"`
	FID          string            `json:"fid"`
	RefreshToken string            `json:"refreshToken"`
	AuthToken    FireBaseAuthToken `json:"authToken"`
}

type FireBaseAuthToken struct {
	Token      string `json:"token"`
	Expiration string `json:"expiresin"`
}

type NotifyInstallationResponse struct {
	FID         string `json:"fid"`
	AppID       string `json:"appId"`
	AuthVersion string `json:"authVersion"`
	SDKVersion  string `json:"sdkVersion"`
}

type HeaderFiller interface {
	Fill(*http.Request) *http.Request
}

type DefaultHeadersFiller struct {
	Headers map[string]string
}

func (filler *DefaultHeadersFiller) Fill(req *http.Request) *http.Request {

	for key, val := range filler.Headers {
		req.Header[key] = []string{val}
	}

	return req
}

var DefaultHeaders = DefaultHeadersFiller{
	Headers: map[string]string{
		HeaderKeyContentType:  HeaderValueMIMEJSON,
		HeaderKeyAccept:       HeaderValueMIMEJSON,
		HeaderKeyCacheControl: "no-cache",
	},
}

func NotifyInstallationRequest(data *NotifyInstallationRequestBody, filler HeaderFiller, ProjectID, AndroidPackage, AndroidCertificate, GoogAPIKey, FireBaseClient, FireBaseLogType, UserAgent string) *http.Request {
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", Protocol+Host+EndpointProjects+ProjectID+SubEndpointInstallations, bytes.NewBuffer(body))

	req.Header[HeaderKeyAndroidCert] = []string{AndroidCertificate}
	req.Header[HeaderKeyAndroidPackage] = []string{AndroidPackage}
	req.Header[HeaderKeyFireBaseClient] = []string{FireBaseClient}
	req.Header[HeaderKeyGoogAPIKey] = []string{GoogAPIKey}
	req.Header[HeaderKeyFireBaseLogType] = []string{FireBaseLogType}
	req.Header[HeaderKeyUserAgent] = []string{UserAgent}

	req = filler.Fill(req)
	return req
}