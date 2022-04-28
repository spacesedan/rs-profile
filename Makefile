SERVICE_NAME := "profile-api"
SERVICE_TAG := "v1"
ECR_REPO_URL := "871631304903.dkr.ecr.us-east-1.amazonaws.com/$(SERVICE_NAME)"

run:
	go run cmd/*.go

create-repo: ## Create repo in ECR
	aws ecr get-login-password --region us-east-1 --profile rarityshark | docker login --username AWS --password-stdin 871631304903.dkr.ecr.us-east-1.amazonaws.com  \
	&& aws ecr create-repository --repository-name $(SERVICE_NAME) --profile rarityshark || true \

dev:
	docker build -t spacecoupe/$(SERVICE_NAME):$(SERVICE_TAG) -f infra/Dockerfile .; \
	docker push spacecoupe/$(SERVICE_NAME):$(SERVICE_TAG)

docker:
	docker build -t $(SERVICE_NAME):$(SERVICE_TAG) -f infra/Dockerfile .

dockerize: docker
	aws ecr get-login-password --region us-east-1 --profile rarityshark | docker login --username AWS --password-stdin 871631304903.dkr.ecr.us-east-1.amazonaws.com  \
	&& docker tag $(SERVICE_NAME):$(SERVICE_TAG) $(ECR_REPO_URL):$(SERVICE_TAG) \
	&& docker push $(ECR_REPO_URL):$(SERVICE_TAG)

init: ## Init terraform
	cd infra; terraform init

plan : ## Plan Terraform resources
	cd infra; terraform plan -var-file="production.tfvars" -var "docker-image-url=$(ECR_REPO_URL):$(SERVICE_TAG)"

deploy: init docker ## Build docker image and push to ecr
	cd infra; \
    terraform apply -var-file="production.tfvars" -var "docker-image-url=$(ECR_REPO_URL):$(SERVICE_TAG)" --auto-approve

destroy: init
	cd infra;  \
	terraform destroy -var-file="production.tfvars" -var "docker-image-url=$(ECR_REPO_URL):$(SERVICE_TAG)" --auto-approve

redeploy: destroy deploy