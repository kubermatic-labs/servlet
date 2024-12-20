# Copyright 2024 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM docker.io/golang:1.23.4 as builder

# import the GOPROXY variable via an arg and then use
# that arg to define the environment variable later on
ARG GOPROXY=
ENV GOPROXY=$GOPROXY

WORKDIR /go/src/k8c.io/servlet
COPY . .
RUN make clean servlet

FROM gcr.io/distroless/static-debian12:debug
LABEL maintainer="support@kubermatic.com"

COPY --from=builder /go/src/k8c.io/servlet/_build/servlet /usr/local/bin/servlet

USER nobody
ENTRYPOINT [ "servlet" ]
