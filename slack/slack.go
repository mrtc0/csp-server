package slack

import (
	"fmt"
	"os"

	slack "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/mrtc0/csp-server/report"
)

func Send(report report.Report) {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	color := "danger"
	attachment1 := slack.Attachment{Color: &color}

	attachment1.AddField(slack.Field{Title: "Time", Value: fmt.Sprintf("%s", report.Timestamp)})
	attachment1.AddField(slack.Field{Title: "document-uri", Value: report.CspReport.DocumentURI})
	attachment1.AddField(slack.Field{Title: "referrer", Value: report.CspReport.Referrer})
	attachment1.AddField(slack.Field{Title: "blocked-uri", Value: report.CspReport.BlockedURI})
	attachment1.AddField(slack.Field{Title: "violated-directive", Value: report.CspReport.ViolatedDirective})
	attachment1.AddField(slack.Field{Title: "original-policy", Value: report.CspReport.OriginalPolicy})
	attachment1.AddField(slack.Field{Title: "raw", Value: fmt.Sprintf("```%+v```", report)})

	payload := slack.Payload{
		Text:        "*CSPのレポートが届きました*\n\n",
		Username:    "cspbot",
		Channel:     "#csp",
		IconEmoji:   ":robot_face:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
