#!/bin/bash

# Copyright Envoy Gateway Authors
# SPDX-License-Identifier: Apache-2.0
# The full text of the Apache license is available in the LICENSE file at
# the root of the repo.


cd "${SRC}"/gateway
go mod tidy

# compile native-format fuzzers
# https://github.com/google/oss-fuzz/blob/master/infra/base-images/base-builder/compile_native_go_fuzzer
compile_native_go_fuzzer github.com/envoyproxy/gateway/test/fuzz FuzzGatewayAPIToXDS fuzz_gateway_api_to_xds

