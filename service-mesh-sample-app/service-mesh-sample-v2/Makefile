NAME = service-mesh-sample:2.0.1

default: run

image:
	docker build -t $(NAME) .

push:
	docker tag $(NAME) soumiyajit/$(NAME)
	docker push soumiyajit/$(NAME)

run: 
	go run main.go
