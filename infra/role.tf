resource "aws_iam_role" "profile-api-role" {
  name               = "${var.service-name}-iam-role"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow",
        Principal = {
          Service = [
            "ecs.amazonaws.com",
            "ecs-tasks.amazonaws.com",
            "ssm.amazonaws.com"
          ]
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy" "price-scraper-policy" {
  name   = "${var.service-name}-iam-policy"
  role   = aws_iam_role.profile-api-role.id
  policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "ecs:*",
          "ecr:*",
          "ssm:*",
          "cloudwatch:*",
          "logs:*",
          "elasticloadbalancing:*",
          "kms:*",
          "secretsmanager:*"
        ],
        Resource = "*"
      }
    ]
  })
}