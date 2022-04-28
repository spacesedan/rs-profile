data "aws_ssm_parameter" "mongo-uri" {
  name = "${var.parameter-path}/mongo_uri"
}

data "aws_ssm_parameter" "db-name" {
  name = "${var.parameter-path}/db_name"
}

data "aws_ssm_parameter" "fire-prox" {
  name = "${var.parameter-path}/fire_prox"
}

data "aws_ssm_parameter" "os-api-key" {
  name = "${var.parameter-path}/os_api_key"
}