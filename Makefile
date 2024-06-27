TARGET = encoder
OUTDIR = bin
LOGFILE = ${TARGET}.log

all: clean build

build:
	go build -o ${OUTDIR}/${TARGET} main.go

test: clean build
	clear
	go test ./...

clean:
	rm -f ${OUTDIR}/${TARGET}
	rm -f ${OUTDIR}/${LOGFILE}
	clear
