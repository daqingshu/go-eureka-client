package eureka

const (
	STATUS_UP             = "UP"
	STATUS_DOWN           = "DOWN"
	STATUS_STARTING       = "STARTING"
	STATUS_OUT_OF_SERVICE = "OUT_OF_SERVICE"
	STATUS_UNKNOWN        = "UNKNOWN"

	DC_NAME_TYPE_MY_OWN = "MyOwn"
	DC_NAME_TYPE_AMAZON = "Amazon"
)

type (
	InstanceVo struct {
		// Register application instance needed -- BEGIN
		Hostname                      string          `json:"hostName"`
		App                           string          `json:"app"`
		IpAddr                        string          `json:"ipAddr"`
		VipAddress                    string          `json:"vipAddress"`
		SecureVipAddress              string          `json:"secureVipAddress"`
		Status                        string          `json:"status"`
		Port                          *PositiveInt    `json:"port"`
		SecurePort                    *PositiveInt    `json:"securePort"`
		HomePageUrl                   string          `json:"homePageUrl"`
		StatusPageUrl                 string          `json:"statusPageUrl"`
		HealthCheckUrl                string          `json:"healthCheckUrl"`
		DataCenterInfo                *DataCenterInfo `json:"dataCenterInfo"`
		LeaseInfo                     *LeaseInfo      `json:"leaseInfo"`
		Metadata                      *MetaData       `xml:"metadata,omitempty" json:"metadata,omitempty"`
		IsCoordinatingDiscoveryServer bool            `xml:"isCoordinatingDiscoveryServer,omitempty" json:"isCoordinatingDiscoveryServer,omitempty"`
		// Register application instance needed -- END

		InstanceId           string `json:"instanceId,omitempty"`
		OverriddenStatus     string `json:"overriddenstatus,omitempty"`
		LastUpdatedTimestamp int    `json:"lastUpdatedTimestamp,omitempty"`
		LastDirtyTimestamp   int    `json:"lastUpdatedTimestamp,omitempty"`
		ActionType           string `json:"actionType,omitempty"`
		CountryId            int    `xml:"countryId,omitempty" json:"countryId,omitempty"`
	}

	PositiveInt struct {
		Value   int  `json:"$"`
		Enabled bool `json:"@enabled"` // true|false
	}

	DataCenterInfo struct {
		// MyOwn | Amazon
		Name string `json:"name"`
		// metadata is only required if name is Amazon
		Metadata *DataCenterMetadata `json:"metadata,omitempty"`
		Class    string              `json:"@class"`
	}

	DataCenterMetadata struct {
		AmiLaunchIndex   string `xml:"ami-launch-index,omitempty" json:"ami-launch-index,omitempty"`
		LocalHostname    string `xml:"local-hostname,omitempty" json:"local-hostname,omitempty"`
		AvailabilityZone string `xml:"availability-zone,omitempty" json:"availability-zone,omitempty"`
		InstanceId       string `xml:"instance-id,omitempty" json:"instance-id,omitempty"`
		PublicIpv4       string `xml:"public-ipv4,omitempty" json:"public-ipv4,omitempty"`
		PublicHostname   string `xml:"public-hostname,omitempty" json:"public-hostname,omitempty"`
		AmiManifestPath  string `xml:"ami-manifest-path,omitempty" json:"ami-manifest-path,omitempty"`
		LocalIpv4        string `xml:"local-ipv4,omitempty" json:"local-ipv4,omitempty"`
		Hostname         string `xml:"hostname,omitempty" json:"hostname,omitempty"`
		AmiId            string `xml:"ami-id,omitempty" json:"ami-id,omitempty"`
		InstanceType     string `xml:"instance-type,omitempty" json:"instance-type,omitempty"`
	}

	LeaseInfo struct {
		// (optional) if you want to change the length of lease - default if 90 seconds
		EvictionDurationInSecs int `json:"eviction_duration_in_secs,omitempty"`
		RenewalIntervalInSecs  int `xml:"renewalIntervalInSecs,omitempty" json:"renewalIntervalInSecs,omitempty"`
		DurationInSecs         int `xml:"durationInSecs,omitempty" json:"durationInSecs,omitempty"`
		RegistrationTimestamp  int `xml:"registrationTimestamp,omitempty" json:"registrationTimestamp,omitempty"`
		LastRenewalTimestamp   int `xml:"lastRenewalTimestamp,omitempty" json:"lastRenewalTimestamp,omitempty"`
		EvictionTimestamp      int `xml:"evictionTimestamp,omitempty" json:"evictionTimestamp,omitempty"`
		ServiceUpTimestamp     int `xml:"serviceUpTimestamp,omitempty" json:"serviceUpTimestamp,omitempty"`
	}

	Instance struct {
		Instance *InstanceVo `xml:"instance" json:"instance"`
	}

	// application
	ApplicationVo struct {
		Name      string       `json:"name"`
		Instances []InstanceVo `json:"instance"`
	}

	ApplicationsVo struct {
		VersionDelta string          `json:"version__delta"`
		AppsHashCode string          `json:"apps_hash__code"`
		Application  []ApplicationVo `json:"application"`
	}
)

func DefaultInstanceVo() *InstanceVo {
	ip := getLocalIp()
	//hostname, err := os.Hostname()
	//if err != nil {
	//    log.Errorf("Failed to get hostname, err=%s, user ip as hostname, ip=%s", err.Error(), ip)
	//    hostname = ip
	//}
	return &InstanceVo{
		//Hostname:         hostname,
		Hostname:         ip,
		App:              "",
		IpAddr:           ip,
		VipAddress:       ip,
		SecureVipAddress: ip,
		Status:           STATUS_STARTING,
		Port:             &PositiveInt{Value: 8080, Enabled: true},
		SecurePort:       &PositiveInt{Value: 443, Enabled: false},
		HomePageUrl:      "",
		StatusPageUrl:    "",
		HealthCheckUrl:   "",
		DataCenterInfo: &DataCenterInfo{
			Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
			Name:  DC_NAME_TYPE_MY_OWN,
		},
		LeaseInfo: &LeaseInfo{
			EvictionDurationInSecs: 30,
		},
	}
}
