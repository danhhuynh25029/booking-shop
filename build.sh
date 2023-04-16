function start_local(){
  go run ./cmd/service/server.go
}
if [ $1="local" ]
then
  start_local
fi