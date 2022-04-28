variable "aws-region" {
  default = "us-east-1"
}

variable "app-name" {
  type        = string
  description = "Application Name"
}

variable "service-name" {
  type        = string
  description = "Service Name"
}

variable "app-environment" {
  type        = string
  description = "Application Environment"
}

variable "memory" {
  description = "Amount of memory resource"
}

variable "desired-task-number" {
  description = "Number of instances running by default"
}

variable "docker-image-url" {
  description = "Docker image url"
}

variable "docker-container-port" {
  description = "Exposed Docker container port"
}



variable "parameter-path" {
  description = "Base path to parameters on the SSM Parameter Store"
}

variable "org-name" {
  description = "Terraform Cloud Organization name"
}

variable "ws-name" {
  description = "Terraform Cloud Organization name"
}