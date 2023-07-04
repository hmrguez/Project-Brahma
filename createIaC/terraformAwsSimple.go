package createIaC

var terraformawssimple = `
# Configure the AWS provider
provider "aws" {
  region = "us-west-2"
  access_key = "<access-key>"
  secret_key = "<secret-key>"
}

# Create a VPC
resource "aws_vpc" "ecommerce_vpc" {
  cidr_block = "10.0.0.0/16"
}

# Create a public subnet for the EC2 instances
resource "aws_subnet" "public_subnet" {
  vpc_id     = aws_vpc.ecommerce_vpc.id
  cidr_block = "10.0.1.0/24"
}

# Create a private subnet for the RDS instance
resource "aws_subnet" "private_subnet" {
  vpc_id     = aws_vpc.ecommerce_vpc.id
  cidr_block = "10.0.2.0/24"
}

# Create a security group for the EC2 instances
resource "aws_security_group" "webapp_security_group" {
  name_prefix = "webapp_"
  vpc_id      = aws_vpc.ecommerce_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create a security group for the RDS instance
resource "aws_security_group" "rds_security_group" {
  name_prefix = "rds_"
  vpc_id      = aws_vpc.ecommerce_vpc.id

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["10.0.1.0/24"] # Allow traffic from EC2 instances
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create an RDS instance
resource "aws_db_instance" "webapp_db" {
  identifier             = "webapp-db"
  engine                 = "mysql"
  engine_version         = "5.7"
  instance_class         = "db.t2.micro"
  allocated_storage      = 20
  storage_type           = "gp2"
  storage_encrypted      = true
  multi_az               = false
  publicly_accessible    = false
  db_subnet_group_name   = aws_db_subnet_group.webapp_db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.rds_security_group.id]

  tags = {
    Environment = "production"
  }
}

# Create a DB subnet group for the RDS instance
resource "aws_db_subnet_group" "webapp_db_subnet_group" {
  name        = "webapp-db-subnet-group"
  subnet_ids  = [aws_subnet.private_subnet.id]
  description = "Subnet group for webapp-db"
}

# Create an EC2 instance
resource "aws_instance" "webapp_instance" {
  count         = 3
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  subnet_id        = aws_subnet.public_subnet.id
  vpc_security_group_ids = [aws_security_group.webapp_security_group.id]

  user_data = <<-EOF
              #!/bin/bash
              # Install and configure the web app
              EOF
}

# Create a load balancer
resource "aws_lb" "webapp_lb" {
  name               = "webapp-lb"
  internal           = false
  load_balancer_type = "application"
  subnets            = [aws_subnet.public_subnet.id]

  security_groups = [aws_security_group.webapp_security_group.id]

  tags = {
    Environment = "production"
  }
}

# Attach the EC2 instances to the load balancer
resource "aws_lb_target_group_attachment" "webapp_lb_attachment" {
  count            = 3
  target_group_arn = aws_lb_target_group.webapp_target_group.arn
  target_id        = aws_instance.webapp_instance.*.id[count.index]
  port             = 80
}

# Create a target group for the load balancer
resource "aws_lb_target_group" "webapp_target_group" {
  name_prefix      = "watg-"
  port             = 80
  protocol         = "HTTP"
  target_type      = "instance"
  vpc_id           = aws_vpc.ecommerce_vpc.id

  health_check {
    healthy_threshold   = 2
    interval            = 30
    protocol            = "HTTP"
    timeout            = 5
    unhealthy_threshold = 2
    path                = "/"
  }

  stickiness {
    type    = "lb_cookie"
    enabled = true
  }
}
`
