provider "aws" {
    region = "us-east-1"
    shared_credentials_files = ["$HOME/.aws/credentials"]
}

resource "aws_instance" "instancia_linux" {
    ami = "ami-053b0d53c279acc90"
    instance_type = var.tipo_de_instancia
    key_name = "saitama"
    tags = {
        Name = "Instancia_desde_terraform_test"
        Ambiente = "Ambiente de desarrollo"
    }
}