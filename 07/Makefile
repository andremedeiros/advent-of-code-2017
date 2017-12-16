ragel = ragel

parser.go: parser.go.rl
	$(ragel) -Z -T0 -o $@ $<

clean:
	rm -rf parser.go
