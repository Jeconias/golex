COMMAND := ""
PARAMS 	:= ""

run:
	@docker run -it golex go ${COMMAND} ${PARAMS}
build:
	@docker build -t golex .
