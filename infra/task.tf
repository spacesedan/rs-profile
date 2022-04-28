resource "aws_ecs_task_definition" "profile-api" {
  container_definitions = jsonencode([
    {
      name        = var.service-name,
      image       = var.docker-image-url,
      essential   = true,
      environment = [
        {
          name  = "APP_ENV",
          value = "production"
        }
      ],
      secrets = [
        {
          name      = "MONGO_URI",
          valueFrom = data.aws_ssm_parameter.mongo-uri.arn
        },
        {
          name      = "DB",
          valueFrom = data.aws_ssm_parameter.db-name.arn

        },
        {
          name      = "FIRE_PROX",
          valueFrom = data.aws_ssm_parameter.fire-prox.arn
        },
        {
          name      = "OS_KEY",
          valueFrom = data.aws_ssm_parameter.os-api-key.arn
        }
      ],
      portMappings = [
        {
          containerPort = var.docker-container-port
        }
      ],
      logConfiguration = {
        logDriver = "awslogs",
        options   = {
          awslogs-group         = aws_cloudwatch_log_group.log-group.id
          awslogs-region        = var.aws-region,
          awslogs-stream-prefix = "${var.service-name}-lg"
        }
      }
    }
  ])
  family                   = "${var.service-name}-task"
  cpu                      = 256
  memory                   = var.memory
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.profile-api-role.arn
  task_role_arn            = aws_iam_role.profile-api-role.arn
}

data "aws_ecs_task_definition" "main" {
  task_definition = aws_ecs_task_definition.profile-api.family
}

resource "aws_cloudwatch_log_group" "log-group" {
  name = "${var.service-name}-lg"
}

resource "aws_ecs_service" "profile-api-service" {
  name = "${var.service-name}-srv"
  task_definition = "${aws_ecs_task_definition.profile-api.family}:${max(aws_ecs_task_definition.profile-api.revision, data.aws_ecs_task_definition.main.revision)}"
  desired_count = var.desired-task-number
  cluster = data.tfe_outputs.platform.values.ecs_cluster_name
  launch_type = "FARGATE"

  network_configuration {
    subnets = split(",", join(",", data.tfe_outputs.platform.values.ecs_public_subnets))
    security_groups = [aws_security_group.profile-api-sg.id]
    assign_public_ip = true
  }

  load_balancer {
    container_name = var.service-name
    container_port = var.docker-container-port
    target_group_arn = aws_alb_target_group.profile-api-tg.arn
  }
}