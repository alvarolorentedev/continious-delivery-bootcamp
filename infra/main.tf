terraform {
}

output "hello_world" {
    value = "Hello, World!"
}

data "archive_file" "zip" {
  type = "zip"
  output_path = "${path.root}/../output/lambda.zip"
  source_dir = "${path.root}/.."
  excludes = [ "infra", ".github", "output", ".git" ]
}

data "aws_iam_policy_document" "policy" {
  statement {
    sid = ""
    effect="Allow"
    principals {
        identifiers = ["lambda.amazonaws.com"]
        type        = "Service"
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"
  assume_role_policy = data.aws_iam_policy_document.policy.json
}

resource "aws_lambda_function" "lambda" {
    function_name = "my_function"
    filename = data.archive_file.zip.output_path
    source_code_hash = data.archive_file.zip.output_base64sha256
    role = aws_iam_role.iam_for_lambda.arn
    handler = "index.handler"
    runtime = "nodejs16.x"
  
}

output "rest_api_id" {
    value = ""
}

output "stage_name" {
    value = ""
}

