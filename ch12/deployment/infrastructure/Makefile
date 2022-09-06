CONN = $(shell terraform output -raw db_conn)

GIT_SHA = $(shell git rev-parse HEAD)
ECR_URL = $(shell terraform output -raw ecr_url)

# Use this command to easily run all of the infrastructure deployment steps
# There are costs associated with this demo; Remember to tear everything down!
deploy: init validate plan apply
	$(MAKE) setup-db
	$(MAKE) push-services

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
	terraform destroy

# Use this command to populate ~/.kube/config so that you can use k9s to browse the cluster
kubeconfig:
	aws eks update-kubeconfig --region $(shell terraform output region) --name $(shell terraform output project)


setup-db:
	@echo Initializing the Mallbots service DBs
	@echo ...
	@psql --file="./sql/create_commondb.sql" ${CONN}/postgres
	@psql --file="./sql/init_commondb.sql" ${CONN}/commondb
	@psql --file="./sql/create_services.sql" ${CONN}/postgres

	@psql --file="./sql/init_baskets.sql" ${CONN}/baskets
	@psql --file="./sql/init_cosec.sql" ${CONN}/cosec
	@psql --file="./sql/init_customers.sql" ${CONN}/customers
	@psql --file="./sql/init_depot.sql" ${CONN}/depot
	@psql --file="./sql/init_notifications.sql" ${CONN}/notifications
	@psql --file="./sql/init_ordering.sql" ${CONN}/ordering
	@psql --file="./sql/init_payments.sql" ${CONN}/payments
	@psql --file="./sql/init_search.sql" ${CONN}/search
	@psql --file="./sql/init_stores.sql" ${CONN}/stores

push-services: baskets cosec customers depot notifications ordering payments search stores

baskets:
	docker build -t ${ECR_URL}/baskets:${GIT_SHA} -t ${ECR_URL}/baskets:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=baskets ../..
	docker push -a ${ECR_URL}/baskets

cosec:
	docker build -t ${ECR_URL}/cosec:${GIT_SHA} -t ${ECR_URL}/cosec:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=cosec ../..
	docker push -a ${ECR_URL}/cosec

customers:
	docker build -t ${ECR_URL}/customers:${GIT_SHA} -t ${ECR_URL}/customers:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=customers ../..
	docker push -a ${ECR_URL}/customers

depot:
	docker build -t ${ECR_URL}/depot:${GIT_SHA} -t ${ECR_URL}/depot:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=depot ../..
	docker push -a ${ECR_URL}/depot

notifications:
	docker build -t ${ECR_URL}/notifications:${GIT_SHA} -t ${ECR_URL}/notifications:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=notifications ../..
	docker push -a ${ECR_URL}/notifications

ordering:
	docker build -t ${ECR_URL}/ordering:${GIT_SHA} -t ${ECR_URL}/ordering:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=ordering ../..
	docker push -a ${ECR_URL}/ordering

payments:
	docker build -t ${ECR_URL}/payments:${GIT_SHA} -t ${ECR_URL}/payments:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=payments ../..
	docker push -a ${ECR_URL}/payments

search:
	docker build -t ${ECR_URL}/search:${GIT_SHA} -t ${ECR_URL}/search:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=search ../..
	docker push -a ${ECR_URL}/search

stores:
	docker build -t ${ECR_URL}/stores:${GIT_SHA} -t ${ECR_URL}/stores:latest --file "../../docker/Dockerfile.microservices" --build-arg=service=stores ../..
	docker push -a ${ECR_URL}/stores
