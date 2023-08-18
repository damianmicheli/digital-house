variable "region" {
    type = string
    default = "us-east-1"
}

variable "credentials" {
    type = string
    default = "$HOME/.aws/credentials"
}

variable "image_id" {
    type = string
    default = "ami-053b0d53c279acc90"
}

variable "type" {
    type = string
    default = "t2.micro"
}

variable "key" {
    type = string
    default = "saitama"
}

variable "ambiente" {
    type = string
    default = "Ambiente de desarrollo"
}

variable "nombre_1" {
    type = string
    default = "Instancia_1_linux_desde_terraform_damian"
}

variable "nombre_2" {
    type = string
    default = "Instancia_2_linux_desde_terraform_damian"
}

variable "nombre_3" {
    type = string
    default = "Un_bucket_desde_terraform_damian"
}

variable "nombre_bucket" {
    type = string
    default = "my-tf-test-bucket-damian-nombre-unico-2"
}

