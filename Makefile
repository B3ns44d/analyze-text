rm:
	docker rm analyze_text
run:
	docker run --env-file .env --name analyze_text -p 3008:3008 analyze_text:latest

build:
	docker build -t analyze_text:latest .