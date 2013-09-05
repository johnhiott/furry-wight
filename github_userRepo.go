package main

import "fmt"
import "net/http"
import "io/ioutil"
import "flag"
import "encoding/json"

const sURL = "https://api.github.com/users/";

func parseJSON(b []byte){

	var repo []map[string]interface{}

	err := json.Unmarshal(b, &repo)
	if err != nil{
		fmt.Println("error:", err)
		return;
	}

	for i := 0; i<len(repo); i++{
		fmt.Print(repo[i]["name"], "\n");
	}
}

func generateURL(s string) string{
	return sURL + s + "/repos"
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) < 1{
		fmt.Println("No parameters, specify username")
		return
	}

	url := generateURL(args[0])
	
	resp, err := http.Get(url);
	if err != nil {
		//TODO: Handle Error
	}

	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);

	parseJSON(body);

}