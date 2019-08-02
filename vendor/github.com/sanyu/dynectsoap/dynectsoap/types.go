package dynectsoap

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Messages struct {
	// XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ messages"`

	Source string `xml:"source,omitempty"`

	Lvl string `xml:"lvl,omitempty"`

	Err_cd string `xml:"err_cd,omitempty"`

	Info string `xml:"info,omitempty"`
}

type WeightData struct {
	A_weight []string `xml:"a_weight,omitempty"`

	Aaaa_weight []string `xml:"aaaa_weight,omitempty"`

	Cname_weight []string `xml:"cname_weight,omitempty"`
}

type ServeCountData struct {
	A_serve_count string `xml:"a_serve_count,omitempty"`

	Aaaa_serve_count string `xml:"aaaa_serve_count,omitempty"`
}

type TTLData struct {
	A_ttl int32 `xml:"a_ttl,omitempty"`

	Aaaa_ttl int32 `xml:"aaaa_ttl,omitempty"`

	Cert_ttl int32 `xml:"cert_ttl,omitempty"`

	Cname_ttl int32 `xml:"cname_ttl,omitempty"`

	Mx_ttl int32 `xml:"mx_ttl,omitempty"`

	Txt_ttl int32 `xml:"txt_ttl,omitempty"`

	Spf_ttl int32 `xml:"spf_ttl,omitempty"`

	Ptr_ttl int32 `xml:"ptr_ttl,omitempty"`

	Loc_ttl int32 `xml:"loc_ttl,omitempty"`

	Srv_ttl int32 `xml:"srv_ttl,omitempty"`

	Rp_ttl int32 `xml:"rp_ttl,omitempty"`

	Key_ttl int32 `xml:"key_ttl,omitempty"`

	Dnskey_ttl int32 `xml:"dnskey_ttl,omitempty"`

	Sshfp_ttl int32 `xml:"sshfp_ttl,omitempty"`

	Dhcid_ttl int32 `xml:"dhcid_ttl,omitempty"`

	Nsap_ttl int32 `xml:"nsap_ttl,omitempty"`

	Px_ttl int32 `xml:"px_ttl,omitempty"`
}

type LabelData struct {
	A_label []string `xml:"a_label,omitempty"`

	Aaaa_label []string `xml:"aaaa_label,omitempty"`

	Cert_label []string `xml:"cert_label,omitempty"`

	Cname_label []string `xml:"cname_label,omitempty"`

	Mx_label []string `xml:"mx_label,omitempty"`

	Txt_label []string `xml:"txt_label,omitempty"`

	Spf_label []string `xml:"spf_label,omitempty"`

	Ptr_label []string `xml:"ptr_label,omitempty"`

	Loc_label []string `xml:"loc_label,omitempty"`

	Srv_label []string `xml:"srv_label,omitempty"`

	Rp_label []string `xml:"rp_label,omitempty"`

	Key_label []string `xml:"key_label,omitempty"`

	Dnskey_label []string `xml:"dnskey_label,omitempty"`

	Sshfp_label []string `xml:"sshfp_label,omitempty"`

	Dhcid_label []string `xml:"dhcid_label,omitempty"`

	Nsap_label []string `xml:"nsap_label,omitempty"`

	Px_label []string `xml:"px_label,omitempty"`
}

type GeoRegionGroup struct {
	Name string `xml:"name,omitempty"`

	Rdata *ANYRData `xml:"rdata,omitempty"`

	Countries []string `xml:"countries,omitempty"`

	Serve_count *ServeCountData `xml:"serve_count,omitempty"`

	Weight *WeightData `xml:"weight,omitempty"`

	Ttl *TTLData `xml:"ttl,omitempty"`

	Label *LabelData `xml:"label,omitempty"`
}

type GeoRegionGroupData struct {
	Name string `xml:"name,omitempty"`

	Rdata *ANYRData `xml:"rdata,omitempty"`

	Countries []string `xml:"countries,omitempty"`

	Serve_count *ServeCountData `xml:"serve_count,omitempty"`

	Weight *WeightData `xml:"weight,omitempty"`

	Ttl *TTLData `xml:"ttl,omitempty"`

	Label *LabelData `xml:"label,omitempty"`

	Service_name string `xml:"service_name,omitempty"`
}

type GeoNode struct {
	Fqdn string `xml:"fqdn,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type Geo struct {
	Name string `xml:"name,omitempty"`

	Groups []*GeoRegionGroup `xml:"groups,omitempty"`

	Nodes []*GeoNode `xml:"nodes,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Active string `xml:"active,omitempty"`
}

type DSFRData struct {
	Type_ string `xml:"type,omitempty"`

	Data *GenericRData `xml:"data,omitempty"`

	Ttl string `xml:"ttl,omitempty"`
}

type DSFRecord struct {
	Master_line string `xml:"master_line,omitempty"`

	// Rdata to update the svc record with
	Rdata *ANYOneRData `xml:"rdata,omitempty"`

	Label string `xml:"label,omitempty"`

	Weight string `xml:"weight,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Endpoints []string `xml:"endpoints,omitempty"`

	Endpoint_up_count string `xml:"endpoint_up_count,omitempty"`

	Eligible string `xml:"eligible,omitempty"`
}

type DSFRecordData struct {
	Dsf_record_id string `xml:"dsf_record_id,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	Label string `xml:"label,omitempty"`

	Master_line string `xml:"master_line,omitempty"`

	Weight int32 `xml:"weight,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Endpoints []string `xml:"endpoints,omitempty"`

	Endpoint_up_count int32 `xml:"endpoint_up_count,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Rdata_class string `xml:"rdata_class,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Rdata []*DSFRData `xml:"rdata,omitempty"`

	Status string `xml:"status,omitempty"`

	Response_time int32 `xml:"response_time,omitempty"`

	Torpidity int32 `xml:"torpidity,omitempty"`

	Last_monitored int32 `xml:"last_monitored,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`
}

type DSFRecordSet struct {
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	Failover string `xml:"failover,omitempty"`

	Rdata_class string `xml:"rdata_class,omitempty"`

	Label string `xml:"label,omitempty"`

	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Serve_count string `xml:"serve_count,omitempty"`

	Fail_count string `xml:"fail_count,omitempty"`

	Trouble_count string `xml:"trouble_count,omitempty"`

	Torpidity_max string `xml:"torpidity_max,omitempty"`

	Index string `xml:"index,omitempty"`

	Records []*DSFRecord `xml:"records,omitempty"`
}

type DSFRecordSetData struct {
	Service_id string `xml:"service_id,omitempty"`

	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	Label string `xml:"label,omitempty"`

	Rdata_class string `xml:"rdata_class,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Serve_count string `xml:"serve_count,omitempty"`

	Fail_count string `xml:"fail_count,omitempty"`

	Trouble_count string `xml:"trouble_count,omitempty"`

	Torpidity_max string `xml:"torpidity_max,omitempty"`

	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`

	Ttl_derived string `xml:"ttl_derived,omitempty"`

	Records []*DSFRecordData `xml:"records,omitempty"`

	Status string `xml:"status,omitempty"`

	Last_monitored string `xml:"last_monitored,omitempty"`
}

type DSFRecordSetFailoverChain struct {
	Label string `xml:"label,omitempty"`

	Core string `xml:"core,omitempty"`

	Record_sets []*DSFRecordSet `xml:"record_sets,omitempty"`
}

type DSFRecordSetFailoverChainData struct {
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	Label string `xml:"label,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Core string `xml:"core,omitempty"`

	Status string `xml:"status,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`

	Record_sets []*DSFRecordSetData `xml:"record_sets,omitempty"`
}

type DSFCriteria struct {
	Geoip *DSFGeoIPCriteria `xml:"geoip,omitempty"`
}

type DSFGeoIPCriteria struct {
	Country []string `xml:"country,omitempty"`

	Region []string `xml:"region,omitempty"`

	Province []string `xml:"province,omitempty"`
}

type DSFRulesetData struct {
	Service_id string `xml:"service_id,omitempty"`

	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	Label string `xml:"label,omitempty"`

	Criteria_type string `xml:"criteria_type,omitempty"`

	Criteria *DSFCriteria `xml:"criteria,omitempty"`

	Ordering string `xml:"ordering,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Response_pools []*DSFResponsePoolData `xml:"response_pools,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`
}

type DSFRuleset struct {
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	Label string `xml:"label,omitempty"`

	// The type of criteria contained within this Pool
	Criteria_type string `xml:"criteria_type,omitempty"`

	// Required based on criteria_type. Filtered in API/BLL
	Criteria *DSFCriteria `xml:"criteria,omitempty"`

	// where in the list of rulesets this should be. Defaults to last.
	Ordering string `xml:"ordering,omitempty"`

	Response_pools []*DSFResponsePool `xml:"response_pools,omitempty"`
}

type DSFResponsePoolData struct {
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Label string `xml:"label,omitempty"`

	// Notify Object
	Notifier string `xml:"notifier,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Core_set_count string `xml:"core_set_count,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`

	Rs_chains []*DSFRecordSetFailoverChainData `xml:"rs_chains,omitempty"`

	Rulesets []*DSFRulesetData `xml:"rulesets,omitempty"`

	Status string `xml:"status,omitempty"`

	Last_monitored string `xml:"last_monitored,omitempty"`
}

type DSFResponsePool struct {
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	Failover string `xml:"failover,omitempty"`

	Label string `xml:"label,omitempty"`

	Notifier string `xml:"notifier,omitempty"`

	Core_set_count string `xml:"core_set_count,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Rs_chains []*DSFRecordSetFailoverChain `xml:"rs_chains,omitempty"`

	Ruleset string `xml:"ruleset,omitempty"`

	Index string `xml:"index,omitempty"`
}

type DSFData struct {
	Service_id string `xml:"service_id,omitempty"`

	Active string `xml:"active,omitempty"`

	Label string `xml:"label,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Nodes []*DSFNode `xml:"nodes,omitempty"`

	Notifiers []*NotifierLinkData `xml:"notifiers,omitempty"`

	Rulesets []*DSFRulesetData `xml:"rulesets,omitempty"`

	Pending_change string `xml:"pending_change,omitempty"`
}

