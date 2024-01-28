package main

import (
    "context"
    "io/ioutil"
    "net/http"
    "strconv"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    version := getLatestVer()
    return &events.APIGatewayProxyResponse{
        StatusCode: 200,
        Headers: map[string]string{
            "Content-Type":   "text/plain",
            "Content-Length": strconv.FormatInt(int64(len(version)), 10),
            "Cache-Control":  "public, max-age=86400",
        },
        Body:            version,
        IsBase64Encoded: false,
    }, nil
}

func getLatestVer() string {
    version := "?"
    resp, err := http.Get("https://go.dev/VERSION?m=text")
    if err != nil {
        return version
    }
    resBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return version
    }
    return string(resBody)
}

func main() {
    lambda.Start(handler)
}

