package infra

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/files"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func Setup() {
	files.CopyFile("./provider.tf", "./provider.tf.backup")
	files.CopyFile("./test/artifacts/provider.tf", "./provider.tf")
}

func CleanUp() {
	defer os.RemoveAll("./terraform.tfstate")
	defer os.RemoveAll("./terraform.tfstate.backup")
	defer os.RemoveAll("./terraform")
	files.CopyFile("./provider.tf.backup", "./provider.tf")
	defer os.Remove("./provider.tf.backup")
}

func TestLambdaIsAccesible(t *testing.T) {
	Setup()
	defer CleanUp()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: ".",
	})
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	restApiId := terraform.Output(t, terraformOptions, "rest_api_id")
	stageName := terraform.Output(t, terraformOptions, "stage_name")
	url := fmt.Sprintf("http://localhost:4566/restapis/%s/%s/_user_request_/health", restApiId, stageName)
	http_helper.HttpGetWithRetry(t, url, nil, 200, "", 10, 5*time.Second)
}
