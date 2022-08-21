package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/javadmohebbi/gobdgz"
)

var (
	gz, gzDst *gobdgz.GravityZoneAPI
	rq, rqDst gobdgz.Request

	srcServer *string
	dstServer *string
	srcPolicy *string
	dstPolicy *string
	srcAPI    *string
	dstAPI    *string

	debug *bool
)

func main() {

	// export policy
	rq.Method = "exportPolicies"
	rq.Params = map[string]interface{}{
		"policyNames": []string{
			*srcPolicy,
		},
	}
	gz.Policy.SetRequest(rq)
	resp, err := gz.Policy.ExportPolicies()
	if err != nil {
		log.Fatal(err)
	}

	doTheImport(resp)
	// strResp, err := json.Marshal(resp.Result.Items[0])
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = os.WriteFile("/tmp/json.json", strResp, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }

}

func doTheImport(resp gobdgz.ExportPolicyResponse) {

	// import policy
	if *dstPolicy == "" {
		*dstPolicy = *srcPolicy
	}
	resp.Result.Items[0].Name = *dstPolicy
	rqDst.Method = "importPolicies"
	// type record struct {
	// 	Name       string            `json:"name"`
	// 	UISettings gobdgz.UISettings `json:"uiSettings"`
	// 	Service    int               `json:"service"`
	// }
	// var records []record
	// var rec record

	// rec.Name = *dstPolicy
	// rec.UISettings = resp.Result.Items[0].UISettings

	// rec.Service = resp.Result.Items[0].Service
	// records = append(records, rec)

	// fmt.Println(resp.Result.Items)

	rqDst.Params = map[string]interface{}{
		"records":       resp.Result.Items,
		"compatVersion": resp.Result.CompatVersion,
	}
	gzDst.Policy.SetRequest(rqDst)
	respDst, err := gzDst.Policy.ImportPolicies()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(respDst.Result.Success)
}

func init() {

	// initialize flags
	initFlags()

	init_src_server()
	init_dst_server()

}

func init_src_server() {
	// create new instance of GravityZone API
	gz = &gobdgz.GravityZoneAPI{
		BaseURL: fmt.Sprintf("https://%s/api/v1.0/jsonrpc/", *srcServer),
	}

	// Create Default Request struct
	rq = gobdgz.Request{
		Debug:      *debug,
		APIKey:     *srcAPI,
		JSONRPC:    "2.0",
		URL:        gz.BaseURL + "/policies",
		HttpMethod: "POST",
		ID:         "98409cc1-93cc-415a-9f77-1d4f681000b3",
	}
}

func init_dst_server() {

	// create new instance of GravityZone API
	gzDst = &gobdgz.GravityZoneAPI{
		BaseURL: fmt.Sprintf("https://%s/api/v1.0/jsonrpc/", *dstServer),
	}

	// Create Default Request struct
	rqDst = gobdgz.Request{
		Debug:      *debug,
		APIKey:     *dstAPI,
		JSONRPC:    "2.0",
		URL:        gzDst.BaseURL + "/policies",
		HttpMethod: "POST",
		ID:         "98409cc1-93cc-415a-9f77-1d4f681000b3",
	}
}

func initFlags() {
	srcServer = flag.String("src-server", "", "Source Server (FQDN / IP)")
	dstServer = flag.String("dst-server", "", "Destination Server (FQDN / IP)")
	srcAPI = flag.String("src-server-api", "", "Source Server API key")
	dstAPI = flag.String("dst-server-api", "", "Destination Server API key")
	srcPolicy = flag.String("src-policy-name", "", "Source Policy Name")
	dstPolicy = flag.String("dst-policy-name", "", "Source Policy Name")

	debug = flag.Bool("debug", false, "This parameter will enable debugging mode")

	flag.Parse()

	if *srcServer == "" {
		fmt.Println("-src-server must be provided")
		flag.PrintDefaults()
		os.Exit(127)
	}
	if *dstServer == "" {
		fmt.Println("-dst-server must be provided")
		flag.PrintDefaults()
		os.Exit(127)
	}
	if *srcAPI == "" {
		fmt.Println("-src-server-api must be provided")
		flag.PrintDefaults()
		os.Exit(127)
	}
	if *dstAPI == "" {
		fmt.Println("-dst-server-api must be provided")
		flag.PrintDefaults()
		os.Exit(127)
	}
	if *srcPolicy == "" {
		fmt.Println("-src-policy-name must be provided")
		flag.PrintDefaults()
		os.Exit(127)
	}

}
