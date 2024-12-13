resource "aws_instance" "ec2" {
  for_each        = { for idx, ec in var.ec2_list : idx => ec }
  ami             = each.value.image_id
  instance_type   = each.value.instance_type
  security_groups = var.security_group_names
  tags = {
    Name = each.value.name
  }
  associate_public_ip_address = true
  key_name         = var.key_pair_name
}

resource "aws_eip" "ec2_eip" {
  count    = length(var.ec2_list)
  instance = aws_instance.ec2[count.index].id
}
