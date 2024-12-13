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
  description = "Key Pair Name"
  type        = string
}
