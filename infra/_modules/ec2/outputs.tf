output "ec2_public_ip" {
   value = [for instance in aws_eip.ec2_eip : instance.public_ip]
   description = "List of instance public IP addresses"
}
