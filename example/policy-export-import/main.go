package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/javadmohebbi/gobdgz"
	"github.com/manifoldco/promptui"
)

type Config struct {
	Src SRCDST `json:"SRC"`
	Dst SRCDST `json:"DST"`
}
type SRCDST struct {
	Server   string   `json:"SERVER"`
	APIKey   string   `json:"API_KEY"`
	Policies []string `json:"POLICIES,omitempty"`
}

var (
	gz, gzDst *gobdgz.GravityZoneAPI
	rq, rqDst gobdgz.Request

	debug bool

	result Config
)

func main() {

	fmt.Println("*** Bitdefender GravityZone Policy Migration Tool ***")
	fmt.Println("Author: https://openintelligence24.com | javad@openintelligence24.com")
	fmt.Println("Online Help: https://github.com/javadmohebbi/gobdgz/tree/master/example/policy-export-import")
	fmt.Println("Github: https://github.com/javadmohebbi/gobdgz")

	fmt.Println("- - - - - - - - - - - - - - - - - - - ")
	fmt.Printf("You are going to migrate (%d) policy/policies \nfrom Source Server (%s) to Destination Server (%s)\n",
		len(result.Src.Policies), result.Src.Server, result.Dst.Server,
	)
	fmt.Println("*** All migrated policies will get '-clone' extention")
	fmt.Println("*** Policies with the same name (YourPolicy-clone) will be OVERWRITTEN in the destination server")
	fmt.Println("- - - - - - - - - - - - - - - - - - - ")
	if yes := yesNo(); !yes {
		fmt.Println("Have a nice day! goodbye ;-)")
		os.Exit(0)
	}

	for _, pol := range result.Src.Policies {

		log.Printf("Migrating policy (%s)...\n", pol)

		// export policy
		rq.Method = "exportPolicies"
		rq.Params = map[string]interface{}{
			"policyNames": []string{
				pol,
			},
		}
		gz.Policy.SetRequest(rq)
		resp, err := gz.Policy.ExportPolicies()
		if err != nil {
			log.Fatal(err)
		}

		// import policy
		if ok := doTheImport(resp, pol); ok {
			log.Printf(".....(%s -> %s-clone) Migrated!\n", pol, pol)
		}
		time.Sleep(1 * time.Second)
	}

	// strResp, err := json.Marshal(resp.Result.Items[0])
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = os.WriteFile("/tmp/json.json", strResp, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }

}

func yesNo() bool {
	prompt := promptui.Select{
		Label: "Do you want to contine (Yes / No)",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func doTheImport(resp gobdgz.ExportPolicyResponse, pol string) bool {

	// import policy

	resp.Result.Items[0].Name = pol + "-clone"
	rqDst.Method = "importPolicies"

	rqDst.Params = map[string]interface{}{
		"records":       resp.Result.Items,
		"compatVersion": resp.Result.CompatVersion,
	}
	gzDst.Policy.SetRequest(rqDst)
	respDst, err := gzDst.Policy.ImportPolicies()
	if err != nil {
		log.Fatal("===> IMPORT ERROR: ", err)
	}
	return respDst.Result.Success
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
		BaseURL: fmt.Sprintf("https://%s/api/v1.0/jsonrpc/", result.Src.Server),
	}

	// Create Default Request struct
	rq = gobdgz.Request{
		Debug:      debug,
		APIKey:     result.Src.APIKey,
		JSONRPC:    "2.0",
		URL:        gz.BaseURL + "/policies",
		HttpMethod: "POST",
		ID:         "98409cc1-93cc-415a-9f77-1d4f681000b3",
	}
}

func init_dst_server() {

	// create new instance of GravityZone API
	gzDst = &gobdgz.GravityZoneAPI{
		BaseURL: fmt.Sprintf("https://%s/api/v1.0/jsonrpc/", result.Dst.Server),
	}

	// Create Default Request struct
	rqDst = gobdgz.Request{
		Debug:      debug,
		APIKey:     result.Dst.APIKey,
		JSONRPC:    "2.0",
		URL:        gzDst.BaseURL + "/policies",
		HttpMethod: "POST",
		ID:         "98409cc1-93cc-415a-9f77-1d4f681000b3",
	}
}

func initFlags() {

	configFile := flag.String("config", "", "config.json configuration file (https://github.com/javadmohebbi/gobdgz/tree/master/example/policy-export-import)")
	flag.Parse()
	// debug = true

	if *configFile == "" {
		fmt.Println("-config {file.json} is required!")
		flag.PrintDefaults()
		os.Exit(127)
	}

	file, err := os.Open(*configFile)
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatalln("JSON error:", err)
	}

	// log.Println(result)

}
