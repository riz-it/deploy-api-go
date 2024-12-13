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

#Key Pair
variable "key_pair_name" {
  description = "key_pair_name"
  type        = string
}

# Record List
variable "ec2_list" {
  type = list(object({
    name          = string
    instance_type = string
    image_id      = string
  }))
}

# Security Group List
variable "security_group_names" {
  type = list(string)
}
