output "security_group_names" {
  value       = [for sg in aws_security_group.security_group : sg.name]
  description = "List of security group"
}
