package xml

import (
    "fmt"
	"log"
    "net/http"
  	"math/big"
	"encoding/xml"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ReportAid/app/src/server/internal/contracts"
	"github.com/ReportAid/app/src/server/internal/configs"
)

// ActivitiesList - a list if all activities
type ActivitiesList struct {
	ID	[]string `xml:"activities-id"`
}

// Activities - the XML for the activities file
type Activities struct {
	XMLName   xml.Name `xml:"iati-activities"`
	Version 		string `xml:"version,attr"`
	Time 				string `xml:"generated-datetime,attr"`
	LinkedData  string `xml:"linked-data-default,attr"`
}

// TotalActivities - get the total number of activities
type TotalActivities struct {
	Total 		int64 `xml:"total-activities"`
}

// GetNumActivites - Get the total number of activities
func numActivites(contract *contracts.IATIActivities) (int64) {

	num, err := contract.GetNumActivities(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("GetNumActivities Failed: %v", err)
		return 0
	}
	smallNum := num.Int64()
	return smallNum
}

// NumActivites - get total activities
func NumActivites (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

	num := numActivites(contract)
	totalXML := &TotalActivities{Total: num}
	thisXML, _ := xml.MarshalIndent(totalXML, "", " ")
	w.Write(thisXML)
	w.Write([]byte("\n"))
}

// ActivitiesID get specific acvtitie
func ActivitiesID (c *ethclient.Client, contract *contracts.IATIActivities, ref [32]byte, w http.ResponseWriter, r *http.Request) {

	//fmt.Printf("Activities Reference: %x\n", ref)

	activities, err := contract.GetActivities(&bind.CallOpts{}, ref)
	if err != nil {
		log.Fatalf("GetActivities Failed: %v", err)
	}

	thisVersion := activities.Version
	sliceVersion := thisVersion[:]
	trimmedVersion := common.TrimRightZeroes(sliceVersion)
	version := hexutil.Encode(trimmedVersion)
	stringVersion, _ := hexutil.Decode(version)

	thisTime := activities.GeneratedTime
	sliceTime := thisTime[:]
	trimmedTime := common.TrimRightZeroes(sliceTime)
	time := hexutil.Encode(trimmedTime)
	stringTime, _ := hexutil.Decode(time)

	thisLink := activities.LinkedData
	sliceLink := thisLink[:]
	trimmedLink := common.TrimRightZeroes(sliceLink)
	link := hexutil.Encode(trimmedLink)
	stringLink, _ := hexutil.Decode(link)

	/* fmt.Printf("Version: %s\n", stringVersion)
	fmt.Printf("Generated Time: %s\n", stringTime)
	fmt.Printf("Linked Data: %s\n\n", stringLink)*/

	activitiesXML := &Activities{Version: string(stringVersion), Time: string(stringTime), LinkedData: string(stringLink)}
	thisXML, _ := xml.MarshalIndent(activitiesXML, " ", " ")
	w.Write(thisXML)
}


// ListActivities - list all activities
func ListActivities (c *ethclient.Client, contract *contracts.IATIActivities, w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("\n"))
	num := numActivites(contract)
	var i int64
	activitiesIDs := &ActivitiesList{}
	for ; i < num; i++ {
		ref, err := contract.GetReference(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Fatalf("GetReference Failed: %v", err)
		}
		thisRef := fmt.Sprintf("%x", ref)
		activitiesIDs.ID = append(activitiesIDs.ID, thisRef)
	}
	thisXML, _ := xml.MarshalIndent(activitiesIDs, " ", " ")
	w.Write(thisXML)
}

// ActivitesContract - get activities contract address
func ActivitesContract (c *ethclient.Client) (*contracts.IATIActivities){

	contract, err := contracts.NewIATIActivities(common.HexToAddress(string(configs.ActivitiesContractAddr)), c)
	if err != nil {
		log.Fatalf("Failed to instantiate activities contract: %v", err)
		return nil
	}
	return contract
}
