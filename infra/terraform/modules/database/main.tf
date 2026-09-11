resource "aws_db_subnet_group" "main" {
  name       = "cloudboard-db-subnet-group"
  subnet_ids = var.db_subnet_ids

  tags = {
    Name = "cloudboard-db-subnet-group"
  }
}

resource "aws_security_group" "db" {
  name        = "cloudboard-db-sg"
  description = "Allow inbound traffic from VPC to PostgreSQL"
  vpc_id      = var.vpc_id

  ingress {
    description = "PostgreSQL from VPC"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [var.vpc_cidr]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "cloudboard-db-sg"
  }
}

resource "aws_db_instance" "main" {
  allocated_storage      = 20
  engine                 = "postgres"
  engine_version         = "15"
  instance_class         = "db.t3.micro"
  db_name                = var.db_name
  username               = var.db_username
  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.main.name
  vpc_security_group_ids = [aws_security_group.db.id]
  skip_final_snapshot    = true
  publicly_accessible    = false

  tags = {
    Name = "cloudboard-db"
  }
}
