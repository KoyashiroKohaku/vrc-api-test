package vrcapi

import (
	"time"
)

type Config struct {
	MessageOfTheDay string `json:"messageOfTheDay"`
	Events          struct {
		DistanceClose             int `json:"distanceClose"`
		DistanceFactor            int `json:"distanceFactor"`
		DistanceFar               int `json:"distanceFar"`
		GroupDistance             int `json:"groupDistance"`
		MaximumBunchSize          int `json:"maximumBunchSize"`
		NotVisibleFactor          int `json:"notVisibleFactor"`
		PlayerOrderBucketSize     int `json:"playerOrderBucketSize"`
		PlayerOrderFactor         int `json:"playerOrderFactor"`
		SlowUpdateFactorThreshold int `json:"slowUpdateFactorThreshold"`
		ViewSegmentLength         int `json:"viewSegmentLength"`
	} `json:"events"`
	DisCountdown                                  time.Time `json:"dis-countdown"`
	HomepageRedirectTarget                        string    `json:"homepageRedirectTarget"`
	YoutubedlHash                                 string    `json:"youtubedl-hash"`
	YoutubedlVersion                              string    `json:"youtubedl-version"`
	TimeOutWorldID                                string    `json:"timeOutWorldId"`
	GearDemoRoomID                                string    `json:"gearDemoRoomId"`
	ReleaseServerVersionStandalone                string    `json:"releaseServerVersionStandalone"`
	DownloadLinkWindows                           string    `json:"downloadLinkWindows"`
	ReleaseAppVersionStandalone                   string    `json:"releaseAppVersionStandalone"`
	DevAppVersionStandalone                       string    `json:"devAppVersionStandalone"`
	DevServerVersionStandalone                    string    `json:"devServerVersionStandalone"`
	DevDownloadLinkWindows                        string    `json:"devDownloadLinkWindows"`
	CurrentTOSVersion                             int       `json:"currentTOSVersion"`
	ReleaseSdkURL                                 string    `json:"releaseSdkUrl"`
	ReleaseSdkVersion                             string    `json:"releaseSdkVersion"`
	DevSdkURL                                     string    `json:"devSdkUrl"`
	DevSdkVersion                                 string    `json:"devSdkVersion"`
	WhiteListedAssetUrls                          []string  `json:"whiteListedAssetUrls"`
	ClientAPIKey                                  string    `json:"clientApiKey"`
	ViveWindowsURL                                string    `json:"viveWindowsUrl"`
	SdkUnityVersion                               string    `json:"sdkUnityVersion"`
	HubWorldID                                    string    `json:"hubWorldId"`
	HomeWorldID                                   string    `json:"homeWorldId"`
	TutorialWorldID                               string    `json:"tutorialWorldId"`
	DisableEventStream                            bool      `json:"disableEventStream"`
	DisableAvatarGating                           bool      `json:"disableAvatarGating"`
	DisableFeedbackGating                         bool      `json:"disableFeedbackGating"`
	DisableRegistration                           bool      `json:"disableRegistration"`
	DisableUpgradeAccount                         bool      `json:"disableUpgradeAccount"`
	DisableCommunityLabs                          bool      `json:"disableCommunityLabs"`
	DisableCommunityLabsPromotion                 bool      `json:"disableCommunityLabsPromotion"`
	DisableTwoFactorAuth                          bool      `json:"disableTwoFactorAuth"`
	DisableSteamNetworking                        bool      `json:"disableSteamNetworking"`
	DisableHello                                  bool      `json:"disableHello"`
	DisableUdon                                   bool      `json:"disableUdon"`
	Plugin                                        string    `json:"plugin"`
	SdkNotAllowedToPublishMessage                 string    `json:"sdkNotAllowedToPublishMessage"`
	SdkDeveloperFaqURL                            string    `json:"sdkDeveloperFaqUrl"`
	SdkDiscordURL                                 string    `json:"sdkDiscordUrl"`
	NotAllowedToSelectAvatarInPrivateWorldMessage string    `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	UserVerificationTimeout                       int       `json:"userVerificationTimeout"`
	UserUpdatePeriod                              int       `json:"userUpdatePeriod"`
	UserVerificationDelay                         int       `json:"userVerificationDelay"`
	UserVerificationRetry                         int       `json:"userVerificationRetry"`
	WorldUpdatePeriod                             int       `json:"worldUpdatePeriod"`
	ModerationQueryPeriod                         int       `json:"moderationQueryPeriod"`
	ClientDisconnectTimeout                       int       `json:"clientDisconnectTimeout"`
	DefaultAvatar                                 string    `json:"defaultAvatar"`
	DynamicWorldRows                              []struct {
		Name          string `json:"name"`
		SortHeading   string `json:"sortHeading"`
		SortOwnership string `json:"sortOwnership"`
		SortOrder     string `json:"sortOrder"`
		Platform      string `json:"platform"`
		Index         int    `json:"index"`
		Tag           string `json:"tag,omitempty"`
	} `json:"dynamicWorldRows"`
	DisableAvatarCopying bool `json:"disableAvatarCopying"`
	Announcements        []struct {
		Name string `json:"name"`
		Text string `json:"text"`
	} `json:"announcements"`
	UseReliableUDPForVoice   bool `json:"useReliableUdpForVoice"`
	UpdateRateMsMaximum      int  `json:"updateRateMsMaximum"`
	UpdateRateMsMinimum      int  `json:"updateRateMsMinimum"`
	UpdateRateMsNormal       int  `json:"updateRateMsNormal"`
	ClientBPSCeiling         int  `json:"clientBPSCeiling"`
	ClientReservedPlayerBPS  int  `json:"clientReservedPlayerBPS"`
	ClientSentCountAllowance int  `json:"clientSentCountAllowance"`
	UploadAnalysisPercent    int  `json:"uploadAnalysisPercent"`
	DownloadUrls             struct {
		Sdk2        string `json:"sdk2"`
		Sdk3Worlds  string `json:"sdk3-worlds"`
		Sdk3Avatars string `json:"sdk3-avatars"`
	} `json:"downloadUrls"`
	Address         string `json:"address"`
	ContactEmail    string `json:"contactEmail"`
	SupportEmail    string `json:"supportEmail"`
	JobsEmail       string `json:"jobsEmail"`
	CopyrightEmail  string `json:"copyrightEmail"`
	ModerationEmail string `json:"moderationEmail"`
	DisableEmail    bool   `json:"disableEmail"`
	AppName         string `json:"appName"`
	ServerName      string `json:"serverName"`
	DeploymentGroup string `json:"deploymentGroup"`
	BuildVersionTag string `json:"buildVersionTag"`
	APIKey          string `json:"apiKey"`
}
