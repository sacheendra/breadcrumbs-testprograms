go build -a -x -work 2.\ loopy.go 1> transcript.txt 2>&1

# Get importcfg from transcript.txt
GOSSAFUNC=loopy ../goroot/pkg/tool/linux_amd64/compile -importcfg /tmp/go-build4212046273/b008/importcfg 2.\ loopy.go

dlv exec --listen=:2345 --headless --api-version=2 ../goroot/pkg/tool/linux_amd64/compile -- -importcfg /tmp/go-build3820912063/b008/importcfg -l  8.\ ptrparamtest.go
