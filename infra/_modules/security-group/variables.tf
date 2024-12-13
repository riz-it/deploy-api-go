# AWS Region
variable "aws_region" {
  type        = string
  description = "AWS Region"
}

# AWS Access Key
variable "aws_access_key" {
  type        = string
  description = "AWS Access Key"
}

# AWS Secret Access Key
variable "aws_secret_key" {
  type        = string
  description = "AWS Secret Key"
  sensitive   = true
}

# VPC ID
variable "aws_vpc_id" {
  type        = string
  description = "AWS VPC ID"
}

# Security Group List
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

