FROM golang:1.22-bookworm AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

RUN make worker

FROM continuumio/miniconda3:24.1.2-0 AS app

# prepare environment
RUN apt update -y && apt upgrade -y
RUN apt install -y \
    libgl1-mesa-glx \
    curl \
    wget \
    make \
    libssl-dev \
    libffi-dev \
    libopenblas-dev \
    git

# install vapoursynth
RUN conda install conda-forge::vapoursynth=69 -y

# install vapoursynth C++ plugins
RUN conda install tongyuantongyu::vapoursynth-bestsource=5 -y
RUN conda install tongyuantongyu::vapoursynth-fmtconv=30 -y

# install vapoursynth python plugins
RUN conda install tongyuantongyu::vapoursynth-mvsfunc=10.10 -y
RUN conda install tongyuantongyu::vapoursynth-vsutil=0.8.0 -y
RUN pip install git+https://github.com/HomeOfVapourSynthEvolution/havsfunc.git

# install python packages
RUN conda install conda-forge::numpy=1.26.4 -y
RUN conda install fastai::opencv-python-headless=4.10.0.82 -y

# worker app
WORKDIR /app

ENV TZ=Asia/Shanghai

ENV FINALRIP_EASYTIER_HOST 10.126.126.251

COPY --from=builder /build/worker/worker /app/
COPY --from=builder /build/conf/finalrip.yml /app/conf/

CMD ["/app/worker", "encode"]
