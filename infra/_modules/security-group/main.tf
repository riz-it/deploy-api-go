resource "aws_security_group" "security_group" {
  for_each = { for idx, sg in var.sg_list : idx => sg }

  name   = each.value.name
  vpc_id = var.aws_vpc_id

  ingress {
    description = each.value.ingress.description
    from_port   = each.value.ingress.from_port
    to_port     = each.value.ingress.to_port
    protocol    = each.value.ingress.protocol
    cidr_blocks = each.value.ingress.cidr_blocks
  }

  egress {
    description = each.value.egress.description
    from_port   = each.value.egress.from_port
    to_port     = each.value.egress.to_port
    protocol    = each.value.egress.protocol
    cidr_blocks = each.value.egress.cidr_blocks
  }
}
