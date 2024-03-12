package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "io/ioutil"
    "os"
    "encoding/json"
    "os/exec"


)


//Product defines a structure for an item in product catalog
type GameData struct {
    GameDate string `json:"GameDate"`
    TeamName string `json:"TeamName"`
   // TeamWinner string `json:"team_winner"`
    //TotalRebounds string `json:"total_rebounds"`
    //FieldGoalsAttempted string `json:"field_goals_attempted"`
    //FreeThrowsAttempted string `json:"free_throws_attempted"`
    //Steals string `json:"steals"`
}

// GutProductHandler is used to get data inside the products defined on our product catalog
func PutProductHandler() http.HandlerFunc {
    return func(rw http.ResponseWriter, r * http.Request) {
	    fmt.Printf("requested")
        // Read incoming JSON from request body
        data, err := ioutil.ReadAll(r.Body)
        // If no body is associated return with StatusBadRequest
        if err != nil {
	rw.WriteHeader(http.StatusBadRequest)
            return
                }

	//load the input data, error if format is wrong
        var game GameData 
	err = json.Unmarshal([]byte(data), &game)
		if err != nil {
			fmt.Printf("could not ready json: %s\n", err)
			return
		}
	fmt.Printf("json map: %s\n", data)

        //put the JSON data in an environment variables
        //fmt.Println('game date %s \n', d.GameDate)
	err = os.Setenv("date_input", game.GameDate)
        if err != nil {
            fmt.Println("Error setting game_input:", err)
            return
        }
        fmt.Println("new date_input $date_input")

	err = os.Setenv("team_input", game.TeamName)
        if err != nil {
            fmt.Println("Error setting team_input:", err)
            return
        }

        // run the bq command makes a new ML input table
        fmt.Printf("bq setting_input \n")
	out, err := exec.Command("./get_input.sh").Output() 
        if err != nil {
            fmt.Println("Error setting ML input:", err)
            return    //log.Fatal(err)
        }
        fmt.Printf("output is %s\n", out)
        fmt.Printf("bq predicting \n")
        pred_out, err := exec.Command("./predict.sh").Output()
        if err != nil {
            fmt.Println("Error setting ML input:", err)
                //log.Fatal(err)
        }
	fmt.Printf("prediction done")
        // Write the body with JSON data
        rw.Header().Add("content-type", "application/json")
                rw.WriteHeader(http.StatusFound)
        rw.Write(pred_out)
    }
}



// Create new Router

func main() {
    // Create new Router
    router := mux.NewRouter()

    // route properly to respective handlers
        router.Handle("/", PutProductHandler()).Methods("PUT")

    // Create new server and assign the router
    server := http.Server {
        Addr: ":9090",
        Handler: router,
    }
    fmt.Println("Starting NCAA Game Predictor on Port 9090")
    // Start Server on defined port/host.
    server.ListenAndServe()
}
