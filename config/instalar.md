## Compartir clipboard entre maquina virtual y host
sudo apt install open-vm-tools-desktop

## credenciales
Datos de acceso a AWS: https://dh-infra2-0723.signin.aws.amazon.com/console

Usuario:0723-grupo20
Contraseña:*********


## aws configure
aws configure
AWS Access Key ID [None]: ************************
AWS Secret Access Key [None]: ****************************
Default region name [None]: us-east-1
Default output format [None]: json



## configurar git
git config --global user.name "Damián Micheli"
git config --global user.email damianmicheli@gmail.com

## github cli para no tener que cargar siempre la contraseña (token)
sudo apt install gh
gh auth login

## crear llave ssh
ssh-keygen



## ver todas las interfaces de red
ip a


## Hacer ejecutable un script
chmod +x script.sh 

## instalar ssh para usar winscp
 sudo apt-get install openssh-server


## instalar go
### Descargar desde: https://go.dev/doc/install

sudo rm -rf /usr/local/go 
tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
nano $HOME/.profile
### y agregar esta linea al final:
export PATH=$PATH:/usr/local/go/bin


## habilitar una clave pem
chmod 400 pochita2.pem 

## conectarme por ssh a instancia ec2
ssh -i pochita.pem ubuntu@3.91.218.181

## si no anda bien ansible
sudo apt install python3-pip
pip install boto3
ansible-galaxy collection install amazon.aws

### esto no se si hace falta
sudo apt-add-repository ppa:ansible/ansible
pip install boto


## editar hosts de ansible (si no existe crearlo)
sudo nano /etc/ansible/hosts

## Cool Retro Term
sudo apt install cool-retro-term

## GO
### Instalar GO
sudo apt install golang-go

### editar ~/.profile:
export PATH=$PATH:/usr/local/go/bin
