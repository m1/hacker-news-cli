default: build

build:
	docker build -t hn .
	docker create --name hn hn:latest -f
	docker cp hn:/hn hacker-news-cli
	docker rm -f hn

run-container:
	docker run -d -p 8282:80 --rm -it gousto