TARGET = encoder
OUTDIR = bin
LOGFILE = ${TARGET}.log
TESTFILE = "Big Buck Bunny (2014) Orig.mp4"

all: clean build

build:
	go build -o ${OUTDIR}/${TARGET} \
		main.go \
		mediainfo.go \
		logger.go \
		mediafile.go

test: clean build
	clear
	./${OUTDIR}/${TARGET} ${TESTFILE}

clean:
	rm -f ${OUTDIR}/${TARGET}
	rm -f ${OUTDIR}/${LOGFILE}
	clear
