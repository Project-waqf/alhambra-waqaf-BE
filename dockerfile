FROM golang:1.19

# SET TIMEZONE
RUN apt-get update && \
    apt-get install -yq tzdata && \
    ln -fs /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata

ENV TZ="Asia/Jakarta"

# MEMBUAT FOLDER APP
RUN mkdir /app

# SET DIREKTORI APP
WORKDIR /app

# COPY FILE KE FOLDER APP
ADD . .

# BUAT FILE EXE
RUN go build -o main

#RUN EXE
CMD [ "./main" ]

