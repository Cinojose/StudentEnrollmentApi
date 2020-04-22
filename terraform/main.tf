// AWS provider

provider "aws" {
  region = "us-east-1"
}

terraform {
    backend "s3" {
    bucket="tfstate-pf20"
    key="terraform/dev/terraform_dev.tfstate"
    region="ap-southeast-1"
    }
}

// Fing the ubuntu AMI to use

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

data "http" "myip" {
  url = "http://ipv4.icanhazip.com"
}

// Security group  that allow only 8000
resource "aws_security_group" "allow_home_ip" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"
  #vpc_id      = "${aws_vpc.main.id}"

  ingress {
    description = "Default API port"
    from_port   = 8001
    to_port     = 8001
    protocol    = "tcp"
    cidr_blocks = ["${chomp(data.http.myip.body)}/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "allow_tls"
  }
}

data "template_file" "userdata" {
  template = "${file("${path.module}/startup.sh.tpl")}"
  vars = {
    mysql_image_name = "student-mysql"
    mysql_password = "XXXXXXXXX"
  }
}

// Create a key to attach to the instance
resource "aws_key_pair" "web-server-key" {
  key_name   = "webserver-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDJ4XGmk3FrbMYPhZSvUYOImDi0RJRbl0Jv0IwjKiBjqp71Ave+CSJza5BI9IkBeGKmaEPzOJEs26pU1Gx005Pc7s67MIvuH0MGwW28sC8xHo/hhTsNpBXiAh1sveX5SMt7hFRLJKcRo170JjmOu6ZOwkt/qLIhKBX27w+uPuZ0Gc2tfc8Z4ONbNZEkxlnoQWfIXufGgsxXaul+6ZAHQyEKKOEr1RgxXiVbqv61IrKTRhzIkV5be2hiK3UhDoiaHEuzUmqJUm44J+gu4XmUrTUbU09zxR7cMMlKJ9ICfnO2MpLqg/uiUm4DWuUxczy9ctnpQoLyeekei3aKkUkMXYm1 cjff20@gmail.com"
}

// New EC2 instance for web server
resource "aws_instance" "webserver" {
   ami           = "${data.aws_ami.ubuntu.id}"
   instance_type = "t2.micro"
   // user data
   user_data = "${data.template_file.userdata.rendered}"
   // key to login
   key_name =  "${aws_key_pair.web-server-key.key_name}"
   //  Security group
   security_groups = ["${aws_security_group.allow_home_ip.id}"]
   tags = {
    Name = "StudentEntollmentAPI"
  }
}



