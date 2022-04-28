resource "aws_security_group" "profile-api-sg" {
  name        = "${var.service-name}-sg"
  description = "Security group that allows internet traffic in and out of the profile-api service"
  vpc_id      = data.tfe_outputs.platform.values.vpc_id

  ingress {
    from_port   = 8080
    protocol    = "TCP"
    to_port     = 8080
    cidr_blocks = [data.tfe_outputs.platform.values.vpc_cidr_block]
  }


  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.service-name}-sg"
    Environment = var.app-environment
  }
}

resource "aws_alb_target_group" "profile-api-tg" {
  name        = "${var.service-name}-tg"
  port        = var.docker-container-port
  protocol    = "HTTP"
  vpc_id      = data.tfe_outputs.platform.values.vpc_id
  target_type = "ip"

  health_check {
    path                = "/health"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 60
    timeout             = 30
    unhealthy_threshold = "3"
    healthy_threshold   = "3"
  }

  tags = {
    Name        = "${var.service-name}-tg"
    Environment = var.app-environment
  }
}

resource "aws_alb_listener_rule" "profile-api-lr" {
  listener_arn = data.tfe_outputs.platform.values.ecs_alb_https_listener_arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.profile-api-tg.arn
  }
  condition {
    host_header {
      values = ["${var.service-name}.alb.${data.tfe_outputs.platform.values.ecs_domain_name}"]
    }
  }
}