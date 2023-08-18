resource "aws_instance" "instancia_1_linux" {
    ami = var.image_id
    instance_type = var.type
    key_name = var.key
    tags = {
        Name = var.nombre_1
        Ambiente = var.ambiente
    }
}

resource "aws_instance" "instancia_2_linux" {
    ami = var.image_id
    instance_type = var.type
    key_name = var.key
    tags = {
        Name = var.nombre_2
        Ambiente = var.ambiente
    }
}

resource "aws_s3_bucket" "bucket_s3" {
  bucket = var.nombre_bucket
  tags = {
    Name = var.nombre_3
    Ambiente = var.ambiente
  }
}