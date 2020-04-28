# Copyright Greg Haskins All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
FROM hyperledger/fabric-baseimage:amd64-0.4.18
RUN npm install -g npm@latest
COPY payload/chaintool payload/protoc-gen-go /usr/local/bin/
ADD payload/goshim.tar.bz2 $GOPATH/src/
RUN mkdir -p /chaincode/input /chaincode/output
LABEL org.hyperledger.fabric.version=1.4.7 \
      org.hyperledger.fabric.base.version=0.4.18