type DSFNode struct {
	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type DSFMonitorData struct {

	// ID
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	// Label for the DSF Monitor
	Label string `xml:"label,omitempty"`

	// Indicates whether or not the DSF Monitor is active
	Active string `xml:"active,omitempty"`

	// Num of responses to determine status
	Response_count string `xml:"response_count,omitempty"`

	// Interval, in seconds, between probes
	Probe_interval string `xml:"probe_interval,omitempty"`

	// number of attempted retries on failure before giving up
	Retries string `xml:"retries,omitempty"`

	// name of the protocol to monitor
	Protocol string `xml:"protocol,omitempty"`

	// options pertaining the monitor
	Options *DSFMonitorOptions `xml:"options,omitempty"`

	// IDs of attached notifiers
	Notifier []string `xml:"notifier,omitempty"`

	// Endpoints to monitor
	Endpoints []*DSFMonitorEndpoint `xml:"endpoints,omitempty"`

	// how are agents chosen?
	Agent_scheme string `xml:"agent_scheme,omitempty"`

	// IDs of attached services
	Services []string `xml:"services,omitempty"`
}

type DSFMonitorEndpoint struct {

	// Indicates whether or not the end point is active
	Active string `xml:"active,omitempty"`

	// label of the endpoint
	Label string `xml:"label,omitempty"`

	// address for the endpoint
	Address string `xml:"address,omitempty"`

	// ordered list of preferred sites for monitoring
	Site_prefs []string `xml:"site_prefs,omitempty"`
}

type DSFMonitorOptions struct {

	// time, in seconds, before the check timeout
	Timeout string `xml:"timeout,omitempty"`

	// an alternate port to connect to
	Port string `xml:"port,omitempty"`

	// a specific path to request
	Path string `xml:"path,omitempty"`

	// a value to pass into the `HOST:` header
	Host string `xml:"host,omitempty"`

	// additional header fields
	Header string `xml:"header,omitempty"`

	// a string to search for in the response
	Expected string `xml:"expected,omitempty"`

	Host_override []*DSFMonitorHostOverride `xml:"host_override,omitempty"`
}

type DSFMonitorHostOverride struct {

	// address of an endpoint
	Address string `xml:"address,omitempty"`

	// host to use when checking that endpoint
	Host string `xml:"host,omitempty"`
}

type DSFMonitorSitesData struct {

	// information for a site performing DSF monitoring
	DSFMonitorSites []*DSFMonitorSite `xml:"DSFMonitorSites,omitempty"`
}

type DSFMonitorSite struct {

	// site code for the monitoring machine
	Code string `xml:"code,omitempty"`

	// description of the machine's location
	Description string `xml:"description,omitempty"`

	// CIDR of the monitoring machine(s)
	Address string `xml:"address,omitempty"`
}

type Notifier struct {

	// Label for the Notifier object
	Label string `xml:"label,omitempty"`

	// List of Recipients attached to the Notifier
	Recipients []*Recipient `xml:"recipients,omitempty"`

	// List of services attached to the Notifier
	Services []*Service `xml:"services,omitempty"`
}

type NotifierLink struct {

	// Public_id of the Notifier to link to
	Notifier_id string `xml:"notifier_id,omitempty"`

	// filters on when services should fire the notifier
	Filters []string `xml:"filters,omitempty"`
}

type NotifierDataAlt struct {

	// Public ID of the Notifier object
	Notifier_id string `xml:"notifier_id,omitempty"`

	// Label for the Notifier object
	Label string `xml:"label,omitempty"`

	// Indicates whether or not the Notifier is active
	Active string `xml:"active,omitempty"`

	// List of Recepients attached to the Notifier
	Recipients []string `xml:"recipients,omitempty"`

	// List of services attached to the Notifier
	Services []*Service `xml:"services,omitempty"`
}

type NotifierData struct {

	// Public ID of the Notifier object
	Notifier_id string `xml:"notifier_id,omitempty"`

	// Label for the Notifier object
	Label string `xml:"label,omitempty"`

	// Indicates whether or not the Notifier is active
	Active string `xml:"active,omitempty"`

	// List of Recepients attached to the Notifier
	Recipients []*Recipient `xml:"recipients,omitempty"`

	// List of services attached to the Notifier
	Services []*Service `xml:"services,omitempty"`
}

type NotifierLinkData struct {

	// Public ID of the Notifier object
	Link_id string `xml:"link_id,omitempty"`

	// Indicates whether or not the Notifier link is active
	Active string `xml:"active,omitempty"`

	// filters on when services should fire the notifier
	Filters []string `xml:"filters,omitempty"`

	Notifier *NotifierSummaryData `xml:"notifier,omitempty"`
}

type NotifierSummaryData struct {

	// Public ID of the Notifier object
	Notifier_id string `xml:"notifier_id,omitempty"`

	Label string `xml:"label,omitempty"`

	// List of Recepients attached to the Notifier
	Recipients []string `xml:"recipients,omitempty"`

	// Indicates whether or not the Notifier is active
	Active string `xml:"active,omitempty"`
}

type Recipient struct {

	// email or syslog
	Format string `xml:"format,omitempty"`

	// For email, valid address or contact name. Syslog - address or hostname
	Recipient string `xml:"recipient,omitempty"`

	// hash of options
	Details *RecipientDetail `xml:"details,omitempty"`

	// List of string. For email - detail and summary are valid
	Features []string `xml:"features,omitempty"`
}

type RecipientDetail struct {

	// syslog port
	Port string `xml:"port,omitempty"`

	// syslog ident
	Ident string `xml:"ident,omitempty"`

	// syslog facility
	Facility string `xml:"facility,omitempty"`

	// syslog tls
	Tls string `xml:"tls,omitempty"`
}

type Service struct {

	// Valid entries - DSF or Monitor
	Service_class string `xml:"service_class,omitempty"`

	// public_id of the specified service
	Service_id string `xml:"service_id,omitempty"`

	// filters on when services should fire the notifier
	Filters []string `xml:"filters,omitempty"`

	// Indicates whether or not the link to the service is active
	Active string `xml:"active,omitempty"`
}

type ConfigLimitData struct {
	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type PermissionZone struct {
	Zone_name string `xml:"zone_name,omitempty"`

	Recurse string `xml:"recurse,omitempty"`

	// This field is returned in responses from the API, it should not be included in requests.
	Reason []string `xml:"reason,omitempty"`
}

type PermissionResponse struct {
	Admin_override string `xml:"admin_override,omitempty"`

	Allowed []*PermissionData `xml:"allowed,omitempty"`

	Forbidden []*PermissionData `xml:"forbidden,omitempty"`
}

type PermissionData struct {
	Name string `xml:"name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`

	Reason []string `xml:"reason,omitempty"`
}

type PermissionGroupData struct {
	Group_name string `xml:"group_name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Description string `xml:"description,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	User_name []string `xml:"user_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type TSIGKeyData struct {
	Name string `xml:"name,omitempty"`

	Algorithm string `xml:"algorithm,omitempty"`

	Secret string `xml:"secret,omitempty"`
}

type HostStatFlagsData struct {

	// empty for customer-wide default
	Zone_name string `xml:"zone_name,omitempty"`

	// Y or N
	Active string `xml:"active,omitempty"`
}

type ZoneData struct {

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// code indicating how serial numbers are constructed on publish
	Serial_style string `xml:"serial_style,omitempty"`

	// current serial number
	Serial int32 `xml:"serial,omitempty"`

	// Type of zone. Primary or Secondary
	Zone_type string `xml:"zone_type,omitempty"`
}

type SecondaryData struct {
	Zone string `xml:"zone,omitempty"`

	Masters []string `xml:"masters,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Contact_nickname string `xml:"contact_nickname,omitempty"`

	Active string `xml:"active,omitempty"`

	Task_id string `xml:"task_id,omitempty"`
}

type SessionLoginData struct {
	Token string `xml:"token,omitempty"`

	Version string `xml:"version,omitempty"`
}

type QueryStatsData struct {
	Csv string `xml:"csv,omitempty"`
}

type ARecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataA `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type AAAARecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataAAAA `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type ALIASRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataALIAS `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CAARecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCAA `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CDNSKEYRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDNSKEY `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CDSRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDS `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CERTRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCERT `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CNAMERecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCNAME `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CSYNCRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCSYNC `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DHCIDRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDHCID `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DNAMERecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNAME `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DNSKEYRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNSKEY `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DSRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDS `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type IPSECKEYRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataIPSECKEY `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type KEYRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKEY `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type KXRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKX `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type LOCRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataLOC `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type MXRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataMX `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type NAPTRRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNAPTR `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type NSAPRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNSAP `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type POLICYRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPOLICY `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type PTRRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPTR `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type PXRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPX `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type RPRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataRP `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type SPFRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSPF `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type SRVRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSRV `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type SSHFPRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSSHFP `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type TLSARecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTLSA `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type TXTRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTXT `xml:"rdata,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type SOARecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataSOA `xml:"rdata,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Serial_style string `xml:"serial_style,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type RDataSOAUpdate struct {
	Rname string `xml:"rname,omitempty"`

	Refresh string `xml:"refresh,omitempty"`

	Retry string `xml:"retry,omitempty"`

	Expire string `xml:"expire,omitempty"`

	Minimum string `xml:"minimum,omitempty"`
}

type NSRecordData struct {
	Ttl int32 `xml:"ttl,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNS `xml:"rdata,omitempty"`

	Service_class string `xml:"service_class,omitempty"`

	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type UserData struct {
	User_name string `xml:"user_name,omitempty"`

	User_id string `xml:"user_id,omitempty"`

	Password string `xml:"password,omitempty"`

	Group_name []string `xml:"group_name,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Status string `xml:"status,omitempty"`
}

type ContactData struct {
	Nickname string `xml:"nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type CustomerNSData struct {
	Primary string `xml:"primary,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type CustomerInterfaceData struct {
	Name string `xml:"name,omitempty"`

	Address []string `xml:"address,omitempty"`
}

type CustomerData struct {
	Id string `xml:"id,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Status string `xml:"status,omitempty"`

	Pool_id string `xml:"pool_id,omitempty"`

	Activated string `xml:"activated,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Level string `xml:"level,omitempty"`

	Owner_contact string `xml:"owner_contact,omitempty"`

	Billing_contact string `xml:"billing_contact,omitempty"`

	Primary_sales_agent string `xml:"primary_sales_agent,omitempty"`

	Salesforce_id string `xml:"salesforce_id,omitempty"`

	Default_ns []*CustomerNSData `xml:"default_ns,omitempty"`

	Interfaces []*CustomerInterfaceData `xml:"interfaces,omitempty"`

	Admin_user_id string `xml:"admin_user_id,omitempty"`

	Compartment_id string `xml:"compartment_id,omitempty"`

	Tenant string `xml:"tenant,omitempty"`
}

type CustomerAdminData struct {
	User_name string `xml:"user_name,omitempty"`

	Password string `xml:"password,omitempty"`

	Email string `xml:"email,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type CustomerPrefData struct {

	// name of the pref setting
	Name string `xml:"name,omitempty"`

	// value of the setting
	Value string `xml:"value,omitempty"`

	// Y/N: pref is not explicitly set, this is the default value
	Default_ string `xml:"default,omitempty"`
}

type CustomerIPACL struct {

	// comma or space-delimited list of netmasks, in CIDR form; no '/' assumes exact address
	Netmasks string `xml:"netmasks,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// 'web'/'api', default 'web'
	Scope string `xml:"scope,omitempty"`
}

type CustomerOracleMetadataData struct {

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// customer id
	Cust_id string `xml:"cust_id,omitempty"`

	// tenant id
	Tenant string `xml:"tenant,omitempty"`
}

type ZoneOracleMetadataData struct {

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// public id
	Public_id string `xml:"public_id,omitempty"`

	// zone id
	Zone_id string `xml:"zone_id,omitempty"`
}

type DDNSData struct {

	// an IP address, either v4 or v6
	Address string `xml:"address,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// last time this service was update by a Dynamic DNS client
	Last_updated string `xml:"last_updated,omitempty"`

	// count of excessive updates
	Abuse_count string `xml:"abuse_count,omitempty"`

	// 'Y', 'N'
	Active string `xml:"active,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type UpdateUserPasswordData struct {
	Password string `xml:"password,omitempty"`
}

type UpdateUser struct {
	User_name string `xml:"user_name,omitempty"`

	Password string `xml:"password,omitempty"`

	Group_name []string `xml:"group_name,omitempty"`

	Status string `xml:"status,omitempty"`
}

type DDNSHostData struct {
	Ddns *DDNSData `xml:"ddns,omitempty"`

	New_user *UpdateUser `xml:"new_user,omitempty"`
}

type FailoverData struct {

	// normally served address
	Address string `xml:"address,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip' or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// address or CNAME to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// 'ok', 'trouble', 'failover'
	Status string `xml:"status,omitempty"`

	// 'Y', 'N'
	Active string `xml:"active,omitempty"`

	// monitoring changes task
	Task_id string `xml:"task_id,omitempty"`

	// Recent monitoring results. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type MonitorData struct {

	// HTTP, HTTPS, PING, SMTP, TCP, NONE
	Protocol string `xml:"protocol,omitempty"`

	// time between checks, in minutes
	Interval int32 `xml:"interval,omitempty"`

	// times to retest on failure
	Retries int32 `xml:"retries,omitempty"`

	// if different from service default
	Port int32 `xml:"port,omitempty"`

	// e.g: http://host/path
	Path string `xml:"path,omitempty"`

	// if different from fqdn
	Host string `xml:"host,omitempty"`

	// check response for specific text
	Expected string `xml:"expected,omitempty"`

	// additional HTTP headers
	Header string `xml:"header,omitempty"`

	// test timeout
	Timeout int32 `xml:"timeout,omitempty"`
}

type MonitorLogData struct {

	// 'up', 'down', 'unk'
	Status string `xml:"status,omitempty"`

	// more details on error
	Message string `xml:"message,omitempty"`

	// unix timestamp of monitor
	Time int32 `xml:"time,omitempty"`

	// "airport" code
	Site_code string `xml:"site_code,omitempty"`
}

type LoadBalanceData struct {

	// pool of IP addresses to balance
	Pool []*LoadBalanceAddress `xml:"pool,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// 'ok', 'trouble', 'failover'
	Status string `xml:"status,omitempty"`

	// 'Y', 'N'
	Active string `xml:"active,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type LoadBalanceAddress struct {

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`
}

type LoadBalancePoolEntry struct {
	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`
}

type GSLBData struct {

	// per-region addresses and configuration
	Region []*GSLBRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// 'ok', 'trouble', 'failover'
	Status string `xml:"status,omitempty"`

	// 'Y', 'N'
	Active string `xml:"active,omitempty"`

	// monitoring changes task
	Task_id string `xml:"task_id,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GSLBRegion struct {

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// number of 'ok' addresses before region fails over
	Min_healthy int32 `xml:"min_healthy,omitempty"`

	// pool of IP addresses to balance
	Pool []*GSLBAddress `xml:"pool,omitempty"`
}

type GSLBAddress struct {

	// an IP address or FQDN to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`
}

type GSLBRegionData struct {

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// number of 'ok' addresses before region fails over
	Min_healthy int32 `xml:"min_healthy,omitempty"`

	// pool of IP addresses to balance
	Pool []*GSLBAddress `xml:"pool,omitempty"`

	// 'ok', 'trouble', 'failover'
	Status string `xml:"status,omitempty"`

	// monitoring changes task
	Task_id string `xml:"task_id,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GSLBRegionPoolEntry struct {
	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// an IP address or FQDN to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// monitoring changes task This field is returned in responses from the API, it should not be included in requests.
	Task_id string `xml:"task_id,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`
}

type RTTMData struct {

	// per-region addresses and configuration
	Region []*RTTMRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// for custom syslog messages
	Syslog_rttm_fmt string `xml:"syslog_rttm_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// details about performance monitoring
	Performance_monitor *MonitorData `xml:"performance_monitor,omitempty"`

	// 'ok', 'unk', 'trouble', 'failover'
	Status string `xml:"status,omitempty"`

	// 'Y', 'N'
	Active string `xml:"active,omitempty"`

	// monitoring changes task
	Task_id string `xml:"task_id,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type RTTMRegion struct {

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'Y' or 'N', if 'Y', region will filled in with global settings
	Autopopulate string `xml:"autopopulate,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// pool_count, number of addresses to be included in the serve pool
	Ep int32 `xml:"ep,omitempty"`

	// 'ip', 'global', 'region', default 'global'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' mode, address to serve on failover. For 'region' mode, region_code of the region to failover to.
	Failover_data string `xml:"failover_data,omitempty"`

	// failover_count, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Apmc int32 `xml:"apmc,omitempty"`

	// failover_count2, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Epmc int32 `xml:"epmc,omitempty"`

	// pool of IP addresses to balance
	Pool []*RTTMAddress `xml:"pool,omitempty"`

	// 'ok, 'unk', 'trouble', 'failover', This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// monitoring changes task This field is returned in responses from the API, it should not be included in requests.
	Task_id string `xml:"task_id,omitempty"`
}

type RTTMAddress struct {

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`
}

type RTTMLogData struct {

	// zone serial at which this status was made
	Serial string `xml:"serial,omitempty"`

	// timestamp in UTC at which this status was made
	Change_ts string `xml:"change_ts,omitempty"`

	// type of status change. 'health', 'perf', or 'user'
	Change_type string `xml:"change_type,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'up', 'trouble', 'failover', or 'reg_remove'
	Region_status string `xml:"region_status,omitempty"`

	// If change_type is 'user', the user that made the change
	User_name string `xml:"user_name,omitempty"`

	// number of addresses in each DNS response
	Serve_count string `xml:"serve_count,omitempty"`

	// 'A', 'AAAA', or 'CNAME'
	Rdata_type string `xml:"rdata_type,omitempty"`

	// List of rdata being served for this region
	Region_rdata []*RTTMLogRData `xml:"region_rdata,omitempty"`
}

type RTTMLogRData struct {

	// how often this is served relative to others in pool
	Weight string `xml:"weight,omitempty"`

	Rdata_a *RDataA `xml:"rdata_a,omitempty"`

	Rdata_aaaa *RDataAAAA `xml:"rdata_aaaa,omitempty"`

	Rdata_cname *RDataCNAME `xml:"rdata_cname,omitempty"`
}

type RTTMRegionData struct {

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'Y' or 'N', if 'Y', region will filled in with global settings
	Autopopulate string `xml:"autopopulate,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// pool_count, number of addresses to be included in the serve pool
	Ep int32 `xml:"ep,omitempty"`

	// 'ip', 'global', 'region', default 'global'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' mode, address to serve on failover. For 'region' mode, region_code of the region to failover to.
	Failover_data string `xml:"failover_data,omitempty"`

	// failover_count, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Apmc int32 `xml:"apmc,omitempty"`

	// failover_count2, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Epmc int32 `xml:"epmc,omitempty"`

	// pool of IP addresses to balance
	Pool []*RTTMAddress `xml:"pool,omitempty"`

	// monitoring changes task This field is returned in responses from the API, it should not be included in requests.
	Task_id string `xml:"task_id,omitempty"`

	// 'ok, 'unk', 'trouble', 'failover',
	Status string `xml:"status,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type RTTMRegionPoolEntry struct {
	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`

	// current monitoring status This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`

	// Recent monitoring results for this address. This field is returned in responses from the API, it should not be included in requests.
	Log []*MonitorLogData `xml:"log,omitempty"`

	// monitoring changes task This field is returned in responses from the API, it should not be included in requests.
	Task_id string `xml:"task_id,omitempty"`
}

type HTTPRedirectData struct {

	// URL requests are redirecto to
	Url string `xml:"url,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// should redirected URL include requested URL
	Keep_uri string `xml:"keep_uri,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type AdvRedirectRuleData struct {

	// Public ID of rule.
	Public_id string `xml:"public_id,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// host portion of URL to match
	Host_prefix string `xml:"host_prefix,omitempty"`

	// path portion of URL to match
	Path string `xml:"path,omitempty"`

	// replacement pattern
	Url_pattern string `xml:"url_pattern,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// Public ID of next AdvRedirect rule to be processed. (default to end of list)
	Next_public_id string `xml:"next_public_id,omitempty"`
}

type AdvRedirectData struct {

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// List of AdvRedirectRules
	Rules []*AdvRedirectRuleData `xml:"rules,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CDNManagerData struct {

	// per region hostnames and configuration
	Region []*CDNRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ok' or 'failover'
	Status string `xml:"status,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type CDNRegion struct {

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// pool of CNAMEs to balance
	Pool []*CDNHost `xml:"pool,omitempty"`

	// 'ok' or 'failover' This field is returned in responses from the API, it should not be included in requests.
	Status string `xml:"status,omitempty"`
}

type CDNHost struct {

	// hostname or address to redirect to
	Cname string `xml:"cname,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// is this record currently being served This field is returned in responses from the API, it should not be included in requests.
	Active string `xml:"active,omitempty"`

	// is this record for failover or normal operation
	Failover string `xml:"failover,omitempty"`
}

type RDataA struct {
	Address string `xml:"address,omitempty"`
}

type RDataAAAA struct {
	Address string `xml:"address,omitempty"`
}

type RDataALIAS struct {
	Alias string `xml:"alias,omitempty"`
}

type RDataCAA struct {
	Flags int32 `xml:"flags,omitempty"`

	Tag string `xml:"tag,omitempty"`

	Value string `xml:"value,omitempty"`
}

type RDataCDNSKEY struct {
	Flags int32 `xml:"flags,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Protocol int32 `xml:"protocol,omitempty"`

	Public_key string `xml:"public_key,omitempty"`
}

type RDataCDS struct {
	Keytag int32 `xml:"keytag,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Digtype int32 `xml:"digtype,omitempty"`

	Digest string `xml:"digest,omitempty"`
}

type RDataCERT struct {
	Format int32 `xml:"format,omitempty"`

	Tag int32 `xml:"tag,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Certificate string `xml:"certificate,omitempty"`
}

type RDataCNAME struct {
	Cname string `xml:"cname,omitempty"`
}

type RDataCSYNC struct {
	Soa_serial int32 `xml:"soa_serial,omitempty"`

	Flags string `xml:"flags,omitempty"`

	Types string `xml:"types,omitempty"`
}

type RDataDNSKEY struct {
	Flags int32 `xml:"flags,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Protocol int32 `xml:"protocol,omitempty"`

	Public_key string `xml:"public_key,omitempty"`
}

type RDataDHCID struct {
	Digest string `xml:"digest,omitempty"`
}

type RDataDNAME struct {
	Dname string `xml:"dname,omitempty"`
}

type RDataDS struct {
	Keytag int32 `xml:"keytag,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Digtype int32 `xml:"digtype,omitempty"`

	Digest string `xml:"digest,omitempty"`
}

type RDataIPSECKEY struct {
	Precedence int32 `xml:"precedence,omitempty"`

	Gatetype int32 `xml:"gatetype,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Gateway string `xml:"gateway,omitempty"`

	Public_key string `xml:"public_key,omitempty"`
}

type RDataKEY struct {
	Flags int32 `xml:"flags,omitempty"`

	Algorithm int32 `xml:"algorithm,omitempty"`

	Protocol int32 `xml:"protocol,omitempty"`

	Public_key string `xml:"public_key,omitempty"`
}

type RDataKX struct {
	Preference int32 `xml:"preference,omitempty"`

	Exchange string `xml:"exchange,omitempty"`
}

type RDataLOC struct {
	Latitude string `xml:"latitude,omitempty"`

	Longitude string `xml:"longitude,omitempty"`

	Altitude int32 `xml:"altitude,omitempty"`

	Horiz_pre int32 `xml:"horiz_pre,omitempty"`

	Vert_pre int32 `xml:"vert_pre,omitempty"`

	Size int32 `xml:"size,omitempty"`

	Version int32 `xml:"version,omitempty"`
}

type RDataMX struct {
	Preference int32 `xml:"preference,omitempty"`

	Exchange string `xml:"exchange,omitempty"`
}

type RDataNAPTR struct {
	Order int32 `xml:"order,omitempty"`

	Preference int32 `xml:"preference,omitempty"`

	Flags string `xml:"flags,omitempty"`

	Services string `xml:"services,omitempty"`

	Regexp string `xml:"regexp,omitempty"`

	Replacement string `xml:"replacement,omitempty"`
}

type RDataNS struct {
	Nsdname string `xml:"nsdname,omitempty"`
}

type RDataNSAP struct {
	Nsap string `xml:"nsap,omitempty"`
}

type RDataPOLICY struct {
	Name string `xml:"name,omitempty"`

	Rtype string `xml:"rtype,omitempty"`

	Api_url string `xml:"api_url,omitempty"`

	Gui_url string `xml:"gui_url,omitempty"`

	Policy string `xml:"policy,omitempty"`
}

type RDataPTR struct {
	Ptrdname string `xml:"ptrdname,omitempty"`
}

type RDataPX struct {
	Preference int32 `xml:"preference,omitempty"`

	Map822 string `xml:"map822,omitempty"`

	Mapx400 string `xml:"mapx400,omitempty"`
}

type RDataRP struct {
	Mbox string `xml:"mbox,omitempty"`

	Txtdname string `xml:"txtdname,omitempty"`
}

type RDataSPF struct {
	Txtdata string `xml:"txtdata,omitempty"`
}

type RDataSSHFP struct {
	Algorithm int32 `xml:"algorithm,omitempty"`

	Fptype int32 `xml:"fptype,omitempty"`

	Fingerprint string `xml:"fingerprint,omitempty"`
}

type RDataSRV struct {
	Priority int32 `xml:"priority,omitempty"`

	Weight int32 `xml:"weight,omitempty"`

	Port int32 `xml:"port,omitempty"`

	Target string `xml:"target,omitempty"`
}

type RDataTLSA struct {
	Cert_usage int32 `xml:"cert_usage,omitempty"`

	Selector int32 `xml:"selector,omitempty"`

	Match_type int32 `xml:"match_type,omitempty"`

	Certificate string `xml:"certificate,omitempty"`
}

type RDataTXT struct {
	Txtdata string `xml:"txtdata,omitempty"`
}

type RDataSOA struct {
	Mname string `xml:"mname,omitempty"`

	Rname string `xml:"rname,omitempty"`

	Serial int32 `xml:"serial,omitempty"`

	Refresh int32 `xml:"refresh,omitempty"`

	Retry int32 `xml:"retry,omitempty"`

	Expire int32 `xml:"expire,omitempty"`

	Minimum int32 `xml:"minimum,omitempty"`
}

type GenericRData struct {
	Rdata_a *RDataA `xml:"rdata_a,omitempty"`

	Rdata_aaaa *RDataAAAA `xml:"rdata_aaaa,omitempty"`

	Rdata_alias *RDataALIAS `xml:"rdata_alias,omitempty"`

	Rdata_caa *RDataCAA `xml:"rdata_caa,omitempty"`

	Rdata_cdnskey *RDataCDNSKEY `xml:"rdata_cdnskey,omitempty"`

	Rdata_cds *RDataCDS `xml:"rdata_cds,omitempty"`

	Rdata_cert *RDataCERT `xml:"rdata_cert,omitempty"`

	Rdata_cname *RDataCNAME `xml:"rdata_cname,omitempty"`

	Rdata_csync *RDataCSYNC `xml:"rdata_csync,omitempty"`

	Rdata_dhcid *RDataDHCID `xml:"rdata_dhcid,omitempty"`

	Rdata_dname *RDataDNAME `xml:"rdata_dname,omitempty"`

	Rdata_dnskey *RDataDNSKEY `xml:"rdata_dnskey,omitempty"`

	Rdata_ds *RDataDS `xml:"rdata_ds,omitempty"`

	Rdata_ipseckey *RDataIPSECKEY `xml:"rdata_ipseckey,omitempty"`

	Rdata_key *RDataKEY `xml:"rdata_key,omitempty"`

	Rdata_kx *RDataKX `xml:"rdata_kx,omitempty"`

	Rdata_loc *RDataLOC `xml:"rdata_loc,omitempty"`

	Rdata_mx *RDataMX `xml:"rdata_mx,omitempty"`

	Rdata_naptr *RDataNAPTR `xml:"rdata_naptr,omitempty"`

	Rdata_ns *RDataNS `xml:"rdata_ns,omitempty"`

	Rdata_nsap *RDataNSAP `xml:"rdata_nsap,omitempty"`

	Rdata_policy *RDataPOLICY `xml:"rdata_policy,omitempty"`

	Rdata_ptr *RDataPTR `xml:"rdata_ptr,omitempty"`

	Rdata_px *RDataPX `xml:"rdata_px,omitempty"`

	Rdata_rp *RDataRP `xml:"rdata_rp,omitempty"`

	Rdata_spf *RDataSPF `xml:"rdata_spf,omitempty"`

	Rdata_sshfp *RDataSSHFP `xml:"rdata_sshfp,omitempty"`

	Rdata_srv *RDataSRV `xml:"rdata_srv,omitempty"`

	Rdata_tlsa *RDataTLSA `xml:"rdata_tlsa,omitempty"`

	Rdata_txt *RDataTXT `xml:"rdata_txt,omitempty"`

	Rdata_soa *RDataSOA `xml:"rdata_soa,omitempty"`
}

type QNames struct {
	Names []string `xml:"names,omitempty"`
}

type ANYRecordData struct {
	A_records []*ARecordData `xml:"a_records,omitempty"`

	Aaaa_records []*AAAARecordData `xml:"aaaa_records,omitempty"`

	Alias_records []*ALIASRecordData `xml:"alias_records,omitempty"`

	Caa_records []*CAARecordData `xml:"caa_records,omitempty"`

	Cdnskey_records []*CDNSKEYRecordData `xml:"cdnskey_records,omitempty"`

	Cds_records []*CDSRecordData `xml:"cds_records,omitempty"`

	Cert_records []*CERTRecordData `xml:"cert_records,omitempty"`

	Cname_records []*CNAMERecordData `xml:"cname_records,omitempty"`

	Csync_records []*CSYNCRecordData `xml:"csync_records,omitempty"`

	Dhcid_records []*DHCIDRecordData `xml:"dhcid_records,omitempty"`

	Dname_records []*DNAMERecordData `xml:"dname_records,omitempty"`

	Dnskey_records []*DNSKEYRecordData `xml:"dnskey_records,omitempty"`

	Ds_records []*DSRecordData `xml:"ds_records,omitempty"`

	Ipseckey_records []*IPSECKEYRecordData `xml:"ipseckey_records,omitempty"`

	Key_records []*KEYRecordData `xml:"key_records,omitempty"`

	Kx_records []*KXRecordData `xml:"kx_records,omitempty"`

	Loc_records []*LOCRecordData `xml:"loc_records,omitempty"`

	Mx_records []*MXRecordData `xml:"mx_records,omitempty"`

	Naptr_records []*NAPTRRecordData `xml:"naptr_records,omitempty"`

	Nsap_records []*NSAPRecordData `xml:"nsap_records,omitempty"`

	Policy_records []*POLICYRecordData `xml:"policy_records,omitempty"`

	Ptr_records []*PTRRecordData `xml:"ptr_records,omitempty"`

	Px_records []*PXRecordData `xml:"px_records,omitempty"`

	Rp_records []*RPRecordData `xml:"rp_records,omitempty"`

	Spf_records []*SPFRecordData `xml:"spf_records,omitempty"`

	Srv_records []*SRVRecordData `xml:"srv_records,omitempty"`

	Sshfp_records []*SSHFPRecordData `xml:"sshfp_records,omitempty"`

	Tlsa_records []*TLSARecordData `xml:"tlsa_records,omitempty"`

	Txt_records []*TXTRecordData `xml:"txt_records,omitempty"`

	Soa_records []*SOARecordData `xml:"soa_records,omitempty"`

	Ns_records []*NSRecordData `xml:"ns_records,omitempty"`
}

type ANYRData struct {
	A_rdata []*RDataA `xml:"a_rdata,omitempty"`

	Aaaa_rdata []*RDataAAAA `xml:"aaaa_rdata,omitempty"`

	Alias_rdata []*RDataALIAS `xml:"alias_rdata,omitempty"`

	Caa_rdata []*RDataCAA `xml:"caa_rdata,omitempty"`

	Cdnskey_rdata []*RDataCDNSKEY `xml:"cdnskey_rdata,omitempty"`

	Cds_rdata []*RDataCDS `xml:"cds_rdata,omitempty"`

	Cert_rdata []*RDataCERT `xml:"cert_rdata,omitempty"`

	Cname_rdata []*RDataCNAME `xml:"cname_rdata,omitempty"`

	Csync_rdata []*RDataCSYNC `xml:"csync_rdata,omitempty"`

	Dhcid_rdata []*RDataDHCID `xml:"dhcid_rdata,omitempty"`

	Dname_rdata []*RDataDNAME `xml:"dname_rdata,omitempty"`

	Dnskey_rdata []*RDataDNSKEY `xml:"dnskey_rdata,omitempty"`

	Ds_rdata []*RDataDS `xml:"ds_rdata,omitempty"`

	Ipseckey_rdata []*RDataIPSECKEY `xml:"ipseckey_rdata,omitempty"`

	Key_rdata []*RDataKEY `xml:"key_rdata,omitempty"`

	Kx_rdata []*RDataKX `xml:"kx_rdata,omitempty"`

	Loc_rdata []*RDataLOC `xml:"loc_rdata,omitempty"`

	Mx_rdata []*RDataMX `xml:"mx_rdata,omitempty"`

	Naptr_rdata []*RDataNAPTR `xml:"naptr_rdata,omitempty"`

	Nsap_rdata []*RDataNSAP `xml:"nsap_rdata,omitempty"`

	Policy_rdata []*RDataPOLICY `xml:"policy_rdata,omitempty"`

	Ptr_rdata []*RDataPTR `xml:"ptr_rdata,omitempty"`

	Px_rdata []*RDataPX `xml:"px_rdata,omitempty"`

	Rp_rdata []*RDataRP `xml:"rp_rdata,omitempty"`

	Spf_rdata []*RDataSPF `xml:"spf_rdata,omitempty"`

	Srv_rdata []*RDataSRV `xml:"srv_rdata,omitempty"`

	Sshfp_rdata []*RDataSSHFP `xml:"sshfp_rdata,omitempty"`

	Tlsa_rdata []*RDataTLSA `xml:"tlsa_rdata,omitempty"`

	Txt_rdata []*RDataTXT `xml:"txt_rdata,omitempty"`

	Soa_rdata []*RDataSOA `xml:"soa_rdata,omitempty"`

	Ns_rdata []*RDataNS `xml:"ns_rdata,omitempty"`
}

type ANYOneRData struct {
	A_rdata *RDataA `xml:"a_rdata,omitempty"`

	Aaaa_rdata *RDataAAAA `xml:"aaaa_rdata,omitempty"`

	Alias_rdata *RDataALIAS `xml:"alias_rdata,omitempty"`

	Caa_rdata *RDataCAA `xml:"caa_rdata,omitempty"`

	Cdnskey_rdata *RDataCDNSKEY `xml:"cdnskey_rdata,omitempty"`

	Cds_rdata *RDataCDS `xml:"cds_rdata,omitempty"`

	Cert_rdata *RDataCERT `xml:"cert_rdata,omitempty"`

	Cname_rdata *RDataCNAME `xml:"cname_rdata,omitempty"`

	Csync_rdata *RDataCSYNC `xml:"csync_rdata,omitempty"`

	Dhcid_rdata *RDataDHCID `xml:"dhcid_rdata,omitempty"`

	Dname_rdata *RDataDNAME `xml:"dname_rdata,omitempty"`

	Dnskey_rdata *RDataDNSKEY `xml:"dnskey_rdata,omitempty"`

	Ds_rdata *RDataDS `xml:"ds_rdata,omitempty"`

	Ipseckey_rdata *RDataIPSECKEY `xml:"ipseckey_rdata,omitempty"`

	Key_rdata *RDataKEY `xml:"key_rdata,omitempty"`

	Kx_rdata *RDataKX `xml:"kx_rdata,omitempty"`

	Loc_rdata *RDataLOC `xml:"loc_rdata,omitempty"`

	Mx_rdata *RDataMX `xml:"mx_rdata,omitempty"`

	Naptr_rdata *RDataNAPTR `xml:"naptr_rdata,omitempty"`

	Nsap_rdata *RDataNSAP `xml:"nsap_rdata,omitempty"`

	Policy_rdata *RDataPOLICY `xml:"policy_rdata,omitempty"`

	Ptr_rdata *RDataPTR `xml:"ptr_rdata,omitempty"`

	Px_rdata *RDataPX `xml:"px_rdata,omitempty"`

	Rp_rdata *RDataRP `xml:"rp_rdata,omitempty"`

	Spf_rdata *RDataSPF `xml:"spf_rdata,omitempty"`

	Srv_rdata *RDataSRV `xml:"srv_rdata,omitempty"`

	Sshfp_rdata *RDataSSHFP `xml:"sshfp_rdata,omitempty"`

	Tlsa_rdata *RDataTLSA `xml:"tlsa_rdata,omitempty"`

	Txt_rdata *RDataTXT `xml:"txt_rdata,omitempty"`

	Soa_rdata *RDataSOA `xml:"soa_rdata,omitempty"`

	Ns_rdata *RDataNS `xml:"ns_rdata,omitempty"`
}

type ZoneChangeData struct {
	Id int64 `xml:"id,omitempty"`

	User_id int64 `xml:"user_id,omitempty"`

	Rdata_type string `xml:"rdata_type,omitempty"`

	Rdata *GenericRData `xml:"rdata,omitempty"`

	Serial int32 `xml:"serial,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type ZoneNoteData struct {
	Zone string `xml:"zone,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Serial int32 `xml:"serial,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Note string `xml:"note,omitempty"`

	Timestamp string `xml:"timestamp,omitempty"`
}

type ZoneTransferStatus struct {
	Zone string `xml:"zone,omitempty"`

	Master_ip string `xml:"master_ip,omitempty"`

	Status string `xml:"status,omitempty"`

	Message string `xml:"message,omitempty"`
}

type ZoneConfigOptionData struct {
	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`

	Target string `xml:"target,omitempty"`
}

type PublishZoneData struct {
	Zone string `xml:"zone,omitempty"`

	Serial int32 `xml:"serial,omitempty"`

	Serial_style string `xml:"serial_style,omitempty"`

	Zone_type string `xml:"zone_type,omitempty"`

	Task_id string `xml:"task_id,omitempty"`
}

type IPTrackData struct {

	// A, Dynamic_A, AAAA, Dynamic_AAAA
	Record_types []string `xml:"record_types,omitempty"`

	// List of hostnames to watch for records
	Hosts []string `xml:"hosts,omitempty"`

	// 'match', 'default', or a valid ttl
	Ttl string `xml:"ttl,omitempty"`

	// Mask that records should match
	Netmask string `xml:"netmask,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	Active string `xml:"active,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DNSSECData struct {
	Zone string `xml:"zone,omitempty"`

	// ,      contact that gets key notifications
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	Keys []*DNSSECKey `xml:"keys,omitempty"`

	Active string `xml:"active,omitempty"`
}

type DNSSECKey struct {

	// 'KSK' or 'ZSK'
	Type_ string `xml:"type,omitempty"`

	// 'RSA/SHA-1', 'RSA/SHA-256', 'RSA/SHA-512', 'DSA'
	Algorithm string `xml:"algorithm,omitempty"`

	Bits string `xml:"bits,omitempty"`

	Start_ts int32 `xml:"start_ts,omitempty"`

	Lifetime int32 `xml:"lifetime,omitempty"`

	Overlap int32 `xml:"overlap,omitempty"`

	Expire_ts int32 `xml:"expire_ts,omitempty"`

	// Only for updates: 'rollover', 'rollover_now', 'remove'
	Action string `xml:"action,omitempty"`

	Dnssec_key_id int64 `xml:"dnssec_key_id,omitempty"`

	// This field is returned in responses from the API, it should not be included in requests.
	Dnskey *RDataDNSKEY `xml:"dnskey,omitempty"`

	// preserved for compatibility This field is returned in responses from the API, it should not be included in requests.
	Ds *RDataDS `xml:"ds,omitempty"`

	All_ds []*RDataDS `xml:"all_ds,omitempty"`
}

type DNSSECTimelineEvent struct {
	Scheduled_ts int32 `xml:"scheduled_ts,omitempty"`

	Status string `xml:"status,omitempty"`

	Event string `xml:"event,omitempty"`

	Message string `xml:"message,omitempty"`

	Send_notify string `xml:"send_notify,omitempty"`

	User string `xml:"user,omitempty"`

	Dnssec_key_id int64 `xml:"dnssec_key_id,omitempty"`
}

type TaskArgData struct {
	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type TaskIDData struct {
	Task_id string `xml:"task_id,omitempty"`
}

type TaskData struct {
	Task_id string `xml:"task_id,omitempty"`

	// identifies the task operation
	Name string `xml:"name,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Zone_name string `xml:"zone_name,omitempty"`

	// ready running waiting complete failed canceled stalled
	Status string `xml:"status,omitempty"`

	// in a multi-step process, how far along
	Step_count int32 `xml:"step_count,omitempty"`

	// total number of steps in multi-step process
	Total_steps int32 `xml:"total_steps,omitempty"`

	// Y/N - does this task block further zone operations?
	Blocking string `xml:"blocking,omitempty"`

	Message string `xml:"message,omitempty"`

	Debug string `xml:"debug,omitempty"`

	Created_ts int64 `xml:"created_ts,omitempty"`

	Modified_ts int64 `xml:"modified_ts,omitempty"`

	// other arguments passed to the task
	Args []*TaskArgData `xml:"args,omitempty"`
}

type ExtNameserverData struct {

	// can be empty or 'default'
	Zone string `xml:"zone,omitempty"`

	// Y/N - does this block requests or add them
	Deny string `xml:"deny,omitempty"`

	Hosts []*ExtNSEntry `xml:"hosts,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Active string `xml:"active,omitempty"`
}

type ExtNSEntry struct {

	// address or CIDR
	Address string `xml:"address,omitempty"`

	// Y/N - do we send NOTIFYs to this host?
	Notifies string `xml:"notifies,omitempty"`

	// Y/N - do we accept [AI]XFRs from this host?
	Transfers string `xml:"transfers,omitempty"`
}
type JobResponseInterface interface {
	status() string
}
type GetAllRecordsResponseType struct {
	// XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAllRecordsResponse"`
	Job_id string         `xml:"job_id,omitempty"`
	Status string         `xml:"status,omitempty"`
	Msgs   []*Messages    `xml:"msgs,omitempty"`
	Data   *ANYRecordData `xml:"data,omitempty"`
}

func (e *GetAllRecordsResponseType) status() string {
	return e.Status
}

type ErrorResponseType struct {
	// XMLName xml.Name    `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ErrorResponse"`
	Job_id string      `xml:"job_id,omitempty"`
	Status string      `xml:"status,omitempty"`
	Msgs   []*Messages `xml:"msgs,omitempty"`
	// should be empty and can be ignored
	Data string `xml:"data,omitempty"`
}

func (e *ErrorResponseType) status() string {
	return e.Status
}

type GetJobRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetJobRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Job_id string `xml:"job_id,omitempty"`
}

type GetJobResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetJobResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	// Contains the response data.  Can be any type as GetJob is request-agnostic.
	Data interface{} `xml:"data,omitempty"`
}

type SessionLoginRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionLoginRequest"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Password string `xml:"password,omitempty"`
}

type SessionLoginResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionLoginResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	// identifies your session token (needed for all other Dynect API calls) and the API protocol version
	Data *SessionLoginData `xml:"data,omitempty"`
}

type SessionLogoutRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionLogoutRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type SessionLogoutResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionLogoutResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SessionIsAliveRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionIsAliveRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type SessionIsAliveResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionIsAliveResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SessionKeepAliveRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionKeepAliveRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type SessionKeepAliveResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SessionKeepAliveResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ScopeInRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ScopeInRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	// defaults to admin user
	User_name string `xml:"user_name,omitempty"`
}

type ScopeInResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ScopeInResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ScopeAsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ScopeAsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	// defaults to admin user
	User_name string `xml:"user_name,omitempty"`
}

type ScopeAsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ScopeAsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type UnscopeRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UnscopeRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type UnscopeResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UnscopeResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetQueryStatsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetQueryStatsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// The timestamp indicating the beginning of the period to report on
	Start_ts int32 `xml:"start_ts,omitempty"`

	// The timestamp indicating the end of the period to report on
	End_ts int32 `xml:"end_ts,omitempty"`

	// The fields to break down the data with
	Breakdown []string `xml:"breakdown,omitempty"`

	// A list of specific hostnames to report on.  A hostname beginning with '!' will cause that hostname to be excluded
	Hosts []string `xml:"hosts,omitempty"`

	// A list of specific nameservers to report on.  A nameserver beginning with '!' will cause that nameserver to be excluded
	Nameservers []string `xml:"nameservers,omitempty"`

	// A list of record types to report on.  A record type beginning with '!' will cause that record type to be excluded
	Rrecs []string `xml:"rrecs,omitempty"`

	// A list of zone names to report on.  A zone name beginning with '!' will cause that zone to be excluded.
	Zones []string `xml:"zones,omitempty"`
}

type GetQueryStatsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetQueryStatsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	// The CSV data containing the requested statistics
	Data *QueryStatsData `xml:"data,omitempty"`
}

type CreateGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the service we want to create
	Name string `xml:"name,omitempty"`

	// List of region groups that contain a list of countries and records to be served
	Groups []*GeoRegionGroup `xml:"groups,omitempty"`

	// List of zone name, node name pairs to link a node to the Geo Service
	Nodes []*GeoNode `xml:"nodes,omitempty"`

	// Default TTL for records
	Ttl int32 `xml:"ttl,omitempty"`
}

type CreateGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type UpdateGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Name string `xml:"name,omitempty"`

	// The new name to assign to the Geo Service
	New_name string `xml:"new_name,omitempty"`

	// List of region groups that contain a list of countries and records to be served
	Groups []*GeoRegionGroup `xml:"groups,omitempty"`

	// List of zone name, node name pairs to link a node to the Geo Service
	Nodes []*GeoNode `xml:"nodes,omitempty"`

	// Default TTL for records
	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type GetGeosRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeosRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name (partial) of the Geo Service to find
	Name string `xml:"name,omitempty"`

	// Name (partial) of the Geo Region Group to find
	Group_name string `xml:"group_name,omitempty"`
}

type GetGeosResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeosResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*Geo `xml:"data,omitempty"`
}

type GetOneGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name (partial) of the Geo Service to find
	Name string `xml:"name,omitempty"`
}

type GetOneGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type DeleteOneGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to delete
	Name string `xml:"name,omitempty"`
}

type DeleteOneGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to activate
	Name string `xml:"name,omitempty"`
}

type ActivateGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type DeactivateGeoRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateGeoRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to deactivate
	Name string `xml:"name,omitempty"`
}

type DeactivateGeoResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateGeoResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type CreateGeoRegionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoRegionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Service_name string `xml:"service_name,omitempty"`

	// ,				Name of the Region Group
	Name string `xml:"name,omitempty"`

	// Rdata to update the Region Group with
	Rdata *ANYRData `xml:"rdata,omitempty"`

	// List of countries to update the Region Group with
	Countries []string `xml:"countries,omitempty"`

	// Optional weights to accompany the rdata
	Weight *WeightData `xml:"weight,omitempty"`

	// Optional serve counts to accompany the rdata
	Serve_count *ServeCountData `xml:"serve_count,omitempty"`

	// Optional Default TTL values for each record
	Ttl *TTLData `xml:"ttl,omitempty"`

	// Optional labels for the rdata
	Label *LabelData `xml:"label,omitempty"`
}

type CreateGeoRegionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoRegionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type UpdateGeoRegionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGeoRegionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Service_name string `xml:"service_name,omitempty"`

	// ,				Name of the Region Group
	Name string `xml:"name,omitempty"`

	// Rdata to update the Region Group with
	Rdata *ANYRData `xml:"rdata,omitempty"`

	// List of countries to update the Region Group with
	Countries []string `xml:"countries,omitempty"`

	// Optional weights to accompany the rdata
	Weight *WeightData `xml:"weight,omitempty"`

	// Optional serve counts to accompany the rdata
	Serve_count *ServeCountData `xml:"serve_count,omitempty"`

	// Optional Default TTL values for each record
	Ttl *TTLData `xml:"ttl,omitempty"`

	// Optional labels for the rdata
	Label *LabelData `xml:"label,omitempty"`
}

type UpdateGeoRegionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGeoRegionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type DeleteOneGeoRegionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoRegionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Service_name string `xml:"service_name,omitempty"`

	// ,			Name of the Region Group
	Name string `xml:"name,omitempty"`
}

type DeleteOneGeoRegionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoRegionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetGeoRegionGroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeoRegionGroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Service_name string `xml:"service_name,omitempty"`
}

type GetGeoRegionGroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeoRegionGroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*GeoRegionGroupData `xml:"data,omitempty"`
}

type GetOneGeoRegionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGeoRegionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to update
	Service_name string `xml:"service_name,omitempty"`

	// ,			Name of the Region Group
	Name string `xml:"name,omitempty"`
}

type GetOneGeoRegionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGeoRegionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GeoRegionGroupData `xml:"data,omitempty"`
}

type CreateGeoNodeRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoNodeRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to add the nodes to
	Service_name string `xml:"service_name,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type CreateGeoNodeResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGeoNodeResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *Geo `xml:"data,omitempty"`
}

type DeleteOneGeoNodeRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoNodeRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service to delete the nodes from
	Service_name string `xml:"service_name,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type DeleteOneGeoNodeResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGeoNodeResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetGeoNodesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeoNodesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Name of the Geo Service
	Service_name string `xml:"service_name,omitempty"`
}

type GetGeoNodesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGeoNodesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []string `xml:"data,omitempty"`
}

type CreateDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// DSF Label
	Label string `xml:"label,omitempty"`

	// Default TTL to be used in this service
	Ttl string `xml:"ttl,omitempty"`

	// A list of fqdn,zone pairs to identify nodes to attach
	Nodes []*DSFNode `xml:"nodes,omitempty"`

	// A list of DSF Rulesets defined for the service
	Rulesets []*DSFRuleset `xml:"rulesets,omitempty"`

	// A list of notifier links
	Notifiers []*NotifierLink `xml:"notifiers,omitempty"`

	// If 'Y', service will be published on creation
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFData `xml:"data,omitempty"`
}

type UpdateDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF service to update
	Service_id string `xml:"service_id,omitempty"`

	// A new label for the service
	Label string `xml:"label,omitempty"`

	// Default TTL to be used
	Ttl string `xml:"ttl,omitempty"`

	// A list of fqdn,zone pairs to identify nodes to attach
	Nodes []*DSFNode `xml:"nodes,omitempty"`

	Rulesets []*DSFRuleset `xml:"rulesets,omitempty"`

	// A list of notifier links
	Notifiers []*NotifierLink `xml:"notifiers,omitempty"`

	// If true, the service is immediately published
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFData `xml:"data,omitempty"`
}

type GetDSFsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label, ID, or wildcard
	Cust string `xml:"cust,omitempty"`

	// label or ID
	Service_id string `xml:"service_id,omitempty"`

	// wildcard
	Label string `xml:"label,omitempty"`

	// linker
	Linker string `xml:"linker,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFData `xml:"data,omitempty"`
}

type GetDSFNotifiersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFNotifiersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Identifier for DSF service to search on
	Service_id string `xml:"service_id,omitempty"`
}

type GetDSFNotifiersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFNotifiersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NotifierDataAlt `xml:"data,omitempty"`
}

type DeleteOneDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF service to delete
	Service_id string `xml:"service_id,omitempty"`
}

type DeleteOneDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetOneDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label or ID
	Service_id string `xml:"service_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFData `xml:"data,omitempty"`
}

type RevertDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RevertDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label or ID
	Service_id string `xml:"service_id,omitempty"`
}

type RevertDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RevertDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFData `xml:"data,omitempty"`
}

type PublishDSFRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PublishDSFRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label or ID
	Service_id string `xml:"service_id,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type PublishDSFResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PublishDSFResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFData `xml:"data,omitempty"`
}

type AddDSFNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Identifier for DSF service to search on
	Service_id string `xml:"service_id,omitempty"`

	// Public_id of the Notifier to link to
	Notifier_id string `xml:"notifier_id,omitempty"`

	// filters on when services should fire the notifier
	Filters []string `xml:"filters,omitempty"`

	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type AddDSFNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierLinkData `xml:"data,omitempty"`
}

type RemoveDSFNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveDSFNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Identifier for DSF service to search on
	Service_id string `xml:"service_id,omitempty"`

	// Public_id of the Notifier to link to
	Link_id string `xml:"link_id,omitempty"`

	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type RemoveDSFNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveDSFNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierLinkData `xml:"data,omitempty"`
}

type CreateDSFRulesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRulesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// Ruleset label
	Label string `xml:"label,omitempty"`

	// The type of criteria contained within this Pool
	Criteria_type string `xml:"criteria_type,omitempty"`

	// Required based on criteria_type. Filtered in API/BLL
	Criteria *DSFCriteria `xml:"criteria,omitempty"`

	// Where in the chain does the ruleset land. Defautls to the last.
	Ordering string `xml:"ordering,omitempty"`

	// A list of DSF Reponse Pools that comprise the Ruleset
	Response_pools []*DSFResponsePool `xml:"response_pools,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFRulesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRulesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRulesetData `xml:"data,omitempty"`
}

type UpdateDSFRulesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRulesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID of the ruleset to update
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	// Ruleset label
	Label string `xml:"label,omitempty"`

	// The type of criteria contained within this Pool
	Criteria_type string `xml:"criteria_type,omitempty"`

	// Required based on criteria_type. Filtered in API/BLL
	Criteria *DSFCriteria `xml:"criteria,omitempty"`

	// Where in the chain does the ruleset land. Defautls to the last.
	Ordering string `xml:"ordering,omitempty"`

	// A list of DSF Reponse Pools that comprise the Ruleset
	Response_pools []*DSFResponsePool `xml:"response_pools,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFRulesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRulesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRulesetData `xml:"data,omitempty"`
}

type GetDSFRulesetsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRulesetsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// Ruleset label to search on. Can be wildcarded
	Label string `xml:"label,omitempty"`

	// Can be wildcarded...must be string representation
	Criteria string `xml:"criteria,omitempty"`

	// The type of criteria contained within this Pool
	Criteria_type string `xml:"criteria_type,omitempty"`

	// Where in the chain does the ruleset land. Defautls to the last.
	Ordering string `xml:"ordering,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFRulesetsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRulesetsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFRulesetData `xml:"data,omitempty"`
}

type GetOneDSFRulesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRulesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID of the ruleset to update
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFRulesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRulesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRulesetData `xml:"data,omitempty"`
}

type DeleteOneDSFRulesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRulesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID of the ruleset to update
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFRulesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRulesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRulesetData `xml:"data,omitempty"`
}

type CreateDSFResponsePoolRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFResponsePoolRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// Response Pool label
	Label string `xml:"label,omitempty"`

	Core_set_count string `xml:"core_set_count,omitempty"`

	Automation string `xml:"automation,omitempty"`

	// Defaults to true
	Eligible string `xml:"eligible,omitempty"`

	// ID or label of the DSF Ruleset to join
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	// Index within the specified DSF Ruleset
	Index string `xml:"index,omitempty"`

	Rs_chains []*DSFRecordSetFailoverChain `xml:"rs_chains,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFResponsePoolResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFResponsePoolResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFResponsePoolData `xml:"data,omitempty"`
}

type UpdateDSFResponsePoolRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFResponsePoolRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the Response Pool to update
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// Response Pool label
	Label string `xml:"label,omitempty"`

	Core_set_count string `xml:"core_set_count,omitempty"`

	Automation string `xml:"automation,omitempty"`

	// Defaults to true
	Eligible string `xml:"eligible,omitempty"`

	// Entire chain must be specified
	Rs_chains []*DSFRecordSetFailoverChain `xml:"rs_chains,omitempty"`

	// ID or label of the DSF Ruleset to join
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	// If true, removes record-sets that are no longer referenced by anyone
	Remove_orphans string `xml:"remove_orphans,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFResponsePoolResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFResponsePoolResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFResponsePoolData `xml:"data,omitempty"`
}

type GetDSFResponsePoolsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFResponsePoolsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	Label string `xml:"label,omitempty"`

	// ID or label of a DSF Ruleset that could contain
	Dsf_ruleset_id string `xml:"dsf_ruleset_id,omitempty"`

	Core_set_count string `xml:"core_set_count,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Status string `xml:"status,omitempty"`

	// bool, if true, finds pools that do not exist in any Rulesets
	No_ruleset string `xml:"no_ruleset,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFResponsePoolsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFResponsePoolsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFResponsePoolData `xml:"data,omitempty"`
}

type GetOneDSFResponsePoolRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFResponsePoolRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or the label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the desired response pool
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFResponsePoolResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFResponsePoolResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFResponsePoolData `xml:"data,omitempty"`
}

type DeleteOneDSFResponsePoolRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFResponsePoolRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or the label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the desired response pool
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// If 'Y', Pool will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFResponsePoolResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFResponsePoolResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFResponsePoolData `xml:"data,omitempty"`
}

type CreateDSFRecordSetFailoverChainRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordSetFailoverChainRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID of the DSF Response Pool that the Rec Set Fail chain belongs to
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// Label of the DSF Record Set Failover Chain
	Label string `xml:"label,omitempty"`

	// Optional, defaults to false, indicates whether enclosed Record Sets are Core
	Core string `xml:"core,omitempty"`

	// A list of record sets to be included in this chain
	Record_sets []*DSFRecordSet `xml:"record_sets,omitempty"`

	// If 'Y', RS Chain will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFRecordSetFailoverChainResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordSetFailoverChainResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetFailoverChainData `xml:"data,omitempty"`
}

type UpdateDSFRecordSetFailoverChainRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordSetFailoverChainRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Rec Set Failover Chain
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// Label of the DSF Record Set Failover Chain
	Label string `xml:"label,omitempty"`

	// Optional, defaults to false, indicates whether enclosed Record Sets are Core
	Core string `xml:"core,omitempty"`

	// A list of record sets to be included in this chain
	Record_sets []*DSFRecordSet `xml:"record_sets,omitempty"`

	// If 'Y', RS Chain will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFRecordSetFailoverChainResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordSetFailoverChainResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetFailoverChainData `xml:"data,omitempty"`
}

type GetDSFRecordSetFailoverChainsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordSetFailoverChainsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Response Pool that the Rec Set Fail chain belongs to
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// Label of the DSF Record Set Failover Chain
	Label string `xml:"label,omitempty"`

	// Search for core DSF Record Set Failover Chains that are core
	Core string `xml:"core,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFRecordSetFailoverChainsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordSetFailoverChainsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFRecordSetFailoverChainData `xml:"data,omitempty"`
}

type GetOneDSFRecordSetFailoverChainRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordSetFailoverChainRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Rec Set Failover Chain
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFRecordSetFailoverChainResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordSetFailoverChainResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetFailoverChainData `xml:"data,omitempty"`
}

type DeleteOneDSFRecordSetFailoverChainRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordSetFailoverChainRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Rec Set Failover Chain
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// If 'Y', Pool will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFRecordSetFailoverChainResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordSetFailoverChainResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetFailoverChainData `xml:"data,omitempty"`
}

type CreateDSFRecordSetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordSetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// class of rdata that the set will contain
	Rdata_class string `xml:"rdata_class,omitempty"`

	// Record Set label
	Label string `xml:"label,omitempty"`

	// ID or label of the associated monitor
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Serve_count string `xml:"serve_count,omitempty"`

	Fail_count string `xml:"fail_count,omitempty"`

	Trouble_count string `xml:"trouble_count,omitempty"`

	Torpidity_max string `xml:"torpidity_max,omitempty"`

	// list of hashes that contain information to create DSF Records
	Records []*DSFRecord `xml:"records,omitempty"`

	// ID or label to associate the new RS with an existing RP
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// Index of the RS in the specified chain
	Index string `xml:"index,omitempty"`

	// ID or label to associate the new RS with an existing RP
	Dsf_response_pool_id string `xml:"dsf_response_pool_id,omitempty"`

	// Defaults to true
	Eligible string `xml:"eligible,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFRecordSetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordSetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetData `xml:"data,omitempty"`
}

type UpdateDSFRecordSetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordSetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Record Set
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	// new label for the DSF Record Set
	Label string `xml:"label,omitempty"`

	// ID or label of the associated monitor
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Serve_count string `xml:"serve_count,omitempty"`

	Fail_count string `xml:"fail_count,omitempty"`

	Trouble_count string `xml:"trouble_count,omitempty"`

	Torpidity_max string `xml:"torpidity_max,omitempty"`

	// Defaults to true
	Eligible string `xml:"eligible,omitempty"`

	// hash of information to create DSF Records
	Records []*DSFRecord `xml:"records,omitempty"`

	// boolean - immediately save change and publish
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFRecordSetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordSetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetData `xml:"data,omitempty"`
}

type GetOneDSFRecordSetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordSetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Record Set
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFRecordSetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordSetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetData `xml:"data,omitempty"`
}

type GetDSFRecordSetsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordSetsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// Record Set label - can be wildcarded
	Label string `xml:"label,omitempty"`

	Ttl string `xml:"ttl,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Serve_count string `xml:"serve_count,omitempty"`

	Fail_count string `xml:"fail_count,omitempty"`

	Trouble_count string `xml:"trouble_count,omitempty"`

	Torpidity_max string `xml:"torpidity_max,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Status string `xml:"status,omitempty"`

	// class of rdata that the set will contain
	Rdata_class string `xml:"rdata_class,omitempty"`

	// ID or label to associate the RS with an existing DSF Record Set Failover Chain
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// ID or label of the associated monitor
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFRecordSetsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordSetsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFRecordSetData `xml:"data,omitempty"`
}

type DeleteOneDSFRecordSetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordSetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID or label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// ID or label of the DSF Record Set
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	// If 'Y', Record Set will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFRecordSetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordSetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordSetData `xml:"data,omitempty"`
}

type CreateDSFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// label of the DSF Record Set
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	Master_line string `xml:"master_line,omitempty"`

	// Rdata to create the svc record with
	Rdata *ANYOneRData `xml:"rdata,omitempty"`

	// name of the DSF record
	Label string `xml:"label,omitempty"`

	Weight string `xml:"weight,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Endpoints []string `xml:"endpoints,omitempty"`

	// number of endpoints that need to be up
	Endpoint_up_count string `xml:"endpoint_up_count,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	// boolean, if true add and immediately publish
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type CreateDSFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordData `xml:"data,omitempty"`
}

type UpdateDSFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// label of the DSF Record
	Dsf_record_id string `xml:"dsf_record_id,omitempty"`

	Master_line string `xml:"master_line,omitempty"`

	// Rdata to update the svc record with
	Rdata *ANYOneRData `xml:"rdata,omitempty"`

	// name of the DSF record
	Label string `xml:"label,omitempty"`

	Weight string `xml:"weight,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Endpoints []string `xml:"endpoints,omitempty"`

	// number of endpoints that need to be up
	Endpoint_up_count string `xml:"endpoint_up_count,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	// boolean, if true add and immediately publish
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordData `xml:"data,omitempty"`
}

type GetOneDSFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// label of the DSF Record
	Dsf_record_id string `xml:"dsf_record_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetOneDSFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordData `xml:"data,omitempty"`
}

type GetDSFRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// DSF Record Set Failover Chain ID to search on
	Dsf_record_set_failover_chain_id string `xml:"dsf_record_set_failover_chain_id,omitempty"`

	// Record set id to search on
	Dsf_record_set_id string `xml:"dsf_record_set_id,omitempty"`

	// Wildcard allowed
	Label string `xml:"label,omitempty"`

	// Wildcard allowed
	Master_line string `xml:"master_line,omitempty"`

	// Wildcard allowed
	Endpoints string `xml:"endpoints,omitempty"`

	Endpoint_up_count string `xml:"endpoint_up_count,omitempty"`

	Weight string `xml:"weight,omitempty"`

	Automation string `xml:"automation,omitempty"`

	Eligible string `xml:"eligible,omitempty"`

	Status string `xml:"status,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFRecordData `xml:"data,omitempty"`
}

type DeleteOneDSFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// label of the DSF Record
	Dsf_record_id string `xml:"dsf_record_id,omitempty"`

	// If 'Y', Record will be deleted on execution
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFRecordData `xml:"data,omitempty"`
}

type AddDSFNodeRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFNodeRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// an fqdn, zone pair
	Node *DSFNode `xml:"node,omitempty"`

	// If 'Y', change is published immediately
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type AddDSFNodeResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFNodeResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFNode `xml:"data,omitempty"`
}

type UpdateDSFNodesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFNodesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// A list of fqdn, zone pairs
	Nodes []*DSFNode `xml:"nodes,omitempty"`

	// If 'Y', change is published immediately
	Publish string `xml:"publish,omitempty"`

	// Optional Publish Notes.
	Notes string `xml:"notes,omitempty"`
}

type UpdateDSFNodesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFNodesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFNode `xml:"data,omitempty"`
}

type GetDSFNodesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFNodesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// response will include pending changes
	Pending_changes string `xml:"pending_changes,omitempty"`
}

type GetDSFNodesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFNodesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFNode `xml:"data,omitempty"`
}

type DeleteOneDSFNodeRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFNodeRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// label of the DSF Service
	Service_id string `xml:"service_id,omitempty"`

	// an fqdn, zone pair
	Node *DSFNode `xml:"node,omitempty"`

	// If 'Y', change is published immediately
	Publish string `xml:"publish,omitempty"`

	// Optional notes
	Notes string `xml:"notes,omitempty"`
}

type DeleteOneDSFNodeResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFNodeResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFNode `xml:"data,omitempty"`
}

type CreateDSFMonitorRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFMonitorRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Label for the DSF Monitor
	Label string `xml:"label,omitempty"`

	// Num of responses to determine status
	Response_count string `xml:"response_count,omitempty"`

	// Interval, in seconds, between probes
	Probe_interval string `xml:"probe_interval,omitempty"`

	// number of attempted retries on failure before giving up
	Retries string `xml:"retries,omitempty"`

	// name of the protocol to monitor
	Protocol string `xml:"protocol,omitempty"`

	// indicates if the monitor is active, default is N
	Active string `xml:"active,omitempty"`

	// options pertaining the monitor
	Options *DSFMonitorOptions `xml:"options,omitempty"`

	// Endpoints to monitor
	Endpoints []*DSFMonitorEndpoint `xml:"endpoints,omitempty"`
}

type CreateDSFMonitorResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSFMonitorResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFMonitorData `xml:"data,omitempty"`
}

type UpdateDSFMonitorRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFMonitorRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	// New label for the DSF Monitor
	Label string `xml:"label,omitempty"`

	// Num of responses to determine status
	Response_count string `xml:"response_count,omitempty"`

	// Interval, in seconds, between probes
	Probe_interval string `xml:"probe_interval,omitempty"`

	// number of attempted retries on failure before giving up
	Retries string `xml:"retries,omitempty"`

	// name of the protocol to monitor
	Protocol string `xml:"protocol,omitempty"`

	// options pertaining the monitor
	Options *DSFMonitorOptions `xml:"options,omitempty"`

	// Endpoints to monitor
	Endpoints []*DSFMonitorEndpoint `xml:"endpoints,omitempty"`
}

type UpdateDSFMonitorResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSFMonitorResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFMonitorData `xml:"data,omitempty"`
}

type GetOneDSFMonitorRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFMonitorRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`
}

type GetOneDSFMonitorResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSFMonitorResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFMonitorData `xml:"data,omitempty"`
}

type GetDSFMonitorsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFMonitorsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Label for the DSF Monitor
	Label string `xml:"label,omitempty"`

	// Interval, in seconds, between probes
	Probe_interval string `xml:"probe_interval,omitempty"`

	// number of attempted retries on failure before giving up
	Retries string `xml:"retries,omitempty"`

	// name of the protocol to monitor
	Protocol string `xml:"protocol,omitempty"`
}

type GetDSFMonitorsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFMonitorsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSFMonitorData `xml:"data,omitempty"`
}

type DeleteOneDSFMonitorRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFMonitorRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// ID
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`
}

type DeleteOneDSFMonitorResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSFMonitorResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddDSFMonitorNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFMonitorNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Monitor ID
	Dsf_monitor_id string `xml:"dsf_monitor_id,omitempty"`

	// Notifier ID passed in for existing Notifier, or the follow params used to create
	Dsf_notify_id string `xml:"dsf_notify_id,omitempty"`

	// filters on when services should fire the notifier
	Filters []string `xml:"filters,omitempty"`
}

type AddDSFMonitorNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDSFMonitorNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierData `xml:"data,omitempty"`
}

type GetDSFMonitorSitesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFMonitorSitesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetDSFMonitorSitesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSFMonitorSitesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSFMonitorSitesData `xml:"data,omitempty"`
}

type CreateNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Label for the Notifier
	Label string `xml:"label,omitempty"`

	// List of Recipients attached to the Notifier
	Recipients []*Recipient `xml:"recipients,omitempty"`

	// List of Services attached to the Notifier
	Services []*Service `xml:"services,omitempty"`
}

type CreateNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierData `xml:"data,omitempty"`
}

type UpdateNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Public_id of the Notifier to update
	Notifier_id string `xml:"notifier_id,omitempty"`

	// Label for the DSF Notify
	Label string `xml:"label,omitempty"`

	// List of Recipients attached to the Notifier
	Recipients []*Recipient `xml:"recipients,omitempty"`

	// List of Services attached to the Notifier
	Services []*Service `xml:"services,omitempty"`
}

type UpdateNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierData `xml:"data,omitempty"`
}

type GetOneNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Public_id of the Notifier to retrieve
	Notifier_id string `xml:"notifier_id,omitempty"`
}

type GetOneNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NotifierData `xml:"data,omitempty"`
}

type GetNotifiersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNotifiersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Label for the DSF Notify
	Label string `xml:"label,omitempty"`

	// Search for active or inactive notifiers
	Active string `xml:"active,omitempty"`

	// Search for email or syslog recipients
	Format string `xml:"format,omitempty"`

	// Search on the recipient email, hostname, or contact
	Recipient string `xml:"recipient,omitempty"`

	// Search for active or inactive recipients
	Recipient_active string `xml:"recipient_active,omitempty"`

	// Must be specified with service public id
	Service_class string `xml:"service_class,omitempty"`

	// Public_id of the service_class item to search for
	Service_id string `xml:"service_id,omitempty"`

	// Search for active or inactive services
	Service_active string `xml:"service_active,omitempty"`
}

type GetNotifiersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNotifiersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NotifierDataAlt `xml:"data,omitempty"`
}

type DeleteOneNotifierRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNotifierRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// Public_id of the Notifier to delete
	Notifier_id string `xml:"notifier_id,omitempty"`
}

type DeleteOneNotifierResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNotifierResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateConfigLimitRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateConfigLimitRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CreateConfigLimitResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateConfigLimitResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ConfigLimitData `xml:"data,omitempty"`
}

type GetOneConfigLimitRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneConfigLimitRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Name string `xml:"name,omitempty"`
}

type GetOneConfigLimitResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneConfigLimitResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ConfigLimitData `xml:"data,omitempty"`
}

type GetConfigLimitsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetConfigLimitsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Service_id string `xml:"service_id,omitempty"`
}

type GetConfigLimitsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetConfigLimitsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ConfigLimitData `xml:"data,omitempty"`
}

type UpdateConfigLimitRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateConfigLimitRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type UpdateConfigLimitResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateConfigLimitResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ConfigLimitData `xml:"data,omitempty"`
}

type DeleteOneConfigLimitRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneConfigLimitRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Service_id string `xml:"service_id,omitempty"`

	Name string `xml:"name,omitempty"`
}

type DeleteOneConfigLimitResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneConfigLimitResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreatePermissionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePermissionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	All_users string `xml:"all_users,omitempty"`

	Description string `xml:"description,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	User_name []string `xml:"user_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type CreatePermissionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePermissionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionGroupData `xml:"data,omitempty"`
}

type GetOnePermissionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePermissionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`
}

type GetOnePermissionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePermissionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionGroupData `xml:"data,omitempty"`
}

type GetPermissionGroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPermissionGroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetPermissionGroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPermissionGroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*PermissionGroupData `xml:"data,omitempty"`
}

type DeleteOnePermissionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePermissionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`
}

type DeleteOnePermissionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePermissionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type UpdatePermissionGroupRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePermissionGroupRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	New_group_name string `xml:"new_group_name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	All_users string `xml:"all_users,omitempty"`

	Description string `xml:"description,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	User_name []string `xml:"user_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type UpdatePermissionGroupResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePermissionGroupResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionGroupData `xml:"data,omitempty"`
}

type GetCustomerPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`
}

type GetCustomerPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionResponse `xml:"data,omitempty"`
}

type GetUserPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUserPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`
}

type GetUserPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUserPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionResponse `xml:"data,omitempty"`
}

type CheckPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CheckPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	Zone_name string `xml:"zone_name,omitempty"`
}

type CheckPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CheckPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PermissionResponse `xml:"data,omitempty"`
}

type AddPermissionGroupUsersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupUsersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	User_name []string `xml:"user_name,omitempty"`
}

type AddPermissionGroupUsersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupUsersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetPermissionGroupUsersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupUsersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	User_name []string `xml:"user_name,omitempty"`
}

type SetPermissionGroupUsersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupUsersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemovePermissionGroupUsersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupUsersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	User_name []string `xml:"user_name,omitempty"`
}

type RemovePermissionGroupUsersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupUsersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddPermissionGroupSubgroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupSubgroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`
}

type AddPermissionGroupSubgroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupSubgroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetPermissionGroupSubgroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupSubgroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`
}

type SetPermissionGroupSubgroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupSubgroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemovePermissionGroupSubgroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupSubgroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Subgroup []string `xml:"subgroup,omitempty"`
}

type RemovePermissionGroupSubgroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupSubgroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddPermissionGroupPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type AddPermissionGroupPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetPermissionGroupPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type SetPermissionGroupPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemovePermissionGroupPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type RemovePermissionGroupPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddPermissionGroupZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type AddPermissionGroupZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddPermissionGroupZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetPermissionGroupZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type SetPermissionGroupZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetPermissionGroupZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemovePermissionGroupZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Group_name string `xml:"group_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type RemovePermissionGroupZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemovePermissionGroupZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddUserGroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserGroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Group []string `xml:"group,omitempty"`
}

type AddUserGroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserGroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetUserGroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserGroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Group []string `xml:"group,omitempty"`
}

type SetUserGroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserGroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveUserGroupsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserGroupsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Group []string `xml:"group,omitempty"`
}

type RemoveUserGroupsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserGroupsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddUserZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type AddUserZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetUserZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type SetUserZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveUserZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`
}

type RemoveUserZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddUserPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type AddUserPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetUserPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type SetUserPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveUserPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type RemoveUserPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddUserForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Forbid []*PermissionData `xml:"forbid,omitempty"`
}

type AddUserForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddUserForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetUserForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Forbid []*PermissionData `xml:"forbid,omitempty"`
}

type SetUserForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetUserForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveUserForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Forbid []*PermissionData `xml:"forbid,omitempty"`
}

type RemoveUserForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveUserForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddCustomerPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddCustomerPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type AddCustomerPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddCustomerPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetCustomerPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type SetCustomerPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveCustomerPermissionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveCustomerPermissionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Permission []string `xml:"permission,omitempty"`
}

type RemoveCustomerPermissionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveCustomerPermissionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type AddCustomerForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddCustomerForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Forbid []string `xml:"forbid,omitempty"`
}

type AddCustomerForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddCustomerForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type SetCustomerForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Forbid []string `xml:"forbid,omitempty"`
}

type SetCustomerForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RemoveCustomerForbidsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveCustomerForbidsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Forbid []string `xml:"forbid,omitempty"`
}

type RemoveCustomerForbidsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RemoveCustomerForbidsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetHostStatsFlagsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetHostStatsFlagsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`
}

type GetHostStatsFlagsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetHostStatsFlagsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*HostStatFlagsData `xml:"data,omitempty"`
}

type SetHostStatsFlagsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetHostStatsFlagsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Host_stats []*HostStatFlagsData `xml:"host_stats,omitempty"`
}

type SetHostStatsFlagsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetHostStatsFlagsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*HostStatFlagsData `xml:"data,omitempty"`
}

type CreateTSIGKeyRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTSIGKeyRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Name string `xml:"name,omitempty"`

	Algorithm string `xml:"algorithm,omitempty"`

	Secret string `xml:"secret,omitempty"`
}

type CreateTSIGKeyResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTSIGKeyResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TSIGKeyData `xml:"data,omitempty"`
}

type GetOneTSIGKeyRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTSIGKeyRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Name string `xml:"name,omitempty"`
}

type GetOneTSIGKeyResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTSIGKeyResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TSIGKeyData `xml:"data,omitempty"`
}

type GetTSIGKeysRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTSIGKeysRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetTSIGKeysResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTSIGKeysResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TSIGKeyData `xml:"data,omitempty"`
}

type UpdateTSIGKeyRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTSIGKeyRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Name string `xml:"name,omitempty"`

	New_name string `xml:"new_name,omitempty"`

	Secret string `xml:"secret,omitempty"`

	Algorithm string `xml:"algorithm,omitempty"`
}

type UpdateTSIGKeyResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTSIGKeyResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TSIGKeyData `xml:"data,omitempty"`
}

type DeleteOneTSIGKeyRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTSIGKeyRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Name string `xml:"name,omitempty"`
}

type DeleteOneTSIGKeyResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTSIGKeyResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// address of responsible party, per SOA
	Rname string `xml:"rname,omitempty"`

	// default TTL (Time-to-Live) for records
	Ttl int32 `xml:"ttl,omitempty"`

	// code indicating how serial numbers are constructed on publish
	Serial_style string `xml:"serial_style,omitempty"`
}

type CreateZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneData `xml:"data,omitempty"`
}

type GetOneZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetOneZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneData `xml:"data,omitempty"`
}

type GetZonesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZonesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetZonesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZonesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneData `xml:"data,omitempty"`
}

type DeleteOneZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteOneZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateSecondaryZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSecondaryZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Masters []string `xml:"masters,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Contact_nickname string `xml:"contact_nickname,omitempty"`
}

type CreateSecondaryZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSecondaryZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type UpdateSecondaryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSecondaryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Masters []string `xml:"masters,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Contact_nickname string `xml:"contact_nickname,omitempty"`
}

type UpdateSecondaryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSecondaryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type ActivateSecondaryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateSecondaryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type ActivateSecondaryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateSecondaryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type DeactivateSecondaryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateSecondaryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type DeactivateSecondaryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateSecondaryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type RetransferSecondaryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RetransferSecondaryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type RetransferSecondaryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RetransferSecondaryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type GetOneSecondaryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSecondaryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type GetOneSecondaryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSecondaryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SecondaryData `xml:"data,omitempty"`
}

type GetSecondariesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSecondariesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetSecondariesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSecondariesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SecondaryData `xml:"data,omitempty"`
}

type GetZoneApexRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneApexRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// FQDN of a node
	Node string `xml:"node,omitempty"`
}

type GetZoneApexResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneApexResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneData `xml:"data,omitempty"`
}

type CreateARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataA `xml:"rdata,omitempty"`
}

type CreateARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ARecordData `xml:"data,omitempty"`
}

type GetOneARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ARecordData `xml:"data,omitempty"`
}

type GetARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ARecordData `xml:"data,omitempty"`
}

type UpdateARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ARecordData `xml:"data,omitempty"`
}

type DeleteARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataA `xml:"rdata,omitempty"`
}

type DeleteOneARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateAAAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAAAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataAAAA `xml:"rdata,omitempty"`
}

type CreateAAAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAAAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AAAARecordData `xml:"data,omitempty"`
}

type GetOneAAAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAAAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataAAAA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneAAAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAAAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AAAARecordData `xml:"data,omitempty"`
}

type GetAAAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAAAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetAAAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAAAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*AAAARecordData `xml:"data,omitempty"`
}

type UpdateAAAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAAAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataAAAA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateAAAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAAAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AAAARecordData `xml:"data,omitempty"`
}

type DeleteAAAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteAAAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteAAAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteAAAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneAAAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAAAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataAAAA `xml:"rdata,omitempty"`
}

type DeleteOneAAAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAAAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateALIASRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateALIASRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataALIAS `xml:"rdata,omitempty"`
}

type CreateALIASRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateALIASRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ALIASRecordData `xml:"data,omitempty"`
}

type GetOneALIASRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneALIASRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataALIAS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneALIASRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneALIASRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ALIASRecordData `xml:"data,omitempty"`
}

type GetALIASRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetALIASRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetALIASRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetALIASRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ALIASRecordData `xml:"data,omitempty"`
}

type UpdateALIASRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateALIASRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataALIAS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateALIASRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateALIASRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ALIASRecordData `xml:"data,omitempty"`
}

type DeleteALIASRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteALIASRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteALIASRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteALIASRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneALIASRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneALIASRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataALIAS `xml:"rdata,omitempty"`
}

type DeleteOneALIASRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneALIASRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCAA `xml:"rdata,omitempty"`
}

type CreateCAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CAARecordData `xml:"data,omitempty"`
}

type GetOneCAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCAA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CAARecordData `xml:"data,omitempty"`
}

type GetCAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CAARecordData `xml:"data,omitempty"`
}

type UpdateCAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCAA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CAARecordData `xml:"data,omitempty"`
}

type DeleteCAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCAARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCAARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCAA `xml:"rdata,omitempty"`
}

type DeleteOneCAARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCAARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCDNSKEY `xml:"rdata,omitempty"`
}

type CreateCDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNSKEYRecordData `xml:"data,omitempty"`
}

type GetOneCDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDNSKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNSKEYRecordData `xml:"data,omitempty"`
}

type GetCDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CDNSKEYRecordData `xml:"data,omitempty"`
}

type UpdateCDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDNSKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNSKEYRecordData `xml:"data,omitempty"`
}

type DeleteCDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDNSKEY `xml:"rdata,omitempty"`
}

type DeleteOneCDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCDS `xml:"rdata,omitempty"`
}

type CreateCDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDSRecordData `xml:"data,omitempty"`
}

type GetOneCDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDSRecordData `xml:"data,omitempty"`
}

type GetCDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CDSRecordData `xml:"data,omitempty"`
}

type UpdateCDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDSRecordData `xml:"data,omitempty"`
}

type DeleteCDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCDS `xml:"rdata,omitempty"`
}

type DeleteOneCDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCERTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCERTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCERT `xml:"rdata,omitempty"`
}

type CreateCERTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCERTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CERTRecordData `xml:"data,omitempty"`
}

type GetOneCERTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCERTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCERT `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCERTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCERTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CERTRecordData `xml:"data,omitempty"`
}

type GetCERTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCERTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCERTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCERTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CERTRecordData `xml:"data,omitempty"`
}

type UpdateCERTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCERTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCERT `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCERTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCERTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CERTRecordData `xml:"data,omitempty"`
}

type DeleteCERTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCERTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCERTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCERTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCERTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCERTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCERT `xml:"rdata,omitempty"`
}

type DeleteOneCERTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCERTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCNAME `xml:"rdata,omitempty"`
}

type CreateCNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CNAMERecordData `xml:"data,omitempty"`
}

type GetOneCNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCNAME `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CNAMERecordData `xml:"data,omitempty"`
}

type GetCNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CNAMERecordData `xml:"data,omitempty"`
}

type UpdateCNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCNAME `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CNAMERecordData `xml:"data,omitempty"`
}

type DeleteCNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCNAME `xml:"rdata,omitempty"`
}

type DeleteOneCNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateCSYNCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCSYNCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataCSYNC `xml:"rdata,omitempty"`
}

type CreateCSYNCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCSYNCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CSYNCRecordData `xml:"data,omitempty"`
}

type GetOneCSYNCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCSYNCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCSYNC `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCSYNCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCSYNCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CSYNCRecordData `xml:"data,omitempty"`
}

type GetCSYNCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCSYNCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCSYNCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCSYNCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CSYNCRecordData `xml:"data,omitempty"`
}

type UpdateCSYNCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCSYNCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCSYNC `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCSYNCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCSYNCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CSYNCRecordData `xml:"data,omitempty"`
}

type DeleteCSYNCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCSYNCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteCSYNCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCSYNCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneCSYNCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCSYNCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataCSYNC `xml:"rdata,omitempty"`
}

type DeleteOneCSYNCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCSYNCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateDHCIDRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDHCIDRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataDHCID `xml:"rdata,omitempty"`
}

type CreateDHCIDRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDHCIDRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DHCIDRecordData `xml:"data,omitempty"`
}

type GetOneDHCIDRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDHCIDRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDHCID `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneDHCIDRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDHCIDRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DHCIDRecordData `xml:"data,omitempty"`
}

type GetDHCIDRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDHCIDRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetDHCIDRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDHCIDRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DHCIDRecordData `xml:"data,omitempty"`
}

type UpdateDHCIDRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDHCIDRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDHCID `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateDHCIDRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDHCIDRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DHCIDRecordData `xml:"data,omitempty"`
}

type DeleteDHCIDRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDHCIDRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteDHCIDRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDHCIDRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneDHCIDRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDHCIDRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDHCID `xml:"rdata,omitempty"`
}

type DeleteOneDHCIDRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDHCIDRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateDNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataDNAME `xml:"rdata,omitempty"`
}

type CreateDNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNAMERecordData `xml:"data,omitempty"`
}

type GetOneDNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNAME `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneDNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNAMERecordData `xml:"data,omitempty"`
}

type GetDNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetDNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNAMERecordData `xml:"data,omitempty"`
}

type UpdateDNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNAME `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateDNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNAMERecordData `xml:"data,omitempty"`
}

type DeleteDNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteDNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneDNAMERecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNAMERecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNAME `xml:"rdata,omitempty"`
}

type DeleteOneDNAMERecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNAMERecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataDNSKEY `xml:"rdata,omitempty"`
}

type CreateDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSKEYRecordData `xml:"data,omitempty"`
}

type GetOneDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNSKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSKEYRecordData `xml:"data,omitempty"`
}

type GetDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNSKEYRecordData `xml:"data,omitempty"`
}

type UpdateDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNSKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSKEYRecordData `xml:"data,omitempty"`
}

type DeleteDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneDNSKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNSKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDNSKEY `xml:"rdata,omitempty"`
}

type DeleteOneDNSKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNSKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataDS `xml:"rdata,omitempty"`
}

type CreateDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSRecordData `xml:"data,omitempty"`
}

type GetOneDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSRecordData `xml:"data,omitempty"`
}

type GetDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSRecordData `xml:"data,omitempty"`
}

type UpdateDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DSRecordData `xml:"data,omitempty"`
}

type DeleteDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneDSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataDS `xml:"rdata,omitempty"`
}

type DeleteOneDSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateIPSECKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateIPSECKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataIPSECKEY `xml:"rdata,omitempty"`
}

type CreateIPSECKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateIPSECKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPSECKEYRecordData `xml:"data,omitempty"`
}

type GetOneIPSECKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneIPSECKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataIPSECKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneIPSECKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneIPSECKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPSECKEYRecordData `xml:"data,omitempty"`
}

type GetIPSECKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetIPSECKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetIPSECKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetIPSECKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*IPSECKEYRecordData `xml:"data,omitempty"`
}

type UpdateIPSECKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateIPSECKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataIPSECKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateIPSECKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateIPSECKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPSECKEYRecordData `xml:"data,omitempty"`
}

type DeleteIPSECKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteIPSECKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteIPSECKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteIPSECKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneIPSECKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneIPSECKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataIPSECKEY `xml:"rdata,omitempty"`
}

type DeleteOneIPSECKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneIPSECKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataKEY `xml:"rdata,omitempty"`
}

type CreateKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KEYRecordData `xml:"data,omitempty"`
}

type GetOneKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KEYRecordData `xml:"data,omitempty"`
}

type GetKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*KEYRecordData `xml:"data,omitempty"`
}

type UpdateKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKEY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KEYRecordData `xml:"data,omitempty"`
}

type DeleteKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneKEYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneKEYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKEY `xml:"rdata,omitempty"`
}

type DeleteOneKEYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneKEYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateKXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateKXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataKX `xml:"rdata,omitempty"`
}

type CreateKXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateKXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KXRecordData `xml:"data,omitempty"`
}

type GetOneKXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneKXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneKXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneKXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KXRecordData `xml:"data,omitempty"`
}

type GetKXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetKXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetKXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetKXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*KXRecordData `xml:"data,omitempty"`
}

type UpdateKXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateKXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateKXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateKXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *KXRecordData `xml:"data,omitempty"`
}

type DeleteKXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteKXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteKXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteKXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneKXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneKXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataKX `xml:"rdata,omitempty"`
}

type DeleteOneKXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneKXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateLOCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLOCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataLOC `xml:"rdata,omitempty"`
}

type CreateLOCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLOCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LOCRecordData `xml:"data,omitempty"`
}

type GetOneLOCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLOCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataLOC `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneLOCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLOCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LOCRecordData `xml:"data,omitempty"`
}

type GetLOCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLOCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetLOCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLOCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*LOCRecordData `xml:"data,omitempty"`
}

type UpdateLOCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLOCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataLOC `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateLOCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLOCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LOCRecordData `xml:"data,omitempty"`
}

type DeleteLOCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteLOCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteLOCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteLOCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneLOCRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLOCRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataLOC `xml:"rdata,omitempty"`
}

type DeleteOneLOCRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLOCRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateMXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateMXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataMX `xml:"rdata,omitempty"`
}

type CreateMXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateMXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *MXRecordData `xml:"data,omitempty"`
}

type GetOneMXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneMXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataMX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneMXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneMXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *MXRecordData `xml:"data,omitempty"`
}

type GetMXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetMXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetMXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetMXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*MXRecordData `xml:"data,omitempty"`
}

type UpdateMXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateMXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataMX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateMXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateMXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *MXRecordData `xml:"data,omitempty"`
}

type DeleteMXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteMXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteMXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteMXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneMXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneMXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataMX `xml:"rdata,omitempty"`
}

type DeleteOneMXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneMXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateNAPTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNAPTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataNAPTR `xml:"rdata,omitempty"`
}

type CreateNAPTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNAPTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NAPTRRecordData `xml:"data,omitempty"`
}

type GetOneNAPTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNAPTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNAPTR `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneNAPTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNAPTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NAPTRRecordData `xml:"data,omitempty"`
}

type GetNAPTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNAPTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetNAPTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNAPTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NAPTRRecordData `xml:"data,omitempty"`
}

type UpdateNAPTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNAPTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNAPTR `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateNAPTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNAPTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NAPTRRecordData `xml:"data,omitempty"`
}

type DeleteNAPTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNAPTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteNAPTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNAPTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneNAPTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNAPTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNAPTR `xml:"rdata,omitempty"`
}

type DeleteOneNAPTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNAPTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateNSAPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNSAPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataNSAP `xml:"rdata,omitempty"`
}

type CreateNSAPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNSAPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSAPRecordData `xml:"data,omitempty"`
}

type GetOneNSAPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNSAPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNSAP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneNSAPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNSAPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSAPRecordData `xml:"data,omitempty"`
}

type GetNSAPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNSAPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetNSAPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNSAPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NSAPRecordData `xml:"data,omitempty"`
}

type UpdateNSAPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNSAPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNSAP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateNSAPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNSAPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSAPRecordData `xml:"data,omitempty"`
}

type DeleteNSAPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNSAPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteNSAPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNSAPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneNSAPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNSAPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNSAP `xml:"rdata,omitempty"`
}

type DeleteOneNSAPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNSAPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreatePOLICYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePOLICYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataPOLICY `xml:"rdata,omitempty"`
}

type CreatePOLICYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePOLICYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *POLICYRecordData `xml:"data,omitempty"`
}

type GetOnePOLICYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePOLICYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPOLICY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOnePOLICYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePOLICYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *POLICYRecordData `xml:"data,omitempty"`
}

type GetPOLICYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPOLICYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetPOLICYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPOLICYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*POLICYRecordData `xml:"data,omitempty"`
}

type UpdatePOLICYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePOLICYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPOLICY `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdatePOLICYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePOLICYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *POLICYRecordData `xml:"data,omitempty"`
}

type DeletePOLICYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePOLICYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeletePOLICYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePOLICYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOnePOLICYRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePOLICYRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPOLICY `xml:"rdata,omitempty"`
}

type DeleteOnePOLICYRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePOLICYRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreatePTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataPTR `xml:"rdata,omitempty"`
}

type CreatePTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PTRRecordData `xml:"data,omitempty"`
}

type GetOnePTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPTR `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOnePTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PTRRecordData `xml:"data,omitempty"`
}

type GetPTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetPTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*PTRRecordData `xml:"data,omitempty"`
}

type UpdatePTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPTR `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdatePTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PTRRecordData `xml:"data,omitempty"`
}

type DeletePTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeletePTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOnePTRRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePTRRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPTR `xml:"rdata,omitempty"`
}

type DeleteOnePTRRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePTRRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreatePXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataPX `xml:"rdata,omitempty"`
}

type CreatePXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreatePXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PXRecordData `xml:"data,omitempty"`
}

type GetOnePXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOnePXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOnePXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PXRecordData `xml:"data,omitempty"`
}

type GetPXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetPXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetPXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*PXRecordData `xml:"data,omitempty"`
}

type UpdatePXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPX `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdatePXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdatePXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PXRecordData `xml:"data,omitempty"`
}

type DeletePXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeletePXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeletePXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOnePXRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePXRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataPX `xml:"rdata,omitempty"`
}

type DeleteOnePXRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOnePXRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateRPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataRP `xml:"rdata,omitempty"`
}

type CreateRPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RPRecordData `xml:"data,omitempty"`
}

type GetOneRPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataRP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneRPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RPRecordData `xml:"data,omitempty"`
}

type GetRPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetRPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RPRecordData `xml:"data,omitempty"`
}

type UpdateRPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataRP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateRPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RPRecordData `xml:"data,omitempty"`
}

type DeleteRPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteRPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteRPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteRPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneRPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataRP `xml:"rdata,omitempty"`
}

type DeleteOneRPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateSPFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSPFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataSPF `xml:"rdata,omitempty"`
}

type CreateSPFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSPFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SPFRecordData `xml:"data,omitempty"`
}

type GetOneSPFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSPFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSPF `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneSPFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSPFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SPFRecordData `xml:"data,omitempty"`
}

type GetSPFRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSPFRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetSPFRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSPFRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SPFRecordData `xml:"data,omitempty"`
}

type UpdateSPFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSPFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSPF `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateSPFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSPFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SPFRecordData `xml:"data,omitempty"`
}

type DeleteSPFRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSPFRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteSPFRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSPFRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneSPFRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSPFRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSPF `xml:"rdata,omitempty"`
}

type DeleteOneSPFRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSPFRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateSRVRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSRVRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataSRV `xml:"rdata,omitempty"`
}

type CreateSRVRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSRVRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SRVRecordData `xml:"data,omitempty"`
}

type GetOneSRVRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSRVRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSRV `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneSRVRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSRVRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SRVRecordData `xml:"data,omitempty"`
}

type GetSRVRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSRVRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetSRVRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSRVRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SRVRecordData `xml:"data,omitempty"`
}

type UpdateSRVRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSRVRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSRV `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateSRVRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSRVRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SRVRecordData `xml:"data,omitempty"`
}

type DeleteSRVRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSRVRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteSRVRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSRVRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneSRVRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSRVRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSRV `xml:"rdata,omitempty"`
}

type DeleteOneSRVRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSRVRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateSSHFPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSSHFPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataSSHFP `xml:"rdata,omitempty"`
}

type CreateSSHFPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateSSHFPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SSHFPRecordData `xml:"data,omitempty"`
}

type GetOneSSHFPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSSHFPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSSHFP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneSSHFPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSSHFPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SSHFPRecordData `xml:"data,omitempty"`
}

type GetSSHFPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSSHFPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetSSHFPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSSHFPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SSHFPRecordData `xml:"data,omitempty"`
}

type UpdateSSHFPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSSHFPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSSHFP `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateSSHFPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSSHFPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SSHFPRecordData `xml:"data,omitempty"`
}

type DeleteSSHFPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSSHFPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteSSHFPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteSSHFPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneSSHFPRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSSHFPRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSSHFP `xml:"rdata,omitempty"`
}

type DeleteOneSSHFPRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneSSHFPRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateTLSARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTLSARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataTLSA `xml:"rdata,omitempty"`
}

type CreateTLSARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTLSARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TLSARecordData `xml:"data,omitempty"`
}

type GetOneTLSARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTLSARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTLSA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneTLSARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTLSARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TLSARecordData `xml:"data,omitempty"`
}

type GetTLSARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTLSARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetTLSARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTLSARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TLSARecordData `xml:"data,omitempty"`
}

type UpdateTLSARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTLSARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTLSA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateTLSARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTLSARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TLSARecordData `xml:"data,omitempty"`
}

type DeleteTLSARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteTLSARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteTLSARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteTLSARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneTLSARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTLSARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTLSA `xml:"rdata,omitempty"`
}

type DeleteOneTLSARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTLSARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateTXTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTXTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataTXT `xml:"rdata,omitempty"`
}

type CreateTXTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateTXTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TXTRecordData `xml:"data,omitempty"`
}

type GetOneTXTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTXTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTXT `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneTXTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTXTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TXTRecordData `xml:"data,omitempty"`
}

type GetTXTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTXTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetTXTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTXTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TXTRecordData `xml:"data,omitempty"`
}

type UpdateTXTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTXTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTXT `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateTXTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateTXTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TXTRecordData `xml:"data,omitempty"`
}

type DeleteTXTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteTXTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteTXTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteTXTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneTXTRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTXTRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataTXT `xml:"rdata,omitempty"`
}

type DeleteOneTXTRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneTXTRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetOneSOARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSOARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataSOA `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneSOARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneSOARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SOARecordData `xml:"data,omitempty"`
}

type GetSOARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSOARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetSOARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetSOARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SOARecordData `xml:"data,omitempty"`
}

type UpdateSOARecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSOARecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Serial_style string `xml:"serial_style,omitempty"`

	Rdata *RDataSOAUpdate `xml:"rdata,omitempty"`
}

type UpdateSOARecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateSOARecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *SOARecordData `xml:"data,omitempty"`
}

type CreateNSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Rdata *RDataNS `xml:"rdata,omitempty"`

	Service_class string `xml:"service_class,omitempty"`
}

type CreateNSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateNSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSRecordData `xml:"data,omitempty"`
}

type GetOneNSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneNSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneNSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSRecordData `xml:"data,omitempty"`
}

type GetNSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetNSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NSRecordData `xml:"data,omitempty"`
}

type UpdateNSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNS `xml:"rdata,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Ttl int32 `xml:"ttl,omitempty"`

	Service_class string `xml:"service_class,omitempty"`
}

type UpdateNSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateNSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *NSRecordData `xml:"data,omitempty"`
}

type DeleteNSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteNSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteNSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteOneNSRecordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNSRecordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`

	Rdata *RDataNS `xml:"rdata,omitempty"`
}

type DeleteOneNSRecordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneNSRecordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ReplaceARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	ARecords []*ARecordData `xml:"ARecords,omitempty"`
}

type ReplaceARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ARecordData `xml:"data,omitempty"`
}

type ReplaceAAAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceAAAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	AAAARecords []*AAAARecordData `xml:"AAAARecords,omitempty"`
}

type ReplaceAAAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceAAAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*AAAARecordData `xml:"data,omitempty"`
}

type ReplaceALIASRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceALIASRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	ALIASRecords []*ALIASRecordData `xml:"ALIASRecords,omitempty"`
}

type ReplaceALIASRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceALIASRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ALIASRecordData `xml:"data,omitempty"`
}

type ReplaceCAARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCAARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CAARecords []*CAARecordData `xml:"CAARecords,omitempty"`
}

type ReplaceCAARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCAARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CAARecordData `xml:"data,omitempty"`
}

type ReplaceCDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CDNSKEYRecords []*CDNSKEYRecordData `xml:"CDNSKEYRecords,omitempty"`
}

type ReplaceCDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CDNSKEYRecordData `xml:"data,omitempty"`
}

type ReplaceCDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CDSRecords []*CDSRecordData `xml:"CDSRecords,omitempty"`
}

type ReplaceCDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CDSRecordData `xml:"data,omitempty"`
}

type ReplaceCERTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCERTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CERTRecords []*CERTRecordData `xml:"CERTRecords,omitempty"`
}

type ReplaceCERTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCERTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CERTRecordData `xml:"data,omitempty"`
}

type ReplaceCNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CNAMERecords []*CNAMERecordData `xml:"CNAMERecords,omitempty"`
}

type ReplaceCNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CNAMERecordData `xml:"data,omitempty"`
}

type ReplaceCSYNCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCSYNCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	CSYNCRecords []*CSYNCRecordData `xml:"CSYNCRecords,omitempty"`
}

type ReplaceCSYNCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceCSYNCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CSYNCRecordData `xml:"data,omitempty"`
}

type ReplaceDHCIDRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDHCIDRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	DHCIDRecords []*DHCIDRecordData `xml:"DHCIDRecords,omitempty"`
}

type ReplaceDHCIDRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDHCIDRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DHCIDRecordData `xml:"data,omitempty"`
}

type ReplaceDNAMERecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDNAMERecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	DNAMERecords []*DNAMERecordData `xml:"DNAMERecords,omitempty"`
}

type ReplaceDNAMERecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDNAMERecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNAMERecordData `xml:"data,omitempty"`
}

type ReplaceDNSKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDNSKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	DNSKEYRecords []*DNSKEYRecordData `xml:"DNSKEYRecords,omitempty"`
}

type ReplaceDNSKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDNSKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNSKEYRecordData `xml:"data,omitempty"`
}

type ReplaceDSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	DSRecords []*DSRecordData `xml:"DSRecords,omitempty"`
}

type ReplaceDSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceDSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DSRecordData `xml:"data,omitempty"`
}

type ReplaceIPSECKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceIPSECKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	IPSECKEYRecords []*IPSECKEYRecordData `xml:"IPSECKEYRecords,omitempty"`
}

type ReplaceIPSECKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceIPSECKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*IPSECKEYRecordData `xml:"data,omitempty"`
}

type ReplaceKEYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceKEYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	KEYRecords []*KEYRecordData `xml:"KEYRecords,omitempty"`
}

type ReplaceKEYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceKEYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*KEYRecordData `xml:"data,omitempty"`
}

type ReplaceKXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceKXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	KXRecords []*KXRecordData `xml:"KXRecords,omitempty"`
}

type ReplaceKXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceKXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*KXRecordData `xml:"data,omitempty"`
}

type ReplaceLOCRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceLOCRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	LOCRecords []*LOCRecordData `xml:"LOCRecords,omitempty"`
}

type ReplaceLOCRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceLOCRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*LOCRecordData `xml:"data,omitempty"`
}

type ReplaceMXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceMXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	MXRecords []*MXRecordData `xml:"MXRecords,omitempty"`
}

type ReplaceMXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceMXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*MXRecordData `xml:"data,omitempty"`
}

type ReplaceNAPTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNAPTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	NAPTRRecords []*NAPTRRecordData `xml:"NAPTRRecords,omitempty"`
}

type ReplaceNAPTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNAPTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NAPTRRecordData `xml:"data,omitempty"`
}

type ReplaceNSAPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNSAPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	NSAPRecords []*NSAPRecordData `xml:"NSAPRecords,omitempty"`
}

type ReplaceNSAPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNSAPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NSAPRecordData `xml:"data,omitempty"`
}

type ReplacePOLICYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePOLICYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	POLICYRecords []*POLICYRecordData `xml:"POLICYRecords,omitempty"`
}

type ReplacePOLICYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePOLICYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*POLICYRecordData `xml:"data,omitempty"`
}

type ReplacePTRRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePTRRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	PTRRecords []*PTRRecordData `xml:"PTRRecords,omitempty"`
}

type ReplacePTRRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePTRRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*PTRRecordData `xml:"data,omitempty"`
}

type ReplacePXRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePXRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	PXRecords []*PXRecordData `xml:"PXRecords,omitempty"`
}

type ReplacePXRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplacePXRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*PXRecordData `xml:"data,omitempty"`
}

type ReplaceRPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceRPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	RPRecords []*RPRecordData `xml:"RPRecords,omitempty"`
}

type ReplaceRPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceRPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RPRecordData `xml:"data,omitempty"`
}

type ReplaceSPFRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSPFRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	SPFRecords []*SPFRecordData `xml:"SPFRecords,omitempty"`
}

type ReplaceSPFRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSPFRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SPFRecordData `xml:"data,omitempty"`
}

type ReplaceSRVRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSRVRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	SRVRecords []*SRVRecordData `xml:"SRVRecords,omitempty"`
}

type ReplaceSRVRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSRVRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SRVRecordData `xml:"data,omitempty"`
}

type ReplaceSSHFPRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSSHFPRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	SSHFPRecords []*SSHFPRecordData `xml:"SSHFPRecords,omitempty"`
}

type ReplaceSSHFPRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceSSHFPRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*SSHFPRecordData `xml:"data,omitempty"`
}

type ReplaceTLSARecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceTLSARecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	TLSARecords []*TLSARecordData `xml:"TLSARecords,omitempty"`
}

type ReplaceTLSARecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceTLSARecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TLSARecordData `xml:"data,omitempty"`
}

type ReplaceTXTRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceTXTRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	TXTRecords []*TXTRecordData `xml:"TXTRecords,omitempty"`
}

type ReplaceTXTRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceTXTRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TXTRecordData `xml:"data,omitempty"`
}

type ReplaceNSRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNSRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	NSRecords []*NSRecordData `xml:"NSRecords,omitempty"`
}

type ReplaceNSRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ReplaceNSRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*NSRecordData `xml:"data,omitempty"`
}

type GetANYRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetANYRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type GetANYRecordsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetANYRecordsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ANYRecordData `xml:"data,omitempty"`
}

type GetAllRecordsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAllRecordsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type GetAllAliasQNamesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAllAliasQNamesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetAllAliasQNamesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAllAliasQNamesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *QNames `xml:"data,omitempty"`
}

type GetOneUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of user
	User_name string `xml:"user_name,omitempty"`
}

type GetOneUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type DeleteOneUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of user
	User_name string `xml:"user_name,omitempty"`
}

type DeleteOneUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Password string `xml:"password,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	Group_name []string `xml:"group_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`

	Forbid []*PermissionData `xml:"forbid,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type CreateUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type UpdateUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	New_user_name string `xml:"new_user_name,omitempty"`

	Password string `xml:"password,omitempty"`

	Require_pw_change string `xml:"require_pw_change,omitempty"`

	Status string `xml:"status,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	Group_name []string `xml:"group_name,omitempty"`

	Zone []*PermissionZone `xml:"zone,omitempty"`

	Forbid []*PermissionData `xml:"forbid,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type UpdateUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type GetUsersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUsersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Search string `xml:"search,omitempty"`
}

type GetUsersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUsersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*UserData `xml:"data,omitempty"`
}

type GetUpdateUsersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUpdateUsersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetUpdateUsersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUpdateUsersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*UserData `xml:"data,omitempty"`
}

type UpdateUpdateUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUpdateUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`

	Password string `xml:"password,omitempty"`
}

type UpdateUpdateUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUpdateUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type DeleteOneUpdateUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneUpdateUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`
}

type DeleteOneUpdateUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneUpdateUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type UpdateUserPasswordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUserPasswordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Password string `xml:"password,omitempty"`
}

type UpdateUserPasswordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateUserPasswordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type BlockUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ BlockUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`
}

type BlockUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ BlockUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type UnblockUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UnblockUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`
}

type UnblockUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UnblockUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UserData `xml:"data,omitempty"`
}

type CreateContactRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateContactRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type CreateContactResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateContactResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ContactData `xml:"data,omitempty"`
}

type GetOneContactRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneContactRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Nickname string `xml:"nickname,omitempty"`
}

type GetOneContactResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneContactResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ContactData `xml:"data,omitempty"`
}

type GetContactsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetContactsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetContactsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetContactsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ContactData `xml:"data,omitempty"`
}

type DeleteOneContactRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneContactRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Nickname string `xml:"nickname,omitempty"`
}

type DeleteOneContactResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneContactResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type UpdateContactRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateContactRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	New_nickname string `xml:"new_nickname,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Phone string `xml:"phone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Email string `xml:"email,omitempty"`

	Notify_email string `xml:"notify_email,omitempty"`

	Pager_email string `xml:"pager_email,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_2 string `xml:"address_2,omitempty"`

	City string `xml:"city,omitempty"`

	State string `xml:"state,omitempty"`

	Post_code string `xml:"post_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Website string `xml:"website,omitempty"`

	Organization string `xml:"organization,omitempty"`
}

type UpdateContactResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateContactResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ContactData `xml:"data,omitempty"`
}

type CreateCustomerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCustomerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Status string `xml:"status,omitempty"`

	Pool_id string `xml:"pool_id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Level string `xml:"level,omitempty"`

	Primary_sales_agent string `xml:"primary_sales_agent,omitempty"`

	Salesforce_id string `xml:"salesforce_id,omitempty"`

	Owner *CustomerAdminData `xml:"owner,omitempty"`

	Billing *CustomerAdminData `xml:"billing,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	Forbid []string `xml:"forbid,omitempty"`
}

type CreateCustomerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCustomerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerData `xml:"data,omitempty"`
}

type UpdateCustomerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCustomerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	New_customer_name string `xml:"new_customer_name,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Status string `xml:"status,omitempty"`

	Note string `xml:"note,omitempty"`

	Pool_id string `xml:"pool_id,omitempty"`

	Activated string `xml:"activated,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Level string `xml:"level,omitempty"`

	Primary_sales_agent string `xml:"primary_sales_agent,omitempty"`

	Salesforce_id string `xml:"salesforce_id,omitempty"`

	Owner_contact string `xml:"owner_contact,omitempty"`

	Billing_contact string `xml:"billing_contact,omitempty"`

	Permission []string `xml:"permission,omitempty"`

	Forbid []string `xml:"forbid,omitempty"`
}

type UpdateCustomerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCustomerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerData `xml:"data,omitempty"`
}

type GetOneCustomerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCustomerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`
}

type GetOneCustomerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCustomerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerData `xml:"data,omitempty"`
}

type GetCustomersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Simple string `xml:"simple,omitempty"`

	Search string `xml:"search,omitempty"`
}

type GetCustomersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CustomerData `xml:"data,omitempty"`
}

type DeleteOneCustomerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCustomerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Note string `xml:"note,omitempty"`
}

type DeleteOneCustomerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCustomerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetCustomerPrefsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerPrefsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// preference name; may be '*' to list all
	Name string `xml:"name,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`
}

type GetCustomerPrefsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerPrefsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CustomerPrefData `xml:"data,omitempty"`
}

type SetCustomerPrefsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerPrefsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Prefs []*CustomerPrefData `xml:"prefs,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`
}

type SetCustomerPrefsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerPrefsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetCustomerIPACLRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerIPACLRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// customer name or ID to see ACLs for, defaults to current customer
	Customer_name string `xml:"customer_name,omitempty"`

	// scope of the ACL to retrieve
	Scope string `xml:"scope,omitempty"`
}

type GetCustomerIPACLResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerIPACLResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CustomerIPACL `xml:"data,omitempty"`
}

type SetCustomerIPACLRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerIPACLRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// customer name or ID to set ACLs for, defaults to current customer
	Customer_name string `xml:"customer_name,omitempty"`

	Acl *CustomerIPACL `xml:"acl,omitempty"`
}

type SetCustomerIPACLResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetCustomerIPACLResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CustomerIPACL `xml:"data,omitempty"`
}

type CreateCustomerOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCustomerOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of customer
	Customer_name string `xml:"customer_name,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// tenant id
	Tenant string `xml:"tenant,omitempty"`
}

type CreateCustomerOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCustomerOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerOracleMetadataData `xml:"data,omitempty"`
}

type UpdateCustomerOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCustomerOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of customer
	Customer_name string `xml:"customer_name,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// tenant id
	Tenant string `xml:"tenant,omitempty"`
}

type UpdateCustomerOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCustomerOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerOracleMetadataData `xml:"data,omitempty"`
}

type GetCustomerOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of customer
	Customer_name string `xml:"customer_name,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// tenant id
	Tenant string `xml:"tenant,omitempty"`
}

type GetCustomerOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCustomerOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CustomerOracleMetadataData `xml:"data,omitempty"`
}

type DeleteCustomerOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCustomerOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of customer
	Customer_name string `xml:"customer_name,omitempty"`
}

type DeleteCustomerOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteCustomerOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CustomerOracleMetadataData `xml:"data,omitempty"`
}

type CreateZoneOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateZoneOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// public_id
	Public_id string `xml:"public_id,omitempty"`
}

type CreateZoneOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateZoneOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneOracleMetadataData `xml:"data,omitempty"`
}

type UpdateZoneOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateZoneOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// public_id
	Public_id string `xml:"public_id,omitempty"`
}

type UpdateZoneOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateZoneOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneOracleMetadataData `xml:"data,omitempty"`
}

type GetZoneOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// compartment id
	Compartment_id string `xml:"compartment_id,omitempty"`

	// public id
	Public_id string `xml:"public_id,omitempty"`
}

type GetZoneOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneOracleMetadataData `xml:"data,omitempty"`
}

type DeleteZoneOracleMetadataRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteZoneOracleMetadataRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type DeleteZoneOracleMetadataResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteZoneOracleMetadataResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneOracleMetadataData `xml:"data,omitempty"`
}

type CreateDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// an IP address, either v4 or v6
	Address string `xml:"address,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type GetOneDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type GetDDNSsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDDNSsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetDDNSsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDDNSsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DDNSData `xml:"data,omitempty"`
}

type UpdateDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// an IP address, either v4 or v6
	Address string `xml:"address,omitempty"`
}

type UpdateDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type DeleteOneDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`
}

type ActivateDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type DeactivateDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`
}

type DeactivateDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type ResetDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ResetDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`
}

type ResetDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ResetDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type GetUpdateUserPasswordRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUpdateUserPasswordRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	User_name string `xml:"user_name,omitempty"`
}

type GetUpdateUserPasswordResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetUpdateUserPasswordResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UpdateUserPasswordData `xml:"data,omitempty"`
}

type CreateDDNSHostRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDDNSHostRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// an IP address, either v4 or v6
	Address string `xml:"address,omitempty"`

	// either A or AAAA
	Record_type string `xml:"record_type,omitempty"`

	// name of update user
	User string `xml:"user,omitempty"`
}

type CreateDDNSHostResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDDNSHostResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSHostData `xml:"data,omitempty"`
}

type CreateUpdateUserRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateUpdateUserRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Password string `xml:"password,omitempty"`
}

type CreateUpdateUserResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateUpdateUserResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *UpdateUser `xml:"data,omitempty"`
}

type AddDDNSRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDDNSRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	Record_id int64 `xml:"record_id,omitempty"`
}

type AddDDNSResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ AddDDNSResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DDNSData `xml:"data,omitempty"`
}

type CreateFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// normally served address
	Address string `xml:"address,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip' or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// address or CNAME to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type GetOneFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type GetFailoversRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetFailoversRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetFailoversResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetFailoversResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*FailoverData `xml:"data,omitempty"`
}

type UpdateFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// normally served address
	Address string `xml:"address,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip' or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// address or CNAME to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`
}

type UpdateFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type DeleteOneFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type ActivateFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type DeactivateFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeactivateFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type RecoverFailoverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverFailoverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type RecoverFailoverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverFailoverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *FailoverData `xml:"data,omitempty"`
}

type CreateLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// pool of IP addresses to balance
	Pool []*LoadBalanceAddress `xml:"pool,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type GetOneLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type GetLoadBalancesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLoadBalancesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetLoadBalancesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLoadBalancesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*LoadBalanceData `xml:"data,omitempty"`
}

type UpdateLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// pool of IP addresses to balance
	Pool []*LoadBalanceAddress `xml:"pool,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`
}

type UpdateLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type DeleteOneLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type ActivateLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type DeactivateLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeactivateLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type RecoverLoadBalanceRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverLoadBalanceRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type RecoverLoadBalanceResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverLoadBalanceResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type RecoverLoadBalanceIPRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverLoadBalanceIPRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Address string `xml:"address,omitempty"`
}

type RecoverLoadBalanceIPResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverLoadBalanceIPResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalanceData `xml:"data,omitempty"`
}

type CreateLoadBalancePoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLoadBalancePoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type CreateLoadBalancePoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateLoadBalancePoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalancePoolEntry `xml:"data,omitempty"`
}

type UpdateLoadBalancePoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLoadBalancePoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// The IP of the pool entry to update
	Address string `xml:"address,omitempty"`

	// If specified, the new IP address for this entry
	New_address string `xml:"new_address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type UpdateLoadBalancePoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateLoadBalancePoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalancePoolEntry `xml:"data,omitempty"`
}

type GetOneLoadBalancePoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLoadBalancePoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// The IP of the pool entry to get
	Address string `xml:"address,omitempty"`
}

type GetOneLoadBalancePoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneLoadBalancePoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *LoadBalancePoolEntry `xml:"data,omitempty"`
}

type GetLoadBalancePoolEntriesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLoadBalancePoolEntriesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetLoadBalancePoolEntriesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetLoadBalancePoolEntriesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*LoadBalancePoolEntry `xml:"data,omitempty"`
}

type DeleteOneLoadBalancePoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLoadBalancePoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// The IP of the pool entry to delete
	Address string `xml:"address,omitempty"`
}

type DeleteOneLoadBalancePoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneLoadBalancePoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// per-region addresses and configuration
	Region []*GSLBRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type GetOneGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type GetGSLBsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetGSLBsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*GSLBData `xml:"data,omitempty"`
}

type UpdateGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// per-region addresses and configuration
	Region []*GSLBRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`
}

type UpdateGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type DeleteOneGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type ActivateGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type DeactivateGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeactivateGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type RecoverGSLBRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverGSLBRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type RecoverGSLBResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverGSLBResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type RecoverGSLBIPRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverGSLBIPRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Address string `xml:"address,omitempty"`
}

type RecoverGSLBIPResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverGSLBIPResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBData `xml:"data,omitempty"`
}

type CreateGSLBRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// number of 'ok' addresses before region fails over
	Min_healthy int32 `xml:"min_healthy,omitempty"`

	// pool of IP addresses to balance
	Pool []*GSLBAddress `xml:"pool,omitempty"`
}

type CreateGSLBRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionData `xml:"data,omitempty"`
}

type GetOneGSLBRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneGSLBRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionData `xml:"data,omitempty"`
}

type GetGSLBRegionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBRegionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetGSLBRegionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBRegionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*GSLBRegionData `xml:"data,omitempty"`
}

type UpdateGSLBRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'ip', 'global', or 'cname'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' or 'cname', what to serve on failover
	Failover_data string `xml:"failover_data,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// number of 'ok' addresses before region fails over
	Min_healthy int32 `xml:"min_healthy,omitempty"`

	// pool of IP addresses to balance
	Pool []*GSLBAddress `xml:"pool,omitempty"`
}

type UpdateGSLBRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionData `xml:"data,omitempty"`
}

type DeleteOneGSLBRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneGSLBRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateGSLBRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// an IP address or FQDN to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type CreateGSLBRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateGSLBRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionPoolEntry `xml:"data,omitempty"`
}

type UpdateGSLBRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP address or FQDN of the pool entry to update
	Address string `xml:"address,omitempty"`

	// If specified, the new IP address for this entry
	New_address string `xml:"new_address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type UpdateGSLBRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateGSLBRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionPoolEntry `xml:"data,omitempty"`
}

type GetOneGSLBRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP address or FQDN of the pool entry to get
	Address string `xml:"address,omitempty"`
}

type GetOneGSLBRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneGSLBRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *GSLBRegionPoolEntry `xml:"data,omitempty"`
}

type GetGSLBRegionPoolEntriesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBRegionPoolEntriesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`
}

type GetGSLBRegionPoolEntriesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetGSLBRegionPoolEntriesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*GSLBRegionPoolEntry `xml:"data,omitempty"`
}

type DeleteOneGSLBRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP of the pool entry to delete
	Address string `xml:"address,omitempty"`
}

type DeleteOneGSLBRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneGSLBRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// per-region addresses and configuration
	Region []*RTTMRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// for custom syslog messages
	Syslog_rttm_fmt string `xml:"syslog_rttm_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// details about performance monitoring
	Performance_monitor *MonitorData `xml:"performance_monitor,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type GetOneRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type GetRTTMsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetRTTMsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RTTMData `xml:"data,omitempty"`
}

type UpdateRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// per-region addresses and configuration
	Region []*RTTMRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// restore normal address automatically (Y)
	Auto_recover string `xml:"auto_recover,omitempty"`

	// The number of consecutive monitoring intervals to delay before placing an IP address back in service
	Recovery_delay int32 `xml:"recovery_delay,omitempty"`

	// contact that gets status notification
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	// The IP or hostname of a syslog server to send monitor events to
	Syslog_server string `xml:"syslog_server,omitempty"`

	// The port of the syslog server. Defaults to 514 if not present
	Syslog_port string `xml:"syslog_port,omitempty"`

	// The syslog ident to use. Defaults to 'dynect'
	Syslog_ident string `xml:"syslog_ident,omitempty"`

	// The syslog facility to use. Defaults to 'daemon'
	Syslog_facility string `xml:"syslog_facility,omitempty"`

	// When to deliver syslog message; 'change' or 'all'
	Syslog_delivery string `xml:"syslog_delivery,omitempty"`

	// for custom syslog messages
	Syslog_probe_fmt string `xml:"syslog_probe_fmt,omitempty"`

	// for custom syslog messages
	Syslog_status_fmt string `xml:"syslog_status_fmt,omitempty"`

	// for custom syslog messages
	Syslog_rttm_fmt string `xml:"syslog_rttm_fmt,omitempty"`

	// details about monitoring
	Monitor *MonitorData `xml:"monitor,omitempty"`

	// details about performance monitoring
	Performance_monitor *MonitorData `xml:"performance_monitor,omitempty"`
}

type UpdateRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type DeleteOneRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type ActivateRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type DeactivateRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeactivateRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type RecoverRTTMRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverRTTMRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type RecoverRTTMResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverRTTMResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type RecoverRTTMIPRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverRTTMIPRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Address string `xml:"address,omitempty"`
}

type RecoverRTTMIPResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverRTTMIPResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMData `xml:"data,omitempty"`
}

type GetRTTMLogsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMLogsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// The timestamp indicating the beginning of the period to report on
	Start_ts int32 `xml:"start_ts,omitempty"`

	// The timestamp indicating the end of the period to report on
	End_ts int32 `xml:"end_ts,omitempty"`
}

type GetRTTMLogsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMLogsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RTTMLogData `xml:"data,omitempty"`
}

type GetRTTMRRSetsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRRSetsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// The timestamp indicating the period to report on
	Ts int32 `xml:"ts,omitempty"`
}

type GetRTTMRRSetsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRRSetsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RTTMLogData `xml:"data,omitempty"`
}

type CreateRTTMRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// 'Y' or 'N', if 'Y', region will filled in with global settings
	Autopopulate string `xml:"autopopulate,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// pool_count, number of addresses to be included in the serve pool
	Ep int32 `xml:"ep,omitempty"`

	// 'ip', 'global', 'region', default 'global'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' mode, address to serve on failover. For 'region' mode, region_code of the region to failover to.
	Failover_data string `xml:"failover_data,omitempty"`

	// failover_count, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Apmc int32 `xml:"apmc,omitempty"`

	// failover_count2, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Epmc int32 `xml:"epmc,omitempty"`

	// pool of IP addresses to balance
	Pool []*RTTMAddress `xml:"pool,omitempty"`
}

type CreateRTTMRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionData `xml:"data,omitempty"`
}

type GetOneRTTMRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneRTTMRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionData `xml:"data,omitempty"`
}

type GetRTTMRegionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRegionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetRTTMRegionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRegionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RTTMRegionData `xml:"data,omitempty"`
}

type UpdateRTTMRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'Y' or 'N', if 'Y', region will filled in with global settings
	Autopopulate string `xml:"autopopulate,omitempty"`

	// number of addresses in each DNS response
	Serve_count int32 `xml:"serve_count,omitempty"`

	// pool_count, number of addresses to be included in the serve pool
	Ep int32 `xml:"ep,omitempty"`

	// 'ip', 'global', 'region', default 'global'
	Failover_mode string `xml:"failover_mode,omitempty"`

	// for 'ip' mode, address to serve on failover. For 'region' mode, region_code of the region to failover to.
	Failover_data string `xml:"failover_data,omitempty"`

	// failover_count, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Apmc int32 `xml:"apmc,omitempty"`

	// failover_count2, number of addresses that must be 'ok', otherwise a region is considered 'failover'
	Epmc int32 `xml:"epmc,omitempty"`

	// pool of IP addresses to balance
	Pool []*RTTMAddress `xml:"pool,omitempty"`
}

type UpdateRTTMRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionData `xml:"data,omitempty"`
}

type DeleteOneRTTMRegionRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMRegionRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Region_code string `xml:"region_code,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneRTTMRegionResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMRegionResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateRTTMRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// an IP address to monitor and publish
	Address string `xml:"address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type CreateRTTMRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateRTTMRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionPoolEntry `xml:"data,omitempty"`
}

type UpdateRTTMRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP of the pool entry to update
	Address string `xml:"address,omitempty"`

	// If specified, the new IP address for this entry
	New_address string `xml:"new_address,omitempty"`

	// a human-readable label
	Label string `xml:"label,omitempty"`

	// how often this is served relative to others in pool
	Weight int32 `xml:"weight,omitempty"`

	// how this address reponds to monitoring: obey,remove,always,no
	Serve_mode string `xml:"serve_mode,omitempty"`
}

type UpdateRTTMRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateRTTMRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionPoolEntry `xml:"data,omitempty"`
}

type GetOneRTTMRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP of the pool entry to get
	Address string `xml:"address,omitempty"`
}

type GetOneRTTMRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneRTTMRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *RTTMRegionPoolEntry `xml:"data,omitempty"`
}

type GetRTTMRegionPoolEntriesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRegionPoolEntriesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`
}

type GetRTTMRegionPoolEntriesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetRTTMRegionPoolEntriesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*RTTMRegionPoolEntry `xml:"data,omitempty"`
}

type DeleteOneRTTMRegionPoolEntryRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMRegionPoolEntryRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`

	// The IP of the pool entry to delete
	Address string `xml:"address,omitempty"`
}

type DeleteOneRTTMRegionPoolEntryResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneRTTMRegionPoolEntryResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateHTTPRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateHTTPRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// URL requests are redirecto to
	Url string `xml:"url,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// should redirected URL include requested URL
	Keep_uri string `xml:"keep_uri,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateHTTPRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateHTTPRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *HTTPRedirectData `xml:"data,omitempty"`
}

type GetOneHTTPRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneHTTPRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneHTTPRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneHTTPRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *HTTPRedirectData `xml:"data,omitempty"`
}

type GetHTTPRedirectsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetHTTPRedirectsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetHTTPRedirectsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetHTTPRedirectsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*HTTPRedirectData `xml:"data,omitempty"`
}

type UpdateHTTPRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateHTTPRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// URL requests are redirecto to
	Url string `xml:"url,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// should redirected URL include requested URL
	Keep_uri string `xml:"keep_uri,omitempty"`
}

type UpdateHTTPRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateHTTPRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *HTTPRedirectData `xml:"data,omitempty"`
}

type DeleteOneHTTPRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneHTTPRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	Publish string `xml:"publish,omitempty"`
}

type DeleteOneHTTPRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneHTTPRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type CreateAdvRedirectRuleRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAdvRedirectRuleRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// host portion of URL to match
	Host_prefix string `xml:"host_prefix,omitempty"`

	// path portion of URL to match
	Path string `xml:"path,omitempty"`

	// replacement pattern
	Url_pattern string `xml:"url_pattern,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// Public ID of next AdvRedirect rule to be processed. (default to end of list)
	Next_public_id string `xml:"next_public_id,omitempty"`
}

type CreateAdvRedirectRuleResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAdvRedirectRuleResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectRuleData `xml:"data,omitempty"`
}

type UpdateAdvRedirectRuleRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAdvRedirectRuleRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// public_id of Rule
	Public_id string `xml:"public_id,omitempty"`

	// either '301' (temporary) or '302' (permanent)
	Code string `xml:"code,omitempty"`

	// host portion of URL to match
	Host_prefix string `xml:"host_prefix,omitempty"`

	// path portion of URL to match
	Path string `xml:"path,omitempty"`

	// replacement pattern
	Url_pattern string `xml:"url_pattern,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// Public ID of next AdvRedirect rule to be processed. (default to end of list)
	Next_public_id string `xml:"next_public_id,omitempty"`
}

type UpdateAdvRedirectRuleResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAdvRedirectRuleResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectRuleData `xml:"data,omitempty"`
}

type GetOneAdvRedirectRuleRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAdvRedirectRuleRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// public_id of Rule
	Public_id string `xml:"public_id,omitempty"`
}

type GetOneAdvRedirectRuleResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAdvRedirectRuleResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectRuleData `xml:"data,omitempty"`
}

type GetAdvRedirectRulesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAdvRedirectRulesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetAdvRedirectRulesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAdvRedirectRulesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*AdvRedirectRuleData `xml:"data,omitempty"`
}

type DeleteOneAdvRedirectRuleRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAdvRedirectRuleRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// public_id of Rule
	Public_id string `xml:"public_id,omitempty"`
}

type DeleteOneAdvRedirectRuleResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAdvRedirectRuleResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectRuleData `xml:"data,omitempty"`
}

type CreateAdvRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAdvRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// List of AdvRedirectRules
	Rules []*AdvRedirectRuleData `xml:"rules,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateAdvRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateAdvRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectData `xml:"data,omitempty"`
}

type GetOneAdvRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAdvRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneAdvRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneAdvRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectData `xml:"data,omitempty"`
}

type GetAdvRedirectsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAdvRedirectsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	Rules string `xml:"rules,omitempty"`
}

type GetAdvRedirectsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetAdvRedirectsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*AdvRedirectData `xml:"data,omitempty"`
}

type UpdateAdvRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAdvRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// List of AdvRedirectRules
	Rules []*AdvRedirectRuleData `xml:"rules,omitempty"`
}

type UpdateAdvRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateAdvRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectData `xml:"data,omitempty"`
}

type DeleteOneAdvRedirectRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAdvRedirectRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// 'Y'/'N', default 'Y'
	Active string `xml:"active,omitempty"`

	// List of AdvRedirectRules
	Rules []*AdvRedirectRuleData `xml:"rules,omitempty"`
}

type DeleteOneAdvRedirectResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneAdvRedirectResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *AdvRedirectData `xml:"data,omitempty"`
}

type CreateCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// per region hostnames and configuration
	Region []*CDNRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNManagerData `xml:"data,omitempty"`
}

type GetOneCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNManagerData `xml:"data,omitempty"`
}

type GetCDNManagersRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDNManagersRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetCDNManagersResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetCDNManagersResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*CDNManagerData `xml:"data,omitempty"`
}

type UpdateCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// per region hostnames and configuration
	Region []*CDNRegion `xml:"region,omitempty"`

	// TTL (time-to-live)
	Ttl int32 `xml:"ttl,omitempty"`
}

type UpdateCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNManagerData `xml:"data,omitempty"`
}

type DeleteOneCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type FailoverCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ FailoverCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`
}

type FailoverCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ FailoverCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNManagerData `xml:"data,omitempty"`
}

type RecoverCDNManagerRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverCDNManagerRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Id int64 `xml:"id,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`

	// 'global' or specific region code: 'US West', 'US Central', 'US East', 'EU West', 'EU Central', 'Asia',
	Region_code string `xml:"region_code,omitempty"`
}

type RecoverCDNManagerResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RecoverCDNManagerResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *CDNManagerData `xml:"data,omitempty"`
}

type GetNodeListRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNodeListRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type GetNodeListResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetNodeListResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []string `xml:"data,omitempty"`
}

type PublishZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PublishZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Notes string `xml:"notes,omitempty"`
}

type PublishZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PublishZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *PublishZoneData `xml:"data,omitempty"`
}

type PruneZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PruneZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Fqdn string `xml:"fqdn,omitempty"`
}

type PruneZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ PruneZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneData `xml:"data,omitempty"`
}

type FreezeZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ FreezeZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type FreezeZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ FreezeZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ThawZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ThawZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type ThawZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ThawZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type RestoreZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RestoreZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type RestoreZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ RestoreZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type BlockZoneRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ BlockZoneRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type BlockZoneResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ BlockZoneResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DeleteZoneChangesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteZoneChangesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type DeleteZoneChangesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteZoneChangesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetZoneChangesetRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneChangesetRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type GetZoneChangesetResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneChangesetResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneChangeData `xml:"data,omitempty"`
}

type GetZoneNotesRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneNotesRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Offset int32 `xml:"offset,omitempty"`
}

type GetZoneNotesResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneNotesResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneNoteData `xml:"data,omitempty"`
}

type UploadZoneFileRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UploadZoneFileRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	File string `xml:"file,omitempty"`

	Create string `xml:"create,omitempty"`
}

type UploadZoneFileResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UploadZoneFileResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TaskIDData `xml:"data,omitempty"`
}

type TransferZoneInRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ TransferZoneInRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Master_ip string `xml:"master_ip,omitempty"`
}

type TransferZoneInResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ TransferZoneInResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TaskIDData `xml:"data,omitempty"`
}

type GetTransferStatusRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTransferStatusRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type GetTransferStatusResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTransferStatusResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ZoneTransferStatus `xml:"data,omitempty"`
}

type GetZoneConfigOptionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneConfigOptionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type GetZoneConfigOptionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetZoneConfigOptionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneConfigOptionData `xml:"data,omitempty"`
}

type SetZoneConfigOptionsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetZoneConfigOptionsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	Option []*ZoneConfigOptionData `xml:"option,omitempty"`
}

type SetZoneConfigOptionsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ SetZoneConfigOptionsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ZoneConfigOptionData `xml:"data,omitempty"`
}

type CreateIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// A, Dynamic_A, AAAA, Dynamic_AAAA
	Record_types []string `xml:"record_types,omitempty"`

	// List of hostnames to watch for records
	Hosts []string `xml:"hosts,omitempty"`

	// 'match', 'default', or a valid ttl
	Ttl string `xml:"ttl,omitempty"`

	// Mask that records should match
	Netmask string `xml:"netmask,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type CreateIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPTrackData `xml:"data,omitempty"`
}

type GetOneIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type GetOneIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPTrackData `xml:"data,omitempty"`
}

type GetIPTracksRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetIPTracksRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`
}

type GetIPTracksResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetIPTracksResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*IPTrackData `xml:"data,omitempty"`
}

type UpdateIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`

	// A, Dynamic_A, AAAA, Dynamic_AAAA
	Record_types []string `xml:"record_types,omitempty"`

	// List of hostnames to watch for records
	Hosts []string `xml:"hosts,omitempty"`

	// 'match', 'default', or a valid ttl
	Ttl string `xml:"ttl,omitempty"`

	// Mask that records should match
	Netmask string `xml:"netmask,omitempty"`
}

type UpdateIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPTrackData `xml:"data,omitempty"`
}

type DeleteOneIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeleteOneIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type ActivateIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPTrackData `xml:"data,omitempty"`
}

type DeactivateIPTrackRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateIPTrackRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Iptrack_id int64 `xml:"iptrack_id,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// hostname
	Fqdn string `xml:"fqdn,omitempty"`
}

type DeactivateIPTrackResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateIPTrackResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *IPTrackData `xml:"data,omitempty"`
}

type CreateDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	// ,      contact that gets key notifications
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	Keys []*DNSSECKey `xml:"keys,omitempty"`
}

type CreateDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSSECData `xml:"data,omitempty"`
}

type GetOneDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type GetOneDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSSECData `xml:"data,omitempty"`
}

type GetDNSSECsRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSSECsRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetDNSSECsResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSSECsResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNSSECData `xml:"data,omitempty"`
}

type UpdateDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`

	// ,      contact that gets key notifications
	Contact_nickname string `xml:"contact_nickname,omitempty"`

	// when notifications are sent
	Notify_events string `xml:"notify_events,omitempty"`

	Keys []*DNSSECKey `xml:"keys,omitempty"`
}

type UpdateDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSSECData `xml:"data,omitempty"`
}

type DeleteOneDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type DeleteOneDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}

type ActivateDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type ActivateDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ ActivateDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSSECData `xml:"data,omitempty"`
}

type DeactivateDNSSECRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateDNSSECRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Zone string `xml:"zone,omitempty"`
}

type DeactivateDNSSECResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeactivateDNSSECResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *DNSSECData `xml:"data,omitempty"`
}

type GetDNSSECTimelineRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSSECTimelineRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// name of zone
	Zone string `xml:"zone,omitempty"`

	// an epoch time, or 'now'
	Start_ts string `xml:"start_ts,omitempty"`

	// an epoch time, or 'now'
	End_ts string `xml:"end_ts,omitempty"`
}

type GetDNSSECTimelineResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetDNSSECTimelineResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*DNSSECTimelineEvent `xml:"data,omitempty"`
}

type GetTasksRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTasksRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Name string `xml:"name,omitempty"`

	Customer_name string `xml:"customer_name,omitempty"`

	Zone_name string `xml:"zone_name,omitempty"`

	Status string `xml:"status,omitempty"`
}

type GetTasksResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetTasksResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*TaskData `xml:"data,omitempty"`
}

type GetOneTaskRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTaskRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Task_id string `xml:"task_id,omitempty"`
}

type GetOneTaskResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneTaskResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TaskData `xml:"data,omitempty"`
}

type CancelTaskRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CancelTaskRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	Task_id string `xml:"task_id,omitempty"`
}

type CancelTaskResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CancelTaskResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *TaskData `xml:"data,omitempty"`
}

type CreateExtNameserverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateExtNameserverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// can be empty or 'default'
	Zone string `xml:"zone,omitempty"`

	// Y/N - does this block requests or add them
	Deny string `xml:"deny,omitempty"`

	Hosts []*ExtNSEntry `xml:"hosts,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Active string `xml:"active,omitempty"`
}

type CreateExtNameserverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ CreateExtNameserverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ExtNameserverData `xml:"data,omitempty"`
}

type GetOneExtNameserverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneExtNameserverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// can be empty or 'default'
	Zone string `xml:"zone,omitempty"`
}

type GetOneExtNameserverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetOneExtNameserverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ExtNameserverData `xml:"data,omitempty"`
}

type GetExtNameserversRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetExtNameserversRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`
}

type GetExtNameserversResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ GetExtNameserversResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data []*ExtNameserverData `xml:"data,omitempty"`
}

type UpdateExtNameserverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateExtNameserverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// can be empty or 'default'
	Zone string `xml:"zone,omitempty"`

	// Y/N - does this block requests or add them
	Deny string `xml:"deny,omitempty"`

	Hosts []*ExtNSEntry `xml:"hosts,omitempty"`

	Tsig_key_name string `xml:"tsig_key_name,omitempty"`

	Active string `xml:"active,omitempty"`
}

type UpdateExtNameserverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ UpdateExtNameserverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data *ExtNameserverData `xml:"data,omitempty"`
}

type DeleteOneExtNameserverRequestType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneExtNameserverRequest"`

	Token string `xml:"token,omitempty"`

	Fault_incompat int32 `xml:"fault_incompat,omitempty"`

	// can be empty or 'default'
	Zone string `xml:"zone,omitempty"`
}

type DeleteOneExtNameserverResponseType struct {
	XMLName xml.Name `xml:"https://api2.dynect.net/wsdl/3.7.11/Dynect/ DeleteOneExtNameserverResponse"`

	Job_id string `xml:"job_id,omitempty"`

	Status string `xml:"status,omitempty"`

	Msgs []*Messages `xml:"msgs,omitempty"`

	Data string `xml:"data,omitempty"`
}
