SERVER1_PATH=frontend/cmd/web
SERVER2_PATH=backend-service/cmd/api
SERVER3_PATH=authentication-service/cmd/api


start_server1:
	cd $(SERVER1_PATH) && go run main.go handler.go

start_server2:
	cd $(SERVER2_PATH) && go run main.go handlerfunctions.go

start_server3:
	cd $(SERVER3_PATH) && go run main.go routes.go databaseFunctionalities.go


start_all_servers:
	make start_server1 &
	make start_server2 &
	make start_server3 &
	wait
