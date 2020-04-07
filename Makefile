all:
	rm -rf demo/docs/*
	rm -rf demo/docs/html/*
	go run . -o demo/markdown demo/simple2.sysl
	go run . --type=html -o demo/html demo/simple2.sysl
	cp -r demo/html/* docs/
	printf '%s\n%s\n' "<h1>This is an example of serving specs from sysl catalog</h1><a href=\"http://github.com/anz-bank/sysl-catalog\"></a>" "$$(cat docs/index.html)" > docs/index.html