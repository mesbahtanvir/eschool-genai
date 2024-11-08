# main.tf

provider "aws" {
  region = "us-west-2" # Replace with your desired AWS region
}

resource "aws_instance" "app_server" {
  ami           = "ami-0c55b159cbfafe1f0" # Replace with a suitable AMI ID for your region
  instance_type = "t2.micro"              # Choose an instance type
  key_name      = "your-key-pair"         # Replace with your existing key pair name

  # Security group
  vpc_security_group_ids = [aws_security_group.app_sg.id]

  # Add a startup script to install backend and frontend services
  user_data = <<-EOF
              #!/bin/bash
              sudo apt update -y
              sudo apt install -y nginx python3-pip

              # Clone your backend and frontend services
              git clone https://github.com/yourusername/your-backend-repo.git /home/ubuntu/backend
              git clone https://github.com/yourusername/your-frontend-repo.git /home/ubuntu/frontend

              # Install dependencies for backend and start it
              cd /home/ubuntu/backend
              pip3 install -r requirements.txt
              nohup python3 app.py & # Replace with the backend start command

              # Serve the frontend with Nginx
              sudo cp -r /home/ubuntu/frontend/* /var/www/html/

              sudo systemctl restart nginx
              EOF

  tags = {
    Name = "app-server"
  }
}

# Security group to allow HTTP and SSH access
resource "aws_security_group" "app_sg" {
  name        = "app-sg"
  description = "Allow HTTP and SSH inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Only for demo, restrict in production
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
