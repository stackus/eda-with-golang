ECR_URL = $(shell terraform output -raw ecr_url)
REGION = $(shell terraform output -raw region)

ready: init validate

# Use this command to easily run all of the infrastructure deployment steps
# There are costs associated with this demo; Remember to tear everything down!
deploy: plan apply
	$(MAKE) kubeconfig
	$(MAKE) push-services
	@echo !!!
	@echo !!!
	@echo !!! There are cost associated with running this demo! Remember to tear everything down!
	@echo !!!
	@echo !!! Use 'make destroy' OR use 'terraform destroy' to tear everything down.
	@echo !!!
	@echo !!!

init:
	terraform init -upgrade

validate:
	terraform validate

plan:
	terraform plan -out infra.tfplan

apply:
	terraform apply infra.tfplan
	@echo !!!
	@echo !!!
	@echo !!! There are cost associated with running this demo! Remember to tear everything down!
	@echo !!!
	@echo !!! Use 'make destroy' OR use 'terraform destroy' to tear everything down.
	@echo !!!
	@echo !!!

destroy:
	terraform destroy -auto-approve

# Use this command to populate ~/.kube/config so that you can use k9s to browse the cluster
kubeconfig:
	aws eks update-kubeconfig --region ${REGION} --name $(shell terraform output project)

push-services: ecr-login baskets cosec customers depot notifications ordering payments search stores

ecr-login:
	@aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR_URL}

baskets:
	@docker build -t ${ECR_URL}/baskets:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=baskets ../..
	@docker push -a ${ECR_URL}/baskets

cosec:
	@docker build -t ${ECR_URL}/cosec:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=cosec ../..
	@docker push -a ${ECR_URL}/cosec

customers:
	@docker build -t ${ECR_URL}/customers:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=customers ../..
	@docker push -a ${ECR_URL}/customers

depot:
	@docker build -t ${ECR_URL}/depot:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=depot ../..
	@docker push -a ${ECR_URL}/depot

notifications:
	@docker build -t ${ECR_URL}/notifications:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=notifications ../..
	@docker push -a ${ECR_URL}/notifications

ordering:
	@docker build -t ${ECR_URL}/ordering:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=ordering ../..
	@docker push -a ${ECR_URL}/ordering

payments:
	@docker build -t ${ECR_URL}/payments:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=payments ../..
	@docker push -a ${ECR_URL}/payments

search:
	@docker build -t ${ECR_URL}/search:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=search ../..
	@docker push -a ${ECR_URL}/search

stores:
	@docker build -t ${ECR_URL}/stores:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=stores ../..
	@docker push -a ${ECR_URL}/stores
