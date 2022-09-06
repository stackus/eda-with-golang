GIT_SHA := $(shell git rev-parse HEAD)
ECR_URL := $(shell terraform output ecr_url)

# Use this command to easily run all of the infrastructure deployment steps
# There are costs associated with this demo; Remember to tear everything down!
deploy: init validate plan apply

init:
	terraform init -upgrade

validate:
	terraform validate

plan:
	terraform plan -out app.tfplan

apply:
	terraform apply app.tfplan
	@echo !!!
	@echo !!!
	@echo !!! There are cost associated with running this demo! Remember to tear everything down!
	@echo !!!
	@echo !!! Use 'make destroy' OR use 'terraform destroy' to tear everything down.
	@echo !!!
	@echo !!!

destroy:
	terraform destroy

baskets:
	docker build -t ${ECR_URL}/baskets:${GIT_SHA} -t ${ECR_URL}/baskets:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=baskets ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/baskets

cosec:
	docker build -t ${ECR_URL}/cosec:${GIT_SHA} -t ${ECR_URL}/cosec:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=cosec ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/cosec

customers:
	docker build -t ${ECR_URL}/customers:${GIT_SHA} -t ${ECR_URL}/customers:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=customers ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/customers

depot:
	docker build -t ${ECR_URL}/depot:${GIT_SHA} -t ${ECR_URL}/depot:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=depot ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/depot

notifications:
	docker build -t ${ECR_URL}/notifications:${GIT_SHA} -t ${ECR_URL}/notifications:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=notifications ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/notifications

ordering:
	docker build -t ${ECR_URL}/ordering:${GIT_SHA} -t ${ECR_URL}/ordering:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=ordering ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/ordering

payments:
	docker build -t ${ECR_URL}/payments:${GIT_SHA} -t ${ECR_URL}/payments:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=payments ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/payments

search:
	docker build -t ${ECR_URL}/search:${GIT_SHA} -t ${ECR_URL}/search:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=search ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/search

stores:
	docker build -t ${ECR_URL}/stores:${GIT_SHA} -t ${ECR_URL}/stores:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=stores ../..
	docker push -a 357852354913.dkr.ecr.us-east-1.amazonaws.com/stores
