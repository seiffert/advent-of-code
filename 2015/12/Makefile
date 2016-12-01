build:
	go build

test: build
	echo "[1,2,3]" | ./12 
	echo "[[[3]]]" | ./12 
	echo '{"a":[-1,1]}' | ./12 
	echo '[]' | ./12 
	echo '[1,{"c":"red","b":2},3]' | ./12
	echo '{"d":"red","e":[1,2,3,4],"f":5}' | ./12

run: build
	cat input | ./12
