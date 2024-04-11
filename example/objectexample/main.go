package main

import (
	"fmt"
	jsonql "github.com/NubeIO/jql"
)

func main() {
	jsonString := `
	{
	    "id": "rubix-manager",
	    "info": {
	        "objectID": "rubix-manager",
	        "objectType": "rubix-network",
	        "category": "rubix",
	        "pluginName": "main",
	        "workingGroup": "rubix",
	        "workingGroupLeader": "rubix-manager",
	        "workingGroupObjects": [
	            "rubix-manager",
	            "rubix-network",
	            "rubix-mapping"
	        ],
	        "workingGroupChildObjects": [
	            "rubix-network"
	        ],
	        "objectTags": [
	            "drivers",
	            "mqtt",
	            "rest",
	            "rubix"
	        ],
	        "permissions": {
	            "canBeUpdated": true
	        },
	        "requirements": {
	            "maxOne": true,
	            "mustLiveInObjectType": true,
	            "requiresLogger": true
	        }
	    },
	    "outputs": [
	        {
	            "id": "err",
	            "name": "err",
	            "portUUID": "ieae7f63e1961810e",
	            "direction": "output",
	            "dataType": "float",
	            "defaultPosition": 1
	        }
	    ],
	    "meta": {
	        "objectUUID": "d3383e797a6043b78ac2591662281125",
	        "objectName": "rubix-manager"
	    }
	}
	`

	j := jsonql.New()
	j.NewStringData(jsonString)
	q := "id=rubix-manager"
	fmt.Println("query", q)
	fmt.Println(j.Query("id='rubix-manager'"))

	//fmt.Println(parser.Query("info.objectType='rubix-network' && info.category='rubix'"))
}
