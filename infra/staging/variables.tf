variable "aws_access_key" {
  type      = string
  sensitive = true
}

variable "aws_secret_key" {
  type      = string
  sensitive = true
}

variable "aws_region" {
  type = string
}

variable "aws_vpc_id" {
  type = string
}

variable "aws_user" {
  description = "User"
  type        = string
  default     = "ubuntu"
}

variable "key_pair_name" {
  description = "Key Pair Name"
  type        = string
}

variable "ec2_list" {
  type = list(object({
    name          = string
    instance_type = string
    image_id      = string
  }))
}

variable "sg_list" {
  type = list(object({
    name = string
    ingress = object({
      description = string
      from_port   = number
      to_port     = number
      protocol    = string
      cidr_blocks = list(string)
    })
    egress = object({
      description = string
      from_port   = number
      to_port     = number
      protocol    = string
      cidr_blocks = list(string)
    })
  }))
}

# Gitlab Access Token
variable "gitlab_access_token" {
  type        = string
  description = "Gitlab Access Token"
  sensitive   = true
}

# Repository List
variable "repository_list" {
  type = list(object({
    name        = string
    description = string
    visibility  = string
  }))
}
