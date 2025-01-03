variable "region" {
  description = "The AWS region to deploy the resources"
  type        = string
  default     = "ca-central-1"
}

variable "subnet" {
  description = "The subnet to deploy the resources"
  type        = string
  default     = "ca-central-1a"
}

terraform {
  backend "s3" {
    bucket = "terraform-state-autocross-timer"
    key    = "terraform.tfstate"
    region = "ca-central-1"
  }
}

provider "aws" {
  region = var.region
}

data "aws_ami" "amzn-linux-2023-ami" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["al2023-ami-2023.*-x86_64"]
  }
}

resource "aws_vpc" "autocross_timer_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "autocross_timer_vpc"
  }
}

resource "aws_subnet" "autocross_timer_subnet" {
  vpc_id            = aws_vpc.autocross_timer_vpc.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = var.subnet

  tags = {
    Name = "autocross_timer_subnet"
  }
}

resource "aws_instance" "autocross_timer" {
  ami           = data.aws_ami.amzn-linux-2023-ami.id
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.autocross_timer_subnet.id
  key_name      = aws_key_pair.github_actions.key_name
  associate_public_ip_address = true
  vpc_security_group_ids = [aws_security_group.autocross_timer_sg.id]
  iam_instance_profile = aws_iam_instance_profile.ssm_profile.name
  depends_on = [ aws_internet_gateway.autocross_timer_igw ]

  ebs_block_device {
    device_name = "/dev/xvda"
    volume_size = 30
    volume_type = "gp3"
  }

  tags = {
    Name = "autocross_timer"
  }
}

resource "aws_key_pair" "github_actions" {
  key_name   = "github-actions"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDaE8/AJreu4jpGeA75+O8ZTUKm6uRtgRYumX5k/oj1B3oOWxCNjMqflTg+2Fwr1DgUxldzTb+th9bDS9E2VQmAnaEFNfOrSUsdbcAdDtI16ACnZod8XCj4rGn4LTCDBaUjIzl3ydBQ209ui2lbVP7ymj5nJ0eEI+j2HaDwgeNd/xMwUthszqXPvr7uVlVWFG17Y/8D/vaG7lPt1KR9JiYXRcAqYfq8hiAJvXVkqveFghcH21CsonVIdfvwzih9erefDb3s45kAHVujo7IBtoqasCtnjReOETRsHTPc0jN0uaukl49vtxEpYiE8Gu5cMGwTRjM6TqLlusWFqr6u9hqY56LckIzodx4VZlYpk3Tar6RN7xn15ZtB2SUQ1IvUFW0PBOz2NIfskUGlKLjTebdkEWux3g7T5L49SrNv8K24q9dxpo5sOsQ8xfdqFQvhjalpDOo7wsxZdTlwJmUzUcsyDJYNIEXn2JIbEMiwNviafAt0a6DIz0YN7tFYFZ9WbjqDkBsf0yJS4ETX4GIGeF6o7TsXLuJ9BmG8+CTHds9NwbJVlVWiQTDe/vinVChe2zHL12g6pgjqFscGLN/CtwebCMz5fHpgYTNINX3e+CcipISiDGlWQhemIbjWaH5RUzm8I213M2svVmhxYG2oKdcGDrzur1CHccVnFaU5CIXlJw== githubactions@gh"
}

resource "aws_security_group" "autocross_timer_sg" {
  vpc_id = aws_vpc.autocross_timer_vpc.id
  name = "autocross_timer_sg"

  tags = {
    Name = "autocross_timer_sg"
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_vpc_security_group_ingress_rule" "allow_https" {
  security_group_id = aws_security_group.autocross_timer_sg.id
  cidr_ipv4         = "0.0.0.0/0"
  from_port         = 443
  ip_protocol       = "tcp"
  to_port           = 443
}

resource "aws_iam_role" "ssm_role" {
  name = "SSMRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ssm_attach" {
  role       = aws_iam_role.ssm_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_iam_instance_profile" "ssm_profile" {
  name = "SSMInstanceProfile"
  role = aws_iam_role.ssm_role.name
}

resource "aws_internet_gateway" "autocross_timer_igw" {
  vpc_id = aws_vpc.autocross_timer_vpc.id

  tags = {
    Name = "autocross_timer_igw"
  }
}

resource "aws_route" "route" {
  route_table_id         = aws_vpc.autocross_timer_vpc.main_route_table_id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.autocross_timer_igw.id
}

resource "aws_ssm_document" "install_docker" {
  name = "install_docker"
  document_type = "Command"
  content = <<DOC
  {
    "schemaVersion": "2.2",
    "description": "Install Docker",
    "parameters": {},
    "mainSteps": [
      {
        "action": "aws:runShellScript",
        "name": "install-docker",
        "inputs": {
          "runCommand": [
            "sudo yum update -y",
            "sudo yum install -y docker",
            "sudo service docker start",
            "sudo usermod -a -G docker ec2-user"
          ]
        }
      }
    ]
  }
  DOC
}