package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NexClipper/logger"
	"github.com/gin-gonic/gin"
	"github.com/ralfyang/pkg/klevr/manager/model"
)

type apiDef struct {
	method   string
	uri      string
	function func(*gin.Context)
}

var apiSlice []apiDef

func (api *API) initAPI() {
	logger.Debug("API InitLegacy - init URI")
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/ackprimary", api.ackprimary})
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/hostsinfo", api.hostsinfo})
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/primaryinfo", api.primaryinfo})
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/hostsmgt", api.hostsmgt})
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/job/:JOB/ticket/:TICKET/:MSG", api.callback})
	apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/hostname/:H/hostinfo", api.hostinfo})
	// apiSlice = append(apiSlice, apiDef{"any", "/group/:G/user/:U/zone/:Z/platform/:P/hostname/:H/:I/:TTL/:MLO", api.alivetime})
	apiSlice = append(apiSlice, apiDef{"any", "/systems/platform_types/:P", api.initAgent})
	apiSlice = append(apiSlice, apiDef{"any", "/groups/:G/users/:U/zones/:Z/platforms/:P/aliveagent", api.statusReciever})
}

// InitLegacy initialize legacy API
func (api *API) InitLegacy(legacy *gin.RouterGroup) {
	logger.Debug("API InitLegacy")

	api.initAPI()

	for _, def := range apiSlice {
		switch def.method {
		case "any":
			legacy.Any(def.uri, def.function)
		case "get":
			legacy.GET(def.uri, def.function)
		case "post":
			legacy.POST(def.uri, def.function)
		case "put":
			legacy.PUT(def.uri, def.function)
		case "delete":
			legacy.DELETE(def.uri, def.function)
		case "patch":
			legacy.PATCH(def.uri, def.function)
		}
	}
}

func (api *API) ackprimary(c *gin.Context) {
	group, _ := strconv.ParseUint(c.Param("G"), 10, 64)
	user, _ := strconv.ParseUint(c.Param("U"), 10, 64)
	zone := c.Param("Z")
	platform := c.Param("P")

	c.JSON(200, gin.H{
		"accessTime": api.PutPrimaryAck(group, user, zone, platform, fmt.Sprint(time.Now().Unix())),
	})
	// GetHost()
}

func (api *API) hostsinfo(c *gin.Context) {
	c.String(200, "test2")
}

func (api *API) primaryinfo(c *gin.Context) {

}

func (api *API) hostsmgt(c *gin.Context) {

}

func (api *API) callback(c *gin.Context) {

}

func (api *API) alivetime(c *gin.Context) {

}

func (api *API) hostinfo(c *gin.Context) {

}

func (api *API) initAgent(c *gin.Context) {

}

func (api *API) statusReciever(c *gin.Context) {

}

// GetProvisionScript For custom scripts when the agent download & install
func GetProvisionScript() string {
	// Http_body_buffer := communicator.Get_http(ConsulURL+"/v1/kv/klevr/form?raw=1", API_key_string)
	// if len(string(Http_body_buffer)) == 0 {
	// 	/// Set Script for instruction
	// 	uri := "/v1/kv/klevr/form"
	// 	data := "bash -s 'echo \"hello world\"'" /// Temporary use
	// 	communicator.Put_http(ConsulURL+uri, data, API_key_string)
	// 	/// Read again
	// 	API_provision_script = communicator.Get_http(ConsulURL+"/v1/kv/klevr/form?raw=1", API_key_string)
	// } else {
	// 	API_provision_script = communicator.Get_http(ConsulURL+"/v1/kv/klevr/form?raw=1", API_key_string)
	// }
	// return API_provision_script
	return ""
}

// LandingPage Default Landing page for http
func LandingPage(w http.ResponseWriter, r *http.Request) {
	// // w.Write([]byte("<a href='https://bit.ly/startdocker' target='blank'>Download Klever agent</a>"))
	// GetProvisionScript()
	// w.Write([]byte("curl -sL " + AgentDownload + " | " + API_provision_script))
}

// SetParam Get Config variable when the webconsole start
func SetParam() string {
	// //Parsing by Flag
	// port := flag.String("port", Service_port, "Set port number for Service")
	// api_server := flag.String("apiserver", ConsulURL, "Set API Server URI for comunication")
	// flag.Parse()
	// Service_port = *port
	// ConsulURL = *api_server
	// return Service_port
	return ""
}

// GetPrimary company user zone platform
//%s/+group+"/users/"+user+"/zones/"+zone+/+group+"\/groups/"+group+"/users/"+user+"/zones/"+zone+"\/zones/"+zone+/g
/// Get Primary server infomation for secondary agent control
func GetPrimary(group, user, zone, platform string) string {
	// Primary_info = communicator.Get_http(ConsulURL+"/v1/kv/klevr/groups/"+group+"/users/"+user+"/zones/"+zone+"/platforms/"+platform+"/primarys?raw=1", API_key_string)
	// if len(Primary_info) == 0 {
	// 	Primary_info = "Not yet"
	// }
	// return Primary_info

	return ""
}

