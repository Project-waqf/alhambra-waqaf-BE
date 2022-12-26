FROM golang:1.19

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

