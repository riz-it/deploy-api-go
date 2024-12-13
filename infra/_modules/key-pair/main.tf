resource "tls_private_key" "rsa_4096" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "ssh-key" {
  key_name   = var.key_pair_name
  public_key = tls_private_key.rsa_4096.public_key_openssh
  depends_on = [ tls_private_key.rsa_4096 ]
}

resource "local_file" "private_key" {
  content  = tls_private_key.rsa_4096.private_key_pem
   filename = "${var.key_pair_name}_pem"
}

resource "null_resource" "fix_key_permissions" {
  depends_on = [local_file.private_key]

  provisioner "local-exec" {
    command = "chmod 600 ${local_file.private_key.filename}"
  }
}
