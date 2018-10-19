package clients

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-resty/resty"
	"github.com/spf13/viper"
	"github.com/steviebiddles/jira-timesheets/models"
)

func initClient() {
	resty.SetHostURL(viper.GetString("host"))
	resty.SetHeader("Accept", "application/json")
	resty.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	resty.SetBasicAuth(
		viper.GetString("email"),
		viper.GetString("apiToken"),
	)
}

func GetIssueWorklogs(issue string) (worklogs *models.Worklogs) {
	initClient()

	worklogs = &models.Worklogs{}

	resp, _ := resty.R().
		SetResult(worklogs).
		Get(fmt.Sprintf("/rest/api/3/issue/%s/worklog", issue))

	switch resp.StatusCode() {
		case 200:
			color.Green(resp.Status())
		case 400:
			color.Yellow(resp.Status())
			fmt.Println(resp)
		default:
			color.Red(resp.Status())
	}

	return worklogs
}

func PostIssueWorklog(issue string, worklog models.Worklog) {
	initClient()

	resp, _ := resty.R().
		SetBody(worklog).
		Post(fmt.Sprintf("/rest/api/3/issue/%s/worklog", issue))

	switch resp.StatusCode() {
		case 201:
			color.Green(resp.Status())
			fmt.Println("---")
			fmt.Println(fmt.Sprintf("Location: %s", resp.Header().Get("Location")))
		case 400:
			color.Yellow(resp.Status())
			fmt.Println(resp)
		default:
			color.Red(resp.Status())
	}
}

func DeleteIssueWorklog(issue string, id string) {
	initClient()

	resp, _ := resty.R().
		Delete(fmt.Sprintf("/rest/api/3/issue/%s/worklog/%s", issue, id))

	switch resp.StatusCode() {
	case 204:
		color.Green(resp.Status())
	case 400:
		color.Yellow(resp.Status())
		fmt.Println(resp)
	default:
		color.Red(resp.Status())
	}
}
