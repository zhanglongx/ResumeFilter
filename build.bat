set OS=windows
set ARCH=amd64

go build -ldflags -H=windowsgui -o resumefilter_%OS%_%ARCH%.exe . 
