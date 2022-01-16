FROM localstack/localstack:0.13.3

RUN apt update -y && apt install jq -y
