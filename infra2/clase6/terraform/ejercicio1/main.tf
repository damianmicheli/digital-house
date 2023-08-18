provider "aws" {
    region = "us-east-1"
    shared_credentials_files = ["$HOME/.aws/credentials"]
}

resource "aws_instance" "instancia_1_linux" {
    ami = "ami-053b0d53c279acc90"
    instance_type = "t2.micro"
    key_name = "saitama"
    tags = {
        Name = "Instancia_1_linux_desde_terraform_damian"
        Ambiente = "Ambiente de desarrollo"
    }
}

resource "aws_instance" "instancia_2_linux" {
    ami = "ami-053b0d53c279acc90"
    instance_type = "t2.micro"
    key_name = "saitama"
    tags = {
        Name = "Instancia_2_linux_desde_terraform_damian"
        Ambiente = "Ambiente de desarrollo"
    }
}

resource "aws_s3_bucket" "bucket_s3" {
  bucket = "my-tf-test-bucket-damian-nombre-unico"
  tags = {
    Name        = "Un_bucket_desde_terraform_damian"
    Ambiente = "Ambiente de desarrollo"
  }

}