// LogRequest ..
func LogRequest(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s | url: %s", r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// GetHost Get Hostlist
func (api *API) GetHost(group uint64, user uint64, zone, platform, priyes string) string {
	api.DB.LogMode(true)

	api.DB.Joins

	return ""
}

// GetInfoPrimary ..
func GetInfoPrimary(group, user, zone, platform string) {
	/// initial primary info
	// GetHost(group, user, zone, platform, "")
	GetPrimary(group, user, zone, platform)
}

// PutPlatformInit ..
func PutPlatformInit(platform, data string) {
	// uri := "/v1/kv/klevr/systems/platform_types/" + platform
	// communicator.Put_http(ConsulURL+uri, data, API_key_string)
}

// PutPrimaryAck ..
func (api *API) PutPrimaryAck(group uint64, user uint64, zone, platform, ack string) time.Time {
	logger.Debug(fmt.Sprintf("group : %d, user : %d", group, user))

	var ma = &model.PrimaryAgents{
		GroupId:        group,
		AgentId:        user,
		LastAccessTime: time.Now().UTC(),
	}

	api.DB.LogMode(true)

	api.DB.Where(&model.PrimaryAgents{
		GroupId: group,
		AgentId: user,
	}).FirstOrCreate(&ma)

	logger.Debug(ma)

	accessTime := time.Now().UTC()

	api.DB.Model(&ma).Updates(model.PrimaryAgents{LastAccessTime: accessTime})

	return accessTime
}

// HostpoolMgt Old hostlist purge
func HostpoolMgt(group, user, zone, platform string) string {
	// /// Define variables
	// var arr []string
	// var queue, target_key string
	// Host_purge_result = "\n"

	// /// Get Hostlist with Keys
	// dataJson := communicator.Get_http(ConsulURL+"/v1/kv/klevr/groups/"+group+"/users/"+user+"/zones/"+zone+"/platforms/"+platform+"/hosts/?keys", API_key_string)
	// _ = json.Unmarshal([]byte(dataJson), &arr)
	// for i := 0; i < len(arr); i++ {
	// 	var target_txt, time_arry []string
	// 	var time_string string
	// 	endpoint := arr[i][strings.LastIndex(arr[i], "/")+1:]
	// 	if endpoint == "health" {
	// 		queue = communicator.Get_http(ConsulURL+"/v1/kv/"+arr[i]+"?raw=1", API_key_string) /// Endpoing value will be "~/health" part of API
	// 		get_data := arr[i]

	// 		/// Get value of each hosts
	// 		target_key = ConsulURL + "/v1/kv/" + get_data
	// 		println("target_key=", target_key) ///////////  Test output
	// 		/// Parsing the Key/value of host_info
	// 		target_txt = strings.Split(string(queue), "&")
	// 		time_arry = strings.Split(target_txt[0], "=")

	// 		/// Parsing the Key/value for Unix Time
	// 		time_string = string(time_arry[1])
	// 		time_parsing, err := strconv.ParseInt(time_string, 10, 64)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		/// Duration check
	// 		tm := time.Unix(time_parsing, 0)
	// 		if time.Since(tm).Hours() > 1 {
	// 			/// Delete old host via API server
	// 			Host_purge_result = Host_purge_result + "Overtime: " + get_data + "\n"
	// 			communicator.Delete_http(ConsulURL+"/v1/kv/"+get_data, API_key_string)
	// 		} else {
	// 			Host_purge_result = Host_purge_result + "It's ok: " + get_data + "\n"
	// 		}
	// 	}
	// }
	// return Host_purge_result

	return ""
}

// ClientReceiver ..
func ClientReceiver(group, user, zone, hostname, hostIP, platform, hostAlive, primaryAlive string) string {
	// uri := "/v1/kv/klevr/groups/" + group + "/users/" + user + "/zones/" + zone + "/platforms/" + platform + "/hosts/" + hostname + "/health"
	// data := "last_check=" + host_alive + "&ip=" + host_ip + "&clientType=" + platform + "&primaryConnection=" + primary_alive
	// communicator.Put_http(ConsulURL+uri, data, API_key_string)
	// Buffer_result = data
	// return Buffer_result

	return ""
}

// PutHostinfo ..
func PutHostinfo(group, user, zone, platform, hostname, body string) string {
	// uri := "/v1/kv/klevr/groups/" + group + "/users/" + user + "/zones/" + zone + "/platforms/" + platform + "/hosts/" + hostname + "/hostinfo"
	// data := body
	// communicator.Put_http(ConsulURL+uri, data, API_key_string)
	// Buffer_result = data
	// return Buffer_result

	return ""
}
