TARGET = encoder
LOGFILE = ${TARGET}.log
TESTFILE = example1.mp4

all: clean build

build:
	go build -o bin/${TARGET} main.go mediainfo.go logger.go mediafile.go

test: clean build
	./bin/${TARGET} ${TESTFILE}

clean:
	rm -f ${TARGET}
	rm -f ${LOGFILE}
	clear
